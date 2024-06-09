package database

import (
	"fmt"
	"os"
	"time"

	"github.com/flakrimjusufi/grpc-with-rest/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Conn *gorm.DB
}

// NewDB - connects to the DB and returns a new DB struct
func NewDB() (*DB, error) {
	// Load environment variables from .env file if not already set
	if os.Getenv("DB_USERNAME") == "" {
		if err := godotenv.Load(); err != nil {
			return nil, fmt.Errorf("failed to load .env file: %w", err)
		}
	}

	// Fetch environment variables
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOSTNAME")
	dbPort := os.Getenv("DB_PORT")

	// Construct the Data Source Name (DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, username, password, dbName, dbPort)

	// Open a connection to the database using GORM
	dbConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Obtain the underlying sql.DB object to set connection pool settings
	sqlDB, err := dbConnect.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Set the maximum number of idle connections
	sqlDB.SetMaxIdleConns(10)

	// Set the maximum number of open connections
	sqlDB.SetMaxOpenConns(100)

	// Set the maximum lifetime of connections
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Perform automatic migrations for your models
	if handleErr := dbConnect.AutoMigrate(&models.User{}, &models.CreditCards{}, &models.CreditCardApplication{}); handleErr != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", handleErr)
	}

	// Return a new DB struct instance containing the gorm.DB connection
	return &DB{Conn: dbConnect}, nil
}
