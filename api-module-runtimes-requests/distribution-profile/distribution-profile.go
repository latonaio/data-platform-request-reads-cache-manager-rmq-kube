package apiModuleRuntimesRequestsDistributionProfile

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type DistributionProfileReq struct {
	DistributionProfile  DistributionProfile   `json:"DistributionProfile"`
	DistributionProfiles []DistributionProfile `json:"DistributionProfiles"`
	Accepter             []string              `json:"accepter"`
}

type DistributionProfile struct {
	DistributionProfile string  `json:"DistributionProfile"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Text                []Text  `json:"Text"`
}

type Text struct {
	DistributionProfile     string  `json:"DistributionProfile"`
	Language                string  `json:"Language"`
	DistributionProfileName *string `json:"DistributionProfileName"`
	CreationDate            *string `json:"CreationDate"`
	LastChangeDate          *string `json:"LastChangeDate"`
	IsMarkedForDeletion     *bool   `json:"IsMarkedForDeletion"`
}

func CreateDistributionProfileRequestDistributionProfilesByDistributionProfiles(
	requestPram *apiInputReader.Request,
	input []DistributionProfile,
) DistributionProfileReq {
	req := DistributionProfileReq{
		DistributionProfiles: input,
		Accepter: []string{
			"DistributionProfilesByDistributionProfiles",
		},
	}
	return req
}

func CreateDistributionProfileRequestDistributionProfiles(
	requestPram *apiInputReader.Request,
	input apiInputReader.DistributionProfile,
) DistributionProfileReq {
	isMarkedForDeletion := false

	req := DistributionProfileReq{
		DistributionProfiles: []DistributionProfile{
			{
				DistributionProfile: input.DistributionProfile,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"DistributionProfiles",
		},
	}
	return req
}

func CreateDistributionProfileRequestText(
	requestPram *apiInputReader.Request,
	input DistributionProfile,
) DistributionProfileReq {
	isMarkedForDeletion := false

	req := DistributionProfileReq{
		DistributionProfile: DistributionProfile{
			DistributionProfile: input.DistributionProfile,
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

func CreateDistributionProfileRequestTexts(
	requestPram *apiInputReader.Request,
	input DistributionProfile,
) DistributionProfileReq {
	isMarkedForDeletion := false

	req := DistributionProfileReq{
		DistributionProfile: DistributionProfile{
			DistributionProfile: input.DistributionProfile,
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

func DistributionProfileReadsDistributionProfiles(
	requestPram *apiInputReader.Request,
	input apiInputReader.DistributionProfile,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_DISTRIBUTION_PROFILE_SRV"
	aPIType := "reads"

	var request DistributionProfileReq

	request = CreateDistributionProfileRequestDistributionProfiles(
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

func DistributionProfileReadsDistributionProfilesByDistributionProfiles(
	requestPram *apiInputReader.Request,
	input []DistributionProfile,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_DISTRIBUTION_PROFILE_SRV"
	aPIType := "reads"

	var request DistributionProfileReq

	request = CreateDistributionProfileRequestDistributionProfilesByDistributionProfiles(
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

func DistributionProfileReadsText(
	requestPram *apiInputReader.Request,
	input DistributionProfile,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_DISTRIBUTION_PROFILE_SRV"
	aPIType := "reads"

	var request DistributionProfileReq

	request = CreateDistributionProfileRequestText(
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

func DistributionProfileReadsTexts(
	requestPram *apiInputReader.Request,
	input DistributionProfile,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_DISTRIBUTION_PROFILE_SRV"
	aPIType := "reads"

	var request DistributionProfileReq

	request = CreateDistributionProfileRequestTexts(
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
