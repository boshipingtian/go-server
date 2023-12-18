package core

import (
	"go-server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func Gorm() {
	global.DB = MysqlConnector()
}

func MysqlConnector() *gorm.DB {
	if global.Config.Datasource.Host == "" {
		return nil
	}
	dsn := global.Config.Datasource.GenerateUrl()

	var mysqlLogger logger.Interface
	if global.Config.System.LogLevel == "dev" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最大可容纳连接数
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 最大可复用时间
	return db
}
