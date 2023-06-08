package models

type CurrencyReq struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"api_type"`
	Currency          Currency `json:"Currency"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type Currency struct {
	Currency     string         `json:"Currency"`
	CurrencyText []CurrencyText `json:"CurrencyText"`
}

type CurrencyText struct {
	Currency         string  `json:"Currency"`
	Language         string  `json:"Language"`
	CurrencyName     *string `json:"CurrencyName"`
	CurrencyLongName *string `json:"CurrencyLongName"`
}
