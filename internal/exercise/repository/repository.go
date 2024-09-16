package repository

import (
	"context"

	exm "github.com/bellyachx/hercules-be/internal/exercise/model"
	"gorm.io/gorm"
)

type DBStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *DBStore {
	return &DBStore{db: db}
}

func (d *DBStore) SaveExercise(_ context.Context, ex *exm.Exercise) error {
	return d.db.Create(&ex).Error
}

func (d *DBStore) GetExercises(_ context.Context) ([]exm.Exercise, error) {
	var exercises []exm.Exercise
	result := d.db.Find(&exercises)
	return exercises, result.Error
}
