package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func ServerRouter(server *Server) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(middleware.Heartbeat("/ping"))

	r.Route("/products", func(r chi.Router) {
		r.Get("/", server.ListProducts)
		r.Post("/", server.CreateProduct)
	})

	return r
}
