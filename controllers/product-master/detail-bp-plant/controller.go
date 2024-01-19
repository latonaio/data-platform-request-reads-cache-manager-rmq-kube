package controllersProductMasterDetailBPPlant

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ProductMasterDetailBPPlantController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ProductMasterDetailBPPlantController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	product := controller.GetString("product")
	//businessPartner, _ := controller.GetInt("businessPartner")
	//plant := controller.GetString("plant")
	redisKeyCategory1 := "product-master"
	redisKeyCategory2 := "detail-bp-plant"
	redisKeyCategory3 := product
	userType := controller.GetString(":userType")

	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	productMasterBPPlant := apiInputReader.ProductMaster{
		ProductMasterGeneral: &apiInputReader.ProductMasterGeneral{
			Product: product,
		},
		ProductMasterBPPlant: &apiInputReader.ProductMasterBPPlant{
			Product:             product,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			redisKeyCategory3,
			userType,
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.ProductMaster

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
			controller.request(productMasterBPPlant)
		}()
	} else {
		controller.request(productMasterBPPlant)
	}
}

func (
	controller *ProductMasterDetailBPPlantController,
) createProductMasterRequestGeneral(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductMaster,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsGeneral(
		requestPram,
		apiModuleRuntimesRequestsProductMaster.General{
			Product: input.ProductMasterBPPlant.Product,
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
		controller.CustomLogger.Error("createProductMasterRequestGeneral Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductMasterDetailBPPlantController,
) createProductMasterRequestBPPlants(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductMaster,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsByBPPlants(
		requestPram,
		apiModuleRuntimesRequestsProductMaster.General{
			Product: input.ProductMasterBPPlant.Product,
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
		controller.CustomLogger.Error("ProductMasterReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductMasterDetailBPPlantController,
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
	controller *ProductMasterDetailBPPlantController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	bpPlantRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, len(*bpPlantRes.Message.BPPlant))
	isMarkedForDeletion := false

	for _, v := range *bpPlantRes.Message.BPPlant {
		productDescsByBP = append(productDescsByBP, apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BusinessPartner: v.BusinessPartner,
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
	controller *ProductMasterDetailBPPlantController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	bpPlantsRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*bpPlantsRes.Message.BPPlant))

	for _, v := range *bpPlantsRes.Message.BPPlant {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.BusinessPartner,
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
	controller *ProductMasterDetailBPPlantController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	bpPlantRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*bpPlantRes.Message.BPPlant))
	for i, v := range *bpPlantRes.Message.BPPlant {
		input[i].Plant = v.Plant
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
	controller *ProductMasterDetailBPPlantController,
) request(
	input apiInputReader.ProductMaster,
) {
	defer services.Recover(controller.CustomLogger)

	generalRes := controller.createProductMasterRequestGeneral(
		controller.UserInfo,
		input,
	)

	bPPlantRes := controller.createProductMasterRequestBPPlants(
		controller.UserInfo,
		input,
	)

	productDescByBPRes := controller.createProductMasterRequestProductDescByBP(
		controller.UserInfo,
		bPPlantRes,
	)

	businessPartnerRes := controller.createBusinessPartnerRequest(
		controller.UserInfo,
		bPPlantRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		bPPlantRes,
	)

	controller.fin(
		generalRes,
		bPPlantRes,
		productDescByBPRes,
		businessPartnerRes,
		productDocRes,
		plantRes,
	)
}

func (
	controller *ProductMasterDetailBPPlantController,
) fin(
	generalRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	bPPlantRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productDescByBPRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
) {
	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	productDescByBPMapper := services.ProductDescByBPMapper(
		productDescByBPRes.Message.ProductDescByBP,
	)

	data := apiOutputFormatter.ProductMaster{}

	for _, v := range *generalRes.Message.General {
		//img := services.ReadProductImage(
		//	productDocRes,
		//	v.BusinessPartner,
		//	v.Product,
		//)

		productDescription := fmt.Sprintf("%s", productDescByBPMapper[v.Product].ProductDescription)

		data.ProductMasterGeneral = append(data.ProductMasterGeneral,
			apiOutputFormatter.ProductMasterGeneral{
				Product:           v.Product,
				ProductName:       &productDescription,
				ProductGroup:      v.ProductGroup,
				BaseUnit:          v.BaseUnit,
				ValidityStartDate: v.ValidityStartDate,
				//Images: apiOutputFormatter.Images{
				//	Product: img,
				//},
			},
		)
	}

	for _, v := range *generalRes.Message.BPPlant {
		data.ProductMasterDetailBPPlant = append(data.ProductMasterDetailBPPlant,
			apiOutputFormatter.ProductMasterDetailBPPlant{
				BusinessPartner:                           v.BusinessPartner,
				BusinessPartnerName:                       businessPartnerMapper[v.BusinessPartner].BusinessPartnerName,
				Plant:                                     v.Plant,
				PlantName:                                 plantMapper[v.Plant].PlantName,
				MRPType:                                   v.MRPType,
				MRPController:                             v.MRPController,
				ReorderThresholdQuantityInBaseUnit:        v.ReorderThresholdQuantityInBaseUnit,
				PlanningTimeFenceInDays:                   v.PlanningTimeFenceInDays,
				MRPPlanningCalendar:                       v.MRPPlanningCalendar,
				SafetyStockQuantityInBaseUnit:             v.SafetyStockQuantityInBaseUnit,
				SafetyDuration:                            v.SafetyDuration,
				SafetyDurationUnit:                        v.SafetyDurationUnit,
				MaximumStockQuantityInBaseUnit:            v.MaximumStockQuantityInBaseUnit,
				MinimumDeliveryQuantityInBaseUnit:         v.MinimumDeliveryQuantityInBaseUnit,
				MinimumDeliveryLotSizeQuantityInBaseUnit:  v.MinimumDeliveryLotSizeQuantityInBaseUnit,
				StandardDeliveryQuantityInBaseUnit:        v.StandardDeliveryQuantityInBaseUnit,
				StandardDeliveryLotSizeQuantityInBaseUnit: v.StandardDeliveryLotSizeQuantityInBaseUnit,
				MaximumDeliveryQuantityInBaseUnit:         v.MaximumDeliveryQuantityInBaseUnit,
				MaximumDeliveryLotSizeQuantityInBaseUnit:  v.MaximumDeliveryLotSizeQuantityInBaseUnit,
				DeliveryLotSizeRoundingQuantityInBaseUnit: v.DeliveryLotSizeRoundingQuantityInBaseUnit,
				DeliveryLotSizeIsFixed:                    v.DeliveryLotSizeIsFixed,
				StandardDeliveryDuration:                  v.StandardDeliveryDuration,
				StandardDeliveryDurationUnit:              v.StandardDeliveryDurationUnit,
				IsBatchManagementRequired:                 v.IsBatchManagementRequired,
				BatchManagementPolicy:                     v.BatchManagementPolicy,
				ProfitCenter:                              v.ProfitCenter,
				CreationDate:                              v.CreationDate,
				LastChangeDate:                            v.LastChangeDate,
				IsMarkedForDeletion:                       v.IsMarkedForDeletion,
				//Images: apiOutputFormatter.Images{
				//	Product: img,
				//},
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
