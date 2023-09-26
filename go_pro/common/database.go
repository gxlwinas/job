package common

import (
	"go_pro/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.PartTimeJob{})
	db.AutoMigrate(&model.Job{})
	db.AutoMigrate(&model.JobJson{})

	DB = db

	return db
}

func GetDB() *gorm.DB {
	return DB
}