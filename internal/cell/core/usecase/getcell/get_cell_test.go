package getcell_test

import (
	"context"
	"errors"
	"github.com/julianVelandia/golang-sheets/internal/cell/core/entity"
	"github.com/julianVelandia/golang-sheets/internal/cell/core/query"
	"github.com/julianVelandia/golang-sheets/internal/cell/core/usecase/getcell"
	"github.com/julianVelandia/golang-sheets/internal/platform/sheets/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	cellPositionA1 string = "A1"
	cellPositionA2 string = "A1"
	endColumn      string = "B"
	endRow         int    = 2
	informationA1  string = "informationA1"
	informationA2  string = "informationA1"
	sheetName      string = "name"
	startColumn    string = "A"
	startRow       int    = 1
)

func TestExecuteWhenGetCellsSuccessfullyShouldReturnSuccess(t *testing.T) {
	repositoryReadMock := new(RepositoryReadMock)
	mapperMock := new(MapperMock)
	ctx := context.TODO()
	queryGetCell := givenQuery()
	models := givenModels()
	entitiesCell := givenEntities()

	repositoryReadMock.On("GetByQuery", ctx, queryGetCell).Return(models, nil).Once()
	mapperMock.On("ModelToEntities", models).Return(entitiesCell).Once()
	useCase := getcell.NewUseCase(repositoryReadMock, mapperMock)
	entities, errResult := useCase.Execute(ctx, queryGetCell)

	assert.Equal(t, entities, entitiesCell)
	assert.Nil(t, errResult)
	repositoryReadMock.AssertExpectations(t)
	mapperMock.AssertExpectations(t)
}

func TestExecuteWhenGetCellsFailShouldReturnError(t *testing.T) {
	repositoryReadMock := new(RepositoryReadMock)
	ctx := context.TODO()
	queryGetCell := givenQuery()
	errorReadRepository := errors.New(
		"error in the use case, when read repository layer:use_case_get_Cells " +
			"action:execute_use_case entity_type:read_repository key:name_name!A1:B2",
	)

	repositoryReadMock.On("GetByQuery", ctx, queryGetCell).Return(nil, errorReadRepository).Once()
	useCase := getcell.NewUseCase(repositoryReadMock, nil)
	entities, errResult := useCase.Execute(ctx, queryGetCell)

	assert.Nil(t, entities)
	assert.EqualError(t, errResult, errorReadRepository.Error())
	repositoryReadMock.AssertExpectations(t)
}

func givenQuery() query.GetCells {
	return query.GetCells{
		SheetName:   sheetName,
		StartColumn: startColumn,
		EndColumn:   endColumn,
		StartRow:    startRow,
		EndRow:      endRow,
	}
}

func givenModels() []model.Cell {
	modelCells := make([]model.Cell, 0)
	modelCells = append(
		modelCells,
		model.Cell{
			CellPosition: cellPositionA1,
			Information:  informationA1,
		},
		model.Cell{
			CellPosition: cellPositionA2,
			Information:  informationA2,
		})
	return modelCells
}

func givenEntities() []entity.Cell {
	entityCells := make([]entity.Cell, 0)
	entityCells = append(
		entityCells,
		entity.Cell{
			CellPosition: cellPositionA1,
			Information:  informationA1,
		},
		entity.Cell{
			CellPosition: cellPositionA2,
			Information:  informationA2,
		})
	return entityCells
}
