package deliverydocumentlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateSupplyChainReq(param *dpfm_api_input_reader.DeliveryDocumentListParams, ddRes *apiresponses.DeliveryDocumentRes, sID string, log *logger.Logger) *models.SupplyChainRelationshipReq {
	dp := make([]models.DeliveryPlantRelation, 0)
	for _, v := range *ddRes.Message.Header {
		dp = append(dp, models.DeliveryPlantRelation{
			SupplyChainRelationshipID: *v.SupplyChainRelationshipID,
			Buyer:                     *v.Buyer,
			Seller:                    *v.Seller,
			DeliverToParty:            *v.DeliverToParty,
			DeliverFromParty:          *v.DeliverFromParty,
			DeliverToPlant:            "",
			DeliverFromPlant:          "",
		})
	}
	return &models.SupplyChainRelationshipReq{
		General: models.SCRGeneral{
			DeliveryRelation: []models.DeliveryRelation{
				{
					DeliveryPlantRelation: dp,
				},
			},
		},
		Accepter: []string{
			"DeliveryPlantRelations",
		},
		RuntimeSessionID: sID,
	}
}
