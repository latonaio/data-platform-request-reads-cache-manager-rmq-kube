package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type InvoiceDetailList struct {
	UIKeyGeneralUserID          string                  `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string                  `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string                  `json:"ui_key_general_business_partner"`
	UIFunction                  string                  `json:"ui_function"`
	UIKeyFunctionURL            string                  `json:"ui_key_function_url"`
	RuntimeSessionID            string                  `json:"runtime_session_id"`
	Params                      InvoiceDetailListParams `json:"Params"`
	ReqReceiveQueue             *string                 `json:"responseReceiveQueue"`
}

type InvoiceDetailListParams struct {
	InvoiceDocument int    `json:"InvoiceDocument"`
	User            string `json:"User"`
	BillToParty     *int   `json:"BillToParty"`
	BillFromParty   *int   `json:"BillFromParty"`

	HeaderPaymentBlockStatus bool `json:"HeaderPaymentBlockStatus"`
	IsCancelled              bool `json:"IsCancelled"`
	IsMarkedForDeletion      bool `json:"IsMarkedForDeletion"`

	Language        string `json:"Language"`
	UserID          string `json:"UserId"`
	BusinessPartner int    `json:"BusinessPartner"`
}

func ReadInvoiceDetailList(msg rabbitmq.RabbitmqMessage) *InvoiceDetailList {
	d := InvoiceDetailList{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	// d.UIKeyGeneralUserID = d.UIKeyGeneralUserID[len("invoice/list/userID="):]
	// lang := d.UIKeyGeneralUserLanguage[len("invoice/list/language="):]
	// bp, err := strconv.Atoi(d.UIKeyGeneralBusinessPartner[len("invoice/list/businessPartnerID="):])
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// 	return nil
	// }
	// d.Params.Language = lang
	// d.Params.BusinessPartner = bp
	return &d
}
