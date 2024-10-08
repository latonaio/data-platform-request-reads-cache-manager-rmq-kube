package controllersLoginAuthenticatorInitialSMSAuth

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

type LoginAuthenticatorInitialSMSAuthController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type LoginAuthenticatorInitialSMSAuth struct {
	AuthenticatorInitialSMSAuth []apiOutputFormatter.AuthenticatorInitialSMSAuth `json:"AuthenticatorInitialSMSAuth"`
}

func (controller *LoginAuthenticatorInitialSMSAuthController) Get() {
	mobilePhoneNumber := controller.GetString("mobilePhoneNumber")

	isMarkedForDeletion := false

	AuthenticatorInitialSMSAuth := apiInputReader.Authenticator{}

	AuthenticatorInitialSMSAuth = apiInputReader.Authenticator{
		AuthenticatorInitialSMSAuth: &apiInputReader.AuthenticatorInitialSMSAuth{
			MobilePhoneNumber:   mobilePhoneNumber,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.request(
		AuthenticatorInitialSMSAuth,
	)
}

func (
	controller *LoginAuthenticatorInitialSMSAuthController,
) createAuthenticatorRequestInitialSMSAuth(
	requestPram *apiInputReader.Request,
	inputAuthenticatorInitialSMSAuth apiInputReader.Authenticator,
) *apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes {
	responseJsonData := apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes{}
	responseBody := apiModuleRuntimesRequestsAuthenticator.AuthenticatorReadsInitialSMSAuth(
		requestPram,
		inputAuthenticatorInitialSMSAuth,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createAuthenticatorRequestInitialSMSAuth Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *LoginAuthenticatorInitialSMSAuthController,
) request(
	inputAuthenticatorInitialSMSAuth apiInputReader.Authenticator,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	authenticatorInitialSMSAuthRes := controller.createAuthenticatorRequestInitialSMSAuth(
		controller.UserInfo,
		inputAuthenticatorInitialSMSAuth,
	)

	controller.fin(
		authenticatorInitialSMSAuthRes,
	)
}

func (
	controller *LoginAuthenticatorInitialSMSAuthController,
) fin(
	authenticatorInitialSMSAuthRes *apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes,
) {

	data := LoginAuthenticatorInitialSMSAuth{}

	for _, v := range *authenticatorInitialSMSAuthRes.Message.InitialSMSAuth {
		data.AuthenticatorInitialSMSAuth = append(data.AuthenticatorInitialSMSAuth,
			apiOutputFormatter.AuthenticatorInitialSMSAuth{
				MobilePhoneNumber:  v.MobilePhoneNumber,
				AuthenticationCode: v.AuthenticationCode,
			},
		)
	}

	controller.Data["json"] = &data
	controller.ServeJSON()
}
