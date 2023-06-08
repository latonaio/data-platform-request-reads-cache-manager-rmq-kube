package deliverydocumentlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreatePlantListReq(param *dpfm_api_input_reader.DeliveryDocumentListParams, sID string, log *logger.Logger) *models.PlantReq {
	generals := make(models.PlantGenerals, 1)
	generals[0].Language = &param.Language

	return &models.PlantReq{
		Generals: generals,
		Accepter: []string{
			"Generals",
		},
		RuntimeSessionID: sID,
	}
}
