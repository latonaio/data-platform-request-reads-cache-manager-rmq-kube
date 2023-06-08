package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"
	"strconv"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type OrdersDetailList struct {
	UIKeyGeneralUserID          string                 `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string                 `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string                 `json:"ui_key_general_business_partner"`
	UIFunction                  string                 `json:"ui_function"`
	UIKeyFunctionURL            string                 `json:"ui_key_function_url"`
	RuntimeSessionID            string                 `json:"runtime_session_id"`
	Params                      OrdersDetailListParams `json:"Params"`
	ReqReceiveQueue             *string                `json:"responseReceiveQueue"`
}
type OrdersDetailListParams struct {
	User                          string `json:"User"`
	Buyer                         *int   `json:"Buyer"`
	Seller                        *int   `json:"Seller"`
	UserType                      string `json:"UserType"`
	OrderID                       int    `json:"OrderId"`
	ItemCompleteDeliveryIsDefined *bool  `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus            *bool  `json:"ItemDeliveryStatus"`
	ItemDeliveryBlockStatus       *bool  `json:"ItemDeliveryBlockStatus"`
	IsCancelled                   *bool  `json:"IsCancelled"`
	IsMarkedForDeletion           *bool  `json:"IsMarkedForDeletion"`
	Language                      string `json:"Language"`
	BusinessPartner               int    `json:"BusinessPartner"`
	UserID                        string `json:"UserId"`
}

func ReadOrdersDetailList(msg rabbitmq.RabbitmqMessage) *OrdersDetailList {
	d := OrdersDetailList{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	d.UIKeyGeneralUserID = d.UIKeyGeneralUserID[len("orders/detail/list/userID="):]
	lang := d.UIKeyGeneralUserLanguage[len("orders/detail/list/language="):]
	bp, err := strconv.Atoi(d.UIKeyGeneralBusinessPartner[len("orders/detail/list/businessPartnerID="):])
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	d.Params.Language = lang
	d.Params.BusinessPartner = bp
	return &d
}
