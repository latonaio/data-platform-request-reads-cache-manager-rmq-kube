package controllersAccountInfo

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsCountry "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/country"
	apiModuleRuntimesRequestsLanguage "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/language"
	apiModuleRuntimesRequestsLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-region"
	apiModuleRuntimesRequestsLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-sub-region"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesCountry "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/country"
	apiModuleRuntimesResponsesLanguage "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/language"
	apiModuleRuntimesResponsesLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-region"
	apiModuleRuntimesResponsesLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-sub-region"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
	"sync"
)

type AccountInfoController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type AccountInfo struct {
	BusinessPartnerPerson  []apiOutputFormatter.BusinessPartnerPerson  `json:"BusinessPartnerPerson"`
	BusinessPartnerAddress []apiOutputFormatter.BusinessPartnerAddress `json:"BusinessPartnerAddress"`
	BusinessPartnerBPRole  []apiOutputFormatter.BusinessPartnerBPRole  `json:"BusinessPartnerBPRole"`
}

func (controller *AccountInfoController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)

	businessPartner := *controller.UserInfo.BusinessPartner
	redisKeyCategory1 := "account-info"
	redisKeyCategory2 := businessPartner

	isMarkedForDeletion := false

	BusinessPartner := apiInputReader.BusinessPartner{}

	BusinessPartner = apiInputReader.BusinessPartner{
		BusinessPartnerPerson: &apiInputReader.BusinessPartnerPerson{
			BusinessPartner:     businessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		BusinessPartnerAddress: &apiInputReader.BusinessPartnerAddress{
			BusinessPartner: businessPartner,
		},
		BusinessPartnerBPRole: &apiInputReader.BusinessPartnerBPRole{
			BusinessPartner: businessPartner,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			strconv.Itoa(redisKeyCategory2),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData AccountInfo

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
				BusinessPartner,
			)
		}()
	} else {
		controller.request(
			BusinessPartner,
		)
	}
}

func (
	controller *AccountInfoController,
) createBusinessPartnerRequestPerson(
	requestPram *apiInputReader.Request,
	inputBusinessPartner apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input apiModuleRuntimesRequestsBusinessPartner.Person

	input = apiModuleRuntimesRequestsBusinessPartner.Person{
		BusinessPartner: inputBusinessPartner.BusinessPartnerPerson.BusinessPartner,
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
	controller *AccountInfoController,
) createBusinessPartnerRequestAddress(
	requestPram *apiInputReader.Request,
	inputBusinessPartner apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input apiModuleRuntimesRequestsBusinessPartner.Address

	input = apiModuleRuntimesRequestsBusinessPartner.Address{
		BusinessPartner: inputBusinessPartner.BusinessPartnerPerson.BusinessPartner,
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
	controller *AccountInfoController,
) createBusinessPartnerRequestRole(
	requestPram *apiInputReader.Request,
	inputBusinessPartner apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input apiModuleRuntimesRequestsBusinessPartner.Role

	input = apiModuleRuntimesRequestsBusinessPartner.Role{
		BusinessPartner: inputBusinessPartner.BusinessPartnerPerson.BusinessPartner,
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
	controller *AccountInfoController,
) CreateLocalSubRegionRequestTextBPAddress(
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
		controller.CustomLogger.Error("CreateLocalSubRegionRequestTextBPAddress Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AccountInfoController,
) CreateLocalRegionRequestTextBPAddress(
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
		controller.CustomLogger.Error("CreateLocalRegionRequestTextBPAddress Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AccountInfoController,
) CreateCountryRequestText(
	requestPram *apiInputReader.Request,
	businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) *apiModuleRuntimesResponsesCountry.CountryRes {

	country := &(*businessPartnerPersonRes.Message.Person)[0].Nationality

	var inputCountry *string

	if country != nil {
		inputCountry = country
	}

	input := apiModuleRuntimesRequestsCountry.Country{
		Country: *inputCountry,
	}

	responseJsonData := apiModuleRuntimesResponsesCountry.CountryRes{}
	responseBody := apiModuleRuntimesRequestsCountry.CountryReadsText(
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
		controller.CustomLogger.Error("CreateCountryRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AccountInfoController,
) CreateLanguageRequestText(
	requestPram *apiInputReader.Request,
	businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) *apiModuleRuntimesResponsesLanguage.LanguageRes {

	language := &(*businessPartnerPersonRes.Message.Person)[0].Language

	var inputLanguage *string

	if language != nil {
		inputLanguage = language
	}

	input := apiModuleRuntimesRequestsLanguage.Language{
		Language: *inputLanguage,
	}

	responseJsonData := apiModuleRuntimesResponsesLanguage.LanguageRes{}
	responseBody := apiModuleRuntimesRequestsLanguage.LanguageReadsText(
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
		controller.CustomLogger.Error("CreateLanguageRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AccountInfoController,
) request(
	input apiInputReader.BusinessPartner,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var wg sync.WaitGroup
	wg.Add(3)

	var businessPartnerPersonRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerAddressRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerBPRoleRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes

	var localSubRegionTextResBPAddress *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes
	var localRegionTextResBPAddress *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes

	var countryTextRes *apiModuleRuntimesResponsesCountry.CountryRes

	var languageTextRes *apiModuleRuntimesResponsesLanguage.LanguageRes

	go func() {
		defer wg.Done()
		businessPartnerPersonRes = *controller.createBusinessPartnerRequestPerson(
			controller.UserInfo,
			input,
		)
		countryTextRes = controller.CreateCountryRequestText(
			controller.UserInfo,
			&businessPartnerPersonRes,
		)
		languageTextRes = controller.CreateLanguageRequestText(
			controller.UserInfo,
			&businessPartnerPersonRes,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerAddressRes = *controller.createBusinessPartnerRequestAddress(
			controller.UserInfo,
			input,
		)
		localSubRegionTextResBPAddress = controller.CreateLocalSubRegionRequestTextBPAddress(
			controller.UserInfo,
			&businessPartnerAddressRes,
		)
		localRegionTextResBPAddress = controller.CreateLocalRegionRequestTextBPAddress(
			controller.UserInfo,
			&businessPartnerAddressRes,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerBPRoleRes = *controller.createBusinessPartnerRequestRole(
			controller.UserInfo,
			input,
		)
	}()

	wg.Wait()

	controller.fin(
		&businessPartnerPersonRes,
		&businessPartnerAddressRes,
		localSubRegionTextResBPAddress,
		localRegionTextResBPAddress,
		countryTextRes,
		languageTextRes,
		&businessPartnerBPRoleRes,
	)
}

func (
	controller *AccountInfoController,
) fin(
	businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerAddressRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	localSubRegionTextResBPAddress *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
	localRegionTextResBPAddress *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
	countryTextRes *apiModuleRuntimesResponsesCountry.CountryRes,
	languageTextRes *apiModuleRuntimesResponsesLanguage.LanguageRes,
	businessPartnerBPRoleRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {

	localSubRegionTextMapperBPAddress := services.LocalSubRegionTextMapper(
		localSubRegionTextResBPAddress.Message.Text,
	)

	localRegionTextMapperBPAddress := services.LocalRegionTextMapper(
		localRegionTextResBPAddress.Message.Text,
	)

	countryTextMapper := services.CountryTextMapper(
		countryTextRes.Message.Text,
	)

	languageTextMapper := services.LanguageTextMapper(
		languageTextRes.Message.Text,
	)

	data := AccountInfo{}

	for _, v := range *businessPartnerPersonRes.Message.Person {

		data.BusinessPartnerPerson = append(data.BusinessPartnerPerson,
			apiOutputFormatter.BusinessPartnerPerson{
				BusinessPartner:   v.BusinessPartner,
				FirstName:         v.FirstName,
				LastName:          v.LastName,
				Gender:            v.Gender,
				Language:          v.Language,
				LanguageName:      languageTextMapper[v.Language].LanguageName,
				BirthDate:         v.BirthDate,
				Nationality:       v.Nationality,
				NationalityName:   countryTextMapper[v.Nationality].CountryName,
				EmailAddress:      v.EmailAddress,
				MobilePhoneNumber: v.MobilePhoneNumber,
			},
		)
	}

	for _, v := range *businessPartnerAddressRes.Message.Address {
		data.BusinessPartnerAddress = append(data.BusinessPartnerAddress,
			apiOutputFormatter.BusinessPartnerAddress{
				BusinessPartner:    v.BusinessPartner,
				LocalSubRegion:     v.LocalSubRegion,
				LocalSubRegionName: localSubRegionTextMapperBPAddress[v.LocalSubRegion].LocalSubRegionName,
				LocalRegion:        v.LocalRegion,
				LocalRegionName:    localRegionTextMapperBPAddress[v.LocalRegion].LocalRegionName,
				Country:            &v.Country,
			},
		)
	}

	for _, v := range *businessPartnerBPRoleRes.Message.Role {
		data.BusinessPartnerBPRole = append(data.BusinessPartnerBPRole,
			apiOutputFormatter.BusinessPartnerBPRole{
				BusinessPartner:     v.BusinessPartner,
				BusinessPartnerRole: v.BusinessPartnerRole,
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
