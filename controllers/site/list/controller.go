package controllersSiteList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site"
	apiModuleRuntimesRequestsSiteDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/site/site-doc"
	apiModuleRuntimesResponsesSite "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/site"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type SiteListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *SiteListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	localSubRegion := controller.GetString("localSubRegion")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "site"
	redisKeyCategory2 := "list"
	redisKeyCategory3 := localSubRegion

	isReleased := true
	isMarkedForDeletion := false

	SiteAddress := apiInputReader.Site{
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
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.Site

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
			controller.request(SiteAddress, isReleased, isMarkedForDeletion)
		}()
	} else {
		controller.request(SiteAddress, isReleased, isMarkedForDeletion)
	}
}

func (
	controller *SiteListController,
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

	//if len(*responseJsonData.Message.Address) == 0 {
	//	status := 500
	//	services.HandleError(
	//		&controller.Controller,
	//		"ローカルサブ地域に対してのサイトアドレスが見つかりませんでした",
	//		&status,
	//	)
	//	return nil
	//}

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
	controller *SiteListController,
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

	//if responseJsonData.Message.Header == nil {
	//	status := 500
	//	services.HandleError(
	//		&controller.Controller,
	//		"ローカルサブ地域に対して有効なサイトヘッダデータが見つかりませんでした",
	//		&status,
	//	)
	//	return nil
	//}

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
	controller *SiteListController,
) createSiteRequestCountersBySites(
	requestPram *apiInputReader.Request,
	siteRes *apiModuleRuntimesResponsesSite.SiteRes,
) *apiModuleRuntimesResponsesSite.SiteRes {
	var input []apiModuleRuntimesRequestsSite.Header

	for _, v := range *siteRes.Message.Address {
		input = append(input, apiModuleRuntimesRequestsSite.Header{
			Site: v.Site,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesSite.SiteRes{}
	responseBody := apiModuleRuntimesRequestsSite.SiteReadsCountersByEvents(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	//if responseJsonData.Message.Counter == nil {
	//	status := 500
	//	services.HandleError(
	//		&controller.Controller,
	//		"ローカルサブ地域に対して有効なサイトカウンタデータが見つかりませんでした",
	//		&status,
	//	)
	//	return nil
	//}

	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createSiteRequestCountersBySites Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteListController,
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

	//if responseJsonData.Message.HeaderDoc == nil {
	//	status := 500
	//	services.HandleError(
	//		&controller.Controller,
	//		"サイトヘッダに画像が見つかりませんでした",
	//		&status,
	//	)
	//	return nil
	//}

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
	controller *SiteListController,
) request(
	input apiInputReader.Site,
	isReleased bool,
	isMarkedForDeletion bool,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var headerRes *apiModuleRuntimesResponsesSite.SiteRes
	var counterRes *apiModuleRuntimesResponsesSite.SiteRes
	var headerDocRes *apiModuleRuntimesResponsesSite.SiteDocRes

	addressRes := *controller.createSiteRequestAddressesByLocalSubRegion(
		controller.UserInfo,
		input,
	)

	if addressRes.Message.Address != nil && len(*addressRes.Message.Address) != 0 {
		headerRes = controller.createSiteRequestHeadersBySites(
			controller.UserInfo,
			&addressRes,
			isReleased,
			isMarkedForDeletion,
		)

		counterRes = controller.createSiteRequestCountersBySites(
			controller.UserInfo,
			&addressRes,
		)

		headerDocRes = controller.createSiteDocRequest(
			controller.UserInfo,
			input,
		)
	}

	controller.fin(
		&addressRes,
		headerRes,
		counterRes,
		headerDocRes,
	)
}

func (
	controller *SiteListController,
) fin(
	addressRes *apiModuleRuntimesResponsesSite.SiteRes,
	headerRes *apiModuleRuntimesResponsesSite.SiteRes,
	counterRes *apiModuleRuntimesResponsesSite.SiteRes,
	headerDocRes *apiModuleRuntimesResponsesSite.SiteDocRes,
) {
	data := apiOutputFormatter.Site{}

	if addressRes.Message.Address != nil && len(*addressRes.Message.Address) != 0 {
		siteHeadersMapper := services.SiteHeadersMapper(
			headerRes,
		)

		siteCountersMapper := services.SiteCountersMapper(
			counterRes,
		)

		for _, v := range *addressRes.Message.Address {
			siteType := siteHeadersMapper[strconv.Itoa(v.Site)].SiteType
			validityStartDate := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityStartDate
			validityStartTime := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityStartTime
			validityEndDate := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityEndDate
			validityEndTime := siteHeadersMapper[strconv.Itoa(v.Site)].ValidityEndTime
			description := siteHeadersMapper[strconv.Itoa(v.Site)].Description
			introduction := siteHeadersMapper[strconv.Itoa(v.Site)].Introduction
			tag1 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag1
			tag2 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag2
			tag3 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag3
			tag4 := siteHeadersMapper[strconv.Itoa(v.Site)].Tag4
			lastChangeDate := siteHeadersMapper[strconv.Itoa(v.Site)].LastChangeDate
			lastChangeTime := siteHeadersMapper[strconv.Itoa(v.Site)].LastChangeTime

			numberOfLikes := siteCountersMapper[strconv.Itoa(v.Site)].NumberOfLikes

			img := services.ReadSiteImage(
				headerDocRes,
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
					Description:       description,
					Introduction:      introduction,
					Tag1:              tag1,
					Tag2:              tag2,
					Tag3:              tag3,
					Tag4:              tag4,
					LastChangeDate:    lastChangeDate,
					LastChangeTime:    lastChangeTime,
					NumberOfLikes:     &numberOfLikes,

					Images: apiOutputFormatter.Images{
						Site: img,
					},
				},
			)
		}
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
