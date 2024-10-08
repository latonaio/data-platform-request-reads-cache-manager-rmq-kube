package controllersPointConsumption

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsPointBalance "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/point-balance"
	apiModuleRuntimesRequestsShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/shop/shop"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesPointBalance "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/point-balance"
	apiModuleRuntimesResponsesShop "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/shop"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
	"sync"
)

type PointConsumptionController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

type PointConsumptionGlobal struct {
	PointBalancePointBalanceSender  []apiOutputFormatter.PointBalancePointBalance `json:"PointBalancePointBalanceSender"`
	PointBalanceByShopReceiver      []apiOutputFormatter.PointBalanceByShop       `json:"PointBalanceByShopReceiver"`
	ShopHeader                      []apiOutputFormatter.ShopHeader               `json:"ShopHeader"`
	BusinessPartnerGeneralShopOwner []apiOutputFormatter.BusinessPartnerGeneral   `json:"BusinessPartnerGeneralShopOwner"`
}

func (controller *PointConsumptionController) Get() {
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "businessPartner"
	redisKeyCategory2 := "shop"
	businessPartner, _ := controller.GetInt("businessPartner")
	shopOwner, _ := controller.GetInt("shopOwner")
	pointSymbol := "POYPO"
	shop, _ := controller.GetInt("shop")

	PointBalanceSender := apiInputReader.PointBalanceGlobal{}
	PointBalanceReceiver := apiInputReader.PointBalanceGlobal{}

	Shop := apiInputReader.Shop{}

	isMarkedForDeletion := false

	PointBalanceSender = apiInputReader.PointBalanceGlobal{
		PointBalance: &apiInputReader.PointBalance{
			BusinessPartner: businessPartner,
			PointSymbol:     pointSymbol,
		},
	}

	PointBalanceReceiver = apiInputReader.PointBalanceGlobal{
		PointBalance: &apiInputReader.PointBalance{
			ByShop: []apiInputReader.ByShop{
				{
					BusinessPartner: shopOwner,
					PointSymbol:     pointSymbol,
					Shop:            shop,
				},
			},
		},
	}

	Shop = apiInputReader.Shop{
		ShopHeader: &apiInputReader.ShopHeader{
			Shop:                shop,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(businessPartner),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData PointConsumptionGlobal

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
			controller.request(PointBalanceSender, PointBalanceReceiver, Shop)
		}()
	} else {
		controller.request(PointBalanceSender, PointBalanceReceiver, Shop)
	}
}

func (
	controller *PointConsumptionController,
) createPointBalanceRequestPointBalanceSender(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointBalanceGlobal,
) *apiModuleRuntimesResponsesPointBalance.PointBalanceRes {
	responseJsonData := apiModuleRuntimesResponsesPointBalance.PointBalanceRes{}
	responseBody := apiModuleRuntimesRequestsPointBalance.PointBalanceReadsPointBalance(
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
		controller.CustomLogger.Error("createPointBalanceRequestPointBalanceSender Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointConsumptionController,
) createPointBalanceRequestByShopReceiver(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointBalanceGlobal,
) *apiModuleRuntimesResponsesPointBalance.PointBalanceRes {
	responseJsonData := apiModuleRuntimesResponsesPointBalance.PointBalanceRes{}
	responseBody := apiModuleRuntimesRequestsPointBalance.PointBalanceReadsByShop(
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
		controller.CustomLogger.Error("createPointBalanceRequestPointBalanceSender Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointConsumptionController,
) createPointBalanceRequestByShopSender(
	requestPram *apiInputReader.Request,
	input apiInputReader.PointBalanceGlobal,
) *apiModuleRuntimesResponsesPointBalance.PointBalanceRes {
	responseJsonData := apiModuleRuntimesResponsesPointBalance.PointBalanceRes{}
	responseBody := apiModuleRuntimesRequestsPointBalance.PointBalanceReadsByShop(
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
		controller.CustomLogger.Error("createPointBalanceRequestByShopSender Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointConsumptionController,
) createShopRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.Shop,
) *apiModuleRuntimesResponsesShop.ShopRes {
	responseJsonData := apiModuleRuntimesResponsesShop.ShopRes{}
	responseBody := apiModuleRuntimesRequestsShop.ShopReads(
		requestPram,
		input,
		&controller.Controller,
		"Header",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createShopRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointConsumptionController,
) createBusinessPartnerRequestGeneralShopOwner(
	requestPram *apiInputReader.Request,
	shopRes *apiModuleRuntimesResponsesShop.ShopRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*shopRes.Message.Header))

	for _, v := range *shopRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.ShopOwner,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsGeneralsByBusinessPartners(
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
		controller.CustomLogger.Error("createBusinessPartnerRequestGeneralShopOwner Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointConsumptionController,
) request(
	inputPointBalanceSender apiInputReader.PointBalanceGlobal,
	inputPointBalanceReceiver apiInputReader.PointBalanceGlobal,
	inputShop apiInputReader.Shop,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	var wg sync.WaitGroup
	wg.Add(3)

	var pointBalanceSenderRes apiModuleRuntimesResponsesPointBalance.PointBalanceRes
	var pointBalanceReceiverRes apiModuleRuntimesResponsesPointBalance.PointBalanceRes

	var shopRes apiModuleRuntimesResponsesShop.ShopRes
	var businessPartnerGeneralResShopOwner apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes

	go func() {
		defer wg.Done()
		pointBalanceSenderRes = *controller.createPointBalanceRequestPointBalanceSender(
			controller.UserInfo,
			inputPointBalanceSender,
		)
	}()

	go func() {
		defer wg.Done()
		pointBalanceReceiverRes = *controller.createPointBalanceRequestByShopReceiver(
			controller.UserInfo,
			inputPointBalanceReceiver,
		)
	}()

	go func() {
		defer wg.Done()
		shopRes = *controller.createShopRequestHeader(
			controller.UserInfo,
			inputShop,
		)
		businessPartnerGeneralResShopOwner = *controller.createBusinessPartnerRequestGeneralShopOwner(
			controller.UserInfo,
			&shopRes,
		)
	}()

	wg.Wait()

	controller.fin(
		&pointBalanceSenderRes,
		&pointBalanceReceiverRes,
		&shopRes,
		&businessPartnerGeneralResShopOwner,
	)
}

func (
	controller *PointConsumptionController,
) fin(
	pointBalanceSenderRes *apiModuleRuntimesResponsesPointBalance.PointBalanceRes,
	pointBalanceReceiverRes *apiModuleRuntimesResponsesPointBalance.PointBalanceRes,
	shopRes *apiModuleRuntimesResponsesShop.ShopRes,
	businessPartnerGeneralResShopOwner *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {

	data := PointConsumptionGlobal{}

	for _, v := range *pointBalanceSenderRes.Message.PointBalance {

		data.PointBalancePointBalanceSender = append(data.PointBalancePointBalanceSender,
			apiOutputFormatter.PointBalancePointBalance{
				BusinessPartner: v.BusinessPartner,
				PointSymbol:     v.PointSymbol,
				CurrentBalance:  v.CurrentBalance,
				LimitBalance:    v.LimitBalance,
			},
		)
	}

	for _, v := range *pointBalanceReceiverRes.Message.ByShop {

		data.PointBalanceByShopReceiver = append(data.PointBalanceByShopReceiver,
			apiOutputFormatter.PointBalanceByShop{
				BusinessPartner: v.BusinessPartner,
				PointSymbol:     v.PointSymbol,
				Shop:            v.Shop,
				CurrentBalance:  v.CurrentBalance,
				LimitBalance:    v.LimitBalance,
			},
		)
	}

	for _, v := range *shopRes.Message.Header {

		data.ShopHeader = append(data.ShopHeader,
			apiOutputFormatter.ShopHeader{
				Shop:        v.Shop,
				ShopOwner:   v.ShopOwner,
				Description: v.Description,
			},
		)
	}

	for _, v := range *businessPartnerGeneralResShopOwner.Message.General {
		data.BusinessPartnerGeneralShopOwner = append(data.BusinessPartnerGeneralShopOwner,
			apiOutputFormatter.BusinessPartnerGeneral{
				BusinessPartner:     v.BusinessPartner,
				BusinessPartnerName: v.BusinessPartnerName,
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
