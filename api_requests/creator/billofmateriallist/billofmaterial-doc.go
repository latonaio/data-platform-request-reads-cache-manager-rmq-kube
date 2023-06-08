package billofmateriallist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductMasterDocReq(param *dpfm_api_input_reader.BillOfMaterialListParams, pvRes *apiresponses.BillOfMaterialRes, sID string, log *logger.Logger) *models.ProductMasterDocReq {
	return &models.ProductMasterDocReq{
		Product: models.PMDProduct{
			HeaderDoc: models.ProductMasterDocHeaderDoc{
				DocType:                  "IMAGE",
				DocIssuerBusinessPartner: param.BusinessPartner,
			},
		},
		Accepter:         []string{},
		RuntimeSessionID: sID,
	}
}
