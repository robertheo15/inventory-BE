package inventory

import (
	"fmt"
	"inventory-app-be/internal/models"
	"inventory-app-be/internal/repository/postgres/sqlc"
	"inventory-app-be/pkg/http/response"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func setInvoice() string {
	return fmt.Sprintf("KRSTX-%d", time.Now().UnixMilli())
}

func (repo *PostgresInventoryRepository) CreateTransaction(ctx *gin.Context, newTransaction *response.TransactionRequest) (*response.TransactionResponse, error) {
	transaction := sqlc.CreateTransactionParams{
		CID:        newTransaction.CID,
		Invoice:    setInvoice(),
		Status:     newTransaction.Status,
		Type:       newTransaction.Type,
		CreatedBy:  ctx.GetString("full_name"),
		UpdatedByy: ctx.GetString("full_name"),
	}
	// parent
	parent, err := repo.db.CreateTransaction(ctx, transaction)
	if err != nil {
		log.Printf("Transaction Repository: CreateTransaction%s", err)

		return nil, err
	}
	// child
	children, err := repo.createChildTransaction(ctx, parent.ID, newTransaction.Children)
	if err != nil {
		log.Printf("Transaction Repository: createChildTransaction %s", err)

		return nil, err
	}
	// to response
	resultTransaction := &response.TransactionResponse{
		ID:            parent.ID,
		TransactionID: transaction.TransactionID,
		CID:           transaction.CID,
		Invoice:       parent.Invoice,
		Status:        transaction.Status,
		Type:          transaction.Type,
		Children:      children,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		CreatedBy:     transaction.CreatedBy,
		UpdatedBy:     transaction.CreatedBy,
	}

	return resultTransaction, nil
}

func (repo *PostgresInventoryRepository) createChildTransaction(ctx *gin.Context, parentID string,
	childrenTransactions []*response.Transaction) ([]*response.Transaction, error) {
	children := make([]*response.Transaction, 0)

	for _, childTransaction := range childrenTransactions {
		newChildTransaction := sqlc.CreateTransactionParams{
			CID:           childTransaction.CID,
			TransactionID: parentID,
			Invoice:       setInvoice(),
			Status:        childTransaction.Status,
			Type:          childTransaction.Type,
			CreatedBy:     ctx.GetString("full_name"),
			UpdatedByy:    ctx.GetString("full_name"),
		}

		child, err := repo.db.CreateTransaction(ctx, newChildTransaction)
		if err != nil {
			log.Printf("Transaction Repository: %s", err)

			return nil, err
		}

		// detail transaction
		newTransactionDetails, err := repo.createDetailTransaction(ctx, child.ID, childTransaction.TransactionDetails)
		if err != nil {
			log.Printf("Transaction Repository: %s", err)

			return nil, err
		}

		c := &response.Transaction{
			ID:                 child.ID,
			CID:                childTransaction.CID,
			TransactionID:      parentID,
			TransactionDetails: newTransactionDetails,
			Invoice:            child.Invoice,
			Status:             childTransaction.Status,
			Type:               childTransaction.Type,
			CreatedBy:          ctx.GetString("full_name"),
			UpdatedBy:          ctx.GetString("full_name"),
		}
		children = append(children, c)
	}

	return children, nil
}

func (repo *PostgresInventoryRepository) createDetailTransaction(ctx *gin.Context, childID string,
	childrenTransactions []*response.TransactionDetail) ([]*response.TransactionDetail, error) {
	newTransactionDetails := make([]*response.TransactionDetail, 0)

	for _, tDetail := range childrenTransactions {
		transactionDetailSqlc := sqlc.CreateTransactionDetailParams{
			TID:       childID,
			PID:       tDetail.ProductID,
			Price:     tDetail.Price,
			Qty:       tDetail.Qty,
			CreatedBy: ctx.GetString("full_name"),
			UpdatedBy: ctx.GetString("full_name"),
		}
		// get current stock
		product, err := repo.db.GetProductByID(ctx, tDetail.ProductID)
		if err != nil {
			log.Printf("Transaction Repository: GetProductByID  %s", err)

			return nil, err
		}

		//reduce current stock with transaction detail qty
		err = repo.db.UpdateProductStockParentByID(ctx, sqlc.UpdateProductStockParentByIDParams{
			ID:        tDetail.ProductID,
			Stock:     product.Stock - tDetail.Qty,
			Updatedby: ctx.GetString("full_name"),
		})
		if err != nil {
			log.Printf("Transaction Repository: GetProductByID  %s", err)

			return nil, err
		}

		tDetailID, err := repo.db.CreateTransactionDetail(ctx, transactionDetailSqlc)
		if err != nil {
			log.Printf("Transaction Repository: %s", err)

			return nil, err
		}

		newTransactionDetail := &response.TransactionDetail{
			ID:            tDetailID,
			TransactionID: transactionDetailSqlc.TID,
			ProductID:     transactionDetailSqlc.PID,
			Price:         transactionDetailSqlc.Price,
			Qty:           transactionDetailSqlc.Qty,
			CreatedBy:     ctx.GetString("full_name"),
			UpdatedBy:     ctx.GetString("full_name"),
		}

		newTransactionDetails = append(newTransactionDetails, newTransactionDetail)
	}

	return newTransactionDetails, nil
}

func (repo *PostgresInventoryRepository) GetTransactions(ctx *gin.Context) ([]*models.Transaction, error) {
	parentTransactions := make([]*models.Transaction, 0)
	childTransactions := make([]*models.Transaction, 0)

	newTransactions, err := repo.db.GetTransactions(ctx)
	if err != nil {
		log.Printf("Transaction Repository: %s", err)

		return nil, err
	}

	for _, transaction := range newTransactions {
		if transaction.Type == "parent" {
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
			parentTransactions = append(parentTransactions, newTransaction)
		} else {
			childTransaction := &models.Transaction{
				ID:            transaction.ID,
				TransactionID: transaction.TransactionID,
				CID:           transaction.CID,
				Invoice:       transaction.Invoice,
				Status:        transaction.Status,
				Type:          transaction.Type,
				CreatedAt:     transaction.CreatedAt,
				UpdatedAt:     transaction.UpdatedAt,
				CreatedBy:     transaction.CreatedBy,
				UpdatedBy:     transaction.CreatedBy,
			}
			childTransactions = append(childTransactions, childTransaction)
		}
	}

	for _, parentTransaction := range parentTransactions {
		tempTransaction := make([]*models.Transaction, 0)

		for _, childTransaction := range childTransactions {
			if parentTransaction.ID == childTransaction.TransactionID {
				tempTransaction = append(tempTransaction, childTransaction)
			}

			parentTransaction.Children = tempTransaction
		}
	}

	return parentTransactions, nil
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

	childTransactions, err := repo.db.GetTransactionByChildID(ctx, id)
	if err != nil {
		log.Printf("Transaction Repository: GetTransactionByChildID %s", err)

		return nil, err
	}

	children := make([]*models.Transaction, 0)

	for _, childTransaction := range childTransactions {
		newChildren := &models.Transaction{
			ID:            childTransaction.ID,
			TransactionID: childTransaction.TransactionID,
			CID:           childTransaction.CID,
			Invoice:       childTransaction.Invoice,
			Status:        childTransaction.Status,
			Type:          childTransaction.Type,
			Children:      make([]*models.Transaction, 0),
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			CreatedBy:     childTransaction.CreatedBy,
			UpdatedBy:     childTransaction.CreatedBy,
		}
		children = append(children, newChildren)
	}

	newTransaction.Children = children

	return newTransaction, nil
}

func (repo *PostgresInventoryRepository) UpdateTransactionByID(ctx *gin.Context,
	newTransaction *models.Transaction) (*models.Transaction, error) {
	updateTransactionByIDParams := sqlc.UpdateTransactionByIDParams{
		ID:            newTransaction.ID,
		CID:           newTransaction.CID,
		TransactionID: newTransaction.TransactionID,
		Status:        newTransaction.Status,
		Type:          newTransaction.Type,
		CreatedBy:     ctx.GetString("full_name"),
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
