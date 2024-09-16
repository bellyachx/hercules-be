package server

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"
	"testing"

	"github.com/bellyachx/hercules-be/api/exercisepb"
	"github.com/bellyachx/hercules-be/internal/common/logger"
	"github.com/bellyachx/hercules-be/internal/exercise/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	testDB     *gorm.DB
	grpcServer *grpc.Server
	conn       exercisepb.ExerciseServiceClient
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	err = testDB.AutoMigrate(&model.Exercise{})
	if err != nil {
		log.Fatalf("Failed to migrate test database: %v", err)
	}

	lis, err := net.Listen("tcp", ":0") // Listen on random available port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer = grpc.NewServer()
	lgr := logger.GetLogger()
	InitializeService(grpcServer, testDB, lgr)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			lgr.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	addr := lis.Addr().String()
	connGRPC, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		lgr.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer func(connGRPC *grpc.ClientConn) {
		_ = connGRPC.Close()
	}(connGRPC)
	conn = exercisepb.NewExerciseServiceClient(connGRPC)

	code := m.Run()

	grpcServer.Stop()

	os.Exit(code)
}

func resetDB(t *testing.T) {
	if err := testDB.Exec("DELETE FROM exercises WHERE TRUE").Error; err != nil {
		t.Fatalf("Failed to reset database: %v", err)
	}
}

func TestCreateExercise(t *testing.T) {
	resetDB(t)
	req := &exercisepb.Exercise{
		UserId:      uuid.NewString(),
		Name:        "Push Up",
		Description: "A basic push up exercise",
		MuscleGroup: "chest",
		Difficulty:  "medium",
		Type:        "strength",
		SetsCount:   3,
		RepsCount:   15,
		Duration:    0,
	}

	resp, err := conn.CreateExercise(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateExercise failed: %v", err)
	}

	if resp.Message != "exercise created" {
		t.Errorf("Unexpected response message: got %v, want %v", resp.Message, "exercise created")
	}

	var exercise model.Exercise
	result := testDB.First(&exercise, "exercise_id = ?", exercise.ExerciseID)
	if result.Error != nil {
		t.Errorf("Failed to retrieve exercise from DB: %v", result.Error)
	}

	if exercise.Name != req.Name {
		t.Errorf("Exercise name mismatch: got %v, want %v", exercise.Name, req.Name)
	}
}

func TestGetExercises(t *testing.T) {
	resetDB(t)
	exercises := []model.Exercise{
		{
			ExerciseID:  uuid.NewString(),
			UserID:      uuid.NewString(),
			Name:        "Squat",
			Description: "A basic squat exercise",
			MuscleGroup: "legs",
			Difficulty:  "easy",
			Type:        "strength",
			SetsCount:   4,
			RepsCount:   10,
			Duration:    0,
		},
		{
			ExerciseID:  uuid.NewString(),
			UserID:      uuid.NewString(),
			Name:        "Plank",
			Description: "Core strengthening exercise",
			MuscleGroup: "core",
			Difficulty:  "hard",
			Type:        "endurance",
			SetsCount:   3,
			RepsCount:   0,
			Duration:    60,
		},
	}

	if err := testDB.Create(&exercises).Error; err != nil {
		t.Fatalf("Failed to seed database: %v", err)
	}

	resp, err := conn.GetExercises(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Fatalf("GetExercises failed: %v", err)
	}

	if len(resp.Exercises) != 2 {
		t.Errorf("Unexpected number of exercises: got %v, want %v", len(resp.Exercises), 2)
	}

	ex := resp.Exercises[0]
	if ex.Name != "Squat" {
		t.Errorf("First exercise name mismatch: got %v, want %v", ex.Name, "Squat")
	}

	ex = resp.Exercises[1]
	if ex.Name != "Plank" {
		t.Errorf("Second exercise name mismatch: got %v, want %v", ex.Name, "Plank")
	}
}

func TestCreateExercise_InvalidInput(t *testing.T) {
	resetDB(t)
	req := &exercisepb.Exercise{
		UserId:      "",   // Invalid UUID
		Name:        "Pu", // Name too short
		Description: "Invalid exercise",
		MuscleGroup: "",
		Difficulty:  "",
		Type:        "",
		SetsCount:   0,
		RepsCount:   -5,
		Duration:    -10,
	}

	_, err := conn.CreateExercise(context.Background(), req)
	if err == nil {
		t.Fatalf("Expected error for invalid input, got nil")
	}

	st, _ := status.FromError(err)
	if st.Code() != codes.InvalidArgument {
		t.Errorf("Unexpected error code: got %v, want %v", st.Code(), codes.InvalidArgument)
	}
}

func TestGetExercises_Empty(t *testing.T) {
	resetDB(t)
	resp, err := conn.GetExercises(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Fatalf("GetExercises failed: %v", err)
	}

	if len(resp.Exercises) != 0 {
		t.Errorf("Expected 0 exercises, got %v", len(resp.Exercises))
	}
}
