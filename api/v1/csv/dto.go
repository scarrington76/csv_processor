package csv

type requestCSV struct {
	Filename *string `json:"filename"`
}

type requestRow struct {
	Name        *string  `json:"name"`
	Price       *float64 `json:"price"`
	Description *string  `json:"description"`
}
