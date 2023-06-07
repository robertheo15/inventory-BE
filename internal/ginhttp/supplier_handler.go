package ginhttp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inventory-app-be/internal/models"
	pkgHttp "inventory-app-be/pkg/http"
	"net/http"
)

func (s *Server) createSupplier(ctx *gin.Context) {
	var newSupplier *models.Supplier

	err := ctx.ShouldBindJSON(&newSupplier)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	supplier, err := s.service.CreateSupplier(ctx, newSupplier)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, supplier, pkgHttp.Created)
	}
}

func (s *Server) getSuppliers(ctx *gin.Context) {
	suppliers, err := s.service.GetSuppliers(ctx)

	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, suppliers, pkgHttp.Get)
	}
}

func (s *Server) getSupplierByID(ctx *gin.Context) {
	supplierID := ctx.Param("supplierID")

	supplier, err := s.service.GetSupplierByID(ctx, supplierID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, supplier, pkgHttp.Get)
	}
}

func (s *Server) updateSupplierByID(ctx *gin.Context) {
	var newSupplier *models.Supplier

	err := ctx.ShouldBindJSON(&newSupplier)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	supplier, err := s.service.UpdateSupplierByID(ctx, newSupplier)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, supplier, pkgHttp.Updated)
	}
}

func (s *Server) deleteSupplierByID(ctx *gin.Context) {
	supplierID := ctx.Param("supplierID")

	supplierID, err := s.service.DeleteSupplierByID(ctx, supplierID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK,
			fmt.Sprintf("Supplier dengan id: %s berhasil dihapus", supplierID), pkgHttp.Deleted)
	}
}
