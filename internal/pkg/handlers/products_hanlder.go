package handlers

import (
	"context"
	"fmt"
	"github.com/Azatik1000/distsys-hw/internal/pkg/apimodels"
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

// @Summary List all products
// @Description List all products
// @Tags products
// @Accept json
// @Produce json
// @Param: limit query int false "limit"
// @Param: offset query int false "offset"
// @Success 200 {array} apimodels.ProductResponse
// @Router /v1/products [get]
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

// @Summary Create a product
// @Description Create a product
// @Tags products
// @Accept json
// @Produce json
// @Param product body apimodels.ProductRequest true "product"
// @Success 200 {object} apimodels.ProductResponse
// @Router /v1/products [post]
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input apimodels.ProductRequest
	err := render.Decode(r, &input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("input:", input)
	response, err := h.service.CreateProduct(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, response)
}

// @Summary Get a product with provided id
// @Description Get a product with provided id
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "product id"
// @Success 200 {object} apimodels.ProductResponse
// @Router /v1/products/{id} [get]
func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("productID").(uint)
	response, _ := h.service.GetProduct(id)
	render.Respond(w, r, response)
}

// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "product id"
// @Param product body apimodels.ProductRequest true "product"
// @Success 200 {object} apimodels.ProductResponse
// @Router /v1/products/{id} [put]
func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("productID").(uint)

	var input apimodels.ProductRequest
	err := render.Decode(r, &input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.service.UpdateProduct(id, &input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, response)
}

// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "product id"
// @Success 200
// @Router /v1/products/{id} [delete]
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("productID").(uint)
	err := h.service.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ProductCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var productID uint
		var err error

		if productIDArg := chi.URLParam(r, "productID"); productIDArg != "" {
			var tmp uint64
			tmp, err = strconv.ParseUint(productIDArg, 10, 64)

			// TODO: make safe conversion
			if err == nil {
				productID = uint(tmp)
				_, err = h.service.GetProduct(productID)
			}
		} else {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "productID", productID)
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
