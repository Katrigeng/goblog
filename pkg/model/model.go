package model

import (
	"fmt"
	"goblog/pkg/config"
	"goblog/pkg/logger"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	// GORM 的 MySQL 数据库驱动导入
	"gorm.io/driver/mysql"
)

// DB gorm.DB 对象
var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {

	var err error

	gormConfig := mysql.New(mysql.Config{
		// DSN: "root:123456@tcp(192.168.240.182:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
		// DSN: "root:root@tcp(127.0.0.1:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
		DSN: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
			config.GetString("database.username"),
			config.GetString("database.password"),
			config.GetString("database.host"),
			config.GetString("database.port"),
			config.GetString("database.database"),
			config.GetString("database.charset")),
	})

	var level gormlogger.LogLevel
	if config.GetBool("app.debug") {
		level = gormlogger.Warn
	} else {
		level = gormlogger.Error
	}

	// 准备数据库连接池
	DB, err = gorm.Open(gormConfig, &gorm.Config{
		Logger: gormlogger.Default.LogMode(level),
	})

	logger.LogError(err)

	return DB
}
