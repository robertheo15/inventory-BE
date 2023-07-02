package inventory

import (
	"inventory-app-be/internal/models"
	"inventory-app-be/internal/repository/postgres/sqlc"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func (repo *PostgresInventoryRepository) CreateProductVariant(ctx *gin.Context, newProductVariant *models.ProductVariant) (*models.ProductVariant, error) {
	productVariant := sqlc.CreateProductVariantParams{

		PID:       newProductVariant.ProductID,
		PvID:      newProductVariant.ProductVariantID,
		Name:      newProductVariant.Name,
		Colour:    newProductVariant.Colour,
		Type:      newProductVariant.Type,
		Stock:     newProductVariant.Stock,
		Location:  newProductVariant.Location,
		CreatedBy: ctx.GetString("full_name"),
		UpdatedBy: ctx.GetString("full_name"),
	}

	ID, err := repo.db.CreateProductVariant(ctx, productVariant)
	if err != nil {
		log.Printf("Product variant Repository: %s", err)

		return nil, err
	}

	resultProductVariant := &models.ProductVariant{
		ID:               ID,
		ProductID:        productVariant.PID,
		ProductVariantID: productVariant.PvID,
		Name:             productVariant.Name,
		Colour:           productVariant.Colour,
		Stock:            productVariant.Stock,
		Location:         productVariant.Location,
		Type:             productVariant.Type,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		CreatedBy:        newProductVariant.CreatedBy,
		UpdatedBy:        newProductVariant.UpdatedBy,
	}

	return resultProductVariant, nil
}

func (repo *PostgresInventoryRepository) GetProductVariants(ctx *gin.Context) ([]*models.ProductVariant, error) {
	productVariants := make([]*models.ProductVariant, 0)

	pVariants, err := repo.db.GetProductVariants(ctx)
	if err != nil {
		log.Printf("Product variant Repository: %s", err)

		return nil, err
	}

	for _, pVariant := range pVariants {
		//if pVariant.PvStock < 0 {
		//	pVariant.Pv = 0
		//}
		productVariant := &models.ProductVariant{
			ID:               pVariant.PvID,
			ProductVariantID: pVariant.PvPvID,
			ProductID:        pVariant.PID,
			ProductName:      pVariant.PName,
			Name:             pVariant.PvName,
			Colour:           pVariant.PvColour,
			Stock:            pVariant.PvStock,
			Type:             pVariant.PvType,
			Location:         pVariant.PvLocation,
			CreatedAt:        pVariant.PvCreatedAt,
			UpdatedAt:        pVariant.PvUpdatedAt,
			CreatedBy:        pVariant.PvCreatedBy,
			UpdatedBy:        pVariant.PvUpdatedBy,
		}
		productVariants = append(productVariants, productVariant)
	}

	return productVariants, nil
}

func (repo *PostgresInventoryRepository) GetProductVariantsByProductID(ctx *gin.Context, id string) ([]*models.ProductVariant, error) {
	pVariants, err := repo.db.GetProductVariantsByProductID(ctx, id)
	if err != nil {
		log.Printf("Product variant Repository: %s", err)

		return nil, err
	}

	productVariants := make([]*models.ProductVariant, 0)

	for _, pVariant := range pVariants {
		productVariant := &models.ProductVariant{
			ID:        pVariant.ID,
			ProductID: pVariant.PID,
			Name:      pVariant.Name,
			Colour:    pVariant.Colour,
			Stock:     pVariant.Stock,
			Location:  pVariant.Location,
			Type:      pVariant.Type,
			CreatedAt: pVariant.CreatedAt,
			CreatedBy: pVariant.CreatedBy,
			UpdatedAt: pVariant.UpdatedAt,
			UpdatedBy: pVariant.UpdatedBy,
		}
		productVariants = append(productVariants, productVariant)
	}

	return productVariants, nil
}

func (repo *PostgresInventoryRepository) GetProductVariantByID(ctx *gin.Context, id string) (*models.ProductVariant, error) {
	pVariant, err := repo.db.GetProductVariantByID(ctx, id)
	if err != nil {
		log.Printf("Product variant Repository: %s", err)

		return nil, err
	}

	productVariant := &models.ProductVariant{
		ID:        pVariant.ID,
		ProductID: pVariant.PID,
		Name:      pVariant.Name,
		Colour:    pVariant.Colour,
		Stock:     pVariant.Stock,
		Location:  pVariant.Location,
		Type:      pVariant.Type,
		CreatedAt: pVariant.CreatedAt,
		CreatedBy: pVariant.CreatedBy,
		UpdatedAt: pVariant.UpdatedAt,
		UpdatedBy: pVariant.UpdatedBy,
	}

	return productVariant, nil
}

func (repo *PostgresInventoryRepository) UpdateProductVariantByID(ctx *gin.Context, newPVariant *models.ProductVariant) (*models.ProductVariant, error) {
	productVariant, err := repo.db.UpdateProductVariantByID(ctx, sqlc.UpdateProductVariantByIDParams{
		ID:        newPVariant.ID,
		PID:       newPVariant.ProductID,
		Name:      newPVariant.Name,
		Colour:    newPVariant.Colour,
		Stock:     newPVariant.Stock,
		CreatedBy: ctx.GetString("full_name"),
		UpdatedAt: newPVariant.UpdatedAt,
		UpdatedBy: newPVariant.UpdatedBy,
	})
	if err != nil {
		log.Printf("Product variant Repository: %s", err)

		return nil, err
	}

	resultProductVariant := &models.ProductVariant{
		ID:        productVariant.ID,
		ProductID: productVariant.PID,
		Name:      productVariant.Name,
		Colour:    productVariant.Colour,
		Stock:     productVariant.Stock,
		CreatedAt: productVariant.CreatedAt,
		CreatedBy: productVariant.CreatedBy,
		UpdatedAt: productVariant.UpdatedAt,
		UpdatedBy: productVariant.UpdatedBy,
	}

	return resultProductVariant, nil
}

func (repo *PostgresInventoryRepository) UpdateProductVariantStockByID(ctx *gin.Context, currentStock int32, newPVariant *models.ProductVariant) (string, error) {
	productVariantID, err := repo.db.UpdateProductVariantStockByID(ctx, sqlc.UpdateProductVariantStockByIDParams{
		ID:        newPVariant.ID,
		Stock:     currentStock + newPVariant.Stock,
		Updatedby: newPVariant.UpdatedBy,
	})
	if err != nil {
		log.Printf("Product variant Repository: update product variant stock by id %s", err)

		return "", err
	}
	log.Printf("Update stock success %s", productVariantID)

	return "Update stock success", nil
}

func (repo *PostgresInventoryRepository) DeleteProductVariantByID(ctx *gin.Context, id string) (string, error) {
	productID, err := repo.db.DeleteProductVariantByID(ctx, id)
	if err != nil {
		log.Printf("Product variant Repository: %s", err)

		return "", err
	}

	return productID, nil
}
