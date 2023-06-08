package ordersdetaillist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreatePaymentMethodRequest(param *dpfm_api_input_reader.OrdersDetailListParams, sID string, log *logger.Logger) *models.PaymentMethodReq {
	return &models.PaymentMethodReq{
		BusinessPartnerID: &param.BusinessPartner,
		PaymentMethod: models.PaymentMethod{
			PaymentMethodText: []models.PaymentMethodText{
				{Language: param.Language},
			},
		},
		Accepter: []string{
			"PaymentMethodTexts",
		},
		RuntimeSessionID: sID,
	}
}
