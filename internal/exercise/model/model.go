package model

import (
	"time"

	"github.com/google/uuid"
)

type Exercise struct {
	ExerciseID  uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	UserID      uuid.UUID `gorm:"not null" validate:"required,uuid4"`
	ImageID     uuid.UUID `validate:"omitempty,uuid4"`
	GifID       uuid.UUID `validate:"omitempty,uuid4"`
	VideoID     uuid.UUID `validate:"omitempty,uuid4"`
	Name        string    `gorm:"not null" validate:"required,min=3,max=255"`
	Description string    `gorm:"type:text"`
	MuscleGroup string    `gorm:"type:varchar(100)" validate:"required"`
	Difficulty  string    `gorm:"type:varchar(50)" validate:"required"`
	Type        string    `gorm:"type:varchar(100)" validate:"required"`
	SetsCount   int32     `validate:"omitempty,gte=1"`
	RepsCount   int32     `validate:"omitempty,gte=1"`
	Duration    int64     `validate:"omitempty,gte=0"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
