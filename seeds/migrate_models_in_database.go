//go:build exclude
// +build exclude

package main

import (
	"github.com/flakrimjusufi/grpc-with-rest/database"
	"github.com/flakrimjusufi/grpc-with-rest/models"
)

func main() {

	var db = database.Connect()

	if err := db.Debug().AutoMigrate(&models.User{}).Error; err != nil {
		panic(err)
	}
	if err := db.Debug().AutoMigrate(&models.CreditCards{}).Error; err != nil {
		panic(err)
	}
	if err := db.Debug().AutoMigrate(&models.CreditCardApplication{}).Error; err != nil {
		panic(err)
	}
}
