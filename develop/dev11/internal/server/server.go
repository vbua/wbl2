package server

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(handler *http.Handler) *Server {
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: *handler,
	}

	s := &Server{
		server: httpServer,
	}

	return s
}

func (s *Server) Start() {
	err := s.server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *Server) Stop(ctx context.Context) {
	err := s.server.Shutdown(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
