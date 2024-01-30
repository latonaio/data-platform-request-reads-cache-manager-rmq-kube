package controllersOrdersItemSingleUnitMillSheet

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsMillSheetPdf "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/function-mill-sheet-pdf"
	apiModuleRuntimesRequestsInspectionLot "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/inspection-lot/inspection-lot"
	apiModuleRuntimesRequestsOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/orders/orders"
	apiModuleRuntimesRequestsOrdersDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/orders/orders-doc"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesInspectionLot "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/inspection-lot"
	apiModuleRuntimesResponsesMillSheetPdf "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/mill-sheet-pdf"
	apiModuleRuntimesResponsesOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/orders"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type OrdersItemSingleUnitMillSheetController struct {
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

func (controller *OrdersItemSingleUnitMillSheetController) Get() {
	//isReleased, _ := controller.GetBool("isReleased")
	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "orders"
	redisKeyCategory2 := "orders-item-single-unit-mill-sheet"
	orderId, _ := controller.GetInt("orderId")
	orderItem, _ := controller.GetInt("orderItem")
	userType := controller.GetString(":userType")
	pBuyer, _ := controller.GetInt("buyer")
	pSeller, _ := controller.GetInt("seller")

	OrdersItemSingleUnitMillSheet := apiInputReader.Orders{}

	headerCompleteDeliveryIsDefined := false
	headerDeliveryBlockStatus := false
	headerDeliveryStatus := "CL"
	isCancelled := false
	isMarkedForDeletion := false

	itemCompleteDeliveryIsDefined := false
	itemDeliveryBlockStatus := false
	itemDeliveryStatus := "NP"

	docType := "QRCODE"

	if userType == buyer {
		OrdersItemSingleUnitMillSheet = apiInputReader.Orders{
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
			OrdersDocItemDoc: &apiInputReader.OrdersDocItemDoc{
				OrderID:                  orderId,
				OrderItem:                &orderItem,
				DocType:                  &docType,
				DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
			},
		}
	} else {
		OrdersItemSingleUnitMillSheet = apiInputReader.Orders{
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
			OrdersDocItemDoc: &apiInputReader.OrdersDocItemDoc{
				OrderID:                  orderId,
				OrderItem:                &orderItem,
				DocType:                  &docType,
				DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
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

	//cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)
	//
	//if cacheData != nil {
	//	var responseData apiOutputFormatter.Orders
	//
	//	err := json.Unmarshal(cacheData, &responseData)
	//
	//	if err != nil {
	//		services.HandleError(
	//			&controller.Controller,
	//			err,
	//			nil,
	//		)
	//	}
	//
	//	services.Respond(
	//		&controller.Controller,
	//		&responseData,
	//	)
	//}

	//if cacheData != nil {
	//	go func() {
	//		controller.request(OrdersItemSingleUnitMillSheet)
	//	}()
	//} else {
	//	controller.request(OrdersItemSingleUnitMillSheet)
	//}

	controller.request(OrdersItemSingleUnitMillSheet)
}

func (
	controller *OrdersItemSingleUnitMillSheetController,
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
	controller *OrdersItemSingleUnitMillSheetController,
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
	controller *OrdersItemSingleUnitMillSheetController,
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
	controller *OrdersItemSingleUnitMillSheetController,
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
	controller *OrdersItemSingleUnitMillSheetController,
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
	controller *OrdersItemSingleUnitMillSheetController,
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
	controller *OrdersItemSingleUnitMillSheetController,
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
	controller *OrdersItemSingleUnitMillSheetController,
) createInspectionLotRequestHeader(
	requestPram *apiInputReader.Request,
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes {
	input := apiInputReader.InspectionLot{
		InspectionLotHeader: &apiInputReader.InspectionLotHeader{
			InspectionLot: *(*ordersRes.Message.Item)[0].InspectionLot,
		},
	}

	responseJsonData := apiModuleRuntimesResponsesInspectionLot.InspectionLotRes{}
	responseBody := apiModuleRuntimesRequestsInspectionLot.InspectionLotReads(
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
		controller.CustomLogger.Error("createInspectionLotRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersItemSingleUnitMillSheetController,
) createInspectionLotRequestSpecDetails(
	requestPram *apiInputReader.Request,
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes {
	input := apiInputReader.InspectionLot{
		InspectionLotSpecDetail: &apiInputReader.InspectionLotSpecDetail{
			InspectionLot: *(*ordersRes.Message.Item)[0].InspectionLot,
		},
	}

	responseJsonData := apiModuleRuntimesResponsesInspectionLot.InspectionLotRes{}
	responseBody := apiModuleRuntimesRequestsInspectionLot.InspectionLotReads(
		requestPram,
		input,
		&controller.Controller,
		"SpecDetails",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createInspectionLotRequestSpecDetails Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersItemSingleUnitMillSheetController,
) createInspectionLotRequestComponentCompositions(
	requestPram *apiInputReader.Request,
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes {
	input := apiInputReader.InspectionLot{
		InspectionLotComponentComposition: &apiInputReader.InspectionLotComponentComposition{
			InspectionLot: *(*ordersRes.Message.Item)[0].InspectionLot,
		},
	}
	responseJsonData := apiModuleRuntimesResponsesInspectionLot.InspectionLotRes{}
	responseBody := apiModuleRuntimesRequestsInspectionLot.InspectionLotReads(
		requestPram,
		input,
		&controller.Controller,
		"ComponentCompositions",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createInspectionLotRequestComponentCompositions Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersItemSingleUnitMillSheetController,
) createInspectionLotRequestInspections(
	requestPram *apiInputReader.Request,
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes {
	input := apiInputReader.InspectionLot{
		InspectionLotInspection: &apiInputReader.InspectionLotInspection{
			InspectionLot: *(*ordersRes.Message.Item)[0].InspectionLot,
		},
	}

	responseJsonData := apiModuleRuntimesResponsesInspectionLot.InspectionLotRes{}
	responseBody := apiModuleRuntimesRequestsInspectionLot.InspectionLotReads(
		requestPram,
		input,
		&controller.Controller,
		"Inspections",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createInspectionLotRequestInspections Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersItemSingleUnitMillSheetController,
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

	//ordersItemScheduleLinesRes := controller.createOrdersRequestItemScheduleLines(
	//	controller.UserInfo,
	//	input,
	//)
	//
	//ordersItemPricingElementsRes := controller.createOrdersRequestItemPricingElements(
	//	controller.UserInfo,
	//	input,
	//)

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

	inspectionLotHeaderRes := controller.createInspectionLotRequestHeader(
		controller.UserInfo,
		ordersItemRes,
	)

	inspectionLotSpecDetailsRes := controller.createInspectionLotRequestSpecDetails(
		controller.UserInfo,
		ordersItemRes,
	)

	inspectionLotComponentCompositionsRes := controller.createInspectionLotRequestComponentCompositions(
		controller.UserInfo,
		ordersItemRes,
	)

	inspectionLotInspectionsRes := controller.createInspectionLotRequestInspections(
		controller.UserInfo,
		ordersItemRes,
	)

	controller.fin(
		&ordersHeaderRes,
		ordersItemRes,
		//ordersItemScheduleLinesRes,
		//ordersItemPricingElementsRes,
		&businessPartnerRes,
		productDocRes,
		ordersItemDocRes,
		inspectionLotHeaderRes,
		inspectionLotSpecDetailsRes,
		inspectionLotComponentCompositionsRes,
		inspectionLotInspectionsRes,
	)
}

func (
	controller *OrdersItemSingleUnitMillSheetController,
) fin(
	ordersHeaderRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	ordersItemRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	//ordersItemScheduleLinesRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	//ordersItemPricingElementsRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	ordersItemDocRes *apiModuleRuntimesResponsesOrders.OrdersDocRes,
	inspectionLotHeaderRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	inspectionLotSpecDetailsRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	inspectionLotComponentCompositionsRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	inspectionLotInspectionsRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.Orders{}

	for _, v := range *ordersItemRes.Message.Item {
		//img := services.ReadProductImage(
		//	productDocRes,
		//	*controller.UserInfo.BusinessPartner,
		//	v.Product,
		//)

		var orderType *string
		//var conditionCurrency *string

		if ordersHeaderRes != nil && ordersHeaderRes.Message.Header != nil && len(*ordersHeaderRes.Message.Header) > 0 {
			orderType = &(*ordersHeaderRes.Message.Header)[0].OrderType
		} else {
			orderType = nil
		}

		//if ordersItemScheduleLinesRes != nil && ordersItemScheduleLinesRes.Message.ItemScheduleLine != nil && len(*ordersItemScheduleLinesRes.Message.ItemScheduleLine) > 0 {
		//	requestedDeliveryDate = &(*ordersItemScheduleLinesRes.Message.ItemScheduleLine)[0].RequestedDeliveryDate
		//	requestedDeliveryTime = &(*ordersItemScheduleLinesRes.Message.ItemScheduleLine)[0].RequestedDeliveryTime
		//} else {
		//	requestedDeliveryDate = nil
		//	requestedDeliveryTime = nil
		//}

		//if ordersItemPricingElementsRes != nil && ordersItemPricingElementsRes.Message.ItemPricingElement != nil && len(*ordersItemPricingElementsRes.Message.ItemPricingElement) > 0 {
		//	conditionCurrency = (*ordersItemPricingElementsRes.Message.ItemPricingElement)[0].ConditionCurrency
		//} else {
		//	conditionCurrency = nil
		//}

		//qrcode := services.CreateQRCodeOrdersItemDocImage(
		//	ordersItemDocRes,
		//	v.OrderID,
		//	v.OrderItem,
		//)

		data.OrdersItemSingleUnitMillSheetHeader = append(data.OrdersItemSingleUnitMillSheetHeader,
			apiOutputFormatter.OrdersItemSingleUnitMillSheetHeader{
				OrderID:                 v.OrderID,
				OrderItem:               v.OrderItem,
				OrderType:               orderType,
				OrderStatus:             v.OrderStatus,
				Buyer:                   v.Buyer,
				BuyerName:               businessPartnerMapper[v.Buyer].BusinessPartnerName,
				Seller:                  v.Seller,
				SellerName:              businessPartnerMapper[v.Seller].BusinessPartnerName,
				Product:                 v.Product,
				SizeOrDimensionText:     *v.SizeOrDimensionText,
				OrderItemTextByBuyer:    v.OrderItemTextByBuyer,
				OrderItemTextBySeller:   v.OrderItemTextBySeller,
				OrderQuantityInBaseUnit: v.OrderQuantityInBaseUnit,
				RequestedDeliveryDate:   v.RequestedDeliveryDate,
				RequestedDeliveryTime:   v.RequestedDeliveryTime,
				ProductSpecification:    *v.ProductSpecification,
				MarkingOfMaterial:       *v.MarkingOfMaterial,
				ProductionVersion:       v.ProductionVersion,
				ProductionVersionItem:   v.ProductionVersionItem,
				ProductionOrder:         v.ProductionOrder,
				ProductionOrderItem:     v.ProductionOrderItem,
				Contract:                v.Contract,
				ContractItem:            v.ContractItem,
				Project:                 v.Project,
				WBSElement:              v.WBSElement,
				GrossAmount:             v.GrossAmount,
				//ConditionCurrency:       v.ConditionCurrency,
				InspectionLot: *v.InspectionLot,

				//Images: apiOutputFormatter.Images{
				//	Product: img,
				//	QRCode:  qrcode,
				//},
			},
		)
	}

	for _, v := range *inspectionLotHeaderRes.Message.Header {
		data.OrdersItemSingleUnitMillSheetHeaderInspectionLot = append(data.OrdersItemSingleUnitMillSheetHeaderInspectionLot,
			apiOutputFormatter.OrdersItemSingleUnitMillSheetHeaderInspectionLot{
				OrderID:                 (*ordersItemRes.Message.Item)[0].OrderID,
				OrderItem:               (*ordersItemRes.Message.Item)[0].OrderItem,
				InspectionLot:           *(*ordersItemRes.Message.Item)[0].InspectionLot,
				InspectionLotDate:       v.InspectionLotDate,
				InspectionSpecification: v.InspectionSpecification,
			},
		)
	}

	for _, v := range *inspectionLotSpecDetailsRes.Message.SpecDetail {
		data.OrdersItemSingleUnitMillSheetSpecDetails = append(data.OrdersItemSingleUnitMillSheetSpecDetails,
			apiOutputFormatter.OrdersItemSingleUnitMillSheetSpecDetails{
				OrderID:         (*ordersItemRes.Message.Item)[0].OrderID,
				OrderItem:       (*ordersItemRes.Message.Item)[0].OrderItem,
				InspectionLot:   *(*ordersItemRes.Message.Item)[0].InspectionLot,
				SpecType:        v.SpecType,
				UpperLimitValue: &v.UpperLimitValue,
				LowerLimitValue: &v.LowerLimitValue,
				StandardValue:   &v.StandardValue,
				SpecTypeUnit:    &v.SpecTypeUnit,
			},
		)
	}

	for _, v := range *inspectionLotComponentCompositionsRes.Message.ComponentComposition {
		data.OrdersItemSingleUnitMillSheetComponentCompositions = append(data.OrdersItemSingleUnitMillSheetComponentCompositions,
			apiOutputFormatter.OrdersItemSingleUnitMillSheetComponentCompositions{
				OrderID:                                    (*ordersItemRes.Message.Item)[0].OrderID,
				OrderItem:                                  (*ordersItemRes.Message.Item)[0].OrderItem,
				InspectionLot:                              *(*ordersItemRes.Message.Item)[0].InspectionLot,
				ComponentCompositionType:                   v.ComponentCompositionType,
				ComponentCompositionUpperLimitInPercent:    &v.ComponentCompositionUpperLimitInPercent,
				ComponentCompositionLowerLimitInPercent:    &v.ComponentCompositionLowerLimitInPercent,
				ComponentCompositionStandardValueInPercent: &v.ComponentCompositionStandardValueInPercent,
			},
		)
	}
	for _, v := range *inspectionLotInspectionsRes.Message.Inspection {
		data.OrdersItemSingleUnitMillSheetInspections = append(data.OrdersItemSingleUnitMillSheetInspections,
			apiOutputFormatter.OrdersItemSingleUnitMillSheetInspections{
				OrderID:                                  (*ordersItemRes.Message.Item)[0].OrderID,
				OrderItem:                                (*ordersItemRes.Message.Item)[0].OrderItem,
				InspectionLot:                            *(*ordersItemRes.Message.Item)[0].InspectionLot,
				Inspection:                               v.Inspection,
				InspectionType:                           v.InspectionType,
				InspectionTypeCertificateValueInText:     v.InspectionTypeCertificateValueInText,
				InspectionTypeCertificateValueInQuantity: v.InspectionTypeCertificateValueInQuantity,
				InspectionTypeValueUnit:                  v.InspectionTypeValueUnit,
			},
		)
	}

	// ここから generates に rabbitmq で送信
	// accepter 対応
	responseJsonData := apiModuleRuntimesResponsesMillSheetPdf.MillSheetPdfRes{}
	responseBody := apiModuleRuntimesRequestsMillSheetPdf.FunctionMillSheetPdfGeneratesGenerates(
		data,
		&controller.Controller,
		"MillSheet",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("apiModuleRuntimesRequestsMillSheetPdf.FunctionMillSheetPdfGeneratesGenerates Unmarshal error")
	}

	data.MillSheetPdfMountPath = responseJsonData.MountPath

	//err = controller.RedisCache.SetCache(
	//	controller.RedisKey,
	//	data,
	//)
	//if err != nil {
	//	services.HandleError(
	//		&controller.Controller,
	//		err,
	//		nil,
	//	)
	//}

	controller.Data["json"] = &data
	controller.ServeJSON()
}
