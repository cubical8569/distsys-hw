package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name string
	Code string
	Kind ProductKind
}

func NewProduct(name string, code string, kind ProductKind) *Product {
	return &Product{Name: name, Code: code, Kind: kind}
}
