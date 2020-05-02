package server

import (
	_ "github.com/Azatik1000/distsys-hw/docs"
	"github.com/Azatik1000/distsys-hw/internal/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
)

func ServerRouter(handler *handlers.Handler) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(middleware.Heartbeat("/ping"))

	r.Route("/v1/products", func(r chi.Router) {
		r.With(handler.PaginationCtx).Get("/", handler.ListProducts)
		r.Post("/", handler.CreateProduct)

		r.Route("/{productID}", func(r chi.Router) {
			r.Use(handler.ProductCtx)
			r.Get("/", handler.GetProduct)
			r.Put("/", handler.UpdateProduct)
			r.Delete("/", handler.DeleteProduct)
		})
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	return r
}
