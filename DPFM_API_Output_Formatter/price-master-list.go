package dpfm_api_output_formatter

type PriceMasterList struct {
	PriceMasters []PriceMasters `json:"PriceMasters"`
}
type PriceMasters struct {
	SupplyChainRelationshipID int    `json:"SupplyChainRelationshipID"`
	Buyer                     *int   `json:"Buyer"`
	BuyerName                 string `json:"BuyerName"`
	Seller                    *int   `json:"Seller"`
	SellerName                string `json:"SellerName"`
}

// type PriceMasters struct {
// 	Product                   string  `json:"Product"`
// 	ProductionDescription     string  `json:"ProductDescription"`
// 	ConditionRateValue        float32 `json:"ConditionRateValue"`
// 	ConditionScaleQuantity    float32 `json:"ConditionScaleQuantity"`
// 	BaseUnit                  string  `json:"BaseUnit"`
// 	ConditionCurrency         string  `json:"ConditionCurrency"`
// 	ConditionRecord           int     `json:"ConditionRecord"`
// 	ConditionSequentialNumber int     `json:"ConditionSequentialNumber"`
// 	IsMarkedForDeletion       bool    `json:"IsMarkedForDeletion"`
// }
