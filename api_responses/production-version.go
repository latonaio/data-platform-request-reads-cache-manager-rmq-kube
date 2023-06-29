package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type ProductionVersionRes struct {
	ConnectionKey       string                   `json:"connection_key"`
	Result              bool                     `json:"result"`
	RedisKey            string                   `json:"redis_key"`
	Filepath            string                   `json:"filepath"`
	APIStatusCode       int                      `json:"api_status_code"`
	RuntimeSessionID    string                   `json:"runtime_session_id"`
	BusinessPartnerID   *int                     `json:"business_partner"`
	ServiceLabel        string                   `json:"service_label"`
	APIType             string                   `json:"api_type"`
	Message             ProductionVersionMessage `json:"message"`
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

type ProductionVersionMessage struct {
	Header *[]ProductionVersionHeader `json:"Header"`
	Item   *[]ProductionVersionItem   `json:"Item"`
}

type ProductionVersionHeader struct {
	ProductionVersion       int     `json:"ProductionVersion"`
	Product                 string  `json:"Product"`
	OwnerBusinessPartner    int     `json:"OwnerBusinessPartner"`
	OwnerPlant              string  `json:"OwnerPlant"`
	BillOfMaterial          int     `json:"BillOfMaterial"`
	Operations              int     `json:"Operations"`
	ProductionVersionText   *string `json:"ProductionVersionText"`
	ProductionVersionStatus *string `json:"ProductionVersionStatus"`
	ValidityStartDate       *string `json:"ValidityStartDate"`
	ValidityEndDate         *string `json:"ValidityEndDate"`
	CreationDate            *string `json:"CreationDate"`
	LastChangeDate          *string `json:"LastChangeDate"`
	IsLocked                *bool   `json:"IsLocked"`
	IsMarkedForDeletion     *bool   `json:"IsMarkedForDeletion"`
}

type ProductionVersionItem struct {
	ProductionVersion       int     `json:"ProductionVersion"`
	ProductionVersionItem   int     `json:"ProductionVersionItem"`
	Product                 string  `json:"Product"`
	BusinessPartner         int     `json:"BusinessPartner"`
	Plant                   string  `json:"Plant"`
	PlantName               *string `json:"PlantName"`
	BillOfMaterial          int     `json:"BillOfMaterial"`
	Operations              int     `json:"Operations"`
	ProductionVersionText   *string `json:"ProductionVersionText"`
	ProductionVersionStatus *string `json:"ProductionVersionStatus"`
	ValidityStartDate       *string `json:"ValidityStartDate"`
	ValidityEndDate         *string `json:"ValidityEndDate"`
	CreationDate            *string `json:"CreationDate"`
	LastChangeDate          *string `json:"LastChangeDate"`
	IsLocked                *bool   `json:"IsLocked"`
	IsMarkedForDeletion     *bool   `json:"IsMarkedForDeletion"`
}

func CreateProductionVersionRes(msg rabbitmq.RabbitmqMessage) (*ProductionVersionRes, error) {
	res := ProductionVersionRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
