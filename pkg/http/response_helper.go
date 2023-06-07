package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Get     = "Data berhasil diambil"
	Created = "Data berhasil ditambah"
	Updated = "Data berhasil diubah"
	Deleted = "Data berhasil dihapus"
)

func WriteJSONResponse(ctx *gin.Context, status int, payload interface{}, message interface{}) {
	err := false
	if status == http.StatusBadRequest || status == http.StatusUnauthorized {
		err = true
	}

	ctx.JSON(status, gin.H{
		"error":   err,
		"data":    payload,
		"message": message,
	})
}
