package controllersProductionOrderList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesRequestsProductionOrder "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/production-order"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProductionOrder "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/production-order"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type ProductionOrderItemOperationListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ProductionOrderItemOperationListController) Get() {
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	productionOrder, _ := controller.GetInt("productionOrder")
	productionOrderItem, _ := controller.GetInt("productionOrderItem")
	redisKeyCategory1 := "production-order"
	redisKeyCategory2 := "item-operation-list"
	redisKeyCategory3 := productionOrder

	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	isMarkedForDeletion = false
	isReleased := false

	productionOrderParam := apiInputReader.ProductionOrder{
		ProductionOrderHeader: &apiInputReader.ProductionOrderHeader{
			ProductionOrder:     productionOrder,
			IsMarkedForDeletion: &isMarkedForDeletion,
			IsReleased:          &isReleased,
		},
		ProductionOrderItemOperation: &apiInputReader.ProductionOrderItemOperation{
			ProductionOrder:     productionOrder,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		ProductionOrderItemOperationComponent: &apiInputReader.ProductionOrderItemOperationComponent{
			ProductionOrder:     productionOrder,
			ProductionOrderItem: productionOrderItem,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
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
		var responseData apiOutputFormatter.ProductionOrder

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
			controller.request(productionOrderParam)
		}()
	} else {
		controller.request(productionOrderParam)
	}
}

func (
	controller *ProductionOrderItemOperationListController,
) createProductionOrderRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductionOrder,
) *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes {
	responseJsonData := apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes{}
	responseBody := apiModuleRuntimesRequestsProductionOrder.ProductionOrderReads(
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
		controller.CustomLogger.Error("createProductionOrderRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderItemOperationListController,
) createProductionOrderRequestItemOperations(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductionOrder,
) *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes {
	responseJsonData := apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes{}
	responseBody := apiModuleRuntimesRequestsProductionOrder.ProductionOrderReads(
		requestPram,
		input,
		&controller.Controller,
		"ItemOperations",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("ProductionOrderReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderItemOperationListController,
) createProductionOrderRequestItemOperationComponents(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductionOrder,
) *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes {
	responseJsonData := apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes{}
	responseBody := apiModuleRuntimesRequestsProductionOrder.ProductionOrderReads(
		requestPram,
		input,
		&controller.Controller,
		"ItemOperationComponents",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("ProductionOrderReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderItemOperationListController,
) createProductMasterDocRequest(
	requestPram *apiInputReader.Request,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes{}
	responseBody := apiModuleRuntimesRequestsProductMasterDoc.ProductMasterDocReads(
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
	controller *ProductionOrderItemOperationListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	productionOrderRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*productionOrderRes.Message.Header))
	for i, v := range *productionOrderRes.Message.Header {
		input[i].Plant = v.OwnerProductionPlant
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
	controller *ProductionOrderItemOperationListController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	bRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, len(*bRes.Message.Header))
	isMarkedForDeletion := false

	for _, v := range *bRes.Message.Header {
		productDescsByBP = append(productDescsByBP, apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BusinessPartner: v.OwnerProductionPlantBusinessPartner,
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
	controller *ProductionOrderItemOperationListController,
) createBusinessPartnerRequestGeneralsByBusinessPartners(
	requestPram *apiInputReader.Request,
	productionOrderRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*productionOrderRes.Message.ItemOperation))

	for _, v := range *productionOrderRes.Message.ItemOperation {
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
		controller.CustomLogger.Error("BusinessPartnerReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderItemOperationListController,
) request(
	input apiInputReader.ProductionOrder,
) {
	defer services.Recover(controller.CustomLogger)

	headerRes := controller.createProductionOrderRequestHeader(
		controller.UserInfo,
		input,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		headerRes,
	)

	productDescByBPRes := controller.createProductMasterRequestProductDescByBP(
		controller.UserInfo,
		headerRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	itemOperationsRes := controller.createProductionOrderRequestItemOperations(
		controller.UserInfo,
		input,
	)

	itemOperationComponentsRes := controller.createProductionOrderRequestItemOperationComponents(
		controller.UserInfo,
		input,
	)

	businessPartnerRes := controller.createBusinessPartnerRequestGeneralsByBusinessPartners(
		controller.UserInfo,
		itemOperationsRes,
	)

	controller.fin(
		headerRes,
		itemOperationsRes,
		itemOperationComponentsRes,
		businessPartnerRes,
		plantRes,
		productDescByBPRes,
		productDocRes,
	)
}

func (
	controller *ProductionOrderItemOperationListController,
) fin(
	headerRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	itemOperationsRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	itemOperationComponentsRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	productDescByBPRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	descriptionMapper := services.ProductDescByBPMapper(
		productDescByBPRes.Message.ProductDescByBP,
	)

	data := apiOutputFormatter.ProductionOrder{}

	for _, v := range *headerRes.Message.Header {
		img := services.CreateProductImage(
			productDocRes,
			v.OwnerProductionPlantBusinessPartner,
			v.Product,
		)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)
		plantName := fmt.Sprintf("%s", plantMapper[v.OwnerProductionPlant].PlantName)

		data.ProductionOrderHeaderWithItem = append(data.ProductionOrderHeaderWithItem,
			apiOutputFormatter.ProductionOrderHeaderWithItem{
				Product:                                 v.Product,
				ProductionOrder:                         v.ProductionOrder,
				ProductDescription:                      productDescription,
				OwnerProductionPlantBusinessPartnerName: businessPartnerMapper[v.OwnerProductionPlantBusinessPartner].BusinessPartnerName,
				OwnerProductionPlant:                    v.OwnerProductionPlant,
				OwnerProductionPlantName:                plantName,
				ProductionOrderQuantityInBaseUnit:       v.ProductionOrderQuantityInBaseUnit,
				ProductionOrderQuantityInDestinationProductionUnit: v.ProductionOrderQuantityInDestinationProductionUnit,
				ProductionOrderPlannedStartDate:                    v.ProductionOrderPlannedStartDate,
				ProductionOrderPlannedStartTime:                    v.ProductionOrderPlannedStartTime,
				ProductionOrderPlannedEndDate:                      v.ProductionOrderPlannedEndDate,
				ProductionOrderPlannedEndTime:                      v.ProductionOrderPlannedEndTime,
				Images: apiOutputFormatter.Images{
					Product: img,
				},
			},
		)
	}

	for _, v := range *itemOperationsRes.Message.ItemOperation {
		sellerName := fmt.Sprintf("%s", businessPartnerMapper[v.Seller].BusinessPartnerName)

		data.ProductionOrderItemOperation = append(data.ProductionOrderItemOperation,
			apiOutputFormatter.ProductionOrderItemOperation{
				ProductionOrder:      v.ProductionOrder,
				ProductionOrderItem:  v.ProductionOrderItem,
				Operations:           v.Operations,
				OperationsItem:       v.OperationsItem,
				OperationID:          v.OperationID,
				OperationText:        v.OperationText,
				Product:              v.Product,
				Seller:               v.Seller,
				SellerName:           sellerName,
				IsReleased:           v.IsReleased,
				IsMarkedForDeletion:  v.IsMarkedForDeletion,
				IsPartiallyConfirmed: v.IsPartiallyConfirmed,
				IsConfirmed:          v.IsConfirmed,
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
