package apiModuleRuntimesRequestsProductionOrder

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type ProductionOrderReq struct {
	Header   Header   `json:"ProductionOrder"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	ProductionOrder                                    int         `json:"ProductionOrder"`
	SupplyChainRelationshipID                          int         `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID           int         `json:"SupplyChainRelationshipProductionPlantID"`
	SupplyChainRelationshipDeliveryID                  int         `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID             int         `json:"SupplyChainRelationshipDeliveryPlantID"`
	ProductionOrderType                                string      `json:"ProductionOrderType"`
	Product                                            string      `json:"Product"`
	Buyer                                              int         `json:"Buyer"`
	Seller                                             int         `json:"Seller"`
	OwnerProductionPlantBusinessPartner                *int        `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlant                               string      `json:"OwnerProductionPlant"`
	OwnerProductionPlantStorageLocation                string      `json:"OwnerProductionPlantStorageLocation"`
	DepartureDeliverFromParty                          int         `json:"DepartureDeliverFromParty"`
	DepartureDeliverFromPlant                          string      `json:"DepartureDeliverFromPlant"`
	DepartureDeliverFromPlantStorageLocation           string      `json:"DepartureDeliverFromPlantStorageLocation"`
	DestinationDeliverToParty                          int         `json:"DestinationDeliverToParty"`
	DestinationDeliverToPlant                          string      `json:"DestinationDeliverToPlant"`
	DestinationDeliverToPlantStorageLocation           string      `json:"DestinationDeliverToPlantStorageLocation"`
	ProductBaseUnit                                    string      `json:"ProductBaseUnit"`
	MRPArea                                            *string     `json:"MRPArea"`
	MRPController                                      *string     `json:"MRPController"`
	ProductionVersion                                  *int        `json:"ProductionVersion"`
	BillOfMaterial                                     int         `json:"BillOfMaterial"`
	Operations                                         int         `json:"Operations"`
	ProductionOrderQuantityInBaseUnit                  float32     `json:"ProductionOrderQuantityInBaseUnit"`
	ProductionOrderQuantityInDepartureProductionUnit   float32     `json:"ProductionOrderQuantityInDepartureProductionUnit"`
	ProductionOrderQuantityInDestinationProductionUnit float32     `json:"ProductionOrderQuantityInDestinationProductionUnit"`
	ProductionOrderQuantityInDepartureDeliveryUnit     float32     `json:"ProductionOrderQuantityInDepartureDeliveryUnit"`
	ProductionOrderQuantityInDestinationDeliveryUnit   float32     `json:"ProductionOrderQuantityInDestinationDeliveryUnit"`
	ProductionOrderDepartureProductionUnit             string      `json:"ProductionOrderDepartureProductionUnit"`
	ProductionOrderDestinationProductionUnit           string      `json:"ProductionOrderDestinationProductionUnit"`
	ProductionOrderDepartureDeliveryUnit               string      `json:"ProductionOrderDepartureDeliveryUnit"`
	ProductionOrderDestinationDeliveryUnit             string      `json:"ProductionOrderDestinationDeliveryUnit"`
	ProductionOrderPlannedScrapQtyInBaseUnit           *float32    `json:"ProductionOrderPlannedScrapQtyInBaseUnit"`
	ProductionOrderPlannedStartDate                    string      `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime                    string      `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate                      string      `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime                      string      `json:"ProductionOrderPlannedEndTime"`
	ProductionOrderActualReleaseDate                   *string     `json:"ProductionOrderActualReleaseDate"`
	ProductionOrderActualReleaseTime                   *string     `json:"ProductionOrderActualReleaseTime"`
	ProductionOrderActualStartDate                     *string     `json:"ProductionOrderActualStartDate"`
	ProductionOrderActualStartTime                     *string     `json:"ProductionOrderActualStartTime"`
	ProductionOrderActualEndDate                       *string     `json:"ProductionOrderActualEndDate"`
	ProductionOrderActualEndTime                       *string     `json:"ProductionOrderActualEndTime"`
	PlannedOrder                                       *int        `json:"PlannedOrder"`
	OrderID                                            *int        `json:"OrderID"`
	OrderItem                                          *int        `json:"OrderItem"`
	ProductionOrderHeaderText                          *string     `json:"ProductionOrderHeaderText"`
	CreationDate                                       string      `json:"CreationDate"`
	CreationTime                                       string      `json:"CreationTime"`
	LastChangeDate                                     string      `json:"LastChangeDate"`
	LastChangeTime                                     string      `json:"LastChangeTime"`
	IsReleased                                         *bool       `json:"IsReleased"`
	IsPartiallyConfirmed                               *bool       `json:"IsPartiallyConfirmed"`
	IsConfirmed                                        *bool       `json:"IsConfirmed"`
	IsLocked                                           *bool       `json:"IsLocked"`
	IsCancelled                                        *bool       `json:"IsCancelled"`
	IsMarkedForDeletion                                *bool       `json:"IsMarkedForDeletion"`
	Item                                               []Item      `json:"Item"`
}

type Item struct {
	ProductionOrder                               int             `json:"ProductionOrder"`
	ProductionOrderItem                           int             `json:"ProductionOrderItem"`
	PrecedingProductionOrderItem                  *int            `json:"PrecedingProductionOrderItem"`
	FollowingProductionOrderItem                  *int            `json:"FollowingProductionOrderItem"`
	SupplyChainRelationshipID                     *int            `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID      *int            `json:"SupplyChainRelationshipProductionPlantID"`
	SupplyChainRelationshipDeliveryID             *int            `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int            `json:"SupplyChainRelationshipDeliveryPlantID"`
	ProductionOrderType                           *string         `json:"ProductionOrderType"`
	Product                                       *string         `json:"Product"`
	Buyer                                         *int            `json:"Buyer"`
	Seller                                        *int            `json:"Seller"`
	ProductionPlantBusinessPartner                *int            `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string         `json:"ProductionPlant"`
	ProductionPlantStorageLocation                *string         `json:"ProductionPlantStorageLocation"`
	DeliverToParty                                *int            `json:"DeliverToParty"`
	DeliverToPlant                                *string         `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation                 *string         `json:"DeliverToPlantStorageLocation"`
	DeliverFromParty                              *int            `json:"DeliverFromParty"`
	DeliverFromPlant                              *string         `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation               *string         `json:"DeliverFromPlantStorageLocation"`
	MRPArea                                       *string         `json:"MRPArea"`
	MRPController                                 *string         `json:"MRPController"`
	ProductionVersion                             *int            `json:"ProductionVersion"`
	ProductionVersionItem                         *int            `json:"ProductionVersionItem"`
	BillOfMaterial                                *int            `json:"BillOfMaterial"`
	Operations                                    *int            `json:"Operations"`
	ProductionOrderQuantityInBaseUnit             *float32        `json:"ProductionOrderQuantityInBaseUnit"`
	ProductionOrderQuantityInProductionUnit       *float32        `json:"ProductionOrderQuantityInProductionUnit"`
	ProductionOrderQuantityInDeliveryUnit         *float32        `json:"ProductionOrderQuantityInDeliveryUnit"`
	ProductionOrderPlannedScrapQtyInBaseUnit      *float32        `json:"ProductionOrderPlannedScrapQtyInBaseUnit"`
	ProductionOrderMinimumLotSizeQuantity         *float32        `json:"ProductionOrderMinimumLotSizeQuantity"`
	ProductionOrderStandardLotSizeQuantity        *float32        `json:"ProductionOrderStandardLotSizeQuantity"`
	ProductionOrderMaximumLotSizeQuantity         *float32        `json:"ProductionOrderMaximumLotSizeQuantity"`
	ProductionOrderLotSizeRoundingQuantity        *float32        `json:"ProductionOrderLotSizeRoundingQuantity"`
	ProductionOrderLotSizeIsFixed                 *bool           `json:"ProductionOrderLotSizeIsFixed"`
	ProductionOrderPlannedStartDate               *string         `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime               *string         `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate                 *string         `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime                 *string         `json:"ProductionOrderPlannedEndTime"`
	ProductionOrderActualReleaseDate              *string         `json:"ProductionOrderActualReleaseDate"`
	ProductionOrderActualReleaseTime              *string         `json:"ProductionOrderActualReleaseTime"`
	ProductionOrderActualStartDate                *string         `json:"ProductionOrderActualStartDate"`
	ProductionOrderActualStartTime                *string         `json:"ProductionOrderActualStartTime"`
	ProductionOrderActualEndDate                  *string         `json:"ProductionOrderActualEndDate"`
	ProductionOrderActualEndTime                  *string         `json:"ProductionOrderActualEndTime"`
	ConfirmedYieldQuantityInBaseUnit              *float32        `json:"ConfirmedYieldQuantityInBaseUnit"`
	ConfirmedYieldQuantityInProductionUnit        *float32        `json:"ConfirmedYieldQuantityInProductionUnit"`
	ScrappedQuantityInBaseUnit                    *float32        `json:"ScrappedQuantityInBaseUnit"`
	PlannedOrder                                  *int            `json:"PlannedOrder"`
	PlannedOrderItem                              *int            `json:"PlannedOrderItem"`
	OrderID                                       *int            `json:"OrderID"`
	OrderItem                                     *int            `json:"OrderItem"`
	ProductIsBatchManagedInProductionPlant        *bool           `json:"ProductIsBatchManagedInProductionPlant"`
	BatchMgmtPolicyInProductionOrder              *string         `json:"BatchMgmtPolicyInProductionOrder"`
	ProductionOrderTargetedBatch                  *string         `json:"ProductionOrderTargetedBatch"`
	ProductionOrderTargetedBatchValidityStartDate *string         `json:"ProductionOrderTargetedBatchValidityStartDate"`
	ProductionOrderTargetedBatchValidityStartTime *string         `json:"ProductionOrderTargetedBatchValidityStartTime"`
	ProductionOrderTargetedBatchValidityEndDate   *string         `json:"ProductionOrderTargetedBatchValidityEndDate"`
	ProductionOrderTargetedBatchValidityEndTime   *string         `json:"ProductionOrderTargetedBatchValidityEndTime"`
	ProductionOrderItemText                       *string         `json:"ProductionOrderItemText"`
	CreationDate                                  *string         `json:"CreationDate"`
	CreationTime                                  *string         `json:"CreationTime"`
	LastChangeDate                                *string         `json:"LastChangeDate"`
	LastChangeTime                                *string         `json:"LastChangeTime"`
	IsReleased                                    *bool           `json:"IsReleased"`
	IsPartiallyConfirmed                          *bool           `json:"IsPartiallyConfirmed"`
	IsConfirmed                                   *bool           `json:"IsConfirmed"`
	IsLocked                                      *bool           `json:"IsLocked"`
	IsCancelled                                   *bool           `json:"IsCancelled"`
	IsMarkedForDeletion                           *bool           `json:"IsMarkedForDeletion"`
}

func CreateProductionOrderRequestHeaderByOwnerProductionPlantBP(
	requestPram *apiInputReader.Request,
	productionOrderHeader *apiInputReader.ProductionOrderHeader,
) ProductionOrderReq {
	req := ProductionOrderReq{
		Header: Header{
			OwnerProductionPlantBusinessPartner: requestPram.BusinessPartner,
			IsMarkedForDeletion:                 productionOrderHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeaderByOwnerProductionPlantBP",
		},
	}
	return req
}

func CreateProductionOrderRequestItems(
	requestPram *apiInputReader.Request,
	productionOrderItems *apiInputReader.ProductionOrderItems,
) ProductionOrderReq {
	req := ProductionOrderReq{
		Header: Header{
			ProductionOrder: productionOrderItems.ProductionOrder,
			Item: []Item{
				{
					IsMarkedForDeletion: productionOrderItems.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Items",
		},
	}
	return req
}

func ProductionOrderReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductionOrder,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCTION_ORDER_SRV"
	aPIType := "reads"

	var request ProductionOrderReq

	if accepter == "HeaderByOwnerProductionPlantBP" {
		request = CreateProductionOrderRequestHeaderByOwnerProductionPlantBP(
			requestPram,
			&apiInputReader.ProductionOrderHeader{
				//IsMarkedForDeletion: input.ProductionOrderHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Items" {
		request = CreateProductionOrderRequestItems(
			requestPram,
			&apiInputReader.ProductionOrderItems{
				ProductionOrder: input.ProductionOrderItems.ProductionOrder,
				//IsMarkedForDeletion: input.ProductionOrderItems.IsMarkedForDeletion,
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
	)

	return responseBody
}
