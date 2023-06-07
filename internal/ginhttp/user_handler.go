package ginhttp

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"inventory-app-be/internal/models"
	pkgHttp "inventory-app-be/pkg/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) createUser(ctx *gin.Context) {
	var newUser *models.User

	validate := validator.New()

	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	err = validate.Struct(newUser)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	user, err := s.service.CreateUser(ctx, newUser)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, user, pkgHttp.Created)
	}
}

func (s *Server) login(ctx *gin.Context) {
	var user *models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	token, err := s.service.GetUserByEmail(ctx, user)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, token, pkgHttp.Created)
	}
}

func (s *Server) getUserDetailByToken(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		pkgHttp.WriteJSONResponse(ctx, http.StatusUnauthorized, nil, "Unauthorized")
		return
	}

	user, err := s.service.GetUserDetail(token)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, user, pkgHttp.Get)
	}
}

func (s *Server) getUserByID(ctx *gin.Context) {
	userID := ctx.Param("userID")

	user, err := s.service.GetUserByID(ctx, userID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, user, pkgHttp.Get)
	}
}

func (s *Server) updateUserByID(ctx *gin.Context) {
	var newUser *models.User

	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	user, err := s.service.UpdateUserByID(ctx, newUser)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK, user, pkgHttp.Updated)
	}
}

func (s *Server) deleteUserByID(ctx *gin.Context) {
	userID := ctx.Param("userID")

	productID, err := s.service.DeleteUserByID(ctx, userID)
	if err != nil {
		pkgHttp.WriteJSONResponse(ctx, http.StatusBadRequest, nil, err.Error())
	} else {
		pkgHttp.WriteJSONResponse(ctx, http.StatusOK,
			fmt.Sprintf("User dengan id: %s berhasil dihapus", productID), pkgHttp.Deleted)
	}
}
