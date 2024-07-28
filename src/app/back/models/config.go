package models

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Username string
	Password string
	Url      string
}

var instance *AppConfig
var once sync.Once

// GetConfig returns the singleton instance of Config
func GetConfig() *AppConfig {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
		instance = &AppConfig{
			Username: os.Getenv("ZINC_FIRST_ADMIN_USER"),
			Password: os.Getenv("ZINC_FIRST_ADMIN_PASSWORD"),
			Url:      "http://localhost:4080/api/emails",
		}
	})
	return instance
}
