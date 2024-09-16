package server

import (
	"context"
	"errors"

	"github.com/bellyachx/hercules-be/internal/common/logger"
	"github.com/bellyachx/hercules-be/internal/exercise/mapper"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/bellyachx/hercules-be/api/exercisepb"
	"github.com/bellyachx/hercules-be/internal/exercise/repository"
	"github.com/bellyachx/hercules-be/internal/exercise/service"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func InitializeService(grpcServer *grpc.Server, dbConn *gorm.DB, log logger.Logger) {
	if log == nil {
		log = logger.GetLogger()
	}
	repo := repository.NewRepository(dbConn)
	svc := service.NewService(repo, log)
	exercisepb.RegisterExerciseServiceServer(grpcServer, &server{service: svc})
}

type server struct {
	exercisepb.UnimplementedExerciseServiceServer

	service service.Service
}

func (s *server) CreateExercise(ctx context.Context, exercise *exercisepb.Exercise) (*exercisepb.ExerciseCreatedResponse, error) {
	if exercise == nil {
		return nil, status.Error(codes.InvalidArgument, "exercise is empty")
	}
	exModel, err := mapper.MapToModel(exercise)
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot map model")
	}

	err = s.service.CreateExercise(ctx, exModel)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "failed to create exercise: %v", err)
	}
	return &exercisepb.ExerciseCreatedResponse{Message: "exercise created"}, nil
}

func (s *server) GetExercises(ctx context.Context, _ *emptypb.Empty) (*exercisepb.GetExercisesResponse, error) {
	exercises, err := s.service.GetExercises(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve exercises: %v", err)
	}

	exSlice, err := mapper.MapFromModelSlice(exercises)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to map model to pb response: %v", err)
	}
	return &exercisepb.GetExercisesResponse{Exercises: exSlice}, nil
}
