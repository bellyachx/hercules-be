package server

import (
	"context"
	"github.com/bellyachx/hercules-be/api/exercisepb"
	"github.com/bellyachx/hercules-be/internal/exercise"
	"google.golang.org/grpc"
)

type server struct {
	exercisepb.UnimplementedExerciseServiceServer

	service *exercise.Service
}

func (s *server) CreateExercise(ctx context.Context, exercise *exercisepb.Exercise) (*exercisepb.ExerciseCreatedResponse, error) {
	return s.service.CreateExercise(exercise)
}

func (s *server) GetExercises(ctx context.Context, request *exercisepb.GetExercisesRequest) (*exercisepb.GetExercisesResponse, error) {
	return s.service.GetExercises()
}

func RegisterServer(s *grpc.Server, service *exercise.Service) {
	exercisepb.RegisterExerciseServiceServer(s, &server{service: service})
}
