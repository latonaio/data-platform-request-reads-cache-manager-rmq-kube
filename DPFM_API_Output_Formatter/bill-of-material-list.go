package dpfm_api_output_formatter

type BillOfMaterialList struct {
	BillOfMaterials []BillOfMaterial `json:"BillOfMaterials"`
}

type BillOfMaterial struct {
	Product                  string  `json:"Product"`
	BillOfMaterial           int     `json:"BillOfMaterial"`
	ProductDescription       *string `json:"ProductDescription"`
	OwnerProductionPlant     string  `json:"OwnerProductionPlant"`
	OwnerProductionPlantName *string `json:"OwnerProductionPlantName"`
	ValidityStartDate        *string `json:"ValidityStartDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
	Images                   Images  `json:"Images"`
}
