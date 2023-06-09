package service

import (
	"errors"
	"inventory-app-be/internal/middleware"
	"inventory-app-be/internal/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (s *Service) CreateUser(ctx *gin.Context, newUser *models.User) (*models.User, error) {
	hash, err := middleware.HashPassword(newUser.Password)
	if err != nil {
		return nil, err
	}

	newUser.Password = hash

	user, err := s.inventoryRepo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetUserByID(ctx *gin.Context, id string) (*models.User, error) {
	user, err := s.inventoryRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetUserByEmail(ctx *gin.Context, loginUser *models.User) (string, error) {
	user, err := s.inventoryRepo.GetUserByEmail(ctx, loginUser.Email)
	if err != nil {
		return "", err
	}

	comparePass := middleware.ComparePassword(user.Password, loginUser.Password)
	if !comparePass {
		return "", errors.New("email / password is not match")
	}

	token := middleware.GenerateToken(user)

	return token, nil
}

func (s *Service) UpdateUserByID(ctx *gin.Context, newUser *models.User) (*models.User, error) {
	user, err := s.inventoryRepo.UpdateUserByID(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) UpdatePasswordByID(ctx *gin.Context, newUser *models.User) (string, error) {
	user, err := s.inventoryRepo.GetUserByID(ctx, newUser.ID)
	if err != nil {
		return "", err
	}

	comparePass := middleware.ComparePassword(user.Password, newUser.OldPassword)
	if !comparePass {
		return "", errors.New("email / old password is not match")
	}

	newPassword, err := middleware.HashPassword(newUser.Password)
	if err != nil {
		return "", err
	}

	newUser.Password = newPassword

	message, err := s.inventoryRepo.UpdateUserPasswordByID(ctx, newUser)
	if err != nil {
		return "", err
	}

	return message, nil
}

func (s *Service) GetUserDetail(token string) (interface{}, error) {
	bearerToken := strings.Split(token, "Bearer ")[1]

	user, err := middleware.VerifyToken(bearerToken)
	if err != nil {
		return nil, err
	}

	return user.(jwt.Claims), nil
}

func (s *Service) DeActiveUserByID(ctx *gin.Context, newUser *models.User) (*models.User, error) {
	user, err := s.inventoryRepo.DeActiveUserByID(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) DeleteUserByID(ctx *gin.Context, id string) (string, error) {
	productID, err := s.inventoryRepo.DeleteUserByID(ctx, id)
	if err != nil {
		return "", err
	}

	return productID, nil
}
