package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type DeliveryDocumentDetailPagination struct {
	Params           DeliveryDocumentDetailPaginationParams `json:"Params"`
	RuntimeSessionID string                                 `json:"runtime_session_id"`
	UIKeyFunctionURL string                                 `json:"ui_key_function_url"`
	ReqReceiveQueue  *string                                `json:"responseReceiveQueue"`
}
type DeliveryDocumentDetailPaginationParams struct {
	DeliveryDocument int    `json:"DeliveryDocument"`
	BusinessPartner  int    `json:"BusinessPartner"`
	UserType         string `json:"UserType"`
	Language         string `json:"Language"`
}

func ReadDeliveryDocumentDetailPagination(msg rabbitmq.RabbitmqMessage) *DeliveryDocumentDetailPagination {
	d := DeliveryDocumentDetailPagination{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	return &d
}
