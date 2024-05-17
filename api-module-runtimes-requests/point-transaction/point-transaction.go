package apiModuleRuntimesRequestsPointTransaction

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type PointTransactionReq struct {
	Header   Header   `json:"PointTransaction"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	PointTransaction                      int      `json:"PointTransaction"`
	PointTransactionType                  *string  `json:"PointTransactionType"`
	PointTransactionDate                  *string  `json:"PointTransactionDate"`
	PointTransactionTime                  *string  `json:"PointTransactionTime"`
	SenderObjectType					  *string  `json:"SenderObjectType"`
	SenderObject						  *int	   `json:"SenderObject"`
	ReceiverObjectType					  *string  `json:"ReceiverObjectType"`
	ReceiverObject						  *int	   `json:"ReceiverObject"`
	PointSymbol                           *string  `json:"PointSymbol"`
	PlusMinus                             *string  `json:"PlusMinus"`
	PointTransactionAmount                *float32 `json:"PointTransactionAmount"`
	PointTransactionObjectType            *string  `json:"PointTransactionObjectType"`
	PointTransactionObject                *int     `json:"PointTransactionObject"`
	SenderPointBalanceBeforeTransaction   *float32 `json:"SenderPointBalanceBeforeTransaction"`
	SenderPointBalanceAfterTransaction    *float32 `json:"SenderPointBalanceAfterTransaction"`
	ReceiverPointBalanceBeforeTransaction *float32 `json:"ReceiverPointBalanceBeforeTransaction"`
	ReceiverPointBalanceAfterTransaction  *float32 `json:"ReceiverPointBalanceAfterTransaction"`
	Attendance							  *int	   `json:"Attendance"`
	Participation						  *int	   `json:"Participation"`
	CreationDate                          *string  `json:"CreationDate"`
	CreationTime                          *string  `json:"CreationTime"`
	IsCancelled                           *bool    `json:"IsCancelled"`
}

func CreatePointTransactionRequestHeader(
	requestPram *apiInputReader.Request,
	pointTransactionHeader *apiInputReader.PointTransactionHeader,
) PointTransactionReq {
	req := PointTransactionReq{
		Header: Header{
			PointTransaction: pointTransactionHeader.PointTransaction,
			IsCancelled:      pointTransactionHeader.IsCancelled,
		},
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreatePointTransactionRequestHeadersByReceiver(
	requestPram *apiInputReader.Request,
	pointTransactionHeader *apiInputReader.PointTransactionHeader,
) PointTransactionReq {
	req := PointTransactionReq{
		Header: Header{
			ReceiverObjectType:		pointTransactionHeader.ReceiverObjectType,
			ReceiverObject:			pointTransactionHeader.ReceiverObject,
			IsCancelled: 			pointTransactionHeader.IsCancelled,
		},
		Accepter: []string{
			"HeadersByReceiver",
		},
	}
	return req
}

func CreatePointTransactionRequestHeadersBySender(
	requestPram *apiInputReader.Request,
	pointTransactionHeader *apiInputReader.PointTransactionHeader,
) PointTransactionReq {
	req := PointTransactionReq{
		Header: Header{
			SenderObjectType:		pointTransactionHeader.SenderObjectType,
			SenderObject:			pointTransactionHeader.SenderObject,
			IsCancelled: 			pointTransactionHeader.IsCancelled,
		},
		Accepter: []string{
			"HeadersBySender",
		},
	}
	return req
}

func PointTransactionReadsHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointTransaction,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_TRANSACTION_SRV"
	aPIType := "reads"

	var request PointTransactionReq

	request = CreatePointTransactionRequestHeader(
		requestPram,
		&apiInputReader.PointTransactionHeader{
			PointTransaction: input.PointTransactionHeader.PointTransaction,
			IsCancelled:      input.PointTransactionHeader.IsCancelled,
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
	)

	return responseBody
}

func PointTransactionReadsHeadersByReceiver(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointTransaction,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_TRANSACTION_SRV"
	aPIType := "reads"

	var request PointTransactionReq

	request = CreatePointTransactionRequestHeadersByReceiver(
		requestPram,
		&apiInputReader.PointTransactionHeader{
			ReceiverObjectType:		input.PointTransactionHeader.ReceiverObjectType,
			ReceiverObject:			input.PointTransactionHeader.ReceiverObject,
			IsCancelled:			input.PointTransactionHeader.IsCancelled,
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
	)

	return responseBody
}

func PointTransactionReadsHeadersBySender(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointTransaction,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_TRANSACTION_SRV"
	aPIType := "reads"

	var request PointTransactionReq

	request = CreatePointTransactionRequestHeadersBySender(
		requestPram,
		&apiInputReader.PointTransactionHeader{
			SenderObjectType:		input.PointTransactionHeader.SenderObjectType,
			SenderObject:			input.PointTransactionHeader.SenderObject,
			IsCancelled:			input.PointTransactionHeader.IsCancelled,
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
	)

	return responseBody
}
