package ordersdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateBusinessPartnerReq(param *dpfm_api_input_reader.OrdersDetailParams, productRes *apiresponses.ProductMasterRes, oiRes *apiresponses.OrdersRes, sID string, log *logger.Logger) *models.BusinessPartnerReq {
	return &models.BusinessPartnerReq{
		General: models.BPGeneral{
			BusinessPartner: *(*oiRes.Message.Item)[0].ProductionPlantBusinessPartner,
			Language:        &param.Language,
		},
		Accepter: []string{
			"General",
		},
		RuntimeSessionID: sID,
	}
}
