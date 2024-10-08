package apiModuleRuntimesRequestsBusinessPartnerRole

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type BusinessPartnerRoleReq struct {
	BusinessPartnerRole  BusinessPartnerRole   `json:"BusinessPartnerRole"`
	BusinessPartnerRoles []BusinessPartnerRole `json:"BusinessPartnerRoles"`
	Accepter             []string              `json:"accepter"`
}

type BusinessPartnerRole struct {
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
	Text                []Text `json:"Text"`
}

type Text struct {
	BusinessPartnerRole     string `json:"BusinessPartnerRole"`
	Language                string `json:"Language"`
	BusinessPartnerRoleName string `json:"BusinessPartnerRoleName"`
	CreationDate            string `json:"CreationDate"`
	LastChangeDate          string `json:"LastChangeDate"`
	IsMarkedForDeletion     *bool  `json:"IsMarkedForDeletion"`
}

func CreateBusinessPartnerRoleRequestBusinessPartnerRolesByBusinessPartnerRoles(
	requestPram *apiInputReader.Request,
	input []BusinessPartnerRole,
) BusinessPartnerRoleReq {
	req := BusinessPartnerRoleReq{
		BusinessPartnerRoles: input,
		Accepter: []string{
			"BusinessPartnerRolesByBusinessPartnerRoles",
		},
	}
	return req
}

func CreateBusinessPartnerRoleRequestBusinessPartnerRoles(
	requestPram *apiInputReader.Request,
	input apiInputReader.BusinessPartnerRole,
) BusinessPartnerRoleReq {
	isMarkedForDeletion := false

	req := BusinessPartnerRoleReq{
		BusinessPartnerRoles: []BusinessPartnerRole{
			{
				BusinessPartnerRole: input.BusinessPartnerRole,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		},
		Accepter: []string{
			"BusinessPartnerRoles",
		},
	}
	return req
}

func CreateBusinessPartnerRoleRequestText(
	requestPram *apiInputReader.Request,
	input BusinessPartnerRole,
) BusinessPartnerRoleReq {
	isMarkedForDeletion := false

	req := BusinessPartnerRoleReq{
		BusinessPartnerRole: BusinessPartnerRole{
			BusinessPartnerRole: input.BusinessPartnerRole,
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

func CreateBusinessPartnerRoleRequestTexts(
	requestPram *apiInputReader.Request,
	input BusinessPartnerRole,
) BusinessPartnerRoleReq {
	isMarkedForDeletion := false

	req := BusinessPartnerRoleReq{
		BusinessPartnerRole: BusinessPartnerRole{
			BusinessPartnerRole: input.BusinessPartnerRole,
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

func BusinessPartnerRoleReadsBusinessPartnerRoles(
	requestPram *apiInputReader.Request,
	input apiInputReader.BusinessPartnerRole,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_ROLE_SRV"
	aPIType := "reads"

	var request BusinessPartnerRoleReq

	request = CreateBusinessPartnerRoleRequestBusinessPartnerRoles(
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

func BusinessPartnerRoleReadsBusinessPartnerRolesByBusinessPartnerRoles(
	requestPram *apiInputReader.Request,
	input []BusinessPartnerRole,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_ROLE_SRV"
	aPIType := "reads"

	var request BusinessPartnerRoleReq

	request = CreateBusinessPartnerRoleRequestBusinessPartnerRolesByBusinessPartnerRoles(
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

func BusinessPartnerRoleReadsText(
	requestPram *apiInputReader.Request,
	input BusinessPartnerRole,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_ROLE_SRV"
	aPIType := "reads"

	var request BusinessPartnerRoleReq

	request = CreateBusinessPartnerRoleRequestText(
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

func BusinessPartnerRoleReadsTexts(
	requestPram *apiInputReader.Request,
	input BusinessPartnerRole,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_ROLE_SRV"
	aPIType := "reads"

	var request BusinessPartnerRoleReq

	request = CreateBusinessPartnerRoleRequestTexts(
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
