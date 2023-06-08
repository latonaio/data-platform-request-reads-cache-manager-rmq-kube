package models

type PaymentTermsReq struct {
	ConnectionKey     string       `json:"connection_key"`
	Result            bool         `json:"result"`
	RedisKey          string       `json:"redis_key"`
	Filepath          string       `json:"filepath"`
	APIStatusCode     int          `json:"api_status_code"`
	RuntimeSessionID  string       `json:"runtime_session_id"`
	BusinessPartnerID *int         `json:"business_partner"`
	ServiceLabel      string       `json:"service_label"`
	APIType           string       `json:"api_type"`
	PaymentTerms      PaymentTerms `json:"PaymentTerms"`
	APISchema         string       `json:"api_schema"`
	Accepter          []string     `json:"accepter"`
	Deleted           bool         `json:"deleted"`
}

type PaymentTerms struct {
	PaymentTerms                string             `json:"PaymentTerms"`
	BaseDate                    int                `json:"BaseDate"`
	BaseDateCalcAddMonth        *int               `json:"BaseDateCalcAddMonth"`
	BaseDateCalcFixedDate       *int               `json:"BaseDateCalcFixedDate"`
	PaymentDueDateCalcAddMonth  *int               `json:"PaymentDueDateCalcAddMonth"`
	PaymentDueDateCalcFixedDate *int               `json:"PaymentDueDateCalcFixedDate"`
	PaymentTermsText            []PaymentTermsText `json:"PaymentTermsText"`
}

type PaymentTermsText struct {
	PaymentTerms     string  `json:"PaymentTerms"`
	Language         string  `json:"Language"`
	PaymentTermsName *string `json:"PaymentTermsName"`
}
