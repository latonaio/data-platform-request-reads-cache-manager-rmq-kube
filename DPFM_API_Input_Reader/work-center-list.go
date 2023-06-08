package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type WorkCenterList struct {
	UIKeyGeneralUserID          string               `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string               `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string               `json:"ui_key_general_business_partner"`
	UIFunction                  string               `json:"ui_function"`
	UIKeyFunctionURL            string               `json:"ui_key_function_url"`
	RuntimeSessionID            string               `json:"runtime_session_id"`
	Params                      WorkCenterListParams `json:"Params"`
	ReqReceiveQueue             *string              `json:"responseReceiveQueue"`
}
type WorkCenterListParams struct {
	User                string `json:"User"`
	Language            string `json:"Language"`
	BusinessPartner     int    `json:"BusinessPartner"`
	UserID              string `json:"UserId"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

func ReadWorkCenterList(msg rabbitmq.RabbitmqMessage) *WorkCenterList {
	d := WorkCenterList{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	return &d
}
