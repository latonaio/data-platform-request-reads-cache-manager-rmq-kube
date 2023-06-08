package productionorderdetailpagination

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductionOrderDetailPaginationReq(param *dpfm_api_input_reader.ProductionOrderDetailPagenationParams, sID string, log *logger.Logger) *models.ProductionOrderReq {
	req := &models.ProductionOrderReq{
		Header: &models.ProductionOrderHeader{
			ProductionOrder: *param.ProductionOrder,
			Item: []models.ProductionOrderItem{
				{
					ItemIsReleased:          nil,
					ItemIsMarkedForDeletion: nil,
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
