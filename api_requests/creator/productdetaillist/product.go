package productdetaillist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductRequest(param *dpfm_api_input_reader.ProductDetailListParams, sID string, log *logger.Logger) *models.ProductMasterReq {
	return &models.ProductMasterReq{
		BusinessPartnerID: param.BusinessPartner,
		General: models.PMGeneral{
			Product:             param.Product,
			IsMarkedForDeletion: param.IsMarkedForDeletion,
			BusinessPartner: []models.PMBusinessPartner{
				{
					BusinessPartner: *param.BusinessPartner,
					Allergen: []models.Allergen{
						{
							Product: param.Product,
						},
					},
					ProductDescription: []models.ProductDescription{
						{
							Language: *param.Language,
							ProductDescByBP: []models.ProductDescByBP{
								{
									Product:         param.Product,
									BusinessPartner: *param.BusinessPartner,
									Language:        *param.Language,
								},
							},
						},
					},
				},
			},
		},
		Accepter: []string{
			"General", "ProductDescription", "BusinessPartner",
			"ProductDescByBP", "BPPlants", "StorageLocations",
			"StorageBins", "MRPAreas", "WorkSchedulings",
			"Qualities", "Taxes", "Accountings",
		},
		RuntimeSessionID: sID,
	}
}
