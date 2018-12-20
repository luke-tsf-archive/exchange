package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/luke-tsf/exchange/db"
	"github.com/luke-tsf/exchange/model"
)

func UpdateContextUserModel(c *gin.Context, my_user_id uint) {
	var myUserModel model.UserModel
	if my_user_id != 0 {
		db.DB.First(&myUserModel, my_user_id)
	}
	c.Set("my_user_id", my_user_id)
	c.Set("my_user_model", myUserModel)
}
