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
	BillOfMaterialItem                          int      `json:"BillOfMaterialItem"`
	ComponentProduct                            *string  `json:"ComponentProduct"`
	BillOfMaterialItemText                      *string  `json:"BillOfMaterialItemText"`
	StockConfirmationPlantName                  *string  `json:"StockConfirmationPlantName"`
	StockConfirmationPlant                      *string  `json:"StockConfirmationPlant"`
	ComponentProductStandardQuantityInBaseUnuit *float32 `json:"ComponentProductStandardQuantityInBaseUnuit"`
	ComponentProductBaseUnit                    *string  `json:"ComponentProductBaseUnit"`
	ValidityStartDate                           *string  `json:"ValidityStartDate"`
	IsMarkedForDeletion                         *bool    `json:"IsMarkedForDeletion"`
}
