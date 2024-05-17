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
	ProductionOrderDate                                *string	   `json:"ProductionOrderDate"`
	SupplyChainRelationshipID                          *int        `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID           *int        `json:"SupplyChainRelationshipProductionPlantID"`
	SupplyChainRelationshipDeliveryID                  *int        `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID             *int        `json:"SupplyChainRelationshipDeliveryPlantID"`
	ProductionOrderType                                *string     `json:"ProductionOrderType"`
	Product                                            *string     `json:"Product"`
	Buyer                                              *int        `json:"Buyer"`
	Seller                                             *int        `json:"Seller"`
	OwnerProductionPlantBusinessPartner                *int        `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlant                               *string     `json:"OwnerProductionPlant"`
	OwnerProductionPlantStorageLocation                *string     `json:"OwnerProductionPlantStorageLocation"`
	DepartureDeliverFromParty                          *int        `json:"DepartureDeliverFromParty"`
	DepartureDeliverFromPlant                          *string     `json:"DepartureDeliverFromPlant"`
	DepartureDeliverFromPlantStorageLocation           *string     `json:"DepartureDeliverFromPlantStorageLocation"`
	DestinationDeliverToParty                          *int        `json:"DestinationDeliverToParty"`
	DestinationDeliverToPlant                          *string     `json:"DestinationDeliverToPlant"`
	DestinationDeliverToPlantStorageLocation           *string     `json:"DestinationDeliverToPlantStorageLocation"`
	ProductBaseUnit                                    *string     `json:"ProductBaseUnit"`
	MRPArea                                            *string     `json:"MRPArea"`
	MRPController                                      *string     `json:"MRPController"`
	ProductionVersion                                  *int        `json:"ProductionVersion"`
	BillOfMaterial                                     *int        `json:"BillOfMaterial"`
	Operations                                         *int        `json:"Operations"`
	ProductionOrderQuantityInBaseUnit                  *float32    `json:"ProductionOrderQuantityInBaseUnit"`
	ProductionOrderQuantityInDepartureProductionUnit   *float32    `json:"ProductionOrderQuantityInDepartureProductionUnit"`
	ProductionOrderQuantityInDestinationProductionUnit *float32    `json:"ProductionOrderQuantityInDestinationProductionUnit"`
	ProductionOrderQuantityInDepartureDeliveryUnit     *float32    `json:"ProductionOrderQuantityInDepartureDeliveryUnit"`
	ProductionOrderQuantityInDestinationDeliveryUnit   *float32    `json:"ProductionOrderQuantityInDestinationDeliveryUnit"`
	ProductionOrderDepartureProductionUnit             *string     `json:"ProductionOrderDepartureProductionUnit"`
	ProductionOrderDestinationProductionUnit           *string     `json:"ProductionOrderDestinationProductionUnit"`
	ProductionOrderDepartureDeliveryUnit               *string     `json:"ProductionOrderDepartureDeliveryUnit"`
	ProductionOrderDestinationDeliveryUnit             *string     `json:"ProductionOrderDestinationDeliveryUnit"`
	ProductionOrderPlannedScrapQtyInBaseUnit           *float32    `json:"ProductionOrderPlannedScrapQtyInBaseUnit"`
	ProductionOrderPlannedStartDate                    *string     `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime                    *string     `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate                      *string     `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime                      *string     `json:"ProductionOrderPlannedEndTime"`
	ProductionOrderActualReleaseDate                   *string     `json:"ProductionOrderActualReleaseDate"`
	ProductionOrderActualReleaseTime                   *string     `json:"ProductionOrderActualReleaseTime"`
	ProductionOrderActualStartDate                     *string     `json:"ProductionOrderActualStartDate"`
	ProductionOrderActualStartTime                     *string     `json:"ProductionOrderActualStartTime"`
	ProductionOrderActualEndDate                       *string     `json:"ProductionOrderActualEndDate"`
	ProductionOrderActualEndTime                       *string     `json:"ProductionOrderActualEndTime"`
	PlannedOrder                                       *int        `json:"PlannedOrder"`
	OrderID                                            *int        `json:"OrderID"`
	OrderItem                                          *int        `json:"OrderItem"`
	Project                                            *int        `json:"Project"`
	WBSElement                                         *int        `json:"WBSElement"`
	InspectionPlantBusinessPartner                	   *int        `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                               	   *string     `json:"InspectionPlant"`
	InspectionPlan                                	   *int        `json:"InspectionPlan"`
	InspectionLot                                 	   *int        `json:"InspectionLot"`
	ProductionOrderHeaderText                          *string     `json:"ProductionOrderHeaderText"`
	CertificateAuthorityChain        				   *string 	   `json:"CertificateAuthorityChain"`
	UsageControlChain        		 				   *string     `json:"UsageControlChain"`
	CreationDate                                       *string     `json:"CreationDate"`
	CreationTime                                       *string     `json:"CreationTime"`
	LastChangeDate                                     *string     `json:"LastChangeDate"`
	LastChangeTime                                     *string     `json:"LastChangeTime"`
	IsReleased                                         *bool       `json:"IsReleased"`
	IsPartiallyConfirmed                               *bool       `json:"IsPartiallyConfirmed"`
	IsConfirmed                                        *bool       `json:"IsConfirmed"`
	IsLocked                                           *bool       `json:"IsLocked"`
	IsCancelled                                        *bool       `json:"IsCancelled"`
	IsMarkedForDeletion                                *bool       `json:"IsMarkedForDeletion"`
	Item                                               []Item   `json:"Item"`
}

type Item struct {
	ProductionOrder                               int             `json:"ProductionOrder"`
	ProductionOrderItem                           int             `json:"ProductionOrderItem"`
	ProductionOrderItemDate                       *string  		  `json:"ProductionOrderItemDate"`
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
	ProductionPlant                               *string         `json:"Productionlant"`
	ProductionPlantStorageLocation                *string         `json:"ProductionPlantStorageLocation"`
	DeliverFromParty                              *int            `json:"DeliverFromParty"`
	DeliverFromPlant                              *string         `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation               *string         `json:"DeliverFromPlantStorageLocation"`
	DeliverToParty                                *int            `json:"DeliverToParty"`
	DeliverToPlant                                *string         `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation                 *string         `json:"DeliverToPlantStorageLocation"`
	ProductBaseUnit                               *string         `json:"ProductBaseUnit"`
	ProductProductionUnit                         *string         `json:"ProductProductionUnit"`
	ProductDeliveryUnit                           *string         `json:"ProductDeliveryUnit"`
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
	Project                                       *int     		  `json:"Project"`
	WBSElement                                    *int     		  `json:"WBSElement"`
	InspectionPlantBusinessPartner                *int     		  `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                               *string  		  `json:"InspectionPlant"`
	InspectionPlan                                *int     		  `json:"InspectionPlan"`
	InspectionLot                                 *int     		  `json:"InspectionLot"`
	ProductIsBatchManagedInProductionPlant        *bool           `json:"ProductIsBatchManagedInProductionPlant"`
	BatchMgmtPolicyInProductionOrder              *string         `json:"BatchMgmtPolicyInProductionOrder"`
	ProductionOrderTargetedBatch                  *string         `json:"ProductionOrderTargetedBatch"`
	ProductionOrderTargetedBatchValidityStartDate *string         `json:"ProductionOrderTargetedBatchValidityStartDate"`
	ProductionOrderTargetedBatchValidityStartTime *string         `json:"ProductionOrderTargetedBatchValidityStartTime"`
	ProductionOrderTargetedBatchValidityEndDate   *string         `json:"ProductionOrderTargetedBatchValidityEndDate"`
	ProductionOrderTargetedBatchValidityEndTime   *string         `json:"ProductionOrderTargetedBatchValidityEndTime"`
	ProductionOrderItemText                       *string         `json:"ProductionOrderItemText"`
	CreationDate                                  string          `json:"CreationDate"`
	CreationTime                                  string          `json:"CreationTime"`
	LastChangeDate                                string          `json:"LastChangeDate"`
	LastChangeTime                                string          `json:"LastChangeTime"`
	IsReleased                                    *bool           `json:"IsReleased"`
	IsPartiallyConfirmed                          *bool           `json:"IsPartiallyConfirmed"`
	IsConfirmed                                   *bool           `json:"IsConfirmed"`
	IsLocked                                      *bool           `json:"IsLocked"`
	IsCancelled                                   *bool           `json:"IsCancelled"`
	IsMarkedForDeletion                           *bool           `json:"IsMarkedForDeletion"`
	ItemComponent                                 []ItemComponent `json:"ItemComponent"`
	ItemOperation                                 []ItemOperation `json:"ItemOperation"`
}

type ItemComponent struct {
	ProductionOrder                                int      `json:"ProductionOrder"`
	ProductionOrderItem                            int      `json:"ProductionOrderItem"`
	BillOfMaterial                                 int      `json:"BillOfMaterial"`
	BillOfMaterialItem                             int      `json:"BillOfMaterialItem"`
	SupplyChainRelationshipID                      *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID              *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID         *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID        *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	ProductionPlantBusinessPartner                 *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                                *string  `json:"ProductionPlant"`
	MRPArea                                        *string  `json:"MRPArea"`
	MRPController                                  *string  `json:"MRPController"`
	ProductionVersion                              *int     `json:"ProductionVersion"`
	ProductionVersionItem                          *int     `json:"ProductionVersionItem"`
	ComponentProduct                               *string  `json:"ComponentProduct"`
	ComponentProductBuyer                          *int     `json:"ComponentProductBuyer"`
	ComponentProductSeller                         *int     `json:"ComponentProductSeller"`
	ComponentProductDeliverToParty                 *int     `json:"ComponentProductDeliverToParty"`
	ComponentProductDeliverToPlant                 *string  `json:"ComponentProductDeliverToPlant"`
	ComponentProductDeliverFromParty               *int     `json:"ComponentProductDeliverFromParty"`
	ComponentProductDeliverFromPlant               *string  `json:"ComponentProductDeliverFromPlant"`
	ComponentProductBaseUnit                       *string  `json:"ComponentProductBaseUnit"`
	ComponentProductDeliveryUnit                   *string  `json:"ComponentProductDeliveryUnit"`
	ComponentProductRequirementDate                *string  `json:"ComponentProductRequirementDate"`
	ComponentProductRequirementTime                *string  `json:"ComponentProductRequirementTime"`
	ComponentProductRequiredQuantityInBaseUnit     *float32 `json:"ComponentProductRequiredQuantityInBaseUnit"`
	ComponentProductRequiredQuantityInDeliveryUnit *float32 `json:"ComponentProductRequiredQuantityInDeliveryUnit"`
	ComponentProductPlannedScrapInPercent          *float32 `json:"ComponentProductPlannedScrapInPercent"`
	ComponentProductIsMarkedForBackflush           *bool    `json:"ComponentProductIsMarkedForBackflush"`
	StockConfirmationBusinessPartner               *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                         *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantStorageLocation          *string  `json:"StockConfirmationPlantStorageLocation"`
	PlannedOrder                                   *int     `json:"PlannedOrder"`
	PlannedOrderItem                               *int     `json:"PlannedOrderItem"`
	BillOfMaterialItemText                         *string  `json:"BillOfMaterialItemText"`
	ComponentProductBatch                          *string  `json:"ComponentProductBatch"`
	ComponentProductBatchValidityStartDate         *string  `json:"ComponentProductBatchValidityStartDate"`
	ComponentProductBatchValidityStartTime         *string  `json:"ComponentProductBatchValidityStartTime"`
	ComponentProductBatchValidityEndDate           *string  `json:"ComponentProductBatchValidityEndDate"`
	ComponentProductBatchValidityEndTime           *string  `json:"ComponentProductBatchValidityEndTime"`
	ComponentProductCostingPolicy                  *string  `json:"ComponentProductCostingPolicy"`
	ComponentProductPriceUnitQty                   *int     `json:"ComponentProductPriceUnitQty"`
	ComponentProductStandardPrice                  *float32 `json:"ComponentProductStandardPrice"`
	ComponentProductMovingAveragePrice             *float32 `json:"ComponentProductMovingAveragePrice"`
	ComponentProductWithdrawnQuantity              *float32 `json:"ComponentProductWithdrawnQuantity"`
	CreationDate                                   *string  `json:"CreationDate"`
	CreationTime                                   *string  `json:"CreationTime"`
	LastChangeDate                                 *string  `json:"LastChangeDate"`
	LastChangeTime                                 *string  `json:"LastChangeTime"`
	ComponentProductAvailabilityIsNotChecked       *bool    `json:"ComponentProductAvailabilityIsNotChecked"`
	IsReleased                                     *bool    `json:"IsReleased"`
	IsLocked                                       *bool    `json:"IsLocked"`
	IsCancelled                                    *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                            *bool    `json:"IsMarkedForDeletion"`
}

type ItemOperation struct {
	ProductionOrder                                 int      `json:"ProductionOrder"`
	ProductionOrderItem                             int      `json:"ProductionOrderItem"`
	Operations                                      int      `json:"Operations"`
	OperationsItem                                  int      `json:"OperationsItem"`
	OperationID                                     int      `json:"OperationID"`
	OperationType                                   *string  `json:"OperationType"`
	SupplyChainRelationshipID                       *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID               *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID          *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID        *int     `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                         *string  `json:"Product"`
	Buyer                                           *int     `json:"Buyer"`
	Seller                                          *int     `json:"Seller"`
	DeliverFromParty                                *int     `json:"DeliverFromParty"`
	DeliverFromPlant                                *string  `json:"DeliverFromPlant"`
	DeliverToParty                                  *int     `json:"DeliverToParty"`
	DeliverToPlant                                  *string  `json:"DeliverToPlant"`
	ProductionPlantBusinessPartner                  *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                                 *string  `json:"ProductionPlant"`
	MRPArea                                         *string  `json:"MRPArea"`
	MRPController                                   *string  `json:"MRPController"`
	ProductionVersion                               *int     `json:"ProductionVersion"`
	ProductionVersionItem                           *int     `json:"ProductionVersionItem"`
	Sequence                                        *int     `json:"Sequence"`
	SequenceText                                    *string  `json:"SequenceText"`
	OperationText                                   *string  `json:"OperationText"`
	ProductBaseUnit                                 *string  `json:"ProductBaseUnit"`
	ProductProductionUnit                           *string  `json:"ProductProductionUnit"`
	ProductOperationUnit                            *string  `json:"ProductOperationUnit"`
	ProductDeliveryUnit                             *string  `json:"ProductDeliveryUnit"`
	StandardLotSizeQuantity                         *float32 `json:"StandardLotSizeQuantity"`
	MinimumLotSizeQuantity                          *float32 `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity                          *float32 `json:"MaximumLotSizeQuantity"`
	OperationPlannedQuantityInBaseUnit              *float32 `json:"OperationPlannedQuantityInBaseUnit"`
	OperationPlannedQuantityInProductionUnit        *float32 `json:"OperationPlannedQuantityInProductionUnit"`
	OperationPlannedQuantityInOperationUnit         *float32 `json:"OperationPlannedQuantityInOperationUnit"`
	OperationPlannedQuantityInDeliveryUnit          *float32 `json:"OperationPlannedQuantityInDeliveryUnit"`
	OperationPlannedScrapInPercent                  *float32 `json:"OperationPlannedScrapInPercent"`
	ResponsiblePlannerGroup                         *string  `json:"ResponsiblePlannerGroup"`
	PlainLongText                                   *string  `json:"PlainLongText"`
	WorkCenter                                      *int     `json:"WorkCenter"`
	CapacityCategory                                *string  `json:"CapacityCategory"`
	OperationCostingRelevancyType                   *string  `json:"OperationCostingRelevancyType"`
	OperationSetupType                              *string  `json:"OperationSetupType"`
	OperationSetupGroupCategory                     *string  `json:"OperationSetupGroupCategory"`
	OperationSetupGroup                             *string  `json:"OperationSetupGroup"`
	MaximumWaitDuration                             *float32 `json:"MaximumWaitDuration"`
	StandardWaitDuration                            *float32 `json:"StandardWaitDuration"`
	MinimumWaitDuration                             *float32 `json:"MinimumWaitDuration"`
	WaitDurationUnit                                *string  `json:"WaitDurationUnit"`
	MaximumQueDuration                              *float32 `json:"MaximumQueDuration"`
	StandardQueueDuration                           *float32 `json:"StandardQueueDuration"`
	MinimumQueueDuration                            *float32 `json:"MinimumQueueDuration"`
	QueDurationUnit                                 *string  `json:"QueDurationUnit"`
	MaximumMoveDuration                             *float32 `json:"MaximumMoveDuration"`
	StandardMoveDuration                            *float32 `json:"StandardMoveDuration"`
	MinimumMoveDuration                             *float32 `json:"MinimumMoveDuration"`
	MoveDurationUnit                                *string  `json:"MoveDurationUnit"`
	StandardDeliveryDuration                        *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit                    *string  `json:"StandardDeliveryDurationUnit"`
	CostElement                                     *string  `json:"CostElement"`
	OperationErlstSchedldExecStrtDte                *string  `json:"OperationErlstSchedldExecStrtDte"`
	OperationErlstSchedldExecStrtTme                *string  `json:"OperationErlstSchedldExecStrtTme"`
	OperationErlstSchedldExecEndDte                 *string  `json:"OperationErlstSchedldExecEndDte"`
	OperationErlstSchedldExecEndTme                 *string  `json:"OperationErlstSchedldExecEndTme"`
	OperationActualExecutionStartDate               *string  `json:"OperationActualExecutionStartDate"`
	OperationActualExecutionStartTime               *string  `json:"OperationActualExecutionStartTime"`
	OperationActualExecutionEndDate                 *string  `json:"OperationActualExecutionEndDate"`
	OperationActualExecutionEndTime                 *string  `json:"OperationActualExecutionEndTime"`
	OperationConfirmedYieldQuantityInBaseUnit       *float32 `json:"OperationConfirmedYieldQuantityInBaseUnit"`
	OperationConfirmedYieldQuantityInProductionUnit *float32 `json:"OperationConfirmedYieldQuantityInProductionUnit"`
	OperationConfirmedYieldQuantityInOperationUnit  *float32 `json:"OperationConfirmedYieldQuantityInOperationUnit"`
	CreationDate                                    *string  `json:"CreationDate"`
	CreationTime                                    *string  `json:"CreationTime"`
	LastChangeDate                                  *string  `json:"LastChangeDate"`
	LastChangeTime                                  *string  `json:"LastChangeTime"`
	IsReleased                                      *bool    `json:"IsReleased"`
	IsPartiallyConfirmed                            *bool    `json:"IsPartiallyConfirmed"`
	IsConfirmed                                     *bool    `json:"IsConfirmed"`
	IsLocked                                        *bool    `json:"IsLocked"`
	IsCancelled                                     *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion"`
}

func CreateProductionOrderRequestHeader(
	requestPram *apiInputReader.Request,
	productionOrderHeader *apiInputReader.ProductionOrderHeader,
) ProductionOrderReq {
	req := ProductionOrderReq{
		Header: Header{
			ProductionOrder:     productionOrderHeader.ProductionOrder,
			IsReleased:          productionOrderHeader.IsReleased,
			IsCancelled:         productionOrderHeader.IsCancelled,
			IsMarkedForDeletion: productionOrderHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreateProductionOrderRequestHeaderByOwnerProductionPlantBP(
	requestPram *apiInputReader.Request,
	productionOrderHeader *apiInputReader.ProductionOrderHeader,
) ProductionOrderReq {
	req := ProductionOrderReq{
		Header: Header{
			OwnerProductionPlantBusinessPartner: requestPram.BusinessPartner,
			IsReleased:                          productionOrderHeader.IsReleased,
			IsCancelled:                         productionOrderHeader.IsCancelled,
			IsMarkedForDeletion:                 productionOrderHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersByOwnerProductionPlantBP",
		},
	}
	return req
}

func CreateProductionOrderRequestItem(
	requestPram *apiInputReader.Request,
	productionOrderItems *apiInputReader.ProductionOrderItem,
) ProductionOrderReq {
	req := ProductionOrderReq{
		Header: Header{
			ProductionOrder: productionOrderItems.ProductionOrder,
			Item: []Item{
				{
					ProductionOrderItem: productionOrderItems.ProductionOrderItem,
					IsReleased:          productionOrderItems.IsReleased,
					IsCancelled:         productionOrderItems.IsCancelled,
					IsMarkedForDeletion: productionOrderItems.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Item",
		},
	}
	return req
}

func CreateProductionOrderRequestItems(
	requestPram *apiInputReader.Request,
	productionOrderItems *apiInputReader.ProductionOrderItem,
) ProductionOrderReq {
	req := ProductionOrderReq{
		Header: Header{
			ProductionOrder: productionOrderItems.ProductionOrder,
			Item: []Item{
				{
					IsReleased:          productionOrderItems.IsReleased,
					IsCancelled:         productionOrderItems.IsCancelled,
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

func CreateProductionOrderRequestItemComponent(
	requestPram *apiInputReader.Request,
	productionOrderItemComponent *apiInputReader.ProductionOrderItemComponent,
) ProductionOrderReq {
	req := ProductionOrderReq{
		Header: Header{
			ProductionOrder: productionOrderItemComponent.ProductionOrder,
			Item: []Item{
				{
					ProductionOrderItem: productionOrderItemComponent.ProductionOrderItem,
					ItemComponent: []ItemComponent{
						{
							BillOfMaterial:      productionOrderItemComponent.BillOfMaterial,
							BillOfMaterialItem:  productionOrderItemComponent.BillOfMaterialItem,
							IsReleased:          productionOrderItemComponent.IsReleased,
							IsCancelled:         productionOrderItemComponent.IsCancelled,
							IsMarkedForDeletion: productionOrderItemComponent.IsMarkedForDeletion,
						},
					},
				},
			},
		},
		Accepter: []string{
			"ItemComponent",
		},
	}
	return req
}

func CreateProductionOrderRequestItemComponents(
	requestPram *apiInputReader.Request,
	productionOrderItemComponents *apiInputReader.ProductionOrderItemComponent,
) ProductionOrderReq {
	req := ProductionOrderReq{
		Header: Header{
			ProductionOrder: productionOrderItemComponents.ProductionOrder,
			Item: []Item{
				{
					ProductionOrderItem: productionOrderItemComponents.ProductionOrderItem,
					ItemComponent: []ItemComponent{
						{
							IsReleased:          productionOrderItemComponents.IsReleased,
							IsCancelled:         productionOrderItemComponents.IsCancelled,
							IsMarkedForDeletion: productionOrderItemComponents.IsMarkedForDeletion,
						},
					},
				},
			},
		},
		Accepter: []string{
			"ItemComponents",
		},
	}
	return req
}

func CreateProductionOrderRequestItemOperation(
	requestPram *apiInputReader.Request,
	productionOrderItemOperation *apiInputReader.ProductionOrderItemOperation,
) ProductionOrderReq {
	req := ProductionOrderReq{
		Header: Header{
			ProductionOrder: productionOrderItemOperation.ProductionOrder,
			Item: []Item{
				{
					ProductionOrderItem: productionOrderItemOperation.ProductionOrderItem,
					ItemOperation: []ItemOperation{
						{
							Operations:          productionOrderItemOperation.Operations,
							OperationsItem:      productionOrderItemOperation.OperationsItem,
							OperationID:         productionOrderItemOperation.OperationID,
							IsReleased:          productionOrderItemOperation.IsReleased,
							IsCancelled:         productionOrderItemOperation.IsCancelled,
							IsMarkedForDeletion: productionOrderItemOperation.IsMarkedForDeletion,
						},
					},
				},
			},
		},
		Accepter: []string{
			"ItemOperation",
		},
	}
	return req
}

func CreateProductionOrderRequestItemOperations(
	requestPram *apiInputReader.Request,
	productionOrderItemOperations *apiInputReader.ProductionOrderItemOperation,
) ProductionOrderReq {
	req := ProductionOrderReq{
		Header: Header{
			ProductionOrder: productionOrderItemOperations.ProductionOrder,
			Item: []Item{
				{
					ProductionOrderItem: productionOrderItemOperations.ProductionOrderItem,
					ItemOperation: []ItemOperation{
						{
							IsReleased:          productionOrderItemOperations.IsReleased,
							IsCancelled:         productionOrderItemOperations.IsCancelled,
							IsMarkedForDeletion: productionOrderItemOperations.IsMarkedForDeletion,
						},
					},
				},
			},
		},
		Accepter: []string{
			"ItemOperations",
		},
	}
	return req
}

//func CreateProductionOrderRequestItemOperationComponents(
//	requestPram *apiInputReader.Request,
//	productionOrderItemOperationComponent *apiInputReader.ProductionOrderItemOperationComponent,
//) ProductionOrderReq {
//	req := ProductionOrderReq{
//		Header: Header{
//			ProductionOrder: productionOrderItemOperationComponent.ProductionOrder,
//			Item: []Item{
//				{
//					ProductionOrderItem: productionOrderItemOperationComponent.ProductionOrderItem,
//					IsMarkedForDeletion: productionOrderItemOperationComponent.IsMarkedForDeletion,
//				},
//			},
//		},
//		Accepter: []string{
//			"ItemOperationComponents",
//		},
//	}
//	return req
//}

func ProductionOrderReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductionOrder,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCTION_ORDER_SRV"
	aPIType := "reads"

	var request ProductionOrderReq

	if accepter == "Header" {
		request = CreateProductionOrderRequestHeader(
			requestPram,
			&apiInputReader.ProductionOrderHeader{
				ProductionOrder:     input.ProductionOrderHeader.ProductionOrder,
				IsReleased:          input.ProductionOrderHeader.IsReleased,
				IsCancelled:         input.ProductionOrderHeader.IsCancelled,
				IsMarkedForDeletion: input.ProductionOrderHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "HeadersByOwnerProductionPlantBP" {
		request = CreateProductionOrderRequestHeaderByOwnerProductionPlantBP(
			requestPram,
			&apiInputReader.ProductionOrderHeader{
				IsReleased:          input.ProductionOrderHeader.IsReleased,
				IsCancelled:         input.ProductionOrderHeader.IsCancelled,
				IsMarkedForDeletion: input.ProductionOrderHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Item" {
		request = CreateProductionOrderRequestItem(
			requestPram,
			&apiInputReader.ProductionOrderItem{
				ProductionOrder:     input.ProductionOrderItem.ProductionOrder,
				ProductionOrderItem: input.ProductionOrderItem.ProductionOrderItem,
				IsReleased:          input.ProductionOrderItem.IsReleased,
				IsCancelled:         input.ProductionOrderItem.IsCancelled,
				IsMarkedForDeletion: input.ProductionOrderItem.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Items" {
		request = CreateProductionOrderRequestItems(
			requestPram,
			&apiInputReader.ProductionOrderItem{
				ProductionOrder:     input.ProductionOrderItem.ProductionOrder,
				IsReleased:          input.ProductionOrderItem.IsReleased,
				IsCancelled:         input.ProductionOrderItem.IsCancelled,
				IsMarkedForDeletion: input.ProductionOrderItem.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "ItemComponent" {
		request = CreateProductionOrderRequestItemComponent(
			requestPram,
			&apiInputReader.ProductionOrderItemComponent{
				ProductionOrder:     input.ProductionOrderItemComponent.ProductionOrder,
				ProductionOrderItem: input.ProductionOrderItemComponent.ProductionOrderItem,
				BillOfMaterial:      input.ProductionOrderItemComponent.BillOfMaterial,
				BillOfMaterialItem:  input.ProductionOrderItemComponent.BillOfMaterialItem,
				IsReleased:          input.ProductionOrderItemComponent.IsReleased,
				IsCancelled:         input.ProductionOrderItemComponent.IsCancelled,
				IsMarkedForDeletion: input.ProductionOrderItemComponent.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "ItemComponents" {
		request = CreateProductionOrderRequestItemComponents(
			requestPram,
			&apiInputReader.ProductionOrderItemComponent{
				ProductionOrder:     input.ProductionOrderItemComponent.ProductionOrder,
				ProductionOrderItem: input.ProductionOrderItemComponent.ProductionOrderItem,
				IsReleased:          input.ProductionOrderItemComponent.IsReleased,
				IsCancelled:         input.ProductionOrderItemComponent.IsCancelled,
				IsMarkedForDeletion: input.ProductionOrderItemComponent.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "ItemOperation" {
		request = CreateProductionOrderRequestItemOperation(
			requestPram,
			&apiInputReader.ProductionOrderItemOperation{
				ProductionOrder:     input.ProductionOrderItemOperation.ProductionOrder,
				ProductionOrderItem: input.ProductionOrderItemOperation.ProductionOrderItem,
				Operations:          input.ProductionOrderItemOperation.Operations,
				OperationsItem:      input.ProductionOrderItemOperation.OperationsItem,
				OperationID:         input.ProductionOrderItemOperation.OperationID,
				IsReleased:          input.ProductionOrderItemOperation.IsReleased,
				IsCancelled:         input.ProductionOrderItemOperation.IsCancelled,
				IsMarkedForDeletion: input.ProductionOrderItemOperation.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "ItemOperations" {
		request = CreateProductionOrderRequestItemOperations(
			requestPram,
			&apiInputReader.ProductionOrderItemOperation{
				ProductionOrder:     input.ProductionOrderItemOperation.ProductionOrder,
				ProductionOrderItem: input.ProductionOrderItemOperation.ProductionOrderItem,
				IsReleased:          input.ProductionOrderItemOperation.IsReleased,
				IsCancelled:         input.ProductionOrderItemOperation.IsCancelled,
				IsMarkedForDeletion: input.ProductionOrderItemOperation.IsMarkedForDeletion,
			},
		)
	}

	//	if accepter == "ItemOperationComponents" {
	//		request = CreateProductionOrderRequestItemOperationComponents(
	//			requestPram,
	//			&apiInputReader.ProductionOrderItemOperationComponent{
	//				ProductionOrder:     input.ProductionOrderItemOperationComponent.ProductionOrder,
	//				ProductionOrderItem: input.ProductionOrderItemOperationComponent.ProductionOrderItem,
	//				IsMarkedForDeletion: input.ProductionOrderItemOperationComponent.IsMarkedForDeletion,
	//			},
	//		)
	//	}

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
