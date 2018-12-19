package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luke-tsf/exchange/db"
	"github.com/luke-tsf/exchange/model"
)

func MigrateDB() {
	db := db.GetDB()
	db.AutoMigrate(&model.UserModel{})
}

func main() {
	router := gin.Default()
	router.Run()
}
