package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name string      `json:"name"`
	Code string      `json:"code"`
	Type productType `json:"type"`
}
