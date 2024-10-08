package controllersUserProfile

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsActPurpose "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/act-purpose"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsBusinessPartnerDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner-doc"
	apiModuleRuntimesRequestsCountry "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/country"
	apiModuleRuntimesRequestsLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-region"
	apiModuleRuntimesRequestsLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-sub-region"
	apiModuleRuntimesRequestsRank "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/rank"
	apiModuleRuntimesResponsesActPurpose "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/act-purpose"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesCountry "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/country"
	apiModuleRuntimesResponsesLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-region"
	apiModuleRuntimesResponsesLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-sub-region"
	apiModuleRuntimesResponsesRank "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/rank"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
	"sync"
)

type UserProfileController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type UserProfile struct {
	BusinessPartnerPerson  []apiOutputFormatter.BusinessPartnerPerson  `json:"BusinessPartnerPerson"`
	BusinessPartnerAddress []apiOutputFormatter.BusinessPartnerAddress `json:"BusinessPartnerAddress"`
	BusinessPartnerBPRole  []apiOutputFormatter.BusinessPartnerBPRole  `json:"BusinessPartnerBPRole"`
	BusinessPartnerRank    []apiOutputFormatter.BusinessPartnerRank    `json:"BusinessPartnerRank"`
	LocalRegionText        []apiOutputFormatter.LocalRegionText        `json:"LocalRegionText"`
	LocalSubRegionText     []apiOutputFormatter.LocalSubRegionText     `json:"LocalSubRegionText"`
	ActPurposeText         []apiOutputFormatter.ActPurposeText         `json:"ActPurposeText"`
}

func (controller *UserProfileController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	businessPartner, _ := controller.GetInt(":businessPartner")
	redisKeyCategory1 := "user-profile"
	redisKeyCategory2 := businessPartner
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)

	isMarkedForDeletion := false

	BusinessPartner := apiInputReader.BusinessPartner{}

	LocalRegion := apiInputReader.LocalRegionGlobal{}
	LocalSubRegion := apiInputReader.LocalSubRegionGlobal{}

	ActPurpose := apiInputReader.ActPurposeGlobal{}

	BusinessPartner = apiInputReader.BusinessPartner{
		BusinessPartnerPerson: &apiInputReader.BusinessPartnerPerson{
			BusinessPartner:     businessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		BusinessPartnerAddress: &apiInputReader.BusinessPartnerAddress{
			BusinessPartner: businessPartner,
		},
		BusinessPartnerRank: &apiInputReader.BusinessPartnerRank{
			BusinessPartner:     businessPartner,
			RankType:            "PTAP",
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		BusinessPartnerDocGeneralDoc: &apiInputReader.BusinessPartnerDocGeneralDoc{
			BusinessPartner: businessPartner,
		},
	}

	LocalRegion = apiInputReader.LocalRegionGlobal{
		LocalRegionText: &apiInputReader.LocalRegionText{
			Language:            *controller.UserInfo.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	LocalSubRegion = apiInputReader.LocalSubRegionGlobal{
		LocalSubRegionText: &apiInputReader.LocalSubRegionText{
			Language:            *controller.UserInfo.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	ActPurpose = apiInputReader.ActPurposeGlobal{
		ActPurposeText: &apiInputReader.ActPurposeText{
			Language:            *controller.UserInfo.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
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
		var responseData UserProfile

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
				LocalRegion,
				LocalSubRegion,
				ActPurpose,
			)
		}()
	} else {
		controller.request(
			BusinessPartner,
			LocalRegion,
			LocalSubRegion,
			ActPurpose,
		)
	}
}

func (
	controller *UserProfileController,
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
	controller *UserProfileController,
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
	controller *UserProfileController,
) createBusinessPartnerRequestRank(
	requestPram *apiInputReader.Request,
	inputBusinessPartner apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input apiModuleRuntimesRequestsBusinessPartner.Rank

	input = apiModuleRuntimesRequestsBusinessPartner.Rank{
		BusinessPartner: inputBusinessPartner.BusinessPartnerPerson.BusinessPartner,
		RankType:        inputBusinessPartner.BusinessPartnerRank.RankType,
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsRank(
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
		controller.CustomLogger.Error("createBusinessPartnerRequestRank Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *UserProfileController,
) createBusinessPartnerDocRequest(
	input apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes {
	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartnerDoc.BusinessPartnerDocReads(
		nil,
		input,
		&controller.Controller,
		"GeneralDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if responseJsonData.Message.GeneralDoc == nil {
		status := 500
		services.HandleError(
			&controller.Controller,
			"あなたのビジネスパートナヘッダに画像が見つかりませんでした",
			&status,
		)
		return nil
	}

	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createBusinessPartnerDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *UserProfileController,
) createBusinessPartnerDocQRCodeRequest(
	input apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes {
	docType := "QRCODE"

	input = apiInputReader.BusinessPartner{
		BusinessPartnerDocGeneralDoc: &apiInputReader.BusinessPartnerDocGeneralDoc{
			BusinessPartner: input.BusinessPartnerDocGeneralDoc.BusinessPartner,
			DocType:         &docType,
		},
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartnerDoc.BusinessPartnerDocReads(
		nil,
		input,
		&controller.Controller,
		"GeneralDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if responseJsonData.Message.GeneralDoc == nil {
		status := 500
		services.HandleError(
			&controller.Controller,
			"あなたのビジネスパートナヘッダに画像が見つかりませんでした",
			&status,
		)
		return nil
	}

	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createBusinessPartnerDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *UserProfileController,
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
	controller *UserProfileController,
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
	controller *UserProfileController,
) CreateLocalSubRegionRequestTextBPPerson(
	requestPram *apiInputReader.Request,
	businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes {

	localSubRegion := &(*businessPartnerPersonRes.Message.Person)[0].PreferableLocalSubRegion
	localRegion := &(*businessPartnerPersonRes.Message.Person)[0].PreferableLocalRegion
	country := &(*businessPartnerPersonRes.Message.Person)[0].PreferableCountry

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
		controller.CustomLogger.Error("CreateLocalSubRegionRequestTextBPPerson Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *UserProfileController,
) CreateLocalRegionRequestTextBPPerson(
	requestPram *apiInputReader.Request,
	businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes {

	localRegion := &(*businessPartnerPersonRes.Message.Person)[0].PreferableLocalRegion
	country := &(*businessPartnerPersonRes.Message.Person)[0].PreferableCountry

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
		controller.CustomLogger.Error("CreateLocalRegionRequestTextBPPerson Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *UserProfileController,
) CreateActPurposeRequestText(
	requestPram *apiInputReader.Request,
	businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) *apiModuleRuntimesResponsesActPurpose.ActPurposeRes {

	actPurpose := &(*businessPartnerPersonRes.Message.Person)[0].ActPurpose

	var inputActPurpose *string

	if actPurpose != nil {
		inputActPurpose = actPurpose
	}

	input := apiModuleRuntimesRequestsActPurpose.ActPurpose{
		ActPurpose: *inputActPurpose,
	}

	responseJsonData := apiModuleRuntimesResponsesActPurpose.ActPurposeRes{}
	responseBody := apiModuleRuntimesRequestsActPurpose.ActPurposeReadsText(
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
		controller.CustomLogger.Error("CreateActPurposeRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *UserProfileController,
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
	controller *UserProfileController,
) CreateRankRequestText(
	requestPram *apiInputReader.Request,
	businessPartnerRankRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) *apiModuleRuntimesResponsesRank.RankRes {

	rankType := &(*businessPartnerRankRes.Message.Rank)[0].RankType
	rank := (*businessPartnerRankRes.Message.Rank)[0].Rank

	var inputRankType *string
	var inputRank int

	if rankType != nil {
		inputRankType = rankType
		inputRank = rank
	}

	input := apiModuleRuntimesRequestsRank.Rank{
		RankType: *inputRankType,
		Rank:     inputRank,
	}

	responseJsonData := apiModuleRuntimesResponsesRank.RankRes{}
	responseBody := apiModuleRuntimesRequestsRank.RankReadsText(
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
		controller.CustomLogger.Error("CreateRankRequestText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *UserProfileController,
) CreateLocalRegionRequestTexts(
	requestPram *apiInputReader.Request,
	inputLocalRegion apiInputReader.LocalRegionGlobal,
) *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes {
	input := apiModuleRuntimesRequestsLocalRegion.LocalRegion{
		Text: []apiModuleRuntimesRequestsLocalRegion.Text{
			{
				Language:            inputLocalRegion.LocalRegionText.Language,
				IsMarkedForDeletion: inputLocalRegion.LocalRegionText.IsMarkedForDeletion,
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
	controller *UserProfileController,
) CreateLocalSubRegionRequestTexts(
	requestPram *apiInputReader.Request,
	inputLocalSubRegion apiInputReader.LocalSubRegionGlobal,
) *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes {
	input := apiModuleRuntimesRequestsLocalSubRegion.LocalSubRegion{
		Text: []apiModuleRuntimesRequestsLocalSubRegion.Text{
			{
				Language:            inputLocalSubRegion.LocalSubRegionText.Language,
				IsMarkedForDeletion: inputLocalSubRegion.LocalSubRegionText.IsMarkedForDeletion,
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
	controller *UserProfileController,
) CreateActPurposeRequestTexts(
	requestPram *apiInputReader.Request,
	inputActPurpose apiInputReader.ActPurposeGlobal,
) *apiModuleRuntimesResponsesActPurpose.ActPurposeRes {
	input := apiModuleRuntimesRequestsActPurpose.ActPurpose{
		Text: []apiModuleRuntimesRequestsActPurpose.Text{
			{
				Language:            inputActPurpose.ActPurposeText.Language,
				IsMarkedForDeletion: inputActPurpose.ActPurposeText.IsMarkedForDeletion,
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
	controller *UserProfileController,
) request(
	inputBusinessPartner apiInputReader.BusinessPartner,
	inputLocalRegion apiInputReader.LocalRegionGlobal,
	inputLocalSubRegion apiInputReader.LocalSubRegionGlobal,
	inputActPurpose apiInputReader.ActPurposeGlobal,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var wg sync.WaitGroup
	wg.Add(7)

	var businessPartnerPersonRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerAddressRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerRankRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes
	var businessPartnerGeneralDocRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes
	var businessPartnerGeneralQRCodeDocRes apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes

	var localSubRegionTextResBPPerson *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes
	var localRegionTextResBPPerson *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes

	var localSubRegionTextResBPAddress *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes
	var localRegionTextResBPAddress *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes

	var actPurposeTextRes *apiModuleRuntimesResponsesActPurpose.ActPurposeRes
	var countryTextRes *apiModuleRuntimesResponsesCountry.CountryRes

	var rankTextRes *apiModuleRuntimesResponsesRank.RankRes

	var localSubRegionTextsRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes
	var localRegionTextsRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes

	var actPurposeTextsRes *apiModuleRuntimesResponsesActPurpose.ActPurposeRes

	go func() {
		defer wg.Done()
		businessPartnerPersonRes = *controller.createBusinessPartnerRequestPerson(
			controller.UserInfo,
			inputBusinessPartner,
		)
		localSubRegionTextResBPPerson = controller.CreateLocalSubRegionRequestTextBPPerson(
			controller.UserInfo,
			&businessPartnerPersonRes,
		)
		localRegionTextResBPPerson = controller.CreateLocalRegionRequestTextBPPerson(
			controller.UserInfo,
			&businessPartnerPersonRes,
		)
		actPurposeTextRes = controller.CreateActPurposeRequestText(
			controller.UserInfo,
			&businessPartnerPersonRes,
		)
		countryTextRes = controller.CreateCountryRequestText(
			controller.UserInfo,
			&businessPartnerPersonRes,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerAddressRes = *controller.createBusinessPartnerRequestAddress(
			controller.UserInfo,
			inputBusinessPartner,
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
		businessPartnerRankRes = *controller.createBusinessPartnerRequestRank(
			controller.UserInfo,
			inputBusinessPartner,
		)
		rankTextRes = controller.CreateRankRequestText(
			controller.UserInfo,
			&businessPartnerRankRes,
		)
	}()

	go func() {
		defer wg.Done()
		businessPartnerGeneralDocRes = *controller.createBusinessPartnerDocRequest(
			inputBusinessPartner,
		)
		businessPartnerGeneralQRCodeDocRes = *controller.createBusinessPartnerDocQRCodeRequest(
			inputBusinessPartner,
		)
	}()

	go func() {
		defer wg.Done()
		localSubRegionTextsRes = controller.CreateLocalSubRegionRequestTexts(
			controller.UserInfo,
			inputLocalSubRegion,
		)
	}()

	go func() {
		defer wg.Done()
		localRegionTextsRes = controller.CreateLocalRegionRequestTexts(
			controller.UserInfo,
			inputLocalRegion,
		)
	}()

	go func() {
		defer wg.Done()
		actPurposeTextsRes = controller.CreateActPurposeRequestTexts(
			controller.UserInfo,
			inputActPurpose,
		)
	}()

	wg.Wait()

	controller.fin(
		&businessPartnerPersonRes,
		localSubRegionTextResBPPerson,
		localRegionTextResBPPerson,
		actPurposeTextRes,
		countryTextRes,
		&businessPartnerAddressRes,
		localSubRegionTextResBPAddress,
		localRegionTextResBPAddress,
		&businessPartnerRankRes,
		&businessPartnerGeneralDocRes,
		&businessPartnerGeneralQRCodeDocRes,
		rankTextRes,
		localSubRegionTextsRes,
		localRegionTextsRes,
		actPurposeTextsRes,
	)
}

func (
	controller *UserProfileController,
) fin(
	businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	localSubRegionTextResBPPerson *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
	localRegionTextResBPPerson *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
	actPurposeTextRes *apiModuleRuntimesResponsesActPurpose.ActPurposeRes,
	countryTextRes *apiModuleRuntimesResponsesCountry.CountryRes,
	businessPartnerAddressRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	localSubRegionTextResBPAddress *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
	localRegionTextResBPAddress *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
	businessPartnerRankRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	businessPartnerGeneralDocRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes,
	businessPartnerGeneralQRCodeDocRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerDocRes,
	rankTextRes *apiModuleRuntimesResponsesRank.RankRes,
	localSubRegionTextsRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
	localRegionTextsRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
	actPurposeTextsRes *apiModuleRuntimesResponsesActPurpose.ActPurposeRes,
) {

	//businessPartnerRoleTextMapper := services.BusinessPartnerRoleTextMapper(
	//	businessPartnerRoleTextRes.Message.Text,
	//)

	localSubRegionTextMapperBPPerson := services.LocalSubRegionTextMapper(
		localSubRegionTextResBPPerson.Message.Text,
	)

	localRegionTextMapperBPPerson := services.LocalRegionTextMapper(
		localRegionTextResBPPerson.Message.Text,
	)

	actPurposeTextMapper := services.ActPurposeTextMapper(
		actPurposeTextRes.Message.Text,
	)

	countryTextMapper := services.CountryTextMapper(
		countryTextRes.Message.Text,
	)

	localSubRegionTextMapperBPAddress := services.LocalSubRegionTextMapper(
		localSubRegionTextResBPAddress.Message.Text,
	)

	localRegionTextMapperBPAddress := services.LocalRegionTextMapper(
		localRegionTextResBPAddress.Message.Text,
	)

	rankTextMapper := services.RankTextMapper(
		rankTextRes.Message.Text,
	)

	data := UserProfile{}

	for _, v := range *businessPartnerPersonRes.Message.Person {

		img := services.ReadBusinessPartnerImage(
			businessPartnerGeneralDocRes,
			v.BusinessPartner,
		)

		qrcode := services.CreateQRCodeBusinessPartnerDocImage(
			businessPartnerGeneralQRCodeDocRes,
			v.BusinessPartner,
		)

		documentImage := services.ReadDocumentImageBusinessPartner(
			businessPartnerGeneralDocRes,
			v.BusinessPartner,
		)

		data.BusinessPartnerPerson = append(data.BusinessPartnerPerson,
			apiOutputFormatter.BusinessPartnerPerson{
				BusinessPartner:              v.BusinessPartner,
				NickName:                     v.NickName,
				Gender:                       v.Gender,
				Language:                     v.Language,
				ProfileComment:               v.ProfileComment,
				BirthDate:                    v.BirthDate,
				Nationality:                  v.Nationality,
				NationalityName:              countryTextMapper[v.Nationality].CountryName,
				EmailAddress:                 v.EmailAddress,
				MobilePhoneNumber:            v.MobilePhoneNumber,
				PreferableLocalSubRegion:     v.PreferableLocalSubRegion,
				PreferableLocalSubRegionName: localSubRegionTextMapperBPPerson[v.PreferableLocalSubRegion].LocalSubRegionName,
				PreferableLocalRegion:        v.PreferableLocalRegion,
				PreferableLocalRegionName:    localRegionTextMapperBPPerson[v.PreferableLocalRegion].LocalRegionName,
				PreferableCountry:            v.PreferableCountry,
				ActPurpose:                   v.ActPurpose,
				ActPurposeName:               actPurposeTextMapper[v.ActPurpose].ActPurposeName,
				TermsOfUseIsConfirmed:        v.TermsOfUseIsConfirmed,
				Images: apiOutputFormatter.Images{
					BusinessPartner:              img,
					QRCode:                       qrcode,
					DocumentImageBusinessPartner: documentImage,
				},
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

	for _, v := range *businessPartnerRankRes.Message.Rank {
		data.BusinessPartnerRank = append(data.BusinessPartnerRank,
			apiOutputFormatter.BusinessPartnerRank{
				BusinessPartner: v.BusinessPartner,
				RankType:        v.RankType,
				Rank:            v.Rank,
				RankName:        rankTextMapper[v.Rank].RankName,
			},
		)
	}

	for _, v := range *localSubRegionTextsRes.Message.Text {
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

	for _, v := range *localRegionTextsRes.Message.Text {
		data.LocalRegionText = append(data.LocalRegionText,
			apiOutputFormatter.LocalRegionText{
				LocalRegion:     v.LocalRegion,
				Country:         v.Country,
				Language:        v.Language,
				LocalRegionName: v.LocalRegionName,
			},
		)
	}

	for _, v := range *actPurposeTextsRes.Message.Text {
		data.ActPurposeText = append(data.ActPurposeText,
			apiOutputFormatter.ActPurposeText{
				ActPurpose:     v.ActPurpose,
				Language:       v.Language,
				ActPurposeName: v.ActPurposeName,
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
