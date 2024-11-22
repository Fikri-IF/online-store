package configuration

import (
	"fmt"
	"online-store-golang/helper"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(config Config) *gorm.DB {
	username := config.Get("DB_USERNAME")
	password := config.Get("DB_PASSWORD")
	host := config.Get("DB_HOST")
	dbName := config.Get("DB_NAME")
	port := config.Get("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	sql, err := db.DB()
	helper.PanicIfError(err)

	sql.SetMaxIdleConns(5)
	sql.SetMaxOpenConns(20)
	sql.SetConnMaxLifetime(60 * time.Minute)
	sql.SetConnMaxIdleTime(10 * time.Minute)

	return db

}
