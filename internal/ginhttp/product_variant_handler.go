package ginhttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) createProductVariant(ctx *gin.Context) {

}

func (s *Server) getProductVariants(ctx *gin.Context) {
	s.service.GetProducts(ctx)
	ctx.JSON(http.StatusOK, s.service.GetProducts(ctx))

}

func (s *Server) getProductVariantByID(ctx *gin.Context) {

}

func (s *Server) updateProductVariantByID(ctx *gin.Context) {

}

func (s *Server) deleteProductVariantByID(ctx *gin.Context) {

}
