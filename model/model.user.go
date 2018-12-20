package model

import (
	"errors"

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
		Username string
		Email    string
		Salt     string
		PwdHash  string
	}
)

func (u *UserModel) CheckPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	salt := []byte(u.Salt)
	hash := []byte(u.PwdHash)
	if helpers.VerifyHashWithSalt(password, salt, hash) == false {
		return errors.New("Error verifying password")
	}
	return nil
}
