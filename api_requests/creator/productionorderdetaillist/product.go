package productionorderdetaillist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductRequest(param *dpfm_api_input_reader.ProductionOrderDetailListParams, poRes *apiresponses.ProductionOrderRes, sID string, log *logger.Logger) *models.ProductMasterReq {
	pDesc := make([]models.ProductDescByBP, 0)
	for _, v := range *poRes.Message.Item {
		pDesc = append(pDesc, models.ProductDescByBP{
			Product:         *v.Product,
			BusinessPartner: *param.OwnerProductionPlantBusinessPartner,
			Language:        *param.Language,
		})
	}

	return &models.ProductMasterReq{
		BusinessPartnerID: param.BusinessPartner,
		General: models.PMGeneral{
			BusinessPartner: []models.PMBusinessPartner{
				{
					BusinessPartner: *param.BusinessPartner,
					ProductDescription: []models.ProductDescription{
						{
							ProductDescByBP: pDesc,
						},
					},
				},
			},
		},
		Accepter: []string{
			"ProductDescByBP",
		},
		RuntimeSessionID: sID,
	}
}
