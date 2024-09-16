package main

import (
	"github.com/bellyachx/hercules-be/internal/common/logger"
	"log"
	"os"

	"github.com/bellyachx/hercules-be/config"
	"github.com/bellyachx/hercules-be/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found")
	}

	cfg := config.NewConfig()

	logger.Init(cfg.Logging.Level)
	defer logger.Sync()

	if err := server.Start(cfg, logger.GetLogger()); err != nil {
		logger.GetLogger().Fatalf("Server failed to start: %v", err)
		os.Exit(-1)
	}
}
