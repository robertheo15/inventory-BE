package inventory

import (
	"inventory-app-be/internal/models"
	"inventory-app-be/internal/repository/postgres/sqlc"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func (repo *PostgresInventoryRepository) CreateProduct(ctx *gin.Context, newProduct *models.Product) (*models.Product, error) {
	product := sqlc.CreateProductParams{
		ProductID:   newProduct.ProductID,
		Name:        newProduct.Name,
		Brand:       newProduct.Brand,
		Description: newProduct.Description,
		BasePrice:   newProduct.BasePrice,
		PriceEceran: newProduct.EceranPrice,
		PriceGrosir: newProduct.GrosirPrice,
		CreatedBy:   ctx.GetString("full_name"),
		UpdatedBy:   ctx.GetString("full_name"),
	}

	ID, err := repo.db.CreateProduct(ctx, product)
	if err != nil {
		log.Printf("Product Repository: %s", err)

		return nil, err
	}

	resultProduct := &models.Product{
		ID:          ID,
		ProductID:   product.ProductID,
		Name:        product.Name,
		Brand:       product.Brand,
		Description: product.Description,
		BasePrice:   product.BasePrice,
		EceranPrice: product.PriceEceran,
		GrosirPrice: product.PriceGrosir,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   product.CreatedBy,
		UpdatedBy:   product.UpdatedBy,
	}

	return resultProduct, nil
}

func (repo *PostgresInventoryRepository) GetProducts(ctx *gin.Context) ([]*models.Product, error) {
	products := make([]*models.Product, 0)

	newProducts, err := repo.db.GetProducts(ctx)
	if err != nil {
		log.Printf("Product Repository: %s", err)

		return nil, err
	}

	for _, product := range newProducts {
		newProduct := &models.Product{
			ID:          product.ID,
			Name:        product.Name,
			Brand:       product.Brand,
			Description: product.Description,
			BasePrice:   product.BasePrice,
			EceranPrice: product.PriceEceran,
			GrosirPrice: product.PriceGrosir,
			CreatedAt:   product.CreatedAt,
			CreatedBy:   product.CreatedBy,
			UpdatedAt:   product.UpdatedAt,
			UpdatedBy:   product.UpdatedBy,
			Children:    make([]*models.Product, 0),
		}
		products = append(products, newProduct)
	}

	return products, nil
}

func (repo *PostgresInventoryRepository) GetProductsBySupplierID(ctx *gin.Context, id string) ([]*models.Product, error) {
	products := make([]*models.Product, 0)
	productsResult, err := repo.db.GetProductsBySupplierID(ctx, id)
	if err != nil {
		log.Printf("Product Repository: GetProductsBySupplierID %s", err)

		return nil, err
	}

	for _, product := range productsResult {
		newProduct := &models.Product{
			ID:          product.ID,
			ProductID:   product.ProductID,
			Name:        product.Name,
			Brand:       product.Brand,
			Description: product.Description,
			BasePrice:   product.BasePrice,
			EceranPrice: product.PriceEceran,
			GrosirPrice: product.PriceGrosir,
			CreatedAt:   product.CreatedAt,
			CreatedBy:   product.CreatedBy,
			UpdatedAt:   product.UpdatedAt,
			UpdatedBy:   product.UpdatedBy,
			Children:    make([]*models.Product, 0),
		}
		products = append(products, newProduct)
	}

	return products, nil
}

func (repo *PostgresInventoryRepository) GetProductByID(ctx *gin.Context, id string) (*models.Product, error) {
	product, err := repo.db.GetProductByID(ctx, id)
	if err != nil {
		log.Printf("Product Repository: %s", err)

		return nil, err
	}

	newProduct := &models.Product{
		ID:          product.ID,
		ProductID:   product.ProductID,
		Name:        product.Name,
		Brand:       product.Brand,
		Description: product.Description,
		BasePrice:   product.BasePrice,
		EceranPrice: product.PriceEceran,
		GrosirPrice: product.PriceGrosir,
		CreatedAt:   product.CreatedAt,
		CreatedBy:   product.CreatedBy,
		UpdatedAt:   product.UpdatedAt,
		UpdatedBy:   product.UpdatedBy,
		Children:    make([]*models.Product, 0),
	}

	return newProduct, nil
}

func (repo *PostgresInventoryRepository) UpdateProductByID(ctx *gin.Context, newProduct *models.Product) (*models.Product, error) {
	product, err := repo.db.UpdateProductByID(ctx, sqlc.UpdateProductByIDParams{
		ID:          newProduct.ID,
		ProductID:   newProduct.ProductID,
		Name:        newProduct.Name,
		Brand:       newProduct.Brand,
		Description: newProduct.Description,
		Baseprice:   newProduct.BasePrice,
		Priceeceran: newProduct.EceranPrice,
		Pricegrosir: newProduct.GrosirPrice,
		Updatedby:   ctx.GetString("full_name"),
	})
	if err != nil {
		log.Printf("Product Repository: %s", err)

		return nil, err
	}

	resultProduct := &models.Product{
		ID:          product.ID,
		ProductID:   product.ProductID,
		Name:        product.Name,
		Brand:       product.Brand,
		Description: product.Description,
		BasePrice:   product.BasePrice,
		EceranPrice: product.PriceEceran,
		GrosirPrice: product.PriceGrosir,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		CreatedBy:   product.CreatedBy,
		UpdatedBy:   product.UpdatedBy,
	}

	return resultProduct, nil
}

func (repo *PostgresInventoryRepository) DeleteProductByID(ctx *gin.Context, id string) (string, error) {
	productID, err := repo.db.DeleteProductByID(ctx, id)
	if err != nil {
		log.Printf("Product Repository: %s", err)

		return "", err
	}

	return productID, nil
}
