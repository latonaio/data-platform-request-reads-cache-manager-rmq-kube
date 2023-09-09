package apiModuleRuntimesRequestsProductionOrderDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ProductionOrderDocReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	HeaderDoc         HeaderDoc `json:"ProductionOrder"`
	Accepter          []string  `json:"accepter"`
}

type HeaderDoc struct {
	ProductionOrder          *int    `json:"ProductionOrder"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
	ItemDoc                  ItemDoc `json:"ItemDoc"`
}

type ItemDoc struct {
	ProductionOrder          *int    `json:"ProductionOrder"`
	ProductionOrderItem      *int    `json:"ProductionOrderItem"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreateProductionOrderDocRequestHeaderDoc(
	requestPram *apiInputReader.Request,
	// headerDoc *apiInputReader.ProductionOrderDocItemDoc,
) ProductionOrderDocReq {
	req := ProductionOrderDocReq{
		HeaderDoc: HeaderDoc{
			//DocType:                  headerDoc.DocType,
			DocIssuerBusinessPartner: requestPram.BusinessPartner,
		},
		Accepter: []string{
			"HeaderDoc",
		},
	}
	return req
}

func CreateProductionOrderDocRequestItemDoc(
	requestPram *apiInputReader.Request,
	itemDoc apiInputReader.ProductionOrderDocItemDoc,
) ProductionOrderDocReq {
	req := ProductionOrderDocReq{
		HeaderDoc: HeaderDoc{
			ProductionOrder: &itemDoc.ProductionOrder,
			ItemDoc: ItemDoc{
				ProductionOrder:          &itemDoc.ProductionOrder,
				ProductionOrderItem:      &itemDoc.ProductionOrderItem,
				DocType:                  &itemDoc.DocType,
				DocIssuerBusinessPartner: &itemDoc.DocIssuerBusinessPartner,
			},
		},
		Accepter: []string{
			"ItemDoc",
		},
	}
	return req
}

func ProductionOrderDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductionOrder,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCTION_ORDER_DOC_SRV"
	aPIType := "reads"

	var request ProductionOrderDocReq

	if accepter == "HeaderDoc" {
		request = CreateProductionOrderDocRequestHeaderDoc(
			requestPram,
		)
	}

	if accepter == "ItemDoc" {
		request = CreateProductionOrderDocRequestItemDoc(
			requestPram,
			apiInputReader.ProductionOrderDocItemDoc{
				ProductionOrder:          input.ProductionOrderDocItemDoc.ProductionOrder,
				ProductionOrderItem:      input.ProductionOrderDocItemDoc.ProductionOrderItem,
				DocType:                  input.ProductionOrderDocItemDoc.DocType,
				DocIssuerBusinessPartner: input.ProductionOrderDocItemDoc.DocIssuerBusinessPartner,
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
