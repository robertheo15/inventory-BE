package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequestResponse(ctx *gin.Context, payload interface{}) {
	WriteJsonResponse(ctx, http.StatusBadRequest, gin.H{
		"error": true,
		"data":  payload,
	})
}

func InternalServerJsonResponse(ctx *gin.Context, payload interface{}) {
	WriteJsonResponse(ctx, http.StatusInternalServerError, gin.H{
		"error": true,
		"data":  payload,
	})
}

func NotFoundResponse(ctx *gin.Context, payload interface{}) {
	WriteJsonResponse(ctx, http.StatusNotFound, gin.H{
		"error": true,
		"data":  payload,
	})
}

func WriteJsonResponse(ctx *gin.Context, status int, payload interface{}) {
	ctx.JSON(status, payload)
}
