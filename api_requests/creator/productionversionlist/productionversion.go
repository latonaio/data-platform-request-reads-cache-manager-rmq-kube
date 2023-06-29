package productionversionlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductionVersionReq(param *dpfm_api_input_reader.ProductionVersionListParams, sID string, log *logger.Logger) *models.ProductionVersionReq {
	req := &models.ProductionVersionReq{
		ProductionVersion: models.ProductionVersion{
			OwnerBusinessPartner: *param.BusinessPartner,
			IsMarkedForDeletion:  param.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Headers",
		},
		RuntimeSessionID: sID,
	}
	return req
}
