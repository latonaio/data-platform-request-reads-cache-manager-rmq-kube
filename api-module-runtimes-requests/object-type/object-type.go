package apiModuleRuntimesRequestsObjectType

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ObjectTypeReq struct {
	ObjectType  ObjectType   `json:"ObjectType"`
	ObjectTypes []ObjectType `json:"ObjectTypes"`
	Accepter    []string     `json:"accepter"`
}

type ObjectType struct {
	ObjectType          string  `json:"ObjectType"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Text                []Text  `json:"Text"`
}

type Text struct {
	ObjectType          string  `json:"ObjectType"`
	Language            string  `json:"Language"`
	ObjectTypeName      *string `json:"ObjectTypeName"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func CreateObjectTypeRequestObjectTypesByObjectTypes(
	requestPram *apiInputReader.Request,
	input []ObjectType,
) ObjectTypeReq {
	req := ObjectTypeReq{
		ObjectTypes: input,
		Accepter: []string{
			"ObjectTypesByObjectTypes",
		},
	}
	return req
}

func CreateObjectTypeRequestObjectTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.ObjectType,
) ObjectTypeReq {
	isMarkedForDeletion := false

	req := ObjectTypeReq{
		ObjectTypes: []ObjectType{
			{
				ObjectType:          input.ObjectType,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"ObjectTypes",
		},
	}
	return req
}

func CreateObjectTypeRequestText(
	requestPram *apiInputReader.Request,
	input ObjectType,
) ObjectTypeReq {
	isMarkedForDeletion := false

	req := ObjectTypeReq{
		ObjectType: ObjectType{
			ObjectType:          input.ObjectType,
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

func CreateObjectTypeRequestTexts(
	requestPram *apiInputReader.Request,
	input ObjectType,
) ObjectTypeReq {
	isMarkedForDeletion := false

	req := ObjectTypeReq{
		ObjectType: ObjectType{
			ObjectType: input.ObjectType,
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

func ObjectTypeReadsObjectTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.ObjectType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_OBJECT_TYPE_SRV"
	aPIType := "reads"

	var request ObjectTypeReq

	request = CreateObjectTypeRequestObjectTypes(
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

func ObjectTypeReadsObjectTypesByObjectTypes(
	requestPram *apiInputReader.Request,
	input []ObjectType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_OBJECT_TYPE_SRV"
	aPIType := "reads"

	var request ObjectTypeReq

	request = CreateObjectTypeRequestObjectTypesByObjectTypes(
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

func ObjectTypeReadsText(
	requestPram *apiInputReader.Request,
	input ObjectType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_OBJECT_TYPE_SRV"
	aPIType := "reads"

	var request ObjectTypeReq

	request = CreateObjectTypeRequestText(
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

func ObjectTypeReadsTexts(
	requestPram *apiInputReader.Request,
	input ObjectType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_OBJECT_TYPE_SRV"
	aPIType := "reads"

	var request ObjectTypeReq

	request = CreateObjectTypeRequestTexts(
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
