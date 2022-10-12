package dependence

import (
	useCaseCells "github.com/julianVelandia/golang-sheets/internal/cell/core/usecase/getcell"
	useCaseMapperCells "github.com/julianVelandia/golang-sheets/internal/cell/core/usecase/getcell/mapper"
	repositoryRead "github.com/julianVelandia/golang-sheets/internal/cell/infrastructure/repository/sheets/cell"
	platformParams "github.com/julianVelandia/golang-sheets/internal/platform/params"
	"github.com/julianVelandia/golang-sheets/internal/platform/sheets"
	handlerGetCells "github.com/julianVelandia/golang-sheets/src/api/handler/getcell"
	mapperGetCells "github.com/julianVelandia/golang-sheets/src/api/handler/getcell/mapper"
	handlerPing "github.com/julianVelandia/golang-sheets/src/api/handler/ping"
)

type HandlerContainer struct {
	GetCellsHandler handlerGetCells.Handler
	PingHandler     handlerPing.Handler
}

func NewWire() HandlerContainer {
	sheetsClients := sheets.ClientSheets{}
	repository := repositoryRead.NewSheetsRepository(sheetsClients)
	mapperUseCase := useCaseMapperCells.Mapper{}
	useCaseGetCell := useCaseCells.NewUseCase(repository, mapperUseCase)
	return HandlerContainer{
		PingHandler:     newWirePingHandler(),
		GetCellsHandler: newWireGetCellsHandler(useCaseGetCell),
	}
}

func newWirePingHandler() handlerPing.Handler {
	return *handlerPing.NewHandler()
}

func newWireGetCellsHandler(useCase handlerGetCells.UseCase) handlerGetCells.Handler {
	return *handlerGetCells.NewHandler(
		useCase,
		mapperGetCells.HandlerMapper{},
		platformParams.NewParamValidation(getParamsValidationDefault()),
	)
}

func getParamsValidationDefault() map[string]platformParams.ValidationParams {
	paramsMap := make(map[string]platformParams.ValidationParams)
	paramsMap[platformParams.SheetsNameValidator{}.KeyParam()] = platformParams.SheetsNameValidator{IsRequired: true}
	paramsMap[platformParams.StartColumnValidator{}.KeyParam()] = platformParams.StartColumnValidator{IsRequired: true}
	paramsMap[platformParams.StartRowValidator{}.KeyParam()] = platformParams.StartRowValidator{IsRequired: true}
	paramsMap[platformParams.EndColumnValidator{}.KeyParam()] = platformParams.EndColumnValidator{IsRequired: true}
	paramsMap[platformParams.EndRowValidator{}.KeyParam()] = platformParams.EndRowValidator{IsRequired: true}
	return paramsMap
}
