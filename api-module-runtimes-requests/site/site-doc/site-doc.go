package apiModuleRuntimesRequestsSiteDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type SiteDocReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	HeaderDoc         HeaderDoc `json:"Site"`
	Accepter          []string  `json:"accepter"`
}

type HeaderDoc struct {
	Site                     *int    `json:"Site"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreateSiteDocRequestHeaderDoc(
	requestPram *apiInputReader.Request,
	headerDoc apiInputReader.SiteDocHeaderDoc,
) SiteDocReq {
	docIssuerBusinessPartner := 1001 // TODO 暫定対応
	docType := "IMAGE"

	req := SiteDocReq{
		HeaderDoc: HeaderDoc{
			Site: &headerDoc.Site,
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

func SiteDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Site,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_SITE_DOC_SRV"
	aPIType := "reads"

	var request SiteDocReq

	if accepter == "HeaderDoc" {
		request = CreateSiteDocRequestHeaderDoc(
			requestPram,
			apiInputReader.SiteDocHeaderDoc{
				Site:                     input.SiteDocHeaderDoc.Site,
				DocType:                  input.SiteDocHeaderDoc.DocType,
				DocIssuerBusinessPartner: input.SiteDocHeaderDoc.DocIssuerBusinessPartner,
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
