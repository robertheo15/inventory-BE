package inventory

import (
	"inventory-app-be/internal/models"
	"inventory-app-be/internal/repository/postgres/sqlc"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func (repo *PostgresInventoryRepository) CreateSupplier(ctx *gin.Context, newSupplier *models.Supplier) (*models.Supplier, error) {
	supplier := sqlc.CreateSupplierParams{
		BrandName:   newSupplier.BrandName,
		PhoneNumber: newSupplier.Phone,
		Email:       newSupplier.Email,
		Address:     newSupplier.Address,
		CreatedBy:   newSupplier.CreatedBy,
		UpdatedBy:   newSupplier.UpdatedBy,
	}

	ID, err := repo.db.CreateSupplier(ctx, supplier)
	if err != nil {
		log.Printf("Supplier Repository: %s", err)

		return nil, err
	}

	resultCustomer := &models.Supplier{
		ID:        ID,
		BrandName: supplier.BrandName,
		Email:     supplier.Email,
		Phone:     supplier.PhoneNumber,
		Address:   supplier.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: supplier.CreatedBy,
		UpdatedBy: supplier.UpdatedBy,
	}

	return resultCustomer, nil
}

func (repo *PostgresInventoryRepository) GetSuppliers(ctx *gin.Context) ([]*models.Supplier, error) {
	suppliers := make([]*models.Supplier, 0)

	newSuppliers, err := repo.db.GetSuppliers(ctx)
	if err != nil {
		log.Printf("Supplier Repository: %s", err)

		return nil, err
	}

	for _, supplier := range newSuppliers {
		newSupplier := &models.Supplier{
			ID:        supplier.ID,
			BrandName: supplier.BrandName,
			Phone:     supplier.PhoneNumber,
			Address:   supplier.Address,
			Email:     supplier.Email,
			CreatedAt: supplier.CreatedAt,
			UpdatedAt: supplier.UpdatedAt,
			CreatedBy: supplier.CreatedBy,
			UpdatedBy: supplier.UpdatedBy,
		}
		suppliers = append(suppliers, newSupplier)
	}

	return suppliers, nil
}

func (repo *PostgresInventoryRepository) GetSupplierByID(ctx *gin.Context, id string) (*models.Supplier, error) {
	supplier, err := repo.db.GetSupplierByID(ctx, id)
	if err != nil {
		log.Printf("Supplier Repository: %s", err)

		return nil, err
	}

	newSupplier := &models.Supplier{
		ID:        supplier.ID,
		BrandName: supplier.BrandName,
		Email:     supplier.Email,
		Phone:     supplier.PhoneNumber,
		Address:   supplier.Address,
		CreatedBy: supplier.CreatedBy,
		UpdatedBy: supplier.UpdatedBy,
	}

	return newSupplier, nil
}

func (repo *PostgresInventoryRepository) UpdateSupplierByID(ctx *gin.Context,
	newSupplier *models.Supplier) (*models.Supplier, error) {
	supplier, err := repo.db.UpdateSupplierByID(ctx, sqlc.UpdateSupplierByIDParams{
		ID:          newSupplier.ID,
		BrandName:   newSupplier.BrandName,
		PhoneNumber: newSupplier.Phone,
		Email:       newSupplier.Email,
		Address:     newSupplier.Address,
		UpdatedBy:   newSupplier.UpdatedBy,
	})
	if err != nil {
		log.Printf("Supplier Repository: %s", err)

		return nil, err
	}

	resultCustomer := &models.Supplier{
		ID:        supplier.ID,
		BrandName: supplier.BrandName,
		Phone:     supplier.PhoneNumber,
		Email:     supplier.Email,
		Address:   supplier.Address,
		CreatedAt: newSupplier.CreatedAt,
		UpdatedAt: time.Now(),
		CreatedBy: supplier.CreatedBy,
		UpdatedBy: supplier.UpdatedBy,
	}

	return resultCustomer, nil
}

func (repo *PostgresInventoryRepository) DeleteSupplierByID(ctx *gin.Context, id string) (string, error) {
	supplierID, err := repo.db.DeleteSupplierByID(ctx, id)
	if err != nil {
		log.Printf("Supplier Repository: %s", err)

		return "", err
	}

	return supplierID, nil
}
