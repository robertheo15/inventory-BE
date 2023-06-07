package service

import (
	"inventory-app-be/internal/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateProductVariant(ctx *gin.Context, newProductVariant *models.ProductVariant) (*models.ProductVariant, error) {
	productVariant, err := s.inventoryRepo.CreateProductVariant(ctx, newProductVariant)
	if err != nil {
		return nil, err
	}

	return productVariant, nil
}

func (s *Service) GetProductVariants(ctx *gin.Context) ([]*models.ProductVariant, error) {
	productVariant, err := s.inventoryRepo.GetProductVariants(ctx)
	if err != nil {
		return nil, err
	}

	return productVariant, nil
}

func (s *Service) GetProductVariantByID(ctx *gin.Context, id string) (*models.ProductVariant, error) {
	productVariant, err := s.inventoryRepo.GetProductVariantByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return productVariant, nil
}

func (s *Service) GetProductVariantsByProductID(ctx *gin.Context, id string) ([]*models.ProductVariant, error) {
	productVariant, err := s.inventoryRepo.GetProductVariantsByProductID(ctx, id)
	if err != nil {
		return nil, err
	}

	return productVariant, nil
}

func (s *Service) UpdateProductVariantByID(ctx *gin.Context, newProductVariant *models.ProductVariant) (*models.ProductVariant, error) {
	productVariant, err := s.inventoryRepo.UpdateProductVariantByID(ctx, newProductVariant)
	if err != nil {
		return nil, err
	}

	return productVariant, nil
}

func (s *Service) DeleteProductVariantByID(ctx *gin.Context, id string) (string, error) {
	productID, err := s.inventoryRepo.DeleteProductVariantByID(ctx, id)
	if err != nil {
		return "", err
	}

	return productID, nil
}
