package apiModuleRuntimesRequestsDeilveryDocumentDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type DeliveryDocumentDocReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	HeaderDoc         HeaderDoc `json:"DeliveryDocument"`
	Accepter          []string  `json:"accepter"`
}

type HeaderDoc struct {
	DeliveryDocument         *int    `json:"DeliveryDocument"`
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
	DeliveryDocument         *int    `json:"DeliveryDocument"`
	DeliveryDocumentItem     *int    `json:"DeliveryDocumentItem"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreateDeliveryDocumentDocRequestItemDoc(
	requestPram *apiInputReader.Request,
	itemDoc apiInputReader.DeliveryDocumentDocItemDoc,
) DeliveryDocumentDocReq {
	req := DeliveryDocumentDocReq{
		HeaderDoc: HeaderDoc{
			DeliveryDocument: &itemDoc.DeliveryDocument,
			ItemDoc: ItemDoc{
				DeliveryDocument:         &itemDoc.DeliveryDocument,
				DeliveryDocumentItem:     &itemDoc.DeliveryDocumentItem,
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

func DeliveryDocumentDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.DeliveryDocument,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_DELIVERY_DOCUMENT_DOC_SRV"
	aPIType := "reads"

	var request DeliveryDocumentDocReq

	if accepter == "ItemDoc" {
		request = CreateDeliveryDocumentDocRequestItemDoc(
			requestPram,
			apiInputReader.DeliveryDocumentDocItemDoc{
				DeliveryDocument:         input.DeliveryDocumentDocItemDoc.DeliveryDocument,
				DeliveryDocumentItem:     input.DeliveryDocumentDocItemDoc.DeliveryDocumentItem,
				DocType:                  input.DeliveryDocumentDocItemDoc.DocType,
				DocIssuerBusinessPartner: input.DeliveryDocumentDocItemDoc.DocIssuerBusinessPartner,
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
