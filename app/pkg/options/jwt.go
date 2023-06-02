package options

import (
	"fmt"
	"github.com/spf13/pflag"
	"time"
	"unicode/utf8"
)

type JwtOptions struct {
	// 底层使用的 github.com/appleboy/gin-jwt/v2 包中需要的参数
	Realm      string        `json:"realm"       mapstructure:"realm"`       // Realm name to display to the user.
	Key        string        `json:"key"         mapstructure:"key"`         // Private key used to sign jwt token.
	Timeout    time.Duration `json:"timeout"     mapstructure:"timeout"`     // JWT token timeout.
	MaxRefresh time.Duration `json:"max-refresh" mapstructure:"max-refresh"` // This field allows clients to refresh their token until MaxRefresh has passed.
}

func NewJwtOptions() *JwtOptions {
	return &JwtOptions{
		Realm:      "JWT",
		Key:        "Gd%YCfP1agNHo5x6xm2Qs33Bf!B#Gi!o",
		Timeout:    24 * time.Hour,
		MaxRefresh: 7 * 24 * time.Hour,
	}
}

func (s *JwtOptions) Validate() []error {
	var errs []error
	strLength := utf8.RuneCountInString(s.Key)

	if strLength < 6 || strLength > 32 {
		errs = append(errs, fmt.Errorf("--secret-key must larger than 5 and little than 33"))
	}

	return errs
}

func (s *JwtOptions) AddFlags(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}

	fs.StringVar(&s.Realm, "jwt.realm", s.Realm, "Realm name to display to the user.")
	fs.StringVar(&s.Key, "jwt.key", s.Key, "Private key used to sign jwt token.")
	fs.DurationVar(&s.Timeout, "jwt.timeout", s.Timeout, "JWT token timeout.")

	fs.DurationVar(&s.MaxRefresh, "jwt.max-refresh", s.MaxRefresh, ""+
		"This field allows clients to refresh their token until MaxRefresh has passed.")
}
