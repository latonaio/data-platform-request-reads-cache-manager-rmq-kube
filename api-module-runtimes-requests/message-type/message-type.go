package apiModuleRuntimesRequestsMessageType

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type MessageTypeReq struct {
	MessageType     MessageType    `json:"MessageType"`
	MessageTypes    []MessageType  `json:"MessageTypes"`
	Accepter        []string       `json:"accepter"`
}

type MessageType struct {
	MessageType             string  `json:"MessageType"`
	CreationDate            *string `json:"CreationDate"`
	LastChangeDate          *string `json:"LastChangeDate"`
	IsMarkedForDeletion     *bool   `json:"IsMarkedForDeletion"`
	Text                    []Text  `json:"Text"`
}

type Text struct {
	MessageType             string  `json:"MessageType"`
	Language                string  `json:"Language"`
	MessageTypeName         *string `json:"MessageTypeName"`
	CreationDate            *string `json:"CreationDate"`
	LastChangeDate          *string `json:"LastChangeDate"`
	IsMarkedForDeletion     *bool   `json:"IsMarkedForDeletion"`
}

func CreateMessageTypeRequestMessageTypesByMessageTypes(
	requestPram *apiInputReader.Request,
	input []MessageType,
) MessageTypeReq {
	req := MessageTypeReq{
		MessageTypes: input,
		Accepter: []string{
			"MessageTypesByMessageTypes",
		},
	}
	return req
}

func CreateMessageTypeRequestMessageTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.MessageType,
) MessageTypeReq {
	isMarkedForDeletion := false

	req := MessageTypeReq{
		MessageTypes: []MessageType{
			{
				MessageType:         input.MessageType,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"MessageTypes",
		},
	}
	return req
}

func CreateMessageTypeRequestText(
	requestPram *apiInputReader.Request,
	input MessageType,
) MessageTypeReq {
	isMarkedForDeletion := false

	req := MessageTypeReq{
		MessageType: MessageType{
			MessageType:            input.MessageType,
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

func CreateMessageTypeRequestTexts(
	requestPram *apiInputReader.Request,
	input MessageType,
) MessageTypeReq {
	isMarkedForDeletion := false

	req := MessageTypeReq{
		MessageType:        MessageType{
			MessageType:    input.MessageType,
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

func MessageTypeReadsMessageTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.MessageType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_MESSAGE_TYPE_SRV"
	aPIType := "reads"

	var request MessageTypeReq

	request = CreateMessageTypeRequestMessageTypes(
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

func MessageTypeReadsMessageTypesByMessageTypes(
	requestPram *apiInputReader.Request,
	input []MessageType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_MESSAGE_TYPE_SRV"
	aPIType := "reads"

	var request MessageTypeReq

	request = CreateMessageTypeRequestMessageTypesByMessageTypes(
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

func MessageTypeReadsText(
	requestPram *apiInputReader.Request,
	input MessageType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_MESSAGE_TYPE_SRV"
	aPIType := "reads"

	var request MessageTypeReq

	request = CreateMessageTypeRequestText(
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

func MessageTypeReadsTexts(
	requestPram *apiInputReader.Request,
	input MessageType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_MESSAGE_TYPE_SRV"
	aPIType := "reads"

	var request MessageTypeReq

	request = CreateMessageTypeRequestTexts(
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
