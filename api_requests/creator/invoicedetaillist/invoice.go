package invoicedetaillist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateInvoiceReq(param *dpfm_api_input_reader.InvoiceDetailListParams, sID string, log *logger.Logger) *models.InvoiceReq {
	return &models.InvoiceReq{
		Header: models.InvoiceHeader{
			InvoiceDocument: param.InvoiceDocument,
		},
		Accepter: []string{
			"Items",
		},
		RuntimeSessionID: sID,
	}
}
