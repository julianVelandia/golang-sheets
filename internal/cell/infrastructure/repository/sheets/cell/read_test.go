package cell_test

import (
	"context"
	"errors"
	"testing"

	"github.com/julianVelandia/golang-sheets/internal/cell/core/query"
	"github.com/julianVelandia/golang-sheets/internal/cell/infrastructure/repository/sheets/cell"
	"github.com/julianVelandia/golang-sheets/internal/platform/sheets/model"
	"github.com/stretchr/testify/assert"
)

const (
	cellPositionA1    string = "A1"
	cellPositionA2    string = "A1"
	endColumn         string = "B"
	endRow            int    = 2
	informationA1     string = "informationA1"
	informationA2     string = "informationA1"
	sheetName         string = "name"
	startColumn       string = "A"
	startRow          int    = 1
	pathCredentials   string = "internal/platform/sheets/environment/credentials.json"
	pathSpreadsheetID string = "internal/platform/sheets/environment/spreadsheetID.json"
)

func TestGetByQueryWhenRepositoryOKShouldResponseSuccess(t *testing.T) {
	ctx := context.TODO()
	clientSheetsMock := new(ClientSheetsMock)
	queryGetCell := givenQuery()
	models := givenModels()

	clientSheetsMock.On("Read", ctx, pathCredentials, pathSpreadsheetID, queryGetCell).Return(models, nil).Once()
	repository := cell.NewSheetsRepository(clientSheetsMock)
	modelsResult, errResult := repository.GetByQuery(ctx, queryGetCell)

	assert.Equal(t, modelsResult, models)
	assert.Nil(t, errResult)
	clientSheetsMock.AssertExpectations(t)
}

func TestGetByQueryWhenRepositoryFailShouldResponseError(t *testing.T) {
	ctx := context.TODO()
	clientSheetsMock := new(ClientSheetsMock)
	queryGetCell := givenQuery()
	errorReadRepository := errors.New(
		"error in the use case, when read repository layer:use_case_get_Cells action:execute_use_case " +
			"entity_type:read_repository " +
			"key:internal/platform/sheets/environment/credentials.json_internal/platform/sheets/environment/spreadsheetID.json_name!A1:B2",
	)

	clientSheetsMock.On(
		"Read",
		ctx,
		pathCredentials,
		pathSpreadsheetID,
		queryGetCell,
	).Return(
		nil,
		errorReadRepository,
	).Once()
	repository := cell.NewSheetsRepository(clientSheetsMock)
	modelsResult, errResult := repository.GetByQuery(ctx, queryGetCell)

	assert.Nil(t, modelsResult)
	assert.EqualError(t, errResult, errorReadRepository.Error())
	clientSheetsMock.AssertExpectations(t)
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
