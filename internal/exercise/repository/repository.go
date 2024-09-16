package repository

import (
	"context"

	"github.com/bellyachx/hercules-be/internal/exercise/model"
	"gorm.io/gorm"
)

type Repository interface {
	SaveExercise(ctx context.Context, exercise *model.Exercise) error
	GetExercises(ctx context.Context) ([]model.Exercise, error)
}

type DBRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *DBRepository {
	return &DBRepository{db: db}
}

func (d *DBRepository) SaveExercise(_ context.Context, ex *model.Exercise) error {
	return d.db.Create(&ex).Error
}

func (d *DBRepository) GetExercises(_ context.Context) ([]model.Exercise, error) {
	var exercises []model.Exercise
	result := d.db.Find(&exercises)
	return exercises, result.Error
}
