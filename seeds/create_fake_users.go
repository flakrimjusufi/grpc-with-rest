package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jinzhu/gorm"
	"log"
	"server/main.go/database"
)

type User struct {
	gorm.Model
	Name        string
	Email       string
	PhoneNumber string
}

func main() {

	var db = database.Connect()
	db.AutoMigrate(&User{})

	numberOfRecords := 100000
	count := 1
	for i := 1; i <= numberOfRecords; i++ {
		gofakeit.Seed(0)
		name := gofakeit.Name()
		email := gofakeit.Email()
		phone := gofakeit.Phone()

		db.Create(&User{Name: name, Email: email, PhoneNumber: phone})
		count = i
	}
	defer db.Close()
	log.Printf("Seed executed successfully. Number or records created: %d", count)
}
