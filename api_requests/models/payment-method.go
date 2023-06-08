package models

type PaymentMethodReq struct {
	ConnectionKey     string        `json:"connection_key"`
	Result            bool          `json:"result"`
	RedisKey          string        `json:"redis_key"`
	Filepath          string        `json:"filepath"`
	APIStatusCode     int           `json:"api_status_code"`
	RuntimeSessionID  string        `json:"runtime_session_id"`
	BusinessPartnerID *int          `json:"business_partner"`
	ServiceLabel      string        `json:"service_label"`
	APIType           string        `json:"api_type"`
	PaymentMethod     PaymentMethod `json:"PaymentMethod"`
	APISchema         string        `json:"api_schema"`
	Accepter          []string      `json:"accepter"`
	Deleted           bool          `json:"deleted"`
}

type PaymentMethod struct {
	PaymentMethod     string              `json:"PaymentMethod"`
	PaymentMethodText []PaymentMethodText `json:"PaymentMethodText"`
}

type PaymentMethodText struct {
	PaymentMethod     string  `json:"PaymentMethod"`
	Language          string  `json:"Language"`
	PaymentMethodName *string `json:"PaymentMethodName"`
}
