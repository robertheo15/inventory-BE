package service

import (
	"inventory-app-be/internal/models"
	"inventory-app-be/pkg/http/response"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateTransactionSupplier(ctx *gin.Context, newTransaction *response.TransactionSupplierRequest) (*response.TransactionSupplierResponse, error) {
	transactionSupplier, err := s.inventoryRepo.CreateTransactionSupplier(ctx, newTransaction)
	if err != nil {
		return nil, err
	}

	return transactionSupplier, nil
}

func (s *Service) GetTransactionSuppliers(ctx *gin.Context) ([]*models.Transaction, error) {
	transactionSuppliers, err := s.inventoryRepo.GetTransactionSupplier(ctx)
	if err != nil {
		return nil, err
	}

	return transactionSuppliers, nil
}

func (s *Service) GetTransactionSupplierByStatus(ctx *gin.Context, status string) ([]*models.Transaction, error) {
	transactions, err := s.inventoryRepo.GetTransactionSuppliersByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *Service) UpdateStatusTransactionByIDAndUpdateStock(ctx *gin.Context, id string) (string, error) {
	message, err := s.inventoryRepo.UpdateStatusTransactionByIDAndUpdateStock(ctx, id)
	if err != nil {
		return "", err
	}

	return message, nil
}
