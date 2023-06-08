package ordersdetaillist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductRequest(param *dpfm_api_input_reader.OrdersDetailListParams, oRes *apiresponses.OrdersRes, sID string, log *logger.Logger) *models.ProductMasterReq {
	return &models.ProductMasterReq{
		BusinessPartnerID: &param.BusinessPartner,
		General: models.PMGeneral{
			BusinessPartner: []models.PMBusinessPartner{
				{BusinessPartner: param.BusinessPartner,
					ProductDescription: []models.ProductDescription{
						{
							ProductDescByBP: []models.ProductDescByBP{
								{
									BusinessPartner: param.BusinessPartner,
									Language:        param.Language,
								},
							},
						},
					},
				},
			},
		},

		Accepter: []string{
			"ProductDescriptions",
		},
		RuntimeSessionID: sID,
	}
}
