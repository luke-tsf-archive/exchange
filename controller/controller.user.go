package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luke-tsf/exchange/db"
	"github.com/luke-tsf/exchange/helpers"
	"github.com/luke-tsf/exchange/model"
)

func UserRegistration(c *gin.Context) {
	userModelValidator := model.NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, helpers.NewValidatorError(err))
		return
	}

	if err := db.SaveOne(userModelValidator.GetUserModel()); err != nil {
		c.JSON(http.StatusUnprocessableEntity, helpers.NewError("database", err))
		return
	}
	c.Set("my_user_model", userModelValidator.GetUserModel())
	serializer := model.UserSerializer{C: c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}
