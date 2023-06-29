package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type BillOfMaterialDetailList struct {
	UIKeyGeneralUserID          string                         `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string                         `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string                         `json:"ui_key_general_business_partner"`
	UIFunction                  string                         `json:"ui_function"`
	UIKeyFunctionURL            string                         `json:"ui_key_function_url"`
	RuntimeSessionID            string                         `json:"runtime_session_id"`
	Params                      BillOfMaterialDetailListParams `json:"Params"`
	ReqReceiveQueue             *string                        `json:"responseReceiveQueue"`
}

type BillOfMaterialDetailListParams struct {
	UserID              string  `json:"UserID"`
	User                string  `json:"User"`
	Language            *string `json:"Language"`
	BusinessPartner     int     `json:"BusinessPartner"`
	BillOfMaterial      int     `json:"BillOfMaterial"`
	ValidityStartDate   *string `json:"ValidityStartDate"`
	OwnerPlantName      *string `json:"OwnerPlantName"`
	ProductDescription  *string `json:"ProductDescription"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func ReadBillOfMaterialDetailList(msg rabbitmq.RabbitmqMessage) *BillOfMaterialDetailList {
	d := BillOfMaterialDetailList{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	return &d
}
