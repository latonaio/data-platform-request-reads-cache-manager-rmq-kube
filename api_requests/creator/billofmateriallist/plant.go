package billofmateriallist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreatePlantReq(param *dpfm_api_input_reader.BillOfMaterialListParams, bomRes *apiresponses.BillOfMaterialRes, sID string, log *logger.Logger) *models.PlantReq {
	generals := make(models.PlantGenerals, len(*bomRes.Message.Header))
	for i, v := range *bomRes.Message.Header {
		generals[i].Plant = v.OwnerPlant
		generals[i].Language = param.Language
	}

	req := &models.PlantReq{
		Generals: generals,
		Accepter: []string{
			"Generals",
		},
		RuntimeSessionID: sID,
	}

	return req
}
