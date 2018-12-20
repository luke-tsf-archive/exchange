package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luke-tsf/exchange/db"
	"github.com/luke-tsf/exchange/helpers"
	"github.com/luke-tsf/exchange/middleware"
	"github.com/luke-tsf/exchange/model"
)

func UsersRegistration(c *gin.Context) {
	userModelValidator := model.NewUserModelValidator()
	log.Printf("Register: Verify user information")
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

func UsersLogin(c *gin.Context) {
	loginValidator := model.NewLoginValidator()
	log.Println("Login: Verify user information")
	if err := loginValidator.Bind(c); err != nil {
		log.Printf("Error in Login %+v \n", err)
		c.JSON(http.StatusUnprocessableEntity, helpers.NewValidatorError(err))
		return
	}
	userModel, err := db.FindOneUser(model.UserModel{Email: loginValidator.User.Email})
	if err != nil {
		log.Printf("Error in Login %+v \n", err)
		c.JSON(http.StatusForbidden, helpers.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	if err = userModel.CheckPassword(loginValidator.User.Password); err != nil {
		log.Printf("Error in Login %+v \n", err)
		c.JSON(http.StatusForbidden, helpers.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	log.Printf("Login: Update Context for UserModel %s\n", userModel.ID)
	middleware.UpdateContextUserModel(c, userModel.ID)
	serializer := model.UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}
