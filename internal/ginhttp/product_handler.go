package ginhttp

import (
	"fmt"

	"inventory-app-be/internal/models"
	pkgHttp "inventory-app-be/pkg/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) createProduct(ctx *gin.Context) {
	var newProduct *models.Product

	err := ctx.ShouldBindJSON(&newProduct)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	product, err := s.service.CreateProduct(ctx, newProduct)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, product, pkgHttp.Created)
	}
}

func (s *Server) getProducts(ctx *gin.Context) {
	products, err := s.service.GetProducts(ctx)

	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, products, pkgHttp.Get)
	}
}

func (s *Server) getProductByID(ctx *gin.Context) {
	productID := ctx.Param("productID")

	product, err := s.service.GetProductByID(ctx, productID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, product, pkgHttp.Get)
	}
}

func (s *Server) getProductBySupplierID(ctx *gin.Context) {
	supplierID := ctx.Param("supplierID")

	products, err := s.service.GetProductBySupplierID(ctx, supplierID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, products, pkgHttp.Get)
	}
}

func (s *Server) updateProductByID(ctx *gin.Context) {
	var newProduct *models.Product

	err := ctx.ShouldBindJSON(&newProduct)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	product, err := s.service.UpdateProductByID(ctx, newProduct)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, product, pkgHttp.Updated)
	}
}

func (s *Server) deleteProductByID(ctx *gin.Context) {
	productID := ctx.Param("productID")

	productID, err := s.service.DeleteProductByID(ctx, productID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK,
			fmt.Sprintf("Product dengan id: %s berhasil dihapus", productID), pkgHttp.Deleted)
	}
}
