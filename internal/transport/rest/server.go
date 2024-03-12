package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/charmingruby/upl/internal/config"
	"github.com/gorilla/mux"
)

type Server struct {
	*http.Server
	Config *config.Config
}

func NewServer(cfg *config.Config, router *mux.Router, visibleRoutes bool) (*Server, error) {
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

	if visibleRoutes {
		router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			t, err := route.GetPathTemplate()
			if err != nil {
				return err
			}
			fmt.Println(t)
			return nil
		})
	}

	return &Server{
		Server: server,
		Config: cfg,
	}, nil
}

func (s *Server) Start() error {
	s.Config.Logger.Info("HTTP server is running.")

	if err := s.Server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
