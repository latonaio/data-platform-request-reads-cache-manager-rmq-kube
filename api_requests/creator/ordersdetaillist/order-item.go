package ordersdetaillist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateOrdersItemsReq(param *dpfm_api_input_reader.OrdersDetailListParams, sID string, log *logger.Logger) *models.OrdersReq {
	return &models.OrdersReq{
		Header: &models.OrdersHeader{
			OrderID:             param.OrderID,
			Buyer:               param.Buyer,
			Seller:              param.Seller,
			IsMarkedForDeletion: param.IsMarkedForDeletion,
			Item: []models.OrdersItem{
				{
					OrderID:                 param.OrderID,
					ItemDeliveryBlockStatus: param.ItemDeliveryBlockStatus,
					// ItemDeliveryStatus:            param.ItemDeliveryStatus,
					ItemCompleteDeliveryIsDefined: param.ItemDeliveryBlockStatus,
					IsCancelled:                   param.IsCancelled,
					IsMarkedForDeletion:           param.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Header",
			"Items",
			"ItemPricingElements",
		},
		RuntimeSessionID: sID,
	}
}
