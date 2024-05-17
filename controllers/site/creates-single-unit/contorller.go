package controllersSiteCreatesSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsSiteType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site-type"
	apiModuleRuntimesRequestsSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site"
	apiModuleRuntimesRequestsSiteDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/site"
	apiModuleRuntimesResponsesSiteType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/site-type"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type SiteCreatesSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type SiteCreatesSingleUnit struct {
	BusinessPartnerPerson []apiOutputFormatter.BusinessPartnerPerson `json:"BusinessPartnerPerson"`
	SiteTypeText          []apiOutputFormatter.SiteTypeText          `json:"SiteTypeText"`
	SiteAddress           []apiOutputFormatter.SiteAddress           `json:"SiteAddress"`
	SiteHeader            []apiOutputFormatter.SiteHeader            `json:"SiteHeader"`
	SiteAddressWithHeader []apiOutputFormatter.SiteAddressWithHeader `json:"SiteAddressWithHeader"`
}

func (controller *SiteCreatesSingleUnitController) Get() {
	//isReleased, _ := controller.GetBool("isReleased")
	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "businessPartner"
	redisKeyCategory2 := "site-creates-single-unit"

	businessPartner, _ := controller.GetInt("businessPartner")

	localSubRegion := controller.GetString("localSubRegion")

	SiteCreatesSingleUnitBP := apiInputReader.BusinessPartner{}
	SiteCreatesSingleUnitSiteType := apiInputReader.SiteTypeGlobal{}
	SiteCreatesSingleUnitSiteAddress := apiInputReader.Site{}

	isReleased := false
	isMarkedForDeletion := false

	//docType := "QRCODE"

	SiteCreatesSingleUnitBP = apiInputReader.BusinessPartner{
		BusinessPartnerPerson: &apiInputReader.BusinessPartnerPerson{
			BusinessPartner:     businessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	SiteCreatesSingleUnitSiteType = apiInputReader.SiteTypeGlobal{
		SiteTypeText: &apiInputReader.SiteTypeText{
			Language:            *controller.UserInfo.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	SiteCreatesSingleUnitSiteAddress = apiInputReader.Site{
		SiteAddress: &apiInputReader.SiteAddress{
			LocalSubRegion: &localSubRegion,
		},
		SiteDocHeaderDoc: &apiInputReader.SiteDocHeaderDoc{},
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
		var responseData SiteCreatesSingleUnit

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
				SiteCreatesSingleUnitBP,
				SiteCreatesSingleUnitSiteType,
				SiteCreatesSingleUnitSiteAddress,
				isReleased,
				isMarkedForDeletion,
			)
		}()
	} else {
		controller.request(
			SiteCreatesSingleUnitBP,
			SiteCreatesSingleUnitSiteType,
			SiteCreatesSingleUnitSiteAddress,
			isReleased,
			isMarkedForDeletion,
		)
	}
}

func (
	controller *SiteCreatesSingleUnitController,
) createBusinessPartnerRequestPerson(
	requestPram *apiInputReader.Request,
	inputSiteCreatesSingleUnitBP apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input apiModuleRuntimesRequestsBusinessPartner.Person

	input = apiModuleRuntimesRequestsBusinessPartner.Person{
		BusinessPartner: inputSiteCreatesSingleUnitBP.BusinessPartnerPerson.BusinessPartner,
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
	controller *SiteCreatesSingleUnitController,
) CreateSiteTypeRequestTexts(
	requestPram *apiInputReader.Request,
	inputSiteCreatesSingleUnitSiteType apiInputReader.SiteTypeGlobal,
) *apiModuleRuntimesResponsesSiteType.SiteTypeRes {
	input := apiModuleRuntimesRequestsSiteType.SiteType{
		Text: []apiModuleRuntimesRequestsSiteType.Text{
			{
				Language:            inputSiteCreatesSingleUnitSiteType.SiteTypeText.Language,
				IsMarkedForDeletion: inputSiteCreatesSingleUnitSiteType.SiteTypeText.IsMarkedForDeletion,
			},
		},
	}

	responseJsonData := apiModuleRuntimesResponsesSiteType.SiteTypeRes{}
	responseBody := apiModuleRuntimesRequestsSiteType.SiteTypeReadsTexts(
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
		controller.CustomLogger.Error("CreateSiteTypeRequestTexts Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteCreatesSingleUnitController,
) createSiteRequestAddressesByLocalSubRegion(
	requestPram *apiInputReader.Request,
	input apiInputReader.Site,
) *apiModuleRuntimesResponsesSite.SiteRes {
	responseJsonData := apiModuleRuntimesResponsesSite.SiteRes{}
	responseBody := apiModuleRuntimesRequestsSite.SiteReads(
		requestPram,
		input,
		&controller.Controller,
		"AddressesByLocalSubRegion",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if len(*responseJsonData.Message.Address) == 0 {
		status := 500
		services.HandleError(
			&controller.Controller,
			"ローカルサブ地域に対してのサイトが見つかりませんでした",
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
		controller.CustomLogger.Error("createSiteRequestAddressesByLocalSubRegion Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteCreatesSingleUnitController,
) createSiteRequestHeadersBySites(
	requestPram *apiInputReader.Request,
	siteRes *apiModuleRuntimesResponsesSite.SiteRes,
	isReleased bool,
	isMarkedForDeletion bool,
) *apiModuleRuntimesResponsesSite.SiteRes {
	var input []apiModuleRuntimesRequestsSite.Header

	for _, v := range *siteRes.Message.Address {
		input = append(input, apiModuleRuntimesRequestsSite.Header{
			Site:                v.Site,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesSite.SiteRes{}
	responseBody := apiModuleRuntimesRequestsSite.SiteReadsHeadersBySites(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	if responseJsonData.Message.Header == nil {
		status := 500
		services.HandleError(
			&controller.Controller,
			"ローカルサブ地域に対して有効なサイトヘッダデータが見つかりませんでした",
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
		controller.CustomLogger.Error("createSiteRequestHeadersBySites Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteCreatesSingleUnitController,
) createSiteDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.Site,
) *apiModuleRuntimesResponsesSite.SiteDocRes {
	responseJsonData := apiModuleRuntimesResponsesSite.SiteDocRes{}
	responseBody := apiModuleRuntimesRequestsSiteDoc.SiteDocReads(
		requestPram,
		input,
		&controller.Controller,
		"HeaderDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createSiteDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteCreatesSingleUnitController,
) request(
	//	input apiInputReader.Site,
	inputSiteCreatesSingleUnitBP apiInputReader.BusinessPartner,
	inputSiteCreatesSingleUnitSiteType apiInputReader.SiteTypeGlobal,
	inputSiteCreatesSingleUnitSiteAddress apiInputReader.Site,
	isReleased bool,
	isMarkedForDeletion bool,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	businessPartnerPersonRes := *controller.createBusinessPartnerRequestPerson(
		controller.UserInfo,
		inputSiteCreatesSingleUnitBP,
	)

	siteTypeTextRes := *controller.CreateSiteTypeRequestTexts(
		controller.UserInfo,
		inputSiteCreatesSingleUnitSiteType,
	)

	siteAddressRes := *controller.createSiteRequestAddressesByLocalSubRegion(
		controller.UserInfo,
		inputSiteCreatesSingleUnitSiteAddress,
	)

	siteHeaderRes := *controller.createSiteRequestHeadersBySites(
		controller.UserInfo,
		&siteAddressRes,
		isReleased,
		isMarkedForDeletion,
	)

	siteHeaderDocRes := controller.createSiteDocRequest(
		controller.UserInfo,
		inputSiteCreatesSingleUnitSiteAddress,
	)

	controller.fin(
		&businessPartnerPersonRes,
		&siteTypeTextRes,
		&siteAddressRes,
		&siteHeaderRes,
		siteHeaderDocRes,
	)
}

func (
	controller *SiteCreatesSingleUnitController,
) fin(
	businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	siteTypeTextRes *apiModuleRuntimesResponsesSiteType.SiteTypeRes,
	siteAddressRes *apiModuleRuntimesResponsesSite.SiteRes,
	siteHeaderRes *apiModuleRuntimesResponsesSite.SiteRes,
	siteHeaderDocRes *apiModuleRuntimesResponsesSite.SiteDocRes,
) {
	siteTypeTextMapper := services.SiteTypeTextMapper(
		siteTypeTextRes.Message.Text,
	)

	siteHeadersMapper := services.SiteHeadersMapper(
		siteHeaderRes,
	)

	data := SiteCreatesSingleUnit{}

	for _, v := range *businessPartnerPersonRes.Message.Person {
		data.BusinessPartnerPerson = append(data.BusinessPartnerPerson,
			apiOutputFormatter.BusinessPartnerPerson{
				BusinessPartner: v.BusinessPartner,
				NickName:        v.NickName,
			},
		)
	}

	for _, v := range *siteTypeTextRes.Message.Text {
		data.SiteTypeText = append(data.SiteTypeText,
			apiOutputFormatter.SiteTypeText{
				SiteType:     v.SiteType,
				Language:     v.Language,
				SiteTypeName: siteTypeTextMapper[v.SiteType].SiteTypeName,
			},
		)
	}

	for _, v := range *siteHeaderRes.Message.Header {
		data.SiteHeader = append(data.SiteHeader,
			apiOutputFormatter.SiteHeader{
				Site:        v.Site,
				Description: v.Description,
			},
		)
	}

	for _, v := range *siteAddressRes.Message.Address {
		siteType := siteHeadersMapper[strconv.Itoa(v.Site)].SiteType
		validityStartDate := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityStartDate
		validityStartTime := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityStartTime
		validityEndDate := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityEndDate
		validityEndTime := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityEndTime
		introduction := siteHeadersMapper[strconv.Itoa(v.Site)].Introduction
		tag1 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag1
		tag2 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag2
		tag3 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag3
		tag4 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag4
		img := services.ReadSiteImage(
			siteHeaderDocRes,
			v.Site,
		)

		data.SiteAddressWithHeader = append(data.SiteAddressWithHeader,
			apiOutputFormatter.SiteAddressWithHeader{
				Site:              v.Site,
				AddressID:         v.AddressID,
				LocalSubRegion:    v.LocalSubRegion,
				LocalRegion:       v.LocalRegion,
				SiteType:          siteType,
				ValidityStartDate: validityStartDate,
				ValidityStartTime: validityStartTime,
				ValidityEndDate:   validityEndDate,
				ValidityEndTime:   validityEndTime,
				Introduction:      introduction,
				Tag1:              tag1,
				Tag2:              tag2,
				Tag3:              tag3,
				Tag4:              tag4,
				Images: apiOutputFormatter.Images{
					Site: img,
				},
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
