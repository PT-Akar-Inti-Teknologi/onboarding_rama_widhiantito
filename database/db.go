package database

import (
	"fmt"

	"gorm.io/driver/postgres" // Import the PostgreSQL driver
	"gorm.io/gorm"
)

// InitDB initializes a database connection and returns a pointer to the DB instance.
func InitDB() (*gorm.DB, error) {
	// Read database connection details from environment variables
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "postgres123"
	dbName := "postgres"

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Automigrate your models here if needed
	// For example, if you have an Employee model:
	// db.AutoMigrate(&Employee{})

	return db, nil
}
