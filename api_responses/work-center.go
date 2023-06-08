package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type WorkCenterRes struct {
	ConnectionKey       string            `json:"connection_key"`
	Result              bool              `json:"result"`
	RedisKey            string            `json:"redis_key"`
	Filepath            string            `json:"filepath"`
	APIStatusCode       int               `json:"api_status_code"`
	RuntimeSessionID    string            `json:"runtime_session_id"`
	BusinessPartnerID   *int              `json:"business_partner"`
	ServiceLabel        string            `json:"service_label"`
	APIType             string            `json:"api_type"`
	Message             WorkCenterMessage `json:"message"`
	APISchema           string            `json:"api_schema"`
	Accepter            []string          `json:"accepter"`
	Deleted             bool              `json:"deleted"`
	SQLUpdateResult     *bool             `json:"sql_update_result"`
	SQLUpdateError      string            `json:"sql_update_error"`
	SubfuncResult       *bool             `json:"subfunc_result"`
	SubfuncError        string            `json:"subfunc_error"`
	ExconfResult        *bool             `json:"exconf_result"`
	ExconfError         string            `json:"exconf_error"`
	APIProcessingResult *bool             `json:"api_processing_result"`
	APIProcessingError  string            `json:"api_processing_error"`
}

type WorkCenterMessage struct {
	General *[]General `json:"General"`
}

type General struct {
	WorkCenter                   int     `json:"WorkCenter"`
	WorkCenterType               string  `json:"WorkCenterType"`
	WorkCenterName               string  `json:"WorkCenterName"`
	BusinessPartner              int     `json:"BusinessPartner"`
	Plant                        string  `json:"Plant"`
	WorkCenterCategory           *string `json:"WorkCenterCategory"`
	WorkCenterResponsible        *string `json:"WorkCenterResponsible"`
	SupplyArea                   *string `json:"SupplyArea"`
	WorkCenterUsage              *string `json:"WorkCenterUsage"`
	MatlCompIsMarkedForBackflush *bool   `json:"MatlCompIsMarkedForBackflush"`
	WorkCenterLocation           *string `json:"WorkCenterLocation"`
	CapacityInternalID           *string `json:"CapacityInternalID"`
	CapacityCategoryCode         *string `json:"CapacityCategoryCode"`
	ValidityStartDate            *string `json:"ValidityStartDate"`
	ValidityEndDate              *string `json:"ValidityEndDate"`
	IsMarkedForDeletion          *bool   `json:"IsMarkedForDeletion"`
}

func CreateWorkCenterRes(msg rabbitmq.RabbitmqMessage) (*WorkCenterRes, error) {
	res := WorkCenterRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
