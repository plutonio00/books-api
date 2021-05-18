package server

import (
	"github.com/plutonio00/books-api/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(conf *config.Config) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: ":" + conf.ServerConfig.Port,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
