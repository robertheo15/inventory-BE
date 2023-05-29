package ginhttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) createProduct(ctx *gin.Context) {
}

func (s *Server) getProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, s.service.GetProducts(ctx))
}

func (s *Server) getProductByID(ctx *gin.Context) {
}

func (s *Server) updateProductByID(ctx *gin.Context) {
}

func (s *Server) deleteProductByID(ctx *gin.Context) {
}
