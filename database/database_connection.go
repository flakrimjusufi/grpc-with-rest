package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
	"server/main.go/models"
	"time"
)

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

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username,
		dbName, password) // connection string

	database, err := gorm.Open(dbType, dbUri)
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
