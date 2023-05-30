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

//func setMessageHttp(data any, ) string {
//	return fmt.Sprintf("%s", data)
//}

func WriteJSONResponse(ctx *gin.Context, status int, payload interface{}, message interface{}) {
	err := false
	if status == http.StatusBadRequest {
		err = true
	}

	ctx.JSON(status, gin.H{
		"error":   err,
		"data":    payload,
		"message": message,
	})
}
