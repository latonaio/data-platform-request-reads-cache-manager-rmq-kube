package supplychainrelationshipdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateSupplyChainRelationshipRequest(
	param *dpfm_api_input_reader.SupplyChainRelationshipDetailParams,
	sID string,
	accepter []string,
	log *logger.Logger,
) *models.SupplyChainRelationshipReq {
	return &models.SupplyChainRelationshipReq{
		General: models.SCRGeneral{
			SupplyChainRelationshipID: *param.SupplyChainRelationshipID,
			IsMarkedForDeletion:       param.IsMarkedForDeletion,
			Buyer:                     param.Buyer,
			Seller:                    param.Seller,
		},
		Accepter:         accepter,
		RuntimeSessionID: sID,
	}
}
