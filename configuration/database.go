package configuration

import (
	"database/sql"
	"fmt"
	"log"
	"online-store-golang/errs"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// var db *sql.DB
var err error

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
func HandleRequiredTables(db *sql.DB) {
	const (
		createTableUserQuery = `
			CREATE TABLE IF NOT EXISTS
				user
					(
						user_id int NOT NULL AUTO_INCREMENT,
						username varchar(255) NOT NULL,
						password varchar(255) NOT NULL,
						created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
						updated_at timestamp NULL DEFAULT NULL,
						PRIMARY KEY (user_id),
						UNIQUE KEY users_un (username)
					)
		`
		createTableCategoryQuery = `
			CREATE TABLE IF NOT EXISTS
				category 
						(
							category_id int NOT NULL AUTO_INCREMENT,
							name varchar(255) NOT NULL,
							PRIMARY KEY (category_id),
							UNIQUE KEY category_un (name)
						)
		`
		createTableProductQuery = `
			CREATE TABLE IF NOT EXISTS
				product 
					(
						product_id int NOT NULL AUTO_INCREMENT,
						name varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
						price bigint NOT NULL,
						category int DEFAULT NULL,
						stock int NOT NULL,
						PRIMARY KEY (product_id),
						KEY product_FK (category),
						CONSTRAINT product_FK FOREIGN KEY (category) REFERENCES category (category_id) ON DELETE RESTRICT ON UPDATE CASCADE
					)
		`
		createTableCartQuery = `
			CREATE TABLE IF NOT EXISTS
				cart 
					(
						user_id int NOT NULL,
						product_id int NOT NULL,
						quantity int NOT NULL DEFAULT '1',
						KEY cart_FK (user_id),
						KEY cart_FK_1 (product_id),
						CONSTRAINT cart_FK FOREIGN KEY (user_id) REFERENCES user (user_id) ON DELETE CASCADE ON UPDATE CASCADE,
						CONSTRAINT cart_FK_1 FOREIGN KEY (product_id) REFERENCES product (product_id) ON DELETE CASCADE ON UPDATE CASCADE
				)
		`
	)

	_, err = db.Exec(createTableUserQuery)

	if err != nil {
		log.Panic("error while creating users table: ", err.Error())
		return
	}

	_, err = db.Exec(createTableCategoryQuery)

	if err != nil {
		log.Panic("error while creating category table: ", err.Error())
		return
	}

	_, err = db.Exec(createTableProductQuery)

	if err != nil {
		log.Panic("error while creating social_media table: ", err.Error())
		return
	}

	_, err = db.Exec(createTableCartQuery)

	if err != nil {
		log.Panic("error while creating comments table: ", err.Error())
		return
	}
}
