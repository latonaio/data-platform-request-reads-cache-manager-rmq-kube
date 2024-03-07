package controllersInspectionLotComponentComposition

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsInspectionLot "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/inspection-lot/inspection-lot"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesInspectionLot "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/inspection-lot"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type InspectionLotComponentCompositionController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

//const (
//	buyer  = "buyer"
//	seller = "seller"
//)

func (controller *InspectionLotComponentCompositionController) Get() {
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "inspection-lot"
	redisKeyCategory2 := "inspection-lot-component-composition"
	inspectionLot, _ := controller.GetInt("inspectionLot")
	//userType := controller.GetString(":userType")
	//	pBuyer, _ := controller.GetInt("buyer")
	//	pSeller, _ := controller.GetInt("seller")

	InspectionLotSingleUnit := apiInputReader.InspectionLot{}

	IsReleased := true
	isMarkedForDeletion := false

	InspectionLotSingleUnit = apiInputReader.InspectionLot{
		InspectionLotHeader: &apiInputReader.InspectionLotHeader{
			InspectionLot:       inspectionLot,
			IsReleased:          &IsReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		InspectionLotComponentCompositions: &apiInputReader.InspectionLotComponentCompositions{
			InspectionLot:       inspectionLot,
			IsReleased:          &IsReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(inspectionLot),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.InspectionLot

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
			controller.request(InspectionLotSingleUnit)
		}()
	} else {
		controller.request(InspectionLotSingleUnit)
	}
}

func (
	controller *InspectionLotComponentCompositionController,
) createInspectionLotRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.InspectionLot,
) *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes {
	responseJsonData := apiModuleRuntimesResponsesInspectionLot.InspectionLotRes{}
	responseBody := apiModuleRuntimesRequestsInspectionLot.InspectionLotReads(
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
		controller.CustomLogger.Error("InspectionLotReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *InspectionLotComponentCompositionController,
) createInspectionLotRequestComponentCompositions(
	requestPram *apiInputReader.Request,
	input apiInputReader.InspectionLot,
) *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes {
	responseJsonData := apiModuleRuntimesResponsesInspectionLot.InspectionLotRes{}
	responseBody := apiModuleRuntimesRequestsInspectionLot.InspectionLotReads(
		requestPram,
		input,
		&controller.Controller,
		"ComponentCompositions",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createInspectionLotRequestComponentCompositions Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *InspectionLotComponentCompositionController,
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
	controller *InspectionLotComponentCompositionController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	inspectionLotHeaderRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*inspectionLotHeaderRes.Message.Header))

	for _, v := range *inspectionLotHeaderRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.InspectionPlantBusinessPartner,
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
	controller *InspectionLotComponentCompositionController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	InspectionLotHeaderRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*InspectionLotHeaderRes.Message.Header))
	for i, v := range *InspectionLotHeaderRes.Message.Header {
		input[i].Plant = v.InspectionPlant
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
	controller *InspectionLotComponentCompositionController,
) request(
	input apiInputReader.InspectionLot,
) {
	defer services.Recover(controller.CustomLogger)

	inspectionLotHeaderRes := *controller.createInspectionLotRequestHeader(
		controller.UserInfo,
		input,
	)

	inspectionLotComponentCompositionsRes := controller.createInspectionLotRequestComponentCompositions(
		controller.UserInfo,
		input,
	)

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		&inspectionLotHeaderRes,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		&inspectionLotHeaderRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	controller.fin(
		&inspectionLotHeaderRes,
		inspectionLotComponentCompositionsRes,
		&businessPartnerRes,
		plantRes,
		productDocRes,
	)
}

func (
	controller *InspectionLotComponentCompositionController,
) fin(
	inspectionLotHeaderRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	inspectionLotComponentCompositionsRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	data := apiOutputFormatter.InspectionLot{}

	for _, v := range *inspectionLotHeaderRes.Message.Header {
		img := services.ReadProductImage(
			productDocRes,
			*controller.UserInfo.BusinessPartner,
			v.Product,
		)

		inspectionPlantName := fmt.Sprintf("%s", plantMapper[v.InspectionPlant].PlantName)

		data.InspectionLotSingleUnit = append(data.InspectionLotSingleUnit, apiOutputFormatter.InspectionLotSingleUnit{
			InspectionLot:                      v.InspectionLot,
			InspectionLotDate:                  v.InspectionLotDate,
			InspectionPlantBusinessPartner:     v.InspectionPlantBusinessPartner,
			InspectionPlantBusinessPartnerName: businessPartnerMapper[v.InspectionPlantBusinessPartner].BusinessPartnerName,
			InspectionPlant:                    v.InspectionPlant,
			InspectionPlantName:                inspectionPlantName,
			Product:                            v.Product,
			ProductSpecification:               v.ProductSpecification,
			ProductionOrder:                    v.ProductionOrder,
			ProductionOrderItem:                v.ProductionOrderItem,
			Images: apiOutputFormatter.Images{
				Product: img,
			},
		})
	}

	for _, v := range *inspectionLotComponentCompositionsRes.Message.ComponentComposition {
		data.InspectionLotComponentComposition = append(data.InspectionLotComponentComposition,
			apiOutputFormatter.InspectionLotComponentComposition{
				InspectionLot:                              v.InspectionLot,
				ComponentCompositionType:                   v.ComponentCompositionType,
				ComponentCompositionUpperLimitInPercent:    v.ComponentCompositionUpperLimitInPercent,
				ComponentCompositionLowerLimitInPercent:    v.ComponentCompositionLowerLimitInPercent,
				ComponentCompositionStandardValueInPercent: v.ComponentCompositionStandardValueInPercent,
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
