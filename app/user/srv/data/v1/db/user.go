package db

import (
	"context"
	"github.com/coderi421/gframework/app/pkg/code"
	"github.com/coderi421/gframework/pkg/errors"
	"github.com/coderi421/goshop/app/pkg/options"
	"gorm.io/gorm"

	metav1 "github.com/coderi421/gframework/pkg/common/meta/v1"
	dv1 "github.com/coderi421/goshop/app/user/srv/data/v1"
)

type user struct {
	db *gorm.DB
}

// DB 重新初始化db 方式污染db
func (u *user) DB() *gorm.DB {
	// 如果使用 db 的时候，初始化链式方法，然后拼接，最后调用 结束方法的时候，最好复制一份db 因为如果并发会出问题
	// 如果直接使用链式+结束方法，则不会出现问题 u.db.Where("goodsname = ?", i).Find(m)
	return u.db.Session(&gorm.Session{})
}

var _ dv1.UserStore = (*user)(nil)
var _ options.GORMUser = (*user)(nil)

func NewUser(db *gorm.DB) *user {
	return &user{db: db}
}

// List
//
//	@Description: 获取用户列表
//	@receiver u *user
//	@param ctx 上下文
//	@param orderby 排序 ["age desc", "name"] -> ORDER BY age desc, name
//	@param opts 分页
//	@return *dv1.UserDOList 用户列表
//	@return error 错误
func (u *user) List(ctx context.Context, orderby []string, opts metav1.ListMeta) (*dv1.UserDOList, error) {
	res := dv1.UserDOList{}

	// 分页
	var limit, offset int
	if opts.PageSize == 0 || opts.PageSize > 100 {
		limit = 10
	} else {
		limit = opts.PageSize
	}

	if opts.Page > 0 {
		offset = (opts.Page - 1) * limit
	}

	// 这里重新赋值了，所以不会影响到原来的db
	// 如果使用 db 的时候，初始化链式方法，然后拼接，最后调用 结束方法的时候，最好复制一份db 因为如果并发会出问题
	// 如果直接使用链式+结束方法，则不会出现问题 u.db.Where("goodsname = ?", i).Find(m)
	query := u.DB()
	// 使用db的DB方法，返回一个新的db
	//排序
	for _, value := range orderby {
		//坑：如果db改掉了？
		//u.db=u.db.Order(value)
		query = query.Order(value)
	}

	//查询 - 发起多个请求
	d := query.Offset(offset).Limit(limit).Find(&res.Items).Count(&res.TotalCount)
	if d.Error != nil {
		return nil, errors.WithCode(code.ErrDatabase, d.Error.Error())
	}
	return &res, nil
}

// GetByMobile
//
//	@Description: 通过手机号获取用户信息
//	@receiver u *user
//	@param ctx 上下文
//	@param mobile 手机号
//	@return *dv1.UserDO 用户信息
//	@return error 错误
func (u *user) GetByMobile(ctx context.Context, mobile string) (*dv1.UserDO, error) {
	//TODO implement me
	panic("implement me")
}

func (u *user) GetByID(ctx context.Context, id uint64) (*dv1.UserDO, error) {
	//TODO implement me
	panic("implement me")
}

func (u *user) Create(ctx context.Context, user *dv1.UserDO) error {
	//TODO implement me
	panic("implement me")
}

func (u *user) Update(ctx context.Context, user *dv1.UserDO) error {
	//TODO implement me
	panic("implement me")
}
