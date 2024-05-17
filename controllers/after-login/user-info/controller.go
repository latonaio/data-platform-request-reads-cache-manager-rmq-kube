package controllersAfterLoginUserInfo

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsAuthenticator "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/authenticator"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesResponsesAuthenticator "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/authenticator"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type AfterLoginUserInfoController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type AfterLoginUserInfo struct {
	AuthenticatorUser      []apiOutputFormatter.AuthenticatorUser      `json:"AuthenticatorUser"`
	BusinessPartnerAddress []apiOutputFormatter.BusinessPartnerAddress `json:"BusinessPartnerAddress"`
	BusinessPartnerPerson  []apiOutputFormatter.BusinessPartnerPerson  `json:"BusinessPartnerPerson"`
}

func (controller *AfterLoginUserInfoController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	userID := controller.GetString("userID")
	redisKeyCategory1 := "userID"

	isMarkedForDeletion := false

	AuthenticatorUser := apiInputReader.Authenticator{}
	//BusinessPartnerAddress := apiInputReader.BusinessPartner{}
	//BusinessPartnerPerson := apiInputReader.BusinessPartner{}

	AuthenticatorUser = apiInputReader.Authenticator{
		AuthenticatorUser: &apiInputReader.AuthenticatorUser{
			UserID:              userID,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	//BusinessPartnerAddress = apiInputReader.BusinessPartner{
	//	BusinessPartnerAddress: &apiInputReader.BusinessPartnerAddress{
	//		BusinessPartner: v.BusinessPartner,
	//	},
	//}
	//
	//BusinessPartnerPerson = apiInputReader.BusinessPartner{
	//	BusinessPartnerPerson: &apiInputReader.BusinessPartnerPerson{
	//		BusinessPartner:     v.BusinessPartner,
	//		IsMarkedForDeletion: &isMarkedForDeletion,
	//	},
	//}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData AfterLoginUserInfo

		err := json.Unmarshal(cacheData, &responseData)

		if err != nil {
			services.HandleError(
				&controller.Controller,
				err,
				nil,
			)
		}

		services.Respond(
			&controller.Controller,
			&responseData,
		)
	}

	if cacheData != nil {
		go func() {
			controller.request(
				AuthenticatorUser,
				//BusinessPartnerAddress,
				//BusinessPartnerPerson,
			)
		}()
	} else {
		controller.request(
			AuthenticatorUser,
			//BusinessPartnerAddress,
			//BusinessPartnerPerson,
		)
	}
}

func (
	controller *AfterLoginUserInfoController,
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
	controller *AfterLoginUserInfoController,
) createBusinessPartnerRequestAddress(
	requestPram *apiInputReader.Request,
	authenticatorRes apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input = apiModuleRuntimesRequestsBusinessPartner.Address{}

	for _, v := range *authenticatorRes.Message.User {
		input = apiModuleRuntimesRequestsBusinessPartner.Address{
			BusinessPartner: v.BusinessPartner,
		}
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsAddresses(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createBusinessPartnerRequestAddress Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AfterLoginUserInfoController,
) createBusinessPartnerRequestPerson(
	requestPram *apiInputReader.Request,
	authenticatorRes apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input = apiModuleRuntimesRequestsBusinessPartner.Person{}

	for _, v := range *authenticatorRes.Message.User {
		input = apiModuleRuntimesRequestsBusinessPartner.Person{
			BusinessPartner: v.BusinessPartner,
		}
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsPerson(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createBusinessPartnerRequestPerson Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AfterLoginUserInfoController,
) request(
	inputAuthenticatorUser apiInputReader.Authenticator,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	authenticatorUserRes := controller.createAuthenticatorRequestUser(
		controller.UserInfo,
		inputAuthenticatorUser,
	)

	businessPartnerAddressRes := controller.createBusinessPartnerRequestAddress(
		controller.UserInfo,
		*authenticatorUserRes,
	)

	businessPartnerPersonRes := controller.createBusinessPartnerRequestPerson(
		controller.UserInfo,
		*authenticatorUserRes,
	)

	controller.fin(
		authenticatorUserRes,
		businessPartnerAddressRes,
		businessPartnerPersonRes,
	)
}

func (
	controller *AfterLoginUserInfoController,
) fin(
	authenticatorUserRes *apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes,
	businessPartnerAddressRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {

	data := AfterLoginUserInfo{}

	for _, v := range *authenticatorUserRes.Message.User {
		data.AuthenticatorUser = append(data.AuthenticatorUser,
			apiOutputFormatter.AuthenticatorUser{
				UserID:          v.UserID,
				BusinessPartner: v.BusinessPartner,
				Language:        v.Language,
				LastLoginDate:   v.LastLoginDate,
				LastLoginTime:   v.LastLoginTime,
			},
		)
	}

	for _, v := range *businessPartnerAddressRes.Message.Address {
		data.BusinessPartnerAddress = append(data.BusinessPartnerAddress,
			apiOutputFormatter.BusinessPartnerAddress{
				BusinessPartner: v.BusinessPartner,
				LocalSubRegion:  v.LocalSubRegion,
				LocalRegion:     v.LocalRegion,
				Country:         &v.Country,
			},
		)
	}

	for _, v := range *businessPartnerPersonRes.Message.Person {
		data.BusinessPartnerPerson = append(data.BusinessPartnerPerson,
			apiOutputFormatter.BusinessPartnerPerson{
				BusinessPartner:          v.BusinessPartner,
				NickName:                 v.NickName,
				ProfileComment:           v.ProfileComment,
				PreferableLocalSubRegion: v.PreferableLocalSubRegion,
				PreferableLocalRegion:    v.PreferableLocalRegion,
				PreferableCountry:        v.PreferableCountry,
				ActPurpose:               v.ActPurpose,
			},
		)
	}

	err := controller.RedisCache.SetCache(
		controller.RedisKey,
		data,
	)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
	}

	controller.Data["json"] = &data
	controller.ServeJSON()
}
