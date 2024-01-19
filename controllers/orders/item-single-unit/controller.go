package controllersOrdersSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/orders/orders"
	apiModuleRuntimesRequestsOrdersDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/orders/orders-doc"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/orders"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type OrdersSingleUnitController struct {
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

func (controller *OrdersSingleUnitController) Get() {
	//isReleased, _ := controller.GetBool("isReleased")
	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "orders"
	redisKeyCategory2 := "orders-item-single-unit"
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
			OrdersItem: &apiInputReader.OrdersItem{
				OrderID:                       orderId,
				OrderItem:                     orderItem,
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
			OrdersDocItemDoc: &apiInputReader.OrdersDocItemDoc{
				OrderID:                  orderId,
				OrderItem:                orderItem,
				DocType:                  "QRCODE",
				DocIssuerBusinessPartner: *controller.UserInfo.BusinessPartner,
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
			OrdersItem: &apiInputReader.OrdersItem{
				OrderID:                       orderId,
				OrderItem:                     orderItem,
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
			OrdersDocItemDoc: &apiInputReader.OrdersDocItemDoc{
				OrderID:                  orderId,
				OrderItem:                orderItem,
				DocType:                  "QRCODE",
				DocIssuerBusinessPartner: *controller.UserInfo.BusinessPartner,
			},
		}
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
	controller *OrdersSingleUnitController,
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
	controller *OrdersSingleUnitController,
) createOrdersRequestItem(
	requestPram *apiInputReader.Request,
	input apiInputReader.Orders,
) *apiModuleRuntimesResponsesOrders.OrdersRes {
	responseJsonData := apiModuleRuntimesResponsesOrders.OrdersRes{}
	responseBody := apiModuleRuntimesRequestsOrders.OrdersReads(
		requestPram,
		input,
		&controller.Controller,
		"Item",
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
	controller *OrdersSingleUnitController,
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
	controller *OrdersSingleUnitController,
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
	controller *OrdersSingleUnitController,
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
	controller *OrdersSingleUnitController,
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
	controller *OrdersSingleUnitController,
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
	controller *OrdersSingleUnitController,
) request(
	input apiInputReader.Orders,
) {
	defer services.Recover(controller.CustomLogger)

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

	ordersItemPricingElementsRes := controller.createOrdersRequestItemPricingElements(
		controller.UserInfo,
		input,
	)

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		ordersItemRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	ordersItemDocRes := controller.createOrdersDocRequest(
		controller.UserInfo,
		input,
	)

	controller.fin(
		&ordersHeaderRes,
		ordersItemRes,
		ordersItemScheduleLinesRes,
		ordersItemPricingElementsRes,
		&businessPartnerRes,
		productDocRes,
		ordersItemDocRes,
	)
}

func (
	controller *OrdersSingleUnitController,
) fin(
	ordersHeaderRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	ordersItemRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	ordersItemScheduleLinesRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	ordersItemPricingElementsRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	ordersItemDocRes *apiModuleRuntimesResponsesOrders.OrdersDocRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.Orders{}

	for _, v := range *ordersItemRes.Message.Item {
		img := services.ReadProductImage(
			productDocRes,
			*controller.UserInfo.BusinessPartner,
			v.Product,
		)

		var orderType *string
		var transactionCurrency *string

		if ordersHeaderRes != nil && ordersHeaderRes.Message.Header != nil && len(*ordersHeaderRes.Message.Header) > 0 {
			orderType = &(*ordersHeaderRes.Message.Header)[0].OrderType
			transactionCurrency = &(*ordersHeaderRes.Message.Header)[0].TransactionCurrency
		} else {
			orderType = nil
			transactionCurrency = nil
		}

		qrcode := services.CreateQRCodeOrdersItemDocImage(
			ordersItemDocRes,
			v.OrderID,
			v.OrderItem,
		)

		data.OrdersSingleUnit = append(data.OrdersSingleUnit,
			apiOutputFormatter.OrdersSingleUnit{
				OrderID:               v.OrderID,
				OrderItem:             v.OrderItem,
				OrderStatus:           v.OrderStatus,
				Buyer:                 v.Buyer,
				OrderType:             orderType,
				BuyerName:             businessPartnerMapper[v.Buyer].BusinessPartnerName,
				Seller:                v.Seller,
				SellerName:            businessPartnerMapper[v.Seller].BusinessPartnerName,
				Product:               v.Product,
				OrderItemTextByBuyer:  v.OrderItemTextByBuyer,
				OrderItemTextBySeller: v.OrderItemTextBySeller,
				GrossAmount:           v.GrossAmount,
				TransactionCurrency:   *transactionCurrency,
				RequestedDeliveryDate: v.RequestedDeliveryDate,
				RequestedDeliveryTime: v.RequestedDeliveryTime,

				Images: apiOutputFormatter.Images{
					Product: img,
					QRCode:  qrcode,
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
