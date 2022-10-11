package contract

type URLParams struct {
	SheetName   string `json:"sheet_name" binding:"required" validate:"required"`
	StartColumn string `json:"start_column" binding:"required" validate:"required"`
	StartRow    string `json:"start_row" binding:"required" validate:"required"`
	EndColumn   string `json:"end_column" binding:"required" validate:"required"`
	EndRow      string `json:"end_row" binding:"required" validate:"required"`
}
