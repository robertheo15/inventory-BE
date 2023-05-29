package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestResponse(ctx *gin.Context, payload interface{}) {
	WriteJSONResponse(ctx, http.StatusBadRequest, gin.H{
		"error": true,
		"data":  payload,
	})
}

func InternalServerJSONResponse(ctx *gin.Context, payload interface{}) {
	WriteJSONResponse(ctx, http.StatusInternalServerError, gin.H{
		"error": true,
		"data":  payload,
	})
}

func NotFoundResponse(ctx *gin.Context, payload interface{}) {
	WriteJSONResponse(ctx, http.StatusNotFound, gin.H{
		"error": true,
		"data":  payload,
	})
}

func WriteJSONResponse(ctx *gin.Context, status int, payload interface{}) {
	ctx.JSON(status, payload)
}
