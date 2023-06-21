package db

import (
	"fmt"
	v1 "github.com/coderi421/goshop/app/user/srv/data/v1"
	"sync"

	"github.com/coderi421/gframework/pkg/errors"
	"github.com/coderi421/goshop/app/pkg/code"
	"github.com/coderi421/goshop/app/pkg/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbFactory v1.DataFactory
	once      sync.Once
)

type mysqlFactory struct {
	db *gorm.DB
}

func (mf *mysqlFactory) User() v1.UserStore {
	return newUser(mf)
}

var _ v1.DataFactory = &mysqlFactory{}

func GetDBFactoryOr(mysqlOpts *options.MySQLOptions) (v1.DataFactory, error) {
	if mysqlOpts == nil && dbFactory == nil {
		return nil, errors.WithCode(code.ErrConnectDB, "failed to get mysql store factory")

	}
	var err error
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mysqlOpts.Username,
			mysqlOpts.Password,
			mysqlOpts.Host,
			mysqlOpts.Port,
			mysqlOpts.Database)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}

		dbFactory = &mysqlFactory{
			db: db,
		}
		sqlDB, _ := db.DB()
		//允许连接多少个mysql
		sqlDB.SetMaxOpenConns(mysqlOpts.MaxOpenConnections)
		//允许最大的空闲的连接数
		sqlDB.SetMaxIdleConns(mysqlOpts.MaxIdleConnections)
		//重用连接的最大时长
		sqlDB.SetConnMaxLifetime(mysqlOpts.MaxConnectionLifetime)

	})
	if dbFactory == nil || err != nil {
		return nil, errors.WithCode(code.ErrConnectDB, "failed to get mysql store factory")
	}
	return dbFactory, nil
}
