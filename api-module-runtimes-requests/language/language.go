package apiModuleRuntimesRequestsLanguage

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type LanguageReq struct {
	Language  Language   `json:"Language"`
	Languages []Language `json:"Languages"`
	Accepter  []string   `json:"accepter"`
}

type Language struct {
	Language            string `json:"Language"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
	Text                []Text `json:"Text"`
}

type Text struct {
	Language               string `json:"Language"`
	CorrespondenceLanguage string `json:"CorrespondenceLanguage"`
	LanguageName           string `json:"LanguageName"`
	CreationDate           string `json:"CreationDate"`
	LastChangeDate         string `json:"LastChangeDate"`
	IsMarkedForDeletion    *bool  `json:"IsMarkedForDeletion"`
}

func CreateLanguageRequestLanguages(
	requestPram *apiInputReader.Request,
	input apiInputReader.Language,
) LanguageReq {
	isMarkedForDeletion := false

	req := LanguageReq{
		Languages: []Language{
			{
				Language:            input.Language,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"Languages",
		},
	}
	return req
}

func CreateLanguageRequestText(
	requestPram *apiInputReader.Request,
	input Language,
) LanguageReq {
	isMarkedForDeletion := false

	req := LanguageReq{
		Language: Language{
			Language:            input.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
			Text: []Text{
				{
					CorrespondenceLanguage: "JA", // TODO 暫定で固定値を設定
					IsMarkedForDeletion:    &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Text",
		},
	}
	return req
}

func CreateLanguageRequestTexts(
	requestPram *apiInputReader.Request,
	input Language,
) LanguageReq {
	isMarkedForDeletion := false

	req := LanguageReq{
		Language: Language{
			Text: []Text{
				{
					CorrespondenceLanguage: *requestPram.Language,
					IsMarkedForDeletion:    &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Texts",
		},
	}
	return req
}

func LanguageReadsLanguages(
	requestPram *apiInputReader.Request,
	input apiInputReader.Language,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_LANGUAGE_SRV"
	aPIType := "reads"

	var request LanguageReq

	request = CreateLanguageRequestLanguages(
		requestPram,
		input,
	)

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

func LanguageReadsText(
	requestPram *apiInputReader.Request,
	input Language,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_LANGUAGE_SRV"
	aPIType := "reads"

	var request LanguageReq

	request = CreateLanguageRequestText(
		requestPram,
		input,
	)

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

func LanguageReadsTexts(
	requestPram *apiInputReader.Request,
	input Language,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_LANGUAGE_SRV"
	aPIType := "reads"

	var request LanguageReq

	request = CreateLanguageRequestTexts(
		requestPram,
		input,
	)

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
