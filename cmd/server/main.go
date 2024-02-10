package main

import (
	"log"

	"github.com/charmingruby/make-it-survey/config"
	"github.com/charmingruby/make-it-survey/pkg/database/postgresql"
	"github.com/charmingruby/make-it-survey/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	db, err := postgresql.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	cfg.AssignDatabaseConn(db)

	logger := logger.SetupLogger()
	cfg.AssignLogger(logger)
}
