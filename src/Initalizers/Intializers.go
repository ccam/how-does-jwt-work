package Initalizers

import (
	"fmt"

	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := sqlite3.Open("sqlite3", "./../gorm.db")
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	DB = db
	return DB
}
