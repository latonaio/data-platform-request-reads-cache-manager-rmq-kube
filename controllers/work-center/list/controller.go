package controllersWorkCenterList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequests "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests"
	apiModuleRuntimesRequestsWorkCenter "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/work-center"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/produ\ct-master"
	apiModuleRuntimesResponsesWorkCenter "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/work-center"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master-doc"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"io/ioutil"
	"strings"
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
	userType := controller.GetString("userType")

	workCenter := apiInputReader.WorkCenter{
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
) *apiModuleRuntimesResponsesProductMasterDoc.ProductMasterDocRes {
	responseJsonData := apiModuleRuntimesResponsesProductMasterDoc.ProductMasterDocRes{}
	responseBody := apiModuleRuntimesRequests.ProductMasterDocReads(
		requestPram,
		&controller.Controller,
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
) request(
	input apiInputReader.WorkCenter,
) {
	defer services.Recover(controller.CustomLogger)

	bRes := controller.createWorkCenterRequestHeaderByBusinessPartner(
		controller.UserInfo,
		input,
	)

	plRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		bRes,
	)

	controller.fin(
		bRes,
		plRes,
	)
}

func (
	controller *WorkCenterListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	workCenterRes *apiModuleRuntimesResponsesWorkCenter.WorkCenterRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	generals := make(apiModuleRuntimesRequests.PlantGenerals, len(*workCenterRes.Message.Header))
	for i, v := range *workCenterRes.Message.Header {
		generals[i].Plant = &v.Plant
		generals[i].Language = requestPram.Language
	}

	aPIServiceName := "DPFM_API_PLANT_SRV"
	aPIType := "reads"
	responseJsonData := apiModuleRuntimesResponsesPlant.PlantRes{}

	request := apiModuleRuntimesRequests.
		CreatePlantRequestGenerals(
			requestPram,
			generals,
		)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		&controller.Controller,
	)

	err = json.Unmarshal(responseBody, &responseJsonData)
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
) fin(
	bRes *apiModuleRuntimesResponsesWorkCenter.WorkCenterRes,
	plRes *apiModuleRuntimesResponsesPlant.PlantRes,
) {

	plantMapper := services.PlantMapper(
		plRes.Message.Generals,
	)

	data := apiOutputFormatter.WorkCenter{}

	for _, v := range *bRes.Message.Header {

		data.WorkCenterGeneral = append(data.WorkCenterGeneral,
			apiOutputFormatter.WorkCenterGeneral{
				WorkCenter:               			&v.WorkCenter,
				WorkCenterName:						v.WorkCenterName,
				PlantName:				  			plantMapper[v.Plant].PlantName,
				ComponentIsMarkedForBackflush:		v.ComponentIsMarkedForBackflush,
				CapacityInternalID:      			v.CapacityInternalID,
				CapacityCategoryCode:      			v.CapacityCategoryCode,
				IsMarkedForDeletion:      			v.IsMarkedForDeletion,
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
