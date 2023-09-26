package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name            string `gorm:"varchar(20);not null"`
	Username        string
	Password        string `gorm:"size:255;not null"`
	Email           *string
	isAdministrator bool
	isPublisher     bool
	Brithday        *time.Time
}
