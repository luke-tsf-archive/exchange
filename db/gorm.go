package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open("mysql", "root:exchange@123@/exchange?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to db")
	}
	log.Println("Successfully connect to database")
}

func SaveOne(data interface{}) error {
	err := DB.Save(data).Error
	return err
}
