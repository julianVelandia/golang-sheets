package mapper

import (
	"github.com/julianVelandia/golang-sheets/internal/cell/core/entity"
	"github.com/julianVelandia/golang-sheets/internal/platform/sheets/model"
)

type Mapper struct{}

func (m Mapper) ReadToEntities(information []model.Cell) []entity.Cell {
	response := make([]entity.Cell, len(information))
	for i := range information {
		response[i] = entity.Cell{
			CellPosition: information[i].CellPosition,
			Information:  information[i].Information,
		}
	}
	return response
}
