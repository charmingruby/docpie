package main

import (
	"os"

	"github.com/charmingruby/upl/config"
	"github.com/charmingruby/upl/internal/database/postgres"
	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/transport/rest"
	"github.com/charmingruby/upl/pkg/database/postgresql"
	"github.com/charmingruby/upl/pkg/logger"
	"github.com/gorilla/mux"
)

func main() {
	// Setup
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

	collectionTagsRepository, err := postgres.NewCollectionTagsRepository(cfg.Logger, cfg.Database.DatabaseConn)
	if err != nil {
		logger.Errorf("error initializing collection tags postgres repository: %s", err.Error())
		os.Exit(1)
	}

	collectionsRepository, err := postgres.NewCollectionsRepository(cfg.Logger, cfg.Database.DatabaseConn)
	if err != nil {
		logger.Errorf("error initializing collection postgres repository: %s", err.Error())
		os.Exit(1)
	}

	collectionMembersRepository, err := postgres.NewCollectionMembersRepository(cfg.Logger, cfg.Database.DatabaseConn)
	if err != nil {
		logger.Errorf("error initializing collection postgres repository: %s", err.Error())
		os.Exit(1)
	}

	// Initialize services
	accountsService := accounts.NewAccountService(accountsRepository)
	collectionTagsService := collections.NewCollectionTagsService(collectionTagsRepository)
	collectionsService := collections.NewCollectionService(collectionsRepository, collectionTagsRepository, collectionMembersRepository, accountsRepository)

	// Initialize REST server
	router := mux.NewRouter().StrictSlash(true)

	// Initialize the routes
	rest.NewAccountsHandler(cfg.Logger, accountsService).Register(router)
	rest.NewCollectionsHandler(cfg.Logger, collectionsService, collectionTagsService).Register(router)
	rest.NewPingHandler().Register(router)

	// Initialize the server
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
