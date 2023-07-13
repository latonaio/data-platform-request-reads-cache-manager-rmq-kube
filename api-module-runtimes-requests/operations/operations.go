package apiModuleRuntimesRequestsOperations

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type OperationsReq struct {
	Header   Header   `json:"Operations"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Operations                               int     `json:"Operations"`
	SupplyChainRelationshipID                *int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        *int    `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   *int    `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID *int    `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  *string `json:"Product"`
	Buyer                                    *int    `json:"Buyer"`
	Seller                                   *int    `json:"Seller"`
	DestinationDeliverToParty                *int    `json:"DestinationDeliverToParty"`
	DestinationDeliverToPlant                *string `json:"DestinationDeliverToPlant"`
	DepartureDeliverFromParty                *int    `json:"DepartureDeliverFromParty"`
	DepartureDeliverFromPlant                *string `json:"DepartureDeliverFromPlant"`
	OwnerProductionPlantBusinessPartner      *int    `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlant                     *string `json:"OwnerProductionPlant"`
	ProductBaseUnit                          *string `json:"ProductBaseUnit"`
	ProductDeliveryUnit                      *string `json:"ProductDeliveryUnit"`
	ProductProductionUnit                    *string `json:"ProductProductionUnit"`
	OperationsText                           *string `json:"OperationsText"`
	OperationsStatus                         *string `json:"OperationsStatus"`
	ResponsiblePlannerGroup                  *string `json:"ResponsiblePlannerGroup"`
	PlainLongText                            *string `json:"PlainLongText"`
	ValidityStartDate                        *string `json:"ValidityStartDate"`
	ValidityEndDate                          *string `json:"ValidityEndDate"`
	CreationDate                             *string `json:"CreationDate"`
	LastChangeDate                           *string `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool   `json:"IsMarkedForDeletion"`
	Item                                     []Item  `json:"Item"`
}

type Item struct {
	Operations                               int      `json:"Operations"`
	OperationsItem                           int      `json:"OperationsItem"`
	SupplyChainRelationshipID                *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID *int     `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  *string  `json:"Product"`
	Buyer                                    *int     `json:"Buyer"`
	Seller                                   *int     `json:"Seller"`
	DeliverToParty                           *int     `json:"DeliverToParty"`
	DeliverToPlant                           *string  `json:"DeliverToPlant"`
	DeliverFromParty                         *int     `json:"DeliverFromParty"`
	DeliverFromPlant                         *string  `json:"DeliverFromPlant"`
	ProductionPlantBusinessPartner           *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          *string  `json:"ProductionPlant"`
	OperationsText                           *string  `json:"OperationsText"`
	BillOfMaterial                           *int     `json:"BillOfMaterial"`
	OperationsStatus                         *string  `json:"OperationsStatus"`
	ResponsiblePlannerGroup                  *string  `json:"ResponsiblePlannerGroup"`
	OperationsUnit                           *string  `json:"OperationsUnit"`
	StandardLotSizeQuantity                  *float32 `json:"StandardLotSizeQuantity"`
	MinimumLotSizeQuantity                   *float32 `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity                   *float32 `json:"MaximumLotSizeQuantity"`
	PlainLongText                            *string  `json:"PlainLongText"`
	WorkCenter                               *int     `json:"WorkCenter"`
	ValidityStartDate                        *string  `json:"ValidityStartDate"`
	ValidityEndDate                          *string  `json:"ValidityEndDate"`
	CreationDate                             *string  `json:"CreationDate"`
	LastChangeDate                           *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}

func CreateOperationsRequestHeaderByOwnerProductionPlantBP(
	requestPram *apiInputReader.Request,
	operationsHeader *apiInputReader.OperationsHeader,
) OperationsReq {
	req := OperationsReq{
		Header: Header{
			OwnerProductionPlantBusinessPartner: requestPram.BusinessPartner,
			//IsMarkedForDeletion:                 operationsHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeaderByOwnerProductionPlantBP",
		},
	}
	return req
}

func CreateOperationsRequestHeader(
	requestPram *apiInputReader.Request,
	operationsHeader *apiInputReader.OperationsHeader,
) OperationsReq {
	req := OperationsReq{
		Header: Header{
			Operations: operationsHeader.Operations,
		},
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreateOperationsRequestItems(
	requestPram *apiInputReader.Request,
	operationsItems *apiInputReader.OperationsItems,
) OperationsReq {
	req := OperationsReq{
		Header: Header{
			Operations: operationsItems.Operations,
			Item: []Item{
				{
					IsMarkedForDeletion: operationsItems.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Items",
		},
	}
	return req
}

func OperationsReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Operations,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_OPERATIONS_SRV"
	aPIType := "reads"

	var request OperationsReq

	if accepter == "HeaderByOwnerProductionPlantBP" {
		request = CreateOperationsRequestHeaderByOwnerProductionPlantBP(
			requestPram,
			&apiInputReader.OperationsHeader{
				//IsMarkedForDeletion: input.OperationsHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Items" {
		request = CreateOperationsRequestItems(
			requestPram,
			&apiInputReader.OperationsItems{
				Operations: input.OperationsItems.Operations,
				//IsMarkedForDeletion: input.OperationsItems.IsMarkedForDeletion,
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
