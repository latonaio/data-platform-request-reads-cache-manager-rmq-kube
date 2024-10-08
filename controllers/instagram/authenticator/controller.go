package controllersInstagramAuthenticator

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsAuthenticator "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/authenticator"
	apiModuleRuntimesResponsesAuthenticator "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/authenticator"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type InstagramAuthenticatorController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type InstagramAuthenticator struct {
	AuthenticatorInstagramAuth []apiOutputFormatter.AuthenticatorInstagramAuth `json:"AuthenticatorInstagramAuth"`
}

func (controller *InstagramAuthenticatorController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	userID := controller.GetString("userID")

	AuthenticatorInstagramAuth := apiInputReader.Authenticator{}

	AuthenticatorInstagramAuth = apiInputReader.Authenticator{
		AuthenticatorUser: &apiInputReader.AuthenticatorUser{
			UserID: userID,
		},
	}

	controller.request(
		AuthenticatorInstagramAuth,
	)
}

func (
controller *InstagramAuthenticatorController,
) createAuthenticatorRequestInstagramAuth(
	requestPram *apiInputReader.Request,
	inputAuthenticatorInstagramAuth apiInputReader.Authenticator,
) *apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes {
	responseJsonData := apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes{}
	responseBody := apiModuleRuntimesRequestsAuthenticator.AuthenticatorReadsInstagramAuth(
		requestPram,
		inputAuthenticatorInstagramAuth,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createAuthenticatorRequestInstagramAuth Unmarshal error")
	}

	return &responseJsonData
}

func (
controller *InstagramAuthenticatorController,
) request(
	inputAuthenticatorInstagramAuth apiInputReader.Authenticator,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	authenticatorInstagramAuthRes := controller.createAuthenticatorRequestInstagramAuth(
		controller.UserInfo,
		inputAuthenticatorInstagramAuth,
	)

	controller.fin(
		authenticatorInstagramAuthRes,
	)
}

func (
controller *InstagramAuthenticatorController,
) fin(
	authenticatorInstagramAuthRes *apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes,
) {

	data := InstagramAuthenticator{}

	for _, v := range *authenticatorInstagramAuthRes.Message.InstagramAuth {
		data.AuthenticatorInstagramAuth = append(data.AuthenticatorInstagramAuth,
			apiOutputFormatter.AuthenticatorInstagramAuth{
				UserID:      v.UserID,
				InstagramID: v.InstagramID,
				AccessToken: v.AccessToken,
			},
		)
	}

	controller.Data["json"] = &data
	controller.ServeJSON()
}
