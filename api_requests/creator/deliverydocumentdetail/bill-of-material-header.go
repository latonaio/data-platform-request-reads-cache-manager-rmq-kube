package deliverydocumentdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateBillOfMaterialHeaderReq(param *dpfm_api_input_reader.DeliveryDocumentDetailParams, ddRes *apiresponses.DeliveryDocumentRes, sID string, log *logger.Logger) *models.BillOfMaterialReq {
	item := (*ddRes.Message.Item)[0]
	return &models.BillOfMaterialReq{
		Header: models.BillOfMaterialHeader{
			Product:              item.Product,
			OwnerBusinessPartner: item.DeliverFromParty,
			OwnerPlant:           item.ProductionPlant,
		},
		Accepter: []string{
			"Header",
		},
		RuntimeSessionID: sID,
	}
}
