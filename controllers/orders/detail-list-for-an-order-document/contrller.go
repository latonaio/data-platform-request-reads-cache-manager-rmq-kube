package controllersOrdersDetailListForAnOrderDocument

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsIncoterms "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/incoterms"
	apiModuleRuntimesRequestsOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/orders/orders"
	apiModuleRuntimesRequestsPaymentTerms "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/payment-terms"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProject "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/project"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesIncoterms "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/incoterms"
	apiModuleRuntimesResponsesOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/orders"
	apiModuleRuntimesResponsesPaymentTerms "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/payment-terms"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProject "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/project"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"io/ioutil"
	"strconv"
	"strings"
)

type OrdersDetailListForAnOrderDocumentController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *OrdersDetailListForAnOrderDocumentController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	orderID, _ := controller.GetInt("orderId")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "orders"
	redisKeyCategory2 := "detail-list-for-an-order-document"

	ordersHeader := apiInputReader.Orders{}

	ordersHeader = apiInputReader.Orders{
		OrdersHeader: &apiInputReader.OrdersHeader{
			OrderID: orderID,
		},
		OrdersItems: &apiInputReader.OrdersItems{
			OrderID: orderID,
		},
		OrdersItemPricingElements: &apiInputReader.OrdersItemPricingElements{
			OrderID:   orderID,
			OrderItem: 1,
		},
		OrdersPartner: &apiInputReader.OrdersPartner{
			OrderID: orderID,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(orderID),
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
	controller *OrdersDetailListForAnOrderDocumentController,
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
	controller *OrdersDetailListForAnOrderDocumentController,
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
	controller *OrdersDetailListForAnOrderDocumentController,
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
	controller *OrdersDetailListForAnOrderDocumentController,
) createOrdersRequestPartners(
	requestPram *apiInputReader.Request,
	input apiInputReader.Orders,
) *apiModuleRuntimesResponsesOrders.OrdersRes {
	responseJsonData := apiModuleRuntimesResponsesOrders.OrdersRes{}
	responseBody := apiModuleRuntimesRequestsOrders.OrdersReads(
		requestPram,
		input,
		&controller.Controller,
		"Partners",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createOrdersRequestPartners Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersDetailListForAnOrderDocumentController,
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
	controller *OrdersDetailListForAnOrderDocumentController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	var input []apiModuleRuntimesRequestsPlant.General

	for _, v := range *ordersRes.Message.Item {
		if v.DeliverToParty != nil {
			input = append(input, apiModuleRuntimesRequestsPlant.General{
				BusinessPartner: *v.DeliverToParty,
				Plant:           *v.DeliverToPlant,
			})
			input = append(input, apiModuleRuntimesRequestsPlant.General{
				BusinessPartner: *v.DeliverFromParty,
				Plant:           *v.DeliverFromPlant,
			})
		}

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
	controller *OrdersDetailListForAnOrderDocumentController,
) createProjectRequestProjects(
	requestPram *apiInputReader.Request,
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesProject.ProjectRes {
	var input []apiModuleRuntimesRequestsProject.Project
	//input := make([]apiModuleRuntimesRequestsProject.Project, len(*ordersRes.Message.Item))

	for _, v := range *ordersRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsProject.Project{
			Project: *v.Project,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesProject.ProjectRes{}
	responseBody := apiModuleRuntimesRequestsProject.ProjectReadsProjectsByProjects(
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
		controller.CustomLogger.Error("ProjectReadsProjectsByProjects Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersDetailListForAnOrderDocumentController,
) CreateProjectRequestWBSElement(
	requestPram *apiInputReader.Request,
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesProject.ProjectRes {
	var input apiModuleRuntimesRequestsProject.WBSElement
	//input := make([]apiModuleRuntimesRequestsProject.Project, len(*ordersRes.Message.Item))

	for _, v := range *ordersRes.Message.Header {
		input = apiModuleRuntimesRequestsProject.WBSElement{
			Project:    *v.Project,
			WBSElement: *v.WBSElement,
		}
	}

	responseJsonData := apiModuleRuntimesResponsesProject.ProjectRes{}
	responseBody := apiModuleRuntimesRequestsProject.ProjectReadsWBSElement(
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
		controller.CustomLogger.Error("ProjectReadsWBSElement Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersDetailListForAnOrderDocumentController,
) CreateIncotermsRequestText(
	requestPram *apiInputReader.Request,
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesIncoterms.IncotermsRes {
	incoterms := &(*ordersRes.Message.Header)[0].Incoterms
	var inputIncoterms *string
	if incoterms != nil {
		inputIncoterms = *incoterms
	}

	input := apiModuleRuntimesRequestsIncoterms.Incoterms{
		Incoterms: *inputIncoterms,
	}

	responseJsonData := apiModuleRuntimesResponsesIncoterms.IncotermsRes{}
	responseBody := apiModuleRuntimesRequestsIncoterms.IncotermsReadsText(
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
		controller.CustomLogger.Error("IncotermsReadsText Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersDetailListForAnOrderDocumentController,
) CreatePaymentTermsRequestText(
	requestPram *apiInputReader.Request,
	ordersRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesPaymentTerms.PaymentTermsRes {
	paymentTerms := &(*ordersRes.Message.Header)[0].PaymentTerms
	var inputPaymentTerms *string
	if paymentTerms != nil {
		inputPaymentTerms = paymentTerms
	}

	input := apiModuleRuntimesRequestsPaymentTerms.PaymentTerms{
		PaymentTerms: *inputPaymentTerms,
	}

	responseJsonData := apiModuleRuntimesResponsesPaymentTerms.PaymentTermsRes{}
	responseBody := apiModuleRuntimesRequestsPaymentTerms.PaymentTermsReadsText(
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
		controller.CustomLogger.Error("PaymentTermsReadsText Unmarshal error")
	}

	return &responseJsonData
}

func functionOrderPdfGenerates(
	input apiOutputFormatter.Orders,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_FUNCTION_ORDER_PDF_SRV"
	aPIType := "generates"

	if accepter == "Order" {
		input.Accepter = []string{
			"Order",
		}
	}

	marshaledRequest, err := json.Marshal(input)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}

func (
	controller *OrdersDetailListForAnOrderDocumentController,
) request(
	input apiInputReader.Orders,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := apiModuleRuntimesResponsesOrders.OrdersRes{}
	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	headerRes = *controller.createOrdersRequestHeader(
		controller.UserInfo,
		input,
	)

	businessPartnerRes = *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		&headerRes,
	)

	itemRes := controller.createOrdersRequestItems(
		controller.UserInfo,
		input,
	)

	itemPricingElementRes := controller.createOrdersRequestItemPricingElements(
		controller.UserInfo,
		input,
	)

	partnerRes := controller.createOrdersRequestPartners(
		controller.UserInfo,
		input,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		itemRes,
	)

	projectRes := controller.createProjectRequestProjects(
		controller.UserInfo,
		&headerRes,
	)

	wBSElementRes := controller.CreateProjectRequestWBSElement(
		controller.UserInfo,
		&headerRes,
	)

	incotermsTextRes := controller.CreateIncotermsRequestText(
		controller.UserInfo,
		&headerRes,
	)

	paymentTermsTextRes := controller.CreatePaymentTermsRequestText(
		controller.UserInfo,
		&headerRes,
	)

	controller.fin(
		&headerRes,
		itemRes,
		itemPricingElementRes,
		partnerRes,
		&businessPartnerRes,
		plantRes,
		projectRes,
		wBSElementRes,
		incotermsTextRes,
		paymentTermsTextRes,
	)
}

func (
	controller *OrdersDetailListForAnOrderDocumentController,
) fin(
	headerRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	itemRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	itemPricingElementRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	partnerRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	projectRes *apiModuleRuntimesResponsesProject.ProjectRes,
	wBSElementRes *apiModuleRuntimesResponsesProject.ProjectRes,
	incotermsTextRes *apiModuleRuntimesResponsesIncoterms.IncotermsRes,
	paymentTermsTextRes *apiModuleRuntimesResponsesPaymentTerms.PaymentTermsRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	projectMapper := services.ProjectMapper(
		projectRes.Message.Project,
	)

	wBSElementMapper := services.WBSElementMapper(
		wBSElementRes.Message.WBSElement,
	)

	incotermsTextMapper := services.IncotermsTextMapper(
		incotermsTextRes.Message.Text,
	)

	paymentTermsTextMapper := services.PaymentTermsTextMapper(
		paymentTermsTextRes.Message.Text,
	)

	data := apiOutputFormatter.Orders{}

	for _, v := range *headerRes.Message.Header {
		projectDescription := projectMapper[*v.Project].ProjectDescription

		var incotermsName *string

		incotermsTextMapperName := incotermsTextMapper[*v.Incoterms].IncotermsName
		if &incotermsTextMapperName != nil {
			incotermsName = &incotermsTextMapperName
		}

		data.OrdersHeaderWithItem = append(data.OrdersHeaderWithItem,
			apiOutputFormatter.OrdersHeaderWithItem{
				OrderID:                         v.OrderID,
				OrderDate:                       v.OrderDate,
				OrderType:                       v.OrderType,
				Buyer:                           v.Buyer,
				BuyerName:                       businessPartnerMapper[v.Buyer].BusinessPartnerName,
				Seller:                          v.Seller,
				SellerName:                      businessPartnerMapper[v.Seller].BusinessPartnerName,
				RequestedDeliveryDate:           v.RequestedDeliveryDate,
				RequestedDeliveryTime:           v.RequestedDeliveryTime,
				TotalGrossAmount:                v.TotalGrossAmount,
				Contract:                        v.Contract,
				ContractItem:                    v.ContractItem,
				Project:                         v.Project,
				ProjectDescription:              &projectDescription,
				WBSElement:                      v.WBSElement,
				WBSElementResponsiblePersonName: wBSElementMapper[*v.Project].ResponsiblePersonName,
				Incoterms:                       v.Incoterms,
				IncotermsName:                   incotermsName,
				PricingDate:                     v.PricingDate,
				PaymentTerms:                    v.PaymentTerms,
				PaymentTermsName:                paymentTermsTextMapper[v.PaymentTerms].PaymentTermsName,
				PaymentMethod:                   v.PaymentMethod,
				TransactionCurrency:             v.TransactionCurrency,
				HeaderText:                      v.HeaderText,
			},
		)
	}

	for _, v := range *itemRes.Message.Item {
		data.OrdersItem = append(data.OrdersItem,
			apiOutputFormatter.OrdersItem{
				OrderID:   v.OrderID,
				OrderItem: v.OrderItem,
				//				DeliverToParty:              v.DeliverToParty,
				//				DeliverToPartyName:          businessPartnerMapper[v.DeliverToParty].BusinessPartnerName,
				DeliverToPlant:     *v.DeliverToPlant,
				DeliverToPlantName: plantMapper[strconv.Itoa(*v.DeliverToParty)].PlantName,
				//				DeliverFromParty:            v.DeliverFromParty,
				//				DeliverFromPartyName:        businessPartnerMapper[v.DeliverFromParty].BusinessPartnerName,
				DeliverFromPlant:            *v.DeliverFromPlant,
				DeliverFromPlantName:        plantMapper[strconv.Itoa(*v.DeliverFromParty)].PlantName,
				Product:                     v.Product,
				ProductSpecification:        v.ProductSpecification,
				SizeOrDimensionText:         v.SizeOrDimensionText,
				OrderItemText:               v.OrderItemText,
				OrderItemTextByBuyer:        v.OrderItemTextByBuyer,
				OrderItemTextBySeller:       v.OrderItemTextBySeller,
				OrderQuantityInBaseUnit:     v.OrderQuantityInBaseUnit,
				OrderQuantityInDeliveryUnit: v.OrderQuantityInDeliveryUnit,
				BaseUnit:                    v.BaseUnit,
				DeliveryUnit:                v.DeliveryUnit,
				RequestedDeliveryDate:       v.RequestedDeliveryDate,
				RequestedDeliveryTime:       v.RequestedDeliveryTime,
				NetAmount:                   v.NetAmount,
				TaxAmount:                   v.TaxAmount,
				GrossAmount:                 v.GrossAmount,
				ProductNetWeight:            v.ProductNetWeight,
			},
		)
	}

	for _, v := range *itemPricingElementRes.Message.ItemPricingElement {
		data.OrdersItemPricingElement = append(data.OrdersItemPricingElement,
			apiOutputFormatter.OrdersItemPricingElement{
				OrderID:                 v.OrderID,
				OrderItem:               v.OrderItem,
				PricingProcedureCounter: v.PricingProcedureCounter,
				ConditionType:           v.ConditionType,
				ConditionRateValue:      v.ConditionRateValue,
			},
		)
	}

	for _, v := range *partnerRes.Message.Partner {
		data.OrdersPartner = append(data.OrdersPartner,
			apiOutputFormatter.OrdersPartner{
				OrderID:                 v.OrderID,
				PartnerFunction:         v.PartnerFunction,
				BusinessPartner:         v.BusinessPartner,
				BusinessPartnerFullName: v.BusinessPartnerFullName,
				BusinessPartnerName:     v.BusinessPartnerName,
			},
		)
	}

	// ここから generates に rabbitmq で送信
	// accepter 対応
	responseJsonData := apiOutputFormatter.Orders{}
	responseBody := functionOrderPdfGenerates(
		data,
		&controller.Controller,
		"Order",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("apiModuleRuntimesRequestsOrderPdf.FunctionOrderPdfGenerates Unmarshal error")
	}

	data.MountPath = responseJsonData.MountPath

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
