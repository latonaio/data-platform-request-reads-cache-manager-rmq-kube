package controllersLoginAuthenticatorInitialGoogleAccountAuth

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

type LoginAuthenticatorInitialGoogleAccountAuthController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type LoginAuthenticatorInitialGoogleAccountAuth struct {
	AuthenticatorInitialGoogleAccountAuth []apiOutputFormatter.AuthenticatorInitialGoogleAccountAuth `json:"AuthenticatorInitialGoogleAccountAuth"`
}

func (controller *LoginAuthenticatorInitialGoogleAccountAuthController) Get() {
	emailAddress := controller.GetString("emailAddress")

	isMarkedForDeletion := false

	AuthenticatorInitialGoogleAccountAuth := apiInputReader.Authenticator{}

	AuthenticatorInitialGoogleAccountAuth = apiInputReader.Authenticator{
		AuthenticatorInitialGoogleAccountAuth: &apiInputReader.AuthenticatorInitialGoogleAccountAuth{
			EmailAddress:        emailAddress,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.request(
		AuthenticatorInitialGoogleAccountAuth,
	)
}

func (
	controller *LoginAuthenticatorInitialGoogleAccountAuthController,
) createAuthenticatorRequestInitialGoogleAccountAuth(
	requestPram *apiInputReader.Request,
	inputAuthenticatorInitialGoogleAccountAuth apiInputReader.Authenticator,
) *apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes {
	responseJsonData := apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes{}
	responseBody := apiModuleRuntimesRequestsAuthenticator.AuthenticatorReadsInitialGoogleAccountAuth(
		requestPram,
		inputAuthenticatorInitialGoogleAccountAuth,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createAuthenticatorRequestInitialGoogleAccountAuth Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *LoginAuthenticatorInitialGoogleAccountAuthController,
) request(
	inputAuthenticatorInitialGoogleAccountAuth apiInputReader.Authenticator,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	authenticatorInitialGoogleAccountAuthRes := controller.createAuthenticatorRequestInitialGoogleAccountAuth(
		controller.UserInfo,
		inputAuthenticatorInitialGoogleAccountAuth,
	)

	controller.fin(
		authenticatorInitialGoogleAccountAuthRes,
	)
}

func (
	controller *LoginAuthenticatorInitialGoogleAccountAuthController,
) fin(
	authenticatorInitialGoogleAccountAuthRes *apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes,
) {

	data := LoginAuthenticatorInitialGoogleAccountAuth{}

	for _, v := range *authenticatorInitialGoogleAccountAuthRes.Message.InitialGoogleAccountAuth {
		data.AuthenticatorInitialGoogleAccountAuth = append(data.AuthenticatorInitialGoogleAccountAuth,
			apiOutputFormatter.AuthenticatorInitialGoogleAccountAuth{
				EmailAddress:    v.EmailAddress,
				GoogleID:        v.GoogleID,
                AccessToken:     v.AccessToken,
			},
		)
	}

	controller.Data["json"] = &data
	controller.ServeJSON()
}
