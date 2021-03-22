package server

import (
	"context"
	"net/http"

	"github.com/Valeriy-Totubalin/test_project/internal/app/interfaces/config_interfaces"
)

type Server struct {
	httpServer *http.Server
	Config     config_interfaces.ServerConfig
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + s.Config.GetPort(),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 mb
		ReadTimeout:    s.Config.GetReadTimeout(),
		WriteTimeout:   s.Config.GetWriteTimeout(),
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
