package pricemasterdetaillist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateDescriptionReq(param *dpfm_api_input_reader.PriceMasterDetailListParams, pmdRes *apiresponses.PriceMasterDetailRes, sID string, log *logger.Logger) *models.ProductMasterReq {
	descByBP := make([]models.ProductDescByBP, 0)
	for _, v := range pmdRes.Message.PriceMasterDetail {
		descByBP = append(descByBP, models.ProductDescByBP{
			Product:         *v.Product,
			BusinessPartner: v.Buyer,
			Language:        param.Language,
		})
	}
	return &models.ProductMasterReq{
		BusinessPartnerID: &param.BusinessPartner,
		General: models.PMGeneral{
			IsMarkedForDeletion: &param.IsMarkedForDeletion,
			BusinessPartner: []models.PMBusinessPartner{
				{
					BusinessPartner: param.BusinessPartner,
					ProductDescription: []models.ProductDescription{
						{
							ProductDescByBP: descByBP,
						},
					},
				},
			},
		},
		Accepter: []string{
			"ProductDescByBP",
		},
		RuntimeSessionID: sID,
	}
}
