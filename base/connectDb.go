package base

import (
	"wiwieie011/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func ConnectionDB(){
	dsn := "host=localhost user=postgres password=tekhiev dbname=fiber port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err !=nil{
		panic("fail connection to database")
	}

	db.AutoMigrate(&models.User{})
	DB = db

}