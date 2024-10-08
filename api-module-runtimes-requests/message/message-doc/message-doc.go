package apiModuleRuntimesRequestsMessageDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type MessageDocReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	HeaderDoc         HeaderDoc `json:"Message"`
	Accepter          []string  `json:"accepter"`
}

type HeaderDoc struct {
	Message                  *int    `json:"Message"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreateMessageDocRequestHeaderDoc(
	requestPram *apiInputReader.Request,
	headerDoc apiInputReader.MessageDocHeaderDoc,
) MessageDocReq {
	docIssuerBusinessPartner := 1001 // TODO 暫定対応
	docType := "IMAGE"

	req := MessageDocReq{
		HeaderDoc: HeaderDoc{
			Message: &headerDoc.Message,
			//DocType:                    headerDoc.DocType,
			DocType:                      &docType,
			DocIssuerBusinessPartner:     &docIssuerBusinessPartner,
			//DocIssuerBusinessPartner:   headerDoc.DocIssuerBusinessPartner,
		},
		Accepter: []string{
			"HeaderDoc",
		},
	}
	return req
}

func MessageDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Message,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_MESSAGE_DOC_SRV"
	aPIType := "reads"

	var request MessageDocReq

	if accepter == "HeaderDoc" {
		request = CreateMessageDocRequestHeaderDoc(
			requestPram,
			apiInputReader.MessageDocHeaderDoc{
				Message:                  input.MessageDocHeaderDoc.Message,
				DocType:                  input.MessageDocHeaderDoc.DocType,
				DocIssuerBusinessPartner: input.MessageDocHeaderDoc.DocIssuerBusinessPartner,
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
