package deliverydocumentdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductStockToReq(params *dpfm_api_input_reader.DeliveryDocumentDetailParams, ddRes *apiresponses.DeliveryDocumentRes, accountingRes *apiresponses.ProductMasterRes, sID string, log *logger.Logger) *models.ProductStockReq {
	return &models.ProductStockReq{
		ProductStock: models.ProductStock{
			Product:         params.Product,
			BusinessPartner: *(*ddRes.Message.Item)[0].DeliverToParty,
			Plant:           *(*ddRes.Message.Item)[0].DeliverToPlant,
			ProductStockAvailability: []models.ProductStockAvailability{
				{
					ProductStockAvailabilityDate: *(*accountingRes.Message.Accounting)[0].PriceLastChangeDate,
				},
			},
		},
		Accepter: []string{
			"ProductStock",
			"ProductStockAvailability",
		},
		RuntimeSessionID: sID,
	}
}
