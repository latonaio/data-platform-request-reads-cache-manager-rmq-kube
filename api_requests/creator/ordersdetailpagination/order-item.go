package ordersdetailpagination

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateOrdersItemsReq(param *dpfm_api_input_reader.OrdersDetailPaginationParams, sID string, log *logger.Logger) *models.OrdersReq {
	m := &models.OrdersReq{
		Header: &models.OrdersHeader{
			OrderID: param.OrderID,
			Item: []models.OrdersItem{
				{
					OrderID: param.OrderID,
					// ItemDeliveryStatus:            param.ItemDeliveryStatus,
					// ItemCompleteDeliveryIsDefined: param.ItemDeliveryBlockStatus,
					// ItemIsCancelled:               param.IsCancelled,
					// IsMarkedForDeletion:           param.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Items",
		},
		RuntimeSessionID: sID,
	}
	if strings.ToLower(param.UserType) == "buyer" {
		m.Header.Buyer = &param.BusinessPartner
	} else if strings.ToLower(param.UserType) == "seller" {
		m.Header.Seller = &param.BusinessPartner
	}
	return m
}
