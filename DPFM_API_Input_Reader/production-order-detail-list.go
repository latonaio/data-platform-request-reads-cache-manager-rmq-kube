package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type ProductionOrderDetailList struct {
	UIKeyGeneralUserID          string                          `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string                          `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string                          `json:"ui_key_general_business_partner"`
	UIFunction                  string                          `json:"ui_function"`
	UIKeyFunctionURL            string                          `json:"ui_key_function_url"`
	RuntimeSessionID            string                          `json:"runtime_session_id"`
	Params                      ProductionOrderDetailListParams `json:"Params"`
	ReqReceiveQueue             *string                         `json:"responseReceiveQueue"`
}

type ProductionOrderDetailListParams struct {
	ProductionOrder                     *int    `json:"ProductionOrder"`
	ProductionOrderIsReleased           *bool   `json:"ProductionOrderIsReleased"`
	OwnerProductionPlantBusinessPartner *int    `json:"OwnerProductionPlantBusinessPartner"`
	BusinessPartner                     *int    `json:"BusinessPartner"`
	UserID                              *string `json:"UserId"`
	User                                *string `json:"User"`
	Language                            *string `json:"Language"`
	IsCancelled                         *bool   `json:"IsCancelled"`
	IsMarkedForDeletion                 *bool   `json:"IsMarkedForDeletion"`
}

func ReadProductionOrderDetailList(msg rabbitmq.RabbitmqMessage) *ProductionOrderDetailList {
	d := ProductionOrderDetailList{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	// lang := d.UIKeyGeneralUserLanguage[len("production/order/list/language="):]
	// d.UIKeyGeneralUserID = d.UIKeyGeneralUserID[len("production/order/list/userID="):]
	// bp, err := strconv.Atoi(d.UIKeyGeneralBusinessPartner[len("orders/list/businessPartnerID="):])
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// 	return nil
	// }
	// d.Params.Language = lang
	// d.Params.BusinessPartner = bp
	return &d
}
