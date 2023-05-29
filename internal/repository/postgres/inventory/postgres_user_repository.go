package inventory

import (
	"inventory-app-be/internal/models"
	"inventory-app-be/internal/repository/postgres/sqlc"
	"log"

	"github.com/gin-gonic/gin"
)

func (repo *PostgresInventoryRepository) CreateUser(ctx *gin.Context) (*models.User, error) {
	user, err := repo.db.CreateUser(ctx, sqlc.CreateUserParams{
		FullName: "robert",
	})
	if err != nil {
		log.Fatal("Create user repository error: ", err)

		return nil, err
	}

	newUser := &models.User{
		ID:       user.ID,
		FullName: user.FullName.String,
	}

	return newUser, nil
}

func (repo *PostgresInventoryRepository) GetUserByID(ctx *gin.Context) (*models.User, error) {
	user, err := repo.db.GetUserByID(ctx, "d0f11f6c-1721-4f56-befe-cc530e47ce0f")
	if err != nil {
		log.Fatal("Get user by id repository error: ", err)

		return nil, err
	}

	newUser := &models.User{
		ID:       user.ID,
		FullName: user.FullName.String,
	}

	return newUser, nil
}
