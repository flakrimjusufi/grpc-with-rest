package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
}

type CreditCards struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Zip         string `json:"zip"`
	CVV         string `json:"CVV"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type CreditCardApplication struct {
	ID                   uint      `gorm:"primary_key"`
	FirstName            string    `json:"firstName"`
	LastName             string    `json:"lastName"`
	DateOfBirth          time.Time `json:"dateOfBirth"`
	PhoneNumber          string    `json:"phoneNumber"`
	SocialSecurityNumber string    `json:"socialSecurityNumber"`
	EmploymentType       string    `json:"employmentType"`
	Occupation           string    `json:"occupation"`
	MonthlyIncome        float64   `json:"monthlyIncome"`
	YearsEmployed        int       `json:"yearsEmployed"`
	StreetAddress        string    `json:"streetAddress"`
	YearsAtAddress       int       `json:"yearsAtAddress"`
	City                 string    `json:"city"`
	State                string    `json:"state"`
	Zip                  string    `json:"zip"`
	Country              string    `json:"country"`
	Ownership            string    `json:"ownership"`
	MonthlyPayment       float64   `json:"monthlyPayment"`
	CardName             string    `json:"cardName"`
	CardType             string    `json:"cardType"`
	Branch               string    `json:"branch"`
	CardBranding         string    `json:"cardBranding"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            time.Time
}
