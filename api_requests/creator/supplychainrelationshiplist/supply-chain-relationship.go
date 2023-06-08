package supplychainrelationshiplist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateSupplyChainRelationshipRequest(param *dpfm_api_input_reader.SupplyChainRelationshipListParams, sID string, log *logger.Logger) *models.SupplyChainRelationshipReq {
	return &models.SupplyChainRelationshipReq{
		BusinessPartnerID: param.BusinessPartner,
		General: models.SCRGeneral{
			Buyer:               param.Buyer,
			Seller:              param.Seller,
			IsMarkedForDeletion: param.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Generals",
		},
		RuntimeSessionID: sID,
	}
}
