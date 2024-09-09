package exercise

import (
	"github.com/bellyachx/hercules-be/api/exercisepb"
	exm "github.com/bellyachx/hercules-be/internal/model/exercise"
	"gorm.io/gorm"
)

type DBStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *DBStore {
	return &DBStore{db: db}
}

func (d *DBStore) SaveExercise(ex *exercisepb.Exercise) {
	model := exm.Exercise{
		UserID:      ex.UserId,
		Name:        ex.Name,
		Description: ex.Description,
		MuscleGroup: ex.MuscleGroup,
		Difficulty:  ex.Difficulty,
		Type:        ex.Type,
		SetsCount:   ex.SetsCount,
		RepsCount:   ex.RepsCount,
		Duration:    ex.Duration,
	}
	d.db.Create(&model)
}

func (d *DBStore) GetExercises() []exm.Exercise {
	var exercises []exm.Exercise
	d.db.Find(&exercises)

	return exercises
}
