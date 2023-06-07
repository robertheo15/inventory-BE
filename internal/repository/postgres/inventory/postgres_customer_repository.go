package inventory

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inventory-app-be/internal/models"
	"inventory-app-be/internal/repository/postgres/sqlc"
	"log"
	"time"
)

func (repo *PostgresInventoryRepository) CreateCustomer(ctx *gin.Context, newCustomer *models.Customer) (*models.Customer, error) {
	customer := sqlc.CreateCustomerParams{
		FullName:    newCustomer.FullName,
		PhoneNumber: newCustomer.Phone,
		Address:     newCustomer.Address,
		CreatedBy:   newCustomer.CreatedBy,
		UpdatedBy:   newCustomer.UpdatedBy,
	}

	ID, err := repo.db.CreateCustomer(ctx, customer)
	if err != nil {
		log.Printf(fmt.Sprintf("Product Repository: %s", err))

		return nil, err
	}

	resultCustomer := &models.Customer{
		ID:        ID,
		FullName:  customer.FullName,
		Phone:     customer.PhoneNumber,
		Address:   customer.Address,
		CreatedBy: customer.CreatedBy,
		UpdatedBy: customer.UpdatedBy,
	}

	return resultCustomer, nil
}

func (repo *PostgresInventoryRepository) GetCustomers(ctx *gin.Context) ([]*models.Customer, error) {
	customers := make([]*models.Customer, 0)

	newCustomers, err := repo.db.GetCustomers(ctx)
	if err != nil {
		log.Println(fmt.Sprintf("Product Repository: %s", err))

		return nil, err
	}

	for _, customer := range newCustomers {
		newCustomer := &models.Customer{
			ID:        customer.ID,
			FullName:  customer.FullName,
			Phone:     customer.PhoneNumber,
			Address:   customer.Address,
			CreatedBy: customer.CreatedBy,
			UpdatedBy: customer.UpdatedBy,
		}
		customers = append(customers, newCustomer)
	}

	return customers, nil
}

func (repo *PostgresInventoryRepository) GetCustomerByID(ctx *gin.Context, id string) (*models.Customer, error) {
	customer, err := repo.db.GetCustomersByID(ctx, id)
	if err != nil {
		log.Printf(fmt.Sprintf("Product Repository: %s", err))

		return nil, err
	}

	newCustomer := &models.Customer{
		ID:        customer.ID,
		FullName:  customer.FullName,
		Phone:     customer.PhoneNumber,
		Address:   customer.Address,
		CreatedBy: customer.CreatedBy,
		UpdatedBy: customer.UpdatedBy,
	}

	return newCustomer, nil
}

func (repo *PostgresInventoryRepository) UpdateCustomerByID(ctx *gin.Context,
	newCustomer *models.Customer) (*models.Customer, error) {
	customer, err := repo.db.UpdateCustomerByID(ctx, sqlc.UpdateCustomerByIDParams{
		ID:        newCustomer.ID,
		FullName:  newCustomer.FullName,
		Address:   newCustomer.Address,
		CreatedAt: newCustomer.CreatedAt,
		UpdatedBy: newCustomer.UpdatedBy,
	})
	if err != nil {
		log.Printf(fmt.Sprintf("Product Repository: %s", err))

		return nil, err
	}

	resultCustomer := &models.Customer{
		ID:        customer.ID,
		FullName:  customer.FullName,
		Phone:     customer.PhoneNumber,
		Address:   customer.Address,
		CreatedAt: newCustomer.CreatedAt,
		UpdatedAt: time.Now(),
		CreatedBy: customer.CreatedBy,
		UpdatedBy: customer.UpdatedBy,
	}

	return resultCustomer, nil
}

func (repo *PostgresInventoryRepository) DeleteCustomerByID(ctx *gin.Context, id string) (string, error) {
	customerID, err := repo.db.DeleteCustomerByID(ctx, id)
	if err != nil {
		log.Printf(fmt.Sprintf("Product Repository: %s", err))

		return "", err
	}

	return customerID, nil
}
