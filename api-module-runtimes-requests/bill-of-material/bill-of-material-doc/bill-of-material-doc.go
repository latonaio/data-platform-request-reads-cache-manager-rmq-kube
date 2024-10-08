package apiModuleRuntimesRequestsBillOfMaterialDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type BillOfMaterialDocReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	HeaderDoc         HeaderDoc `json:"BillOfMaterial"`
	Accepter          []string  `json:"accepter"`
}

type HeaderDoc struct {
	BillOfMaterial           *int    `json:"BillOfMaterial"`
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
	BillOfMaterial           *int    `json:"BillOfMaterial"`
	BillOfMaterialItem       *int    `json:"BillOfMaterialItem"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreateBillOfMaterialDocRequestHeaderDoc(
	requestPram *apiInputReader.Request,
	headerDoc apiInputReader.BillOfMaterialHeaderDoc,
) BillOfMaterialDocReq {
	docIssuerBusinessPartner := 201 // TODO 暫定対応

	req := BillOfMaterialDocReq{
		HeaderDoc: HeaderDoc{
			BillOfMaterial:           &headerDoc.BillOfMaterial,
			DocType:                  headerDoc.DocType,
			DocIssuerBusinessPartner: &docIssuerBusinessPartner,
			//DocIssuerBusinessPartner:	headerDoc.DocIssuerBusinessPartner,
		},
		Accepter: []string{
			"HeaderDoc",
		},
	}
	return req
}

func CreateBillOfMaterialDocRequestItemDoc(
	requestPram *apiInputReader.Request,
	itemDoc apiInputReader.BillOfMaterialItemDoc,
) BillOfMaterialDocReq {
	docIssuerBusinessPartner := 201 // TODO 暫定対応

	req := BillOfMaterialDocReq{
		HeaderDoc: HeaderDoc{
			BillOfMaterial: &itemDoc.BillOfMaterial,
			ItemDoc: ItemDoc{
				BillOfMaterial:           &itemDoc.BillOfMaterial,
				BillOfMaterialItem:       &itemDoc.BillOfMaterialItem,
				DocType:                  itemDoc.DocType,
				DocIssuerBusinessPartner: &docIssuerBusinessPartner,
				//DocIssuerBusinessPartner:	itemDoc.DocIssuerBusinessPartner,
			},
		},
		Accepter: []string{
			"ItemDoc",
		},
	}
	return req
}

func BillOfMaterialDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.BillOfMaterial,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_BILL_OF_MATERIAL_DOC_SRV"
	aPIType := "reads"

	var request BillOfMaterialDocReq

	if accepter == "HeaderDoc" {
		request = CreateBillOfMaterialDocRequestHeaderDoc(
			requestPram,
			apiInputReader.BillOfMaterialHeaderDoc{
				BillOfMaterial:           input.BillOfMaterialHeaderDoc.BillOfMaterial,
				DocType:                  input.BillOfMaterialHeaderDoc.DocType,
				DocIssuerBusinessPartner: input.BillOfMaterialHeaderDoc.DocIssuerBusinessPartner,
			},
		)
	}

	if accepter == "ItemDoc" {
		request = CreateBillOfMaterialDocRequestItemDoc(
			requestPram,
			apiInputReader.BillOfMaterialItemDoc{
				BillOfMaterial:           input.BillOfMaterialItemDoc.BillOfMaterial,
				BillOfMaterialItem:       input.BillOfMaterialItemDoc.BillOfMaterialItem,
				DocType:                  input.BillOfMaterialItemDoc.DocType,
				DocIssuerBusinessPartner: input.BillOfMaterialItemDoc.DocIssuerBusinessPartner,
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
