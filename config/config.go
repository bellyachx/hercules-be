package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   Server
	Database Database
	Logging  Logging
}

type Server struct {
	Port string
}

type Database struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Logging struct {
	Level string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}
	port := getEnv("PORT", "8080")
	logLevel := getEnv("LOG_LEVEL", "development")
	return &Config{
		Server: Server{
			Port: port,
		},
		Database: Database{
			Driver:   os.Getenv("DB_DRIVER"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
		Logging: Logging{
			Level: logLevel,
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
