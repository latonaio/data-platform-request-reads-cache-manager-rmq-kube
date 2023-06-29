package dpfm_api_output_formatter

type ProductionVersionList struct {
	ProductionVersions []ProductionVersion `json:"ProductionVersions"`
}

type ProductionVersion struct {
	Product             *string `json:"Product"`
	ProductionVersion   int     `json:"ProductionVersion"`
	ProductDescription  *string `json:"ProductDescription"`
	OwnerPlant          *string `json:"OwnerPlant"`
	OwnerPlantName      *string `json:"OwnerPlantName"`
	BillOfMaterial      *int    `json:"BillOfMaterial"`
	Operations          *int    `json:"Operations"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	ValidityStartDate   *string `json:"ValidityStartDate"`
	Images              Images  `json:"Images"`
}
