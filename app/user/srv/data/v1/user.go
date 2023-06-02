package v1

import (
	"context"
	"gorm.io/gorm"
	"time"

	metav1 "github.com/coderi421/gframework/pkg/common/meta/v1"
)

type BaseModel struct {
	ID        int32          `gorm:"primary_key;comment:ID"`
	CreatedAt time.Time      `gorm:"column:add_time;comment:创建时间"`
	UpdatedAt time.Time      `gorm:"column:update_time;comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"comment:删除时间"`
	IsDeleted bool           `gorm:"comment:是否删除"`
}

type UserDO struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null;comment:手机号"`
	Password string     `gorm:"type:varchar(100);not null;comment:密码"`
	NickName string     `gorm:"type:varchar(20);comment:账号名称"`
	Birthday *time.Time `gorm:"type:datetime;comment:出生日期"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6);comment:femail表示女,male表示男"`
	Role     int        `gorm:"column:role;default:1;type:int;comment:1表示普通用户,2表示管理员"`
}

func (u *UserDO) TableName() string {
	return "user"
}

type UserDOList struct {
	TotalCount int64     `json:"totalCount,omitempty"`
	Items      []*UserDO `json:"data"`
}

type UserStore interface {
	/*
		有数据访问的方法，一定要有error
		参数中最好有ctx 可能需要cancel / telemetry等
	*/
	//用户列表 - 后台管理系统
	List(ctx context.Context, orderby []string, opts metav1.ListMeta) (*UserDOList, error)

	//通过手机号码查询用户
	GetByMobile(ctx context.Context, mobile string) (*UserDO, error)

	//通过用户ID查询用户
	GetByID(ctx context.Context, id uint64) (*UserDO, error)

	//创建用户
	Create(ctx context.Context, user *UserDO) error

	//更新用户
	Update(ctx context.Context, user *UserDO) error
}
