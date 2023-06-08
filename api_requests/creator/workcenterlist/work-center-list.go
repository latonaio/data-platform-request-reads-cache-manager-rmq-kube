package workcenterlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateWorkCenterListReq(param *dpfm_api_input_reader.WorkCenterListParams, sID string, log *logger.Logger) *models.WorkCenterReq {
	return &models.WorkCenterReq{
		General: models.WorkCenterGeneral{
			BusinessPartner:     &param.BusinessPartner,
			IsMarkedForDeletion: param.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Generals",
		},
		RuntimeSessionID: sID,
	}
}
