package controllersWorkCenterList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesRequestsWorkCenter "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/work-center"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesWorkCenter "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/work-center"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type WorkCenterListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const (
	BusinessPartner = "businessPartner"
)

func (controller *WorkCenterListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "workCenter"
	redisKeyCategory2 := "list"
	//userType := BusinessPartner
	workCenter, _ := controller.GetInt("workCenter")
	userType := controller.GetString(":userType")

	workCenterGeneral := apiInputReader.WorkCenter{
		WorkCenterGeneral: &apiInputReader.WorkCenterGeneral{
			WorkCenter:          workCenter,
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
		var responseData apiOutputFormatter.WorkCenter

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
			controller.request(workCenterGeneral)
		}()
	} else {
		controller.request(workCenterGeneral)
	}
}

func (
	controller *WorkCenterListController,
) createWorkCenterRequestHeaderByBusinessPartner(
	requestPram *apiInputReader.Request,
	input apiInputReader.WorkCenter,
) *apiModuleRuntimesResponsesWorkCenter.WorkCenterRes {
	responseJsonData := apiModuleRuntimesResponsesWorkCenter.WorkCenterRes{}
	responseBody := apiModuleRuntimesRequestsWorkCenter.WorkCenterReads(
		requestPram,
		input,
		&controller.Controller,
		"HeaderByBusinessPartner",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("WorkCenterReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *WorkCenterListController,
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
	controller *WorkCenterListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	workCenterRes *apiModuleRuntimesResponsesWorkCenter.WorkCenterRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*workCenterRes.Message.General))
	for i, v := range *workCenterRes.Message.General {
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
	controller *WorkCenterListController,
) request(
	input apiInputReader.WorkCenter,
) {
	defer services.Recover(controller.CustomLogger)

	generalRes := controller.createWorkCenterRequestHeaderByBusinessPartner(
		controller.UserInfo,
		input,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		generalRes,
	)

	controller.fin(
		generalRes,
		plantRes,
	)
}

func (
	controller *WorkCenterListController,
) fin(
	generalRes *apiModuleRuntimesResponsesWorkCenter.WorkCenterRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
) {

	plantMapper := services.PlantMapper(
		plantRes.Message.Generals,
	)

	data := apiOutputFormatter.WorkCenter{}

	for _, v := range *generalRes.Message.General {

		data.WorkCenterGeneral = append(data.WorkCenterGeneral,
			apiOutputFormatter.WorkCenterGeneral{
				WorkCenter:          v.WorkCenter,
				WorkCenterName:      v.WorkCenterName,
				Plant:               v.Plant,
				PlantName:           plantMapper[v.Plant].PlantName,
				WorkCenterLocation:  v.WorkCenterLocation,
				CapacityCategory:    v.CapacityCategory,
				ValidityStartDate:   v.ValidityStartDate,
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
