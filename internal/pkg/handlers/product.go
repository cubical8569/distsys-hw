package handlers

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/models"
)

type CreateProductRequest struct {
	models.Product
}

type CreateProductResponse struct {
}

func (s *Service) CreateProduct(input *CreateProductRequest) (*CreateProductResponse, error) {
	err := s.storage.AddProduct(&input.Product)
	if err != nil {
		return nil, err
	}

	return &CreateProductResponse{}, nil
}

type ListProductsResponse struct {
	Products []models.Product `json:",inline"`
}

func (s *Service) ListProducts() (*ListProductsResponse, error) {
	products, err := s.storage.Products()
	if err != nil {
		return nil, err
	}

	return &ListProductsResponse{Products: products}, nil
}


