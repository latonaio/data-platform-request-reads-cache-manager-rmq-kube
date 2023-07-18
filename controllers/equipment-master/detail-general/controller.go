package controllersEquipmentMasterDetailGeneral

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsEquipmentMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/equipment-master"
	apiModuleRuntimesRequests "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesEquipmentMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/equipment-master"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"io/ioutil"
	"strconv"
	"strings"
)

type EquipmentMasterDetailGeneralController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *EquipmentMasterDetailGeneralController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	equipmentMaster, _ := controller.GetInt("equipmentMaster")
	redisKeyCategory1 := "equipment-master"
	redisKeyCategory2 := "detail-general"
	redisKeyCategory3 := equipmentMaster
	userType := controller.GetString(":userType")

	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	equipmentMasterItems := apiInputReader.EquipmentMaster{
		EquipmentMasterGeneral: &apiInputReader.EquipmentMasterGeneral{
			Equipment: equipmentMaster,
		},
		//EquipmentMasterItems: &apiInputReader.EquipmentMasterItems{
		//	EquipmentMaster:      equipmentMaster,
		//	IsMarkedForDeletion: &isMarkedForDeletion,
		//},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(redisKeyCategory3),
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
			controller.request(equipmentMasterItems)
		}()
	} else {
		controller.request(equipmentMasterItems)
	}
}

func (
	controller *EquipmentMasterDetailGeneralController,
) createEquipmentMasterRequestGeneral(
	requestPram *apiInputReader.Request,
	input apiInputReader.EquipmentMaster,
) *apiModuleRuntimesResponsesEquipmentMaster.EquipmentMasterRes {
	equipmentMasterGenerals := make([]apiModuleRuntimesRequestsEquipmentMaster.General, 0)

	equipmentMasterGenerals = append(equipmentMasterGenerals, apiModuleRuntimesRequestsEquipmentMaster.General{
		Equipment: input.EquipmentMasterGeneral.Equipment,
	})

	responseJsonData := apiModuleRuntimesResponsesEquipmentMaster.EquipmentMasterRes{}
	responseBody := apiModuleRuntimesRequestsEquipmentMaster.EquipmentMasterReads(
		requestPram,
		equipmentMasterGenerals,
		&controller.Controller,
		"Generals",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createEquipmentMasterRequestGeneral Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *EquipmentMasterDetailGeneralController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	equipmentMasterResRes *apiModuleRuntimesResponsesEquipmentMaster.EquipmentMasterRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, 0)
	for i, v := range *equipmentMasterResRes.Message.General {
		input[i].Plant = v.MaintenancePlant
	}

	aPIServiceName := "DPFM_API_PLANT_SRV"
	aPIType := "reads"
	responseJsonData := apiModuleRuntimesResponsesPlant.PlantRes{}

	request := apiModuleRuntimesRequests.PlantReadsGeneralsByPlants(
		requestPram,
		input,
		&controller.Controller,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createPlantRequestGenerals error")
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
		controller.CustomLogger.Error("createPlantRequestGenerals error")
	}

	return &responseJsonData
}

func (
	controller *EquipmentMasterDetailGeneralController,
) createBusinessPartnerRequestGeneralsByBusinessPartners(
	requestPram *apiInputReader.Request,
	equipmentMasterRes *apiModuleRuntimesResponsesEquipmentMaster.EquipmentMasterRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, 0)

	for _, v := range *equipmentMasterRes.Message.General {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.MaintenancePlantBusinessPartner,
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
	controller *EquipmentMasterDetailGeneralController,
) request(
	input apiInputReader.EquipmentMaster,
) {
	defer services.Recover(controller.CustomLogger)

	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	equipmentGeneralRes := controller.createEquipmentMasterRequestGeneral(
		controller.UserInfo,
		input,
	)

	businessPartnerRes = *controller.createBusinessPartnerRequestGeneralsByBusinessPartners(
		controller.UserInfo,
		equipmentGeneralRes,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		equipmentGeneralRes,
	)

	controller.fin(
		equipmentGeneralRes,
		&businessPartnerRes,
		plantRes,
	)
}

func (
	controller *EquipmentMasterDetailGeneralController,
) fin(
	equipmentGeneralRes *apiModuleRuntimesResponsesEquipmentMaster.EquipmentMasterRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	plantMapper := services.PlantMapper(
		plantRes.Message.Generals,
	)

	data := apiOutputFormatter.EquipmentMaster{}

	for _, v := range *equipmentGeneralRes.Message.General {
		data.EquipmentMasterGeneral = append(data.EquipmentMasterGeneral,
			apiOutputFormatter.EquipmentMasterGeneral{
				Equipment:     v.Equipment,
				EquipmentName: v.EquipmentName,
				EquipmentType: *v.EquipmentType,
				//				EquipmentTypeName:       v.EquipmentTypeName,
				MaintenancePlant:     v.MaintenancePlant,
				MaintenancePlantName: plantMapper[v.MaintenancePlant].PlantName,
				ValidityStartDate:    v.ValidityStartDate,
				//Images: apiOutputFormatter.Images{
				//	Product: img,
				//},
			},
		)

		data.EquipmentMasterDetailGeneral = append(data.EquipmentMasterDetailGeneral,
			apiOutputFormatter.EquipmentMasterDetailGeneral{
				EquipmentCategory:                   v.EquipmentCategory,
				TechnicalObjectType:                 v.TechnicalObjectType,
				GrossWeight:                         v.GrossWeight,
				NetWeight:                           v.NetWeight,
				WeightUnit:                          v.WeightUnit,
				SizeOrDimensionText:                 v.SizeOrDimensionText,
				InventoryNumber:                     v.InventoryNumber,
				OperationStartDate:                  v.OperationStartDate,
				OperationStartTime:                  v.OperationStartTime,
				OperationEndDate:                    v.OperationEndDate,
				OperationEndTime:                    v.OperationEndTime,
				EquipmentStandardID:                 v.EquipmentStandardID,
				EquipmentIndustryStandardName:       v.EquipmentIndustryStandardName,
				AcquisitionDate:                     v.AcquisitionDate,
				Manufacturer:                        v.Manufacturer,
				ManufacturerCountry:                 v.ManufacturerCountry,
				ManufacturerPartNmbr:                v.ManufacturerPartNmbr,
				ManufacturerSerialNumber:            v.ManufacturerSerialNumber,
				MaintenancePlantBusinessPartner:     v.MaintenancePlantBusinessPartner,
				MaintenancePlantBusinessPartnerName: businessPartnerMapper[v.MaintenancePlantBusinessPartner].BusinessPartnerName,
				MaintenancePlant:                    v.MaintenancePlant,
				MaintenancePlantName:                plantMapper[v.MaintenancePlant].PlantName,
				WorkCenter:                          v.WorkCenter,
				Project:                             v.Project,
				WBSElement:                          v.WBSElement,
				FunctionalLocation:                  v.FunctionalLocation,
				SuperordinateEquipment:              v.SuperordinateEquipment,
				EquipmentIsAvailable:                v.EquipmentIsAvailable,
				EquipmentIsInstalled:                v.EquipmentIsInstalled,
				EquipHasSubOrdinateEquipment:        v.EquipHasSubOrdinateEquipment,
				MasterFixedAsset:                    v.MasterFixedAsset,
				FixedAsset:                          v.FixedAsset,
				CreationDate:                        v.CreationDate,
				LastChangeDate:                      v.LastChangeDate,
				IsMarkedForDeletion:                 v.IsMarkedForDeletion,
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
