package main

import (
	"github.com/bellyachx/hercules-be/internal/exercise"
	exsrv "github.com/bellyachx/hercules-be/internal/exercise/server"
	"github.com/bellyachx/hercules-be/internal/infrastructure/db"
	exm "github.com/bellyachx/hercules-be/internal/model/exercise"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found")
	}
	if err := run(); err != nil {
		log.Fatalln("failed to start server")
	}
}

func run() error {
	port, exists := os.LookupEnv("PORT")
	port = ":" + port
	if !exists {
		port = ":8080"
	}
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("failed to create listener", err)
	}

	s := grpc.NewServer()
	defer s.GracefulStop()
	reflection.Register(s)

	initServer(s)

	return s.Serve(listener)
}

func initServer(s *grpc.Server) {
	dbConn, err := db.Open()
	if err != nil {
		log.Fatalln("cannot connect to db", err)
	}
	err = dbConn.AutoMigrate(&exm.Exercise{})
	if err != nil {
		log.Fatalln("automigrate failed", err)
	}

	store := exercise.NewStore(dbConn)
	service := exercise.NewService(store)
	exsrv.RegisterServer(s, service)
}
