package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name        string
	Email       string
	PhoneNumber string
}
