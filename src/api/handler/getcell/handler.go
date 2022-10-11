package getcell

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/golang-sheets/internal/cell/core/entity"
	"github.com/julianVelandia/golang-sheets/internal/cell/core/query"
	"github.com/julianVelandia/golang-sheets/internal/platform/constant"
	"github.com/julianVelandia/golang-sheets/internal/platform/log"
	ErrorResponse "github.com/julianVelandia/golang-sheets/src/api/handler"
	"github.com/julianVelandia/golang-sheets/src/api/handler/getcell/contract"
)

const (
	actionExecuteUseCase string          = "execute_use_case"
	errorCells           log.LogsMessage = "error in the creation handler"
	entityType           string          = "get_cell"
	layer                string          = "handler_Cells"
)

type UseCase interface {
	Execute(ctx context.Context, GetCells query.GetCells) ([]entity.Cell, error)
}

type Mapper interface {
	RequestToQuery(request contract.URLParams) query.GetCells
	EntityToResponse(entities []entity.Cell) contract.Response
}

type Handler struct {
	useCase UseCase
	mapper  Mapper
}

func NewHandler(useCase UseCase, mapper Mapper) *Handler {
	return &Handler{
		useCase: useCase,
		mapper:  mapper,
	}
}

func (h Handler) Handler(ginCTX *gin.Context) {

	requestParam := &contract.URLParams{}

	fullQuery := h.mapper.RequestToQuery(*requestParam)

	cells, errorUseCase := h.useCase.Execute(ginCTX, fullQuery)

	if errorUseCase != nil {
		message := errorCells.GetMessageWithTagParams(
			log.NewTagParams(layer, actionExecuteUseCase,
				log.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s_%v_%v`,
						requestParam.SheetName,
						requestParam.StartColumn,
						requestParam.EndColumn,
						requestParam.StartRow,
						requestParam.EndRow,
					),
					constant.EntityType: entityType,
				}))
		ginCTX.JSON(http.StatusInternalServerError, ErrorResponse.Response{
			Status:  http.StatusInternalServerError,
			Message: message,
		})
		return
	}

	response := h.mapper.EntityToResponse(cells)

	ginCTX.JSON(http.StatusOK, response)
}
