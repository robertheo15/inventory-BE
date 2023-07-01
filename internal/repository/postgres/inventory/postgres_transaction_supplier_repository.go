package inventory

import (
	"github.com/pkg/errors"
	"inventory-app-be/internal/models"
	"inventory-app-be/internal/repository/postgres/sqlc"
	"inventory-app-be/pkg/http/response"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	sedangDikemas = "sedang dikemas"
	sedangDikirim = "sedang dikirim"
	selesai       = "selesai"
)

var (
	supplier = "supplier"
	parent   = "parent"
	child    = "child"
)

var (
	indent = "indent"
	gudang = "gudang"
	toko   = "toko"
)

func (repo *PostgresInventoryRepository) CreateTransactionSupplier(ctx *gin.Context,
	newTransactionSupplier *response.TransactionSupplierRequest) (*response.TransactionSupplierResponse, error) {
	transaction := sqlc.CreateTransactionParams{
		Invoice:   newTransactionSupplier.Invoice,
		CID:       "",
		Status:    sedangDikirim,
		Type:      supplier,
		Methode:   indent,
		CreatedBy: ctx.GetString("full_name"),
		UpdatedBy: ctx.GetString("full_name"),
	}
	// parent
	newTransaction, err := repo.db.CreateTransaction(ctx, transaction)
	if err != nil {
		log.Printf("Transaction Repository: CreateTransactionCustomer %s", err)

		return nil, err
	}

	detailTransaction, err := repo.createDetailTransactionSupplier(ctx, newTransaction.ID, newTransactionSupplier.TransactionDetails)
	if err != nil {
		log.Printf("Transaction Repository: createDetailTransactionSupplier %s", err)

		return nil, err
	}

	// to response
	resultTransactionSupplier := &response.TransactionSupplierResponse{
		ID:                 newTransaction.ID,
		Invoice:            newTransaction.Invoice,
		Status:             transaction.Status,
		Type:               transaction.Type,
		TransactionDetails: detailTransaction,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		CreatedBy:          transaction.CreatedBy,
		UpdatedBy:          transaction.CreatedBy,
	}

	return resultTransactionSupplier, nil
}

func (repo *PostgresInventoryRepository) createDetailTransactionSupplier(ctx *gin.Context, transactionID string,
	detailTransactions []*response.TransactionDetail) ([]*response.TransactionDetail, error) {
	newTransactionDetails := make([]*response.TransactionDetail, 0)

	for _, tDetail := range detailTransactions {
		transactionDetailSqlc := sqlc.CreateTransactionDetailParams{
			TID:       transactionID,
			PID:       tDetail.ProductID,
			PvID:      tDetail.ProductVariantID,
			Price:     tDetail.Price,
			Qty:       tDetail.Qty,
			CreatedBy: ctx.GetString("full_name"),
			UpdatedBy: ctx.GetString("full_name"),
		}

		tDetailID, err := repo.db.CreateTransactionDetail(ctx, transactionDetailSqlc)
		if err != nil {
			log.Printf("Transaction Repository: %s", err)

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

func (repo *PostgresInventoryRepository) GetTransactionSupplier(ctx *gin.Context) ([]*models.Transaction, error) {
	transactions := make([]*models.Transaction, 0)

	newTransactions, err := repo.db.GetTransactions(ctx, []string{"supplier"})
	if err != nil {
		log.Printf("Transaction supplier Repository: GetTransactionSupplier %s", err)

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
		transactions = append(transactions, newTransaction)
	}

	return transactions, nil
}

func (repo *PostgresInventoryRepository) UpdateStatusTransactionByIDAndUpdateStock(ctx *gin.Context, id string) (string, error) {
	transaction, err := repo.GetTransactionSupplierByID(ctx, id)
	if err != nil {
		log.Printf("Transaction supplier Repository: update status transaction, get transaction by id %s", err)

		return "", err
	}

	if transaction.Status != sedangDikirim {
		return "", errors.Errorf("Transaction status is not sedang dikirim")
	}

	tID, err := repo.db.UpdateStatusTransactionByID(ctx, sqlc.UpdateStatusTransactionByIDParams{
		ID:        id,
		Status:    selesai,
		UpdatedBy: ctx.GetString("full_name"),
	})
	if err != nil {
		log.Printf("Transaction supplier Repository: update status transaction %s %s", tID, err)

		return "", err
	}

	transactionDetails, err := repo.GetTransactionDetailsByTID(ctx, transaction.ID)
	if err != nil {
		log.Printf("Transaction supplier Repository: update status transaction, get transaction details by transaction id %s", err)

		return "", err
	}

	for _, detail := range transactionDetails {
		productVariant, err := repo.GetProductVariantByID(ctx, detail.ProductVariantID)
		if err != nil {
			log.Printf("Transaction supplier Repository: update status transaction, get product variant by id %s", err)
			return "", err
		}

		message, err := repo.UpdateProductVariantStockByID(ctx, detail.Qty, productVariant)
		if err != nil {
			log.Printf("Transaction supplier Repository: update status transaction, update product variant stock by id %s", err)
			return "", err
		}

		log.Printf("Transaction supplier repository : %s", message)
	}

	return "Successfully update the status", nil
}

func (repo *PostgresInventoryRepository) GetTransactionSupplierByID(ctx *gin.Context, id string) (*models.Transaction, error) {
	transaction, err := repo.db.GetTransactionByID(ctx, sqlc.GetTransactionByIDParams{
		ID:      id,
		Column1: []string{"supplier"},
	})
	if err != nil {
		log.Printf("Transaction supplier Repository: get transaction by id %s", err)

		return nil, err
	}

	newTransaction := &models.Transaction{
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
	transactionDetails, err := repo.GetTransactionDetailsByTID(ctx, newTransaction.ID)
	if err != nil {
		log.Printf("Transaction supplier Repository: GetTransactionDetailsByTID %s", err)
	}

	newTransaction.TransactionDetails = transactionDetails

	return newTransaction, nil
}

func (repo *PostgresInventoryRepository) GetTransactionSuppliersByStatus(ctx *gin.Context, status string) ([]*models.Transaction, error) {
	transactions := make([]*models.Transaction, 0)

	newTransactions, err := repo.db.GetTransactionsSupplierByStatus(ctx, sqlc.GetTransactionsSupplierByStatusParams{
		Column1: []string{"supplier"},
		Status:  status,
	})
	if err != nil {
		log.Printf("Transaction supplier Repository: GetTransactionSupplier %s", err)

		return nil, err
	}

	for _, transaction := range newTransactions {
		newTransaction := &models.Transaction{
			ID:            transaction.TID,
			TransactionID: transaction.TTransactionID,
			CID:           transaction.TCID,
			Invoice:       transaction.TInvoice,
			Status:        transaction.TStatus,
			Type:          transaction.TType,
			Supplier:      &models.Supplier{BrandName: transaction.SBrandName},
			CreatedAt:     transaction.TCreatedAt,
			UpdatedAt:     transaction.TUpdatedAt,
			CreatedBy:     transaction.TCreatedBy,
			UpdatedBy:     transaction.TCreatedBy,
		}
		transactionDetails, err := repo.GetTransactionDetailsByTID(ctx, newTransaction.ID)
		if err != nil {
			log.Printf("Transaction supplier Repository: GetTransactionDetailsByTID %s", err)
		}

		newTransaction.TransactionDetails = transactionDetails
		transactions = append(transactions, newTransaction)
	}

	return transactions, nil
}
