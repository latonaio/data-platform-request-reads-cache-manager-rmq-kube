package apiModuleRuntimesRequestsCountry

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type CountryReq struct {
	Country   Country   `json:"Country"`
	Countries []Country `json:"Countries"`
	Accepter  []string  `json:"accepter"`
}

type Country struct {
	Country             string  `json:"Country"`
	GlobalRegion        *string `json:"GlobalRegion"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Text                []Text  `json:"Text"`
}

type Text struct {
	Country             string  `json:"Country"`
	Language            string  `json:"Language"`
	CountryName         *string `json:"CountryName"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func CreateCountryRequestCountriesByCountries(
	requestPram *apiInputReader.Request,
	input []Country,
) CountryReq {
	req := CountryReq{
		Countries: input,
		Accepter: []string{
			"CountriesByCountries",
		},
	}
	return req
}

func CreateCountryRequestCountries(
	requestPram *apiInputReader.Request,
	input apiInputReader.Country,
) CountryReq {
	isMarkedForDeletion := false

	req := CountryReq{
		Countries: []Country{
			{
				Country:             input.Country,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"Countries",
		},
	}
	return req
}

func CreateCountryRequestText(
	requestPram *apiInputReader.Request,
	input Country,
) CountryReq {
	isMarkedForDeletion := false

	req := CountryReq{
		Country: Country{
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

func CreateCountryRequestTexts(
	requestPram *apiInputReader.Request,
	input Country,
) CountryReq {
	isMarkedForDeletion := false

	req := CountryReq{
		Country: Country{
			Country: input.Country,
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

func CountryReadsCountries(
	requestPram *apiInputReader.Request,
	input apiInputReader.Country,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_COUNTRY_SRV"
	aPIType := "reads"

	var request CountryReq

	request = CreateCountryRequestCountries(
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

func CountryReadsCountriesByCountries(
	requestPram *apiInputReader.Request,
	input []Country,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_COUNTRY_SRV"
	aPIType := "reads"

	var request CountryReq

	request = CreateCountryRequestCountriesByCountries(
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

func CountryReadsText(
	requestPram *apiInputReader.Request,
	input Country,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_COUNTRY_SRV"
	aPIType := "reads"

	var request CountryReq

	request = CreateCountryRequestText(
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

func CountryReadsTexts(
	requestPram *apiInputReader.Request,
	input Country,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_COUNTRY_SRV"
	aPIType := "reads"

	var request CountryReq

	request = CreateCountryRequestTexts(
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
