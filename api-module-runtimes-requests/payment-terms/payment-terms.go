package apiModuleRuntimesRequestsPaymentTerms

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type PaymentTermsReq struct {
	PaymentTerms   PaymentTerms   `json:"PaymentTerms"`
	PaymentTermses []PaymentTerms `json:"PaymentTermses"`
	Accepter       []string       `json:"accepter"`
}

type PaymentTerms struct {
	PaymentTerms        string `json:"PaymentTerms"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
	Text                []Text `json:"Text"`
}

type Text struct {
	PaymentTerms        string `json:"PaymentTerms"`
	Language            string `json:"Language"`
	PaymentTermsName    string `json:"PaymentTermsName"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

func CreatePaymentTermsRequestPaymentTermsesByPaymentTermses(
	requestPram *apiInputReader.Request,
	input []PaymentTerms,
) PaymentTermsReq {
	req := PaymentTermsReq{
		PaymentTermses: input,
		Accepter: []string{
			"PaymentTermsesByPaymentTermses",
		},
	}
	return req
}

func CreatePaymentTermsRequestPaymentTermses(
	requestPram *apiInputReader.Request,
	input apiInputReader.PaymentTerms,
) PaymentTermsReq {
	isMarkedForDeletion := false

	req := PaymentTermsReq{
		PaymentTermses: []PaymentTerms{
			{
				PaymentTerms:        input.PaymentTerms,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"PaymentTermses",
		},
	}
	return req
}

func CreatePaymentTermsRequestText(
	requestPram *apiInputReader.Request,
	input PaymentTerms,
) PaymentTermsReq {
	isMarkedForDeletion := false

	req := PaymentTermsReq{
		PaymentTerms: PaymentTerms{
			PaymentTerms:        input.PaymentTerms,
			IsMarkedForDeletion: &isMarkedForDeletion,
			Text: []Text{
				{
					Language:            "JA", // TODO 暫定で固定値を設定
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

func CreatePaymentTermsRequestTexts(
	requestPram *apiInputReader.Request,
	input PaymentTerms,
) PaymentTermsReq {
	isMarkedForDeletion := false

	req := PaymentTermsReq{
		PaymentTerms: PaymentTerms{
			PaymentTerms: input.PaymentTerms,
			Text: []Text{
				{
					Language:            "JA", // TODO 暫定で固定値を設定
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

func PaymentTermsReadsPaymentTermses(
	requestPram *apiInputReader.Request,
	input apiInputReader.PaymentTerms,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PAYMENT_TERMS_SRV"
	aPIType := "reads"

	var request PaymentTermsReq

	request = CreatePaymentTermsRequestPaymentTermses(
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

func PaymentTermsReadsPaymentTermsesByPaymentTermses(
	requestPram *apiInputReader.Request,
	input []PaymentTerms,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PAYMENT_TERMS_SRV"
	aPIType := "reads"

	var request PaymentTermsReq

	request = CreatePaymentTermsRequestPaymentTermsesByPaymentTermses(
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

func PaymentTermsReadsText(
	requestPram *apiInputReader.Request,
	input PaymentTerms,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PAYMENT_TERMS_SRV"
	aPIType := "reads"

	var request PaymentTermsReq

	request = CreatePaymentTermsRequestText(
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

func PaymentTermsReadsTexts(
	requestPram *apiInputReader.Request,
	input PaymentTerms,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PAYMENT_TERMS_SRV"
	aPIType := "reads"

	var request PaymentTermsReq

	request = CreatePaymentTermsRequestTexts(
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
