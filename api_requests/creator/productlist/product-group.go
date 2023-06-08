package productlist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductGroupReq(param *dpfm_api_input_reader.ProductListParams, pmRes *apiresponses.ProductMasterRes, sID string, log *logger.Logger) *models.ProductGroupReq {
	pg := make([]models.ProductGroup, 0)
	for _, v := range *pmRes.Message.General {
		pg = append(pg, models.ProductGroup{
			ProductGroup: *v.ProductGroup,
			ProductGroupText: models.ProductGroupText{
				ProductGroup: *v.ProductGroup,
				Language:     *param.Language,
			},
		})
	}

	return &models.ProductGroupReq{
		BusinessPartnerID: param.BusinessPartner,
		ProductGroup:      pg,
		Accepter: []string{
			"ProductGroupTexts",
		},
		RuntimeSessionID: sID,
	}
}
