package storage

import "github.com/Azatik1000/distsys-hw/internal/pkg/models"

type Memory struct {
	products []models.Product
}

func (m *Memory) AddProduct(product *models.Product) error {
	m.products = append(m.products, *product)
	return nil
}

func (m *Memory) Products() ([]models.Product, error) {
	return m.products, nil
}
