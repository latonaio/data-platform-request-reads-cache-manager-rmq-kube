package apiModuleRuntimesRequestsPostDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type PostDocReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	HeaderDoc         HeaderDoc `json:"Post"`
	Accepter          []string  `json:"accepter"`
}

type HeaderDoc struct {
	Post                     *int    `json:"Post"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreatePostDocRequestHeaderDoc(
	requestPram *apiInputReader.Request,
	headerDoc apiInputReader.PostDocHeaderDoc,
) PostDocReq {
	docIssuerBusinessPartner := 1001 // TODO 暫定対応

	req := PostDocReq{
		HeaderDoc: HeaderDoc{
			Post: &headerDoc.Post,
			//DocType:                  headerDoc.DocType,
			DocIssuerBusinessPartner: &docIssuerBusinessPartner,
			//DocIssuerBusinessPartner: headerDoc.DocIssuerBusinessPartner,
		},
		Accepter: []string{
			"HeaderDoc",
		},
	}
	return req
}

func PostDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Post,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_POST_DOC_SRV"
	aPIType := "reads"

	var request PostDocReq

	if accepter == "HeaderDoc" {
		request = CreatePostDocRequestHeaderDoc(
			requestPram,
			apiInputReader.PostDocHeaderDoc{
				Post:                     input.PostDocHeaderDoc.Post,
				DocType:                  input.PostDocHeaderDoc.DocType,
				DocIssuerBusinessPartner: input.PostDocHeaderDoc.DocIssuerBusinessPartner,
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
