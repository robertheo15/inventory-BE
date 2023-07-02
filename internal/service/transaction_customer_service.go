package service

import (
	"inventory-app-be/internal/models"
	"inventory-app-be/pkg/http/response"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateTransaction(ctx *gin.Context, newTransaction *response.TransactionCustomerRequest) (*response.TransactionCustomerResponse, error) {
	transaction, err := s.inventoryRepo.CreateTransactionCustomer(ctx, newTransaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *Service) GetTransactionCustomers(ctx *gin.Context) ([]*models.Transaction, error) {
	transactions, err := s.inventoryRepo.GetTransactionCustomers(ctx)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *Service) GetTransactionCustomersByStatus(ctx *gin.Context, status string) ([]*models.Transaction, error) {
	transactions, err := s.inventoryRepo.GetTransactionCustomersByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *Service) GetTransactionByID(ctx *gin.Context, id string) (*models.Transaction, error) {
	transaction, err := s.inventoryRepo.GetTransactionCustomerParentByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *Service) UpdateTransactionByID(ctx *gin.Context, newTransaction *models.Transaction) (*models.Transaction, error) {
	transaction, err := s.inventoryRepo.UpdateTransactionCustomerByID(ctx, newTransaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *Service) UpdateStatusTransactionSedangDikirimByID(ctx *gin.Context, id string) (string, error) {
	message, err := s.inventoryRepo.UpdateStatusTransactionCSedangDikirimByID(ctx, id)
	if err != nil {
		return "", err
	}

	return message, nil

}

func (s *Service) UpdateStatusTransactionSelesaiByID(ctx *gin.Context, id string) (string, error) {
	message, err := s.inventoryRepo.UpdateStatusTransactionCSelesaiByID(ctx, id)
	if err != nil {
		return "", err
	}

	return message, nil
}

func (s *Service) DeleteTransactionByID(ctx *gin.Context, id string) (string, error) {
	transactionID, err := s.inventoryRepo.DeleteTransactionCustomerByID(ctx, id)
	if err != nil {
		return "", err
	}

	return transactionID, nil
}
