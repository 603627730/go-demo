package config

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlConn *gorm.DB

func init() {
	dsn := "root:lilishop@tcp(127.0.0.1:3306)/lilishop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	MysqlConn = db
}
