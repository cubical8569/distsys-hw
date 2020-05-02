package handlers

import (
	"context"
	"github.com/Azatik1000/distsys-hw/internal/pkg/models"
	"github.com/Azatik1000/distsys-hw/internal/pkg/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"net/url"
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
	limitArg := r.Context().Value("limit")
	offsetArg := r.Context().Value("offset")

	params := service.ListProductParams{}

	if limitArg != nil {
		limit := limitArg.(int)
		params.Limit = &limit
	}

	if offsetArg != nil {
		offset := offsetArg.(int)
		params.Offset = &offset
	}

	products, err := h.service.ListProducts(&params)

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

func ContextWithQueryArg(ctx context.Context, url *url.URL,
	argName string, argKey interface{},
	parseArg func(string) (int, error)) (context.Context, error) {
	argStr := url.Query().Get(argName)

	if argStr == "" {
		return ctx, nil
	}

	arg, err := parseArg(argStr)
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, argKey, arg), nil
}

func (h *Handler) PaginationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx, err := ContextWithQueryArg(ctx, r.URL, "limit", "limit", strconv.Atoi)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx, err = ContextWithQueryArg(ctx, r.URL, "offset", "offset", strconv.Atoi)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
