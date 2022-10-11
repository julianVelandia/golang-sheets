package cell

import (
	"context"
	"fmt"
	"github.com/julianVelandia/golang-sheets/internal/cell/core/query"
	"github.com/julianVelandia/golang-sheets/internal/platform/sheets/model"

	ErrorUseCase "github.com/julianVelandia/golang-sheets/internal/cell/core/error"
	"github.com/julianVelandia/golang-sheets/internal/platform/constant"
	"github.com/julianVelandia/golang-sheets/internal/platform/log"
	sheet "github.com/julianVelandia/golang-sheets/internal/platform/sheets"
)

const (
	action              string          = "execute_use_case"
	errorReadRepository log.LogsMessage = "error in the use case, when read repository"
	entityType          string          = "read_repository"
	pathCredentials     string          = "internal/platform/sheets/environment/credentials.json"
	pathSpreadsheetID   string          = "internal/platform/sheets/environment/spreadsheetID.json"
	layer               string          = "use_case_get_Cells"
)

type RepositoryClient struct {
	client sheet.Client
}

func NewRepositoryClient(client sheet.Client) *RepositoryClient {
	return &RepositoryClient{client: client}
}

func (rc RepositoryClient) GetByQuery(ctx context.Context, queryValue query.GetCells) ([]model.Cell, error) {

	item, errReadClient := rc.client.Read(ctx, pathCredentials, pathSpreadsheetID, queryValue)

	if errReadClient != nil {
		message := errorReadRepository.GetMessageWithTagParams(
			log.NewTagParams(layer, action,
				log.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s`,
						pathCredentials,
						pathSpreadsheetID,
						queryValue.Value(),
					),
					constant.EntityType: entityType,
				}))
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     errReadClient,
		}
	}

	return item, nil
}
