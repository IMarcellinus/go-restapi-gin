package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDB() {
	database, err := gorm.Open(mysql.Open("root:root123@tcp(localhost:3306)/data_buku"))
	if err != nil {
		panic(err.Error())
	}
	database.AutoMigrate(&Book{})

	DB = database
}
