package controllersPlantDetailList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type PlantDetailListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *PlantDetailListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	businessPartner, _ := controller.GetInt("businessPartner")
	plant := controller.GetString("plant")
	redisKeyCategory1 := "plant"
	redisKeyCategory2 := "detail-list"
	redisKeyCategory3 := plant
	userType := controller.GetString(":userType")

	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	plantStorageLocations := apiInputReader.Plant{
		PlantGeneral: &apiInputReader.PlantGeneral{
			BusinessPartner: businessPartner,
			Plant:           plant,
		},
		PlantStorageLocations: &apiInputReader.PlantStorageLocations{
			BusinessPartner:     businessPartner,
			Plant:               plant,
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
		var responseData apiOutputFormatter.Plant

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
			controller.request(plantStorageLocations)
		}()
	} else {
		controller.request(plantStorageLocations)
	}
}

func (
	controller *PlantDetailListController,
) createPlantRequestGeneral(
	requestPram *apiInputReader.Request,
	input apiInputReader.Plant,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	responseJsonData := apiModuleRuntimesResponsesPlant.PlantRes{}
	responseBody := apiModuleRuntimesRequestsPlant.PlantReadsGenerals(
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
		controller.CustomLogger.Error("createPlantRequestGeneral Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PlantDetailListController,
) createPlantRequestStorageLocations(
	requestPram *apiInputReader.Request,
	input apiInputReader.Plant,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	responseJsonData := apiModuleRuntimesResponsesPlant.PlantRes{}
	responseBody := apiModuleRuntimesRequestsPlant.PlantReadsStorageLocations(
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
	controller *PlantDetailListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*plantRes.Message.General))

	for _, v := range *plantRes.Message.General {
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
	controller *PlantDetailListController,
) request(
	input apiInputReader.Plant,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	generalRes := controller.createPlantRequestGeneral(
		controller.UserInfo,
		input,
	)

	storageLocationRes := controller.createPlantRequestStorageLocations(
		controller.UserInfo,
		input,
	)

	businessPartnerRes := controller.createBusinessPartnerRequest(
		controller.UserInfo,
		generalRes,
	)

	controller.fin(
		generalRes,
		storageLocationRes,
		businessPartnerRes,
	)
}

func (
	controller *PlantDetailListController,
) fin(
	generalRes *apiModuleRuntimesResponsesPlant.PlantRes,
	storageLocationRes *apiModuleRuntimesResponsesPlant.PlantRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.Plant{}

	for _, v := range *generalRes.Message.General {

		data.PlantGeneral = append(data.PlantGeneral,
			apiOutputFormatter.PlantGeneral{
				BusinessPartner:     v.BusinessPartner,
				BusinessPartnerName: businessPartnerMapper[v.BusinessPartner].BusinessPartnerName,
				Plant:               v.Plant,
				PlantName:           v.PlantName,
			},
		)
	}

	for _, v := range *generalRes.Message.StorageLocation {
		data.PlantStorageLocation = append(data.PlantStorageLocation,
			apiOutputFormatter.PlantStorageLocation{
				StorageLocation:     v.StorageLocation,
				StorageLocationName: v.StorageLocationName,
				IsMarkedForDeletion: v.IsMarkedForDeletion,
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
