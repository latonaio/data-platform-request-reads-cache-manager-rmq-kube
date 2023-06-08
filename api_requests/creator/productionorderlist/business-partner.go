package productionorderlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateBusinessPartnerReq(param *dpfm_api_input_reader.ProductionOrderListParams, sID string, log *logger.Logger) *models.BusinessPartnerReq {
	return &models.BusinessPartnerReq{
		General: models.BPGeneral{
			BusinessPartner: *param.OwnerProductionPlantBusinessPartner,
			Language:        param.Language,
		},
		Accepter: []string{
			"General",
		},
		RuntimeSessionID: sID,
	}
}
