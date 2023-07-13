package controllersOrdersList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/orders"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/orders"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type OrdersListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const (
	buyer  = "buyer"
	seller = "seller"
)

func (controller *OrdersListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "orders"
	redisKeyCategory2 := "list"
	userType := controller.GetString("userType") // buyer or seller
	buyerValue, _ := controller.GetInt("buyer")
	sellerValue, _ := controller.GetInt("seller")

	ordersHeader := apiInputReader.Orders{}

	if userType == buyer {
		ordersHeader = apiInputReader.Orders{
			OrdersHeader: &apiInputReader.OrdersHeader{
				Buyer:               &buyerValue,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		}
	}

	if userType == seller {
		ordersHeader = apiInputReader.Orders{
			OrdersHeader: &apiInputReader.OrdersHeader{
				Seller:              &sellerValue,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		}
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			userType,
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.Orders

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
			controller.request(ordersHeader)
		}()
	} else {
		controller.request(ordersHeader)
	}
}

func (
	controller *OrdersListController,
) createOrdersRequestHeaderByBuyer(
	requestPram *apiInputReader.Request,
	input apiInputReader.Orders,
) *apiModuleRuntimesResponsesOrders.OrdersRes {
	responseJsonData := apiModuleRuntimesResponsesOrders.OrdersRes{}
	responseBody := apiModuleRuntimesRequestsOrders.OrdersReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersByBuyer",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("OrdersReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersListController,
) createOrdersRequestHeaderBySeller(
	requestPram *apiInputReader.Request,
	input apiInputReader.Orders,
) *apiModuleRuntimesResponsesOrders.OrdersRes {
	responseJsonData := apiModuleRuntimesResponsesOrders.OrdersRes{}
	responseBody := apiModuleRuntimesRequestsOrders.OrdersReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersBySeller",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("OrdersReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersListController,
) createBusinessPartnerRequestByBuyer(
	requestPram *apiInputReader.Request,
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, 0)

	for _, v := range *ordersRes.Message.Header {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Buyer,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReads(
		requestPram,
		generals,
		&controller.Controller,
		"GeneralsByBusinessPartners",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("BusinessPartnerReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersListController,
) createBusinessPartnerRequestBySeller(
	requestPram *apiInputReader.Request,
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, 0)

	for _, v := range *ordersRes.Message.Header {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Seller,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReads(
		requestPram,
		generals,
		&controller.Controller,
		"GeneralsByBusinessPartners",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("BusinessPartnerReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersListController,
) request(
	input apiInputReader.Orders,
) {
	defer services.Recover(controller.CustomLogger)

	ordersRes := apiModuleRuntimesResponsesOrders.OrdersRes{}
	//productMasterRes := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	if input.OrdersHeader.Buyer != nil {
		ordersRes = *controller.createOrdersRequestHeaderByBuyer(
			controller.UserInfo,
			input,
		)
		//productMasterRes = *controller.createProductMasterRequestProductDescByBPByBuyer(
		//	controller.UserInfo,
		//	&ordersRes,
		//)
		businessPartnerRes = *controller.createBusinessPartnerRequestByBuyer(
			controller.UserInfo,
			&ordersRes,
		)
	}

	if input.OrdersHeader.Seller != nil {
		ordersRes = *controller.createOrdersRequestHeaderBySeller(
			controller.UserInfo,
			input,
		)
		//productMasterRes = *controller.createProductMasterRequestProductDescByBPBySeller(
		//	controller.UserInfo,
		//	&ordersRes,
		//)
		businessPartnerRes = *controller.createBusinessPartnerRequestBySeller(
			controller.UserInfo,
			&ordersRes,
		)
	}

	//pMDocRes := controller.createProductMasterDocRequest(
	//	controller.UserInfo,
	//)

	controller.fin(
		&ordersRes,
		//&productMasterRes,
		&businessPartnerRes,
		//pMDocRes,
	)
}

func (
	controller *OrdersListController,
) fin(
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	//productMasterRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	// pMDocRes *apiModuleRuntimesResponses.ProductMasterDocRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.Orders{}

	for _, v := range *ordersRes.Message.Header {

		data.OrdersHeader = append(data.OrdersHeader,
			apiOutputFormatter.OrdersHeader{
				OrderID:				   v.OrderID,
				Buyer:                     v.Buyer,
				BuyerName:                 businessPartnerMapper[v.Buyer].BusinessPartnerName,
				Seller:                    v.Seller,
				SellerName:                businessPartnerMapper[v.Seller].BusinessPartnerName,
				HeaderDeliveryStatus	   v.HeaderDeliveryStatus,
				IsCancelled				   v.IsCancelled,
				IsMarkedForDeletion		   v.IsMarkedForDeletion,
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
