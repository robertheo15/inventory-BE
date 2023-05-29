package service

import (
	"inventory-app-be/internal/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateUser(ctx *gin.Context) (*models.User, error) {
	user, err := s.inventoryRepo.CreateUser(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetUserByID(ctx *gin.Context) (*models.User, error) {
	user, err := s.inventoryRepo.GetUserByID(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
