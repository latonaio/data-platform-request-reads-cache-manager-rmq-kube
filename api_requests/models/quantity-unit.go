package models

type QuantityUnitReq struct {
	ConnectionKey     string       `json:"connection_key"`
	Result            bool         `json:"result"`
	RedisKey          string       `json:"redis_key"`
	Filepath          string       `json:"filepath"`
	APIStatusCode     int          `json:"api_status_code"`
	RuntimeSessionID  string       `json:"runtime_session_id"`
	BusinessPartnerID *int         `json:"business_partner"`
	ServiceLabel      string       `json:"service_label"`
	APIType           string       `json:"api_type"`
	QuantityUnit      QuantityUnit `json:"QuantityUnit"`
	APISchema         string       `json:"api_schema"`
	Accepter          []string     `json:"accepter"`
	Deleted           bool         `json:"deleted"`
}

type QuantityUnit struct {
	QuantityUnit     string             `json:"QuantityUnit"`
	QuantityUnitText []QuantityUnitText `json:"QuantityUnitText"`
}

type QuantityUnitText struct {
	QuantityUnit     string  `json:"QuantityUnit"`
	Language         string  `json:"Language"`
	QuantityUnitName *string `json:"QuantityUnitName"`
}
