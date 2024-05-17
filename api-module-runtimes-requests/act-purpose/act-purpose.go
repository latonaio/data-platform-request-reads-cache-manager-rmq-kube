package apiModuleRuntimesRequestsActPurpose

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ActPurposeReq struct {
	ActPurpose   ActPurpose    `json:"ActPurpose"`
	ActPurposes  []ActPurpose  `json:"ActPurposes"`
	Accepter     []string      `json:"accepter"`
}

type ActPurpose struct {
	ActPurpose          string  `json:"ActPurpose"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Text                []Text  `json:"Text"`
}

type Text struct {
	ActPurpose          string  `json:"ActPurpose"`
	Language            string  `json:"Language"`
	ActPurposeName      *string `json:"ActPurposeName"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func CreateActPurposeRequestActPurposesByActPurposes(
	requestPram *apiInputReader.Request,
	input []ActPurpose,
) ActPurposeReq {
	req := ActPurposeReq{
		ActPurposes: input,
		Accepter: []string{
			"ActPurposesByActPurposes",
		},
	}
	return req
}

func CreateActPurposeRequestActPurposes(
	requestPram *apiInputReader.Request,
	input apiInputReader.ActPurpose,
) ActPurposeReq {
	isMarkedForDeletion := false

	req := ActPurposeReq{
		ActPurposes: []ActPurpose{
			{
				ActPurpose:          input.ActPurpose,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"ActPurposes",
		},
	}
	return req
}

func CreateActPurposeRequestText(
	requestPram *apiInputReader.Request,
	input ActPurpose,
) ActPurposeReq {
	isMarkedForDeletion := false

	req := ActPurposeReq{
		ActPurpose: ActPurpose{
			ActPurpose:          input.ActPurpose,
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

func CreateActPurposeRequestTexts(
	requestPram *apiInputReader.Request,
	input ActPurpose,
) ActPurposeReq {
	isMarkedForDeletion := false

	req := ActPurposeReq{
		ActPurpose: ActPurpose{
			ActPurpose: input.ActPurpose,
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

func ActPurposeReadsActPurposes(
	requestPram *apiInputReader.Request,
	input apiInputReader.ActPurpose,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ACT_PURPOSE_SRV"
	aPIType := "reads"

	var request ActPurposeReq

	request = CreateActPurposeRequestActPurposes(
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

func ActPurposeReadsActPurposesByActPurposes(
	requestPram *apiInputReader.Request,
	input []ActPurpose,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ACT_PURPOSE_SRV"
	aPIType := "reads"

	var request ActPurposeReq

	request = CreateActPurposeRequestActPurposesByActPurposes(
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

func ActPurposeReadsText(
	requestPram *apiInputReader.Request,
	input ActPurpose,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ACT_PURPOSE_SRV"
	aPIType := "reads"

	var request ActPurposeReq

	request = CreateActPurposeRequestText(
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

func ActPurposeReadsTexts(
	requestPram *apiInputReader.Request,
	input ActPurpose,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ACT_PURPOSE_SRV"
	aPIType := "reads"

	var request ActPurposeReq

	request = CreateActPurposeRequestTexts(
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
