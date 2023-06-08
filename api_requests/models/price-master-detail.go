package models

type PriceMasterDetailReq struct {
	ConnectionKey     string            `json:"connection_key"`
	Result            bool              `json:"result"`
	RedisKey          string            `json:"redis_key"`
	Filepath          string            `json:"filepath"`
	APIStatusCode     int               `json:"api_status_code"`
	RuntimeSessionID  string            `json:"runtime_session_id"`
	BusinessPartnerID *int              `json:"business_partner"`
	ServiceLabel      string            `json:"service_label"`
	APIType           string            `json:"api_type"`
	PriceMasterDetail PriceMasterDetail `json:"PriceMaster"`
	APISchema         string            `json:"api_schema"`
	Accepter          []string          `json:"accepter"`
	Deleted           bool              `json:"deleted"`
}

type PriceMasterDetail struct {
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID"`
	Buyer                      int      `json:"Buyer"`
	Seller                     int      `json:"Seller"`
	ConditionRecordCategory    string   `json:"ConditionRecordCategory"`
	ConditionRecord            int      `json:"ConditionRecord"`
	ConditionSequentialNumber  int      `json:"ConditionSequentialNumber"`
	ConditionValidityEndDate   string   `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate string   `json:"ConditionValidityStartDate"`
	Product                    string   `json:"Product"`
	ConditionType              string   `json:"ConditionType"`
	CreationDate               *string  `json:"CreationDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit     *string  `json:"ConditionRateValueUnit"`
	ConditionRateRatio         *float32 `json:"ConditionRateRatio"`
	ConditionRateRatioUnit     *string  `json:"ConditionRateRatioUnit"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	BaseUnit                   *string  `json:"BaseUnit"`
	ConditionIsDeleted         *bool    `json:"ConditionIsDeleted"`
}
