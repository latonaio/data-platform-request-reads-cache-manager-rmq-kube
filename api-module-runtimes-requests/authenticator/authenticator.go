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
	UserID              string  `json:"UserID"`
	BusinessPartner     *int    `json:"BusinessPartner"`
	Password            *string `json:"Password"`
	Qos                 *string `json:"Qos"`
	IsEncrypt           *bool   `json:"IsEncrypt"`
	Language            *string `json:"Language"`
	LastLoginDate       *string `json:"LastLoginDate"`
	LastLoginTime       *string `json:"LastLoginTime"`
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
	)

	return responseBody
}
