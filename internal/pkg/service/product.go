package service

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/models"
)

func (s *Service) CreateProduct(product *models.Product) error {
	return s.storage.AddProduct(product)
}

func (s *Service) ListProducts() ([]models.Product, error) {
	return s.storage.Products()
}

func (s *Service) UpdateProduct(product *models.Product) error {
	return s.storage.UpdateProduct(product)
}

func (s *Service) DeleteProduct(product *models.Product) error {
	return s.storage.DeleteProduct(product.ID)
}


