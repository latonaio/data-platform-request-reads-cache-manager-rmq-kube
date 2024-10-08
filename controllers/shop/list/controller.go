package controllersShopList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop"
	apiModuleRuntimesRequestsShopDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop-doc"
	apiModuleRuntimesResponsesShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/shop"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type ShopListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ShopListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	localSubRegion := controller.GetString("localSubRegion")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "shop"
	redisKeyCategory2 := "list"
	redisKeyCategory3 := localSubRegion

	isReleased := true
	isMarkedForDeletion := false

	ShopAddress := apiInputReader.Shop{
		ShopAddress: &apiInputReader.ShopAddress{
			LocalSubRegion: &localSubRegion,
		},
		ShopDocHeaderDoc: &apiInputReader.ShopDocHeaderDoc{},
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
		var responseData apiOutputFormatter.Shop

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
			controller.request(ShopAddress, isReleased, isMarkedForDeletion)
		}()
	} else {
		controller.request(ShopAddress, isReleased, isMarkedForDeletion)
	}
}

func (
	controller *ShopListController,
) createShopRequestAddressesByLocalSubRegion(
	requestPram *apiInputReader.Request,
	input apiInputReader.Shop,
) *apiModuleRuntimesResponsesShop.ShopRes {
	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
	responseBody := apiModuleRuntimesRequestsShop.ShopReads(
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
	//		"ローカルサブ地域に対してのサイトが見つかりませんでした",
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
		controller.CustomLogger.Error("createShopRequestAddressesByLocalSubRegion Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ShopListController,
) createShopRequestHeadersByShops(
	requestPram *apiInputReader.Request,
	shopRes *apiModuleRuntimesResponsesShop.ShopRes,
	isReleased bool,
	isMarkedForDeletion bool,
) *apiModuleRuntimesResponsesShop.ShopRes {
	var input []apiModuleRuntimesRequestsShop.Header

	for _, v := range *shopRes.Message.Address {
		input = append(input, apiModuleRuntimesRequestsShop.Header{
			Shop:                v.Shop,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
	responseBody := apiModuleRuntimesRequestsShop.ShopReadsHeadersByShops(
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
		controller.CustomLogger.Error("createShopRequestHeadersByShops Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ShopListController,
) createShopDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.Shop,
) *apiModuleRuntimesResponsesShop.ShopDocRes {
	responseJsonData := apiModuleRuntimesResponsesShop.ShopDocRes{}
	responseBody := apiModuleRuntimesRequestsShopDoc.ShopDocReads(
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
		controller.CustomLogger.Error("createShopDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ShopListController,
) request(
	input apiInputReader.Shop,
	isReleased bool,
	isMarkedForDeletion bool,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var headerRes *apiModuleRuntimesResponsesShop.ShopRes
	var headerDocRes *apiModuleRuntimesResponsesShop.ShopDocRes

	addressRes := *controller.createShopRequestAddressesByLocalSubRegion(
		controller.UserInfo,
		input,
	)

	if addressRes.Message.Address != nil && len(*addressRes.Message.Address) != 0 {
		headerRes = *controller.createShopRequestHeadersByShops(
			controller.UserInfo,
			&addressRes,
			isReleased,
			isMarkedForDeletion,
		)

		headerDocRes = controller.createShopDocRequest(
			controller.UserInfo,
			input,
		)
	}

	controller.fin(
		&addressRes,
		&headerRes,
		headerDocRes,
	)
}

func (
	controller *ShopListController,
) fin(
	addressRes *apiModuleRuntimesResponsesShop.ShopRes,
	headerRes *apiModuleRuntimesResponsesShop.ShopRes,
	headerDocRes *apiModuleRuntimesResponsesShop.ShopDocRes,
) {

	data := apiOutputFormatter.Shop{}

	if addressRes.Message.Address != nil && len(*addressRes.Message.Address) != 0 {
	
		shopHeadersMapper := services.ShopHeadersMapper(
			headerRes,
		)

		for _, v := range *addressRes.Message.Address {
			shopType := shopHeadersMapper[strconv.Itoa(v.Shop)].ShopType
			validityStartDate := shopHeadersMapper[strconv.Itoa(v.Shop)].ValidityStartDate
			validityStartTime := shopHeadersMapper[strconv.Itoa(v.Shop)].ValidityStartTime
			validityEndDate := shopHeadersMapper[strconv.Itoa(v.Shop)].ValidityEndDate
			validityEndTime := shopHeadersMapper[strconv.Itoa(v.Shop)].ValidityEndTime
			description := shopHeadersMapper[strconv.Itoa(v.Shop)].Description
			introduction := shopHeadersMapper[strconv.Itoa(v.Shop)].Introduction
			tag1 := shopHeadersMapper[strconv.Itoa(v.Shop)].Tag1
			tag2 := shopHeadersMapper[strconv.Itoa(v.Shop)].Tag2
			tag3 := shopHeadersMapper[strconv.Itoa(v.Shop)].Tag3
			tag4 := shopHeadersMapper[strconv.Itoa(v.Shop)].Tag4
			lastChangeDate := shopHeadersMapper[strconv.Itoa(v.Shop)].LastChangeDate
			lastChangeTime := shopHeadersMapper[strconv.Itoa(v.Shop)].LastChangeTime

			img := services.ReadShopImage(
				headerDocRes,
				v.Shop,
			)

			data.ShopAddressWithHeader = append(data.ShopAddressWithHeader,
				apiOutputFormatter.ShopAddressWithHeader{
					Shop:              v.Shop,
					AddressID:         v.AddressID,
					LocalSubRegion:    v.LocalSubRegion,
					LocalRegion:       v.LocalRegion,
					ShopType:          shopType,
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

					Images: apiOutputFormatter.Images{
						Shop: img,
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
