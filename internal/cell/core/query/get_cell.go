package query

import (
	"fmt"
	"strings"
)

type GetCells struct {
	SheetName   string
	StartColumn string
	EndColumn   string
	StartRow    int
	EndRow      int
}

func (g GetCells) Value() string {
	return fmt.Sprintf(
		`%s!%s:%s`,
		g.SheetName,
		strings.ToUpper(strings.TrimSpace(g.StartColumn))+string(rune(g.StartRow)),
		strings.ToUpper(strings.TrimSpace(g.EndColumn))+string(rune(g.EndRow)),
	)
}
