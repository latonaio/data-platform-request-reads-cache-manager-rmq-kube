package orderslist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateItemPricingReq(param *dpfm_api_input_reader.OrdersListParams, sID string, log *logger.Logger) *models.OrdersReq {
	req := &models.OrdersReq{
		Header: &models.OrdersHeader{
			Buyer:                     param.Buyer,
			Seller:                    param.Seller,
			OrderID:                   param.OrderID,
			HeaderDeliveryBlockStatus: &param.HeaderDeliveryBlockStatus,
			// HeaderDeliveryStatus:            &param.HeaderDeliveryStatus,
			HeaderCompleteDeliveryIsDefined: &param.HeaderCompleteDeliveryIsDefined,
			IsCancelled:                     param.IsCancelled,
			IsMarkedForDeletion:             param.IsMarkedForDeletion,
		},
		Accepter:         []string{},
		RuntimeSessionID: sID,
	}
	// if param.User == "Buyer" {
	// 	req.Header.Buyer = &param.BusinessPartner
	// } else if param.User == "Seller" {
	// 	req.Header.Seller = &param.BusinessPartner
	// }

	return req
}
