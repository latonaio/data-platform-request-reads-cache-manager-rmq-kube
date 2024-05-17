package controllersProductionOrderConfHeaderSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesRequestsProductionOrder "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/production-order/production-order"
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

type ProductionOrderConfHeaderSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ProductionOrderConfHeaderSingleUnitController) Get() {
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	productionOrder, _ := controller.GetInt("productionOrder")
	productionOrderItem, _ := controller.GetInt("productionOrderItem")
	operations, _ := controller.GetInt("operations")
	operationsItem, _ := controller.GetInt("operationsItem")
	operationID, _ := controller.GetInt("operationID")
	redisKeyCategory1 := "production-order-conf"
	redisKeyCategory2 := "header-single-unit"
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
		ProductionOrderItem: &apiInputReader.ProductionOrderItem{
			ProductionOrder:     productionOrder,
			ProductionOrderItem: productionOrderItem,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		ProductionOrderItemOperation: &apiInputReader.ProductionOrderItemOperation{
			ProductionOrder:     productionOrder,
			ProductionOrderItem: productionOrderItem,
			Operations:          operations,
			OperationsItem:      operationsItem,
			OperationID:         operationID,
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
	controller *ProductionOrderConfHeaderSingleUnitController,
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
	controller *ProductionOrderConfHeaderSingleUnitController,
) createProductionOrderRequestItemOperation(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductionOrder,
) *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes {
	responseJsonData := apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes{}
	responseBody := apiModuleRuntimesRequestsProductionOrder.ProductionOrderReads(
		requestPram,
		input,
		&controller.Controller,
		"ItemOperation",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductionOrderRequestItemOperation Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderConfHeaderSingleUnitController,
) createProductionOrderRequestItem(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductionOrder,
) *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes {
	responseJsonData := apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes{}
	responseBody := apiModuleRuntimesRequestsProductionOrder.ProductionOrderReads(
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
		controller.CustomLogger.Error("createProductionOrderRequestItem Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderConfHeaderSingleUnitController,
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
		controller.CustomLogger.Error("createProductionOrderRequestItemOperationComponents Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderConfHeaderSingleUnitController,
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
	controller *ProductionOrderConfHeaderSingleUnitController,
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
	controller *ProductionOrderConfHeaderSingleUnitController,
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
	controller *ProductionOrderConfHeaderSingleUnitController,
) createBusinessPartnerRequestGeneralsByBusinessPartners(
	requestPram *apiInputReader.Request,
	productionOrderHeaderRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	productionOrderItemOperationRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make(
		[]apiModuleRuntimesRequestsBusinessPartner.General,
		len(*productionOrderHeaderRes.Message.Header)+len(*productionOrderItemOperationRes.Message.ItemOperation),
	)

	for _, v := range *productionOrderHeaderRes.Message.Header {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Seller,
		})
	}

	for _, v := range *productionOrderItemOperationRes.Message.ItemOperation {
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
	controller *ProductionOrderConfHeaderSingleUnitController,
) request(
	input apiInputReader.ProductionOrder,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

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

	itemRes := controller.createProductionOrderRequestItem(
		controller.UserInfo,
		input,
	)

	itemOperationRes := controller.createProductionOrderRequestItemOperation(
		controller.UserInfo,
		input,
	)

	//itemOperationComponentsRes := controller.createProductionOrderRequestItemOperationComponents(
	//	controller.UserInfo,
	//	input,
	//)

	businessPartnerRes := controller.createBusinessPartnerRequestGeneralsByBusinessPartners(
		controller.UserInfo,
		headerRes,
		itemOperationRes,
	)

	controller.fin(
		headerRes,
		itemRes,
		itemOperationRes,
		//itemOperationComponentsRes,
		businessPartnerRes,
		plantRes,
		productDescByBPRes,
		productDocRes,
	)
}

func (
	controller *ProductionOrderConfHeaderSingleUnitController,
) fin(
	headerRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	itemRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	itemOperationRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	//itemOperationComponentsRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
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
		img := services.ReadProductImage(
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

	itemOperation := (*itemOperationRes.Message.ItemOperation)[0]

	sellerName := fmt.Sprintf("%s", businessPartnerMapper[itemOperation.Seller].BusinessPartnerName)

	data.ProductionOrderItemOperation = append(data.ProductionOrderItemOperation,
		apiOutputFormatter.ProductionOrderItemOperation{
			ProductionOrder:     itemOperation.ProductionOrder,
			ProductionOrderItem: itemOperation.ProductionOrderItem,
			Operations:          itemOperation.Operations,
			OperationsItem:      itemOperation.OperationsItem,
			OperationID:         itemOperation.OperationID,
			OperationText:       itemOperation.OperationText,
			Product:             itemOperation.Product,
			Seller:              itemOperation.Seller,
			SellerName:          sellerName,
			WorkCenter:          itemOperation.WorkCenter,
		},
	)

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
