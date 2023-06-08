package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type PriceMasterDetailRes struct {
	ConnectionKey       string                   `json:"connection_key"`
	Result              bool                     `json:"result"`
	RedisKey            string                   `json:"redis_key"`
	Filepath            string                   `json:"filepath"`
	APIStatusCode       int                      `json:"api_status_code"`
	RuntimeSessionID    string                   `json:"runtime_session_id"`
	BusinessPartnerID   *int                     `json:"business_partner"`
	ServiceLabel        string                   `json:"service_label"`
	APIType             string                   `json:"api_type"`
	Message             PriceMasterDetailMessage `json:"message"`
	APISchema           string                   `json:"api_schema"`
	Accepter            []string                 `json:"accepter"`
	Deleted             bool                     `json:"deleted"`
	SQLUpdateResult     *bool                    `json:"sql_update_result"`
	SQLUpdateError      string                   `json:"sql_update_error"`
	SubfuncResult       *bool                    `json:"subfunc_result"`
	SubfuncError        string                   `json:"subfunc_error"`
	ExconfResult        *bool                    `json:"exconf_result"`
	ExconfError         string                   `json:"exconf_error"`
	APIProcessingResult *bool                    `json:"api_processing_result"`
	APIProcessingError  string                   `json:"api_processing_error"`
}

type PriceMasterDetailMessage struct {
	PriceMasterDetail []PriceMasterDetail `json:"General"`
}

type PriceMasterDetail struct {
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID"`
	Buyer                      int      `json:"Buyer"`
	Seller                     int      `json:"Seller"`
	ConditionRecord            int      `json:"ConditionRecord"`
	ConditionSequentialNumber  int      `json:"ConditionSequentialNumber"`
	ConditionValidityEndDate   string   `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate string   `json:"ConditionValidityStartDate"`
	Product                    *string  `json:"Product"`
	ConditionType              *string  `json:"ConditionType"`
	CreationDate               *string  `json:"CreationDate"`
	LastChangeDate             *string  `json:"LastChangeDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit     *string  `json:"ConditionRateValueUnit"`
	ConditionScaleQuantity     *float32 `json:"ConditionScaleQuantity"`
	ConditionRateRatio         *float32 `json:"ConditionRateRatio"`
	ConditionRateRatioUnit     *string  `json:"ConditionRateRatioUnit"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	BaseUnit                   *string  `json:"BaseUnit"`
	ConditionIsDeleted         *bool    `json:"ConditionIsDeleted"`
}

func CreatePriceMasterDetailRes(msg rabbitmq.RabbitmqMessage) (*PriceMasterDetailRes, error) {
	res := PriceMasterDetailRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
