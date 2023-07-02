package inventory

import (
	"inventory-app-be/internal/models"
	"log"

	"github.com/gin-gonic/gin"
)

func (repo *PostgresInventoryRepository) GetTransactionDetailsByTID(ctx *gin.Context, TID string) ([]*models.TransactionDetail, error) {
	transactionDetails := make([]*models.TransactionDetail, 0)

	sqlcTDetails, err := repo.db.GetTransactionDetailByTID(ctx, TID)
	if err != nil {
		log.Printf("Transaction Repository: %s", err)

		return nil, err
	}

	for _, transaction := range sqlcTDetails {
		newTransactionDetail := &models.TransactionDetail{
			ID:               transaction.ID,
			TransactionID:    transaction.TID,
			ProductID:        transaction.PID,
			ProductVariantID: transaction.PvID,
			Price:            transaction.Price,
			Qty:              transaction.Qty,
			CreatedAt:        transaction.CreatedAt,
			UpdatedAt:        transaction.UpdatedAt,
			CreatedBy:        transaction.CreatedBy,
			UpdatedBy:        transaction.CreatedBy,
		}
		transactionDetails = append(transactionDetails, newTransactionDetail)
	}

	return transactionDetails, nil
}
