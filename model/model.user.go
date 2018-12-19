package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/luke-tsf/exchange/db"
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
	User struct {
		gorm.Model
		Username  string    `json:"Username"`
		Email     string    `json:"Email"`
		Salt      string    `json:"Salt"`
		PwdHash   string    `json:"PasswordHash"`
		LastLogin time.Time `json:"LastLoginDate"`
	}
)

func MigrateDB() {
	db := db.GetDB()
	db.AutoMigrate(&User{})
}
