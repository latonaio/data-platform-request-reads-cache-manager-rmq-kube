package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type ProductDetailList struct {
	UIKeyGeneralUserID          string                  `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string                  `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string                  `json:"ui_key_general_business_partner"`
	UIFunction                  string                  `json:"ui_function"`
	UIKeyFunctionURL            string                  `json:"ui_key_function_url"`
	RuntimeSessionID            string                  `json:"runtime_session_id"`
	Params                      ProductDetailListParams `json:"Params"`
	ReqReceiveQueue             *string                 `json:"responseReceiveQueue"`
}

type ProductDetailListParams struct {
	Product             string  `json:"Product"`
	BusinessPartner     *int    `json:"BusinessPartner"`
	UserID              *string `json:"UserId"`
	User                *string `json:"User"`
	Language            *string `json:"Language"`
	IsCancelled         *bool   `json:"IsCancelled"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func ReadProductDetailList(msg rabbitmq.RabbitmqMessage) *ProductDetailList {
	d := ProductDetailList{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	return &d
}
