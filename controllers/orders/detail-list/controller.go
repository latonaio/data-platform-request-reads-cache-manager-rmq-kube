package controllersOrdersDetailList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/orders/orders"
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
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type OrdersDetailListController struct {
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

func (controller *OrdersDetailListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	orderID, _ := controller.GetInt("orderId")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "orders"
	redisKeyCategory2 := "detail-list"
	userType := controller.GetString(":userType") // buyer or seller
	buyerValue, _ := controller.GetInt("buyer")
	sellerValue, _ := controller.GetInt("seller")

	ordersHeader := apiInputReader.Orders{}

	headerCompleteDeliveryIsDefined := false
	headerDeliveryBlockStatus := false
	headerDeliveryStatus := "CL"
	isCancelled := false
	isMarkedForDeletion := false

	itemCompleteDeliveryIsDefined := false
	itemDeliveryBlockStatus := false
	itemDeliveryStatus := "CL"

	if userType == buyer {
		ordersHeader = apiInputReader.Orders{
			OrdersHeader: &apiInputReader.OrdersHeader{
				OrderID:                         orderID,
				Buyer:                           &buyerValue,
				HeaderCompleteDeliveryIsDefined: &headerCompleteDeliveryIsDefined,
				HeaderDeliveryBlockStatus:       &headerDeliveryBlockStatus,
				HeaderDeliveryStatus:            &headerDeliveryStatus,
				IsCancelled:                     &isCancelled,
				IsMarkedForDeletion:             &isMarkedForDeletion,
			},
			OrdersItems: &apiInputReader.OrdersItems{
				OrderID:                       orderID,
				ItemCompleteDeliveryIsDefined: &itemCompleteDeliveryIsDefined,
				ItemDeliveryBlockStatus:       &itemDeliveryBlockStatus,
				ItemDeliveryStatus:            &itemDeliveryStatus,
				IsCancelled:                   &isCancelled,
				IsMarkedForDeletion:           &isMarkedForDeletion,
			},
		}
	}

	if userType == seller {
		ordersHeader = apiInputReader.Orders{
			OrdersHeader: &apiInputReader.OrdersHeader{
				OrderID:                         orderID,
				Seller:                          &sellerValue,
				HeaderCompleteDeliveryIsDefined: &headerCompleteDeliveryIsDefined,
				HeaderDeliveryBlockStatus:       &headerDeliveryBlockStatus,
				HeaderDeliveryStatus:            &headerDeliveryStatus,
				IsCancelled:                     &isCancelled,
				IsMarkedForDeletion:             &isMarkedForDeletion,
			},
			OrdersItems: &apiInputReader.OrdersItems{
				OrderID:                       orderID,
				ItemCompleteDeliveryIsDefined: &itemCompleteDeliveryIsDefined,
				ItemDeliveryBlockStatus:       &itemDeliveryBlockStatus,
				ItemDeliveryStatus:            &itemDeliveryStatus,
				IsCancelled:                   &isCancelled,
				IsMarkedForDeletion:           &isMarkedForDeletion,
			},
		}
	}

	//ordersItems := apiInputReader.Orders{
	//	OrdersItems: &apiInputReader.OrdersItems{
	//		OrderID:      		 				orderID,
	//		ItemCompleteDeliveryIsDefined:		&itemCompleteDeliveryIsDefined,
	//		ItemDeliveryBlockStatus:			&itemDeliveryBlockStatus,
	//		ItemDeliveryStatus:					&itemDeliveryStatus,
	//		IsCancelled:						&isCancelled,
	//		IsMarkedForDeletion:	 			&isMarkedForDeletion,
	//	},
	//}

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
	controller *OrdersDetailListController,
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
	controller *OrdersDetailListController,
) createOrdersRequestItems(
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
		controller.CustomLogger.Error("OrdersReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersDetailListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*ordersRes.Message.Header))

	for _, v := range *ordersRes.Message.Header {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Buyer,
		})
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Seller,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsGeneralsByBusinessPartners(
		requestPram,
		generals,
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
	controller *OrdersDetailListController,
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
		controller.CustomLogger.Error("ProductMasterDocReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersDetailListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*ordersRes.Message.Item))

	for _, v := range *ordersRes.Message.Item {
		input = append(input, apiModuleRuntimesRequestsPlant.General{
			BusinessPartner: *v.DeliverToParty,
			Plant:           *v.DeliverToPlant,
		})
		input = append(input, apiModuleRuntimesRequestsPlant.General{
			BusinessPartner: *v.DeliverFromParty,
			Plant:           *v.DeliverFromPlant,
		})
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
		controller.CustomLogger.Error("PlantReadsGeneralsByPlants Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersDetailListController,
) request(
	input apiInputReader.Orders,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := apiModuleRuntimesResponsesOrders.OrdersRes{}
	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	if input.OrdersHeader.Buyer != nil {
		headerRes = *controller.createOrdersRequestHeader(
			controller.UserInfo,
			input,
		)

		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&headerRes,
		)
	}

	if input.OrdersHeader.Seller != nil {
		headerRes = *controller.createOrdersRequestHeader(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&headerRes,
		)
	}

	itemRes := controller.createOrdersRequestItems(
		controller.UserInfo,
		input,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		itemRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	controller.fin(
		&headerRes,
		itemRes,
		&businessPartnerRes,
		plantRes,
		productDocRes,
	)
}

func (
	controller *OrdersDetailListController,
) fin(
	headerRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	itemRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	//plantMapper := services.PlantMapper(
	//	plantRes.Message.Generals,
	//)

	data := apiOutputFormatter.Orders{}

	for _, v := range *headerRes.Message.Header {
		//img := services.ReadProductImage(
		//	productDocRes,
		//	v.Buyer,	//Sellerの対応が必要
		//	v.Product,
		//)

		data.OrdersHeaderWithItem = append(data.OrdersHeaderWithItem,
			apiOutputFormatter.OrdersHeaderWithItem{
				OrderID:      v.OrderID,
				OrderDate:    v.OrderDate,
				PaymentTerms: v.PaymentTerms,
				//PaymentTermsName:		v.PaymentTermsName, //CallerとMapperが必要
				PaymentMethod:       v.PaymentMethod,
				TransactionCurrency: v.TransactionCurrency,
				OrderType:           v.OrderType,
				Buyer:               v.Buyer,
				BuyerName:           businessPartnerMapper[v.Buyer].BusinessPartnerName,
				Seller:              v.Seller,
				SellerName:          businessPartnerMapper[v.Seller].BusinessPartnerName,
			},
		)
	}

	for _, v := range *itemRes.Message.Item {
		data.OrdersItem = append(data.OrdersItem,
			apiOutputFormatter.OrdersItem{
				OrderItem:                   v.OrderItem,
				Product:                     v.Product,
				OrderItemTextByBuyer:        v.OrderItemTextByBuyer,
				OrderItemTextBySeller:       v.OrderItemTextBySeller,
				OrderQuantityInDeliveryUnit: v.OrderQuantityInDeliveryUnit,
				DeliveryUnit:                v.DeliveryUnit,
				RequestedDeliveryDate:       v.RequestedDeliveryDate,
				NetAmount:                   v.NetAmount,
				IsCancelled:                 v.IsCancelled,
				IsMarkedForDeletion:         v.IsMarkedForDeletion,
				//Images: apiOutputFormatter.Images{
				//	Product: img,
				//},
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
