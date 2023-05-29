package inventory

import (
	"inventory-app-be/internal/models"

	"github.com/gin-gonic/gin"
)

func (repo *PostgresInventoryRepository) GetProducts() []*models.Product {
	products := make([]*models.Product, 0)
	return products
}

func (repo *PostgresInventoryRepository) CreateProduct(ctx *gin.Context) (*models.Product, error) {
	return nil, nil
}
