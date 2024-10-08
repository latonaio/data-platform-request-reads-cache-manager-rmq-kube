package apiModuleRuntimesRequestsArticleDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ArticleDocReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	HeaderDoc         HeaderDoc `json:"Article"`
	Accepter          []string  `json:"accepter"`
}

type HeaderDoc struct {
	Article                  *int    `json:"Article"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreateArticleDocRequestHeaderDoc(
	requestPram *apiInputReader.Request,
	headerDoc apiInputReader.ArticleDocHeaderDoc,
) ArticleDocReq {
	docIssuerBusinessPartner := 1001 // TODO 暫定対応
	docType := "IMAGE"

	req := ArticleDocReq{
		HeaderDoc: HeaderDoc{
			Article: &headerDoc.Article,
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

func ArticleDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Article,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_ARTICLE_DOC_SRV"
	aPIType := "reads"

	var request ArticleDocReq

	if accepter == "HeaderDoc" {
		request = CreateArticleDocRequestHeaderDoc(
			requestPram,
			apiInputReader.ArticleDocHeaderDoc{
				Article:                  input.ArticleDocHeaderDoc.Article,
				DocType:                  input.ArticleDocHeaderDoc.DocType,
				DocIssuerBusinessPartner: input.ArticleDocHeaderDoc.DocIssuerBusinessPartner,
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
