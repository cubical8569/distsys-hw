package service

import "github.com/Azatik1000/distsys-hw/shop/storage"

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) *Service {
	return &Service{storage: storage}
}

