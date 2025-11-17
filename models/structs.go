package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Name  string `gorm:"size:50;not null" json:"name"`
	Email string `gorm:"size:100;not null" json:"email"`
	Age   int    `gorm:"not null" json:"age"`
}

type InputUser struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required,gte=0"`
}

type UpdateUser struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Age   *int    `json:"age"`
}

type PutUpdateUser struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required,gte=0"`
}
