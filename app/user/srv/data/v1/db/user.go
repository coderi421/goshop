package db

import (
	"context"

	"gorm.io/gorm"

	metav1 "github.com/coderi421/gframework/pkg/common/meta/v1"
	dv1 "github.com/coderi421/goshop/app/user/srv/data/v1"
)

type user struct {
	db *gorm.DB
}

var _ dv1.UserStore = (*user)(nil)

func NewUser(db *gorm.DB) *user {
	return &user{db: db}
}

func (u *user) List(ctx context.Context, opts metav1.ListMeta) (*dv1.UserDOList, error) {
	//实现gorm查询
	return nil, nil
}
