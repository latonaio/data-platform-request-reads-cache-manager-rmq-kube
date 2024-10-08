package apiModuleRuntimesRequestsPointConsumptionType

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type PointConsumptionTypeReq struct {
	PointConsumptionType  PointConsumptionType   `json:"PointConsumptionType"`
	PointConsumptionTypes []PointConsumptionType `json:"PointConsumptionTypes"`
	Accepter              []string               `json:"accepter"`
}

type PointConsumptionType struct {
	PointConsumptionType string  `json:"PointConsumptionType"`
	CreationDate         *string `json:"CreationDate"`
	LastChangeDate       *string `json:"LastChangeDate"`
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
	Text                 []Text  `json:"Text"`
}

type Text struct {
	PointConsumptionType     string  `json:"PointConsumptionType"`
	Language                 string  `json:"Language"`
	PointConsumptionTypeName *string `json:"PointConsumptionTypeName"`
	CreationDate             *string `json:"CreationDate"`
	LastChangeDate           *string `json:"LastChangeDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

func CreatePointConsumptionTypeRequestPointConsumptionTypesByPointConsumptionTypes(
	requestPram *apiInputReader.Request,
	input []PointConsumptionType,
) PointConsumptionTypeReq {
	req := PointConsumptionTypeReq{
		PointConsumptionTypes: input,
		Accepter: []string{
			"PointConsumptionTypesByPointConsumptionTypes",
		},
	}
	return req
}

func CreatePointConsumptionTypeRequestPointConsumptionTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointConsumptionType,
) PointConsumptionTypeReq {
	isMarkedForDeletion := false

	req := PointConsumptionTypeReq{
		PointConsumptionTypes: []PointConsumptionType{
			{
				PointConsumptionType: input.PointConsumptionType,
				IsMarkedForDeletion:  &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"PointConsumptionTypes",
		},
	}
	return req
}

func CreatePointConsumptionTypeRequestText(
	requestPram *apiInputReader.Request,
	input PointConsumptionType,
) PointConsumptionTypeReq {
	isMarkedForDeletion := false

	req := PointConsumptionTypeReq{
		PointConsumptionType: PointConsumptionType{
			PointConsumptionType: input.PointConsumptionType,
			IsMarkedForDeletion:  &isMarkedForDeletion,
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

func CreatePointConsumptionTypeRequestTexts(
	requestPram *apiInputReader.Request,
	input PointConsumptionType,
) PointConsumptionTypeReq {
	isMarkedForDeletion := false

	req := PointConsumptionTypeReq{
		PointConsumptionType: PointConsumptionType{
			PointConsumptionType: input.PointConsumptionType,
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

func PointConsumptionTypeReadsPointConsumptionTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointConsumptionType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_CONSUMPTION_TYPE_SRV"
	aPIType := "reads"

	var request PointConsumptionTypeReq

	request = CreatePointConsumptionTypeRequestPointConsumptionTypes(
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

func PointConsumptionTypeReadsPointConsumptionTypesByPointConsumptionTypes(
	requestPram *apiInputReader.Request,
	input []PointConsumptionType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_CONSUMPTION_TYPE_SRV"
	aPIType := "reads"

	var request PointConsumptionTypeReq

	request = CreatePointConsumptionTypeRequestPointConsumptionTypesByPointConsumptionTypes(
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

func PointConsumptionTypeReadsText(
	requestPram *apiInputReader.Request,
	input PointConsumptionType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_CONSUMPTION_TYPE_SRV"
	aPIType := "reads"

	var request PointConsumptionTypeReq

	request = CreatePointConsumptionTypeRequestText(
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

func PointConsumptionTypeReadsTexts(
	requestPram *apiInputReader.Request,
	input PointConsumptionType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_CONSUMPTION_TYPE_SRV"
	aPIType := "reads"

	var request PointConsumptionTypeReq

	request = CreatePointConsumptionTypeRequestTexts(
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
