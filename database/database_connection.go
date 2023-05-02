package database

import (
	"fmt"
	"github.com/flakrimjusufi/grpc-with-rest/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // imports postgres dialect
	"github.com/joho/godotenv"
	"os"
	"time"
)

// Connect - connects to the DB
func Connect() *gorm.DB {
	if os.Getenv("DB_USERNAME") == "" {
		e := godotenv.Load() //Load .env file for local environment
		if e != nil {
			panic(e)
		}
	}
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOSTNAME")
	dbType := os.Getenv("DB_TYPE")
	dbPort := os.Getenv("DB_PORT")

	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, username,
		password, dbName, dbPort) // connection string

	database, err := gorm.Open(dbType, dbURI)
	if err != nil {
		panic(err)
	}
	database.Debug().AutoMigrate(models.User{})
	database.Debug().AutoMigrate(models.CreditCards{})
	database.Debug().AutoMigrate(models.CreditCardApplication{})

	// SetMaxIdleConnections sets the maximum number of connections in the idle connection pool.
	database.DB().SetMaxIdleConns(10)

	// SetMaxOpenConnections sets the maximum number of open connections to the database.
	database.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	database.DB().SetConnMaxLifetime(time.Hour)

	return database
}
