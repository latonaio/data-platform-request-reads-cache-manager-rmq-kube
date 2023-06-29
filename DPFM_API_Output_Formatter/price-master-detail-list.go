package dpfm_api_output_formatter

type PriceMasterDetailList struct {
	PriceMasterDetailHeader PriceMasterDetailHeader `json:"PriceMasterDetailHeader"`
	PriceMasterDetail       []PriceMasterDetail     `json:"PriceMasterDetail"`
}

type PriceMasterDetailHeader struct {
	Index int    `json:"Index"`
	Key   string `json:"Key"`
}

type PriceMasterDetail struct {
	Product                   string  `json:"Product"`
	ProductionDescription     string  `json:"ProductDescription"`
	ConditionType             *string `json:"ConditionType"`
	ConditionRateValue        float32 `json:"ConditionRateValue"`
	ConditionScaleQuantity    *int    `json:"ConditionScaleQuantity"`
	ConditionRateValueUnit    *int    `json:"ConditionRateValueUnit"`
	ConditionCurrency         *string `json:"ConditionCurrency"`
	ConditionRecord           int     `json:"ConditionRecord"`
	ConditionSequentialNumber int     `json:"ConditionSequentialNumber"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}
