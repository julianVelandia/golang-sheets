package contract

type URLParams struct {
	SheetName   string `json:"sheet_name" binding:"required"`
	StartColumn string `json:"start_column" binding:"required"`
	EndColumn   string `json:"end_column" binding:"required"`
	StartRow    int    `json:"start_row" binding:"required"`
	EndRow      int    `json:"end_row" binding:"required"`
}
