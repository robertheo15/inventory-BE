package service

import (
	"inventory-app-be/internal/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateProduct(ctx *gin.Context, newProduct *models.Product) (*models.Product, error) {
	product, err := s.inventoryRepo.CreateProduct(ctx, newProduct)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) GetProducts(ctx *gin.Context) ([]*models.Product, error) {
	products, err := s.inventoryRepo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Service) GetProductByID(ctx *gin.Context, id string) (*models.Product, error) {
	product, err := s.inventoryRepo.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) GetProductBySupplierID(ctx *gin.Context, id string) ([]*models.Product, error) {
	products, err := s.inventoryRepo.GetProductsBySupplierID(ctx, id)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Service) UpdateProductByID(ctx *gin.Context, newProduct *models.Product) (*models.Product, error) {
	product, err := s.inventoryRepo.UpdateProductByID(ctx, newProduct)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) DeleteProductByID(ctx *gin.Context, id string) (string, error) {
	productID, err := s.inventoryRepo.DeleteProductByID(ctx, id)
	if err != nil {
		return "", err
	}

	return productID, nil
}
