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
		// v1.GET("/", fetchAllTodo)
		// v1.GET("/:id", fetchSingleTodo)
		// v1.PUT("/:id", updateTodo)
		// v1.DELETE("/:id", deleteTodo)
	}
	router.Run()
}
