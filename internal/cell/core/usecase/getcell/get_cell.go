package getcell

import (
	"context"
	"fmt"
	"github.com/julianVelandia/golang-sheets/internal/cell/core/entity"
	ErrorUseCase "github.com/julianVelandia/golang-sheets/internal/cell/core/error"
	"github.com/julianVelandia/golang-sheets/internal/cell/core/query"
	"github.com/julianVelandia/golang-sheets/internal/platform/constant"
	"github.com/julianVelandia/golang-sheets/internal/platform/log"
	"github.com/julianVelandia/golang-sheets/internal/platform/sheets/model"
)

const (
	action              string          = "execute_use_case"
	errorReadRepository log.LogsMessage = "error in the use case, when read repository"
	entityType          string          = "read_repository"
	layer               string          = "use_case_get_Cells"
)

type RepositoryRead interface {
	GetByQuery(ctx context.Context, queryValue query.GetCells) ([]model.Cell, error)
}

type Mapper interface {
	ReadToEntities(information []model.Cell) []entity.Cell
}

type UseCase struct {
	repositoryRead RepositoryRead
	mapper         Mapper
}

func NewUseCase(repositoryRead RepositoryRead, mapper Mapper) UseCase {
	return UseCase{
		repositoryRead: repositoryRead,
		mapper:         mapper,
	}
}

func (uc UseCase) Execute(ctx context.Context, GetCells query.GetCells) ([]entity.Cell, error) {

	Cell, err := uc.repositoryRead.GetByQuery(ctx, GetCells)

	if err != nil {
		message := errorReadRepository.GetMessageWithTagParams(
			log.NewTagParams(layer, action,
				log.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s`,
						GetCells.SheetName,
						GetCells.Value(),
					),
					constant.EntityType: entityType,
				}))
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     err,
		}
	}

	return uc.mapper.ReadToEntities(Cell), nil
}
