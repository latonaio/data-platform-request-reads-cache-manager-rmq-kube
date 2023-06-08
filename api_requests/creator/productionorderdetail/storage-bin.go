package productionorderdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CrerateStorageBinReq(param *dpfm_api_input_reader.ProductionOrderDetailParams, poRes *apiresponses.ProductionOrderRes, sID string, log *logger.Logger) *models.StorageBinReq {
	return &models.StorageBinReq{
		General: models.StorageBin{
			BusinessPartner: (*poRes.Message.Item)[0].ProductionPlantBusinessPartner,
			Plant:           (*poRes.Message.Item)[0].ProductionPlant,
			StorageLocation: *(*poRes.Message.Item)[0].ProductionPlantStorageLocation,
		},

		Accepter: []string{
			"General",
		},
		RuntimeSessionID: sID,
	}
}
