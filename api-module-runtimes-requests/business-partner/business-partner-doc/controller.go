package apiModuleRuntimesRequestsBusinessPartnerDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type BusinessPartnerDocReq struct {
	BusinessPartnerID *int       `json:"business_partner"`
	GeneralDoc        GeneralDoc `json:"BusinessPartner"`
	Accepter          []string   `json:"accepter"`
}

type GeneralDoc struct {
	BusinessPartner          *int    `json:"BusinessPartner"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreateBusinessPartnerDocRequestGeneralDoc(
	requestPram *apiInputReader.Request,
	generalDoc apiInputReader.BusinessPartnerDocGeneralDoc,
) BusinessPartnerDocReq {
	docIssuerBusinessPartner := 1001 // TODO 暫定対応
	docType := "IMAGE"

	if generalDoc.DocType != nil {
		docType = *generalDoc.DocType
	}

	req := BusinessPartnerDocReq{
		GeneralDoc: GeneralDoc{
			BusinessPartner: &generalDoc.BusinessPartner,
			//DocType:                generalDoc.DocType,
			DocType:                  &docType,
			DocIssuerBusinessPartner: &docIssuerBusinessPartner,
			//DocIssuerBusinessPartner: generalDoc.DocIssuerBusinessPartner,
		},
		Accepter: []string{
			"GeneralDoc",
		},
	}
	return req
}

func BusinessPartnerDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.BusinessPartner,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_DOC_SRV"
	aPIType := "reads"

	var request BusinessPartnerDocReq

	if accepter == "GeneralDoc" {
		request = CreateBusinessPartnerDocRequestGeneralDoc(
			requestPram,
			apiInputReader.BusinessPartnerDocGeneralDoc{
				BusinessPartner: input.BusinessPartnerDocGeneralDoc.BusinessPartner,
				DocType:         input.BusinessPartnerDocGeneralDoc.DocType,
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
