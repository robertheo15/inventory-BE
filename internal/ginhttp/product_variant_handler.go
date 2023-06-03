package ginhttp

import (
	"fmt"
	"inventory-app-be/internal/models"
	pkgHttp "inventory-app-be/pkg/http"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func (s *Server) createProductVariant(ctx *gin.Context) {
	var newProductVariant *models.ProductVariant

	err := ctx.ShouldBindJSON(&newProductVariant)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	_, err = valid.ValidateStruct(newProductVariant)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	productVariant, err := s.service.CreateProductVariant(ctx, newProductVariant)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, productVariant, pkgHttp.Created)
	}
}

func (s *Server) getProductVariants(ctx *gin.Context) {
	productVariant, err := s.service.GetProductVariants(ctx)

	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, productVariant, pkgHttp.Get)
	}
}

func (s *Server) getProductVariantByID(ctx *gin.Context) {
	productVariantID := ctx.Param("productVariantID")

	productVariant, err := s.service.GetProductVariantByID(ctx, productVariantID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, productVariant, pkgHttp.Get)
	}
}

func (s *Server) getProductVariantByProductID(ctx *gin.Context) {
	productID := ctx.Param("productID")

	productVariants, err := s.service.GetProductVariantsByProductID(ctx, productID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, productVariants, pkgHttp.Get)
	}
}

func (s *Server) updateProductVariantByID(ctx *gin.Context) {
	var newProductVariant *models.ProductVariant

	err := ctx.ShouldBindJSON(&newProductVariant)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	_, err = valid.ValidateStruct(newProductVariant)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	productVariant, err := s.service.UpdateProductVariantByID(ctx, newProductVariant)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, productVariant, pkgHttp.Updated)
	}
}

func (s *Server) deleteProductVariantByID(ctx *gin.Context) {
	productVariantID := ctx.Param("productVariantID")

	productVariantID, err := s.service.DeleteProductVariantByID(ctx, productVariantID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK,
			fmt.Sprintf("Product dengan id: %s berhasil dihapus", productVariantID), pkgHttp.Deleted)
	}
}
