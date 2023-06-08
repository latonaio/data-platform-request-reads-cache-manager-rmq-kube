package equipmentlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateEquipmentReq(param *dpfm_api_input_reader.EquipmentListParams, sID string, log *logger.Logger) *models.EquipmentReq {
	req := &models.EquipmentReq{
		General: models.EquipmentGeneral{
			BusinessPartner:     &param.BusinessPartner,
			IsMarkedForDeletion: param.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Generals",
		},
		RuntimeSessionID: sID,
	}
	return req
}
