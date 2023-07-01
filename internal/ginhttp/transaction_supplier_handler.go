package ginhttp

import (
	pkgHttp "inventory-app-be/pkg/http"
	"inventory-app-be/pkg/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) createTransactionSupplier(ctx *gin.Context) {
	var newTransactionSupplier *response.TransactionSupplierRequest

	err := ctx.ShouldBindJSON(&newTransactionSupplier)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	transaction, err := s.service.CreateTransactionSupplier(ctx, newTransactionSupplier)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, transaction, pkgHttp.Created)
	}
}

func (s *Server) getTransactionSuppliers(ctx *gin.Context) {
	transactionSuppliers, err := s.service.GetTransactionSuppliers(ctx)

	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, transactionSuppliers, pkgHttp.Get)
	}
}

func (s *Server) getTransactionSupplierByStatus(ctx *gin.Context) {
	var transaction *transactionStatusRequest
	err := ctx.ShouldBindJSON(&transaction)

	transactions, err := s.service.GetTransactionSupplierByStatus(ctx, transaction.Status)

	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, transactions, pkgHttp.Get)
	}
}

func (s *Server) updateStatusTransactionByIDAndUpdateStock(ctx *gin.Context) {
	transactionID := ctx.Param("transactionID")

	message, err := s.service.UpdateStatusTransactionByIDAndUpdateStock(ctx, transactionID)

	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, message, pkgHttp.Updated)
	}
}
