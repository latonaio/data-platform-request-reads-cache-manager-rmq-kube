package dpfm_api_output_formatter

type OperationsList struct {
	Operations []Operations `json:"Operations"`
}

type Operations struct {
	Operations               int     `json:"Operations"`
	Product                  *string `json:"Product"`
	ProductDescription       *string `json:"ProductDescription"`
	OwnerProductionPlantName *string `json:"OwnerProductionPlantName"`
	ValidityStartDate        *string `json:"ValidityStartDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
	Images                   Images  `json:"Images"`
}
