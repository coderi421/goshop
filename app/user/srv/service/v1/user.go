package v1

import (
	"context"
	"github.com/coderi421/gframework/pkg/errors"
	"github.com/coderi421/goshop/app/pkg/code"

	metav1 "github.com/coderi421/gframework/pkg/common/meta/v1"
	dv1 "github.com/coderi421/goshop/app/user/srv/data/v1"
)

type UserDTO struct {
	//不应该这样做，目前这样做是因为跟底层的逻辑一致了。减少冗余而已
	dv1.UserDO
}

type UserDTOList struct {
	TotalCount int64      `json:"totalCount,omitempty"`
	Items      []*UserDTO `json:"data"`
}

type UserSrv interface {
	List(ctx context.Context, orderby []string, opts metav1.ListMeta) (*UserDTOList, error)
	Create(ctx context.Context, user *UserDTO) error
	Update(ctx context.Context, user *UserDTO) error
	GetByID(ctx context.Context, ID uint64) (*UserDTO, error)
	GetByMobile(ctx context.Context, mobile string) (*UserDTO, error)
}

type userService struct {
	userStore dv1.DataFactory
}

var _ UserSrv = (*userService)(nil)

//	func NewUserService(userStore dv1.DataFactory) *userService {
//		return &userService{userStore: userStore}
//	}
func newUserService(store dv1.DataFactory) *userService {
	return &userService{
		userStore: store,
	}
}

func (u *userService) List(ctx context.Context, orderby []string, opts metav1.ListMeta) (*UserDTOList, error) {
	//这里是业务逻辑1

	/*
		1.data层的接口必须先写好
		2.我期望测试的时候每次测试底层的data层的数据按照我期望的返回
		 1.先手动插入一些数据
		 2.去删除一些数据
		 3.如果data层的方法有bug，坑爹 我们的代码想要具备好的可测试性
	*/
	doList, err := u.userStore.User().List(ctx, orderby, opts)
	if err != nil {
		return nil, err
	}
	//业务逻辑2
	//代码不方便写单元测试用例
	var userDTOList UserDTOList
	for _, value := range doList.Items {
		projectDTO := UserDTO{*value}
		userDTOList.Items = append(userDTOList.Items, &projectDTO)
	}
	//业务逻辑3
	return &userDTOList, nil
}

func (u *userService) Create(ctx context.Context, user *UserDTO) error {
	//一个手机号码只能创建一个用户，且只能通过手机号码创建。判断用户是否存在
	userInfo, err := u.userStore.User().GetByMobile(ctx, user.Mobile)
	//只有手机号不存在的情况下才能注册
	if err != nil {
		if errors.IsCode(err, code.ErrUserNotFound) {
			return u.userStore.User().Create(ctx, &user.UserDO)
		}
		return err
	}

	if userInfo != nil {
		return errors.WithCode(code.ErrUserAlreadyExists, "用户已经存在")
	}
	return errors.WithCode(code.ErrCodeInvalidParam, "参数错误")

	//return u.userStore.Create(ctx, &user.UserDO)
}

func (u *userService) Update(ctx context.Context, user *UserDTO) error {
	//先查询用户是否存在 其实可以不用查询 update一般不会报错
	_, err := u.userStore.User().GetByID(ctx, uint64(user.ID))
	if err != nil {
		return err
	}
	return u.userStore.User().Update(ctx, &user.UserDO)
}

func (u *userService) GetByID(ctx context.Context, ID uint64) (*UserDTO, error) {
	userDO, err := u.userStore.User().GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	return &UserDTO{*userDO}, nil
}

func (u *userService) GetByMobile(ctx context.Context, mobile string) (*UserDTO, error) {
	userDO, err := u.userStore.User().GetByMobile(ctx, mobile)
	if err != nil {
		return nil, err
	}
	return &UserDTO{*userDO}, nil
}
