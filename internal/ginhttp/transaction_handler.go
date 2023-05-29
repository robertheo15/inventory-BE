package ginhttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) createTransaction(ctx *gin.Context) {

}

func (s *Server) getTransactions(ctx *gin.Context) {
	s.service.GetProducts(ctx)
	ctx.JSON(http.StatusOK, s.service.GetProducts(ctx))

}

func (s *Server) getTransactionByID(ctx *gin.Context) {

}

func (s *Server) updateTransaction(ctx *gin.Context) {

}
