package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type SupplyChainRelationshipDetail struct {
	UIKeyGeneralUserID          string                              `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string                              `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string                              `json:"ui_key_general_business_partner"`
	UIFunction                  string                              `json:"ui_function"`
	UIKeyFunctionURL            string                              `json:"ui_key_function_url"`
	RuntimeSessionID            string                              `json:"runtime_session_id"`
	Params                      SupplyChainRelationshipDetailParams `json:"Params"`
	ReqReceiveQueue             *string                             `json:"responseReceiveQueue"`
}

type SupplyChainRelationshipDetailParams struct {
	BusinessPartner           *int    `json:"BusinessPartner"`
	UserID                    *string `json:"UserId"`
	User                      *string `json:"User"`
	Language                  *string `json:"Language"`
	SupplyChainRelationshipID *int    `json:"SupplyChainRelationshipID"`
	Buyer                     *int    `json:"Buyer"`
	Seller                    *int    `json:"Seller"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

func ReadSupplyChainRelationshipDetail(msg rabbitmq.RabbitmqMessage) *SupplyChainRelationshipDetail {
	d := SupplyChainRelationshipDetail{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	return &d
}
