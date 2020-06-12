package service

import (
	"github.com/Azatik1000/distsys-hw/shop/apimodels"
	"github.com/Azatik1000/distsys-hw/shop/storage"
)

func (s *Service) CreateProduct(
		request *apimodels.ProductRequest,
	) (*apimodels.ProductResponse, error) {
	product, err := s.storage.AddProduct(productRequestToProduct(request))

	if err != nil {
		return nil, err
	}

	return productToProductResponse(product), nil
}

func (s *Service) GetProduct(id uint) (*apimodels.ProductResponse, error) {
	product, err := s.storage.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return productToProductResponse(product), nil
}

type ListProductParams storage.GetParams

func (s *Service) ListProducts(params *ListProductParams) ([]*apimodels.ProductResponse, error) {
	products, err := s.storage.Products((*storage.GetParams)(params))
	if err != nil {
		return nil, err
	}


	response := []*apimodels.ProductResponse{}
	for _, product := range products {
		response = append(response, productToProductResponse(&product))
	}

	return response, nil
}

func (s *Service) UpdateProduct(
	id uint, request *apimodels.ProductRequest,
	) (*apimodels.ProductResponse, error) {
	product := productRequestToProduct(request)
	product.ID = id

	product, err := s.storage.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return productToProductResponse(product), nil
}

func (s *Service) DeleteProduct(id uint) error {
	return s.storage.DeleteProduct(id)
}
