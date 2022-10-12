package mapper

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/julianVelandia/golang-sheets/internal/platform/sheets/model"
)

const (
	CellPositionA1 string = "A1"
	InformationA1  string = "informationA1"
	CellPositionA2 string = "A1"
	InformationA2  string = "informationA1"
)

func TestMapperModelToEntitiesWhenOk(t *testing.T) {
	modelCells := givenModel()
	mapper := Mapper{}

	execution := mapper.ModelToEntities(modelCells)
	assert.Equal(t, modelCells[0].CellPosition, execution[0].CellPosition)
	assert.Equal(t, modelCells[0].Information, execution[0].Information)
	assert.Equal(t, modelCells[1].CellPosition, execution[1].CellPosition)
	assert.Equal(t, modelCells[1].Information, execution[1].Information)
}

func givenModel() []model.Cell {

	modelCells := make([]model.Cell, 0)
	modelCells = append(
		modelCells,
		model.Cell{
			CellPosition: CellPositionA1,
			Information:  InformationA1,
		},
		model.Cell{
			CellPosition: CellPositionA2,
			Information:  InformationA2,
		})
	return modelCells
}
