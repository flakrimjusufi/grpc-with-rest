package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	colorRed = "\033[31m"
)

func Connect() *gorm.DB {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbType := os.Getenv("db_type")

	if dbHost == "" || dbName == "" || dbType == "" {
		log.Println(colorRed, ".env file is empty. Please add the environment variables in .env file first and run the server again!")
	}

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string

	db, err := gorm.Open(dbType, dbUri)
	if err != nil {
		panic(err)
	}
	return db

}
