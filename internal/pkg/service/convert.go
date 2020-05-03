package service

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/apimodels"
	"github.com/Azatik1000/distsys-hw/internal/pkg/models"
)

func productRequestToProduct(request *apimodels.ProductRequest) *models.Product {
	return models.NewProduct(request.Name, request.Code, request.Kind)
}

func productToProductResponse(product *models.Product) *apimodels.ProductResponse {
	return &apimodels.ProductResponse{
		Id:             product.ID,
		ProductRequest: apimodels.ProductRequest{
			product.Name,
			product.Code,
			product.Kind,
		},
	}
}
