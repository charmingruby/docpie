package main

import (
	"os"

	"github.com/charmingruby/owler/config"
	"github.com/charmingruby/owler/internal/transport/rest"
	"github.com/charmingruby/owler/pkg/database/postgresql"
	"github.com/charmingruby/owler/pkg/logger"
	"github.com/gorilla/mux"
)

func main() {
	// Setup basics
	logger := logger.SetupLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("error loading environment configuration")
		os.Exit(1)
	}
	cfg.AssignLogger(logger)

	db, err := postgresql.LoadDatabase(cfg)
	if err != nil {
		logger.Errorf("error connecting to postgres: %s", err.Error())
		os.Exit(1)
	}
	cfg.AssignDatabaseConn(db)

	// Initialize repos

	// Initialize services

	// Initialize REST server
	router := mux.NewRouter().StrictSlash(true)

	rest.NewPingHandler().Register(router)

	restServer, err := rest.NewServer(cfg, router, true)
	if err != nil {
		logger.Errorf("error instantiating server: %s", err.Error())
		os.Exit(1)
	}

	if err := restServer.Start(); err != nil {
		logger.Errorf("error starting server: %s", err.Error())
		os.Exit(1)
	}
}
