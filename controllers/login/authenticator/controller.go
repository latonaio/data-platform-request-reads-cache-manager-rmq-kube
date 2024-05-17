package controllersLoginAuthenticator

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

type LoginAuthenticatorController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type LoginAuthenticator struct {
	AuthenticatorUser []apiOutputFormatter.AuthenticatorUser `json:"AuthenticatorUser"`
}

func (controller *LoginAuthenticatorController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	userID := controller.GetString("userID")

	isMarkedForDeletion := false

	AuthenticatorUser := apiInputReader.Authenticator{}

	AuthenticatorUser = apiInputReader.Authenticator{
		AuthenticatorUser: &apiInputReader.AuthenticatorUser{
			UserID:              userID,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.request(
		AuthenticatorUser,
	)
}

func (
	controller *LoginAuthenticatorController,
) createAuthenticatorRequestUser(
	requestPram *apiInputReader.Request,
	inputAuthenticatorUser apiInputReader.Authenticator,
) *apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes {
	responseJsonData := apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes{}
	responseBody := apiModuleRuntimesRequestsAuthenticator.AuthenticatorReadsUser(
		requestPram,
		inputAuthenticatorUser,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createAuthenticatorRequestUser Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *LoginAuthenticatorController,
) request(
	inputAuthenticatorUser apiInputReader.Authenticator,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	authenticatorUserRes := controller.createAuthenticatorRequestUser(
		controller.UserInfo,
		inputAuthenticatorUser,
	)

	controller.fin(
		authenticatorUserRes,
	)
}

func (
	controller *LoginAuthenticatorController,
) fin(
	authenticatorUserRes *apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes,
) {

	data := LoginAuthenticator{}

	for _, v := range *authenticatorUserRes.Message.User {
		data.AuthenticatorUser = append(data.AuthenticatorUser,
			apiOutputFormatter.AuthenticatorUser{
				UserID:   v.UserID,
				Password: v.Password,
			},
		)
	}

	controller.Data["json"] = &data
	controller.ServeJSON()
}
