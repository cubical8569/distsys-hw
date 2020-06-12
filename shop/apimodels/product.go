package apimodels

import "github.com/Azatik1000/distsys-hw/shop/models"

type ProductRequest struct {
	Name string             `json:"name"`
	Code string             `json:"code"`
	Kind models.ProductKind `json:"kind"`
}

type ProductResponse struct {
	Id             uint `json:"id"`
	ProductRequest `json:",inline"`
}
