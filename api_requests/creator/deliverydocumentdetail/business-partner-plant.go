package deliverydocumentdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateBPPlantReq(param *dpfm_api_input_reader.DeliveryDocumentDetailParams, productRes *apiresponses.ProductMasterRes, sID string, log *logger.Logger) *models.ProductMasterReq {
	return &models.ProductMasterReq{
		General: models.PMGeneral{
			Product: param.Product,
			BusinessPartner: []models.PMBusinessPartner{
				{BusinessPartner: (*productRes.Message.BusinessPartner)[0].BusinessPartner},
			},
		},
		Accepter: []string{
			"BPPlant",
		},
		RuntimeSessionID: sID,
	}
}
