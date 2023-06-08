package invoicelist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateInvoiceReq(param *dpfm_api_input_reader.InvoiceListParams, sID string, log *logger.Logger) *models.InvoiceReq {
	return &models.InvoiceReq{
		Header: models.InvoiceHeader{
			BillToParty:   param.BillToParty,
			BillFromParty: param.BillFromParty,
		},
		Accepter: []string{
			"Headers",
		},
		RuntimeSessionID: sID,
	}
}
