package v1

import (
	"context"
	"github.com/coderi421/goshop/app/shop/custom/internal/data"
)

type UserDTO struct {
	data.User

	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}
type UserSrv interface {
	MobileLogin(ctx context.Context, mobile, password string) (*UserDTO, error)
	Register(ctx context.Context, mobile, password string, codes string) (*UserDTO, error)
	Updata(ctx context.Context, userDTO *UserDTO) error
	Get(ctx context.Context, userID uint64) (*UserDTO, error)
	GetByMobile(ctx context.Context, mobile string) (*UserDTO, error)
	CheckPassWord(ctx context.Context, password, EncryptedPassword string) (bool, error)
}
