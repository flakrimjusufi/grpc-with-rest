//go:build exclude
// +build exclude

package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/flakrimjusufi/grpc-with-rest/database"
	"github.com/flakrimjusufi/grpc-with-rest/models"
	"log"
	"time"
)

func main() {

	var db = database.Connect()
	db.AutoMigrate(&models.CreditCardApplication{})

	numberOfRecords := 100
	count := 1
	for i := 1; i <= numberOfRecords; i++ {
		gofakeit.Seed(0)
		firstName := gofakeit.Name()
		lastName := gofakeit.LastName()
		dateOfBirth := gofakeit.Date()
		phoneNumber := gofakeit.Phone()
		socialSecurityNumber := gofakeit.SSN()
		employmentType := "Administrator"
		occupation := gofakeit.Job().Title
		yearsEmployed := gofakeit.Number(1, 20)
		monthlyIncome := gofakeit.Price(1000, 5000)
		streetAddress := gofakeit.Address().Address
		yearsAtAddress := gofakeit.Number(1, 20)
		city := gofakeit.City()
		state := gofakeit.State()
		zip := gofakeit.Zip()
		country := gofakeit.Country()
		ownership := "No"
		monthlyPayment := gofakeit.Price(300, 1200)
		cardName := gofakeit.Name()
		cardType := gofakeit.CreditCardType()
		branch := "Bank"
		cardBranding := "Bank card branding"

		db.Create(&models.CreditCardApplication{
			FirstName:            firstName,
			LastName:             lastName,
			DateOfBirth:          dateOfBirth,
			PhoneNumber:          phoneNumber,
			SocialSecurityNumber: socialSecurityNumber,
			EmploymentType:       employmentType,
			Occupation:           occupation,
			MonthlyIncome:        monthlyIncome,
			YearsEmployed:        yearsEmployed,
			StreetAddress:        streetAddress,
			YearsAtAddress:       yearsAtAddress,
			City:                 city,
			State:                state,
			Zip:                  zip,
			Country:              country,
			Ownership:            ownership,
			MonthlyPayment:       monthlyPayment,
			CardName:             cardName,
			CardType:             cardType,
			Branch:               branch,
			CardBranding:         cardBranding,
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
			DeletedAt:            time.Now(),
		})

		count = i
	}
	defer db.Close()
	log.Printf("Seed executed successfully. Number or records created: %d", count)
}
