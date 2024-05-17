package apiModuleRuntimesRequestsEventDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type EventDocReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	HeaderDoc         HeaderDoc `json:"Event"`
	Accepter          []string  `json:"accepter"`
}

type HeaderDoc struct {
	Event                    *int    `json:"Event"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreateEventDocRequestHeaderDoc(
	requestPram *apiInputReader.Request,
	headerDoc apiInputReader.EventDocHeaderDoc,
) EventDocReq {
	docIssuerBusinessPartner := 1001 // TODO 暫定対応
	docType := "IMAGE"

	req := EventDocReq{
		HeaderDoc: HeaderDoc{
			Event: &headerDoc.Event,
			//DocType:                  headerDoc.DocType,
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

func EventDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Event,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_DOC_SRV"
	aPIType := "reads"

	var request EventDocReq

	if accepter == "HeaderDoc" {
		request = CreateEventDocRequestHeaderDoc(
			requestPram,
			apiInputReader.EventDocHeaderDoc{
				Event:                    input.EventDocHeaderDoc.Event,
				DocType:                  input.EventDocHeaderDoc.DocType,
				DocIssuerBusinessPartner: input.EventDocHeaderDoc.DocIssuerBusinessPartner,
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
