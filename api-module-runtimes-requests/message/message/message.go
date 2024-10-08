package apiModuleRuntimesRequestsMessage

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type MessageReq struct {
	Header   Header   `json:"Message"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Message				int		`json:"Message"`
	MessageType			*string	`json:"MessageType"`
	Sender				*int	`json:"Sender"`
	Receiver			*int	`json:"Receiver"`
	Language			*string	`json:"Language"`
	Title				*string	`json:"Title"`
	LongText			*string	`json:"LongText"`
	MessageIsSent		*bool	`json:"MessageIsSent"`
	MessageIsRead		*bool	`json:"MessageIsRead"`
	CreationDate		*string	`json:"CreationDate"`
	CreationTime		*string	`json:"CreationTime"`
	LastChangeDate		*string	`json:"LastChangeDate"`
	LastChangeTime		*string	`json:"LastChangeTime"`
	IsCancelled			*bool	`json:"IsCancelled"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

func CreateMessageRequestHeader(
	requestPram *apiInputReader.Request,
	messageHeader *apiInputReader.MessageHeader,
) MessageReq {
	req := MessageReq{
		Header: Header{
			Message:             messageHeader.Message,
			MessageIsSent:       messageHeader.MessageIsSent,
			IsMarkedForDeletion: messageHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreateMessageRequestHeadersByReceiver(
	requestPram *apiInputReader.Request,
	messageHeader *apiInputReader.MessageHeader,
) MessageReq {
	req := MessageReq{
		Header: Header{
			Receiver:            messageHeader.Receiver,
			MessageIsSent:       messageHeader.MessageIsSent,
			IsMarkedForDeletion: messageHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersByReceiver",
		},
	}
	return req
}

func CreateMessageRequestHeadersBySender(
	requestPram *apiInputReader.Request,
	messageHeader *apiInputReader.MessageHeader,
) MessageReq {
	req := MessageReq{
		Header: Header{
			Sender:              messageHeader.Sender,
			MessageIsSent:       messageHeader.MessageIsSent,
			IsMarkedForDeletion: messageHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersBySender",
		},
	}
	return req
}

func MessageReadsHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.Message,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_MESSAGE_SRV"
	aPIType := "reads"

	var request MessageReq

	request = CreateMessageRequestHeader(
		requestPram,
		&apiInputReader.MessageHeader{
			Message:             input.MessageHeader.Message,
			MessageIsSent:       input.MessageHeader.MessageIsSent,
			IsMarkedForDeletion: input.MessageHeader.IsMarkedForDeletion,
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

func MessageReadsHeadersByReceiver(
	requestPram *apiInputReader.Request,
	input apiInputReader.Message,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_MESSAGE_SRV"
	aPIType := "reads"

	var request MessageReq

	request = CreateMessageRequestHeadersByReceiver(
		requestPram,
		&apiInputReader.MessageHeader{
			Receiver:            input.MessageHeader.Receiver,
			MessageIsSent:       input.MessageHeader.MessageIsSent,
			IsMarkedForDeletion: input.MessageHeader.IsMarkedForDeletion,
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

func MessageReadsHeadersBySender(
	requestPram *apiInputReader.Request,
	input apiInputReader.Message,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_MESSAGE_SRV"
	aPIType := "reads"

	var request MessageReq

	request = CreateMessageRequestHeadersBySender(
		requestPram,
		&apiInputReader.MessageHeader{
			Sender:              input.MessageHeader.Sender,
			MessageIsSent:       input.MessageHeader.MessageIsSent,
			IsMarkedForDeletion: input.MessageHeader.IsMarkedForDeletion,
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
