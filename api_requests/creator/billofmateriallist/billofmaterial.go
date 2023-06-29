package billofmateriallist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateBillOfMaterialReq(param *dpfm_api_input_reader.BillOfMaterialListParams, sID string, log *logger.Logger) *models.BillOfMaterialReq {
	req := &models.BillOfMaterialReq{
		Header: models.BillOfMaterialHeader{
			OwnerBusinessPartner: param.BusinessPartner,
			IsMarkedForDeletion:  param.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeaderByOwnerProductionPlantBP",
		},
		RuntimeSessionID: sID,
	}
	return req
}
