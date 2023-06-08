package deliverydocumentlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateBusinessPartnerReq(param *dpfm_api_input_reader.DeliveryDocumentListParams, ddRes *apiresponses.DeliveryDocumentRes, scRes *apiresponses.SupplyChainRelationshipRes, sID string, log *logger.Logger) *models.BusinessPartnerReq {
	bpIDs := make([]int, 0)
	dupCheck := make(map[int]struct{})
	for _, v := range *ddRes.Message.Header {
		id := *v.DeliverToParty
		if _, ok := dupCheck[id]; ok {
			continue
		}
		dupCheck[id] = struct{}{}
		bpIDs = append(bpIDs, id)
	}
	for _, v := range *ddRes.Message.Header {
		id := *v.DeliverFromParty
		if _, ok := dupCheck[id]; ok {
			continue
		}
		dupCheck[id] = struct{}{}
		bpIDs = append(bpIDs, id)
	}
	for _, v := range *scRes.Message.DeliveryPlantRelation {
		id := v.DeliverToParty
		if _, ok := dupCheck[id]; ok {
			continue
		}
		dupCheck[id] = struct{}{}
		bpIDs = append(bpIDs, id)
	}
	for _, v := range *scRes.Message.DeliveryPlantRelation {
		id := v.DeliverFromParty
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
		General: models.BPGeneral{
			Language: &param.Language,
		},
		Accepter: []string{
			"Generals",
		},
		RuntimeSessionID: sID,
	}
}
