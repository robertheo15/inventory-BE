package service

import (
	"inventory-app-be/internal/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateSupplier(ctx *gin.Context, newSupplier *models.Supplier) (*models.Supplier, error) {
	supplier, err := s.inventoryRepo.CreateSupplier(ctx, newSupplier)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (s *Service) GetSuppliers(ctx *gin.Context) ([]*models.Supplier, error) {
	suppliers, err := s.inventoryRepo.GetSuppliers(ctx)
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (s *Service) GetSupplierByID(ctx *gin.Context, id string) (*models.Supplier, error) {
	supplier, err := s.inventoryRepo.GetSupplierByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (s *Service) UpdateSupplierByID(ctx *gin.Context, newSupplier *models.Supplier) (*models.Supplier, error) {
	supplier, err := s.inventoryRepo.UpdateSupplierByID(ctx, newSupplier)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}

func (s *Service) DeleteSupplierByID(ctx *gin.Context, id string) (string, error) {
	supplierID, err := s.inventoryRepo.DeleteSupplierByID(ctx, id)
	if err != nil {
		return "", err
	}

	return supplierID, nil
}
