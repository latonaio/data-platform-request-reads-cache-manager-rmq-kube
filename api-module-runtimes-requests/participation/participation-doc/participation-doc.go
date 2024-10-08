package apiModuleRuntimesRequestsParticipationDoc

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ParticipationDocReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	HeaderDoc         HeaderDoc `json:"Participation"`
	Accepter          []string  `json:"accepter"`
}

type HeaderDoc struct {
	Participation            *int    `json:"Participation"`
	DocType                  *string `json:"DocType"`
	FileExtension            *string `json:"FileExtension"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FilePath                 *string `json:"FilePath"`
	FileName                 *string `json:"FileName"`
}

func CreateParticipationDocRequestHeaderDoc(
	requestPram *apiInputReader.Request,
	headerDoc apiInputReader.ParticipationDocHeaderDoc,
) ParticipationDocReq {
	docIssuerBusinessPartner := 1001 // TODO 暫定対応
	docType := "IMAGE"

	if headerDoc.DocType != nil {
		docType = *headerDoc.DocType
	}

	req := ParticipationDocReq{
		HeaderDoc: HeaderDoc{
			Participation: &headerDoc.Participation,
			//DocType:                   headerDoc.DocType,
			DocType:                  &docType,
			DocIssuerBusinessPartner: &docIssuerBusinessPartner,
			//DocIssuerBusinessPartner:  headerDoc.DocIssuerBusinessPartner,
		},
		Accepter: []string{
			"HeaderDoc",
		},
	}
	return req
}

func ParticipationDocReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Participation,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_PARTICIPATION_DOC_SRV"
	aPIType := "reads"

	var request ParticipationDocReq

	if accepter == "HeaderDoc" {
		request = CreateParticipationDocRequestHeaderDoc(
			requestPram,
			apiInputReader.ParticipationDocHeaderDoc{
				Participation:            input.ParticipationDocHeaderDoc.Participation,
				DocType:                  input.ParticipationDocHeaderDoc.DocType,
				DocIssuerBusinessPartner: input.ParticipationDocHeaderDoc.DocIssuerBusinessPartner,
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
