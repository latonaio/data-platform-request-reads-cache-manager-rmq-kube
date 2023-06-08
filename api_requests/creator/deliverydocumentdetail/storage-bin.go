package deliverydocumentdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CrerateStorageBinReq(param *dpfm_api_input_reader.DeliveryDocumentDetailParams, ddRes *apiresponses.DeliveryDocumentRes, bpPlantRes *apiresponses.ProductMasterRes, sID string, log *logger.Logger) *models.StorageBinReq {
	return &models.StorageBinReq{
		General: models.StorageBin{
			BusinessPartner: *(*ddRes.Message.Item)[0].DeliverToParty,
			Plant:           (*bpPlantRes.Message.BPPlant)[0].Plant,
			StorageLocation: *(*ddRes.Message.Item)[0].DeliverToPlantStorageLocation,
		},

		Accepter: []string{
			"General",
		},
		RuntimeSessionID: sID,
	}
}
