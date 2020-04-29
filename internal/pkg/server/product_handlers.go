package server

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/models"
	"github.com/go-chi/render"
	"net/http"
)

type ListProductsResponse struct {
	Products []models.Product `json:",inline"`
}

func (s *Server) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := s.service.ListProducts()

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

func (s *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input createProductRequest
	err := render.Decode(r, &input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.service.CreateProduct(&input.Product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	product := r.Context().Value("product").(*models.Product)
	render.Respond(w, r, product)
}

type updateProductRequest struct {
	models.Product
}

func (s *Server) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	oldProduct := r.Context().Value("product").(*models.Product)
	err := s.service.UpdateProduct(oldProduct)

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
	err = s.service.UpdateProduct(newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	product := r.Context().Value("product").(*models.Product)
	err := s.service.DeleteProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
