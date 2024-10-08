package apiModuleRuntimesRequestsSiteType

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type SiteTypeReq struct {
	SiteType  SiteType   `json:"SiteType"`
	SiteTypes []SiteType `json:"SiteTypes"`
	Accepter  []string   `json:"accepter"`
}

type SiteType struct {
	SiteType            string  `json:"SiteType"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Text                []Text  `json:"Text"`
}

type Text struct {
	SiteType            string  `json:"SiteType"`
	Language            string  `json:"Language"`
	SiteTypeName        *string `json:"SiteTypeName"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func CreateSiteTypeRequestSiteTypesBySiteTypes(
	requestPram *apiInputReader.Request,
	input []SiteType,
) SiteTypeReq {
	req := SiteTypeReq{
		SiteTypes: input,
		Accepter: []string{
			"SiteTypesBySiteTypes",
		},
	}
	return req
}

func CreateSiteTypeRequestSiteTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.SiteType,
) SiteTypeReq {
	isMarkedForDeletion := false

	req := SiteTypeReq{
		SiteTypes: []SiteType{
			{
				SiteType:            input.SiteType,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"SiteTypes",
		},
	}
	return req
}

func CreateSiteTypeRequestText(
	requestPram *apiInputReader.Request,
	input SiteType,
) SiteTypeReq {
	isMarkedForDeletion := false

	req := SiteTypeReq{
		SiteType: SiteType{
			SiteType:            input.SiteType,
			IsMarkedForDeletion: &isMarkedForDeletion,
			Text: []Text{
				{
					Language:            *requestPram.Language,
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Text",
		},
	}
	return req
}

func CreateSiteTypeRequestTexts(
	requestPram *apiInputReader.Request,
	input SiteType,
) SiteTypeReq {
	isMarkedForDeletion := false

	req := SiteTypeReq{
		SiteType: SiteType{
			SiteType: input.SiteType,
			Text: []Text{
				{
					Language:            *requestPram.Language,
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Texts",
		},
	}
	return req
}

func SiteTypeReadsSiteTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.SiteType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SITE_TYPE_SRV"
	aPIType := "reads"

	var request SiteTypeReq

	request = CreateSiteTypeRequestSiteTypes(
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

func SiteTypeReadsSiteTypesBySiteTypes(
	requestPram *apiInputReader.Request,
	input []SiteType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SITE_TYPE_SRV"
	aPIType := "reads"

	var request SiteTypeReq

	request = CreateSiteTypeRequestSiteTypesBySiteTypes(
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

func SiteTypeReadsText(
	requestPram *apiInputReader.Request,
	input SiteType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SITE_TYPE_SRV"
	aPIType := "reads"

	var request SiteTypeReq

	request = CreateSiteTypeRequestText(
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

func SiteTypeReadsTexts(
	requestPram *apiInputReader.Request,
	input SiteType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SITE_TYPE_SRV"
	aPIType := "reads"

	var request SiteTypeReq

	request = CreateSiteTypeRequestTexts(
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
