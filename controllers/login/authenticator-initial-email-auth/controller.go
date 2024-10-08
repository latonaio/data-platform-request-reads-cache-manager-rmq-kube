package controllersLoginAuthenticatorInitialEmailAuth

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

type LoginAuthenticatorInitialEmailAuthController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type LoginAuthenticatorInitialEmailAuth struct {
	AuthenticatorInitialEmailAuth []apiOutputFormatter.AuthenticatorInitialEmailAuth `json:"AuthenticatorInitialEmailAuth"`
}

func (controller *LoginAuthenticatorInitialEmailAuthController) Get() {
	emailAddress := controller.GetString("emailAddress")

	isMarkedForDeletion := false

	AuthenticatorInitialEmailAuth := apiInputReader.Authenticator{}

	AuthenticatorInitialEmailAuth = apiInputReader.Authenticator{
		AuthenticatorInitialEmailAuth: &apiInputReader.AuthenticatorInitialEmailAuth{
			EmailAddress:        emailAddress,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.request(
		AuthenticatorInitialEmailAuth,
	)
}

func (
	controller *LoginAuthenticatorInitialEmailAuthController,
) createAuthenticatorRequestInitialEmailAuth(
	requestPram *apiInputReader.Request,
	inputAuthenticatorInitialEmailAuth apiInputReader.Authenticator,
) *apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes {
	responseJsonData := apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes{}
	responseBody := apiModuleRuntimesRequestsAuthenticator.AuthenticatorReadsInitialEmailAuth(
		requestPram,
		inputAuthenticatorInitialEmailAuth,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createAuthenticatorRequestInitialEmailAuth Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *LoginAuthenticatorInitialEmailAuthController,
) request(
	inputAuthenticatorInitialEmailAuth apiInputReader.Authenticator,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	authenticatorInitialEmailAuthRes := controller.createAuthenticatorRequestInitialEmailAuth(
		controller.UserInfo,
		inputAuthenticatorInitialEmailAuth,
	)

	controller.fin(
		authenticatorInitialEmailAuthRes,
	)
}

func (
	controller *LoginAuthenticatorInitialEmailAuthController,
) fin(
	authenticatorInitialEmailAuthRes *apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes,
) {

	data := LoginAuthenticatorInitialEmailAuth{}

	for _, v := range *authenticatorInitialEmailAuthRes.Message.InitialEmailAuth {
		data.AuthenticatorInitialEmailAuth = append(data.AuthenticatorInitialEmailAuth,
			apiOutputFormatter.AuthenticatorInitialEmailAuth{
				EmailAddress:        v.EmailAddress,
			},
		)
	}

	controller.Data["json"] = &data
	controller.ServeJSON()
}
