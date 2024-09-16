package server

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/bellyachx/hercules-be/api/exercisepb"
	"github.com/bellyachx/hercules-be/internal/exercise/model"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) CreateExercise(ctx context.Context, exercise *model.Exercise) error {
	args := m.Called(ctx, exercise)
	return args.Error(0)
}

func (m *MockService) GetExercises(ctx context.Context) ([]model.Exercise, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Exercise), args.Error(1)
}

func getMockPbExerciseRequest() exercisepb.Exercise {
	userId, _ := uuid.NewRandom()
	return exercisepb.Exercise{
		UserId: userId.String(),
	}
}

func Test_server_CreateExercise(t *testing.T) {
	type args struct {
		ctx      context.Context
		exercise *exercisepb.Exercise
	}

	request := getMockPbExerciseRequest()
	tests := []struct {
		name       string
		args       args
		want       *exercisepb.ExerciseCreatedResponse
		wantSvcErr error
		wantErr    error
	}{
		{
			name: "Should not create exercise and return error that request is empty",
			args: args{
				exercise: nil,
			},
			wantErr: errors.New("exercise is empty"),
		},
		{
			name: "Should not create exercise and return error that request is empty",
			args: args{
				exercise: &exercisepb.Exercise{},
			},
			wantSvcErr: validator.ValidationErrors{},
			wantErr:    status.Error(codes.InvalidArgument, ""),
		},
		{
			name: "Should not validate exercise request and return error",
			args: args{
				exercise: &request,
			},
			wantSvcErr: validator.ValidationErrors{},
			wantErr:    status.Error(codes.InvalidArgument, ""),
		},
		{
			name: "Should not create exercise and return internal error",
			args: args{
				exercise: &request,
			},
			wantSvcErr: errors.New("internal error"),
			wantErr:    status.Errorf(codes.Internal, ""),
		},
		{
			name: "Should create exercise and return positive response",
			args: args{
				exercise: &request,
			},
			want: &exercisepb.ExerciseCreatedResponse{
				Message: "exercise created",
			},
			wantSvcErr: nil,
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockService)
			mockService.On("CreateExercise", tt.args.ctx, mock.AnythingOfType(reflect.TypeOf(&model.Exercise{}).Name())).Return(tt.wantSvcErr)
			s := &server{
				service: mockService,
			}
			got, err := s.CreateExercise(tt.args.ctx, tt.args.exercise)
			if tt.wantErr != nil {
				assert.Error(t, err)
				require.Nil(t, got)
				require.ErrorContains(t, err, tt.wantErr.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_server_GetExercises(t *testing.T) {
	tests := []struct {
		name       string
		wantSvc    []model.Exercise
		want       *exercisepb.GetExercisesResponse
		wantSvcErr error
		wantErr    error
	}{
		{
			name:       "Should not return exercises and return error",
			wantSvcErr: errors.New("internal svc error"),
			wantErr:    status.Errorf(codes.Internal, "failed to retrieve exercises"),
		},
		{
			name: "Should return exercises and no error",
			wantSvc: []model.Exercise{
				{
					Name:        "test_ex",
					Description: "test_desc",
					RepsCount:   11,
				},
				{
					Name:        "test_ex1",
					Description: "test_desc1",
					SetsCount:   3,
					Duration:    3000,
					MuscleGroup: "shoulders",
				},
			},
			want: &exercisepb.GetExercisesResponse{
				Exercises: []*exercisepb.Exercise{
					{
						Name:        "test_ex",
						Description: "test_desc",
						RepsCount:   11,
					},
					{
						Name:        "test_ex1",
						Description: "test_desc1",
						SetsCount:   3,
						Duration:    3000,
						MuscleGroup: "shoulders",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockService)
			mockService.On("GetExercises", mock.Anything, mock.Anything).Return(tt.wantSvc, tt.wantSvcErr)
			s := &server{
				service: mockService,
			}
			got, err := s.GetExercises(nil, nil)
			if tt.wantErr != nil {
				assert.Error(t, err)
				require.Nil(t, got)
				require.ErrorContains(t, err, tt.wantSvcErr.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}
