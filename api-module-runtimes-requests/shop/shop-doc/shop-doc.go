package apiModuleRuntimesRequestsShopDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ShopDocReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	HeaderDoc         HeaderDoc `json:"Shop"`
	Accepter          []string  `json:"accepter"`
}

type HeaderDoc struct {
	Shop                     *int    `json:"Shop"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreateShopDocRequestHeaderDoc(
	requestPram *apiInputReader.Request,
	headerDoc apiInputReader.ShopDocHeaderDoc,
) ShopDocReq {
	docIssuerBusinessPartner := 1001 // TODO 暫定対応
	docType := "IMAGE"

	req := ShopDocReq{
		HeaderDoc: HeaderDoc{
			Shop: &headerDoc.Shop,
			//DocType:                headerDoc.DocType,
			DocType:                  &docType,
			DocIssuerBusinessPartner: &docIssuerBusinessPartner,
			//DocIssuerBusinessPartner: headerDoc.DocIssuerBusinessPartner,
		},
		Accepter: []string{
			"HeaderDoc",
		},
	}
	return req
}

func ShopDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Shop,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_SHOP_DOC_SRV"
	aPIType := "reads"

	var request ShopDocReq

	if accepter == "HeaderDoc" {
		request = CreateShopDocRequestHeaderDoc(
			requestPram,
			apiInputReader.ShopDocHeaderDoc{
				Shop:                     input.ShopDocHeaderDoc.Shop,
				DocType:                  input.ShopDocHeaderDoc.DocType,
				DocIssuerBusinessPartner: input.ShopDocHeaderDoc.DocIssuerBusinessPartner,
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
