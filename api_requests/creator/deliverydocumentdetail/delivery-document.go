package deliverydocumentdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateDeliveryDocumentReq(param *dpfm_api_input_reader.DeliveryDocumentDetailParams, sID string, log *logger.Logger) *models.DeliveryDocumentReq {
	req := &models.DeliveryDocumentReq{
		Header: models.DeliveryDocumentHeader{
			DeliveryDocument: param.DeliveryDocument,
			DeliverToParty:   param.DeliverToParty,
			DeliverFromParty: param.DeliverFromParty,
			Item: []models.DeliveryDocumentItem{
				{
					DeliveryDocument:     param.DeliveryDocument,
					DeliveryDocumentItem: param.DeliveryDocumentItem,
				},
			},
		},

		Accepter: []string{
			"Item",
		},
		RuntimeSessionID: sID,
	}

	// if param.User == "Buyer" {
	// 	req.Header.Buyer = &param.BusinessPartner
	// } else if param.User == "Seller" {
	// 	req.Header.Seller = &param.BusinessPartner
	// }
	return req
}
