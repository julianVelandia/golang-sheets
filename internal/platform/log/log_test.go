package log_test

import (
	"fmt"
	"github.com/julianVelandia/GolangSheets/internal/platform/constant"
	"github.com/julianVelandia/GolangSheets/internal/platform/log"
	"testing"
)

const (
	action               string          = "action"
	entityType           string          = "Account"
	errorCardInformation log.LogsMessage = "error consulting get cell."
	key                  string          = "eureka"
	layer                string          = "use_case"
)

func TestGetMessageWhenExecuteShouldReturnString(t *testing.T) {
	expectedMessage := "error consulting get cell."

	message := errorCardInformation.GetMessage()

	assert.Equal(t, expectedMessage, message)
}

func TestGetMessageWhenExecuteShouldReturnTagWhitParams(t *testing.T) {
	expectedMsj := "error consulting get cell.  layer:%s action:%s entity_type:%s key:%s"
	msj := errorCardInformation.GetMessageWithTagParams(
		log.NewTagParams(layer, action,
			log.Params{
				constant.EntityType: entityType,
				constant.Key:        key,
			}))

	assert.Equal(t, msj, fmt.Sprintf(
		expectedMsj,
		layer,
		action,
		entityType,
		key),
	)
}
