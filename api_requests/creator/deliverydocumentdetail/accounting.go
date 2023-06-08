package deliverydocumentdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CrerateAccountingReq(param *dpfm_api_input_reader.DeliveryDocumentDetailParams, productRes *apiresponses.ProductMasterRes, bpPlantRes *apiresponses.ProductMasterRes, sID string, log *logger.Logger) *models.ProductMasterReq {
	return &models.ProductMasterReq{
		General: models.PMGeneral{
			Product: param.Product,
			// BusinessPartner: []models.PMBusinessPartner{
			// 	{BusinessPartner: (*productRes.Message.BusinessPartner)[0].BusinessPartner},
			// },
			BPPlant: []models.PMBPPlant{
				{
					BusinessPartner: (*productRes.Message.BusinessPartner)[0].BusinessPartner,
					Plant:           (*bpPlantRes.Message.BPPlant)[0].Plant},
			},
		},
		Accepter: []string{
			"Accounting",
		},
		RuntimeSessionID: sID,
	}
}
