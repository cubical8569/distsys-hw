package storage

import "github.com/Azatik1000/distsys-hw/internal/pkg/models"

type GetParams struct {
	Limit  *int
	Offset *int
}

type Storage interface {
	AddProduct(product *models.Product) (*models.Product, error)
	GetProduct(id uint) (*models.Product, error)
	UpdateProduct(product *models.Product) (*models.Product, error)
	DeleteProduct(id uint) error
	Products(params *GetParams) ([]models.Product, error)
	Close() error
}
