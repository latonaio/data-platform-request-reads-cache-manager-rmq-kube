package pricemasterdetaillist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreatePriceMasterDetailReq(param *dpfm_api_input_reader.PriceMasterDetailListParams, sID string, log *logger.Logger) *models.PriceMasterDetailReq {
	priceMasterDetail := models.PriceMasterDetail{
		Buyer:                     param.Buyer,
		Seller:                    param.Seller,
		SupplyChainRelationshipID: param.SupplyChainRelationshipID,
	}
	return &models.PriceMasterDetailReq{
		PriceMasterDetail: priceMasterDetail,
		Accepter: []string{
			"PriceMasters",
		},
		RuntimeSessionID: sID,
	}
}
