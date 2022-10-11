package cell

import (
	"context"
	"fmt"

	ErrorUseCase "github.com/julianVelandia/GolangSheets/internal/cell/core/error"
	"github.com/julianVelandia/GolangSheets/internal/platform/constant"
	"github.com/julianVelandia/GolangSheets/internal/platform/log"
	sheet "github.com/julianVelandia/GolangSheets/internal/platform/sheets"
)

const (
	action              string          = "execute_use_case"
	errorReadRepository log.LogsMessage = "error in the use case, when read repository"
	entityType          string          = "read_repository"
	pathCredentials     string          = "internal/platform/sheets/environment/credentials.json"
	pathSpreadsheetID   string          = "1y8rAuDeYfcT7jx6O1KbC_DwO68SeWeoXkeUdYVMCV_0"
	layer               string          = "use_case_get_Cells"
)

//TODO Leer el spreatsheetIDD desde path

type RepositoryClient struct {
	client sheet.Client
}

func NewRepositoryClient(client sheet.Client) *RepositoryClient {
	return &RepositoryClient{client: client}
}

func (rc RepositoryClient) GetByQuery(ctx context.Context, queryValue string) ([]string, error) {

	item, errReadClient := rc.client.Read(ctx, pathCredentials, pathSpreadsheetID, queryValue)

	if errReadClient != nil {
		message := errorReadRepository.GetMessageWithTagParams(
			log.NewTagParams(layer, action,
				log.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s`,
						pathCredentials,
						pathSpreadsheetID,
						queryValue,
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
