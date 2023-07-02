package ginhttp

import (
	"fmt"
	"inventory-app-be/internal/models"
	pkgHttp "inventory-app-be/pkg/http"
	"inventory-app-be/pkg/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) createTransaction(ctx *gin.Context) {
	var newTransaction *response.TransactionCustomerRequest

	err := ctx.ShouldBindJSON(&newTransaction)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	transaction, err := s.service.CreateTransaction(ctx, newTransaction)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, transaction, pkgHttp.Created)
	}
}

func (s *Server) getTransactions(ctx *gin.Context) {
	transactions, err := s.service.GetTransactionCustomers(ctx)

	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, transactions, pkgHttp.Get)
	}
}

type transactionStatusRequest struct {
	Status string `json:"status"`
}

func (s *Server) getTransactionCustomersByStatus(ctx *gin.Context) {
	var transaction *transactionStatusRequest
	err := ctx.ShouldBindJSON(&transaction)

	transactions, err := s.service.GetTransactionCustomersByStatus(ctx, transaction.Status)

	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, transactions, pkgHttp.Get)
	}
}

func (s *Server) getTransactionByID(ctx *gin.Context) {
	transactionID := ctx.Param("transactionID")

	transaction, err := s.service.GetTransactionByID(ctx, transactionID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, transaction, pkgHttp.Get)
	}
}

func (s *Server) updateTransactionByID(ctx *gin.Context) {
	var newTransaction *models.Transaction

	err := ctx.ShouldBindJSON(&newTransaction)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	transaction, err := s.service.UpdateTransactionByID(ctx, newTransaction)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, transaction, pkgHttp.Updated)
	}
}

func (s *Server) deleteTransactionByID(ctx *gin.Context) {
	transactionID := ctx.Param("transactionID")

	transactionID, err := s.service.DeleteTransactionByID(ctx, transactionID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK,
			fmt.Sprintf("Transaction dengan id: %s berhasil dihapus", transactionID), pkgHttp.Deleted)
	}
}

func (s *Server) updateStatusTransactionCustomerSedangDikirimByID(ctx *gin.Context) {
	transactionID := ctx.Param("transactionID")

	message, err := s.service.UpdateStatusTransactionSedangDikirimByID(ctx, transactionID)

	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, message, pkgHttp.Updated)
	}
}

func (s *Server) updateStatusTransactionCustomerSelesaiByID(ctx *gin.Context) {
	transactionID := ctx.Param("transactionID")

	message, err := s.service.UpdateStatusTransactionSelesaiByID(ctx, transactionID)

	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, message, pkgHttp.Updated)
	}
}
