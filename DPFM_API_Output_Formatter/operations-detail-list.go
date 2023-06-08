package dpfm_api_output_formatter

type OperationsDetailList struct {
	Operations []OperationsDetail       `json:"Operations"`
	Details    []OperationsDetailHeader `json:"Details"`
}

type OperationsDetail struct {
	Operations              int     `json:"Operations"`
	Product                 *string `json:"Product"`
	OperationsText          *string `json:"OperationsText"`
	Plant                   *string `json:"Plant"`
	StandardLotSizeQuantity *string `json:"StandardLotSizeQuantity"`
	OperationsUnit          *string `json:"OperationsUnit"`
	ValidityStartDate       *string `json:"ValidityStartDate"`
	IsMarkedForDeletion     *bool   `json:"IsMarkedForDeletion"`
	Images                  Images  `json:"Images"`
}

type OperationsDetailHeader struct {
	Index int    `json:"Index"`
	Key   string `json:"Key"`
}
