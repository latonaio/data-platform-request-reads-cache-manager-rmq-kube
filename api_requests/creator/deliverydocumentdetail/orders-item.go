package deliverydocumentdetail

import (
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CrerateOrdersReq(ddRes *apiresponses.DeliveryDocumentRes, sID string, log *logger.Logger) *models.OrdersReq {
	return &models.OrdersReq{
		Header: &models.OrdersHeader{
			OrderID: *(*ddRes.Message.Item)[0].OrderID,
			Item: []models.OrdersItem{
				{
					OrderID:   *(*ddRes.Message.Item)[0].OrderID,
					OrderItem: *(*ddRes.Message.Item)[0].OrderItem,
				},
			},
		},
		Accepter: []string{
			"ItemPricingElement",
			"Item",
		},
		RuntimeSessionID: sID,
	}
}
