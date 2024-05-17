package apiModuleRuntimesRequestsPointConditionType

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type PointConditionTypeReq struct {
	PointConditionType   PointConditionType    `json:"PointConditionType"`
	PointConditionTypes  []PointConditionType  `json:"PointConditionTypes"`
	Accepter              []string `json:"accepter"`
}

type PointConditionType struct {
	PointConditionType  string  `json:"PointConditionType"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Text                []Text  `json:"Text"`
}

type Text struct {
	PointConditionType         string  `json:"PointConditionType"`
	Language                   string  `json:"Language"`
	PointConditionTypeName     *string `json:"PointConditionTypeName"`
	CreationDate               *string `json:"CreationDate"`
	LastChangeDate             *string `json:"LastChangeDate"`
	IsMarkedForDeletion        *bool   `json:"IsMarkedForDeletion"`
}

func CreatePointConditionTypeRequestPointConditionTypesByPointConditionTypes(
	requestPram *apiInputReader.Request,
	input []PointConditionType,
) PointConditionTypeReq {
	req := PointConditionTypeReq{
		PointConditionTypes: input,
		Accepter: []string{
			"PointConditionTypesByPointConditionTypes",
		},
	}
	return req
}

func CreatePointConditionTypeRequestPointConditionTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointConditionType,
) PointConditionTypeReq {
	isMarkedForDeletion := false

	req := PointConditionTypeReq{
		PointConditionTypes: []PointConditionType{
			{
				PointConditionType:  input.PointConditionType,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"PointConditionTypes",
		},
	}
	return req
}

func CreatePointConditionTypeRequestText(
	requestPram *apiInputReader.Request,
	input PointConditionType,
) PointConditionTypeReq {
	isMarkedForDeletion := false

	req := PointConditionTypeReq{
		PointConditionType: PointConditionType{
			PointConditionType:     input.PointConditionType,
			IsMarkedForDeletion:    &isMarkedForDeletion,
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

func CreatePointConditionTypeRequestTexts(
	requestPram *apiInputReader.Request,
	input PointConditionType,
) PointConditionTypeReq {
	isMarkedForDeletion := false

	req := PointConditionTypeReq{
		PointConditionType:        PointConditionType{
			PointConditionType:    input.PointConditionType,
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

func PointConditionTypeReadsPointConditionTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointConditionType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_CONDITION_TYPE_SRV"
	aPIType := "reads"

	var request PointConditionTypeReq

	request = CreatePointConditionTypeRequestPointConditionTypes(
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

func PointConditionTypeReadsPointConditionTypesByPointConditionTypes(
	requestPram *apiInputReader.Request,
	input []PointConditionType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_CONDITION_TYPE_SRV"
	aPIType := "reads"

	var request PointConditionTypeReq

	request = CreatePointConditionTypeRequestPointConditionTypesByPointConditionTypes(
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

func PointConditionTypeReadsText(
	requestPram *apiInputReader.Request,
	input PointConditionType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_CONDITION_TYPE_SRV"
	aPIType := "reads"

	var request PointConditionTypeReq

	request = CreatePointConditionTypeRequestText(
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

func PointConditionTypeReadsTexts(
	requestPram *apiInputReader.Request,
	input PointConditionType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_CONDITION_TYPE_SRV"
	aPIType := "reads"

	var request PointConditionTypeReq

	request = CreatePointConditionTypeRequestTexts(
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
