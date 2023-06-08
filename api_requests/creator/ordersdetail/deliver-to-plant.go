package ordersdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateDeliverToPlantReq(param *dpfm_api_input_reader.OrdersDetailParams, oiRes *apiresponses.OrdersRes, scRes *apiresponses.SupplyChainRelationshipRes, sID string, log *logger.Logger) *models.PlantReq {
	generals := make(models.PlantGenerals, len(*scRes.Message.DeliveryRelation))
	for i, v := range *scRes.Message.DeliveryRelation {
		generals[i].BusinessPartner = &v.DeliverFromParty
		generals[i].Language = getLangPtr(param.Language)
	}

	return &models.PlantReq{
		General: models.PlantGeneral{
			BusinessPartner: *(*oiRes.Message.Item)[0].DeliverToParty,
			Plant:           *(*oiRes.Message.Item)[0].DeliverToPlant,
		},
		Accepter: []string{
			"General",
			"Generals",
		},
		RuntimeSessionID: sID,
	}
}
