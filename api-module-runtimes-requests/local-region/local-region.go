package apiModuleRuntimesRequestsLocalRegion

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type LocalRegionReq struct {
	LocalRegion  LocalRegion   `json:"LocalRegion"`
	LocalRegions []LocalRegion `json:"LocalRegions"`
	Accepter     []string      `json:"accepter"`
}

type LocalRegion struct {
	LocalRegion         string  `json:"LocalRegion"`
	Country             string  `json:"Country"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Text                []Text  `json:"Text"`
}

type Text struct {
	LocalRegion         string  `json:"LocalRegion"`
	Country             string  `json:"Country"`
	Language            string  `json:"Language"`
	LocalRegionName     *string `json:"LocalRegionName"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func CreateLocalRegionRequestLocalRegionsByLocalRegions(
	requestPram *apiInputReader.Request,
	input []LocalRegion,
) LocalRegionReq {
	req := LocalRegionReq{
		LocalRegions: input,
		Accepter: []string{
			"LocalRegionsByLocalRegions",
		},
	}
	return req
}

func CreateLocalRegionRequestLocalRegions(
	requestPram *apiInputReader.Request,
	input apiInputReader.LocalRegion,
) LocalRegionReq {
	isMarkedForDeletion := false

	req := LocalRegionReq{
		LocalRegions: []LocalRegion{
			{
				LocalRegion:         input.LocalRegion,
				Country:             input.Country,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"LocalRegions",
		},
	}
	return req
}

func CreateLocalRegionRequestText(
	requestPram *apiInputReader.Request,
	input LocalRegion,
) LocalRegionReq {
	isMarkedForDeletion := false

	req := LocalRegionReq{
		LocalRegion: LocalRegion{
			LocalRegion:         input.LocalRegion,
			Country:             input.Country,
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

func CreateLocalRegionRequestTexts(
	requestPram *apiInputReader.Request,
	input LocalRegion,
) LocalRegionReq {
	isMarkedForDeletion := false

	req := LocalRegionReq{
		LocalRegion: LocalRegion{
			LocalRegion: input.LocalRegion,
			Country:     input.Country,
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

func LocalRegionReadsLocalRegions(
	requestPram *apiInputReader.Request,
	input apiInputReader.LocalRegion,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_LOCAL_REGION_SRV"
	aPIType := "reads"

	var request LocalRegionReq

	request = CreateLocalRegionRequestLocalRegions(
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

func LocalRegionReadsLocalRegionsByLocalRegions(
	requestPram *apiInputReader.Request,
	input []LocalRegion,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_LOCAL_REGION_SRV"
	aPIType := "reads"

	var request LocalRegionReq

	request = CreateLocalRegionRequestLocalRegionsByLocalRegions(
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

func LocalRegionReadsText(
	requestPram *apiInputReader.Request,
	input LocalRegion,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_LOCAL_REGION_SRV"
	aPIType := "reads"

	var request LocalRegionReq

	request = CreateLocalRegionRequestText(
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

func LocalRegionReadsTexts(
	requestPram *apiInputReader.Request,
	input LocalRegion,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_LOCAL_REGION_SRV"
	aPIType := "reads"

	var request LocalRegionReq

	request = CreateLocalRegionRequestTexts(
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
