package model

import (
	"github.com/jinzhu/gorm"
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
