package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Open() (*gorm.DB, error) {
	driver := os.Getenv("DB_DRIVER")
	dsn := os.Getenv("DB_DATASOURCE_NAME")
	var d gorm.Dialector
	switch driver {
	case "postgres":
		d = postgres.Open(dsn)
	}
	return gorm.Open(d, &gorm.Config{})
}
