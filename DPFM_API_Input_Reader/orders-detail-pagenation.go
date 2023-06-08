package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type OrdersDetailPagination struct {
	Params           OrdersDetailPaginationParams `json:"Params"`
	RuntimeSessionID string                       `json:"runtime_session_id"`
	UIKeyFunctionURL string                       `json:"ui_key_function_url"`
	ReqReceiveQueue  *string                      `json:"responseReceiveQueue"`
}
type OrdersDetailPaginationParams struct {
	OrderID         int    `json:"OrderId"`
	BusinessPartner int    `json:"BusinessPartner"`
	UserType        string `json:"UserType"`
	Language        string `json:"Language"`
}

func ReadOrdersDetailPagination(msg rabbitmq.RabbitmqMessage) *OrdersDetailPagination {
	d := OrdersDetailPagination{}
	err := json.Unmarshal(msg.Raw(), &d)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	return &d
}
