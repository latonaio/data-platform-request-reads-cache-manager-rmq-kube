package apiOutputFormatter

type Operations struct {
	OperationsHeader []OperationsHeader `json:"Header"`
	OperationsItem   []OperationsItem   `json:"Item"`
}

type OperationsHeader struct {
	Operations               int     `json:"Operations"`
	Product                  *string `json:"Product"`
	ProductDescription       *string `json:"ProductDescription"`
	OwnerProductionPlant     *string `json:"OwnerProductionPlant"`
	OwnerProductionPlantName *string `json:"OwnerProductionPlantName"`
	ValidityStartDate        *string `json:"ValidityStartDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
	Images                   Images  `json:"Images"`
}

type OperationsItem struct {
	OperationsItem          int      `json:"OperationsItem"`
	OperationsText          string   `json:"OperationsText"`
	ProductionPlant         string   `json:"ProductionPlant"`
	ProductionPlantName     string   `json:"ProductionPlantName"`
	StandardLotSizeQuantity *float32 `json:"StandardLotSizeQuantity"`
	OperationsUnit          *string  `json:"OperationsUnit"`
	ValidityStartDate       *string  `json:"ValidityStartDate"`
	IsMarkedForDeletion     *bool    `json:"IsMarkedForDeletion"`
}
