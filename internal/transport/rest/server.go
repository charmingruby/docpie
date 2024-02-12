package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/charmingruby/make-it-survey/config"
	"github.com/gorilla/mux"
)

type Server struct {
	*http.Server
	Config *config.Config
}

func NewServer(cfg *config.Config, router *mux.Router) (*Server, error) {
	if router == nil {
		return nil, fmt.Errorf("invalid server router")
	}

	serverCfg := cfg.Server

	addr := fmt.Sprintf("%s:%s", serverCfg.Host, serverCfg.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: router,

		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	return &Server{
		Server: server,
		Config: cfg,
	}, nil
}

func (s *Server) Start() error {
	s.Config.Logger.Infof("Server is running at %s", s.Server.Addr)

	if err := s.Server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
