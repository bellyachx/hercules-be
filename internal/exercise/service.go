package exercise

import (
	"github.com/bellyachx/hercules-be/api/exercisepb"
	exm "github.com/bellyachx/hercules-be/internal/model/exercise"
)

type Service struct {
	store Store
}

type Store interface {
	SaveExercise(exercise *exercisepb.Exercise)
	GetExercises() []exm.Exercise
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

func (s *Service) CreateExercise(exercise *exercisepb.Exercise) (*exercisepb.ExerciseCreatedResponse, error) {
	s.store.SaveExercise(exercise)

	return &exercisepb.ExerciseCreatedResponse{Message: "exercise saved"}, nil
}

func (s *Service) GetExercises() (*exercisepb.GetExercisesResponse, error) {
	exercises := s.store.GetExercises()
	exSlice := make([]*exercisepb.Exercise, len(exercises))

	for i, val := range exercises {
		exSlice[i] = &exercisepb.Exercise{
			UserId:      val.UserID,
			Name:        val.Name,
			Description: val.Description,
			MuscleGroup: val.MuscleGroup,
			SetsCount:   val.SetsCount,
			RepsCount:   val.RepsCount,
			Duration:    val.Duration,
			Difficulty:  val.Difficulty,
			Type:        val.Type,
		}
	}
	return &exercisepb.GetExercisesResponse{Exercises: exSlice}, nil
}
