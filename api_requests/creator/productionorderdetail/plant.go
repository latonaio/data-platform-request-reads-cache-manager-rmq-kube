package productionorderdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateDeliverToPlantReq(param *dpfm_api_input_reader.ProductionOrderDetailParams, poRes *apiresponses.ProductionOrderRes, sID string, log *logger.Logger) *models.PlantReq {
	header := (*poRes.Message.Header)[0]
	return &models.PlantReq{
		General: models.PlantGeneral{
			BusinessPartner: header.OwnerProductionPlantBusinessPartner,
			Plant:           header.OwnerProductionPlant,
			Language:        param.Language,
			StorageLocation: models.StorageLocation{StorageLocation: *header.OwnerProductionPlantStorageLocation},
		},
		Accepter: []string{
			"General", "StorageLocation",
		},
		RuntimeSessionID: sID,
	}
}
