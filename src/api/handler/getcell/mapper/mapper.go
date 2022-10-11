package mapper

import (
	"github.com/julianVelandia/golang-sheets/internal/cell/core/entity"
	"github.com/julianVelandia/golang-sheets/internal/cell/core/query"
	"github.com/julianVelandia/golang-sheets/src/api/handler/getcell/contract"
	"strconv"
)

type HandlerMapper struct{}

func (hm HandlerMapper) RequestToQuery(request contract.URLParams) (query.GetCells, error) {
	startRow, errStart := strconv.Atoi(request.StartRow)
	endRow, errEnd := strconv.Atoi(request.EndRow)

	if errStart != nil || errEnd != nil {
		err := errStart
		if errEnd != nil {
			err = errEnd
		}
		return query.GetCells{}, err
	}

	return query.GetCells{
		SheetName:   request.SheetName,
		StartColumn: request.StartColumn,
		EndColumn:   request.EndColumn,
		StartRow:    startRow,
		EndRow:      endRow,
	}, nil
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
