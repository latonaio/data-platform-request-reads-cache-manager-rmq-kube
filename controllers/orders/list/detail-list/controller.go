package controllersOrdersDetailList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequests "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/orders"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResonsesOrders "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/orders"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master-doc"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"io/ioutil"
	"strconv"
	"strings"
)

type OrdersDetailListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *OrdersDetailListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	orders, _ := controller.GetInt("orders")
	redisKeyCategory1 := "orders"
	redisKeyCategory2 := "detail-list"
	redisKeyCategory3 := orders
	userType := controller.GetString("userType")

	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	ordersItems := apiInputReader.Orders{
		OrdersHeader: &apiInputReader.OrdersHeader{
			Orders: orders,
		},
		OrdersItems: &apiInputReader.OrdersItems{
			Orders:      orders,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(redisKeyCategory3),
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
			controller.request(ordersItems)
		}()
	} else {
		controller.request(ordersItems)
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
		controller.CustomLogger.Error("createOrdersRequestHeader Unmarshal error")
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
) createProductMasterDocRequest(
	requestPram *apiInputReader.Request,
) *apiModuleRuntimesResponsesProductMasterDoc.ProductMasterDocRes {
	responseJsonData := apiModuleRuntimesResponsesProductMasterDoc.ProductMasterDocRes{}
	responseBody := apiModuleRuntimesRequests.ProductMasterDocReads(
		requestPram,
		&controller.Controller,
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
) request(
	input apiInputReader.Orders,
) {
	defer services.Recover(controller.CustomLogger)

	bHeaderRes := controller.createOrdersRequestHeader(
		controller.UserInfo,
		input,
	)

	bRes := controller.createOrdersRequestItems(
		controller.UserInfo,
		input,
	)

	plRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		bRes,
	)

	pRes := controller.createProductMasterRequestProductDescByBP(
		controller.UserInfo,
		bHeaderRes,
	)

	pdRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	controller.fin(
		bHeaderRes,
		bRes,
		plRes,
		pRes,
		pdRes,
	)
}

func (
	controller *OrdersDetailListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	bRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	generals := make(apiModuleRuntimesRequests.PlantGenerals, len(*bRes.Message.Item))
	for i, v := range *bRes.Message.Item {
		generals[i].Plant = &v.DeliverToPlant            //複数対応が必要(例:StockConfirmationPlant)
		generals[i].Language = requestPram.Language
	}

	aPIServiceName := "DPFM_API_PLANT_SRV"
	aPIType := "reads"
	responseJsonData := apiModuleRuntimesResponsesPlant.PlantRes{}

	request := apiModuleRuntimesRequests.
		CreatePlantRequestGenerals(
			requestPram,
			generals,
		)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createPlantRequestGenerals error")
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		&controller.Controller,
	)

	err = json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createPlantRequestGenerals error")
	}

	return &responseJsonData
}

func (
	controller *OrdersListController,
) createProductMasterRequestProductDescByBPByBuyer(
	requestPram *apiInputReader.Request,
	pdByBuyerRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, 0)
	isMarkedForDeletion := false

	for _, v := range *pdByBuyerRes.Message.Header {
		productDescsByBP = append(productDescsByBP, apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BusinessPartner: v.Buyer,
					ProductDescByBP: []apiModuleRuntimesRequestsProductMaster.ProductDescByBP{
						{
							Language:            *requestPram.Language,
							IsMarkedForDeletion: &isMarkedForDeletion,
						},
					},
				},
			},
		})
	}

	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsProductDescsByBP(
		requestPram,
		productDescsByBP,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("ProductMasterReadsProductDescsByBP Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersListController,
) createProductMasterRequestProductDescByBPBySeller(
	requestPram *apiInputReader.Request,
	pdByBuyerRes *apiModuleRuntimesResponsesOrders.OrdersRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, 0)
	isMarkedForDeletion := false

	for _, v := range *pdByBuyerRes.Message.Header {
		productDescsByBP = append(productDescsByBP, apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BusinessPartner: v.Seller,
					ProductDescByBP: []apiModuleRuntimesRequestsProductMaster.ProductDescByBP{
						{
							Language:            *requestPram.Language,
							IsMarkedForDeletion: &isMarkedForDeletion,
						},
					},
				},
			},
		})
	}

	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsProductDescsByBP(
		requestPram,
		productDescsByBP,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("ProductMasterReadsProductDescsByBP Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersListController,
) createProductMasterDocRequest(
	requestPram *apiInputReader.Request,
) *apiModuleRuntimesResponses.ProductMasterDocRes {
	responseJsonData := apiModuleRuntimesResponses.ProductMasterDocRes{}
	responseBody := apiModuleRuntimesRequests.ProductMasterDocReads(
		requestPram,
		&controller.Controller,
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
) fin(
	bHeaderRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	bRes *apiModuleRuntimesResponsesOrders.OrdersRes,
	plRes *apiModuleRuntimesResponsesPlant.PlantRes,
	pRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	pdRes *apiModuleRuntimesResponsesProductMasterDoc.ProductMasterDocRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)
	
	descriptionMapper := services.ProductDescByBPMapper(
		pRes.Message.ProductDescByBP,
	)

	plantMapper := services.PlantMapper(
		plRes.Message.Generals,
	)

	data := apiOutputFormatter.Orders{}

	for _, v := range *bHeaderRes.Message.Header {
		img := services.CreateProductImage(
			pdRes,
			v.Buyer,	//Sellerの対応が必要
			v.Product,
		)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)

		data.OrdersHeader = append(data.OrdersHeader,
			apiOutputFormatter.OrdersHeader{
				OrderID:                  v.OrderID,
				Orders:                   v.Orders,
				ProductDescription:       &productDescription,
				OwnerProductionPlant:     v.OwnerProductionPlant,
				OwnerProductionPlantName: plantMapper[v.OwnerProductionPlant].PlantName,
				ValidityStartDate:        v.ValidityStartDate,
				IsMarkedForDeletion:      v.IsMarkedForDeletion,
			},
		)
	}

	for _, v := range *bRes.Message.Item {
		data.OrdersItem = append(data.OrdersItem,
			apiOutputFormatter.OrdersItem{
				OrderItem:									v.ComponentProduct,
				Product:                         			v.Product,
				OrderItemTextByBuyer:                     	v.OrderItemTextByBuyer,
				OrderItemTextBySeller:                     	v.OrderItemTextBySeller,
				OrderQuantityInDeliveryUnit:                v.OrderQuantityInDeliveryUnit,
				DeliveryUnit:								v.DeliveryUnit,
				RequestedDeliveryDate:						v.RequestedDeliveryDate,
				NetAmount:									v.NetAmount,
				IsCancelled:								v.IsCancelled,
				IsMarkedForDeletion:                        v.IsMarkedForDeletion,
				Images: apiOutputFormatter.Images{
					Product: img,
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
