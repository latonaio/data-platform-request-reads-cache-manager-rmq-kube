package supplychainrelationshipdetaillist

import (
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	models "data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/models"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"
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

func CreateSupplyChainRelationshipGeneralRequest(
	param *dpfm_api_input_reader.SupplyChainRelationshipExconfListParams,
	scrRes *apiresponses.SupplyChainRelationshipRes,
	accepter []string,
	sID string,
	log *logger.Logger,
) *models.SupplyChainRelationshipGeneralReq {
	return &models.SupplyChainRelationshipGeneralReq{
		SupplyChainRelationshipGeneral: &models.SupplyChainRelationshipGeneral{
			SupplyChainRelationshipID: param.SupplyChainRelationshipID,
			Buyer:                     param.Buyer,
			Seller:                    param.Seller,
		},
		Accepter:         accepter,
		RuntimeSessionID: sID,
	}
}

func CreateSupplyChainRelationshipDeliveryRelationRequest(
	param *dpfm_api_input_reader.SupplyChainRelationshipExconfListParams,
	scrRes *apiresponses.SupplyChainRelationshipRes,
	accepter []string,
	sID string,
	log *logger.Logger,
) *models.SupplyChainRelationshipDeliveryRelationReq {
	return &models.SupplyChainRelationshipDeliveryRelationReq{
		SupplyChainRelationshipDeliveryRelation: &models.SupplyChainRelationshipGeneral{
			SupplyChainRelationshipID: param.SupplyChainRelationshipID,
			Buyer:                     param.Buyer,
			Seller:                    param.Seller,
		},
		Accepter:         accepter,
		RuntimeSessionID: sID,
	}
}

func CreateSupplyChainRelationshipDeliveryPlantRequest(
	param *dpfm_api_input_reader.SupplyChainRelationshipExconfListParams,
	scrRes *apiresponses.SupplyChainRelationshipRes,
	accepter []string,
	sID string,
	log *logger.Logger,
) *models.SupplyChainRelationshipDeliveryPlantReq {
	return &models.SupplyChainRelationshipDeliveryPlantReq{
		SupplyChainRelationshipDeliveryPlantRelation: &models.SupplyChainRelationshipGeneral{
			SupplyChainRelationshipID: param.SupplyChainRelationshipID,
			Buyer:                     param.Buyer,
			Seller:                    param.Seller,
		},
		Accepter:         accepter,
		RuntimeSessionID: sID,
	}
}

func CreateSupplyChainRelationshipBillingRequest(
	param *dpfm_api_input_reader.SupplyChainRelationshipExconfListParams,
	scrRes *apiresponses.SupplyChainRelationshipRes,
	accepter []string,
	sID string,
	log *logger.Logger,
) *models.SupplyChainRelationshipBillingReq {
	return &models.SupplyChainRelationshipBillingReq{
		SupplyChainRelationshipBillingRelation: &models.SupplyChainRelationshipGeneral{
			SupplyChainRelationshipID: param.SupplyChainRelationshipID,
			Buyer:                     param.Buyer,
			Seller:                    param.Seller,
		},
		Accepter:         accepter,
		RuntimeSessionID: sID,
	}
}

func CreateSupplyChainRelationshipPaymentRequest(
	param *dpfm_api_input_reader.SupplyChainRelationshipExconfListParams,
	scrRes *apiresponses.SupplyChainRelationshipRes,
	accepter []string,
	sID string,
	log *logger.Logger,
) *models.SupplyChainRelationshipPaymentReq {
	return &models.SupplyChainRelationshipPaymentReq{
		SupplyChainRelationshipPaymentRelation: &models.SupplyChainRelationshipGeneral{
			SupplyChainRelationshipID: param.SupplyChainRelationshipID,
			Buyer:                     param.Buyer,
			Seller:                    param.Seller,
		},
		Accepter:         accepter,
		RuntimeSessionID: sID,
	}
}

func CreateSupplyChainRelationshipTransactionRequest(
	param *dpfm_api_input_reader.SupplyChainRelationshipExconfListParams,
	scrRes *apiresponses.SupplyChainRelationshipRes,
	accepter []string,
	sID string,
	log *logger.Logger,
) *models.SupplyChainRelationshipTransactionReq {
	return &models.SupplyChainRelationshipTransactionReq{
		SupplyChainRelationshipTransaction: &models.SupplyChainRelationshipGeneral{
			SupplyChainRelationshipID: param.SupplyChainRelationshipID,
			Buyer:                     param.Buyer,
			Seller:                    param.Seller,
		},
		Accepter:         accepter,
		RuntimeSessionID: sID,
	}
}

func CreateSupplyChainRelationshipListRequest(
	param *dpfm_api_input_reader.SupplyChainRelationshipExconfListParams,
	sID string,
	log *logger.Logger,
) *models.SupplyChainRelationshipReq {
	return &models.SupplyChainRelationshipReq{
		BusinessPartnerID: param.BusinessPartner,
		General: models.SCRGeneral{
			SupplyChainRelationshipID: *param.SupplyChainRelationshipID,
			Buyer:                     param.Buyer,
			Seller:                    param.Seller,
		},
		Accepter: []string{
			"General",
		},
		RuntimeSessionID: sID,
	}
}
