package server

import (
	"fmt"
	"github.com/bellyachx/hercules-be/internal/common/logger"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bellyachx/hercules-be/config"
	"github.com/bellyachx/hercules-be/internal/db"
	exerciseServer "github.com/bellyachx/hercules-be/internal/exercise/delivery/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func Start(cfg *config.Config, log logger.Logger) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatalf("failed to create listener on port %s. err: %v", cfg.Server.Port, err)
		return err
	}

	dbConn, err := db.Open(cfg)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
		return err
	}

	s := grpc.NewServer()
	reflection.Register(s)

	registerServices(s, dbConn, log)

	// Graceful shutdown
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve %v", err)
			os.Exit(-1)
		}
	}()
	log.Infof("Server is running on port: %s", cfg.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server gracefully...")
	s.GracefulStop()
	return nil
}

func registerServices(s *grpc.Server, dbConn *gorm.DB, log logger.Logger) {
	exerciseServer.InitializeService(s, dbConn, log)
}
