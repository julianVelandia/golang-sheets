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
	actionValidateParameters string          = "validate_parameters"
	actionExecuteUseCase     string          = "execute_use_case"
	errorCells               log.LogsMessage = "error in the creation handler"
	entityType               string          = "get_cell"
	layer                    string          = "handler_Cells"
)

type UseCase interface {
	Execute(ctx context.Context, GetCells query.GetCells) ([]entity.Cell, error)
}

type Mapper interface {
	RequestToQuery(request contract.URLParams) (query.GetCells, error)
	EntityToResponse(entities []entity.Cell) contract.Response
}

type ValidationParams interface {
	BindParamsAndValidation(obj interface{}, params gin.Params) error
}

type Handler struct {
	useCase          UseCase
	mapper           Mapper
	validationParams ValidationParams
}

func NewHandler(useCase UseCase, mapper Mapper, validationParams ValidationParams) *Handler {
	return &Handler{
		useCase:          useCase,
		mapper:           mapper,
		validationParams: validationParams,
	}
}

func (h Handler) Handler(ginCTX *gin.Context) {

	requestParam := &contract.URLParams{}

	if errValidator := h.validationParams.BindParamsAndValidation(requestParam, ginCTX.Params); errValidator != nil {
		message := errorCells.GetMessageWithTagParams(
			log.NewTagParams(layer, actionValidateParameters,
				log.Params{
					constant.Key:        ginCTX.Params,
					constant.EntityType: entityType,
				}))
		ginCTX.JSON(http.StatusBadRequest, ErrorResponse.Response{
			Status:  http.StatusBadRequest,
			Message: message,
		})
		return
	}

	fullQuery, errMapper := h.mapper.RequestToQuery(*requestParam)
	if errMapper != nil {
		message := errorCells.GetMessageWithTagParams(
			log.NewTagParams(layer, actionValidateParameters,
				log.Params{
					constant.Key:        ginCTX.Params,
					constant.EntityType: entityType,
				}))
		ginCTX.JSON(http.StatusBadRequest, ErrorResponse.Response{
			Status:  http.StatusBadRequest,
			Message: message,
		})
		return
	}

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
