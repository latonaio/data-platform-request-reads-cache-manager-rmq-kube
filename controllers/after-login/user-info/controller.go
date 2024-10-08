package controllersAfterLoginUserInfo

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsAuthenticator "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/authenticator"
	apiModuleRuntimesRequestsBusinessPartnerRole "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner-role"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-region"
	apiModuleRuntimesRequestsLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-sub-region"
	apiModuleRuntimesResponsesAuthenticator "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/authenticator"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesBusinessPartnerRole "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner-role"
	apiModuleRuntimesResponsesLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-region"
	apiModuleRuntimesResponsesLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-sub-region"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"sync"
)

type AfterLoginUserInfoController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type AfterLoginUserInfo struct {
	AuthenticatorUser					[]apiOutputFormatter.AuthenticatorUser					`json:"AuthenticatorUser"`
	BusinessPartnerAddress				[]apiOutputFormatter.BusinessPartnerAddress				`json:"BusinessPartnerAddress"`
	BusinessPartnerPerson  				[]apiOutputFormatter.BusinessPartnerPerson				`json:"BusinessPartnerPerson"`
	BusinessPartnerPersonOrganization	[]apiOutputFormatter.BusinessPartnerPersonOrganization  `json:"BusinessPartnerPersonOrganization"`
	BusinessPartnerBPRole				[]apiOutputFormatter.BusinessPartnerBPRole				`json:"BusinessPartnerBPRole"`
}

func (controller *AfterLoginUserInfoController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	userID := controller.GetString("userID")
	redisKeyCategory1 := "after-login"
	redisKeyCategory2 := "user-info"
	redisKeyCategory3 := userID
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)

	isMarkedForDeletion := false

	AuthenticatorUser := apiInputReader.Authenticator{}

	AuthenticatorUser = apiInputReader.Authenticator{
		AuthenticatorUser: &apiInputReader.AuthenticatorUser{
			UserID:              userID,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			redisKeyCategory3,
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
			)
		}()
	} else {
		controller.request(
			AuthenticatorUser,
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
) createBusinessPartnerRequestPersonOrganization(
	requestPram *apiInputReader.Request,
	authenticatorRes apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input = apiModuleRuntimesRequestsBusinessPartner.PersonOrganization{}

	for _, v := range *authenticatorRes.Message.User {
		input = apiModuleRuntimesRequestsBusinessPartner.PersonOrganization{
			BusinessPartner: v.BusinessPartner,
		}
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsPersonOrganization(
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
		controller.CustomLogger.Error("createBusinessPartnerRequestPersonOrganization Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AfterLoginUserInfoController,
) createBusinessPartnerRequestRole(
	requestPram *apiInputReader.Request,
	authenticatorRes apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input apiModuleRuntimesRequestsBusinessPartner.Role

	for _, v := range *authenticatorRes.Message.User {
		input = apiModuleRuntimesRequestsBusinessPartner.Role{
			BusinessPartner: v.BusinessPartner,
		}
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsRole(
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
		controller.CustomLogger.Error("createBusinessPartnerRequestRole Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AfterLoginUserInfoController,
) CreateBusinessPartnerRoleRequestText(
	requestPram *apiInputReader.Request,
	authenticatorRes apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes,
	businessPartnerRoleRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes {
	for _, v := range *authenticatorRes.Message.User {
		var language *string
		language = &v.Language
		requestPram.Language = language
	}

	businessPartnerRole := &(*businessPartnerRoleRes.Message.Role)[0].BusinessPartnerRole

	var inputBusinessPartnerRole *string

	if businessPartnerRole != nil {
		inputBusinessPartnerRole = businessPartnerRole
	}

	input := apiModuleRuntimesRequestsBusinessPartnerRole.BusinessPartnerRole{
		BusinessPartnerRole: *inputBusinessPartnerRole,
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartnerRole.BusinessPartnerRoleReadsText(
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
		controller.CustomLogger.Error("CreateBusinessPartnerRoleRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AfterLoginUserInfoController,
) CreateLocalSubRegionRequestText(
	requestPram *apiInputReader.Request,
	businessPartnerAddressRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes {

	localSubRegion := &(*businessPartnerAddressRes.Message.Address)[0].LocalSubRegion
	localRegion := &(*businessPartnerAddressRes.Message.Address)[0].LocalRegion
	country := &(*businessPartnerAddressRes.Message.Address)[0].Country

	var inputLocalSubRegion *string
	var inputLocalRegion *string
	var inputCountry *string

	if localRegion != nil {
		inputLocalSubRegion = localSubRegion
		inputLocalRegion = localRegion
		inputCountry = country
	}

	input := apiModuleRuntimesRequestsLocalSubRegion.LocalSubRegion{
		LocalSubRegion: *inputLocalSubRegion,
		LocalRegion:    *inputLocalRegion,
		Country:        *inputCountry,
	}

	responseJsonData := apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes{}
	responseBody := apiModuleRuntimesRequestsLocalSubRegion.LocalSubRegionReadsText(
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
		controller.CustomLogger.Error("CreateLocalSubRegionRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AfterLoginUserInfoController,
) CreateLocalRegionRequestText(
	requestPram *apiInputReader.Request,
	businessPartnerAddressRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes {

	localRegion := &(*businessPartnerAddressRes.Message.Address)[0].LocalRegion
	country := &(*businessPartnerAddressRes.Message.Address)[0].Country

	var inputLocalRegion *string
	var inputCountry *string

	if localRegion != nil {
		inputLocalRegion = localRegion
		inputCountry = country
	}

	input := apiModuleRuntimesRequestsLocalRegion.LocalRegion{
		LocalRegion: *inputLocalRegion,
		Country:     *inputCountry,
	}

	responseJsonData := apiModuleRuntimesResponsesLocalRegion.LocalRegionRes{}
	responseBody := apiModuleRuntimesRequestsLocalRegion.LocalRegionReadsText(
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
		controller.CustomLogger.Error("CreateLocalRegionRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AfterLoginUserInfoController,
) request(
	inputAuthenticatorUser apiInputReader.Authenticator,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var wg sync.WaitGroup
	wg.Add(4)

	var businessPartnerAddressRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	
	var businessPartnerPersonRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerPersonOrganizationRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes

	var businessPartnerRoleRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerRoleTextRes apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes

	var localSubRegionTextRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes
	var localRegionTextRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes

	authenticatorUserRes := controller.createAuthenticatorRequestUser(
		controller.UserInfo,
		inputAuthenticatorUser,
	)

	go func() {
		defer wg.Done()
		businessPartnerAddressRes = *controller.createBusinessPartnerRequestAddress(
			controller.UserInfo,
			*authenticatorUserRes,
		)
		localSubRegionTextRes = controller.CreateLocalSubRegionRequestText(
			controller.UserInfo,
			&businessPartnerAddressRes,
		)

		localRegionTextRes = controller.CreateLocalRegionRequestText(
			controller.UserInfo,
			&businessPartnerAddressRes,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerPersonRes = *controller.createBusinessPartnerRequestPerson(
			controller.UserInfo,
			*authenticatorUserRes,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerPersonOrganizationRes = *controller.createBusinessPartnerRequestPersonOrganization(
			controller.UserInfo,
			*authenticatorUserRes,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerRoleRes = *controller.createBusinessPartnerRequestRole(
			controller.UserInfo,
			*authenticatorUserRes,
		)
		businessPartnerRoleTextRes = *controller.CreateBusinessPartnerRoleRequestText(
			controller.UserInfo,
			*authenticatorUserRes,
			&businessPartnerRoleRes,
		)
	}()

	wg.Wait()

	controller.fin(
		authenticatorUserRes,
		&businessPartnerAddressRes,
		&businessPartnerPersonRes,
		&businessPartnerPersonOrganizationRes,
		&businessPartnerRoleRes,
		&businessPartnerRoleTextRes,
		localSubRegionTextRes,
		localRegionTextRes,
	)
}

func (
	controller *AfterLoginUserInfoController,
) fin(
	authenticatorUserRes *apiModuleRuntimesResponsesAuthenticator.AuthenticatorRes,
	businessPartnerAddressRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerPersonOrganizationRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerRoleRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerRoleTextRes *apiModuleRuntimesResponsesBusinessPartnerRole.BusinessPartnerRoleRes,
	localSubRegionTextRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
	localRegionTextRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
) {

	businessPartnerRoleTextMapper := services.BusinessPartnerRoleTextMapper(
		businessPartnerRoleTextRes.Message.Text,
	)

	localSubRegionTextMapper := services.LocalSubRegionTextMapper(
		localSubRegionTextRes.Message.Text,
	)

	localRegionTextMapper := services.LocalRegionTextMapper(
		localRegionTextRes.Message.Text,
	)

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
				BusinessPartner:	 v.BusinessPartner,
				LocalSubRegion:  	 v.LocalSubRegion,
				LocalSubRegionName:  localSubRegionTextMapper[v.LocalSubRegion].LocalSubRegionName,
				LocalRegion:     	 v.LocalRegion,
				LocalRegionName:     localRegionTextMapper[v.LocalRegion].LocalRegionName,
				Country:         	 &v.Country,
			},
		)
	}

	for _, v := range *businessPartnerPersonRes.Message.Person {
		data.BusinessPartnerPerson = append(data.BusinessPartnerPerson,
			apiOutputFormatter.BusinessPartnerPerson{
				BusinessPartner:          		v.BusinessPartner,
				NickName:                 		v.NickName,
				ProfileComment:           		v.ProfileComment,
				PreferableLocalSubRegion:		v.PreferableLocalSubRegion,
				PreferableLocalSubRegionName:	localSubRegionTextMapper[v.PreferableLocalSubRegion].LocalSubRegionName,
				PreferableLocalRegion:			v.PreferableLocalRegion,
				PreferableLocalRegionName:		localRegionTextMapper[v.PreferableLocalRegion].LocalRegionName,
				PreferableCountry:				v.PreferableCountry,
				ActPurpose:						v.ActPurpose,
			},
		)
	}

	for _, v := range *businessPartnerPersonOrganizationRes.Message.PersonOrganization {
		data.BusinessPartnerPersonOrganization = append(data.BusinessPartnerPersonOrganization,
			apiOutputFormatter.BusinessPartnerPersonOrganization{
				BusinessPartner:          		v.BusinessPartner,
				OrganizationBusinessPartner:	v.OrganizationBusinessPartner,
			},
		)
	}

	for _, v := range *businessPartnerRoleRes.Message.Role {
		data.BusinessPartnerBPRole = append(data.BusinessPartnerBPRole,
			apiOutputFormatter.BusinessPartnerBPRole{
				BusinessPartner:         v.BusinessPartner,
				BusinessPartnerRole:     v.BusinessPartnerRole,
				BusinessPartnerRoleName: businessPartnerRoleTextMapper[v.BusinessPartnerRole].BusinessPartnerRoleName,
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
