package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name        string
	Email       string
	PhoneNumber string
}

type CreditCards struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Email       string
	PhoneNumber string
	Address     string
	Country     string
	City        string
	Zip         string
	CVV         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
