package ginhttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) createProduct(ctx *gin.Context) {

}

func (s *Server) getProducts(ctx *gin.Context) {
	//s.service.GetProducts(ctx)
	ctx.JSON(http.StatusOK, s.service.GetProducts(ctx))

}

func (s *Server) getProductByID(ctx *gin.Context) {

}

func (s *Server) updateProductByID(ctx *gin.Context) {

}

func (s *Server) deleteProductByID(ctx *gin.Context) {

}
