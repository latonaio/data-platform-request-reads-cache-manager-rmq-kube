package apiModuleRuntimesRequestsLocalSubRegion

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type LocalSubRegionReq struct {
	LocalSubRegion  LocalSubRegion   `json:"LocalSubRegion"`
	LocalSubRegions []LocalSubRegion `json:"LocalSubRegions"`
	Accepter        []string         `json:"accepter"`
}

type LocalSubRegion struct {
	LocalSubRegion      string  `json:"LocalSubRegion"`
	LocalRegion         string  `json:"LocalRegion"`
	Country             string  `json:"Country"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Text                []Text  `json:"Text"`
}

type Text struct {
	LocalSubRegion      string  `json:"LocalSubRegion"`
	LocalRegion         string  `json:"LocalRegion"`
	Country             string  `json:"Country"`
	Language            string  `json:"Language"`
	LocalSubRegionName  *string `json:"LocalSubRegionName"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func CreateLocalSubRegionRequestLocalSubRegionsByLocalSubRegions(
	requestPram *apiInputReader.Request,
	input []LocalSubRegion,
) LocalSubRegionReq {
	req := LocalSubRegionReq{
		LocalSubRegions: input,
		Accepter: []string{
			"LocalSubRegionsByLocalSubRegions",
		},
	}
	return req
}

func CreateLocalSubRegionRequestLocalSubRegions(
	requestPram *apiInputReader.Request,
	input apiInputReader.LocalSubRegion,
) LocalSubRegionReq {
	isMarkedForDeletion := false

	req := LocalSubRegionReq{
		LocalSubRegions: []LocalSubRegion{
			{
				LocalSubRegion:      input.LocalSubRegion,
				LocalRegion:         input.LocalRegion,
				Country:             input.Country,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"LocalSubRegions",
		},
	}
	return req
}

func CreateLocalSubRegionRequestText(
	requestPram *apiInputReader.Request,
	input LocalSubRegion,
) LocalSubRegionReq {
	isMarkedForDeletion := false

	req := LocalSubRegionReq{
		LocalSubRegion: LocalSubRegion{
			LocalSubRegion:      input.LocalSubRegion,
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

func CreateLocalSubRegionRequestTexts(
	requestPram *apiInputReader.Request,
	input LocalSubRegion,
) LocalSubRegionReq {
	isMarkedForDeletion := false

	req := LocalSubRegionReq{
		LocalSubRegion: LocalSubRegion{
			LocalSubRegion: input.LocalSubRegion,
			LocalRegion:    input.LocalRegion,
			Country:        input.Country,
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

func LocalSubRegionReadsLocalSubRegions(
	requestPram *apiInputReader.Request,
	input apiInputReader.LocalSubRegion,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_LOCAL_SUB_REGION_SRV"
	aPIType := "reads"

	var request LocalSubRegionReq

	request = CreateLocalSubRegionRequestLocalSubRegions(
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

func LocalSubRegionReadsLocalSubRegionsByLocalSubRegions(
	requestPram *apiInputReader.Request,
	input []LocalSubRegion,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_LOCAL_SUB_REGION_SRV"
	aPIType := "reads"

	var request LocalSubRegionReq

	request = CreateLocalSubRegionRequestLocalSubRegionsByLocalSubRegions(
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

func LocalSubRegionReadsText(
	requestPram *apiInputReader.Request,
	input LocalSubRegion,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_LOCAL_SUB_REGION_SRV"
	aPIType := "reads"

	var request LocalSubRegionReq

	request = CreateLocalSubRegionRequestText(
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

func LocalSubRegionReadsTexts(
	requestPram *apiInputReader.Request,
	input LocalSubRegion,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_LOCAL_SUB_REGION_SRV"
	aPIType := "reads"

	var request LocalSubRegionReq

	request = CreateLocalSubRegionRequestTexts(
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
