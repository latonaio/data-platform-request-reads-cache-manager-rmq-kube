package apiOutputFormatter

type BillOfMaterial struct {
	BillOfMaterialHeader []BillOfMaterialHeader `json:"Header"`
	BillOfMaterialItem   []BillOfMaterialItem   `json:"Item"`
}

type BillOfMaterialHeader struct {
	Product                  string  `json:"Product"`
	BillOfMaterial           int     `json:"BillOfMaterial"`
	ProductDescription       *string `json:"ProductDescription"`
	OwnerProductionPlant     string  `json:"OwnerProductionPlant"`
	OwnerProductionPlantName *string `json:"OwnerProductionPlantName"`
	ValidityStartDate        *string `json:"ValidityStartDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
	Images                   Images  `json:"Images"`
}

type BillOfMaterialItem struct {
	ComponentProduct                           string   `json:"ComponentProduct"`
	BillOfMaterialItem                         int      `json:"BillOfMaterialItem"`
	BillOfMaterialItemText                     string   `json:"BillOfMaterialItemText"`
	StockConfirmationPlant                     *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantName                 *string  `json:"StockConfirmationPlantName"`
	ComponentProductStandardQuantityInBaseUnit *float32 `json:"ComponentProductStandardQuantityInBaseUnit"`
	ComponentProductBaseUnit                   *string  `json:"ComponentProductBaseUnit"`
	ValidityStartDate                          *string  `json:"ValidityStartDate"`
	IsMarkedForDeletion                        *bool    `json:"IsMarkedForDeletion"`
}
