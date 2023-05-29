package service

import (
	"inventory-app-be/internal/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetProducts(ctx *gin.Context) string {
	s.inventoryRepo.GetProducts()
	return "ok"
}

func (s *Service) CreateProduct(ctx *gin.Context) *models.User {
	// user, err := s.inventoryRepo.CreateProduct(ctx)
	// if err != nil {
	//	return nil
	//}
	return nil
}
