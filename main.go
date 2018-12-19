package main

import (
	"github.com/luke-tsf/exchange/controller"
	"github.com/luke-tsf/exchange/db"
	"github.com/luke-tsf/exchange/model"
)

func main() {
	db.Init()
	db.DB.AutoMigrate(&model.UserModel{})
	defer db.DB.Close()
	router := controller.Init()
	router.Run()

}
