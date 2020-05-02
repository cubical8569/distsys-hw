package handlers

import (
	"context"
	"github.com/Azatik1000/distsys-hw/internal/pkg/models"
	"github.com/Azatik1000/distsys-hw/internal/pkg/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

type ListProductsResponse struct {
	Products []models.Product `json:",inline"`
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, products)
}

type createProductRequest struct {
	models.Product
}

type createProductResponse struct {
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input createProductRequest
	err := render.Decode(r, &input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateProduct(&input.Product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	product := r.Context().Value("product").(*models.Product)
	render.Respond(w, r, product)
}

type updateProductRequest struct {
	models.Product
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	oldProduct := r.Context().Value("product").(*models.Product)
	err := h.service.UpdateProduct(oldProduct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var input updateProductRequest
	err = render.Decode(r, &input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newProduct := &input.Product
	newProduct.ID = oldProduct.ID
	err = h.service.UpdateProduct(newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	product := r.Context().Value("product").(*models.Product)
	err := h.service.DeleteProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ProductCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var product *models.Product
		var err error

		if articleIDArg := chi.URLParam(r, "productID"); articleIDArg != "" {
			var articleID int
			articleID, err = strconv.Atoi(articleIDArg)

			if err == nil {
				product, err = h.service.GetProduct(articleID)
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



