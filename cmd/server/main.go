package main

import (
	"os"

	"github.com/charmingruby/upl/config"
	"github.com/charmingruby/upl/internal/database/postgres"
	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/transport/rest"
	"github.com/charmingruby/upl/internal/transport/rest/middlewares"
	"github.com/charmingruby/upl/pkg/database/postgresql"
	"github.com/charmingruby/upl/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Setup
	logger := logger.SetupLogger()

	if err := godotenv.Load(); err != nil {
		logger.Info(".env file not found")
	}

	cfg, err := config.LoadConfig(logger)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	cfg.AssignLogger(logger)

	db, err := postgresql.LoadDatabase(cfg)
	if err != nil {
		logger.Errorf("Error connecting to postgres: %s", err.Error())
		os.Exit(1)
	}
	cfg.AssignDatabaseConn(db)

	// Initialize repos
	logger.Info("Initializing repositories...")
	accountsRepository, err := postgres.NewAccountsRepository(cfg.Logger, cfg.Database.DatabaseConn)
	if err != nil {
		logger.Errorf("Error initializing accounts postgres repository: %s", err.Error())
		os.Exit(1)
	}

	collectionTagsRepository, err := postgres.NewCollectionTagsRepository(cfg.Logger, cfg.Database.DatabaseConn)
	if err != nil {
		logger.Errorf("Error initializing collection tags postgres repository: %s", err.Error())
		os.Exit(1)
	}

	collectionsRepository, err := postgres.NewCollectionsRepository(cfg.Logger, cfg.Database.DatabaseConn)
	if err != nil {
		logger.Errorf("Error initializing collections postgres repository: %s", err.Error())
		os.Exit(1)
	}

	collectionMembersRepository, err := postgres.NewCollectionMembersRepository(cfg.Logger, cfg.Database.DatabaseConn)
	if err != nil {
		logger.Errorf("Error initializing collection members postgres repository: %s", err.Error())
		os.Exit(1)
	}

	uploadsRepository, err := postgres.NewUploadsRepository(cfg.Logger, cfg.Database.DatabaseConn)
	if err != nil {
		logger.Errorf("Error initializing uploads postgres repository: %s", err.Error())
		os.Exit(1)
	}
	logger.Info("Repositories initialized.")

	// Initialize services
	logger.Info("Initializing services...")
	accountsService := accounts.NewAccountService(accountsRepository)
	collectionTagsService := collections.NewCollectionTagsService(collectionTagsRepository)
	collectionsService := collections.NewCollectionService(collectionsRepository, collectionTagsRepository, collectionMembersRepository, accountsRepository)
	collectionMembersService := collections.NewCollectionsMembersService(collectionMembersRepository, accountsRepository, collectionsRepository)
	uploadsService := collections.NewUploadService(uploadsRepository, collectionsRepository, accountsRepository, collectionMembersRepository)
	logger.Info("Services initialized.")

	// Initialize REST server
	logger.Info("Initializing HTTP server...")
	router := mux.NewRouter().StrictSlash(true)
	middlewares := middlewares.NewMiddleware(cfg.Logger, collectionMembersRepository, collectionsRepository)

	// Initialize the routes
	logger.Info("Registering routes...")
	rest.NewPingHandler().Register(router)
	rest.NewAccountsHandler(cfg.Logger, middlewares, accountsService).Register(router)
	rest.NewCollectionsHandler(cfg.Logger, middlewares, collectionsService, collectionTagsService, collectionMembersService, uploadsService).Register(router)
	logger.Info("Routes registered.")

	// Initialize the server
	visibleRoutes := false
	restServer, err := rest.NewServer(cfg, router, visibleRoutes)
	if err != nil {
		logger.Errorf("error instantiating server: %s", err.Error())
		os.Exit(1)
	}

	if err := restServer.Start(); err != nil {
		logger.Errorf("error starting server: %s", err.Error())
		os.Exit(1)
	}
}
