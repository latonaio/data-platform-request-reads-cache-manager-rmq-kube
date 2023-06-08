package dpfm_api_output_formatter

type PriceMasterDetailList struct {
	PriceMasterDetail []PriceMasterDetail `json:"PriceMasterDetail"`
}
type PriceMasterDetail struct {
	Product                   string  `json:"Product"`
	ProductionDescription     string  `json:"ProductDescription"`
	ConditionRateValue        float32 `json:"ConditionRateValue"`
	ConditionScaleQuantity    float32 `json:"ConditionScaleQuantity"`
	BaseUnit                  string  `json:"BaseUnit"`
	ConditionCurrency         string  `json:"ConditionCurrency"`
	ConditionRecord           int     `json:"ConditionRecord"`
	ConditionSequentialNumber int     `json:"ConditionSequentialNumber"`
	IsMarkedForDeletion       bool    `json:"IsMarkedForDeletion"`
}
