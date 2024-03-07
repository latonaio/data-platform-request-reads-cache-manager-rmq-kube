package apiModuleRuntimesRequestsIncoterms

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type IncotermsReq struct {
	Incoterms   Incoterms   `json:"Incoterms"`
	Incotermses []Incoterms `json:"Incotermses"`
	Accepter    []string    `json:"accepter"`
}

type Incoterms struct {
	Incoterms           string `json:"Incoterms"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
	Text                []Text `json:"Text"`
}

type Text struct {
	Incoterms           string `json:"Incoterms"`
	Language            string `json:"Language"`
	IncotermsName       string `json:"IncotermsName"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

func CreateIncotermsRequestIncotermsesByIncotermses(
	requestPram *apiInputReader.Request,
	input []Incoterms,
) IncotermsReq {
	req := IncotermsReq{
		Incotermses: input,
		Accepter: []string{
			"IncotermsesByIncotermses",
		},
	}
	return req
}

func CreateIncotermsRequestIncotermses(
	requestPram *apiInputReader.Request,
	input apiInputReader.Incoterms,
) IncotermsReq {
	isMarkedForDeletion := false

	req := IncotermsReq{
		Incotermses: []Incoterms{
			{
				Incoterms:           input.Incoterms,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"Incotermses",
		},
	}
	return req
}

func CreateIncotermsRequestText(
	requestPram *apiInputReader.Request,
	input Incoterms,
) IncotermsReq {
	isMarkedForDeletion := false

	req := IncotermsReq{
		Incoterms: Incoterms{
			Incoterms:           input.Incoterms,
			IsMarkedForDeletion: &isMarkedForDeletion,
			Text: []Text{
				{
					Language: "JA", // TODO 暫定で固定値を設定
				},
			},
		},
		Accepter: []string{
			"Text",
		},
	}
	return req
}

func CreateIncotermsRequestTexts(
	requestPram *apiInputReader.Request,
	input Incoterms,
) IncotermsReq {
	isMarkedForDeletion := false

	req := IncotermsReq{
		Incoterms: Incoterms{
			Incoterms: input.Incoterms,
			Text: []Text{
				{
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

func IncotermsReadsIncotermses(
	requestPram *apiInputReader.Request,
	input apiInputReader.Incoterms,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_INCOTERMS_SRV"
	aPIType := "reads"

	var request IncotermsReq

	request = CreateIncotermsRequestIncotermses(
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
	)

	return responseBody
}

func IncotermsReadsIncotermsesByIncotermses(
	requestPram *apiInputReader.Request,
	input []Incoterms,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_INCOTERMS_SRV"
	aPIType := "reads"

	var request IncotermsReq

	request = CreateIncotermsRequestIncotermsesByIncotermses(
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
	)

	return responseBody
}

func IncotermsReadsText(
	requestPram *apiInputReader.Request,
	input Incoterms,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_INCOTERMS_SRV"
	aPIType := "reads"

	var request IncotermsReq

	request = CreateIncotermsRequestText(
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
	)

	return responseBody
}

func IncotermsReadsTexts(
	requestPram *apiInputReader.Request,
	input Incoterms,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_INCOTERMS_SRV"
	aPIType := "reads"

	var request IncotermsReq

	request = CreateIncotermsRequestTexts(
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
	)

	return responseBody
}
