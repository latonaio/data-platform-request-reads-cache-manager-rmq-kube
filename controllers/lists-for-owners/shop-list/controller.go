package controllersShopListForOwners

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

type ShopListForOwnersController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ShopListForOwnersController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	organizationBusinessPartner, _ := controller.GetInt("organizationBusinessPartner")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "lists-for-owners"
	redisKeyCategory2 := "shop-list"
	redisKeyCategory3 := organizationBusinessPartner

	//	isReleased := true
	//	isCancelled := false
	isMarkedForDeletion := false

	Shop := apiInputReader.Shop{
		ShopHeader: &apiInputReader.ShopHeader{
			ShopOwner:            &organizationBusinessPartner,
			IsMarkedForDeletion:  &isMarkedForDeletion,
		},
		ShopDocHeaderDoc: &apiInputReader.ShopDocHeaderDoc{},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(redisKeyCategory3),
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
			controller.request(Shop)
		}()
	} else {
		controller.request(Shop)
	}
}

func (
	controller *ShopListForOwnersController,
) createShopRequestHeadersByShopOwner(
	requestPram *apiInputReader.Request,
	input apiInputReader.Shop,
) *apiModuleRuntimesResponsesShop.ShopRes {
	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
	responseBody := apiModuleRuntimesRequestsShop.ShopReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersByShopOwner",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)

	//	if len(*responseJsonData.Message.Address) == 0 {
	//		status := 500
	//		services.HandleError(
	//			&controller.Controller,
	//			"店舗オーナーに対しての店舗が見つかりませんでした",
	//			&status,
	//		)
	//		return nil
	//	}

	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createShopRequestHeadersByShopOwner Unmarshal error")
	}

	return &responseJsonData
}

//func (
//	controller *ShopListForOwnersController,
//) createShopRequestCountersByShops(
//	requestPram *apiInputReader.Request,
//	shopRes *apiModuleRuntimesResponsesShop.ShopRes,
//) *apiModuleRuntimesResponsesShop.ShopRes {
//	var input []apiModuleRuntimesRequestsShop.Header
//
//	for _, v := range *shopRes.Message.Header {
//		input = append(input, apiModuleRuntimesRequestsShop.Header{
//			Shop: v.Shop,
//		})
//	}

//	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
//	responseBody := apiModuleRuntimesRequestsShop.ShopReadsCountersByShops(
//		requestPram,
//		input,
//		&controller.Controller,
//	)

//	err := json.Unmarshal(responseBody, &responseJsonData)

//	if responseJsonData.Message.Counter == nil {
//		status := 500
//		services.HandleError(
//			&controller.Controller,
//			"店舗に対して有効な店舗カウンタデータが見つかりませんでした",
//			&status,
//		)
//		return nil
//	}

//	if err != nil {
//		services.HandleError(
//			&controller.Controller,
//			err,
//			nil,
//		)
//		controller.CustomLogger.Error("createShopRequestCountersByShops Unmarshal error")
//	}

//	return &responseJsonData
//}

func (
	controller *ShopListForOwnersController,
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

//	if responseJsonData.Message.HeaderDoc == nil {
//		status := 500
//		services.HandleError(
//			&controller.Controller,
//			"店舗ヘッダに画像が見つかりませんでした",
//			&status,
//		)
//		return nil
//	}

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
	controller *ShopListForOwnersController,
) request(
	input apiInputReader.Shop,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

//	var counterRes *apiModuleRuntimesResponsesShop.ShopRes
	var headerDocRes *apiModuleRuntimesResponsesShop.ShopDocRes

	headerRes := *controller.createShopRequestHeadersByShopOwner(
		controller.UserInfo,
		input,
	)
    
	if headerRes.Message.Header != nil && len(*headerRes.Message.Header) != 0 {
//      counterRes = *controller.createShopRequestCountersByShops(
//		    controller.UserInfo,
//		    &headerRes,
//	    )
	    headerDocRes = controller.createShopDocRequest(
		    controller.UserInfo,
		    input,
	    )
    }

	controller.fin(
		&headerRes,
//		&counterRes,
		headerDocRes,
	)
}

func (
	controller *ShopListForOwnersController,
) fin(
	headerRes *apiModuleRuntimesResponsesShop.ShopRes,
//	counterRes *apiModuleRuntimesResponsesShop.ShopRes,
	headerDocRes *apiModuleRuntimesResponsesShop.ShopDocRes,
) {

	data := apiOutputFormatter.Shop{}

	if headerRes.Message.Header != nil && len(*headerRes.Message.Header) != 0 {

//		shopCountersMapper := services.ShopCountersMapper(
//			counterRes,
//		)

		for _, v := range *headerRes.Message.Header {

//			numberOfLikes := shopCountersMapper[strconv.Itoa(v.Shop)].NumberOfLikes

			img := services.ReadShopImage(
				headerDocRes,
				v.Shop,
			)

			data.ShopHeader = append(data.ShopHeader,
				apiOutputFormatter.ShopHeader{
					Shop:						v.Shop,
					Description:				v.Description,
					Introduction:				v.Introduction,
					ValidityStartDate:			v.ValidityStartDate,
					ValidityStartTime:			v.ValidityStartTime,
					ValidityEndDate:			v.ValidityEndDate,
					ValidityEndTime:        	v.ValidityEndTime,
					DailyOperationStartTime:	v.DailyOperationStartTime,
					DailyOperationEndTime:		v.DailyOperationEndTime,
					Tag1:						v.Tag1,
					Tag2:						v.Tag2,
					Tag3:						v.Tag3,
					Tag4:						v.Tag4,
					LastChangeDate:				v.LastChangeDate,
					LastChangeTime:				v.LastChangeTime,
//					NumberOfLikes:				&numberOfLikes,

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
