package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() *gorm.DB {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=kimi dbname=testdb password=213630")
	if err != nil {
		panic("Could not connect to db")
	}
	return db

}
