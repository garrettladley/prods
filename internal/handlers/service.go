package handlers

import (
	"github.com/garrettladley/prods/internal/algo"
	"github.com/garrettladley/prods/internal/constants"
	"github.com/garrettladley/prods/internal/storage"
)

type Service struct {
	storage storage.Storage
	algo    *algo.Service
}

func NewService(storage storage.Storage) *Service {
	return &Service{
		storage: storage,
		algo:    algo.NewService(constants.ProductSubset),
	}
}
