package handler

import (
	"net/http"
	"strconv"

	"github.com/luke-tsf/exchange/db"
	"github.com/luke-tsf/exchange/model"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := model.TodoModel{Title: c.PostForm("title"), Completed: completed}
	db.DB.Save(&todo)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully", "resourcedId": todo.ID})

}
