package deliverydocumentdetailpagination

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateDeliveryDocumentReq(param *dpfm_api_input_reader.DeliveryDocumentDetailPaginationParams, sID string, log *logger.Logger) *models.DeliveryDocumentReq {
	req := &models.DeliveryDocumentReq{
		Header: models.DeliveryDocumentHeader{
			DeliveryDocument: param.DeliveryDocument,
		},

		Accepter: []string{
			"Items",
		},
		RuntimeSessionID: sID,
	}

	if strings.ToLower(param.UserType) == "delivertoparty" {
		req.Header.DeliverToParty = &param.BusinessPartner
	} else if strings.ToLower(param.UserType) == "deliverfromparty" {
		req.Header.DeliverFromParty = &param.BusinessPartner
	}
	return req
}
