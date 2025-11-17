package base

import (
	"os"
	"wiwieie011/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func ConnectionDB(){
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err !=nil{
		panic("fail connection to database")
	}

	db.AutoMigrate(&models.User{})
	DB = db

}