package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type EquipmentDetail struct {
	UIKeyGeneralUserID          string                `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string                `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string                `json:"ui_key_general_business_partner"`
	UIFunction                  string                `json:"ui_function"`
	UIKeyFunctionURL            string                `json:"ui_key_function_url"`
	RuntimeSessionID            string                `json:"runtime_session_id"`
	Params                      EquipmentDetailParams `json:"Params"`
	ReqReceiveQueue             *string               `json:"responseReceiveQueue"`
}
type EquipmentDetailParams struct {
	Equipment         int    `json:"Equipment"`
	EquipmentName     string `json:"EquipmentName"`
	EquipmentTypeName string `json:"EquipmentTypeName"`
	PlantName         string `json:"PlantName"`
	ValidityStartDate string `Json:"ValidityStartDate"`

	Language        string  `json:"Language"`
	BusinessPartner int     `json:"BusinessPartner"`
	UserID          string  `json:"UserId"`
	User            *string `json:"User"`

	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

func ReadEquipmentDetail(msg rabbitmq.RabbitmqMessage) *EquipmentDetail {
	d := EquipmentDetail{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}

	return &d
}
