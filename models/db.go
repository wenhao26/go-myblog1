package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"blog/config"
)

var db *gorm.DB
var err error

func InitDb() {
	dns := fmt.Sprintf("%s:`%s`@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPassWord,
		config.DbHost,
		config.DbPort,
		config.DbName)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// 日志模式
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,

		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名
			SingularTable: true,
		},
	})
	if err != nil {
		log.Panic("连接MySQL数据库失败：", err)
	}

	sqlDB, _ := db.DB()
	// 设置连接池最大闲置连接数
	sqlDB.SetMaxIdleConns(10)
	// 设置数据库最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置连接的最大可复用时间
	sqlDB.SetConnMaxLifetime(10e9)
}
