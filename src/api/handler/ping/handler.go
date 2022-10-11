package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/golang-sheets/src/api/handler/ping/contract"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) Handler(ginCTX *gin.Context) {
	pong := contract.Pong{
		Message: "golang sheets Pong",
	}
	ginCTX.JSON(http.StatusOK, pong)
}
