package controllersProductionOrderList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
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
)

type ProductionOrderListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const (
	OwnerProductionPlantBusinessPartner = "ownerProductionPlantBusinessPartner"
)

func (controller *ProductionOrderListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "production-order"
	redisKeyCategory2 := "list"
	//userType := OwnerProductionPlantBusinessPartner
	productionOrder, _ := controller.GetInt("productionOrder")
	userType := controller.GetString(":userType")

	isReleased := false
	isMarkedForDeletion = false

	productionOrderHeader := apiInputReader.ProductionOrder{
		ProductionOrderHeader: &apiInputReader.ProductionOrderHeader{
			ProductionOrder:     productionOrder,
			IsMarkedForDeletion: &isMarkedForDeletion,
			IsReleased:          &isReleased,
		},
	}

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
			controller.request(productionOrderHeader)
		}()
	} else {
		controller.request(productionOrderHeader)
	}
}

func (
	controller *ProductionOrderListController,
) createProductionOrderRequestHeaderByOwnerProductionPlantBP(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductionOrder,
) *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes {
	responseJsonData := apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes{}
	responseBody := apiModuleRuntimesRequestsProductionOrder.ProductionOrderReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersByOwnerProductionPlantBP",
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
	controller *ProductionOrderListController,
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
	controller *ProductionOrderListController,
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
	controller *ProductionOrderListController,
) createBusinessPartnerRequestGeneralsByBusinessPartners(
	requestPram *apiInputReader.Request,
	productionOrderRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*productionOrderRes.Message.Header))

	for _, v := range *productionOrderRes.Message.Header {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.OwnerProductionPlantBusinessPartner,
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
	controller *ProductionOrderListController,
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
	controller *ProductionOrderListController,
) request(
	input apiInputReader.ProductionOrder,
) {
	defer services.Recover(controller.CustomLogger)

	headerRes := controller.createProductionOrderRequestHeaderByOwnerProductionPlantBP(
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

	bpRes := controller.createBusinessPartnerRequestGeneralsByBusinessPartners(
		controller.UserInfo,
		headerRes,
	)

	controller.fin(
		headerRes,
		plantRes,
		productDescByBPRes,
		productDocRes,
		bpRes,
	)
}

func (
	controller *ProductionOrderListController,
) fin(
	headerRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	productDescByBPRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {

	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	descriptionMapper := services.ProductDescByBPMapper(
		productDescByBPRes.Message.ProductDescByBP,
	)

	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.ProductionOrder{}

	for _, v := range *headerRes.Message.Header {
		img := services.ReadProductImage(
			productDocRes,
			v.OwnerProductionPlantBusinessPartner,
			v.Product,
		)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)

		data.ProductionOrderHeader = append(data.ProductionOrderHeader,
			apiOutputFormatter.ProductionOrderHeader{
				ProductionOrder:                         v.ProductionOrder,
				MRPArea:                                 v.MRPArea,
				Product:                                 v.Product,
				ProductDescription:                      productDescription,
				OwnerProductionPlantBusinessPartnerName: businessPartnerMapper[v.OwnerProductionPlantBusinessPartner].BusinessPartnerName,
				OwnerProductionPlantName:                plantMapper[v.OwnerProductionPlant].PlantName,
				ProductionOrderQuantityInBaseUnit:       v.ProductionOrderQuantityInBaseUnit,
				IsReleased:                              v.IsReleased,
				IsPartiallyConfirmed:                    v.IsPartiallyConfirmed,
				IsConfirmed:                             v.IsConfirmed,
				IsCancelled:                             v.IsCancelled,
				IsMarkedForDeletion:                     v.IsMarkedForDeletion,
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
