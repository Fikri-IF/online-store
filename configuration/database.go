package configuration

import (
	"database/sql"
	"fmt"
	"online-store-golang/errs"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase(config Config) *sql.DB {
	username := config.Get("DB_USERNAME")
	password := config.Get("DB_PASSWORD")
	host := config.Get("DB_HOST")
	dbName := config.Get("DB_NAME")
	port := config.Get("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err := sql.Open("mysql", dsn)
	errs.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

}
