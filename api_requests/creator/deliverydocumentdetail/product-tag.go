package deliverydocumentdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductTagReq(param *dpfm_api_input_reader.DeliveryDocumentDetailParams, sID string, log *logger.Logger) *models.ProductTagReq {
	return &models.ProductTagReq{
		ProductTag: models.ProductTag{
			Product: param.Product,
		},
		Accepter: []string{
			"ProductTag",
		},
		RuntimeSessionID: sID,
	}
}
