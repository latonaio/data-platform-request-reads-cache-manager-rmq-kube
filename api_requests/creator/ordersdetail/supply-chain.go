package ordersdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateSupplyChainReq(param *dpfm_api_input_reader.OrdersDetailParams, oiRes *apiresponses.OrdersRes, sID string, log *logger.Logger) *models.SupplyChainRelationshipReq {
	pricing := (*oiRes.Message.ItemPricingElement)[0]
	dr := make([]models.DeliveryRelation, 0)

	for _, v := range *oiRes.Message.Item {
		dr = append(dr, models.DeliveryRelation{
			SupplyChainRelationshipID: *v.SupplyChainRelationshipID,
			Buyer:                     pricing.Buyer,
			Seller:                    pricing.Seller,
			DeliverToParty:            *v.DeliverToParty,
			DeliverFromParty:          *v.DeliverFromParty,
		})
	}
	return &models.SupplyChainRelationshipReq{
		General: models.SCRGeneral{
			DeliveryRelation: dr,
		},
		Accepter: []string{
			"DeliveryRelations",
		},
		RuntimeSessionID: sID,
	}
}
