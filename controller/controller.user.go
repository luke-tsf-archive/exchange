package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luke-tsf/exchange/db"
	"github.com/luke-tsf/exchange/helpers"
	"github.com/luke-tsf/exchange/model"
)

func UsersRegistration(c *gin.Context) {
	userModelValidator := model.NewUserModelValidator()
	log.Printf("Verify user register information")
	if err := userModelValidator.Bind(c); err != nil {
		log.Printf("Error in UsersRegistration %+v \n", err)
		c.JSON(http.StatusUnprocessableEntity, helpers.NewValidatorError(err))
		return
	}
	userModel := userModelValidator.GetUserModel()
	if err := db.DB.Save(&userModel).Error; err != nil {
		c.JSON(http.StatusUnprocessableEntity, helpers.NewError("database", err))
		return
	}
	c.Set("my_user_model", userModelValidator.GetUserModel())
	serializer := model.UserSerializer{C: c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}
