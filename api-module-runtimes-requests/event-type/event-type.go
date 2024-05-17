package apiModuleRuntimesRequestsEventType

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type EventTypeReq struct {
	EventType   EventType    `json:"EventType"`
	EventTypes  []EventType  `json:"EventTypes"`
	Accepter    []string     `json:"accepter"`
}

type EventType struct {
	EventType           string  `json:"EventType"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Text                []Text  `json:"Text"`
}

type Text struct {
	EventType           string  `json:"EventType"`
	Language            string  `json:"Language"`
	EventTypeName       *string `json:"EventTypeName"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func CreateEventTypeRequestEventTypesByEventTypes(
	requestPram *apiInputReader.Request,
	input []EventType,
) EventTypeReq {
	req := EventTypeReq{
		EventTypes: input,
		Accepter: []string{
			"EventTypesByEventTypes",
		},
	}
	return req
}

func CreateEventTypeRequestEventTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.EventType,
) EventTypeReq {
	isMarkedForDeletion := false

	req := EventTypeReq{
		EventTypes: []EventType{
			{
				EventType:           input.EventType,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"EventTypes",
		},
	}
	return req
}

func CreateEventTypeRequestText(
	requestPram *apiInputReader.Request,
	input EventType,
) EventTypeReq {
	isMarkedForDeletion := false

	req := EventTypeReq{
		EventType: EventType{
			EventType:           input.EventType,
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

func CreateEventTypeRequestTexts(
	requestPram *apiInputReader.Request,
	input EventType,
) EventTypeReq {
	isMarkedForDeletion := false

	req := EventTypeReq{
		EventType: EventType{
			EventType: input.EventType,
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

func EventTypeReadsEventTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.EventType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_TYPE_SRV"
	aPIType := "reads"

	var request EventTypeReq

	request = CreateEventTypeRequestEventTypes(
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

func EventTypeReadsEventTypesByEventTypes(
	requestPram *apiInputReader.Request,
	input []EventType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_TYPE_SRV"
	aPIType := "reads"

	var request EventTypeReq

	request = CreateEventTypeRequestEventTypesByEventTypes(
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

func EventTypeReadsText(
	requestPram *apiInputReader.Request,
	input EventType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_TYPE_SRV"
	aPIType := "reads"

	var request EventTypeReq

	request = CreateEventTypeRequestText(
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

func EventTypeReadsTexts(
	requestPram *apiInputReader.Request,
	input EventType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_TYPE_SRV"
	aPIType := "reads"

	var request EventTypeReq

	request = CreateEventTypeRequestTexts(
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
