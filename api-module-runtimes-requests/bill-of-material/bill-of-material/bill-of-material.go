package apiModuleRuntimesRequestsBillOfMaterial

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type BillOfMaterialReq struct {
	Header   Header   `json:"BillOfMaterial"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	BillOfMaterial                           int      `json:"BillOfMaterial"`
	BillOfMaterialType                       *string  `json:"BillOfMaterialType"`
	SupplyChainRelationshipID                *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID *int     `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  *string  `json:"Product"`
	Buyer                                    *int     `json:"Buyer"`
	Seller                                   *int     `json:"Seller"`
	DestinationDeliverToParty                *int     `json:"DestinationDeliverToParty"`
	DestinationDeliverToPlant                *string  `json:"DestinationDeliverToPlant"`
	DepartureDeliverFromParty                *int     `json:"DepartureDeliverFromParty"`
	DepartureDeliverFromPlant                *string  `json:"DepartureDeliverFromPlant"`
	OwnerProductionPlantBusinessPartner      *int     `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlant                     *string  `json:"OwnerProductionPlant"`
	ProductBaseUnit                          *string  `json:"ProductBaseUnit"`
	ProductDeliveryUnit                      *string  `json:"ProductDeliveryUnit"`
	ProductProductionUnit                    *string  `json:"ProductProductionUnit"`
	ProductStandardQuantityInBaseUnit        *float32 `json:"ProductStandardQuantityInBaseUnit"`
	ProductStandardQuantityInDeliveryUnit    *float32 `json:"ProductStandardQuantityInDeliveryUnit"`
	ProductStandardQuantityInProductionUnit  *float32 `json:"ProductStandardQuantityInProductionUnit"`
	BillOfMaterialHeaderText                 *string  `json:"BillOfMaterialHeaderText"`
	ValidityStartDate                        *string  `json:"ValidityStartDate"`
	ValidityEndDate                          *string  `json:"ValidityEndDate"`
	CreationDate                             *string  `json:"CreationDate"`
	LastChangeDate                           *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
	Item                                     []Item   `json:"Item"`
}

type Item struct {
	BillOfMaterial                                 int      `json:"BillOfMaterial"`
	BillOfMaterialItem                             int      `json:"BillOfMaterialItem"`
	SupplyChainRelationshipID                      *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID              *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID         *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID        *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	Product                                        *string  `json:"Product"`
	ProductionPlantBusinessPartner                 *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                                *string  `json:"ProductionPlant"`
	ComponentProduct                               *string  `json:"ComponentProduct"`
	ComponentProductBuyer                          *int     `json:"ComponentProductBuyer"`
	ComponentProductSeller                         *int     `json:"ComponentProductSeller"`
	ComponentProductDeliverToParty                 *int     `json:"ComponentProductDeliverToParty"`
	ComponentProductDeliverToPlant                 *string  `json:"ComponentProductDeliverToPlant"`
	ComponentProductDeliverFromParty               *int     `json:"ComponentProductDeliverFromParty"`
	ComponentProductDeliverFromPlant               *string  `json:"ComponentProductDeliverFromPlant"`
	StockConfirmationBusinessPartner               *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                         *string  `json:"StockConfirmationPlant"`
	ComponentProductStandardQuantityInBaseUnit     *float32 `json:"ComponentProductStandardQuantityInBaseUnit"`
	ComponentProductStandardQuantityInDeliveryUnit *float32 `json:"ComponentProductStandardQuantityInDeliveryUnit"`
	ComponentProductBaseUnit                       *string  `json:"ComponentProductBaseUnit"`
	ComponentProductDeliveryUnit                   *string  `json:"ComponentProductDeliveryUnit"`
	ComponentProductStandardScrapInPercent         *float32 `json:"ComponentProductStandardScrapInPercent"`
	IsMarkedForBackflush                           *bool    `json:"IsMarkedForBackflush"`
	BillOfMaterialItemText                         *string  `json:"BillOfMaterialItemText"`
	ValidityStartDate                              *string  `json:"ValidityStartDate"`
	ValidityEndDate                                *string  `json:"ValidityEndDate"`
	CreationDate                                   *string  `json:"CreationDate"`
	LastChangeDate                                 *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                            *bool    `json:"IsMarkedForDeletion"`
}

func CreateBillOfMaterialRequestHeader(
	requestPram *apiInputReader.Request,
	billOfMaterialHeader *apiInputReader.BillOfMaterialHeader,
) BillOfMaterialReq {
	req := BillOfMaterialReq{
		Header: Header{
			BillOfMaterial: billOfMaterialHeader.BillOfMaterial,
		},
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreateBillOfMaterialRequestHeaderByOwnerProductionPlantBP(
	requestPram *apiInputReader.Request,
	billOfMaterialHeader *apiInputReader.BillOfMaterialHeader,
) BillOfMaterialReq {
	req := BillOfMaterialReq{
		Header: Header{
			OwnerProductionPlantBusinessPartner: requestPram.BusinessPartner,
			//IsMarkedForDeletion:                 billOfMaterialHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeaderByOwnerProductionPlantBP",
		},
	}
	return req
}

func CreateBillOfMaterialRequestItems(
	requestPram *apiInputReader.Request,
	billOfMaterialItems *apiInputReader.BillOfMaterialItems,
) BillOfMaterialReq {
	req := BillOfMaterialReq{
		Header: Header{
			BillOfMaterial: billOfMaterialItems.BillOfMaterial,
			Item: []Item{
				{
					//IsMarkedForDeletion: billOfMaterialItems.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Items",
		},
	}
	return req
}

func BillOfMaterialReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.BillOfMaterial,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_BILL_OF_MATERIAL_SRV"
	aPIType := "reads"

	var request BillOfMaterialReq

	// 一覧
	if accepter == "HeaderByOwnerProductionPlantBP" {
		request = CreateBillOfMaterialRequestHeaderByOwnerProductionPlantBP(
			requestPram,
			&apiInputReader.BillOfMaterialHeader{
				//IsMarkedForDeletion: input.BillOfMaterialHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Header" {
		request = CreateBillOfMaterialRequestHeader(
			requestPram,
			&apiInputReader.BillOfMaterialHeader{
				BillOfMaterial: input.BillOfMaterialHeader.BillOfMaterial,
			},
		)
	}

	// 詳細一覧（明細一覧）
	if accepter == "Items" {
		request = CreateBillOfMaterialRequestItems(
			requestPram,
			&apiInputReader.BillOfMaterialItems{
				BillOfMaterial:      input.BillOfMaterialItems.BillOfMaterial,
				//IsMarkedForDeletion: input.BillOfMaterialItems.IsMarkedForDeletion,
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
