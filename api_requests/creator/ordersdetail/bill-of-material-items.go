package ordersdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateBillOfMaterialItemsReq(param *dpfm_api_input_reader.OrdersDetailParams, bmRes *apiresponses.BillOfMaterialRes, sID string, log *logger.Logger) *models.BillOfMaterialReq {
	header := (*bmRes.Message.Header)[0]
	return &models.BillOfMaterialReq{
		Header: models.BillOfMaterialHeader{
			BillOfMaterial: header.BillOfMaterial,
		},
		Accepter: []string{
			"Items",
		},
		RuntimeSessionID: sID,
	}
}
