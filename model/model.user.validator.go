package model

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luke-tsf/exchange/helpers"
)

// Validator: write the form/json checking rule according to the doc
// https://github.com/go-playground/validator
// https://github.com/gin-gonic/gin#model-binding-and-validation
type UserModelValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
	} `json:"user"`
	userModel UserModel `json:"-"`
}

func (self *UserModelValidator) GetUserModel() UserModel {
	return self.userModel
}
func (self *UserModelValidator) Bind(c *gin.Context) error {
	log.Printf("Binding input information to Model")
	err := helpers.Bind(c, self)
	if err != nil {
		return err
	}
	log.Printf("Process input from user %v", self.User.Username)
	self.userModel.Username = self.User.Username
	self.userModel.Email = self.User.Email
	if self.User.Password != helpers.NBRandomPassword {
		salt, hash, err := helpers.GenerateSaltAndHash(self.User.Username, self.User.Password)
		if err != nil {
			return err
		}
		if helpers.VerifyHashWithSalt(self.User.Password, salt, hash) == false {
			return errors.New("Erroring verify password hash with salt")
		}
		self.userModel.PwdHash = string(hash)
		self.userModel.Salt = string(salt)
	}
	self.userModel.CreatedAt = time.Now()
	return nil
}

func NewUserModelValidator() UserModelValidator {
	return UserModelValidator{}
}

func NewUserModelValidatorFillWith(userModel UserModel) UserModelValidator {
	userModelValidator := NewUserModelValidator()
	userModelValidator.User.Username = userModel.Username
	userModelValidator.User.Email = userModel.Email
	userModelValidator.User.Password = helpers.NBRandomPassword
	return userModelValidator
}
