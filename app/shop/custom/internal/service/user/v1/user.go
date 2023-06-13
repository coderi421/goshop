package v1

import (
	"context"
	"fmt"
	itime "github.com/coderi421/gframework/pkg/common/time"
	"github.com/coderi421/gframework/pkg/log"
	"github.com/coderi421/gframework/pkg/storage"
	"time"

	"github.com/coderi421/gframework/gmicro/server/restserver/middlewares"
	"github.com/coderi421/gframework/pkg/errors"
	"github.com/coderi421/goshop/app/pkg/code"
	"github.com/coderi421/goshop/app/pkg/options"
	"github.com/coderi421/goshop/app/shop/custom/internal/data"
	"github.com/dgrijalva/jwt-go"
)

type UserDTO struct {
	ID        int32     `json:"id"`
	Mobile    string    `json:"mobile"`
	NickName  string    `json:"nick_name"`
	Birthday  time.Time `gorm:"type:datetime"`
	Gender    string    `json:"gender"`
	Role      int32     `json:"role"`
	PassWord  string    `json:"password"`
	Token     string    `json:"token"`
	ExpiresAt int64     `json:"expires_at"`
}

type UserSrv interface {
	MobileLogin(ctx context.Context, mobile, password string) (*UserDTO, error)
	Register(ctx context.Context, userDTO *UserDTO, codes string) (*UserDTO, error)
	Update(ctx context.Context, userDTO *UserDTO) error
	Get(ctx context.Context, userID int32) (*UserDTO, error)
	GetByMobile(ctx context.Context, mobile string) (*UserDTO, error)
	CheckPassWord(ctx context.Context, password, EncryptedPassword string) (bool, error)
}

type userService struct {
	//ud data.UserData
	data data.DataFactory

	jwtOpts *options.JwtOptions
}

func NewUser(data data.DataFactory, jwtOpts *options.JwtOptions) UserSrv {
	return &userService{data: data, jwtOpts: jwtOpts}
}

// MobileLogin
//
//	@Description:
//	@receiver us
//	@param ctx
//	@param mobile
//	@param password
//	@return *UserDTO
//	@return error
func (us *userService) MobileLogin(ctx context.Context, mobile, password string) (*UserDTO, error) {
	user, err := us.data.User().GetByMobile(ctx, mobile)
	if err != nil {
		return nil, err
	}
	//检查密码是否正确
	ok, err := us.data.User().CheckPassWord(ctx, password, user.PassWord)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.WithCode(code.ErrUserPasswordIncorrect, "用户名或密码错误")
	}

	//生成token
	j := middlewares.NewJWT(us.jwtOpts.Key)
	claims := middlewares.CustomClaims{
		//key需要和jwt的key一致
		//这里json能起作用是因为在j.CreateToken里NewWithClaims的时候底层claims会解析成json
		ID:          uint(user.ID),
		NickName:    user.NickName,
		AuthorityId: uint(user.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                                   //签名的生效时间
			ExpiresAt: (time.Now().Local().Add(us.jwtOpts.Timeout)).Unix(), //*天过期
			Issuer:    us.jwtOpts.Realm,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		return nil, err
	}
	//_ = copier.Copy(&resp, user)
	return &UserDTO{
		ID:        user.ID,
		Mobile:    user.Mobile,
		NickName:  user.NickName,
		Birthday:  user.Birthday,
		Gender:    user.Gender,
		Role:      user.Role,
		PassWord:  user.PassWord,
		Token:     token,
		ExpiresAt: (time.Now().Local().Add(us.jwtOpts.Timeout)).Unix(),
	}, nil
}

// Register
//
//	@Description:
//	@receiver us
//	@param ctx
//	@param userDTO
//	@param codes
//	@return *UserDTO
//	@return error
func (us *userService) Register(ctx context.Context, userDTO *UserDTO, codes string) (*UserDTO, error) {
	//验证码校验
	rstore := storage.RedisCluster{}
	value, err := rstore.GetKey(ctx, fmt.Sprintf("%s_%d", userDTO.Mobile, 1))
	if err != nil {
		return nil, errors.WithCode(code.ErrCodeNotExist, "验证码不存在")
	}
	if value != codes {
		return nil, errors.WithCode(code.ErrCodeInCorrect, "验证码不匹配")
	}
	var userDO = &data.User{
		Mobile:   userDTO.Mobile,
		NickName: userDTO.NickName,
		//proto 中定义的不完全
		//Birthday: userDTO.Birthday,
		//Gender:   userDTO.Gender,
		PassWord: userDTO.PassWord,
	}
	userinfo, err := us.data.User().Create(ctx, userDO)
	if err != nil {
		log.Errorf("user register failed: %v", err)
		return nil, err
	}
	//生成token
	j := middlewares.NewJWT(us.jwtOpts.Key)
	claims := middlewares.CustomClaims{
		ID:          uint(userinfo.Id),
		NickName:    userinfo.NickName,
		AuthorityId: uint(userinfo.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                                   //签名的生效时间
			ExpiresAt: (time.Now().Local().Add(us.jwtOpts.Timeout)).Unix(), //*天过期
			Issuer:    us.jwtOpts.Realm,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		return nil, err
	}
	return &UserDTO{
		ID:        userinfo.Id,
		Mobile:    userinfo.Mobile,
		NickName:  userinfo.NickName,
		Birthday:  itime.Time{time.Unix(int64(userinfo.BirthDay), 0)}.Time,
		Gender:    userinfo.Gender,
		Role:      userinfo.Role,
		PassWord:  userinfo.PassWord,
		Token:     token,
		ExpiresAt: (time.Now().Local().Add(us.jwtOpts.Timeout)).Unix(),
	}, nil
}

// Update
//
//	@Description:
//	@receiver us
//	@param ctx
//	@param userDTO
//	@return error
func (us *userService) Update(ctx context.Context, userDTO *UserDTO) error {
	user := &data.User{
		Mobile:   userDTO.Mobile,
		NickName: userDTO.NickName,
		Birthday: userDTO.Birthday,
		Gender:   userDTO.Gender,
	}
	return us.data.User().Update(ctx, user)
}

// Get
//
//	@Description:
//	@receiver us
//	@param ctx
//	@param userID
//	@return *UserDTO
//	@return error
func (us *userService) Get(ctx context.Context, userID int32) (*UserDTO, error) {
	userDO, err := us.data.User().Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &UserDTO{
		ID:       userDO.ID,
		Mobile:   userDO.Mobile,
		NickName: userDO.NickName,
		Birthday: userDO.Birthday,
		Gender:   userDO.Gender,
		Role:     userDO.Role,
		PassWord: userDO.PassWord,
	}, nil
}

// GetByMobile
//
//	@Description:
//	@receiver us
//	@param ctx
//	@param mobile
//	@return *UserDTO
//	@return error
func (us *userService) GetByMobile(ctx context.Context, mobile string) (*UserDTO, error) {
	userDO, err := us.data.User().GetByMobile(ctx, mobile)
	if err != nil {
		return nil, err
	}
	return &UserDTO{
		ID:       userDO.ID,
		Mobile:   userDO.Mobile,
		NickName: userDO.NickName,
		Birthday: userDO.Birthday,
		Gender:   userDO.Gender,
		Role:     userDO.Role,
		PassWord: userDO.PassWord,
	}, nil
}

func (us *userService) CheckPassWord(ctx context.Context, password, EncryptedPassword string) (bool, error) {
	return us.data.User().CheckPassWord(ctx, password, EncryptedPassword)
}
