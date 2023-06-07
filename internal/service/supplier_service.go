package service

import (
	"github.com/gin-gonic/gin"
	"inventory-app-be/internal/models"
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
		return "nil", err
	}

	return supplierID, nil
}
