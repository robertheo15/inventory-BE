package inventory

import (
	"inventory-app-be/internal/models"
	"inventory-app-be/internal/repository/postgres/sqlc"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func (repo *PostgresInventoryRepository) CreatePrice(ctx *gin.Context, newPrice *models.Price) (*models.Price, error) {
	priceSqlcParam := sqlc.CreatePriceParams{
		Eceran:    newPrice.Eceran,
		Grosir:    newPrice.Grosir,
		CreatedBy: newPrice.CreatedBy,
		UpdatedBy: newPrice.UpdatedBy,
	}

	priceID, err := repo.db.CreatePrice(ctx, priceSqlcParam)
	if err != nil {
		log.Printf("price repository error: create price %s\n", err)

		return nil, err
	}

	resultUser := &models.Price{
		ID:        priceID,
		Eceran:    newPrice.Eceran,
		Grosir:    newPrice.Grosir,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: newPrice.CreatedBy,
		UpdatedBy: newPrice.UpdatedBy,
	}

	return resultUser, nil
}

func (repo *PostgresInventoryRepository) GetPrice(ctx *gin.Context) (*models.Price, error) {
	priceResult, err := repo.db.GetPrice(ctx)
	if err != nil {
		log.Printf("Price repository: update price by id %s", err)

		return nil, err
	}

	price := &models.Price{
		ID:        priceResult.ID,
		Eceran:    priceResult.Eceran,
		Grosir:    priceResult.Grosir,
		CreatedAt: priceResult.CreatedAt,
		UpdatedAt: priceResult.UpdatedAt,
		CreatedBy: priceResult.CreatedBy,
		UpdatedBy: priceResult.UpdatedBy,
	}

	return price, nil
}

func (repo *PostgresInventoryRepository) UpdatePriceByID(ctx *gin.Context, newPrice *models.Price) (*models.Price, error) {
	priceSqlcParam := sqlc.UpdatePriceByIDParams{
		Eceran:    newPrice.Eceran,
		Grosir:    newPrice.Grosir,
		UpdatedBy: newPrice.UpdatedBy,
	}

	priceID, err := repo.db.UpdatePriceByID(ctx, priceSqlcParam)
	if err != nil {
		log.Printf("Price repository: update price by id %s", err)

		return nil, err
	}

	price := &models.Price{
		ID:        priceID,
		Eceran:    newPrice.Eceran,
		Grosir:    newPrice.Grosir,
		UpdatedAt: time.Now(),
		CreatedBy: newPrice.CreatedBy,
		UpdatedBy: newPrice.UpdatedBy,
	}

	return price, nil
}

func (repo *PostgresInventoryRepository) DeletePriceByID(ctx *gin.Context, id string) (string, error) {
	priceID, err := repo.db.DeletePriceByID(ctx, id)
	if err != nil {
		log.Printf("Price Repository: Delete Price By ID %s", err)

		return "", err
	}

	return priceID, nil
}
