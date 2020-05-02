package service

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/models"
	"github.com/Azatik1000/distsys-hw/internal/pkg/storage"
)

func (s *Service) CreateProduct(product *models.Product) error {
	return s.storage.AddProduct(product)
}

func (s *Service) GetProduct(id int) (*models.Product, error) {
	return s.storage.GetProduct(id)
}

type ListProductParams storage.GetParams

func (s *Service) ListProducts(params *ListProductParams) ([]models.Product, error) {
	return s.storage.Products((*storage.GetParams)(params))
}

func (s *Service) UpdateProduct(product *models.Product) error {
	return s.storage.UpdateProduct(product)
}

func (s *Service) DeleteProduct(product *models.Product) error {
	return s.storage.DeleteProduct(product.ID)
}


