package service

import (
	"context"
	"errors"
	"github.com/bellyachx/hercules-be/internal/common/logger"
	"github.com/google/uuid"
	"testing"

	"github.com/bellyachx/hercules-be/internal/exercise/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) SaveExercise(ctx context.Context, ex *model.Exercise) error {
	args := m.Called(ctx, ex)
	return args.Error(0)
}

func (m *MockStore) GetExercises(ctx context.Context) ([]model.Exercise, error) {
	args := m.Called(ctx)
	exModel := args.Get(0)
	err := args.Error(1)
	if exModel == nil {
		return nil, err
	}
	return exModel.([]model.Exercise), err
}

type MockLogger struct {
	mock.Mock
	logger.Logger
}

func (m *MockLogger) Errorf(format string, args ...any) {
	m.Called(format, args)
}

// TestService_CreateExercise_Success tests the successful creation of an exercise.
func TestService_CreateExercise_Success(t *testing.T) {
	mockStore := new(MockStore)
	mockLogger := new(MockLogger)

	exercise := &model.Exercise{
		UserID:      uuid.NewString(),
		Name:        "Test Exercise",
		Description: "A test exercise",
		MuscleGroup: "legs",
		Difficulty:  "easy",
		Type:        "running",
		Duration:    30,
	}

	mockStore.On("SaveExercise", mock.Anything, exercise).Return(nil).Once()
	mockLogger.On("Errorf", mock.Anything, mock.Anything).Maybe()

	svc := NewService(mockStore, mockLogger)

	err := svc.CreateExercise(context.Background(), exercise)

	assert.NoError(t, err)
	mockStore.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
}

// TestService_CreateExercise_ValidationError tests the creation of an exercise with invalid data.
func TestService_CreateExercise_ValidationError(t *testing.T) {
	mockStore := new(MockStore)
	mockLogger := new(MockLogger)

	exercise := &model.Exercise{
		Name:        "", // Name is required
		Description: "A test exercise",
		Duration:    -10, // Invalid duration
	}

	mockLogger.On("Errorf", mock.Anything, mock.Anything).Return().Once()

	svc := NewService(mockStore, mockLogger)

	err := svc.CreateExercise(context.Background(), exercise)

	assert.Error(t, err)
	mockStore.AssertNotCalled(t, "SaveExercise", mock.Anything, mock.Anything)
	mockLogger.AssertCalled(t, "Errorf", mock.Anything, mock.Anything)
}

// TestService_CreateExercise_StoreError tests the creation of an exercise when the repository fails.
func TestService_CreateExercise_StoreError(t *testing.T) {
	mockStore := new(MockStore)
	mockLogger := new(MockLogger)

	exercise := &model.Exercise{
		UserID:      uuid.NewString(),
		Name:        "Test Exercise",
		Description: "A test exercise",
		MuscleGroup: "legs",
		Difficulty:  "easy",
		Type:        "running",
		Duration:    30,
	}

	storeErr := errors.New("database error")
	mockStore.On("SaveExercise", mock.Anything, exercise).Return(storeErr)
	mockLogger.On("Errorf", "failed to create exercise %v", mock.Anything).Once()

	svc := NewService(mockStore, mockLogger)

	err := svc.CreateExercise(context.Background(), exercise)

	assert.Error(t, err)
	assert.Equal(t, storeErr, err)
	mockStore.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
}

// TestService_GetExercises_Success tests successfully retrieving exercises.
func TestService_GetExercises_Success(t *testing.T) {
	mockStore := new(MockStore)
	mockLogger := new(MockLogger)

	exercises := []model.Exercise{
		{
			Name:        "Test Exercise",
			Description: "A test exercise",
			Duration:    30,
		},
		{
			Name:        "Another Exercise",
			Description: "Another test exercise",
			Duration:    45,
		},
	}

	mockStore.On("GetExercises", mock.Anything).Return(exercises, nil)

	svc := NewService(mockStore, mockLogger)

	result, err := svc.GetExercises(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, exercises, result)
	mockStore.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
}

// TestService_GetExercises_StoreError tests retrieving exercises when the repository fails.
func TestService_GetExercises_StoreError(t *testing.T) {
	mockStore := new(MockStore)
	mockLogger := new(MockLogger)

	storeErr := errors.New("database error")
	mockStore.On("GetExercises", mock.Anything).Return(nil, storeErr)
	mockLogger.On("Errorf", "failed to retrieve exercises %v", mock.Anything).Once()

	svc := NewService(mockStore, mockLogger)

	result, err := svc.GetExercises(context.Background())

	assert.Error(t, err)
	assert.Nil(t, result)
	mockStore.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
}
