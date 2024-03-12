package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseConfig represents the configuration for the database connection.
type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// NewDatabaseConfig creates a new DatabaseConfig instance from environment variables.
func NewDatabaseConfig() (*DatabaseConfig, error) {
	// load env
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file")
	}
	return &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}, nil
}

// Connect connects to the database using the provided configuration.
func Connect(config *DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db, nil
}
