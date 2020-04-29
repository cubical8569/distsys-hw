package handlers

import "github.com/Azatik1000/distsys-hw/internal/pkg/storage"

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) *Service {
	return &Service{storage: storage}
}

