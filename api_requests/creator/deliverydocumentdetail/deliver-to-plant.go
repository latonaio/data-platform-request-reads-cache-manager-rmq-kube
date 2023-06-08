package deliverydocumentdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateDeliverToPlantReq(param *dpfm_api_input_reader.DeliveryDocumentDetailParams, ddRes *apiresponses.DeliveryDocumentRes, sID string, log *logger.Logger) *models.PlantReq {
	return &models.PlantReq{
		General: models.PlantGeneral{
			BusinessPartner: *(*ddRes.Message.Item)[0].DeliverToParty,
			Plant:           *(*ddRes.Message.Item)[0].DeliverToPlant,
			StorageLocation: models.StorageLocation{
				StorageLocation: *(*ddRes.Message.Item)[0].DeliverToPlantStorageLocation,
			},
		},
		Accepter: []string{
			"General",
			"StorageLocation",
		},
		RuntimeSessionID: sID,
	}
}
