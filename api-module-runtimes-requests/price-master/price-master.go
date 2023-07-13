package apiModuleRuntimesRequestsPriceMaster

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type PriceMasterReq struct {
	Header   Header   `json:"PriceMaster"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID"`
	Buyer                      int      `json:"Buyer"`
	Seller                     int      `json:"Seller"`
	ConditionRecord            int      `json:"ConditionRecord"`
	ConditionSequentialNumber  int      `json:"ConditionSequentialNumber"`
	Product                    string   `json:"Product"`
	ConditionValidityStartDate string   `json:"ConditionValidityStartDate"`
	ConditionValidityEndDate   string   `json:"ConditionValidityEndDate"`
	ConditionType              *string  `json:"ConditionType"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit     *int     `json:"ConditionRateValueUnit"`
	ConditionScaleQuantity     *int     `json:"ConditionScaleQuantity"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	CreationDate               *string  `json:"CreationDate"`
	LastChangeDate             *string  `json:"LastChangeDate"`
	IsMarkedForDeletion        *bool    `json:"IsMarkedForDeletion"`
}

func CreatePriceMasterRequestHeaderByBuyer(
	requestPram *apiInputReader.Request,
	priceMasterHeader *apiInputReader.PriceMasterHeader,
) PriceMasterReq {
	req := PriceMasterReq{
		Header: Header{
			Buyer:               *priceMasterHeader.Buyer,
			IsMarkedForDeletion: priceMasterHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"PriceMastersByBuyer",
		},
	}
	return req
}

func CreatePriceMasterRequestHeaderBySeller(
	requestPram *apiInputReader.Request,
	priceMasterHeader *apiInputReader.PriceMasterHeader,
) PriceMasterReq {
	req := PriceMasterReq{
		Header: Header{
			Seller:              *priceMasterHeader.Seller,
			IsMarkedForDeletion: priceMasterHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"PriceMastersBySeller",
		},
	}
	return req
}

func PriceMasterReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.PriceMaster,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_PRICE_MASTER_SRV"
	aPIType := "reads"

	var request PriceMasterReq

	if accepter == "PriceMastersByBuyer" {
		request = CreatePriceMasterRequestHeaderByBuyer(
			requestPram,
			&apiInputReader.PriceMasterHeader{
				Buyer:               input.PriceMasterHeader.Buyer,
				IsMarkedForDeletion: input.PriceMasterHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "PriceMastersBySeller" {
		request = CreatePriceMasterRequestHeaderBySeller(
			requestPram,
			&apiInputReader.PriceMasterHeader{
				Seller:              input.PriceMasterHeader.Seller,
				IsMarkedForDeletion: input.PriceMasterHeader.IsMarkedForDeletion,
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
