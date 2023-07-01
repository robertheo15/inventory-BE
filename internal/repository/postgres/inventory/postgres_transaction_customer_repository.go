package inventory

import (
	"fmt"
	"github.com/pkg/errors"
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

func (repo *PostgresInventoryRepository) CreateTransactionCustomer(ctx *gin.Context,
	newTransaction *response.TransactionCustomerRequest) (*response.TransactionCustomerResponse, error) {
	transaction := sqlc.CreateTransactionParams{
		CID:       newTransaction.CID,
		Invoice:   newTransaction.Invoice,
		Status:    sedangDikemas,
		Type:      parent,
		Methode:   newTransaction.Methode,
		CreatedBy: ctx.GetString("full_name"),
		UpdatedBy: ctx.GetString("full_name"),
	}
	// parent
	parentTrx, err := repo.db.CreateTransaction(ctx, transaction)
	if err != nil {
		log.Printf("Transaction Repository: CreateTransactionCustomer%s", err)

		return nil, err
	}
	// child
	children, err := repo.createChildTransaction(ctx, parentTrx.ID, newTransaction.Children)
	if err != nil {
		log.Printf("Transaction Repository: createChildTransaction %s", err)

		return nil, err
	}
	// to response
	resultTransaction := &response.TransactionCustomerResponse{
		ID:            parentTrx.ID,
		TransactionID: transaction.TransactionID,
		CID:           transaction.CID,
		Invoice:       parentTrx.Invoice,
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
			Invoice:       childTransaction.Invoice,
			Status:        sedangDikemas,
			Type:          child,
			Methode:       childTransaction.Methode,
			CreatedBy:     ctx.GetString("full_name"),
			UpdatedBy:     ctx.GetString("full_name"),
		}

		childTrx, err := repo.db.CreateTransaction(ctx, newChildTransaction)
		if err != nil {
			log.Printf("Transaction Repository: %s", err)

			return nil, err
		}

		// detail transaction
		newTransactionDetails, err := repo.createDetailTransaction(ctx, childTrx.ID, childTransaction.TransactionDetails, childTransaction.Methode)
		if err != nil {
			log.Printf("Transaction Repository: %s", err)

			return nil, err
		}

		c := &response.Transaction{
			ID:                 childTrx.ID,
			CID:                childTransaction.CID,
			TransactionID:      parentID,
			TransactionDetails: newTransactionDetails,
			Invoice:            childTrx.Invoice,
			Status:             newChildTransaction.Status,
			Type:               newChildTransaction.Type,
			Methode:            childTransaction.Methode,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
			CreatedBy:          ctx.GetString("full_name"),
			UpdatedBy:          ctx.GetString("full_name"),
		}
		children = append(children, c)
	}

	return children, nil
}

func (repo *PostgresInventoryRepository) createDetailTransaction(ctx *gin.Context, childID string,
	childrenTransactions []*response.TransactionDetail, methode string) ([]*response.TransactionDetail, error) {
	newTransactionDetails := make([]*response.TransactionDetail, 0)

	for _, tDetail := range childrenTransactions {
		transactionDetailSqlc := sqlc.CreateTransactionDetailParams{
			TID:       childID,
			PID:       tDetail.ProductID,
			PvID:      tDetail.ProductVariantID,
			Price:     tDetail.Price,
			Qty:       tDetail.Qty,
			CreatedBy: ctx.GetString("full_name"),
			UpdatedBy: ctx.GetString("full_name"),
		}
		// get current stock
		productVariant, err := repo.db.GetProductVariantByID(ctx, tDetail.ProductVariantID)
		if err != nil {
			log.Printf("Transaction Repository: createDetailTransaction, GetProductByID  %s", err)

			return nil, err
		}

		// reduce current stock with transaction detail qty
		newProductVariant := sqlc.UpdateProductVariantStockByIDParams{}
		if methode != indent {
			newProductVariant = sqlc.UpdateProductVariantStockByIDParams{
				ID:        tDetail.ProductVariantID,
				Stock:     productVariant.Stock - tDetail.Qty,
				Updatedby: ctx.GetString("full_name"),
			}

			pvID, err := repo.db.UpdateProductVariantStockByID(ctx, newProductVariant)
			if err != nil {
				log.Printf("Transaction Repository: createDetailTransaction, update product variant stock by id  %s %s", err, pvID)

				return nil, err
			}
		}

		tDetailID, err := repo.db.CreateTransactionDetail(ctx, transactionDetailSqlc)
		if err != nil {
			log.Printf("Transaction Repository: createDetailTransaction, create transaction detail %s", err)

			return nil, err
		}

		newTransactionDetail := &response.TransactionDetail{
			ID:               tDetailID,
			TransactionID:    transactionDetailSqlc.TID,
			ProductID:        transactionDetailSqlc.PID,
			ProductVariantID: transactionDetailSqlc.PvID,
			Price:            transactionDetailSqlc.Price,
			Qty:              transactionDetailSqlc.Qty,
			CreatedBy:        ctx.GetString("full_name"),
			UpdatedBy:        ctx.GetString("full_name"),
		}

		newTransactionDetails = append(newTransactionDetails, newTransactionDetail)
	}

	return newTransactionDetails, nil
}

func (repo *PostgresInventoryRepository) GetTransactionCustomers(ctx *gin.Context) ([]*models.Transaction, error) {
	parentTransactions := make([]*models.Transaction, 0)
	childTransactions := make([]*models.Transaction, 0)

	newTransactions, err := repo.db.GetTransactions(ctx, []string{"parent", "child"})
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

func (repo *PostgresInventoryRepository) GetTransactionCustomersByStatus(ctx *gin.Context, status string) ([]*models.Transaction, error) {
	parentTransactions := make([]*models.Transaction, 0)
	childTransactions := make([]*models.Transaction, 0)

	newTransactions, err := repo.db.GetTransactionsByStatus(ctx,
		sqlc.GetTransactionsByStatusParams{
			Column1: []string{"parent", "child"},
			Status:  status,
		})
	if err != nil {
		log.Printf("Transaction Repository: %s", err)

		return nil, err
	}

	for _, transaction := range newTransactions {
		customer, err := repo.GetCustomerByID(ctx, transaction.CID)
		if err != nil {
			log.Printf("Transaction Repository: get customer by id%s", err)
		}

		if transaction.Type == "parent" {
			newTransaction := &models.Transaction{
				ID:            transaction.ID,
				TransactionID: transaction.TransactionID,
				CID:           transaction.CID,
				Invoice:       transaction.Invoice,
				Status:        transaction.Status,
				Type:          transaction.Type,
				Children:      make([]*models.Transaction, 0),
				Customer:      customer,
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
				Customer:      customer,
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
			for _, c := range parentTransaction.Children {
				tDetails, err := repo.GetTransactionDetailsByTID(ctx, c.ID)
				if err != nil {
					log.Printf("Transaction Repository: GetTransactionCustomersByStatus, GetTransactionDetailsByTID %s", err)
				}
				c.TransactionDetails = tDetails
			}
		}
	}

	return parentTransactions, nil
}

func (repo *PostgresInventoryRepository) GetTransactionCustomerParentByID(ctx *gin.Context, id string) (*models.Transaction, error) {
	transaction, err := repo.db.GetTransactionByID(ctx, sqlc.GetTransactionByIDParams{
		ID:      id,
		Column1: []string{"parent"},
	})
	if err != nil {
		log.Printf("Transaction customer repository: get transaction by id %s", err)

		return nil, err
	}

	newTransaction := &models.Transaction{
		ID:            transaction.ID,
		TransactionID: transaction.TransactionID,
		CID:           transaction.CID,
		Invoice:       transaction.Invoice,
		Status:        transaction.Status,
		Type:          transaction.Type,
		Methode:       transaction.Methode,
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

func (repo *PostgresInventoryRepository) UpdateTransactionCustomerByID(ctx *gin.Context,
	newTransaction *models.Transaction) (*models.Transaction, error) {
	updateTransactionByIDParams := sqlc.UpdateTransactionByIDParams{
		ID:            newTransaction.ID,
		CID:           newTransaction.CID,
		TransactionID: newTransaction.TransactionID,
		Status:        newTransaction.Status,
		Type:          newTransaction.Type,
		UpdatedBy:     ctx.GetString("full_name"),
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

func (repo *PostgresInventoryRepository) GetTransactionCustomerChildByID(ctx *gin.Context, id string) (*models.Transaction, error) {
	transaction, err := repo.db.GetTransactionByID(ctx, sqlc.GetTransactionByIDParams{
		ID:      id,
		Column1: []string{"parent", "child"},
	})
	if err != nil {
		log.Printf("Transaction customer Repository: get transaction by id %s", err)

		return nil, err
	}

	newTransaction := &models.Transaction{
		ID:            transaction.ID,
		TransactionID: transaction.TransactionID,
		CID:           transaction.CID,
		Invoice:       transaction.Invoice,
		Status:        transaction.Status,
		Type:          transaction.Type,
		Methode:       transaction.Methode,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
		CreatedBy:     transaction.CreatedBy,
		UpdatedBy:     transaction.CreatedBy,
	}
	transactionDetails, err := repo.GetTransactionDetailsByTID(ctx, newTransaction.ID)
	if err != nil {
		log.Printf("Transaction supplier Repository: GetTransactionDetailsByTID %s", err)
	}

	newTransaction.TransactionDetails = transactionDetails

	return newTransaction, nil
}

func (repo *PostgresInventoryRepository) UpdateStatusTransactionCSedangDikirimByID(ctx *gin.Context, id string) (string, error) {
	transaction, err := repo.GetTransactionCustomerChildByID(ctx, id)
	if err != nil {
		log.Printf("Transaction customer repository: update status transaction, get transaction by id %s", err)

		return "", err
	}

	if transaction.Status != sedangDikemas {
		return "", errors.Errorf("Transaction status is not sedand dikemas")
	}

	tID, err := repo.db.UpdateStatusTransactionByID(ctx, sqlc.UpdateStatusTransactionByIDParams{
		ID:        transaction.ID,
		Status:    sedangDikirim,
		UpdatedBy: ctx.GetString("full_name"),
	})
	if err != nil {
		log.Printf("Transaction customer repository: update status transaction %s %s", tID, err)

		return "", err
	}

	if transaction.Methode == indent {
		for _, detail := range transaction.TransactionDetails {
			productVariant, err := repo.GetProductVariantByID(ctx, detail.ProductVariantID)
			if err != nil {
				log.Printf("Transaction customer repository: update status transaction, get product variant by id %s", err)
				return "", err
			}

			message, err := repo.UpdateProductVariantStockByID(ctx, detail.Qty, productVariant)
			if err != nil {
				log.Printf("Transaction customer repository: update status transaction, update product variant stock by id %s", err)
				return "", err
			}

			log.Printf("Transaction customer repository : %s", message)
		}
	}

	return "Successfully update the status", nil
}

func (repo *PostgresInventoryRepository) UpdateStatusTransactionCSelesaiByID(ctx *gin.Context, id string) (string, error) {
	transaction, err := repo.GetTransactionCustomerChildByID(ctx, id)
	if err != nil {
		log.Printf("Transaction customer repository: update status transaction, get transaction by id %s", err)

		return "", err
	}

	if transaction.Status != sedangDikirim {
		return "", errors.Errorf("Transaction status is not sedang dikirim")
	}

	tID, err := repo.db.UpdateStatusTransactionByID(ctx, sqlc.UpdateStatusTransactionByIDParams{
		ID:        transaction.ID,
		Status:    selesai,
		UpdatedBy: ctx.GetString("full_name"),
	})
	if err != nil {
		log.Printf("Transaction customer repository: update status transaction %s %s", tID, err)

		return "", err
	}

	return "Successfully update the status", nil
}

func (repo *PostgresInventoryRepository) DeleteTransactionCustomerByID(ctx *gin.Context, id string) (string, error) {
	transactionID, err := repo.db.DeleteTransactionByID(ctx, id)
	if err != nil {
		log.Printf("Transaction Repository: %s", err)

		return "", err
	}

	return transactionID, nil
}
