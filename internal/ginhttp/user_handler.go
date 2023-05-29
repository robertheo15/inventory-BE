package ginhttp

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateUser(ctx *gin.Context) {
	// s.service
	user, err := s.service.CreateUser(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, user)
}

func (s *Server) GetUserByID(ctx *gin.Context) {
	// s.service
	user, err := s.service.GetUserByID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, user)
}
