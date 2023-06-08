package productionorderlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateDeliverToPlantReq(param *dpfm_api_input_reader.ProductionOrderListParams, poRes *apiresponses.ProductionOrderRes, sID string, log *logger.Logger) *models.PlantReq {
	generals := make(models.PlantGenerals, len(*poRes.Message.Header))
	for i, v := range *poRes.Message.Header {
		generals[i].BusinessPartner = &v.OwnerProductionPlantBusinessPartner
		generals[i].Plant = &v.OwnerProductionPlant
		generals[i].Language = param.Language
	}
	return &models.PlantReq{
		Generals: generals,
		Accepter: []string{
			"Generals",
		},
		RuntimeSessionID: sID,
	}
}
