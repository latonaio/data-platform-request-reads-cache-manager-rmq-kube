package supplychainrelationshipdetaillist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func CreateSupplyChainRelationshipRequest(param *dpfm_api_input_reader.SupplyChainRelationshipDetailListParams, sID string, log *logger.Logger) *models.SupplyChainRelationshipReq {
	return &models.SupplyChainRelationshipReq{
		General: models.SCRGeneral{
			IsMarkedForDeletion: param.IsMarkedForDeletion,
			Buyer:               param.Buyer,
			Seller:              param.Seller,
			DeliveryRelation: []models.DeliveryRelation{
				{
					SupplyChainRelationshipID: *param.SupplyChainRelationshipID,
					DeliverToParty:            *param.DeliverToParty,
					DeliverFromParty:          *param.DeliverFromParty,
					DeliveryPlantRelation: []models.DeliveryPlantRelation{
						{
							SupplyChainRelationshipDeliveryPlantID: *param.SupplyChainRelationshipDeliveryPlantID,
							DeliverToPlant:                         *param.DeliverToPlant,
							DeliverFromPlant:                       *param.DeliverFromPlant,
						},
					},
				},
			},
			BillingRelation: []models.BillingRelation{
				{
					SupplyChainRelationshipBillingID: *param.SupplyChainRelationshipBillingID,
					BillToParty:                      *param.BillToParty,
					BillFromParty:                    *param.BillFromParty,
					PaymentRelation: []models.PaymentRelation{
						{
							SupplyChainRelationshipPaymentID: *param.SupplyChainRelationshipPaymentID,
							Payer:                            *param.Payer,
							Payee:                            *param.Payee,
						},
					},
				},
			},
		},
		Accepter: []string{
			"Generals", "DeliveryRelations", "DeliveryPlantRelations",
			"BillingRelation", "PaymentRelation", "Transaction",
		},
		RuntimeSessionID: sID,
	}
}

func CreateSupplyChainRelationshipGeneralRequest(param *dpfm_api_input_reader.SupplyChainRelationshipExconfListParams, sID string, log *logger.Logger) *models.SupplyChainRelationshipGeneralReq {
	return &models.SupplyChainRelationshipGeneralReq{
		SupplyChainRelationshipGeneral: models.SupplyChainRelationshipGeneral{
			SupplyChainRelationshipID: param.SupplyChainRelationshipID,
			Buyer:                     param.Buyer,
			Seller:                    param.Seller,
		},
		RuntimeSessionID: sID,
	}
}
