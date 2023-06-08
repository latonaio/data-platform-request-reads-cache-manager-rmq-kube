package productionversiondetaillist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductionVersionDetailReq(param *dpfm_api_input_reader.ProductionVersionDetailListParams, sID string, log *logger.Logger) *models.ProductionVersionReq {
	req := &models.ProductionVersionReq{
		General: models.ProductionVersionHeader{
			ProductionVersion: *param.ProductionVersion,
			Item: []models.ProductionVersionItem{
				{
					IsMarkedForDeletion: param.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Items",
		},
		RuntimeSessionID: sID,
	}
	return req
}
