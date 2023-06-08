package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type PaymentTermsRes struct {
	ConnectionKey       string              `json:"connection_key"`
	Result              bool                `json:"result"`
	RedisKey            string              `json:"redis_key"`
	Filepath            string              `json:"filepath"`
	APIStatusCode       int                 `json:"api_status_code"`
	RuntimeSessionID    string              `json:"runtime_session_id"`
	BusinessPartnerID   *int                `json:"business_partner"`
	ServiceLabel        string              `json:"service_label"`
	APIType             string              `json:"api_type"`
	Message             PaymentTermsMessage `json:"message"`
	APISchema           string              `json:"api_schema"`
	Accepter            []string            `json:"accepter"`
	Deleted             bool                `json:"deleted"`
	SQLUpdateResult     *bool               `json:"sql_update_result"`
	SQLUpdateError      string              `json:"sql_update_error"`
	SubfuncResult       *bool               `json:"subfunc_result"`
	SubfuncError        string              `json:"subfunc_error"`
	ExconfResult        *bool               `json:"exconf_result"`
	ExconfError         string              `json:"exconf_error"`
	APIProcessingResult *bool               `json:"api_processing_result"`
	APIProcessingError  string              `json:"api_processing_error"`
}

type PaymentTermsMessage struct {
	PaymentTerms     []PaymentTerms     `json:"PaymentTerms"`
	PaymentTermsText []PaymentTermsText `json:"PaymentTermsText"`
}

type PaymentTerms struct {
	PaymentTerms                string `json:"PaymentTerms"`
	BaseDate                    int    `json:"BaseDate"`
	BaseDateCalcAddMonth        *int   `json:"BaseDateCalcAddMonth"`
	BaseDateCalcFixedDate       *int   `json:"BaseDateCalcFixedDate"`
	PaymentDueDateCalcAddMonth  *int   `json:"PaymentDueDateCalcAddMonth"`
	PaymentDueDateCalcFixedDate *int   `json:"PaymentDueDateCalcFixedDate"`
}

type PaymentTermsText struct {
	PaymentTerms     string  `json:"PaymentTerms"`
	Language         string  `json:"Language"`
	PaymentTermsName *string `json:"PaymentTermsName"`
}

func CreatePaymentTermsRes(msg rabbitmq.RabbitmqMessage) (*PaymentTermsRes, error) {
	res := PaymentTermsRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
