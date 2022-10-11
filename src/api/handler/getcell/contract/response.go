package contract

type Response struct {
	Cells []Cell `json:"cells"`
}

type Cell struct {
	CellPosition string `json:"cell_position"`
	Information  string `json:"information"`
}
