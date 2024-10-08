package apiModuleRuntimesRequestsProductStockDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ProductStockDocReq struct {
	BusinessPartnerID *int            `json:"business_partner"`
	ProductStockDoc   ProductStockDoc `json:"ProductStock"`
	Accepter          []string        `json:"accepter"`
}

type ProductStockDoc struct {
	Product                  *string `json:"Product"`
	BusinessPartner          *int    `json:"BusinessPartner"`
	Plant                    *string `json:"Plant"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreateProductionOrderDocRequestProductStockDoc(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductStockDocProductStockDoc,
) ProductStockDocReq {
	req := ProductStockDocReq{
		ProductStockDoc: ProductStockDoc{
			Product:                  input.Product,
			BusinessPartner:          input.BusinessPartner,
			Plant:                    input.Plant,
			DocType:                  &input.DocType,
			DocIssuerBusinessPartner: requestPram.BusinessPartner,
		},
		Accepter: []string{
			"ProductStockDoc",
		},
	}
	return req
}

func ProductStockDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductStock,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCT_STOCK_DOC_SRV"
	aPIType := "reads"

	var request ProductStockDocReq

	if accepter == "ProductStockDoc" {
		request = CreateProductionOrderDocRequestProductStockDoc(
			requestPram,
			apiInputReader.ProductStockDocProductStockDoc{
				Product:                  input.ProductStockDocProductStockDoc.Product,
				BusinessPartner:          input.ProductStockDocProductStockDoc.BusinessPartner,
				DocType:                  input.ProductStockDocProductStockDoc.DocType,
				DocIssuerBusinessPartner: input.ProductStockDocProductStockDoc.DocIssuerBusinessPartner,
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
