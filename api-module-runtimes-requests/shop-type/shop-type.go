package apiModuleRuntimesRequestsShopType

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ShopTypeReq struct {
	ShopType   ShopType    `json:"ShopType"`
	ShopTypes  []ShopType  `json:"ShopTypes"`
	Accepter   []string    `json:"accepter"`
}

type ShopType struct {
	ShopType            string  `json:"ShopType"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Text                []Text  `json:"Text"`
}

type Text struct {
	ShopType            string  `json:"ShopType"`
	Language            string  `json:"Language"`
	ShopTypeName        *string `json:"ShopTypeName"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func CreateShopTypeRequestShopTypesByShopTypes(
	requestPram *apiInputReader.Request,
	input []ShopType,
) ShopTypeReq {
	req := ShopTypeReq{
		ShopTypes: input,
		Accepter: []string{
			"ShopTypesByShopTypes",
		},
	}
	return req
}

func CreateShopTypeRequestShopTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.ShopType,
) ShopTypeReq {
	isMarkedForDeletion := false

	req := ShopTypeReq{
		ShopTypes: []ShopType{
			{
				ShopType:           input.ShopType,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"ShopTypes",
		},
	}
	return req
}

func CreateShopTypeRequestText(
	requestPram *apiInputReader.Request,
	input ShopType,
) ShopTypeReq {
	isMarkedForDeletion := false

	req := ShopTypeReq{
		ShopType: ShopType{
			ShopType:            input.ShopType,
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

func CreateShopTypeRequestTexts(
	requestPram *apiInputReader.Request,
	input ShopType,
) ShopTypeReq {
	isMarkedForDeletion := false

	req := ShopTypeReq{
		ShopType: ShopType{
			ShopType: input.ShopType,
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

func ShopTypeReadsShopTypes(
	requestPram *apiInputReader.Request,
	input apiInputReader.ShopType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SHOP_TYPE_SRV"
	aPIType := "reads"

	var request ShopTypeReq

	request = CreateShopTypeRequestShopTypes(
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

func ShopTypeReadsShopTypesByShopTypes(
	requestPram *apiInputReader.Request,
	input []ShopType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SHOP_TYPE_SRV"
	aPIType := "reads"

	var request ShopTypeReq

	request = CreateShopTypeRequestShopTypesByShopTypes(
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

func ShopTypeReadsText(
	requestPram *apiInputReader.Request,
	input ShopType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SHOP_TYPE_SRV"
	aPIType := "reads"

	var request ShopTypeReq

	request = CreateShopTypeRequestText(
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

func ShopTypeReadsTexts(
	requestPram *apiInputReader.Request,
	input ShopType,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SHOP_TYPE_SRV"
	aPIType := "reads"

	var request ShopTypeReq

	request = CreateShopTypeRequestTexts(
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
