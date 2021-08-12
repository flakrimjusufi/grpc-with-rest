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

type CreditCardApplication struct {
	ID                   uint `gorm:"primary_key"`
	FirstName            string
	LastName             string
	DateOfBirth          time.Time
	PhoneNumber          string
	SocialSecurityNumber string
	EmploymentType       string
	Occupation           string
	MonthlyIncome        float64
	YearsEmployed        int
	StreetAddress        string
	YearsAtAddress       int
	City                 string
	State                string
	Zip                  string
	Country              string
	Ownership            string
	MonthlyPayment       float64
	CardName             string
	CardType             string
	Branch               string
	CardBranding         string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            time.Time
}
