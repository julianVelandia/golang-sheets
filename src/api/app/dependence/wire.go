package dependence

import (
	useCaseCells "github.com/julianVelandia/golang-sheets/internal/cell/core/usecase/getcell"
	RepositoryRead "github.com/julianVelandia/golang-sheets/internal/cell/infrastructure/repository/sheets/cell"
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
	sheetsClients := sheets.Client{}
	repositoryRead := RepositoryRead.NewRepositoryClient(sheetsClients)
	useCaseGetCell := useCaseCells.NewUseCase(repositoryRead)
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
	)
}
