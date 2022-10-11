package getcell

import (
	"context"
	"fmt"
	"github.com/julianVelandia/golang-sheets/internal/cell/core/entity"
	ErrorUseCase "github.com/julianVelandia/golang-sheets/internal/cell/core/error"
	"github.com/julianVelandia/golang-sheets/internal/cell/core/query"
	"github.com/julianVelandia/golang-sheets/internal/platform/constant"
	"github.com/julianVelandia/golang-sheets/internal/platform/log"
)

const (
	action              string          = "execute_use_case"
	errorReadRepository log.LogsMessage = "error in the use case, when read repository"
	entityType          string          = "read_repository"
	layer               string          = "use_case_get_Cells"
)

type RepositoryRead interface {
	GetByQuery(ctx context.Context, queryValue string) ([]entity.Cell, error)
}

// TODO armar y ddevolver la entidad o moddelo (mirar que no rompa la arqu)
type Mapper interface {
	ReadToEntities(ctx context.Context, information []string) []entity.Cell
}

type UseCase struct {
	repositoryRead RepositoryRead
}

func NewUseCase(repositoryRead RepositoryRead) UseCase {
	return UseCase{repositoryRead: repositoryRead}
}

func (uc UseCase) Execute(ctx context.Context, GetCells query.GetCells) ([]entity.Cell, error) {

	Cell, err := uc.repositoryRead.GetByQuery(ctx, GetCells.Value())

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

	return Cell, nil
}
