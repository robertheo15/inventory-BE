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
		Name:      newProductVariant.Name,
		Colour:    newProductVariant.Colour,
		CreatedBy: newProductVariant.CreatedBy,
		UpdatedBy: newProductVariant.UpdatedBy,
	}

	ID, err := repo.db.CreateProductVariant(ctx, productVariant)
	if err != nil {
		log.Printf("Product variant Repository: %s", err)

		return nil, err
	}

	resultProductVariant := &models.ProductVariant{
		ID:        ID,
		ProductID: productVariant.PID,
		Name:      productVariant.Name,
		Colour:    productVariant.Colour,
		CreatedAt: time.Now(),
		CreatedBy: newProductVariant.CreatedBy,
		UpdatedAt: time.Now(),
		UpdatedBy: newProductVariant.UpdatedBy,
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
		productVariant := &models.ProductVariant{
			ID:        pVariant.ID,
			ProductID: pVariant.PID,
			Name:      pVariant.Name,
			Colour:    pVariant.Colour,
			CreatedAt: pVariant.CreatedAt,
			CreatedBy: pVariant.CreatedBy,
			UpdatedAt: pVariant.UpdatedAt,
			UpdatedBy: pVariant.UpdatedBy,
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
		CreatedAt: newPVariant.CreatedAt,
		CreatedBy: newPVariant.CreatedBy,
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
		CreatedAt: productVariant.CreatedAt,
		CreatedBy: productVariant.CreatedBy,
		UpdatedAt: productVariant.UpdatedAt,
		UpdatedBy: productVariant.UpdatedBy,
	}

	return resultProductVariant, nil
}

func (repo *PostgresInventoryRepository) DeleteProductVariantByID(ctx *gin.Context, id string) (string, error) {
	productID, err := repo.db.DeleteProductVariantByID(ctx, id)
	if err != nil {
		log.Printf("Product variant Repository: %s", err)

		return "", err
	}

	return productID, nil
}
