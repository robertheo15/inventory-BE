package ginhttp

import (
	"fmt"
	"inventory-app-be/internal/models"
	pkgHttp "inventory-app-be/pkg/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) createTransaction(ctx *gin.Context) {
	var newTransaction *models.Transaction

	err := ctx.ShouldBindJSON(&newTransaction)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	supplier, err := s.service.CreateTransaction(ctx, newTransaction)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, supplier, pkgHttp.Created)
	}
}

func (s *Server) getTransactions(ctx *gin.Context) {
	suppliers, err := s.service.GetTransactions(ctx)

	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, suppliers, pkgHttp.Get)
	}
}

func (s *Server) getTransactionByID(ctx *gin.Context) {
	supplierID := ctx.Param("transactionID")

	supplier, err := s.service.GetTransactionByID(ctx, supplierID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, supplier, pkgHttp.Get)
	}
}

func (s *Server) updateTransactionByID(ctx *gin.Context) {
	var newTransaction *models.Transaction

	err := ctx.ShouldBindJSON(&newTransaction)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	supplier, err := s.service.UpdateTransactionByID(ctx, newTransaction)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, supplier, pkgHttp.Updated)
	}
}

func (s *Server) deleteTransactionByID(ctx *gin.Context) {
	supplierID := ctx.Param("transactionID")

	supplierID, err := s.service.DeleteTransactionByID(ctx, supplierID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK,
			fmt.Sprintf("Transaction dengan id: %s berhasil dihapus", supplierID), pkgHttp.Deleted)
	}
}
