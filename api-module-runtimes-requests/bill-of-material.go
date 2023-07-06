package apiModuleRuntimesRequests

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type BillOfMaterialReq struct {
	Header   BillOfMaterialHeader `json:"BillOfMaterial"`
	Accepter []string             `json:"accepter"`
}

type BillOfMaterialHeader struct {
	BillOfMaterial              int                        `json:"BillOfMaterial"`
	BillOfMaterialType          *string                    `json:"BillOfMaterialType"`
	Product                     *string                    `json:"Product"`
	OwnerBusinessPartner        *int                       `json:"OwnerBusinessPartner"`
	OwnerPlant                  *string                    `json:"OwnerPlant"`
	BOMAlternativeText          *string                    `json:"BOMAlternativeText"`
	BOMHeaderBaseUnit           *string                    `json:"BOMHeaderBaseUnit"`
	BOMHeaderQuantityInBaseUnit *float32                   `json:"BOMHeaderQuantityInBaseUnit"`
	ValidityStartDate           *string                    `json:"ValidityStartDate"`
	ValidityEndDate             *string                    `json:"ValidityEndDate"`
	CreationDate                *string                    `json:"CreationDate"`
	LastChangeDate              *string                    `json:"LastChangeDate"`
	BillOfMaterialHeaderText    *string                    `json:"BillOfMaterialHeaderText"`
	IsMarkedForDeletion         *bool                      `json:"IsMarkedForDeletion"`
	Item                        []BillOfMaterialHeaderItem `json:"Item"`
}

type BillOfMaterialHeaderItem struct {
	BillOfMaterial                  int      `json:"BillOfMaterial"`
	BillOfMaterialItem              int      `json:"BillOfMaterialItem"`
	ComponentProduct                *string  `json:"ComponentProduct"`
	ComponentProductBusinessPartner *int     `json:"ComponentProductBusinessPartner"`
	StockConfirmationPlant          *string  `json:"StockConfirmationPlant"`
	BOMAlternativeText              *string  `json:"BOMAlternativeText"`
	BOMItemBaseUnit                 *string  `json:"BOMItemBaseUnit"`
	BOMItemQuantityInBaseUnit       *float32 `json:"BOMItemQuantityInBaseUnit"`
	ComponentScrapInPercent         *int     `json:"ComponentScrapInPercent"`
	ValidityStartDate               *string  `json:"ValidityStartDate"`
	ValidityEndDate                 *string  `json:"ValidityEndDate"`
	BillOfMaterialItemText          *string  `json:"BillOfMaterialItemText"`
	IsMarkedForDeletion             *bool    `json:"IsMarkedForDeletion"`
}

func CreateBillOfMaterialRequestHeaderByOwnerProductionPlantBP(
	requestPram *apiInputReader.Request,
	billOfMaterialParams *apiInputReader.BillOfMaterialListParams,
) BillOfMaterialReq {
	req := BillOfMaterialReq{
		Header: BillOfMaterialHeader{
			OwnerBusinessPartner: requestPram.BusinessPartner,
			IsMarkedForDeletion:  billOfMaterialParams.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeaderByOwnerProductionPlantBP",
		},
	}
	return req
}

func BillOfMaterialReads(
	requestPram *apiInputReader.Request,
	isMarkedForDeletion bool,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BILL_OF_MATERIAL_SRV"
	aPIType := "reads"

	request := CreateBillOfMaterialRequestHeaderByOwnerProductionPlantBP(
		&apiInputReader.Request{
			BusinessPartner: requestPram.BusinessPartner,
		},
		&apiInputReader.BillOfMaterialListParams{
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	)

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
