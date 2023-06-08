package billofmaterialdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateBillOfMaterialReq(param *dpfm_api_input_reader.BillOfMaterialDetailListParams, sID string, log *logger.Logger) *models.BillOfMaterialReq {
	req := &models.BillOfMaterialReq{
		Header: models.BillOfMaterialHeader{
			BillOfMaterial: param.BillOfMaterial,
			Item: []models.BillOfMaterialHeaderItem{
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
