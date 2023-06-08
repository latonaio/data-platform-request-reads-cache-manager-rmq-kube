package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type DeliveryDocumentDetail struct {
	UIKeyGeneralUserID          string                       `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string                       `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string                       `json:"ui_key_general_business_partner"`
	UIFunction                  string                       `json:"ui_function"`
	UIKeyFunctionURL            string                       `json:"ui_key_function_url"`
	RuntimeSessionID            string                       `json:"runtime_session_id"`
	Params                      DeliveryDocumentDetailParams `json:"Params"`
	ReqReceiveQueue             *string                      `json:"responseReceiveQueue"`
}
type DeliveryDocumentDetailParams struct {
	DeliveryDocument     int    `json:"DeliveryDocument"`
	DeliveryDocumentItem int    `json:"DeliveryDocumentItem"`
	DeliverToParty       *int   `json:"DeliverToParty"`
	DeliverFromParty     *int   `json:"DeliverFromParty"`
	User                 string `json:"User"`
	Product              string `json:"Product"`

	UserID          string `json:"UserId"`
	BusinessPartner int    `json:"BusinessPartner"`
	Language        string `json:"Language"`
}

func ReadDeliveryDocumentDetail(msg rabbitmq.RabbitmqMessage) *DeliveryDocumentDetail {
	d := DeliveryDocumentDetail{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}

	return &d
}
