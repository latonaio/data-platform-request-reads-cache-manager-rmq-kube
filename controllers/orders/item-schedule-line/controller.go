package controllersOrdersItemScheduleLine

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/orders/orders"
	apiModuleRuntimesRequestsOrdersDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/orders/orders-doc"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/orders"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type OrdersItemScheduleLineController struct {
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

func (controller *OrdersItemScheduleLineController) Get() {
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "orders"
	redisKeyCategory2 := "orders-item-schedule-line"
	orderId, _ := controller.GetInt("orderId")
	orderItem, _ := controller.GetInt("orderItem")
	userType := controller.GetString(":userType")
	pBuyer, _ := controller.GetInt("buyer")
	pSeller, _ := controller.GetInt("seller")

	OrdersSingleUnit := apiInputReader.Orders{}

	headerCompleteDeliveryIsDefined := false
	headerDeliveryBlockStatus := false
	headerDeliveryStatus := "CL"
	isCancelled := false
	isMarkedForDeletion := false

	itemCompleteDeliveryIsDefined := false
	itemDeliveryBlockStatus := false
	itemDeliveryStatus := "NP"

	if userType == buyer {
		OrdersSingleUnit = apiInputReader.Orders{
			OrdersHeader: &apiInputReader.OrdersHeader{
				OrderID:                         orderId,
				Buyer:                           &pBuyer,
				Seller:                          &pSeller,
				HeaderCompleteDeliveryIsDefined: &headerCompleteDeliveryIsDefined,
				HeaderDeliveryBlockStatus:       &headerDeliveryBlockStatus,
				HeaderDeliveryStatus:            &headerDeliveryStatus,
				IsCancelled:                     &isCancelled,
				IsMarkedForDeletion:             &isMarkedForDeletion,
			},
			OrdersItems: &apiInputReader.OrdersItems{
				OrderID:                       orderId,
				OrderItem:                     &orderItem,
				ItemCompleteDeliveryIsDefined: &itemCompleteDeliveryIsDefined,
				ItemDeliveryBlockStatus:       &itemDeliveryBlockStatus,
				ItemDeliveryStatus:            &itemDeliveryStatus,
				IsCancelled:                   &isCancelled,
				IsMarkedForDeletion:           &isMarkedForDeletion,
			},
			OrdersItemScheduleLines: &apiInputReader.OrdersItemScheduleLines{
				OrderID:   orderId,
				OrderItem: orderItem,
			},
		}
	} else {
		OrdersSingleUnit = apiInputReader.Orders{
			OrdersHeader: &apiInputReader.OrdersHeader{
				OrderID:                         orderId,
				Buyer:                           &pBuyer,
				Seller:                          &pSeller,
				HeaderCompleteDeliveryIsDefined: &headerCompleteDeliveryIsDefined,
				HeaderDeliveryBlockStatus:       &headerDeliveryBlockStatus,
				HeaderDeliveryStatus:            &headerDeliveryStatus,
				IsCancelled:                     &isCancelled,
				IsMarkedForDeletion:             &isMarkedForDeletion,
			},
			OrdersItems: &apiInputReader.OrdersItems{
				OrderID:                       orderId,
				OrderItem:                     &orderItem,
				ItemCompleteDeliveryIsDefined: &itemCompleteDeliveryIsDefined,
				ItemDeliveryBlockStatus:       &itemDeliveryBlockStatus,
				ItemDeliveryStatus:            &itemDeliveryStatus,
				IsCancelled:                   &isCancelled,
				IsMarkedForDeletion:           &isMarkedForDeletion,
			},
			OrdersItemScheduleLines: &apiInputReader.OrdersItemScheduleLines{
				OrderID:   orderId,
				OrderItem: orderItem,
			},
		}
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(orderId),
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
			controller.request(OrdersSingleUnit)
		}()
	} else {
		controller.request(OrdersSingleUnit)
	}
}

func (
	controller *OrdersItemScheduleLineController,
) createOrdersRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.Orders,
) *apiModuleRuntimesResponsesOrders.OrdersRes {
	responseJsonData := apiModuleRuntimesResponsesOrders.OrdersRes{}
	responseBody := apiModuleRuntimesRequestsOrders.OrdersReads(
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
		controller.CustomLogger.Error("OrdersReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersItemScheduleLineController,
) createOrdersRequestItem(
	requestPram *apiInputReader.Request,
	input apiInputReader.Orders,
) *apiModuleRuntimesResponsesOrders.OrdersRes {
	responseJsonData := apiModuleRuntimesResponsesOrders.OrdersRes{}
	responseBody := apiModuleRuntimesRequestsOrders.OrdersReads(
		requestPram,
		input,
		&controller.Controller,
		"Items",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createOrdersRequestItem Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersItemScheduleLineController,
) createOrdersRequestItemScheduleLines(
	requestPram *apiInputReader.Request,
	input apiInputReader.Orders,
) *apiModuleRuntimesResponsesOrders.OrdersRes {
	responseJsonData := apiModuleRuntimesResponsesOrders.OrdersRes{}
	responseBody := apiModuleRuntimesRequestsOrders.OrdersReads(
		requestPram,
		input,
		&controller.Controller,
		"ItemScheduleLines",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createOrdersRequestItemScheduleLines Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersItemScheduleLineController,
) createOrdersRequestItemPricingElements(
	requestPram *apiInputReader.Request,
	input apiInputReader.Orders,
) *apiModuleRuntimesResponsesOrders.OrdersRes {
	responseJsonData := apiModuleRuntimesResponsesOrders.OrdersRes{}
	responseBody := apiModuleRuntimesRequestsOrders.OrdersReads(
		requestPram,
		input,
		&controller.Controller,
		"ItemPricingElements",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createOrdersRequestItemPricingElements Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersItemScheduleLineController,
) createOrdersDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.Orders,
) *apiModuleRuntimesResponsesOrders.OrdersDocRes {
	responseJsonData := apiModuleRuntimesResponsesOrders.OrdersDocRes{}
	responseBody := apiModuleRuntimesRequestsOrdersDoc.OrdersDocReads(
		requestPram,
		input,
		&controller.Controller,
		"OrdersDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createOrdersDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersItemScheduleLineController,
) createProductMasterDocRequest(
	requestPram *apiInputReader.Request,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes{}
	responseBody := apiModuleRuntimesRequestsProductMasterDoc.ProductMasterDocReads(
		requestPram,
		&controller.Controller,
		"GeneralDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductMasterDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersItemScheduleLineController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	ordersItemRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*ordersItemRes.Message.Item))

	for _, v := range *ordersItemRes.Message.Item {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Buyer,
		})
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Seller,
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
		controller.CustomLogger.Error("BusinessPartnerGeneralReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersItemScheduleLineController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	ordersItemScheduleLineRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*ordersItemScheduleLineRes.Message.ItemScheduleLine))
	for i, v := range *ordersItemScheduleLineRes.Message.ItemScheduleLine {
		input[i].Plant = v.StockConfirmationPlant
	}

	responseJsonData := apiModuleRuntimesResponsesPlant.PlantRes{}
	responseBody := apiModuleRuntimesRequestsPlant.PlantReadsGeneralsByPlants(
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
		controller.CustomLogger.Error("createPlantRequestGenerals Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersItemScheduleLineController,
) request(
	input apiInputReader.Orders,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	ordersHeaderRes := *controller.createOrdersRequestHeader(
		controller.UserInfo,
		input,
	)

	ordersItemRes := controller.createOrdersRequestItem(
		controller.UserInfo,
		input,
	)

	ordersItemScheduleLinesRes := controller.createOrdersRequestItemScheduleLines(
		controller.UserInfo,
		input,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		ordersItemScheduleLinesRes,
	)

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		ordersItemRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	controller.fin(
		&ordersHeaderRes,
		ordersItemRes,
		ordersItemScheduleLinesRes,
		&businessPartnerRes,
		plantRes,
		productDocRes,
	)
}

func (
	controller *OrdersItemScheduleLineController,
) fin(
	ordersHeaderRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	ordersItemRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	ordersItemScheduleLinesRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	data := apiOutputFormatter.Orders{}

	for _, v := range *ordersItemRes.Message.Item {
		img := services.ReadProductImage(
			productDocRes,
			*controller.UserInfo.BusinessPartner,
			v.Product,
		)

		var requestedDeliveryDate *string
		var requestedDeliveryTime *string

		if ordersHeaderRes != nil && ordersHeaderRes.Message.Header != nil && len(*ordersHeaderRes.Message.Header) > 0 {
			requestedDeliveryDate = &(*ordersHeaderRes.Message.Header)[0].RequestedDeliveryDate
			requestedDeliveryTime = &(*ordersHeaderRes.Message.Header)[0].RequestedDeliveryTime
		} else {
			requestedDeliveryDate = nil
			requestedDeliveryTime = nil
		}

		data.OrdersItem = append(data.OrdersItem, apiOutputFormatter.OrdersItem{
			OrderItem:                   v.OrderItem,
			Product:                     v.Product,
			OrderItemTextByBuyer:        v.OrderItemTextByBuyer,
			OrderItemTextBySeller:       v.OrderItemTextBySeller,
			Buyer:                       v.Buyer,
			BuyerName:                   businessPartnerMapper[v.Buyer].BusinessPartnerName,
			Seller:                      v.Seller,
			SellerName:                  businessPartnerMapper[v.Seller].BusinessPartnerName,
			DeliveryUnit:                v.DeliveryUnit,
			OrderQuantityInDeliveryUnit: v.OrderQuantityInDeliveryUnit,
			RequestedDeliveryDate:       *requestedDeliveryDate,
			RequestedDeliveryTime:       *requestedDeliveryTime,
			Images: apiOutputFormatter.Images{
				Product: img,
			},
		})
	}

	for _, v := range *ordersItemScheduleLinesRes.Message.ItemScheduleLine {
		var pBuyer *int
		var pSeller *int

		if ordersItemRes != nil && ordersItemRes.Message.Item != nil && len(*ordersItemRes.Message.Item) > 0 {
			pBuyer = &(*ordersItemRes.Message.Item)[0].Buyer
			pSeller = &(*ordersItemRes.Message.Item)[0].Seller
		}

		stockConfirmationPlantName := fmt.Sprintf("%s", plantMapper[v.StockConfirmationPlant].PlantName)

		data.OrdersItemScheduleLine = append(data.OrdersItemScheduleLine,
			apiOutputFormatter.OrdersItemScheduleLine{
				OrderID:                             v.OrderID,
				OrderItem:                           v.OrderItem,
				ScheduleLine:                        v.ScheduleLine,
				Buyer:                               pBuyer,
				BuyerName:                           businessPartnerMapper[*pBuyer].BusinessPartnerName,
				Seller:                              pSeller,
				SellerName:                          businessPartnerMapper[*pSeller].BusinessPartnerName,
				Product:                             v.Product,
				RequestedDeliveryDate:               v.RequestedDeliveryDate,
				RequestedDeliveryTime:               v.RequestedDeliveryTime,
				ScheduleLineOrderQuantityInBaseUnit: v.ScheduleLineOrderQuantityInBaseUnit,
				ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit: v.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit,
				StockConfirmationBusinessPartner:                v.StockConfirmationBusinessPartner,
				StockConfirmationBusinessPartnerName:            businessPartnerMapper[v.StockConfirmationBusinessPartner].BusinessPartnerName,
				StockConfirmationPlant:                          v.StockConfirmationPlant,
				StockConfirmationPlantName:                      stockConfirmationPlantName,
				DeliveredQuantityInBaseUnit:                     v.DeliveredQuantityInBaseUnit,
				UndeliveredQuantityInBaseUnit:                   v.UndeliveredQuantityInBaseUnit,
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
