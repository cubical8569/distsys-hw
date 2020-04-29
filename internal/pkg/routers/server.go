package routers

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func ServerRouter(service *handlers.Service) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(middleware.Heartbeat("/ping"))

	r.Route("/products", func(r chi.Router) {
		r.Method("GET",
			"/",
			handlers.NewBindRenderHandler(service.ListProducts))
		r.Method("POST",
			"/",
			handlers.NewBindRenderHandler(service.CreateProduct))

	})

	return r
}
