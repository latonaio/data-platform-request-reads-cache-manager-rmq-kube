package ordersdetail

import (
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductStockToReq(oiRes *apiresponses.OrdersRes, accountingRes *apiresponses.ProductMasterRes, sID string, log *logger.Logger) *models.ProductStockReq {
	item := (*oiRes.Message.Item)[0]
	return &models.ProductStockReq{
		ProductStock: models.ProductStock{
			Product:         *item.Product,
			BusinessPartner: *item.DeliverToParty,
			Plant:           *item.DeliverToPlant,
		},
		Accepter: []string{
			"ProductStock",
		},
		RuntimeSessionID: sID,
	}
}
