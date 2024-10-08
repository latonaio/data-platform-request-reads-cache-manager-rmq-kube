package apiModuleRuntimesRequestsParticipation

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ParticipationReq struct {
	Header   Header   `json:"Participation"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Participation           int     `json:"Participation"`
	ParticipationDate       *string `json:"ParticipationDate"`
	ParticipationTime       *string `json:"ParticipationTime"`
	Participator            *int    `json:"Participator"`
	ParticipationObjectType *string `json:"ParticipationObjectType"`
	ParticipationObject     *int    `json:"ParticipationObject"`
	Attendance              *int    `json:"Attendance"`
	Invitation              *int    `json:"Invitation"`
	CreationDate            *string `json:"CreationDate"`
	CreationTime            *string `json:"CreationTime"`
	IsCancelled             *bool   `json:"IsCancelled"`
}

func CreateParticipationRequestHeader(
	requestPram *apiInputReader.Request,
	participationHeader *apiInputReader.ParticipationHeader,
) ParticipationReq {
	req := ParticipationReq{
		Header: Header{
			Participation: participationHeader.Participation,
			IsCancelled:   participationHeader.IsCancelled,
		},
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreateParticipationRequestHeadersByParticipator(
	requestPram *apiInputReader.Request,
	participationHeader *apiInputReader.ParticipationHeader,
) ParticipationReq {
	req := ParticipationReq{
		Header: Header{
			Participator: participationHeader.Participator,
			IsCancelled:  participationHeader.IsCancelled,
		},
		Accepter: []string{
			"HeadersByParticipator",
		},
	}
	return req
}

func CreateParticipationRequestHeadersByEvent(
	requestPram *apiInputReader.Request,
	participationHeader *apiInputReader.ParticipationHeader,
) ParticipationReq {
	participationObjectType := "EVENT"

	req := ParticipationReq{
		Header: Header{
			ParticipationObjectType: &participationObjectType,
			ParticipationObject:     participationHeader.ParticipationObject,
			IsCancelled:             participationHeader.IsCancelled,
		},
		Accepter: []string{
			"HeadersByEvent",
		},
	}
	return req
}

func ParticipationReadsHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.Participation,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PARTICIPATION_SRV"
	aPIType := "reads"

	var request ParticipationReq

	request = CreateParticipationRequestHeader(
		requestPram,
		&apiInputReader.ParticipationHeader{
			Participation: input.ParticipationHeader.Participation,
			IsCancelled:   input.ParticipationHeader.IsCancelled,
		},
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

func ParticipationReadsHeadersByParticipator(
	requestPram *apiInputReader.Request,
	input apiInputReader.Participation,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PARTICIPATION_SRV"
	aPIType := "reads"

	var request ParticipationReq

	request = CreateParticipationRequestHeadersByParticipator(
		requestPram,
		&apiInputReader.ParticipationHeader{
			Participator: input.ParticipationHeader.Participator,
			IsCancelled:  input.ParticipationHeader.IsCancelled,
		},
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

func ParticipationReadsHeadersByEvent(
	requestPram *apiInputReader.Request,
	input apiInputReader.Participation,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PARTICIPATION_SRV"
	aPIType := "reads"

	var request ParticipationReq

	participationObjectType := "EVENT"

	request = CreateParticipationRequestHeadersByEvent(
		requestPram,
		&apiInputReader.ParticipationHeader{
			ParticipationObjectType: &participationObjectType,
			ParticipationObject:     input.ParticipationHeader.ParticipationObject,
			IsCancelled:             input.ParticipationHeader.IsCancelled,
		},
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
