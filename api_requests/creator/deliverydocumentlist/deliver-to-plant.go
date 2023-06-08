package deliverydocumentlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateDeliverToPlantReq(param *dpfm_api_input_reader.DeliveryDocumentListParams, ddRes *apiresponses.DeliveryDocumentRes, scRes *apiresponses.SupplyChainRelationshipRes, sID string, log *logger.Logger) *models.PlantReq {
	generals := make(models.PlantGenerals, len(*ddRes.Message.Header)+len(*scRes.Message.DeliveryPlantRelation))
	for i, v := range *ddRes.Message.Header {
		generals[i].BusinessPartner = v.DeliverToParty
		generals[i].Plant = v.DeliverToPlant
	}
	for i, v := range *scRes.Message.DeliveryPlantRelation {
		generals[len(*ddRes.Message.Header)+i].BusinessPartner = &v.DeliverToParty
		generals[len(*ddRes.Message.Header)+i].Plant = &v.DeliverFromPlant
	}
	return &models.PlantReq{
		Generals: generals,
		Accepter: []string{
			"Generals",
		},
		RuntimeSessionID: sID,
	}
}
