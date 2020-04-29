package server

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/service"
	"github.com/Azatik1000/distsys-hw/internal/pkg/storage"
	"github.com/go-chi/chi"
	"net/http"
)

type Server struct {
	storage storage.Storage
	service *service.Service
	router  chi.Router
}

func NewServer(storage storage.Storage) *Server {
	service := service.NewService(storage)

	var server Server
	server = Server{
		storage: storage,
		service: service,
		router:  ServerRouter(&server),
	}

	return &server
}

func (s *Server) Run() {
	_ = http.ListenAndServe(":3333", s.router)
}

