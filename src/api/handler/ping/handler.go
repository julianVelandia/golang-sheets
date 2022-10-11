package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/GolangSheets/src/api/handler/ping/contract"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) Handler(ginCTX *gin.Context) {
	pong := contract.Pong{
		Message: "Eureka Pong",
	}
	ginCTX.JSON(http.StatusOK, pong)
}

