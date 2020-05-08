package server

import (
	"context"
	"github.com/Azatik1000/distsys-hw/internal/pkg/handlers"
	"github.com/Azatik1000/distsys-hw/internal/pkg/service"
	"github.com/Azatik1000/distsys-hw/internal/pkg/storage"
	"net/http"
)

type Server http.Server

func NewServer(storage storage.Storage) *Server {
	service := service.NewService(storage)
	handler := handlers.NewHandler(service)
	router := ServerRouter(handler)

	var server Server
	server = Server{
		Addr:    ":3333",
		Handler: router,
	}

	return &server
}

func (s *Server) Run() {
	_ = ((*http.Server)(s)).ListenAndServe()
}

func (s *Server) Shutdown() error {
	return (*http.Server)(s).Shutdown(context.Background())
}
