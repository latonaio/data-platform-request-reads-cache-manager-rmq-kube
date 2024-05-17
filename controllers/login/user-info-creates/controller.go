package controllersUserInfoCreates

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsActPurpose "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/act-purpose"
	apiModuleRuntimesRequestsCountry "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/country"
	apiModuleRuntimesRequestsLanguage "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/language"
	apiModuleRuntimesRequestsLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-region"
	apiModuleRuntimesRequestsLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-sub-region"
	apiModuleRuntimesResponsesActPurpose "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/act-purpose"
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
)

type UserInfoCreatesController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type UserInfoCreates struct {
	LocalRegionText		[]apiOutputFormatter.LocalRegionText    `json:"LocalRegionText"`
	LocalSubRegionText	[]apiOutputFormatter.LocalSubRegionText `json:"LocalSubRegionText"`
	CountryText			[]apiOutputFormatter.CountryText		`json:"CountryText"`
	LanguageText		[]apiOutputFormatter.LanguageText       `json:"LanguageText"`
	ActPurposeText		[]apiOutputFormatter.ActPurposeText		`json:"ActPurposeText"`
}

func (controller *UserInfoCreatesController) Get() {
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "login"
	redisKeyCategory2 := "user-info"

	isMarkedForDeletion := false

	UserInfoCreatesLocalRegion := apiInputReader.LocalRegionGlobal{}
	UserInfoCreatesLocalSubRegion := apiInputReader.LocalSubRegionGlobal{}
	UserInfoCreatesCountry:= apiInputReader.CountryGlobal{}
	UserInfoCreatesLanguage := apiInputReader.LanguageGlobal{}
	UserInfoCreatesActPurpose := apiInputReader.ActPurposeGlobal{}

	UserInfoCreatesLocalRegion = apiInputReader.LocalRegionGlobal{
		LocalRegionText: &apiInputReader.LocalRegionText{
			Language:            *controller.UserInfo.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	UserInfoCreatesLocalSubRegion = apiInputReader.LocalSubRegionGlobal{
		LocalSubRegionText: &apiInputReader.LocalSubRegionText{
			Language:            *controller.UserInfo.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	UserInfoCreatesCountry = apiInputReader.CountryGlobal{
		CountryText: &apiInputReader.CountryText{
			Language:            *controller.UserInfo.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	UserInfoCreatesLanguage = apiInputReader.LanguageGlobal{
		LanguageText: &apiInputReader.LanguageText{
			CorrespondenceLanguage: *controller.UserInfo.Language,
			IsMarkedForDeletion:    &isMarkedForDeletion,
		},
	}

	UserInfoCreatesActPurpose = apiInputReader.ActPurposeGlobal{
		ActPurposeText: &apiInputReader.ActPurposeText{
			Language:            *controller.UserInfo.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData UserInfoCreates

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
				UserInfoCreatesLocalRegion,
				UserInfoCreatesLocalSubRegion,
				UserInfoCreatesCountry,
				UserInfoCreatesLanguage,
				UserInfoCreatesActPurpose,
			)
		}()
	} else {
		controller.request(
			UserInfoCreatesLocalRegion,
			UserInfoCreatesLocalSubRegion,
			UserInfoCreatesCountry,
			UserInfoCreatesLanguage,
			UserInfoCreatesActPurpose,
		)
	}
}

func (
	controller *UserInfoCreatesController,
) CreateLocalRegionRequestTexts(
	requestPram *apiInputReader.Request,
	inputUserInfoCreatesLocalRegion apiInputReader.LocalRegionGlobal,
) *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes {
	input := apiModuleRuntimesRequestsLocalRegion.LocalRegion{
		Text: []apiModuleRuntimesRequestsLocalRegion.Text{
			{
				Language:            inputUserInfoCreatesLocalRegion.LocalRegionText.Language,
				IsMarkedForDeletion: inputUserInfoCreatesLocalRegion.LocalRegionText.IsMarkedForDeletion,
			},
		},
	}

	responseJsonData := apiModuleRuntimesResponsesLocalRegion.LocalRegionRes{}
	responseBody := apiModuleRuntimesRequestsLocalRegion.LocalRegionReadsTexts(
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
		controller.CustomLogger.Error("CreateLocalRegionRequestTexts Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *UserInfoCreatesController,
) CreateLocalSubRegionRequestTexts(
	requestPram *apiInputReader.Request,
	inputUserInfoCreatesLocalSubRegion apiInputReader.LocalSubRegionGlobal,
) *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes {
	input := apiModuleRuntimesRequestsLocalSubRegion.LocalSubRegion{
		Text: []apiModuleRuntimesRequestsLocalSubRegion.Text{
			{
				Language:            inputUserInfoCreatesLocalSubRegion.LocalSubRegionText.Language,
				IsMarkedForDeletion: inputUserInfoCreatesLocalSubRegion.LocalSubRegionText.IsMarkedForDeletion,
			},
		},
	}

	responseJsonData := apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes{}
	responseBody := apiModuleRuntimesRequestsLocalSubRegion.LocalSubRegionReadsTexts(
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
		controller.CustomLogger.Error("CreateLocalSubRegionRequestTexts Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *UserInfoCreatesController,
) CreateCountryRequestTexts(
	requestPram *apiInputReader.Request,
	inputUserInfoCreatesCountry apiInputReader.CountryGlobal,
) *apiModuleRuntimesResponsesCountry.CountryRes {
	input := apiModuleRuntimesRequestsCountry.Country{
		Text: []apiModuleRuntimesRequestsCountry.Text{
			{
				Language:            inputUserInfoCreatesCountry.CountryText.Language,
				IsMarkedForDeletion: inputUserInfoCreatesCountry.CountryText.IsMarkedForDeletion,
			},
		},
	}

	responseJsonData := apiModuleRuntimesResponsesCountry.CountryRes{}
	responseBody := apiModuleRuntimesRequestsCountry.CountryReadsTexts(
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
		controller.CustomLogger.Error("CreateCountryRequestTexts Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *UserInfoCreatesController,
) CreateLanguageRequestTexts(
	requestPram *apiInputReader.Request,
	inputUserInfoCreatesLanguage apiInputReader.LanguageGlobal,
) *apiModuleRuntimesResponsesLanguage.LanguageRes {
	input := apiModuleRuntimesRequestsLanguage.Language{
		Text: []apiModuleRuntimesRequestsLanguage.Text{
			{
				CorrespondenceLanguage: inputUserInfoCreatesLanguage.LanguageText.CorrespondenceLanguage,
				IsMarkedForDeletion:    inputUserInfoCreatesLanguage.LanguageText.IsMarkedForDeletion,
			},
		},
	}

	responseJsonData := apiModuleRuntimesResponsesLanguage.LanguageRes{}
	responseBody := apiModuleRuntimesRequestsLanguage.LanguageReadsTexts(
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
		controller.CustomLogger.Error("CreateLanguageRequestTexts Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *UserInfoCreatesController,
) CreateActPurposeRequestTexts(
	requestPram *apiInputReader.Request,
	inputUserInfoCreatesActPurpose apiInputReader.ActPurposeGlobal,
) *apiModuleRuntimesResponsesActPurpose.ActPurposeRes {
	input := apiModuleRuntimesRequestsActPurpose.ActPurpose{
		Text: []apiModuleRuntimesRequestsActPurpose.Text{
			{
				Language:            inputUserInfoCreatesActPurpose.ActPurposeText.Language,
				IsMarkedForDeletion: inputUserInfoCreatesActPurpose.ActPurposeText.IsMarkedForDeletion,
			},
		},
	}

	responseJsonData := apiModuleRuntimesResponsesActPurpose.ActPurposeRes{}
	responseBody := apiModuleRuntimesRequestsActPurpose.ActPurposeReadsTexts(
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
		controller.CustomLogger.Error("CreateActPurposeRequestTexts Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *UserInfoCreatesController,
) request(
	inputUserInfoCreatesLocalRegion apiInputReader.LocalRegionGlobal,
	inputUserInfoCreatesLocalSubRegion apiInputReader.LocalSubRegionGlobal,
	inputUserInfoCreatesCountry apiInputReader.CountryGlobal,
	inputUserInfoCreatesLanguage apiInputReader.LanguageGlobal,
	inputUserInfoCreatesActPurpose apiInputReader.ActPurposeGlobal,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	localRegionTextRes := *controller.CreateLocalRegionRequestTexts(
		controller.UserInfo,
		inputUserInfoCreatesLocalRegion,
	)

	localSubRegionTextRes := *controller.CreateLocalSubRegionRequestTexts(
		controller.UserInfo,
		inputUserInfoCreatesLocalSubRegion,
	)

	countryTextRes := *controller.CreateCountryRequestTexts(
		controller.UserInfo,
		inputUserInfoCreatesCountry,
	)

	languageTextRes := *controller.CreateLanguageRequestTexts(
		controller.UserInfo,
		inputUserInfoCreatesLanguage,
	)

	actPurposeTextRes := *controller.CreateActPurposeRequestTexts(
		controller.UserInfo,
		inputUserInfoCreatesActPurpose,
	)

	controller.fin(
		&localRegionTextRes,
		&localSubRegionTextRes,
		&countryTextRes,
		&languageTextRes,
		&actPurposeTextRes,
	)
}

func (
	controller *UserInfoCreatesController,
) fin(
	localRegionTextRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
	localSubRegionTextRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
	countryTextRes *apiModuleRuntimesResponsesCountry.CountryRes,
	languageTextRes *apiModuleRuntimesResponsesLanguage.LanguageRes,
	actPurposeTextRes *apiModuleRuntimesResponsesActPurpose.ActPurposeRes,
) {
	data := UserInfoCreates{}

	for _, v := range *localRegionTextRes.Message.Text {
		data.LocalRegionText = append(data.LocalRegionText,
			apiOutputFormatter.LocalRegionText{
				LocalRegion:     v.LocalRegion,
				Country:         v.Country,
				Language:        v.Language,
				LocalRegionName: v.LocalRegionName,
			},
		)
	}

	for _, v := range *localSubRegionTextRes.Message.Text {
		data.LocalSubRegionText = append(data.LocalSubRegionText,
			apiOutputFormatter.LocalSubRegionText{
				LocalSubRegion:     v.LocalSubRegion,
				LocalRegion:        v.LocalRegion,
				Country:            v.Country,
				Language:           v.Language,
				LocalSubRegionName: v.LocalSubRegionName,
			},
		)
	}
	
	for _, v := range *countryTextRes.Message.Text {
		data.CountryText = append(data.CountryText,
			apiOutputFormatter.CountryText{
				Country:		v.Country,
				Language:		v.Language,
				CountryName:	v.CountryName,
			},
		)
	}

	for _, v := range *languageTextRes.Message.Text {
		data.LanguageText = append(data.LanguageText,
			apiOutputFormatter.LanguageText{
				Language:               v.Language,
				CorrespondenceLanguage: v.CorrespondenceLanguage,
				LanguageName:           v.LanguageName,
			},
		)
	}
	
	for _, v := range *actPurposeTextRes.Message.Text {
		data.ActPurposeText = append(data.ActPurposeText,
			apiOutputFormatter.ActPurposeText{
				ActPurpose:		v.ActPurpose,
				Language:		v.Language,
				ActPurposeName:	v.ActPurposeName,
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
