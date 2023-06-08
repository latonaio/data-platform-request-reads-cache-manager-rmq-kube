package pricemasterlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreatePriceMasterReq(param *dpfm_api_input_reader.PriceMasterListParams, sID string, log *logger.Logger) *models.PriceMasterReq {
	priceMaster := models.PriceMaster{
		Buyer:  param.Buyer,
		Seller: param.Seller,
	}
	return &models.PriceMasterReq{
		PriceMaster: priceMaster,
		Accepter: []string{
			"PriceMasters",
		},
		RuntimeSessionID: sID,
	}
}

func CreateBusinessPartnerReq(param *dpfm_api_input_reader.PriceMasterListParams, scrRes *apiresponses.PriceMasterRes, sID string, log *logger.Logger) *models.BusinessPartnerReq {
	bpIDs := make([]int, 0)
	dupCheck := make(map[int]struct{})
	for _, v := range *scrRes.Message.PriceMaster {
		id := *v.Seller
		if _, ok := dupCheck[id]; ok {
			continue
		}
		dupCheck[id] = struct{}{}
		bpIDs = append(bpIDs, id)
	}
	for _, v := range *scrRes.Message.PriceMaster {
		id := *v.Buyer
		if _, ok := dupCheck[id]; ok {
			continue
		}
		dupCheck[id] = struct{}{}
		bpIDs = append(bpIDs, id)
	}

	return &models.BusinessPartnerReq{
		Generals: models.BPGenerals{
			BusinessPartners: bpIDs,
		},
		Accepter: []string{
			"Generals",
		},
		RuntimeSessionID: sID,
	}
}
