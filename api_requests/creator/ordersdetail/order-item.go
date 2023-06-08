package ordersdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateOrdersItemReq(param *dpfm_api_input_reader.OrdersDetailParams, sID string, log *logger.Logger) *models.OrdersReq {
	return &models.OrdersReq{
		Header: &models.OrdersHeader{
			OrderID:             param.OrderID,
			IsMarkedForDeletion: param.IsMarkedForDeletion,
			Item: []models.OrdersItem{
				{
					OrderID:   param.OrderID,
					OrderItem: param.OrderItem,
				},
			},
		},
		Accepter: []string{
			"ItemPricingElement",
			"Item",
			"ItemScheduleLines",
		},
		RuntimeSessionID: sID,
	}
}
