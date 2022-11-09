package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const dsn = "root:@tcp(127.0.0.1:3306)/gormuxapi_db?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func InitDBMigration() {
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("CONNECTED")
	if err != nil {
		fmt.Println(err.Error())
		panic("cannot connect to DB")
	}

	DB.AutoMigrate(&User{})
}
