package service

import (
	"inventory-app-be/internal/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateCustomer(ctx *gin.Context, newCustomer *models.Customer) (*models.Customer, error) {
	customer, err := s.inventoryRepo.CreateCustomer(ctx, newCustomer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *Service) GetCustomers(ctx *gin.Context) ([]*models.Customer, error) {
	customers, err := s.inventoryRepo.GetCustomers(ctx)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (s *Service) GetCustomerByID(ctx *gin.Context, id string) (*models.Customer, error) {
	customer, err := s.inventoryRepo.GetCustomerByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *Service) UpdateCustomerByID(ctx *gin.Context, newCustomer *models.Customer) (*models.Customer, error) {
	customer, err := s.inventoryRepo.UpdateCustomerByID(ctx, newCustomer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *Service) DeleteCustomerByID(ctx *gin.Context, id string) (string, error) {
	customerID, err := s.inventoryRepo.DeleteCustomerByID(ctx, id)
	if err != nil {
		return "", err
	}

	return customerID, nil
}
