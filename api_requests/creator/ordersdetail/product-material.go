package ordersdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateProductMaterialReq(param *dpfm_api_input_reader.OrdersDetailParams, bmRes *apiresponses.BillOfMaterialRes, sID string, log *logger.Logger) *models.ProductMasterReq {
	pdByBPs := make([]models.ProductDescByBP, 0, len(*bmRes.Message.Item))
	for _, v := range *bmRes.Message.Item {
		pdByBPs = append(pdByBPs, models.ProductDescByBP{
			Product:         *v.ComponentProduct,
			BusinessPartner: *v.StockConfirmationBusinessPartner,
			Language:        param.Language,
		})
	}

	return &models.ProductMasterReq{
		General: models.PMGeneral{
			BusinessPartner: []models.PMBusinessPartner{
				{
					// BusinessPartner: param.BusinessPartner,
					ProductDescription: []models.ProductDescription{
						{
							ProductDescByBP: pdByBPs,
						},
					},
				},
			},
		},
		Accepter:         []string{"ProductDescByBP"},
		RuntimeSessionID: sID,
	}
}
