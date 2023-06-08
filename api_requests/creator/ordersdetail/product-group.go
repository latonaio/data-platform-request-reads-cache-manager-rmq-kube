package ordersdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductGroupReq(param *dpfm_api_input_reader.OrdersDetailParams, productRes *apiresponses.ProductMasterRes, sID string, log *logger.Logger) *models.ProductGroupReq {
	return &models.ProductGroupReq{
		ProductGroup: []models.ProductGroup{{
			ProductGroup: *(*productRes.Message.General)[0].ProductGroup,
			ProductGroupText: models.ProductGroupText{
				Language:         param.Language,
				ProductGroupName: (*productRes.Message.General)[0].ProductGroup,
			},
		}},
		Accepter: []string{
			"ProductGroup",
			"ProductGroupText",
		},
		RuntimeSessionID: sID,
	}
}
