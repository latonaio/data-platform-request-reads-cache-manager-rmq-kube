package ordersdetail

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateAccountingReq(param *dpfm_api_input_reader.OrdersDetailParams, productRes *apiresponses.ProductMasterRes, bpPlantRes *apiresponses.ProductMasterRes, accepter []string, sID string, log *logger.Logger) *models.ProductMasterReq {
	bp := make([]models.PMBusinessPartner, 0)
	bpPlant := make([]models.PMBPPlant, 0)

	for i, v := range *bpPlantRes.Message.BPPlant {
		// bp = append(bp, models.PMBusinessPartner{
		// 	BusinessPartner: v.BusinessPartner,
		// })
		bpPlant = append(bpPlant, models.PMBPPlant{
			BusinessPartner: v.BusinessPartner,
			Plant:           (*bpPlantRes.Message.BPPlant)[i].Plant,
		})
	}

	return &models.ProductMasterReq{
		General: models.PMGeneral{
			Product:         param.Product,
			BusinessPartner: bp,
			BPPlant:         bpPlant,
		},
		Accepter: []string{
			"Accounting",
		},
		RuntimeSessionID: sID,
	}
}
