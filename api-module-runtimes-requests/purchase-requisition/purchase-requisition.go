package apiModuleRuntimesRequestsPurchaseRequisition

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type PurchaseRequisitionReq struct {
	Header   Header   `json:"PurchaseRequisition"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	PurchaseRequisition          int     `json:"PurchaseRequisition"`
	PurchaseRequisitionDate      *string `json:"PurchaseRequisitionDate"`
	Buyer                        *int    `json:"Buyer"`
	PurchaseRequisitionType      *string `json:"PurchaseRequisitionType"`
	PlannedOrder                 *int    `json:"PlannedOrder"`
	PlannedOrderItem             *int    `json:"PlannedOrderItem"`
	ProductionOrder              *int    `json:"ProductionOrder"`
	ProductionOrderItem          *int    `json:"ProductionOrderItem"`
	PrecedingOrderID             *int    `json:"PrecedingOrderID"`
	PrecedingOrderItem           *int    `json:"PrecedingOrderItem"`
	Project                      *int    `json:"Project"`
	WBSElement                   *int    `json:"WBSElement"`
	HeaderOrderStatus            *string `json:"HeaderOrderStatus"`
	HeaderCompleteOrderIsDefined *bool   `json:"HeaderCompleteOrderIsDefined"`
	CreationDate                 *string `json:"CreationDate"`
	CreationTime                 *string `json:"CreationTime"`
	LastChangeDate               *string `json:"LastChangeDate"`
	LastChangeTime               *string `json:"LastChangeTime"`
	IsReleased                   *bool   `json:"IsReleased"`
	IsCancelled                  *bool   `json:"IsCancelled"`
	IsMarkedForDeletion          *bool   `json:"IsMarkedForDeletion"`
	Item                         []Item  `json:"Item"`
}

type Item struct {
	PurchaseRequisition                      int      `json:"PurchaseRequisition"`
	PurchaseRequisitionItem                  int      `json:"PurchaseRequisitionItem"`
	PurchaseRequisitionItemCategory          *string  `json:"PurchaseRequisitionItemCategory"`
	SupplyChainRelationshipID                *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID  *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID *int     `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                    int      `json:"Buyer"`
	Seller                                   *int     `json:"Seller"`
	DeliverToParty                           *int     `json:"DeliverToParty"`
	DeliverFromParty                         *int     `json:"DeliverFromParty"`
	DeliverToPlant                           *string  `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation            *string  `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlant                         *string  `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation          *string  `json:"DeliverFromPlantStorageLocation"`
	Product                                  *string  `json:"Product"`
	ProductGroup                             *string  `json:"ProductGroup"`
	RequestedQuantityInBaseUnit              *float32 `json:"RequestedQuantityInBaseUnit"`
	RequestedQuantityInDeliveryUnit          *float32 `json:"RequestedQuantityInDeliveryUnit"`
	BaseUnit                                 *string  `json:"BaseUnit"`
	DeliveryUnit                             *string  `json:"DeliveryUnit"`
	PlannedOrder                             *int     `json:"PlannedOrder"`
	PlannedOrderItem                         *int     `json:"PlannedOrderItem"`
	ProductionOrder                          *int     `json:"ProductionOrder"`
	ProductionOrderItem                      *int     `json:"ProductionOrderItem"`
	PrecedingOrderID                         *int     `json:"PrecedingOrderID"`
	PrecedingOrderItem                       *int     `json:"PrecedingOrderItem"`
	FollowingOrderID                         *int     `json:"FollowingOrderID"`
	FollowingOrderItem                       *int     `json:"FollowingOrderItem"`
	Project                                  *int     `json:"Project"`
	WBSElement                               *int     `json:"WBSElement"`
	PurchaseRequisitionItemText              *string  `json:"PurchaseRequisitionItemText"`
	PurchaseRequisitionItemTextByBuyer       *string  `json:"PurchaseRequisitionItemTextByBuyer"`
	PurchaseRequisitionItemTextBySeller      *string  `json:"PurchaseRequisitionItemTextBySeller"`
	PurchaseRequisitionItemPrice             *float32 `json:"PurchaseRequisitionItemPrice"`
	PurchaseRequisitionItemPriceQuantity     *int     `json:"PurchaseRequisitionItemPriceQuantity"`
	ProductPlannedDeliveryDuration           *float32 `json:"ProductPlannedDeliveryDuration"`
	ProductPlannedDeliveryDurationUnit       *int     `json:"ProductPlannedDeliveryDurationUnit"`
	OrderedQuantityInBaseUnit                *float32 `json:"OrderedQuantityInBaseUnit"`
	OrderedQuantityInDeliveryUnit            *float32 `json:"OrderedQuantityInDeliveryUnit"`
	DeliveryDate                             *string  `json:"DeliveryDate"`
	ItemCompleteOrderIsDefined               *bool    `json:"ItemCompleteOrderIsDefined"`
	TransactionCurrency                      *string  `json:"TransactionCurrency"`
	MRPArea                                  *string  `json:"MRPArea"`
	MRPController                            *string  `json:"MRPController"`
	StockConfirmationBusinessPartner         *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                   *string  `json:"StockConfirmationPlant"`
	ProductionPlantBusinessPartner           *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          *string  `json:"ProductionPlant"`
	GLAccount                                *string  `json:"GLAccount"`
	ItemBlockStatus                          *bool    `json:"ItemBlockStatus"`
	CreationDate                             *string  `json:"CreationDate"`
	CreationTime                             *string  `json:"CreationTime"`
	LastChangeDate                           *string  `json:"LastChangeDate"`
	LastChangeTime                           *string  `json:"LastChangeTime"`
	IsReleased                               *bool    `json:"IsReleased"`
	IsCancelled                              *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}

func CreatePurchaseRequisitionRequestHeaderByBuyer(
	requestPram *apiInputReader.Request,
	purchaseRequisitionHeader *apiInputReader.PurchaseRequisitionHeader,
) PurchaseRequisitionReq {
	req := PurchaseRequisitionReq{
		Header: Header{
			Buyer:                        purchaseRequisitionHeader.Buyer,
			HeaderOrderStatus:            purchaseRequisitionHeader.HeaderOrderStatus,
			HeaderCompleteOrderIsDefined: purchaseRequisitionHeader.HeaderCompleteOrderIsDefined,
			IsCancelled:                  purchaseRequisitionHeader.IsCancelled,
			IsMarkedForDeletion:          purchaseRequisitionHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersByBuyer",
		},
	}
	return req
}

func CreatePurchaseRequisitionRequestItems(
	requestPram *apiInputReader.Request,
	purchaseRequisitionItems *apiInputReader.PurchaseRequisitionItems,
) PurchaseRequisitionReq {
	req := PurchaseRequisitionReq{
		Header: Header{
			PurchaseRequisition: purchaseRequisitionItems.PurchaseRequisition,
			Item: []Item{
				{
					ItemCompleteOrderIsDefined: purchaseRequisitionItems.ItemCompleteOrderIsDefined,
					IsCancelled:                purchaseRequisitionItems.IsCancelled,
					IsMarkedForDeletion:        purchaseRequisitionItems.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Items",
		},
	}
	return req
}

func PurchaseRequisitionReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.PurchaseRequisition,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_PURCHASE_REQUISITON_SRV"
	aPIType := "reads"

	var request PurchaseRequisitionReq

	if accepter == "HeadersByBuyer" {
		request = CreatePurchaseRequisitionRequestHeaderByBuyer(
			requestPram,
			&apiInputReader.PurchaseRequisitionHeader{
				Buyer:                        input.PurchaseRequisitionHeader.Buyer,
				HeaderOrderStatus:            input.PurchaseRequisitionHeader.HeaderOrderStatus,
				HeaderCompleteOrderIsDefined: input.PurchaseRequisitionHeader.HeaderCompleteOrderIsDefined,
				IsCancelled:                  input.PurchaseRequisitionHeader.IsCancelled,
				IsMarkedForDeletion:          input.PurchaseRequisitionHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Items" {
		request = CreatePurchaseRequisitionRequestItems(
			requestPram,
			&apiInputReader.PurchaseRequisitionItems{
				PurchaseRequisition:        input.PurchaseRequisitionItems.PurchaseRequisition,
				ItemCompleteOrderIsDefined: input.PurchaseRequisitionItems.ItemCompleteOrderIsDefined,
				IsCancelled:                input.PurchaseRequisitionItems.IsCancelled,
				IsMarkedForDeletion:        input.PurchaseRequisitionItems.IsMarkedForDeletion,
			},
		)
	}

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}
