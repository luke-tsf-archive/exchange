package model

import (
	"github.com/gin-gonic/gin"
	"github.com/luke-tsf/exchange/helpers"
)

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
		Token:    helpers.GenToken(myUserModel.ID),
	}
	return user
}
