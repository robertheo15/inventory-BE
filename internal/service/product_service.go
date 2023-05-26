package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"inventory-app-be/internal/models"
)

func (s *Service) GetProducts(ctx *gin.Context) string {
	s.inventoryRepo.GetProducts()
	return "ok"
}

func (s *Service) CreateProduct(ctx *gin.Context) *models.User {
	user := &models.User{
		ID:       uuid.New(),
		FullName: "Robert",
		Email:    "Robertheo@gmail.com"}

	s.inventoryRepo.CreateProduct(user)
	return user
}
