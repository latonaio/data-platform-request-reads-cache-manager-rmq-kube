package models

type PriceMasterReq struct {
	ConnectionKey     string      `json:"connection_key"`
	Result            bool        `json:"result"`
	RedisKey          string      `json:"redis_key"`
	Filepath          string      `json:"filepath"`
	APIStatusCode     int         `json:"api_status_code"`
	RuntimeSessionID  string      `json:"runtime_session_id"`
	BusinessPartnerID *int        `json:"business_partner"`
	ServiceLabel      string      `json:"service_label"`
	APIType           string      `json:"api_type"`
	PriceMaster       PriceMaster `json:"PriceMaster"`
	APISchema         string      `json:"api_schema"`
	Accepter          []string    `json:"accepter"`
	Deleted           bool        `json:"deleted"`
}

type PriceMaster struct {
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID"`
	Buyer                      int      `json:"Buyer"`
	Seller                     int      `json:"Seller"`
	ConditionRecord            int      `json:"ConditionRecord"`
	ConditionSequentialNumber  int      `json:"ConditionSequentialNumber"`
	ConditionValidityStartDate string   `json:"ConditionValidityStartDate"`
	ConditionValidityEndDate   string   `json:"ConditionValidityEndDate"`
	Product                    *string  `json:"Product"`
	ConditionType              *string  `json:"ConditionType"`
	CreationDate               *string  `json:"CreationDate"`
	LastChangeDate             *string  `json:"LastChangeDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit     *int     `json:"ConditionRateValueUnit"`
	ConditionScaleQuantity     *int     `json:"ConditionScaleQuantity"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	IsMarkedForDeletion        *bool    `json:"IsMarkedForDeletion"`
}
