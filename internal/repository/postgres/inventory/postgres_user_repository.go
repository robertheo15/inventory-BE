package inventory

import (
	"inventory-app-be/internal/models"
	"inventory-app-be/internal/repository/postgres/sqlc"
	pkgHttp "inventory-app-be/pkg/http"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func (repo *PostgresInventoryRepository) CreateUser(ctx *gin.Context, newUser *models.User) (*models.User, error) {
	userSqlcParam := sqlc.CreateUserParams{
		FullName:    newUser.FullName,
		Password:    newUser.Password,
		PhoneNumber: newUser.PhoneNumber,
		Email:       newUser.Email,
		Role:        newUser.Role,
		Active:      newUser.Active,
		CreatedBy:   ctx.GetString("full_name"),
		UpdatedBy:   ctx.GetString("full_name"),
	}

	userID, err := repo.db.CreateUser(ctx, userSqlcParam)
	if err != nil {
		log.Printf("Create user repository error: %s\n", err)

		return nil, err
	}

	resultUser := &models.User{
		ID:          userID,
		FullName:    newUser.FullName,
		PhoneNumber: newUser.PhoneNumber,
		Email:       newUser.Email,
		Role:        newUser.Role,
		Active:      newUser.Active,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   newUser.CreatedBy,
		UpdatedBy:   newUser.UpdatedBy,
	}

	return resultUser, nil
}

func (repo *PostgresInventoryRepository) GetUsers(ctx *gin.Context) ([]*models.User, error) {
	users := make([]*models.User, 0)
	resUsers, err := repo.db.GetUsers(ctx)
	if err != nil {
		log.Printf("Get user by id repository error: %s", err)

		return nil, err
	}

	for _, resUser := range resUsers {
		user := &models.User{
			FullName:    resUser.FullName,
			PhoneNumber: resUser.PhoneNumber,
			Email:       resUser.Email,
			Role:        resUser.Role,
			Active:      resUser.Active,
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *PostgresInventoryRepository) GetUserByID(ctx *gin.Context, id string) (*models.User, error) {
	user, err := repo.db.GetUserByID(ctx, id)
	if err != nil {
		log.Printf("Get user by id repository error: %s", err)

		return nil, err
	}

	newUser := &models.User{
		ID:       user.ID,
		Password: user.Password,
		FullName: user.FullName,
	}

	return newUser, nil
}

func (repo *PostgresInventoryRepository) GetUserByEmail(ctx *gin.Context, email string) (*models.User, error) {
	user, err := repo.db.GetUserByEmail(ctx, email)
	if err != nil {
		log.Printf("Get user by id repository error: %s", err)

		return nil, err
	}

	newUser := &models.User{
		ID:          user.ID,
		FullName:    user.FullName,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Role:        user.Role,
		Active:      user.Active,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		CreatedBy:   user.CreatedBy,
		UpdatedBy:   user.UpdatedBy,
	}

	return newUser, nil
}

func (repo *PostgresInventoryRepository) UpdateUserByID(ctx *gin.Context, newUser *models.User) (*models.User, error) {
	userSqlcUpdateParams := sqlc.UpdateUserByIDParams{
		ID:          newUser.ID,
		FullName:    newUser.FullName,
		PhoneNumber: newUser.PhoneNumber,
		Email:       newUser.Email,
		Role:        newUser.Role,
		Active:      newUser.Active,
		CreatedBy:   ctx.GetString("full_name"),
		UpdatedBy:   ctx.GetString("full_name"),
	}

	userID, err := repo.db.UpdateUserByID(ctx, userSqlcUpdateParams)
	if err != nil {
		log.Printf("User repository update user by id: %s", err)

		return nil, err
	}

	resultNewUser := &models.User{
		ID:          userID,
		FullName:    newUser.FullName,
		PhoneNumber: newUser.PhoneNumber,
		Email:       newUser.Email,
		Role:        newUser.Role,
		Active:      newUser.Active,
		CreatedAt:   newUser.CreatedAt,
		UpdatedAt:   time.Now(),
		CreatedBy:   newUser.CreatedBy,
		UpdatedBy:   newUser.UpdatedBy,
	}

	return resultNewUser, nil
}

func (repo *PostgresInventoryRepository) UpdateUserPasswordByID(ctx *gin.Context, newUser *models.User) (string, error) {
	_, err := repo.db.UpdateUserPasswordByID(ctx, sqlc.UpdateUserPasswordByIDParams{
		ID:       newUser.ID,
		Password: newUser.Password,
	})
	if err != nil {
		log.Printf("User repository update password by id: %s", err)

		return "", err
	}

	return pkgHttp.PasswordChanged, nil
}

func (repo *PostgresInventoryRepository) DeActiveUserByID(ctx *gin.Context, newUser *models.User) (*models.User, error) {
	userSqlcUpdateParams := sqlc.UpdateUserByIDParams{
		FullName:    newUser.FullName,
		PhoneNumber: newUser.PhoneNumber,
		Email:       newUser.Email,
		Role:        newUser.Role,
		Active:      0,
		CreatedBy:   newUser.CreatedBy,
		UpdatedBy:   newUser.UpdatedBy,
	}

	userID, err := repo.db.UpdateUserByID(ctx, userSqlcUpdateParams)
	if err != nil {
		log.Printf("Get user by id repository error: %s", err)

		return nil, err
	}

	resultNewUser := &models.User{
		ID:          userID,
		FullName:    newUser.FullName,
		PhoneNumber: newUser.PhoneNumber,
		Email:       newUser.Email,
		Role:        newUser.Role,
		Active:      newUser.Active,
		CreatedAt:   newUser.CreatedAt,
		UpdatedAt:   time.Now(),
		CreatedBy:   newUser.CreatedBy,
		UpdatedBy:   newUser.UpdatedBy,
	}

	return resultNewUser, nil
}

func (repo *PostgresInventoryRepository) DeleteUserByID(ctx *gin.Context, id string) (string, error) {
	userID, err := repo.db.DeleteUserByID(ctx, id)
	if err != nil {
		log.Printf("User Repository: %s", err)

		return "", err
	}

	return userID, nil
}
