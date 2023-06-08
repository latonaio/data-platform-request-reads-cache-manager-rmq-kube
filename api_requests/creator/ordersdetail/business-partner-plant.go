package ordersdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateBPPlantReq(param *dpfm_api_input_reader.OrdersDetailParams, oiRes *apiresponses.OrdersRes, sID string, log *logger.Logger) *models.ProductMasterReq {
	item := (*oiRes.Message.Item)[0]
	return &models.ProductMasterReq{
		General: models.PMGeneral{
			Product: param.Product,
			// BusinessPartner: []models.PMBusinessPartner{
			// 	{
			// 		BusinessPartner: *item.DeliverFromParty,
			// 		ProductDescription: []models.ProductDescription{
			// 			{
			// 				ProductDescByBP: []models.ProductDescByBP{
			// 					{
			// 						BusinessPartner: param.BusinessPartner,
			// 						Language:        param.Language,
			// 					},
			// 				},
			// 			},
			// 		},
			// 	},
			// 	{
			// 		BusinessPartner: *item.DeliverToParty,
			// 	},
			// },
			BPPlant: []models.PMBPPlant{
				{
					BusinessPartner: *item.DeliverFromParty,
					Plant:           *item.DeliverFromPlant,
				},
				{
					BusinessPartner: *item.DeliverToParty,
					Plant:           *item.DeliverToPlant,
				},
			},
		},

		Accepter: []string{
			"BPPlant",
		},
		RuntimeSessionID: sID,
	}
}
