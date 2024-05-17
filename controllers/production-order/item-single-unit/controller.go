package controllersProductionOrderItemSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesRequestsProductionOrder "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/production-order/production-order"
	apiModuleRuntimesRequestsProductionOrderDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/production-order/production-order-doc"
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

type ProductionOrderItemSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const ()

func (controller *ProductionOrderItemSingleUnitController) Get() {
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "productionOrder"
	redisKeyCategory2 := "item-single-unit"
	productionOrder, _ := controller.GetInt("productionOrder")
	productionOrderItem, _ := controller.GetInt("productionOrderItem")

	isReleased := false
	isCancelled := false
	isMarkedForDeletion := false

	productionOrderItemSingleUnit := apiInputReader.ProductionOrder{
		//		ProductionOrderHeader: &apiInputReader.ProductionOrderHeader{
		//			ProductionOrder:     productionOrder,
		//			IsReleased:          &isReleased,
		//			isCancelled:         &isCancelled,
		//			IsMarkedForDeletion: &isMarkedForDeletion,
		//		},
		ProductionOrderItem: &apiInputReader.ProductionOrderItem{
			ProductionOrder:     productionOrder,
			ProductionOrderItem: productionOrderItem,
			IsReleased:          &isReleased,
			IsCancelled:         &isCancelled,
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
			strconv.Itoa(productionOrder),
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
) createProductMasterRequestGeneral(
	requestPram *apiInputReader.Request,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsGeneral(
		requestPram,
		apiModuleRuntimesRequestsProductMaster.General{},
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductMasterRequestGeneral Unmarshal error")
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

	bpPlant := apiModuleRuntimesRequestsProductMaster.General{}

	for _, v := range *productionOrderRes.Message.Item {
		bpPlant = apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BPPlant: []apiModuleRuntimesRequestsProductMaster.BPPlant{
						{
							BusinessPartner: v.ProductionPlantBusinessPartner,
							Plant:           v.ProductionPlant,
						},
					},
				},
			},
		}
	}

	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsByBPPlant(
		requestPram,
		bpPlant,
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

	production := apiModuleRuntimesRequestsProductMaster.General{}

	for _, v := range *productionOrderRes.Message.Item {
		production = apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BPPlant: []apiModuleRuntimesRequestsProductMaster.BPPlant{
						{
							BusinessPartner: v.ProductionPlantBusinessPartner,
							Plant:           v.ProductionPlant,
						},
					},
				},
			},
		}
	}

	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsProduction(
		requestPram,
		production,
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
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	productionOrderItemRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*productionOrderItemRes.Message.Item))

	for _, v := range *productionOrderItemRes.Message.Item {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.ProductionPlantBusinessPartner,
		})
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
	controller *ProductionOrderItemSingleUnitController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	productionOrderItemRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	var input []apiModuleRuntimesRequestsPlant.General
	for _, v := range *productionOrderItemRes.Message.Item {
		input = append(input, apiModuleRuntimesRequestsPlant.General{
			Plant:           v.ProductionPlant,
			BusinessPartner: v.ProductionPlantBusinessPartner,
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
		controller.CustomLogger.Error("createPlantRequestGenerals Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderItemSingleUnitController,
) request(
	input apiInputReader.ProductionOrder,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	itemRes := controller.createProductionOrderRequestItem(
		controller.UserInfo,
		input,
	)

	productMasterGeneralRes := controller.createProductMasterRequestGeneral(
		controller.UserInfo,
	)

	productMasterBPPlantRes := controller.createProductMasterRequestBPPlant(
		controller.UserInfo,
		itemRes,
	)

	productMasterProductionRes := controller.createProductMasterRequestProduction(
		controller.UserInfo,
		itemRes,
	)

	productMasterDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		itemRes,
	)

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		itemRes,
	)

	itemDocRes := controller.createProductionOrderDocRequest(
		controller.UserInfo,
		input,
	)

	controller.fin(
		itemRes,
		productMasterGeneralRes,
		productMasterBPPlantRes,
		productMasterProductionRes,
		productMasterDocRes,
		plantRes,
		&businessPartnerRes,
		itemDocRes,
	)
}

func (
	controller *ProductionOrderItemSingleUnitController,
) fin(
	itemRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	productMasterGeneralRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productMasterBPPlantRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productMasterProductionRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productMasterDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	itemDocRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderDocRes,
) {
	//generalMapper := services.GeneralsMapper(
	//	productMasterGeneralRes.Message.General,
	//)

	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	data := apiOutputFormatter.ProductionOrder{}

	for _, v := range *itemRes.Message.Item {
		img := services.ReadProductImage(
			productMasterDocRes,
			v.ProductionPlantBusinessPartner,
			v.Product,
		)

		qrcode := services.CreateQRCodeProductionOrderItemDocImage(
			itemDocRes,
		)

		//sizeOrDimensionText := fmt.Sprintf("%s", *generalMapper[*v.Product].SizeOrDimensionText)
		//internalCapacityQuantity := *generalMapper[*v.Product].InternalCapacityQuantity

		var SafetyStockQuantityInBaseUnit float32
		var ReorderThresholdQuantityInBaseUnit float32
		var StandardProductionLotSizeQuantityInBaseUnit float32

		for _, b := range *productMasterBPPlantRes.Message.BPPlant {
			SafetyStockQuantityInBaseUnit = *b.SafetyStockQuantityInBaseUnit
			ReorderThresholdQuantityInBaseUnit = *b.ReorderThresholdQuantityInBaseUnit
		}

		for _, pd := range *productMasterProductionRes.Message.Production {
			StandardProductionLotSizeQuantityInBaseUnit = pd.StandardProductionLotSizeQuantityInBaseUnit
		}

		var buyerName string

		buyerNameMapperName := businessPartnerMapper[v.Buyer].BusinessPartnerName
		if &buyerNameMapperName != nil {
			buyerName = buyerNameMapperName
		}

		var sellerName string

		sellerNameMapperName := businessPartnerMapper[v.Seller].BusinessPartnerName
		if &sellerNameMapperName != nil {
			sellerName = sellerNameMapperName
		}

		var productionPlantBusinessPartnerName *string

		productionPlantBusinessPartnerNameMapperName := businessPartnerMapper[v.ProductionPlantBusinessPartner].BusinessPartnerName
		if &productionPlantBusinessPartnerName != nil {
			productionPlantBusinessPartnerName = &productionPlantBusinessPartnerNameMapperName
		}

		productionPlantName := fmt.Sprintf("%s", plantMapper[strconv.Itoa(v.ProductionPlantBusinessPartner)].PlantName)

		data.ProductionOrderItemSingleUnit = append(data.ProductionOrderItemSingleUnit,
			apiOutputFormatter.ProductionOrderItemSingleUnit{
				ProductionOrder:                         v.ProductionOrder,
				ProductionOrderItem:                     v.ProductionOrderItem,
				ProductionOrderItemDate:                 v.ProductionOrderItemDate,
				Product:                                 v.Product,
				Buyer:                                   v.Buyer,
				BuyerName:                               buyerName,
				Seller:                                  v.Seller,
				SellerName:                              sellerName,
				InspectionLot:                           v.InspectionLot,
				ProductionPlantBusinessPartner:          v.ProductionPlantBusinessPartner,
				ProductionPlantBusinessPartnerName:      productionPlantBusinessPartnerNameMapperName,
				ProductionPlant:                         v.ProductionPlant,
				ProductionPlantName:                     productionPlantName,
				ProductionOrderQuantityInBaseUnit:       v.ProductionOrderQuantityInBaseUnit,
				ProductionOrderQuantityInProductionUnit: v.ProductionOrderQuantityInProductionUnit,
				//SizeOrDimensionText:                         &sizeOrDimensionText,
				//InternalCapacityQuantity:                    &internalCapacityQuantity,
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
