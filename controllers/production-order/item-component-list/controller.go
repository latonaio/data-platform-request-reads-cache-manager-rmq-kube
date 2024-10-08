package controllersProductionOrderItemComponentList

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

type ProductionOrderItemComponentListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ProductionOrderItemComponentListController) Get() {
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	productionOrder, _ := controller.GetInt("productionOrder")
	productionOrderItem, _ := controller.GetInt("productionOrderItem")
	redisKeyCategory1 := "production-order"
	redisKeyCategory2 := "item-component-list"
	redisKeyCategory3 := productionOrder
	redisKeyCategory4 := productionOrderItem

	isReleased := false
	isCancelled := false
	isMarkedForDeletion := false

	productionOrderParam := apiInputReader.ProductionOrder{
		ProductionOrderItem: &apiInputReader.ProductionOrderItem{
			ProductionOrder:     productionOrder,
			ProductionOrderItem: productionOrderItem,
			IsReleased:          &isReleased,
			IsCancelled:         &isCancelled,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		ProductionOrderItemComponent: &apiInputReader.ProductionOrderItemComponent{
			ProductionOrder:     productionOrder,
			ProductionOrderItem: productionOrderItem,
			IsReleased:          &isReleased,
			IsCancelled:         &isCancelled,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(redisKeyCategory3),
			strconv.Itoa(redisKeyCategory4),
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
	controller *ProductionOrderItemComponentListController,
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
	controller *ProductionOrderItemComponentListController,
) createProductionOrderRequestItemComponents(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductionOrder,
) *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes {
	responseJsonData := apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes{}
	responseBody := apiModuleRuntimesRequestsProductionOrder.ProductionOrderReads(
		requestPram,
		input,
		&controller.Controller,
		"ItemComponents",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductionOrderRequestItemComponents Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderItemComponentListController,
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
	controller *ProductionOrderItemComponentListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	productionOrderRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	var input []apiModuleRuntimesRequestsPlant.General
	for _, v := range *productionOrderRes.Message.Item {
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
	controller *ProductionOrderItemComponentListController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	productDescByBPRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, len(*productDescByBPRes.Message.Item))
	isMarkedForDeletion := false

	for _, v := range *productDescByBPRes.Message.Item {
		productDescsByBP = append(productDescsByBP, apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BusinessPartner: v.ProductionPlantBusinessPartner,
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
		controller.CustomLogger.Error("createProductMasterRequestProductDescByBP Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderItemComponentListController,
) createBusinessPartnerRequestGeneralsByBusinessPartners(
	requestPram *apiInputReader.Request,
	productionOrderItemRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	productionOrderItemComponentRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var generals []apiModuleRuntimesRequestsBusinessPartner.General

	for _, v := range *productionOrderItemRes.Message.Item {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Buyer,
		})
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Seller,
		})
	}

	for _, v := range *productionOrderItemComponentRes.Message.ItemComponent {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.ComponentProductBuyer, // ComponentProductSellerも必要
		})
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.ComponentProductSeller,
		})
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.ProductionPlantBusinessPartner,
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
	controller *ProductionOrderItemComponentListController,
) request(
	input apiInputReader.ProductionOrder,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	itemRes := controller.createProductionOrderRequestItem(
		controller.UserInfo,
		input,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		itemRes,
	)

	productDescByBPRes := controller.createProductMasterRequestProductDescByBP(
		controller.UserInfo,
		itemRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	itemComponentRes := controller.createProductionOrderRequestItemComponents(
		controller.UserInfo,
		input,
	)

	businessPartnerRes := controller.createBusinessPartnerRequestGeneralsByBusinessPartners(
		controller.UserInfo,
		itemRes,
		itemComponentRes,
	)

	controller.fin(
		itemRes,
		itemComponentRes,
		businessPartnerRes,
		plantRes,
		productDescByBPRes,
		productDocRes,
	)
}

func (
	controller *ProductionOrderItemComponentListController,
) fin(
	itemRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	itemComponentRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
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

	//descriptionMapper := services.ProductDescByBPMapper(
	//	productDescByBPRes.Message.ProductDescByBP,
	//)

	data := apiOutputFormatter.ProductionOrder{}

	for _, v := range *itemRes.Message.Item {
		img := services.ReadProductImage(
			productDocRes,
			v.ProductionPlantBusinessPartner,
			v.Product,
		)

		//productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)
		productionPlantName := fmt.Sprintf("%s", plantMapper[strconv.Itoa(v.ProductionPlantBusinessPartner)].PlantName)

		buyerName := fmt.Sprintf("%s", businessPartnerMapper[v.Buyer].BusinessPartnerName)
		sellerName := fmt.Sprintf("%s", businessPartnerMapper[v.Seller].BusinessPartnerName)
		productionPlantBusinessPartnerNameMapperName := fmt.Sprintf("%s", businessPartnerMapper[v.ProductionPlantBusinessPartner].BusinessPartnerName)

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
				Images: apiOutputFormatter.Images{
					Product: img,
				},
			},
		)
	}

	for _, v := range *itemComponentRes.Message.ItemComponent {
		img := services.ReadProductImage(
			productDocRes,
			v.ProductionPlantBusinessPartner,
			v.ComponentProduct,
		)

		componentProductBuyerName := fmt.Sprintf("%s", businessPartnerMapper[v.ComponentProductBuyer].BusinessPartnerName)
		componentProductSellerName := fmt.Sprintf("%s", businessPartnerMapper[v.ComponentProductSeller].BusinessPartnerName)

		data.ProductionOrderItemComponent = append(data.ProductionOrderItemComponent,
			apiOutputFormatter.ProductionOrderItemComponent{
				ProductionOrder:                                v.ProductionOrder,
				ProductionOrderItem:                            v.ProductionOrderItem,
				BillOfMaterial:                                 v.BillOfMaterial,
				BillOfMaterialItem:                             v.BillOfMaterialItem,
				ComponentProduct:                               v.ComponentProduct,
				ComponentProductBuyer:                          v.ComponentProductBuyer,
				ComponentProductBuyerName:                      componentProductBuyerName,
				ComponentProductSeller:                         v.ComponentProductSeller,
				ComponentProductSellerName:                     componentProductSellerName,
				ComponentProductBaseUnit:                       v.ComponentProductBaseUnit,
				ComponentProductDeliveryUnit:                   v.ComponentProductDeliveryUnit,
				ComponentProductRequiredQuantityInBaseUnit:     v.ComponentProductRequiredQuantityInBaseUnit,
				ComponentProductRequiredQuantityInDeliveryUnit: v.ComponentProductRequiredQuantityInDeliveryUnit,
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
