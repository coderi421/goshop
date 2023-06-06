package user

import (
	"context"
	"crypto/sha512"
	"strings"

	"github.com/anaskhan96/go-password-encoder"
	upbv1 "github.com/coderi421/goshop/api/user/v1"
)

func (us *userServer) CheckPassWord(ctx context.Context, info *upbv1.PasswordCheckInfo) (*upbv1.CheckResponse, error) {
	//校验密码
	options := &password.Options{16, 100, 32, sha512.New}
	passwordInfo := strings.Split(info.EncryptedPassword, "$")
	check := password.Verify(info.Password, passwordInfo[2], passwordInfo[3], options)
	return &upbv1.CheckResponse{Success: check}, nil
}
