package main

import (
	"os"

	"github.com/charmingruby/upl/config"
	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/repository/postgres"
	"github.com/charmingruby/upl/internal/transport/rest"
	"github.com/charmingruby/upl/pkg/database/postgresql"
	"github.com/charmingruby/upl/pkg/logger"
	"github.com/gorilla/mux"
)

func main() {
	// Setup basics
	logger := logger.SetupLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error(err.Error())
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
	accountsRepository, err := postgres.NewAccountsRepository(cfg.Logger, cfg.Database.DatabaseConn)
	if err != nil {
		logger.Errorf("error initializing accounts postgres repository: %s", err.Error())
		os.Exit(1)
	}

	// Initialize services
	accountsService := accounts.NewAccountService(accountsRepository)

	// Initialize REST server
	router := mux.NewRouter().StrictSlash(true)

	rest.NewAccountHandler(cfg.Logger, accountsService).Register(router)
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
