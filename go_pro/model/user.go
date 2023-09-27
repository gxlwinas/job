package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model

	Username        string `json:"username"`
	Password        string `gorm:"size:255;not null" json:"password"`
	Email           string `json:"email"`
	isAdministrator bool   `default:"false"`
}

type Newuser struct {
	Username string `json:"username"`
	Password string `gorm:"size:255;not null" json:"password"`
	Email    string `json:"email"`
	Captcha  string
	Deleteat time.Time
}
