package equipmentdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateBusinessPartnerReq(param *dpfm_api_input_reader.EquipmentDetailParams, sID string, log *logger.Logger) *models.BusinessPartnerReq {
	req := &models.BusinessPartnerReq{
		General: models.BPGeneral{
			BusinessPartner: param.BusinessPartner,
		},

		Accepter: []string{
			"General",
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
