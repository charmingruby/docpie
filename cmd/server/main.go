package main

import (
	"os"

	"github.com/charmingruby/make-it-survey/config"
	"github.com/charmingruby/make-it-survey/pkg/database/postgresql"
	"github.com/charmingruby/make-it-survey/pkg/logger"
)

func main() {
	logger := logger.SetupLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("error loading environment configuration")
		os.Exit(1)
	}
	cfg.AssignLogger(logger)

	db, err := postgresql.ConnectDB(cfg)
	if err != nil {
		logger.Errorf("error connecting to postgres: %s", err.Error())
		os.Exit(1)
	}
	cfg.AssignDatabaseConn(db)
}
