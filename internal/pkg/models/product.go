package models

type Product struct {
	Name string      `json:"name"`
	Code string      `json:"code"`
	Type productType `json:"type"`
}
