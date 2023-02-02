package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hxx258456/pyramidel-chain-baas/internal/localconfig"
	"github.com/hxx258456/pyramidel-chain-baas/pkg/utils/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(mysqlConfig *localconfig.TopLevel) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		mysqlConfig.MySqlConfig.User,
		mysqlConfig.MySqlConfig.Password,
		mysqlConfig.MySqlConfig.Host,
		mysqlConfig.MySqlConfig.Port,
		mysqlConfig.MySqlConfig.DB,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		logger.Error("connect DB failed", zap.Error(err))
		return
	} else {
		logger.Info(">>>MySql连接成功")
	}
	db.SetMaxOpenConns(mysqlConfig.MySqlConfig.MaxOpenConns)
	db.SetMaxIdleConns(mysqlConfig.MySqlConfig.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}
