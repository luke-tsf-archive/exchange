package model

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/luke-tsf/exchange/helpers"
)

/*
User ID (incrementing bigint)
User Common Name (to be displayed on the site)
User Email Address
Password Salt (Unique for every user, inserted when the account is created)
Password (Hashed with the salt - MD5 or SHA1, your preference)
Date Account Was Created
*/
type (
	UserModel struct {
		gorm.Model
		Username  string
		Email     string
		Salt      []byte
		PwdHash   []byte
		LastLogin time.Time
	}
)

// Validator: write the form/json checking rule according to the doc
// https://github.com/go-playground/validator
// https://github.com/gin-gonic/gin#model-binding-and-validation
type UserModelValidator struct {
	User struct {
		// Username string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
		// Email    string `form:"email" json:"email" binding:"exists,email"`
		// Password string `form:"password" json:"password" binding:"exists,min=8,max=255"`
		Username string `json:"username" binding:"exists"`
		Email    string `json:"email" binding:"exists"`
		Password string `json:"password" binding:"exists"`
	} `json:"user"`
	userModel UserModel `json:"-"`
}

func (self *UserModelValidator) GetUserModel() UserModel {
	return self.userModel
}
func (self *UserModelValidator) Bind(c *gin.Context) error {
	err := helpers.Bind(c, self)
	if err != nil {
		return err
	}
	self.userModel.Username = self.User.Username
	self.userModel.Email = self.User.Email
	log.Printf("Username %v", self.User.Username)
	log.Printf("Email %v", self.User.Email)
	log.Printf("Password %v", self.User.Password)
	if self.User.Password != helpers.NBRandomPassword {
		salt, hash, err := helpers.GenerateSaltAndHash(self.User.Username, self.User.Password)
		if err != nil {
			return err
		}
		// if helpers.VerifyHashWithSalt(self.User.Username, salt, hash) == false {
		// 	return errors.New("Erroring verify password hash with salt")
		// }
		self.userModel.PwdHash = hash
		self.userModel.Salt = salt
	}
	self.userModel.LastLogin = time.Now()
	self.userModel.CreatedAt = time.Now()
	// self.userModel.UpdatedAt = time.Now()
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

type UserSerializer struct {
	C *gin.Context
}
type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (self *UserSerializer) Response() UserResponse {
	myUserModel := self.C.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		// Token:    helpers.GenToken(myUserModel.ID),
	}
	return user
}
