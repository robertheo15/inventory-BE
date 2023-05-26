package ginhttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) Test(ctx *gin.Context) {
	//s.service
	product := s.service.CreateProduct(ctx)
	ctx.JSON(http.StatusOK, product)

}
