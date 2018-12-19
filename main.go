package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luke-tsf/exchange/handler"
)

func main() {
	router := gin.Default()
	v1 := router.Group("api/v1/todos")
	{
		v1.POST("/", handler.CreateTodo)
		v1.GET("/", handler.FetchAllTodo)
		v1.GET("/:id", handler.FetchSingleTodo)
		v1.PUT("/:id", handler.UpdateTodo)
		v1.DELETE("/:id", handler.DeleteTodo)
	}
	router.Run()
}
