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
		fmt.Sprintf("%v%v", strings.ToUpper(g.StartColumn), g.StartRow),
		fmt.Sprintf("%v%v", strings.ToUpper(g.EndColumn), g.EndRow),
	)
}
