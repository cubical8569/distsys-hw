package server

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/handlers"
	"github.com/Azatik1000/distsys-hw/internal/pkg/routers"
	"github.com/Azatik1000/distsys-hw/internal/pkg/storage"
	"github.com/go-chi/chi"
	"net/http"
)

type server struct {
	service *handlers.Service
	router  chi.Router
}

func Server(storage storage.Storage) *server {
	service := handlers.NewService(storage)

	server := server{
		service: service,
		router:  routers.ServerRouter(service),
	}

	return &server
}

func (s *server) Run() {
	_ = http.ListenAndServe(":3333", s.router)
}