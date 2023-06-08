package ordersdetaillist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateQuantityUnitRequest(param *dpfm_api_input_reader.OrdersDetailListParams, sID string, log *logger.Logger) *models.QuantityUnitReq {
	return &models.QuantityUnitReq{
		BusinessPartnerID: &param.BusinessPartner,
		QuantityUnit: models.QuantityUnit{
			QuantityUnitText: []models.QuantityUnitText{
				{Language: param.Language},
			},
		},
		Accepter: []string{
			"QuantityUnitTexts",
		},
		RuntimeSessionID: sID,
	}
}
