package service

import (
	"context"

	"github.com/bellyachx/hercules-be/internal/common/logger"
	"github.com/bellyachx/hercules-be/internal/exercise/model"
	"github.com/bellyachx/hercules-be/internal/exercise/repository"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateExercise(ctx context.Context, exercise *model.Exercise) error
	GetExercises(ctx context.Context) ([]model.Exercise, error)
}

type service struct {
	repository repository.Repository
	validator  *validator.Validate
	logger     logger.Logger
}

func NewService(repo repository.Repository, log logger.Logger) Service {
	if log == nil {
		log = logger.GetLogger()
	}
	return &service{
		repository: repo,
		validator:  validator.New(),
		logger:     log,
	}
}

func (s *service) CreateExercise(ctx context.Context, exercise *model.Exercise) error {
	if err := s.validator.Struct(exercise); err != nil {
		s.logger.Errorf("validation error %v", err.Error())
		return err
	}

	if err := s.repository.SaveExercise(ctx, exercise); err != nil {
		s.logger.Errorf("failed to create exercise %v", err)
		return err
	}
	return nil
}

func (s *service) GetExercises(ctx context.Context) ([]model.Exercise, error) {
	exercises, err := s.repository.GetExercises(ctx)
	if err != nil {
		s.logger.Errorf("failed to retrieve exercises %v", err)
		return nil, err
	}
	return exercises, nil
}
