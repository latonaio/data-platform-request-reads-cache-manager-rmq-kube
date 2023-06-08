package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type DeliveryDocumentList struct {
	UIKeyGeneralUserID          string                     `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string                     `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string                     `json:"ui_key_general_business_partner"`
	UIFunction                  string                     `json:"ui_function"`
	UIKeyFunctionURL            string                     `json:"ui_key_function_url"`
	RuntimeSessionID            string                     `json:"runtime_session_id"`
	Params                      DeliveryDocumentListParams `json:"Params"`
	ReqReceiveQueue             *string                    `json:"responseReceiveQueue"`
}
type DeliveryDocumentListParams struct {
	User                         string  `json:"User"`
	DeliverToParty               *int    `json:"DeliverToParty"`
	DeliverFromParty             *int    `json:"DeliverFromParty"`
	HeaderBillingStatusException *string `json:"HeaderBillingStatusException"`

	Language        string `json:"Language"`
	BusinessPartner int    `json:"BusinessPartner"`
	UserID          string `json:"UserId"`

	HeaderCompleteDeliveryIsDefined *bool   `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus            *string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus             *bool   `json:"HeaderBillingStatus"`
	HeaderDeliveryBlockStatus       *bool   `json:"HeaderDeliveryBlockStatus"`
	HeaderIssuingBlockStatus        *bool   `json:"HeaderIssuingBlockStatus"`
	HeaderReceivingBlockStatus      *bool   `json:"HeaderReceivingBlockStatus"`
	HeaderBillingBlockStatus        *bool   `json:"HeaderBillingBlockStatus"`
	IsCancelled                     *bool   `json:"IsCancelled"`
	IsMarkedForDeletion             *bool   `json:"IsMarkedForDeletion"`
}

func ReadDeliveryDocumentList(msg rabbitmq.RabbitmqMessage) *DeliveryDocumentList {
	d := DeliveryDocumentList{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	return &d
}
