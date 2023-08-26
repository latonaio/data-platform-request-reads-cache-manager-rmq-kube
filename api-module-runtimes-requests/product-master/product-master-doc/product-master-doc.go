package apiModuleRuntimesRequestsProductMasterDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ProductMasterDocReq struct {
	BusinessPartnerID *int       `json:"business_partner"`
	GeneralDoc        GeneralDoc `json:"ProductMaster"`
	Accepter          []string   `json:"accepter"`
}

type GeneralDoc struct {
	Product                  *string `json:"Product"`
	DocType                  string  `json:"DocType"`
	FileExtension            string  `json:"FileExtension"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 string  `json:"FilePath"`
	FileName                 string  `json:"FileName"`
}

func CreateProductMasterDocRequestGeneralDoc(
	requestPram *apiInputReader.Request,
) ProductMasterDocReq {
	req := ProductMasterDocReq{
		GeneralDoc: GeneralDoc{
			DocType:                  "IMAGE",
			DocIssuerBusinessPartner: requestPram.BusinessPartner,
		},
		Accepter: []string{},
	}
	return req
}

func ProductMasterDocReads(
	requestPram *apiInputReader.Request,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCT_MASTER_DOC_SRV"
	aPIType := "reads"

	var request ProductMasterDocReq

	if accepter == "GeneralDoc" {
		request = CreateProductMasterDocRequestGeneralDoc(
			requestPram,
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
