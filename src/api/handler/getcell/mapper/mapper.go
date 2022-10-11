package mapper

import (
	"github.com/julianVelandia/golang-sheets/internal/cell/core/entity"
	"github.com/julianVelandia/golang-sheets/internal/cell/core/query"
	"github.com/julianVelandia/golang-sheets/src/api/handler/getcell/contract"
)

type HandlerMapper struct{}

func (hm HandlerMapper) RequestToQuery(request contract.URLParams) query.GetCells {

	return query.GetCells{
		SheetName:   request.SheetName,
		StartColumn: request.StartColumn,
		EndColumn:   request.EndColumn,
		StartRow:    request.StartRow,
		EndRow:      request.EndRow,
	}
}

func (hm HandlerMapper) EntityToResponse(entities []entity.Cell) contract.Response {

	cellsToResponse := make([]contract.Cell, 0)

	for i := range entities {

		cellsToResponse = append(cellsToResponse, contract.Cell{
			CellPosition: entities[i].CellPosition,
			Information:  entities[i].Information,
		})
	}
	response := contract.Response{
		Cells: cellsToResponse,
	}

	return response
}
