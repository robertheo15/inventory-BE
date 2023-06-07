package ginhttp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inventory-app-be/internal/models"
	pkgHttp "inventory-app-be/pkg/http"
	"net/http"
)

func (s *Server) createCustomer(ctx *gin.Context) {
	var newCustomer *models.Customer

	err := ctx.ShouldBindJSON(&newCustomer)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	customer, err := s.service.CreateCustomer(ctx, newCustomer)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, customer, pkgHttp.Created)
	}
}

func (s *Server) getCustomers(ctx *gin.Context) {
	customers, err := s.service.GetCustomers(ctx)

	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, customers, pkgHttp.Get)
	}
}

func (s *Server) getCustomerByID(ctx *gin.Context) {
	customerID := ctx.Param("customerID")

	customer, err := s.service.GetCustomerByID(ctx, customerID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, customer, pkgHttp.Get)
	}
}

func (s *Server) updateCustomerByID(ctx *gin.Context) {
	var newCustomer *models.Customer

	err := ctx.ShouldBindJSON(&newCustomer)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	customer, err := s.service.UpdateCustomerByID(ctx, newCustomer)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, customer, pkgHttp.Updated)
	}
}

func (s *Server) deleteCustomerByID(ctx *gin.Context) {
	customerID := ctx.Param("customerID")

	customerID, err := s.service.DeleteCustomerByID(ctx, customerID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK,
			fmt.Sprintf("Customer dengan id: %s berhasil dihapus", customerID), pkgHttp.Deleted)
	}
}
