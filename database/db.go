package data

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_HOST = "localhost"
const DB_PORT = "3306"
const DB_NAME = "belajar_golang_gorm"

var Db *gorm.DB

func InitDb() *gorm.DB { // init export database
	Db = connectDB()

	return Db
}

func connectDB() *gorm.DB { // config database
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connect to database : error=%v", err)

		return nil
	}

	return db
}
