package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type StorageBinRes struct {
	ConnectionKey       string     `json:"connection_key,omitempty"`
	Result              bool       `json:"result,omitempty"`
	RedisKey            string     `json:"redis_key,omitempty"`
	Filepath            string     `json:"filepath,omitempty"`
	APIStatusCode       int        `json:"api_status_code,omitempty"`
	RuntimeSessionID    string     `json:"runtime_session_id,omitempty"`
	BusinessPartnerID   *int       `json:"business_partner,omitempty"`
	ServiceLabel        string     `json:"service_label,omitempty"`
	APIType             string     `json:"api_type,omitempty"`
	Message             StorageBin `json:"message,omitempty"`
	APISchema           string     `json:"api_schema,omitempty"`
	Accepter            []string   `json:"accepter,omitempty"`
	Deleted             bool       `json:"deleted,omitempty"`
	SQLUpdateResult     *bool      `json:"sql_update_result,omitempty"`
	SQLUpdateError      string     `json:"sql_update_error,omitempty"`
	SubfuncResult       *bool      `json:"subfunc_result,omitempty"`
	SubfuncError        string     `json:"subfunc_error,omitempty"`
	ExconfResult        *bool      `json:"exconf_result,omitempty"`
	ExconfError         string     `json:"exconf_error,omitempty"`
	APIProcessingResult *bool      `json:"api_processing_result,omitempty"`
	APIProcessingError  string     `json:"api_processing_error,omitempty"`
}

type StorageBin struct {
	General *[]StorageBinGeneral `json:"General,omitempty"`
}

type StorageBinGeneral struct {
	BusinessPartner    int      `json:"BusinessPartner,omitempty"`
	Plant              string   `json:"Plant,omitempty"`
	StorageLocation    string   `json:"StorageLocation,omitempty"`
	StorageBin         string   `json:"StorageBin,omitempty"`
	StorageType        *string  `json:"StorageType,omitempty"`
	XCoordinates       *float32 `json:"XCoordinates,omitempty"`
	YCoordinates       *float32 `json:"YCoordinates,omitempty"`
	ZCoordinates       *float32 `json:"ZCoordinates,omitempty"`
	Capacity           *float32 `json:"Capacity,omitempty"`
	CapacityUnit       *string  `json:"CapacityUnit,omitempty"`
	CapacityWeight     *float32 `json:"CapacityWeight,omitempty"`
	CapacityWeightUnit *string  `json:"CapacityWeightUnit,omitempty"`
}

func CreateStorageBinRes(msg rabbitmq.RabbitmqMessage) (*StorageBinRes, error) {
	res := StorageBinRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
