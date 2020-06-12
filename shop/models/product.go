package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name string
	Code string
	Kind ProductKind
}

func (p *Product) Equal(other *Product) bool {
	return p.ID == other.ID &&
		p.Name == other.Name &&
		p.Code == other.Code &&
		p.Kind == other.Kind
}

func NewProduct(name string, code string, kind ProductKind) *Product {
	return &Product{Name: name, Code: code, Kind: kind}
}
