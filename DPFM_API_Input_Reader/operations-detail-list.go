package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type OperationsDetailList struct {
	UIKeyGeneralUserID          string                     `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string                     `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string                     `json:"ui_key_general_business_partner"`
	UIFunction                  string                     `json:"ui_function"`
	UIKeyFunctionURL            string                     `json:"ui_key_function_url"`
	RuntimeSessionID            string                     `json:"runtime_session_id"`
	Params                      OperationsDetailListParams `json:"Params"`
	ReqReceiveQueue             *string                    `json:"responseReceiveQueue"`
}

type OperationsDetailListParams struct {
	Operations          int     `json:"Operations"`
	BusinessPartner     int     `json:"BusinessPartner"`
	UserID              string  `json:"UserId"`
	User                string  `json:"User"`
	Language            *string `json:"Language"`
	IsCancelled         *bool   `json:"IsCancelled"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func ReadOperationsDetailList(msg rabbitmq.RabbitmqMessage) *OperationsDetailList {
	d := OperationsDetailList{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	return &d
}
