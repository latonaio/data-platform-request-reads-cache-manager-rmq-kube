package businesspartnerlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateBusinessPartnertReq(param *dpfm_api_input_reader.BusinessPartnerListParams, sID string, log *logger.Logger) *models.BusinessPartnerReq {
	return &models.BusinessPartnerReq{
		Accepter: []string{
			"Generals-all",
		},
		RuntimeSessionID: sID,
	}
}
