package controllersProductionOrderItemSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesRequestsProductionOrder "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/production-order/production-order"
	apiModuleRuntimesRequestsProductionOrderDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/production-order/production-order-doc"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProductionOrder "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/production-order"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ProductionOrderItemSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const ()

func (controller *ProductionOrderItemSingleUnitController) Get() {
	//isReleased, _ := controller.GetBool("isReleased")
	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "productionOrder"
	redisKeyCategory2 := "item-single-unit"
	productionOrder, _ := controller.GetInt("productionOrder")
	productionOrderItem, _ := controller.GetInt("productionOrderItem")

	isReleased := false
	isMarkedForDeletion := false

	productionOrderItemSingleUnit := apiInputReader.ProductionOrder{
		ProductionOrderHeader: &apiInputReader.ProductionOrderHeader{
			ProductionOrder:     productionOrder,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		ProductionOrderItem: &apiInputReader.ProductionOrderItem{
			ProductionOrder:     productionOrder,
			ProductionOrderItem: productionOrderItem,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		ProductionOrderDocItemDoc: &apiInputReader.ProductionOrderDocItemDoc{
			ProductionOrder:          productionOrder,
			ProductionOrderItem:      productionOrderItem,
			DocType:                  "QRCODE",
			DocIssuerBusinessPartner: *controller.UserInfo.BusinessPartner,
		},
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
			controller.request(productionOrderItemSingleUnit)
		}()
	} else {
		controller.request(productionOrderItemSingleUnit)
	}
}

func (
	controller *ProductionOrderItemSingleUnitController,
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
	controller *ProductionOrderItemSingleUnitController,
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
	controller *ProductionOrderItemSingleUnitController,
) createProductMasterRequestGenerals(
	requestPram *apiInputReader.Request,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsGenerals(
		requestPram,
		apiInputReader.ProductMaster{},
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductMasterRequestGenerals Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderItemSingleUnitController,
) createProductionOrderDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductionOrder,
) *apiModuleRuntimesResponsesProductionOrder.ProductionOrderDocRes {
	responseJsonData := apiModuleRuntimesResponsesProductionOrder.ProductionOrderDocRes{}
	responseBody := apiModuleRuntimesRequestsProductionOrderDoc.ProductionOrderDocReads(
		requestPram,
		input,
		&controller.Controller,
		"ItemDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductionOrderDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderItemSingleUnitController,
) createProductMasterRequestBPPlant(
	requestPram *apiInputReader.Request,
	productionOrderRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsByBPPlant(
		requestPram,
		apiModuleRuntimesRequestsProductMaster.General{
			Product: *(*productionOrderRes.Message.Item)[0].Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BPPlant: []apiModuleRuntimesRequestsProductMaster.BPPlant{
						{
							BusinessPartner: *(*productionOrderRes.Message.Item)[0].ProductionPlantBusinessPartner,
							Plant:           *(*productionOrderRes.Message.Item)[0].ProductionPlant,
						},
					},
				},
			},
		},
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductMasterRequestBPPlant Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderItemSingleUnitController,
) createProductMasterRequestProduction(
	requestPram *apiInputReader.Request,
	productionOrderRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsProduction(
		requestPram,
		apiModuleRuntimesRequestsProductMaster.General{
			Product: *(*productionOrderRes.Message.Item)[0].Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BPPlant: []apiModuleRuntimesRequestsProductMaster.BPPlant{
						{
							BusinessPartner: *(*productionOrderRes.Message.Item)[0].ProductionPlantBusinessPartner,
							Plant:           *(*productionOrderRes.Message.Item)[0].ProductionPlant,
						},
					},
				},
			},
		},
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductMasterRequestProduction Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderItemSingleUnitController,
) request(
	input apiInputReader.ProductionOrder,
) {
	defer services.Recover(controller.CustomLogger)

	itemRes := controller.createProductionOrderRequestItem(
		controller.UserInfo,
		input,
	)

	generalsRes := controller.createProductMasterRequestGenerals(
		controller.UserInfo,
	)

	bpPlantRes := controller.createProductMasterRequestBPPlant(
		controller.UserInfo,
		itemRes,
	)

	productionRes := controller.createProductMasterRequestProduction(
		controller.UserInfo,
		itemRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	itemDocRes := controller.createProductionOrderDocRequest(
		controller.UserInfo,
		input,
	)

	controller.fin(
		itemRes,
		generalsRes,
		bpPlantRes,
		productionRes,
		productDocRes,
		itemDocRes,
	)
}

func (
	controller *ProductionOrderItemSingleUnitController,
) fin(
	itemRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	generalsRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	bpPlantRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productionRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	productionOrderDocRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderDocRes,
) {
	generalMapper := services.GeneralsMapper(
		generalsRes.Message.General,
	)

	data := apiOutputFormatter.ProductionOrder{}

	for _, v := range *itemRes.Message.Item {
		img := services.ReadProductImage(
			productDocRes,
			*v.ProductionPlantBusinessPartner,
			*v.Product,
		)

		qrcode := services.CreateQRCodeProductionOrderItemDocImage(
			productionOrderDocRes,
		)

		sizeOrDimensionText := fmt.Sprintf("%s", *generalMapper[*v.Product].SizeOrDimensionText)
		internalCapacityQuantity := *generalMapper[*v.Product].InternalCapacityQuantity

		var SafetyStockQuantityInBaseUnit float32
		var ReorderThresholdQuantityInBaseUnit float32
		var StandardProductionLotSizeQuantityInBaseUnit float32

		for _, b := range *bpPlantRes.Message.BPPlant {
			SafetyStockQuantityInBaseUnit = *b.SafetyStockQuantityInBaseUnit
			ReorderThresholdQuantityInBaseUnit = *b.ReorderThresholdQuantityInBaseUnit
		}

		for _, pd := range *productionRes.Message.Production {
			StandardProductionLotSizeQuantityInBaseUnit = pd.StandardProductionLotSizeQuantityInBaseUnit
		}

		data.ProductionOrderItemSingleUnit = append(data.ProductionOrderItemSingleUnit,
			apiOutputFormatter.ProductionOrderItemSingleUnit{
				SizeOrDimensionText:                         &sizeOrDimensionText,
				InternalCapacityQuantity:                    &internalCapacityQuantity,
				SafetyStockQuantityInBaseUnit:               &SafetyStockQuantityInBaseUnit,
				ReorderThresholdQuantityInBaseUnit:          &ReorderThresholdQuantityInBaseUnit,
				StandardProductionLotSizeQuantityInBaseUnit: &StandardProductionLotSizeQuantityInBaseUnit,
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
