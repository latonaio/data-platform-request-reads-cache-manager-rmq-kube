package ordersdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductRequest(param *dpfm_api_input_reader.OrdersDetailParams, oiRes *apiresponses.OrdersRes, sID string, log *logger.Logger) *models.ProductMasterReq {
	item := (*oiRes.Message.Item)[0]
	return &models.ProductMasterReq{
		BusinessPartnerID: &param.BusinessPartner,
		General: models.PMGeneral{
			Product: param.Product,
			BusinessPartner: []models.PMBusinessPartner{
				{BusinessPartner: param.BusinessPartner,
					ProductDescription: []models.ProductDescription{
						{
							ProductDescByBP: []models.ProductDescByBP{
								{
									Product:         param.Product,
									BusinessPartner: param.BusinessPartner,
									Language:        param.Language,
								},
							},
						},
					},
					Allergen: []models.Allergen{
						{
							Product:         param.Product,
							BusinessPartner: *item.ProductionPlantBusinessPartner,
							// Language:        param.Language,
						},
					},
				},
			},
		},

		Accepter: []string{
			"General",
			"ProductDescByBP",
			"BusinessPartner",
			"Allergens",
		},
		RuntimeSessionID: sID,
	}
}
