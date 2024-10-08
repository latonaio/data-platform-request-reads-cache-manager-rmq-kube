package controllersShopCreatesSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-region"
	apiModuleRuntimesRequestsLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/local-sub-region"
	apiModuleRuntimesRequestsShopType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop-type"
//	apiModuleRuntimesRequestsShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop"
//	apiModuleRuntimesRequestsShopDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop-doc"
	apiModuleRuntimesRequestsSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site"
	apiModuleRuntimesRequestsSiteDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesLocalRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-region"
	apiModuleRuntimesResponsesLocalSubRegion "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/local-sub-region"
//	apiModuleRuntimesResponsesShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/shop"
	apiModuleRuntimesResponsesShopType "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/shop-type"
	apiModuleRuntimesResponsesSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/site"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type ShopCreatesSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type ShopCreatesSingleUnit struct {
	BusinessPartnerPerson []apiOutputFormatter.BusinessPartnerPerson `json:"BusinessPartnerPerson"`
	ShopTypeText          []apiOutputFormatter.ShopTypeText          `json:"ShopTypeText"`
	ShopAddress           []apiOutputFormatter.ShopAddress           `json:"ShopAddress"`
	ShopHeader            []apiOutputFormatter.ShopHeader            `json:"ShopHeader"`
	ShopAddressWithHeader []apiOutputFormatter.ShopAddressWithHeader `json:"ShopAddressWithHeader"`
	LocalRegionText       []apiOutputFormatter.LocalRegionText       `json:"LocalRegionText"`
	LocalSubRegionText    []apiOutputFormatter.LocalSubRegionText    `json:"LocalSubRegionText"`
}

func (controller *ShopCreatesSingleUnitController) Get() {
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)

	businessPartner, _ := controller.GetInt("businessPartner")
	localSubRegion := controller.GetString("localSubRegion")

	redisKeyCategory1 := "shop"
	redisKeyCategory2 := "creates-single-unit"
	redisKeyCategory3 := localSubRegion
	redisKeyCategory4 := businessPartner

	ShopCreatesSingleUnitBP := apiInputReader.BusinessPartner{}
	ShopCreatesSingleUnitShopType := apiInputReader.ShopTypeGlobal{}
	ShopCreatesSingleUnitShopAddress := apiInputReader.Shop{}

	LocalRegion := apiInputReader.LocalRegionGlobal{}
	LocalSubRegion := apiInputReader.LocalSubRegionGlobal{}

	isReleased := false
	isMarkedForDeletion := false

	//docType := "QRCODE"

	ShopCreatesSingleUnitBP = apiInputReader.BusinessPartner{
		BusinessPartnerPerson: &apiInputReader.BusinessPartnerPerson{
			BusinessPartner:     businessPartner,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	ShopCreatesSingleUnitShopType = apiInputReader.ShopTypeGlobal{
		ShopTypeText: &apiInputReader.ShopTypeText{
			Language:            *controller.UserInfo.Language,
			IsMarkedForDeletion: &isMarkedForDeletion,
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

	ShopCreatesSingleUnitSiteAddress = apiInputReader.Site{
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
			redisKeyCategory3,
			strconv.Itoa(redisKeyCategory4),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData ShopCreatesSingleUnit

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
				ShopCreatesSingleUnitBP,
				ShopCreatesSingleUnitShopType,
				LocalRegion,
				LocalSubRegion,
				ShopCreatesSingleUnitSiteAddress,
				isReleased,
				isMarkedForDeletion,
			)
		}()
	} else {
		controller.request(
			ShopCreatesSingleUnitBP,
			ShopCreatesSingleUnitShopType,
			LocalRegion,
			LocalSubRegion,
			ShopCreatesSingleUnitShopAddress,
			isReleased,
			isMarkedForDeletion,
		)
	}
}

func (
controller *ShopCreatesSingleUnitController,
) createBusinessPartnerRequestPerson(
	requestPram *apiInputReader.Request,
	inputShopCreatesSingleUnitBP apiInputReader.BusinessPartner,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input apiModuleRuntimesRequestsBusinessPartner.Person

	input = apiModuleRuntimesRequestsBusinessPartner.Person{
		BusinessPartner: inputShopCreatesSingleUnitBP.BusinessPartnerPerson.BusinessPartner,
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
controller *ShopCreatesSingleUnitController,
) CreateShopTypeRequestTexts(
	requestPram *apiInputReader.Request,
	inputShopCreatesSingleUnitShopType apiInputReader.ShopTypeGlobal,
) *apiModuleRuntimesResponsesShopType.ShopTypeRes {
	input := apiModuleRuntimesRequestsShopType.ShopType{
		Text: []apiModuleRuntimesRequestsShopType.Text{
			{
				Language:            inputShopCreatesSingleUnitShopType.ShopTypeText.Language,
				IsMarkedForDeletion: inputShopCreatesSingleUnitShopType.ShopTypeText.IsMarkedForDeletion,
			},
		},
	}

	responseJsonData := apiModuleRuntimesResponsesShopType.ShopTypeRes{}
	responseBody := apiModuleRuntimesRequestsShopType.ShopTypeReadsTexts(
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
		controller.CustomLogger.Error("CreateShopTypeRequestTexts Unmarshal error")
	}

	return &responseJsonData
}

func (
controller *ShopCreatesSingleUnitController,
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
controller *ShopCreatesSingleUnitController,
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
controller *ShopCreatesSingleUnitController,
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
controller *ShopCreatesSingleUnitController,
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
controller *ShopCreatesSingleUnitController,
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
controller *ShopCreatesSingleUnitController,
) request(
	inputShopCreatesSingleUnitBP apiInputReader.BusinessPartner,
	inputShopCreatesSingleUnitShopType apiInputReader.ShopTypeGlobal,
	inputShopCreatesSingleUnitSiteAddress apiInputReader.Site,
	inputLocalRegion apiInputReader.LocalRegionGlobal,
	inputLocalSubRegion apiInputReader.LocalSubRegionGlobal,
	isReleased bool,
	isMarkedForDeletion bool,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	businessPartnerPersonRes := *controller.createBusinessPartnerRequestPerson(
		controller.UserInfo,
		inputShopCreatesSingleUnitBP,
	)

	shopTypeTextRes := *controller.CreateShopTypeRequestTexts(
		controller.UserInfo,
		inputShopCreatesSingleUnitShopType,
	)

	siteAddressRes := *controller.createSiteRequestAddressesByLocalSubRegion(
		controller.UserInfo,
		inputShopCreatesSingleUnitSiteAddress,
	)

	siteHeaderRes := *controller.createSiteRequestHeadersBySites(
		controller.UserInfo,
		&siteAddressRes,
		isReleased,
		isMarkedForDeletion,
	)

	siteHeaderDocRes := controller.createSiteDocRequest(
		controller.UserInfo,
		inputShopCreatesSingleUnitSiteAddress,
	)

	localRegionTextsRes := controller.CreateLocalRegionRequestTexts(
		controller.UserInfo,
		inputLocalRegion,
	)

	localSubRegionTextsRes := controller.CreateLocalSubRegionRequestTexts(
		controller.UserInfo,
		inputLocalSubRegion,
	)

	controller.fin(
		&businessPartnerPersonRes,
		&shopTypeTextRes,
		&siteAddressRes,
		&siteHeaderRes,
		siteHeaderDocRes,
		localRegionTextsRes,
		localSubRegionTextsRes,
	)
}

func (
controller *ShopCreatesSingleUnitController,
) fin(
	businessPartnerPersonRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	shopTypeTextRes *apiModuleRuntimesResponsesShopType.ShopTypeRes
	siteAddressRes *apiModuleRuntimesResponsesSite.SiteRes,
	siteHeaderRes *apiModuleRuntimesResponsesSite.SiteRes,
	siteHeaderDocRes *apiModuleRuntimesResponsesSite.SiteDocRes,
	localRegionTextsRes *apiModuleRuntimesResponsesLocalRegion.LocalRegionRes,
	localSubRegionTextsRes *apiModuleRuntimesResponsesLocalSubRegion.LocalSubRegionRes,
) {
	shopTypeTextMapper := services.ShopTypeTextMapper(
		shopTypeTextRes.Message.Text,
	)

	siteHeadersMapper := services.SiteHeadersMapper(
		siteHeaderRes,
	)

	data := ShopCreatesSingleUnit{}

	for _, v := range *businessPartnerPersonRes.Message.Person {
		data.BusinessPartnerPerson = append(data.BusinessPartnerPerson,
			apiOutputFormatter.BusinessPartnerPerson{
				BusinessPartner: v.BusinessPartner,
				NickName:        v.NickName,
			},
		)
	}

	for _, v := range *shopTypeTextRes.Message.Text {
		data.ShopTypeText = append(data.ShopTypeText,
			apiOutputFormatter.ShopTypeText{
				ShopType:     v.ShopType,
				Language:     v.Language,
				ShopTypeName: shopTypeTextMapper[v.ShopType].ShopTypeName,
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
