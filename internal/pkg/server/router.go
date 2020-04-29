package server

import (
	"context"
	"github.com/Azatik1000/distsys-hw/internal/pkg/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
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

		r.Route("/{productID}", func(r chi.Router) {
			r.Use(server.ProductCtx)
			r.Get("/", server.GetProduct)
		})
	})

	return r
}

func (s *Server) ProductCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var product *models.Product
		var err error

		if articleIDArg := chi.URLParam(r, "productID"); articleIDArg != "" {
			var articleID int
			articleID, err = strconv.Atoi(articleIDArg)

			if err == nil {
				product, err = s.storage.GetProduct(articleID)
			}
		} else {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "product", product)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
