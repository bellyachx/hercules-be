package server

import (
	"context"
	"errors"
	"github.com/bellyachx/hercules-be/internal/common/logger"

	"github.com/bellyachx/hercules-be/api/exercisepb"
	"github.com/bellyachx/hercules-be/internal/exercise/model"
	"github.com/bellyachx/hercules-be/internal/exercise/repository"
	"github.com/bellyachx/hercules-be/internal/exercise/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func InitializeService(grpcServer *grpc.Server, dbConn *gorm.DB, log logger.Logger) {
	if log == nil {
		log = logger.GetLogger()
	}
	store := repository.NewStore(dbConn)
	svc := service.NewService(store, log)
	exercisepb.RegisterExerciseServiceServer(grpcServer, &server{service: svc})
}

type server struct {
	exercisepb.UnimplementedExerciseServiceServer

	service service.Service
}

func (s *server) CreateExercise(ctx context.Context, exercise *exercisepb.Exercise) (*exercisepb.ExerciseCreatedResponse, error) {
	exModel := &model.Exercise{
		UserID:      uuid.MustParse(exercise.UserId),
		Name:        exercise.Name,
		Description: exercise.Description,
		MuscleGroup: exercise.MuscleGroup,
		Difficulty:  exercise.Difficulty,
		Type:        exercise.Type,
		SetsCount:   exercise.SetsCount,
		RepsCount:   exercise.RepsCount,
		Duration:    exercise.Duration,
	}

	err := s.service.CreateExercise(ctx, exModel)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to create exercise: %v", err)
	}
	return &exercisepb.ExerciseCreatedResponse{Message: "exercise created"}, nil
}

func (s *server) GetExercises(ctx context.Context, _ *exercisepb.GetExercisesRequest) (*exercisepb.GetExercisesResponse, error) {
	exercises, err := s.service.GetExercises(ctx)
	if err != nil {
		return nil, err
	}

	exSlice := make([]*exercisepb.Exercise, len(exercises))

	for i, val := range exercises {
		exSlice[i] = &exercisepb.Exercise{
			UserId:      val.UserID.String(),
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
