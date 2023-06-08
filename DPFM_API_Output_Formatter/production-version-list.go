package dpfm_api_output_formatter

type ProductionVersionList struct {
	ProductionVersions []ProductionVersion `json:"ProductionVersions"`
}

type ProductionVersion struct {
	Product             *string `json:"Product"`
	ProductionVersion   int     `json:"ProductionVersion"`
	ProductDescription  *string `json:"ProductDescription"`
	OwnerPlant          *string `json:"OwnerPlant"`
	BillOfMaterial      *int    `json:"BillOfMaterial"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Images              Images  `json:"Images"`
}
