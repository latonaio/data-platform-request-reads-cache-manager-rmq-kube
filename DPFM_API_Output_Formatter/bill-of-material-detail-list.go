package dpfm_api_output_formatter

type BillOfMaterialDetailList struct {
	BillOfMaterialDetailHeader BillOfMaterialDetailHeader `json:"BillOfMaterialDetailHeader"`
	BillOfMaterialDetail       []BillOfMaterialDetail     `json:"BillOfMaterialDetails"`
}

type BillOfMaterialDetailHeader struct {
	Index int    `json:"Index"`
	Key   string `json:"Key"`
}

type BillOfMaterialDetail struct {
	ComponentProduct          string  `json:"ComponentProject"`
	BillOfMaterialItemText    string  `json:"BillOfMaterialItemText"`
	StockConfirmationPlant    string  `json:"StockConfirmationPlant"`
	BOMItemQuantityInBaseUnit float32 `json:"BOMItemQuantityInBaseUnit"`
	BOMItemBaseUnit           string  `json:"BOMItemBaseUnit"`
	ValidityStartDate         string  `json:"ValidityStartDate"`
	IsMarkedForDeletion       bool    `json:"IsMarkedForDeletion"`
}
