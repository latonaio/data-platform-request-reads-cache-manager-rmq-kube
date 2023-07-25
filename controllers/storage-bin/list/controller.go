package controllersStorageBinList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsStorageBin "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/storage-bin"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesStorageBin "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/storage-bin"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type StorageBinListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

//const (
//	buyer	= "buyer"
//	seller	= "seller"
//)

func (controller *StorageBinListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "storage-bin"
	redisKeyCategory2 := "list"
	//userType :=
	businessPartner, _ := controller.GetInt("businessPartner")
	plant := controller.GetString("plant")
	storageLocation := controller.GetString("storageLocation")
	storageBin := controller.GetString("storageBin")
	userType := controller.GetString(":userType")

	storageBinGeneral := apiInputReader.StorageBin{
		StorageBinGeneral: &apiInputReader.StorageBinGeneral{
			BusinessPartner:     businessPartner,
			Plant:               plant,
			StorageLocation:     storageLocation,
			StorageBin:          storageBin,
			IsMarkedForDeletion: &isMarkedForDeletion,
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
		var responseData apiOutputFormatter.StorageBin

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
			controller.request(storageBinGeneral)
		}()
	} else {
		controller.request(storageBinGeneral)
	}
}

func (
	controller *StorageBinListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	storageBinRes *apiModuleRuntimesResponsesStorageBin.StorageBinRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*storageBinRes.Message.General))

	for _, v := range *storageBinRes.Message.General {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.BusinessPartner,
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
		controller.CustomLogger.Error("BusinessPartnerGeneralReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *StorageBinListController,
) createStorageBinRequestGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.StorageBin,
) *apiModuleRuntimesResponsesStorageBin.StorageBinRes {
	responseJsonData := apiModuleRuntimesResponsesStorageBin.StorageBinRes{}
	responseBody := apiModuleRuntimesRequestsStorageBin.StorageBinReadsGenerals(
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
		controller.CustomLogger.Error("StorageBinReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *StorageBinListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	storageBinRes *apiModuleRuntimesResponsesStorageBin.StorageBinRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*storageBinRes.Message.General))
	for i, v := range *storageBinRes.Message.General {
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
	controller *StorageBinListController,
) request(
	input apiInputReader.StorageBin,
) {
	defer services.Recover(controller.CustomLogger)

	generalRes := controller.createStorageBinRequestGenerals(
		controller.UserInfo,
		input,
	)

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		generalRes,
	)

	plantRes := *controller.createPlantRequestGenerals(
		controller.UserInfo,
		generalRes,
	)

	controller.fin(
		generalRes,
		&businessPartnerRes,
		&plantRes,
	)
}

func (
	controller *StorageBinListController,
) fin(
	generalRes *apiModuleRuntimesResponsesStorageBin.StorageBinRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)
	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)
	//	storageLocationMapper := services.StorageLocationMapper(
	//		storageLocationRes.Message.StorageLocations,
	//	)

	data := apiOutputFormatter.StorageBin{}

	for _, v := range *generalRes.Message.General {

		data.StorageBinGeneral = append(data.StorageBinGeneral,
			apiOutputFormatter.StorageBinGeneral{
				BusinessPartner:     v.BusinessPartner,
				BusinessPartnerName: businessPartnerMapper[v.BusinessPartner].BusinessPartnerName,
				Plant:               v.Plant,
				PlantName:           plantMapper[v.Plant].PlantName,
				StorageLocation:     v.StorageLocation,
				//				StorageLocationName:    storageLocationMapper[v.StorageLocation].StorageLocationName,
				StorageBin:          v.StorageBin,
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
