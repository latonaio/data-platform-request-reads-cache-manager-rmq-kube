package equipmentlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateEquipmentTypeReq(param *dpfm_api_input_reader.EquipmentListParams, eqRes *apiresponses.EquipmentRes, sID string, log *logger.Logger) *models.EquipmentTypeReq {
	eqTypeTxt := make([]models.EquipmentType, 0)
	for _, v := range *eqRes.Message.General {
		eqTypeTxt = append(eqTypeTxt,
			models.EquipmentType{
				EquipmentType: *v.EquipmentType,
				EquipmentTypeText: []models.EquipmentTypeText{
					{
						Language: param.Language,
					},
				},
			},
		)
	}
	req := &models.EquipmentTypeReq{
		EquipmentType: eqTypeTxt,
		Accepter: []string{
			"EquipmentTypeText",
		},
		RuntimeSessionID: sID,
	}

	// if param.User == "Buyer" {
	// 	req.Header.Buyer = &param.BusinessPartner
	// } else if param.User == "Seller" {
	// 	req.Header.Seller = &param.BusinessPartner
	// }
	return req
}
