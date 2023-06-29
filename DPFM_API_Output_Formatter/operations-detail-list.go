package dpfm_api_output_formatter

type OperationsDetailList struct {
	OperationsDetailHeader OperationsDetailHeader `json:"OperationsDetailHeader"`
	OperationsDetail       []OperationsDetail     `json:"OperationsDetail"`
}

type OperationsDetail struct {
	OperationsItem          int      `json:"OperationsItem"`
	OperationsText          *string  `json:"OperationsText"`
	ProductionPlantName     *string  `json:"ProductionPlantName"`
	StandardLotSizeQuantity *float32 `json:"StandardLotSizeQuantity"`
	OperationsUnit          *string  `json:"OperationsUnit"`
	ValidityStartDate       *string  `json:"ValidityStartDate"`
	IsMarkedForDeletion     *bool    `json:"IsMarkedForDeletion"`
}

type OperationsDetailHeader struct {
	Index int    `json:"Index"`
	Key   string `json:"Key"`
}
