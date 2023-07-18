package apiOutputFormatter

type PriceMaster struct {
	PriceMasterHeader       []PriceMasterHeader       `json:"Header"`
	PriceMasterDetailHeader []PriceMasterDetailHeader `json:"DetailHeader"`
}

type PriceMasterHeader struct {
	SupplyChainRelationshipID int    `json:"SupplyChainRelationshipID"`
	Buyer                     int    `json:"Buyer"`
	BuyerName                 string `json:"BuyerName"`
	Seller                    int    `json:"Seller"`
	SellerName                string `json:"SellerName"`
}

type PriceMasterDetailHeader struct {
	Product                   string  `json:"Product"`
	ProductDescription        string  `json:"ProductDescription"`
	ConditionType             string  `json:"ConditionType"`
	ConditionRateValue        float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit    int     `json:"ConditionRateValueUnit"`
	ConditionScaleQuantity    int     `json:"ConditionScaleQuantity"`
	ConditionCurrency         string  `json:"ConditionCurrency"`
	ConditionRecord           int     `json:"ConditionRecord"`
	ConditionSequentialNumber int     `json:"ConditionSequentialNumber"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
	Images                    Images  `json:"Images"`
}
