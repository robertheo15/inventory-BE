package service

import (
	"inventory-app-be/internal/models"
	"inventory-app-be/pkg/http/response"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateTransaction(ctx *gin.Context, newTransaction *response.TransactionRequest) (*response.TransactionResponse, error) {
	transaction, err := s.inventoryRepo.CreateTransaction(ctx, newTransaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *Service) GetTransactions(ctx *gin.Context) ([]*models.Transaction, error) {
	transactions, err := s.inventoryRepo.GetTransactions(ctx)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *Service) GetTransactionByID(ctx *gin.Context, id string) (*models.Transaction, error) {
	transaction, err := s.inventoryRepo.GetTransactionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *Service) UpdateTransactionByID(ctx *gin.Context, newTransaction *models.Transaction) (*models.Transaction, error) {
	transaction, err := s.inventoryRepo.UpdateTransactionByID(ctx, newTransaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *Service) DeleteTransactionByID(ctx *gin.Context, id string) (string, error) {
	transactionID, err := s.inventoryRepo.DeleteTransactionByID(ctx, id)
	if err != nil {
		return "", err
	}

	return transactionID, nil
}
