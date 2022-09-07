//go:build exclude
// +build exclude

package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/flakrimjusufi/grpc-with-rest/database"
	"github.com/flakrimjusufi/grpc-with-rest/models"
	"log"
)

func main() {

	var db = database.Connect()
	db.AutoMigrate(&models.User{})

	numberOfRecords := 100
	count := 1
	for i := 1; i <= numberOfRecords; i++ {
		gofakeit.Seed(0)
		name := gofakeit.Name()
		email := gofakeit.Email()
		phone := gofakeit.Phone()

		db.Create(&models.User{Name: name, Email: email, PhoneNumber: phone})
		count = i
	}
	defer db.Close()
	log.Printf("Seed executed successfully. Number or records created: %d", count)
}
