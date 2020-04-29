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


