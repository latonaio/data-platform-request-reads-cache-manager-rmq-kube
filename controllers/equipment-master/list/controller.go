package controllersEquipmentMasterList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsEquipmentMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/equipment-master"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesResponsesEquipmentMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/equipment-master"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type EquipmentMasterListController struct {
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

func (controller *EquipmentMasterListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "equipment-master"
	redisKeyCategory2 := "list"
	userType := controller.GetString(":userType")

	equipmentMasterGeneral := apiInputReader.EquipmentMaster{
		EquipmentMasterGeneral: &apiInputReader.EquipmentMasterGeneral{},
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
		var responseData apiOutputFormatter.EquipmentMaster

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
			controller.request(equipmentMasterGeneral)
		}()
	} else {
		controller.request(equipmentMasterGeneral)
	}
}

func (
	controller *EquipmentMasterListController,
) createEquipmentMasterRequestGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.EquipmentMaster,
) *apiModuleRuntimesResponsesEquipmentMaster.EquipmentMasterRes {
	responseJsonData := apiModuleRuntimesResponsesEquipmentMaster.EquipmentMasterRes{}
	responseBody := apiModuleRuntimesRequestsEquipmentMaster.EquipmentMasterReadsGenerals(
		requestPram,
		apiInputReader.EquipmentMaster{},
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createEquipmentMasterRequestGenerals Unmarshal error")
	}

	return &responseJsonData
}

//func (
//	controller *EquipmentMasterListController,
//) createEquipmentMasterDocRequest(
//	requestPram *apiInputReader.Request,
//) *apiModuleRuntimesResponsesEquipmentMasterDoc.EquipmentMasterDocRes {
//	responseJsonData := apiModuleRuntimesResponsesEquipmentMasterDoc.EquipmentMasterDocRes{}
//	responseBody := apiModuleRuntimesRequests.EquipmentMasterDocReads(
//		requestPram,
//		&controller.Controller,
//	)
//
//	err := json.Unmarshal(responseBody, &responseJsonData)
//	if err != nil {
//		services.HandleError(
//			&controller.Controller,
//			err,
//			nil,
//		)
//		controller.CustomLogger.Error("EquipmentMasterDocReads Unmarshal error")
//	}
//
//	return &responseJsonData
//}

func (
	controller *EquipmentMasterListController,
) request(
	input apiInputReader.EquipmentMaster,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	gsRes := controller.createEquipmentMasterRequestGenerals(
		controller.UserInfo,
		input,
	)

	plRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		gsRes,
	)

	//eMDocRes := controller.createEquipmentMasterDocRequest(
	//	controller.UserInfo,
	//)

	controller.fin(
		gsRes,
		plRes,
		//eMDocRes,
	)
}

func (
	controller *EquipmentMasterListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	bRes *apiModuleRuntimesResponsesEquipmentMaster.EquipmentMasterRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*bRes.Message.General))
	for i, v := range *bRes.Message.General {
		input[i].Plant = v.MaintenancePlant
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
	controller *EquipmentMasterListController,
) fin(
	gsRes *apiModuleRuntimesResponsesEquipmentMaster.EquipmentMasterRes,
	plRes *apiModuleRuntimesResponsesPlant.PlantRes,
	// eMDocRes *apiModuleRuntimesResponsesEquipmentMasterDoc.EquipmentMasterDocRes,
) {
	plantMapper := services.PlantMapper(
		plRes.Message.General,
	)

	data := apiOutputFormatter.EquipmentMaster{}

	for _, v := range *gsRes.Message.General {

		data.EquipmentMasterGeneral = append(data.EquipmentMasterGeneral,
			apiOutputFormatter.EquipmentMasterGeneral{
				Equipment:     v.Equipment,
				EquipmentName: v.EquipmentName,
				EquipmentType: *v.EquipmentType,
				//				EquipmentTypeName:       v.EquipmentTypeName,
				MaintenancePlant:     v.MaintenancePlant,
				MaintenancePlantName: plantMapper[v.MaintenancePlant].PlantName,
				ValidityStartDate:    v.ValidityStartDate,
				IsMarkedForDeletion:  v.IsMarkedForDeletion,
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
