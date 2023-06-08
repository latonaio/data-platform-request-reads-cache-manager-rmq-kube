package ordersdetaillist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreatePaymentTermsRequest(param *dpfm_api_input_reader.OrdersDetailListParams, sID string, log *logger.Logger) *models.PaymentTermsReq {
	return &models.PaymentTermsReq{
		BusinessPartnerID: &param.BusinessPartner,
		PaymentTerms: models.PaymentTerms{
			PaymentTermsText: []models.PaymentTermsText{
				{Language: param.Language},
			},
		},
		Accepter: []string{
			"PaymentTermsTexts",
		},
		RuntimeSessionID: sID,
	}
}
