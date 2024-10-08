package apiModuleRuntimesRequestsPointTransactionType

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type PointTransactionTypeReq struct {
	PointTransactionType  PointTransactionType   `json:"PointTransactionType"`
	PointTransactionTypes []PointTransactionType `json:"PointTransactionTypes"`
	Accepter              []string               `json:"accepter"`
}

type PointTransactionType struct {
	PointTransactionType string  `json:"PointTransactionType"`
	CreationDate         *string `json:"CreationDate"`
	LastChangeDate       *string `json:"LastChangeDate"`
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
	Text                 []Text  `json:"Text"`
}

type Text struct {
	PointTransactionType     string  `json:"PointTransactionType"`
	Language                 string  `json:"Language"`
	PointTransactionTypeName *string `json:"PointTransactionTypeName"`
	CreationDate             *string `json:"CreationDate"`
	LastChangeDate           *string `json:"LastChangeDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

func CreatePointTransactionTypeRequestPointTransactionTypesByPointTransactionTypes(
	requestPram *apiInputReader.Request,
	input []PointTransactionType,
) PointTransactionTypeReq {
	req := PointTransactionTypeReq{
		PointTransactionTypes: input,
		Accepter: []string{
			"PointTransactionTypesByPointTransactionTypes",
		},
	}
	return req
}

func CreatePointTransactionTypeRequestPointTransactionTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointTransactionType,
) PointTransactionTypeReq {
	isMarkedForDeletion := false

	req := PointTransactionTypeReq{
		PointTransactionTypes: []PointTransactionType{
			{
				PointTransactionType: input.PointTransactionType,
				IsMarkedForDeletion:  &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"PointTransactionTypes",
		},
	}
	return req
}

func CreatePointTransactionTypeRequestText(
	requestPram *apiInputReader.Request,
	input PointTransactionType,
) PointTransactionTypeReq {
	isMarkedForDeletion := false

	req := PointTransactionTypeReq{
		PointTransactionType: PointTransactionType{
			PointTransactionType: input.PointTransactionType,
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

func CreatePointTransactionTypeRequestTexts(
	requestPram *apiInputReader.Request,
	input PointTransactionType,
) PointTransactionTypeReq {
	isMarkedForDeletion := false

	req := PointTransactionTypeReq{
		PointTransactionType: PointTransactionType{
			PointTransactionType: input.PointTransactionType,
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

func PointTransactionTypeReadsPointTransactionTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointTransactionType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_TRANSACTION_TYPE_SRV"
	aPIType := "reads"

	var request PointTransactionTypeReq

	request = CreatePointTransactionTypeRequestPointTransactionTypes(
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

func PointTransactionTypeReadsPointTransactionTypesByPointTransactionTypes(
	requestPram *apiInputReader.Request,
	input []PointTransactionType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_TRANSACTION_TYPE_SRV"
	aPIType := "reads"

	var request PointTransactionTypeReq

	request = CreatePointTransactionTypeRequestPointTransactionTypesByPointTransactionTypes(
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

func PointTransactionTypeReadsText(
	requestPram *apiInputReader.Request,
	input PointTransactionType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_TRANSACTION_TYPE_SRV"
	aPIType := "reads"

	var request PointTransactionTypeReq

	request = CreatePointTransactionTypeRequestText(
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

func PointTransactionTypeReadsTexts(
	requestPram *apiInputReader.Request,
	input PointTransactionType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_TRANSACTION_TYPE_SRV"
	aPIType := "reads"

	var request PointTransactionTypeReq

	request = CreatePointTransactionTypeRequestTexts(
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
