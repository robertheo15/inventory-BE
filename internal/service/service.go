package service

import repository "inventory-app-be/internal/repository/postgres/inventory"

type Service struct {
	inventoryRepo *repository.PostgresInventoryRepository
}

// NewService creates a new isntance of product service.
func NewService(inventoryRepo *repository.PostgresInventoryRepository) *Service {
	return &Service{
		inventoryRepo: inventoryRepo,
	}
}
