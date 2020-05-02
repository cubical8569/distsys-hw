package server

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/handlers"
	"github.com/Azatik1000/distsys-hw/internal/pkg/service"
	"github.com/Azatik1000/distsys-hw/internal/pkg/storage"
	"github.com/go-chi/chi"
	"net/http"
)

type Server struct {
	storage storage.Storage
	handler *handlers.Handler
	router  chi.Router
}

func NewServer(storage storage.Storage) *Server {
	service := service.NewService(storage)
	handler := handlers.NewHandler(service)
	router := ServerRouter(handler)

	var server Server
	server = Server{
		router:  router,
	}

	return &server
}

func (s *Server) Run() {
	_ = http.ListenAndServe(":3333", s.router)
}

