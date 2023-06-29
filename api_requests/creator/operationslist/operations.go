package operationslist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateOperationsRequest(param *dpfm_api_input_reader.OperationsListParams, sID string, log *logger.Logger) *models.OperationsReq {
	return &models.OperationsReq{
		BusinessPartnerID: param.BusinessPartner,
		Header: models.OperationsHeader{
			OwnerProductionPlantBusinessPartner: *param.BusinessPartner,
			IsMarkedForDeletion:                 param.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeaderByOwnerProductionPlantBP",
		},
		RuntimeSessionID: sID,
	}
}
