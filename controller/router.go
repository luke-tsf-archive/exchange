package controller

import (
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func Init() *gin.Engine {
	Router := gin.Default()
	Router.POST("/", UsersRegistration)
	Router.POST("/login", UsersLogin)
	return Router
}
