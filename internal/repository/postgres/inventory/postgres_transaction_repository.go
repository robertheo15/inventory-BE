package inventory

import (
	"inventory-app-be/internal/models"
	"inventory-app-be/internal/repository/postgres/sqlc"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func (repo *PostgresInventoryRepository) CreateTransaction(ctx *gin.Context, newTransaction *models.Transaction) (*models.Transaction, error) {
	transaction := sqlc.CreateTransactionParams{
		CID:           newTransaction.CID,
		TransactionID: newTransaction.TransactionID,
		Invoice:       newTransaction.Invoice,
		Status:        newTransaction.Status,
		Type:          newTransaction.Type,
		CreatedBy:     newTransaction.CreatedBy,
		UpdatedByy:    newTransaction.UpdatedBy,
	}

	ID, err := repo.db.CreateTransaction(ctx, transaction)
	if err != nil {
		log.Printf("Transaction Repository: %s", err)

		return nil, err
	}

	resultTransaction := &models.Transaction{
		ID:            ID,
		TransactionID: transaction.TransactionID,
		CID:           transaction.CID,
		Invoice:       transaction.Invoice,
		Status:        transaction.Status,
		Type:          transaction.Type,
		Children:      nil,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		CreatedBy:     transaction.CreatedBy,
		UpdatedBy:     transaction.CreatedBy,
	}

	return resultTransaction, nil
}

func (repo *PostgresInventoryRepository) GetTransactions(ctx *gin.Context) ([]*models.Transaction, error) {
	transactions := make([]*models.Transaction, 0)

	newTransactions, err := repo.db.GetTransactions(ctx)
	if err != nil {
		log.Printf("Transaction Repository: %s", err)

		return nil, err
	}

	for _, transaction := range newTransactions {
		newTransaction := &models.Transaction{
			ID:            transaction.ID,
			TransactionID: transaction.TransactionID,
			CID:           transaction.CID,
			Invoice:       transaction.Invoice,
			Status:        transaction.Status,
			Type:          transaction.Type,
			Children:      make([]*models.Transaction, 0),
			CreatedAt:     transaction.CreatedAt,
			UpdatedAt:     transaction.UpdatedAt,
			CreatedBy:     transaction.CreatedBy,
			UpdatedBy:     transaction.CreatedBy,
		}
		transactions = append(transactions, newTransaction)
	}

	return transactions, nil
}

func (repo *PostgresInventoryRepository) GetTransactionByID(ctx *gin.Context, id string) (*models.Transaction, error) {
	transaction, err := repo.db.GetTransactionByID(ctx, id)
	if err != nil {
		log.Printf("Transaction Repository: %s", err)

		return nil, err
	}

	newTransaction := &models.Transaction{
		ID:            transaction.ID,
		TransactionID: transaction.TransactionID,
		CID:           transaction.CID,
		Invoice:       transaction.Invoice,
		Status:        transaction.Status,
		Type:          transaction.Type,
		Children:      make([]*models.Transaction, 0),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		CreatedBy:     transaction.CreatedBy,
		UpdatedBy:     transaction.CreatedBy,
	}

	return newTransaction, nil
}

func (repo *PostgresInventoryRepository) UpdateTransactionByID(ctx *gin.Context,
	newTransaction *models.Transaction) (*models.Transaction, error) {
	updateTransactionByIDParams := sqlc.UpdateTransactionByIDParams{
		ID:            newTransaction.ID,
		CID:           newTransaction.CID,
		TransactionID: newTransaction.TransactionID,
		Invoice:       newTransaction.Invoice,
		Status:        newTransaction.Status,
		Type:          newTransaction.Type,
		CreatedBy:     newTransaction.CreatedBy,
		UpdatedBy:     newTransaction.UpdatedBy,
	}

	transaction, err := repo.db.UpdateTransactionByID(ctx, updateTransactionByIDParams)
	if err != nil {
		log.Printf("Transaction Repository: %s", err)

		return nil, err
	}

	resultTransaction := &models.Transaction{
		ID:            transaction.ID,
		TransactionID: transaction.TransactionID,
		CID:           transaction.CID,
		Invoice:       transaction.Invoice,
		Status:        transaction.Status,
		Type:          transaction.Type,
		Children:      make([]*models.Transaction, 0),
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     time.Now(),
		CreatedBy:     transaction.CreatedBy,
		UpdatedBy:     transaction.CreatedBy,
	}

	return resultTransaction, nil
}

func (repo *PostgresInventoryRepository) DeleteTransactionByID(ctx *gin.Context, id string) (string, error) {
	transactionID, err := repo.db.DeleteTransactionByID(ctx, id)
	if err != nil {
		log.Printf("Transaction Repository: %s", err)

		return "", err
	}

	return transactionID, nil
}
