package inventory

import (
	"fmt"
	"inventory-app-be/internal/models"
)

func (repo *PostgresInventoryRepository) GetProducts() []*models.Product {
	products := make([]*models.Product, 0)
	repo.db.Find(&products)
	return products
}

func (repo *PostgresInventoryRepository) CreateProduct(user *models.User) *models.User {

	err := repo.db.Model(&models.User{}).Create(&user).Error
	fmt.Println(err)
	//repo.db.Last(&product)
	return user
}
