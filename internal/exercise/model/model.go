package model

import (
	"time"
)

type Exercise struct {
	ExerciseID  string    `json:"exercise_id" gorm:"primaryKey"`
	UserID      string    `json:"user_id" gorm:"not null" validate:"required,uuid4"`
	ImageID     string    `json:"image_id" validate:"omitempty,uuid4"`
	GifID       string    `json:"gif_id" validate:"omitempty,uuid4"`
	VideoID     string    `json:"video_id" validate:"omitempty,uuid4"`
	Name        string    `json:"name" gorm:"not null" validate:"required,min=3,max=255"`
	Description string    `json:"description" gorm:"type:text"`
	MuscleGroup string    `json:"muscle_group" gorm:"type:varchar(100)" validate:"required"`
	Difficulty  string    `json:"difficulty" gorm:"type:varchar(50)" validate:"required"`
	Type        string    `json:"type" gorm:"type:varchar(100)" validate:"required"`
	SetsCount   int32     `json:"sets_count" validate:"omitempty,gte=1"`
	RepsCount   int32     `json:"reps_count" validate:"omitempty,gte=1"`
	Duration    int64     `json:"duration" validate:"omitempty,gte=0"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
