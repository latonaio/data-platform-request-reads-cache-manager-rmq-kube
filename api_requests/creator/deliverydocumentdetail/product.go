package deliverydocumentdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductRequest(param *dpfm_api_input_reader.DeliveryDocumentDetailParams, ddRes *apiresponses.DeliveryDocumentRes, sID string, log *logger.Logger) *models.ProductMasterReq {
	ddItem := (*ddRes.Message.Item)[0]

	return &models.ProductMasterReq{
		BusinessPartnerID: &param.BusinessPartner,
		General: models.PMGeneral{
			Product: param.Product,
			BusinessPartner: []models.PMBusinessPartner{
				{
					BusinessPartner: *ddItem.DeliverToParty,
					ProductDescription: []models.ProductDescription{
						{
							ProductDescByBP: []models.ProductDescByBP{
								{
									BusinessPartner: param.BusinessPartner,
									Language:        param.Language,
								},
							},
						},
					},
					Allergen: []models.Allergen{
						{
							Product:         param.Product,
							BusinessPartner: *ddItem.ProductionPlantBusinessPartner,
						},
					},
				},
			},
			BPPlant: []models.PMBPPlant{
				{
					BusinessPartner: *ddItem.DeliverToParty,
					Plant:           *ddItem.DeliverToPlant,
				},
			},
		},
		Accepter: []string{
			"General",
			"ProductDescByBP",
			"BusinessPartner",
			"BPPlant",
			"Allergens",
		},
		RuntimeSessionID: sID,
	}
}
