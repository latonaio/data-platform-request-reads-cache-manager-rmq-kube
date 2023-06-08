package invoicedetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateBusinessPartnerReq(param *dpfm_api_input_reader.InvoiceDetailParams, iRes *apiresponses.InvoiceRes, sID string, log *logger.Logger) *models.BusinessPartnerReq {
	bpIDs := make([]int, 0)
	dupCheck := make(map[int]struct{})
	for _, v := range *iRes.Message.Header {
		id := *v.BillFromParty
		if _, ok := dupCheck[id]; ok {
			continue
		}
		dupCheck[id] = struct{}{}
		bpIDs = append(bpIDs, id)
	}
	for _, v := range *iRes.Message.Header {
		id := *v.BillToParty
		if _, ok := dupCheck[id]; ok {
			continue
		}
		dupCheck[id] = struct{}{}
		bpIDs = append(bpIDs, id)
	}

	return &models.BusinessPartnerReq{
		Generals: models.BPGenerals{
			BusinessPartners: bpIDs,
		},
		Accepter: []string{
			"Generals",
		},
		RuntimeSessionID: sID,
	}
}
