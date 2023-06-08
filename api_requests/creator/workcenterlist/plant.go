package workcenterlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreatePlantReq(param *dpfm_api_input_reader.WorkCenterListParams, wcRes *apiresponses.WorkCenterRes, sID string, log *logger.Logger) *models.PlantReq {
	generals := make(models.PlantGenerals, len(*wcRes.Message.General))
	for i, v := range *wcRes.Message.General {
		generals[i].Plant = &v.Plant
		generals[i].Language = &param.Language
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
