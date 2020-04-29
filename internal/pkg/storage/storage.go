package storage

import "github.com/Azatik1000/distsys-hw/internal/pkg/models"

type Storage interface {
	AddProduct(product *models.Product) error
	Products() ([]models.Product, error)
}
