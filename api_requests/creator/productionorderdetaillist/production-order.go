package productionorderdetaillist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductionOrderReq(param *dpfm_api_input_reader.ProductionOrderDetailListParams, sID string, log *logger.Logger) *models.ProductionOrderReq {
	req := &models.ProductionOrderReq{
		Header: &models.ProductionOrderHeader{
			ProductionOrder:                     *param.ProductionOrder,
			OwnerProductionPlantBusinessPartner: param.OwnerProductionPlantBusinessPartner,
		},
		Accepter: []string{
			"Items",
		},
		RuntimeSessionID: sID,
	}
	return req
}
