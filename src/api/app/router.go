package app

import (
	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/golang-sheets/src/api/app/dependence"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	handlers := dependence.NewWire()
	configureMappings(router, handlers)

	return router
}

func configureMappings(router *gin.Engine, handlers dependence.HandlerContainer) {
	apiGroup := router.Group("v1.0")
	apiGroup.GET("/:sheet_name/:start_column/:start_row/:end_column/:end_row",
		handlers.GetCellsHandler.Handler,
	)
	apiGroup.GET("/ping",
		handlers.PingHandler.Handler,
	)
}
