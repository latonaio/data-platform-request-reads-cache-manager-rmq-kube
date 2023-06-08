package equipmentlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductMasterDocReq(param *dpfm_api_input_reader.EquipmentListParams, eqRes *apiresponses.EquipmentRes, sID string, log *logger.Logger) *models.EquipmentReq {
	docs := make([]models.EquipmentGeneralDoc, 0)
	for _, v := range *eqRes.Message.General {
		docs = append(docs, models.EquipmentGeneralDoc{
			Equipment:                v.Equipment,
			DocIssuerBusinessPartner: &param.BusinessPartner,
		})
	}

	return &models.EquipmentReq{
		General: models.EquipmentGeneral{
			GeneralDoc: docs,
		},
		Accepter: []string{
			"GeneralDocs",
		},
		RuntimeSessionID: sID,
	}
}
