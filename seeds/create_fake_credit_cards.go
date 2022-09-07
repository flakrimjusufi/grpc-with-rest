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
	db.AutoMigrate(&models.CreditCards{})

	numberOfRecords := 100
	count := 1
	for i := 1; i <= numberOfRecords; i++ {
		gofakeit.Seed(0)
		name := gofakeit.Name()
		email := gofakeit.Email()
		phone := gofakeit.Phone()
		address := gofakeit.Address().Address
		country := gofakeit.Address().Country
		city := gofakeit.Address().City
		zip := gofakeit.Address().Zip
		cvv := gofakeit.CreditCardCvv()
		createdAt := time.Now()

		db.Create(&models.CreditCards{Name: name, Email: email, PhoneNumber: phone, Address: address, Country: country,
			City: city, Zip: zip, CVV: cvv, CreatedAt: createdAt})
		count = i
	}
	defer db.Close()
	log.Printf("Seed executed successfully. Number or records created: %d", count)
}
