package controllersProductionOrderDetailList

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

type ProductionOrderDetailListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ProductionOrderDetailListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	productionOrder, _ := controller.GetInt("productionOrder")
	redisKeyCategory1 := "production-order"
	redisKeyCategory2 := "detail-list"
	redisKeyCategory3 := productionOrder
	//	userType := controller.GetString(":userType")

	isReleased, _ := controller.GetBool("isReleased")
	isCancelled, _ := controller.GetBool("isCancelled")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	isReleased = false
	isCancelled = false
	isMarkedForDeletion = false

	productionOrderItems := apiInputReader.ProductionOrder{
		ProductionOrderHeader: &apiInputReader.ProductionOrderHeader{
			ProductionOrder:     productionOrder,
			IsCancelled:         &isCancelled,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		ProductionOrderItem: &apiInputReader.ProductionOrderItem{
			ProductionOrder:     productionOrder,
			IsCancelled:         &isCancelled,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(redisKeyCategory3),
			//			userType,
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
			controller.request(productionOrderItems)
		}()
	} else {
		controller.request(productionOrderItems)
	}
}

func (
	controller *ProductionOrderDetailListController,
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
	controller *ProductionOrderDetailListController,
) createProductionOrderRequestItems(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductionOrder,
) *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes {
	responseJsonData := apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes{}
	responseBody := apiModuleRuntimesRequestsProductionOrder.ProductionOrderReads(
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
		controller.CustomLogger.Error("createProductionOrderRequestItems Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderDetailListController,
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
	controller *ProductionOrderDetailListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	plantHeaderRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	plantItemRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	var input []apiModuleRuntimesRequestsPlant.General
	for _, v := range *plantHeaderRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsPlant.General{
			Plant:           v.OwnerProductionPlant,
			BusinessPartner: v.OwnerProductionPlantBusinessPartner,
		})
	}

	for _, v := range *plantItemRes.Message.Item {
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
	controller *ProductionOrderDetailListController,
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
	controller *ProductionOrderDetailListController,
) createBusinessPartnerRequestGeneralsByBusinessPartners(
	requestPram *apiInputReader.Request,
	productionOrderHeaderRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	productionOrderItemRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var generals []apiModuleRuntimesRequestsBusinessPartner.General

	for _, v := range *productionOrderHeaderRes.Message.Header {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Buyer,
		})
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Seller,
		})
	}

	for _, v := range *productionOrderItemRes.Message.Item {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.ProductionPlantBusinessPartner,
		})
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
		controller.CustomLogger.Error("BusinessPartnerReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderDetailListController,
) request(
	input apiInputReader.ProductionOrder,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := controller.createProductionOrderRequestHeader(
		controller.UserInfo,
		input,
	)

	itemRes := controller.createProductionOrderRequestItems(
		controller.UserInfo,
		input,
	)

	businessPartnerRes := controller.createBusinessPartnerRequestGeneralsByBusinessPartners(
		controller.UserInfo,
		headerRes,
		itemRes,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		headerRes,
		itemRes,
	)

	productDescByBPRes := controller.createProductMasterRequestProductDescByBP(
		controller.UserInfo,
		headerRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	controller.fin(
		headerRes,
		itemRes,
		businessPartnerRes,
		plantRes,
		productDescByBPRes,
		productDocRes,
	)
}

func (
	controller *ProductionOrderDetailListController,
) fin(
	headerRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	itemRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
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
		plantName := fmt.Sprintf("%s", plantMapper[strconv.Itoa(v.OwnerProductionPlantBusinessPartner)].PlantName)

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

		data.ProductionOrderHeaderWithItem = append(data.ProductionOrderHeaderWithItem,
			apiOutputFormatter.ProductionOrderHeaderWithItem{
				ProductionOrder:                   v.ProductionOrder,
				ProductionOrderDate:               v.ProductionOrderDate,
				Product:                           v.Product,
				ProductDescription:                productDescription,
				Buyer:                             v.Buyer,
				BuyerName:                         buyerName,
				Seller:                            v.Seller,
				SellerName:                        sellerName,
				OwnerProductionPlant:              v.OwnerProductionPlant,
				OwnerProductionPlantName:          plantName,
				ProductionOrderQuantityInBaseUnit: v.ProductionOrderQuantityInBaseUnit,
				ProductionOrderQuantityInDestinationProductionUnit: v.ProductionOrderQuantityInDestinationProductionUnit,
				ProductionOrderPlannedStartDate:                    v.ProductionOrderPlannedStartDate,
				ProductionOrderPlannedStartTime:                    v.ProductionOrderPlannedStartTime,
				ProductionOrderPlannedEndDate:                      v.ProductionOrderPlannedEndDate,
				ProductionOrderPlannedEndTime:                      v.ProductionOrderPlannedEndTime,
				InspectionLot:                                      v.InspectionLot,
				Images: apiOutputFormatter.Images{
					Product: img,
				},
			},
		)
	}

	for _, v := range *itemRes.Message.Item {
		img := services.ReadProductImage(
			productDocRes,
			v.ProductionPlantBusinessPartner,
			v.Product,
		)

		plantName := fmt.Sprintf("%s", plantMapper[v.ProductionPlant].PlantName)
		productionPlantBusinessPartnerName := fmt.Sprintf("%s", businessPartnerMapper[v.ProductionPlantBusinessPartner].BusinessPartnerName)
		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)

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

		data.ProductionOrderItem = append(data.ProductionOrderItem,
			apiOutputFormatter.ProductionOrderItem{
				ProductionOrder:                         v.ProductionOrder,
				ProductionOrderItem:                     v.ProductionOrderItem,
				MRPArea:                                 v.MRPArea,
				Product:                                 v.Product,
				ProductDescription:                      productDescription,
				Buyer:                                   v.Buyer,
				BuyerName:                               buyerName,
				Seller:                                  v.Seller,
				SellerName:                              sellerName,
				ProductBaseUnit:                         v.ProductBaseUnit,
				ProductProductionUnit:                   v.ProductProductionUnit,
				ProductionPlantBusinessPartner:          v.ProductionPlantBusinessPartner,
				ProductionPlantBusinessPartnerName:      productionPlantBusinessPartnerName,
				ProductionPlantName:                     plantName,
				ProductionOrderQuantityInBaseUnit:       v.ProductionOrderQuantityInBaseUnit,
				ProductionOrderQuantityInProductionUnit: v.ProductionOrderQuantityInProductionUnit,
				ProductionOrderPlannedStartDate:         v.ProductionOrderPlannedStartDate,
				ProductionOrderPlannedStartTime:         v.ProductionOrderPlannedStartTime,
				ProductionOrderPlannedEndDate:           v.ProductionOrderPlannedEndDate,
				ProductionOrderPlannedEndTime:           v.ProductionOrderPlannedEndTime,
				ConfirmedYieldQuantityInBaseUnit:        v.ConfirmedYieldQuantityInBaseUnit,
				InspectionLot:                           v.InspectionLot,
				IsReleased:                              v.IsReleased,
				IsMarkedForDeletion:                     v.IsMarkedForDeletion,
				IsPartiallyConfirmed:                    v.IsPartiallyConfirmed,
				IsConfirmed:                             v.IsConfirmed,
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
