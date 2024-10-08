package apiModuleRuntimesRequestsAuthenticator

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type AuthenticatorReq struct {
	User     User     `json:"Authenticator"`
	Accepter []string `json:"accepter"`
}

type User struct {
	UserID                   string                     `json:"UserID"`
	BusinessPartner          *int                       `json:"BusinessPartner"`
	Password                 *string                    `json:"Password"`
	Qos                      *string                    `json:"Qos"`
	IsEncrypt                *bool                      `json:"IsEncrypt"`
	Language                 *string                    `json:"Language"`
	LastLoginDate            *string                    `json:"LastLoginDate"`
	LastLoginTime            *string                    `json:"LastLoginTime"`
	CreationDate             *string                    `json:"CreationDate"`
	CreationTime             *string                    `json:"CreationTime"`
	LastChangeDate           *string                    `json:"LastChangeDate"`
	LastChangeTime           *string                    `json:"LastChangeTime"`
	IsMarkedForDeletion      *bool                      `json:"IsMarkedForDeletion"`
	InitialEmailAuth         []InitialEmailAuth         `json:"InitialEmailAuth"`
	InitialSMSAuth           []InitialSMSAuth           `json:"InitialSMSAuth"`
	SMSAuth                  []SMSAuth                  `json:"SMSAuth"`
	InitialGoogleAccountAuth []InitialGoogleAccountAuth `json:"InitialGoogleAccountAuth"`
	GoogleAccountAuth        []GoogleAccountAuth        `json:"GoogleAccountAuth"`
	InstagramAuth            []InstagramAuth            `json:"InstagramAuth"`
}

type InitialEmailAuth struct {
	EmailAddress        string  `json:"EmailAddress"`
	CreationDate        *string `json:"CreationDate"`
	CreationTime        *string `json:"CreationTime"`
	LastChangeDate      *string `json:"LastChangeDate"`
	LastChangeTime      *string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type InitialSMSAuth struct {
	MobilePhoneNumber   string  `json:"MobilePhoneNumber"`
	AuthenticationCode  *int    `json:"AuthenticationCode"`
	CreationDate        *string `json:"CreationDate"`
	CreationTime        *string `json:"CreationTime"`
	LastChangeDate      *string `json:"LastChangeDate"`
	LastChangeTime      *string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type SMSAuth struct {
	UserID              string  `json:"UserID"`
	MobilePhoneNumber   *string `json:"MobilePhoneNumber"`
	AuthenticationCode  *int    `json:"AuthenticationCode"`
	CreationDate        *string `json:"CreationDate"`
	CreationTime        *string `json:"CreationTime"`
	LastChangeDate      *string `json:"LastChangeDate"`
	LastChangeTime      *string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type InitialGoogleAccountAuth struct {
	EmailAddress        string  `json:"EmailAddress"`
	GoogleID            *string `json:"GoogleID"`
	AccessToken         *string `json:"AccessToken"`
	CreationDate        *string `json:"CreationDate"`
	CreationTime        *string `json:"CreationTime"`
	LastChangeDate      *string `json:"LastChangeDate"`
	LastChangeTime      *string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type GoogleAccountAuth struct {
	UserID              string  `json:"UserID"`
	EmailAddress        *string `json:"EmailAddress"`
	GoogleID            *string `json:"GoogleID"`
	AccessToken         *string `json:"AccessToken"`
	CreationDate        *string `json:"CreationDate"`
	CreationTime        *string `json:"CreationTime"`
	LastChangeDate      *string `json:"LastChangeDate"`
	LastChangeTime      *string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type InstagramAuth struct {
	UserID              string  `json:"UserID"`
	InstagramID         *string `json:"InstagramID"`
	AccessToken         *string `json:"AccessToken"`
	CreationDate        *string `json:"CreationDate"`
	CreationTime        *string `json:"CreationTime"`
	LastChangeDate      *string `json:"LastChangeDate"`
	LastChangeTime      *string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func CreateAuthenticatorRequestUser(
	requestPram *apiInputReader.Request,
	input apiInputReader.Authenticator,
) AuthenticatorReq {
	isMarkedForDeletion := false

	req := AuthenticatorReq{
		User: User{
			UserID:              input.AuthenticatorUser.UserID,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		Accepter: []string{
			"User",
		},
	}
	return req
}

func AuthenticatorReadsUser(
	requestPram *apiInputReader.Request,
	input apiInputReader.Authenticator,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_AUTHENTICATOR_SRV"
	aPIType := "reads"

	var request AuthenticatorReq

	request = CreateAuthenticatorRequestUser(
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

func CreateAuthenticatorRequestInitialEmailAuth(
	requestPram *apiInputReader.Request,
	input apiInputReader.Authenticator,
) AuthenticatorReq {
	isMarkedForDeletion := false

	req := AuthenticatorReq{
		User: User{
			InitialEmailAuth: []InitialEmailAuth{
				{
					EmailAddress:        input.AuthenticatorInitialEmailAuth.EmailAddress,
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"InitialEmailAuth",
		},
	}
	return req
}

func AuthenticatorReadsInitialEmailAuth(
	requestPram *apiInputReader.Request,
	input apiInputReader.Authenticator,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_AUTHENTICATOR_SRV"
	aPIType := "reads"

	var request AuthenticatorReq

	request = CreateAuthenticatorRequestInitialEmailAuth(
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

func CreateAuthenticatorRequestInitialSMSAuth(
	requestPram *apiInputReader.Request,
	input apiInputReader.Authenticator,
) AuthenticatorReq {
	isMarkedForDeletion := false

	req := AuthenticatorReq{
		User: User{
			InitialSMSAuth: []InitialSMSAuth{
				{
					MobilePhoneNumber:   input.AuthenticatorInitialSMSAuth.MobilePhoneNumber,
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"InitialSMSAuth",
		},
	}
	return req
}

func AuthenticatorReadsInitialSMSAuth(
	requestPram *apiInputReader.Request,
	input apiInputReader.Authenticator,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_AUTHENTICATOR_SRV"
	aPIType := "reads"

	var request AuthenticatorReq

	request = CreateAuthenticatorRequestInitialSMSAuth(
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

func CreateAuthenticatorRequestInitialGoogleAccountAuth(
	requestPram *apiInputReader.Request,
	input apiInputReader.Authenticator,
) AuthenticatorReq {
	isMarkedForDeletion := false

	req := AuthenticatorReq{
		User: User{
			InitialGoogleAccountAuth: []InitialGoogleAccountAuth{
				{
					EmailAddress:        input.AuthenticatorInitialGoogleAccountAuth.EmailAddress,
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"InitialGoogleAccountAuth",
		},
	}
	return req
}

func AuthenticatorReadsInitialGoogleAccountAuth(
	requestPram *apiInputReader.Request,
	input apiInputReader.Authenticator,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_AUTHENTICATOR_SRV"
	aPIType := "reads"

	var request AuthenticatorReq

	request = CreateAuthenticatorRequestInitialGoogleAccountAuth(
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

func CreateAuthenticatorRequestInstagramAuth(
	requestPram *apiInputReader.Request,
	input apiInputReader.Authenticator,
) AuthenticatorReq {
	isMarkedForDeletion := false

	req := AuthenticatorReq{
		User: User{
			UserID: input.AuthenticatorUser.UserID,
			InstagramAuth: []InstagramAuth{
				{
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"InstagramAuth",
		},
	}
	return req
}

func AuthenticatorReadsInstagramAuth(
	requestPram *apiInputReader.Request,
	input apiInputReader.Authenticator,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_AUTHENTICATOR_SRV"
	aPIType := "reads"

	var request AuthenticatorReq

	request = CreateAuthenticatorRequestInstagramAuth(
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
