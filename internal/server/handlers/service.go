package handlers

import "github.com/garrettladley/prods/internal/storage"

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}
