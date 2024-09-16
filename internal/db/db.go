package db

import (
	"fmt"

	"github.com/bellyachx/hercules-be/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(cfg *config.Config) (*gorm.DB, error) {
	driver := cfg.Database.Driver

	var d gorm.Dialector
	switch driver {
	case "postgres":
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)
		pgConfig := postgres.Config{
			DSN: dsn,
		}
		d = postgres.New(pgConfig)
	default:
		return nil, fmt.Errorf("unknown driver: %s", driver)
	}
	return gorm.Open(d, &gorm.Config{})
}
