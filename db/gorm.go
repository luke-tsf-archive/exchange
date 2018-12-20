package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/luke-tsf/exchange/model"
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

func FindOneUser(condition interface{}) (model.UserModel, error) {
	var model model.UserModel
	err := DB.Where(condition).First(&model).Error
	return model, err
}

func VerifyEmailExistence(email string) bool {
	_, err := FindOneUser(&model.UserModel{Email: email})
	if err == nil {
		return false
	}
	return true
}
