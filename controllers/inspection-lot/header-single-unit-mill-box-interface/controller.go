package controllersInspectionLotSingleUnitMillBoxInterface

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsInspectionLot "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/inspection-lot/inspection-lot"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesInspectionLot "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/inspection-lot"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type InspectionLotSingleUnitMillBoxInterfaceController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *InspectionLotSingleUnitMillBoxInterfaceController) Get() {
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "inspection-lot"
	redisKeyCategory2 := "inspection-lot-header-single-unit-mill-box-interface"
	inspectionLot, _ := controller.GetInt("inspectionLot")

	InspectionLotSingleUnitMillBoxInterface := apiInputReader.InspectionLot{}

	InspectionLotSingleUnitMillBoxInterface = apiInputReader.InspectionLot{
		InspectionLotHeader: &apiInputReader.InspectionLotHeader{
			InspectionLot: inspectionLot,
		},
		InspectionLotSpecDetails: &apiInputReader.InspectionLotSpecDetails{
			InspectionLot: inspectionLot,
		},
		InspectionLotComponentCompositions: &apiInputReader.InspectionLotComponentCompositions{
			InspectionLot: inspectionLot,
		},
		InspectionLotInspections: &apiInputReader.InspectionLotInspections{
			InspectionLot: inspectionLot,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
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
			controller.request(InspectionLotSingleUnitMillBoxInterface)
		}()
	} else {
		controller.request(InspectionLotSingleUnitMillBoxInterface)
	}

	controller.request(InspectionLotSingleUnitMillBoxInterface)

}

func (
	controller *InspectionLotSingleUnitMillBoxInterfaceController,
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
	controller *InspectionLotSingleUnitMillBoxInterfaceController,
) createInspectionLotRequestSpecDetails(
	requestPram *apiInputReader.Request,
	input apiInputReader.InspectionLot,
) *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes {
	responseJsonData := apiModuleRuntimesResponsesInspectionLot.InspectionLotRes{}
	responseBody := apiModuleRuntimesRequestsInspectionLot.InspectionLotReads(
		requestPram,
		input,
		&controller.Controller,
		"SpecDetails",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createInspectionLotRequestSpecDetails Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *InspectionLotSingleUnitMillBoxInterfaceController,
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
	controller *InspectionLotSingleUnitMillBoxInterfaceController,
) createInspectionLotRequestInspections(
	requestPram *apiInputReader.Request,
	input apiInputReader.InspectionLot,
) *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes {
	responseJsonData := apiModuleRuntimesResponsesInspectionLot.InspectionLotRes{}
	responseBody := apiModuleRuntimesRequestsInspectionLot.InspectionLotReads(
		requestPram,
		input,
		&controller.Controller,
		"Inspections",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createInspectionLotRequestInspections Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *InspectionLotSingleUnitMillBoxInterfaceController,
) createInspectionLotRequestMillBoxInterface(
	requestPram *apiInputReader.Request,
	input apiInputReader.InspectionLot,
) *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes {
	responseJsonData := apiModuleRuntimesResponsesInspectionLot.InspectionLotRes{}
	responseBody := apiModuleRuntimesRequestsInspectionLot.InspectionLotReads(
		requestPram,
		input,
		&controller.Controller,
		"MillBoxInterface",
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
	controller *InspectionLotSingleUnitMillBoxInterfaceController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	inspectionLotRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*inspectionLotRes.Message.Header))

	for _, v := range *inspectionLotRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsPlant.General{
			BusinessPartner: v.InspectionPlantBusinessPartner,
			Plant:           v.InspectionPlant,
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
		controller.CustomLogger.Error("PlantReadsGeneralsByPlants Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *InspectionLotSingleUnitMillBoxInterfaceController,
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
	controller *InspectionLotSingleUnitMillBoxInterfaceController,
) request(
	input apiInputReader.InspectionLot,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	inspectionLotHeaderRes := *controller.createInspectionLotRequestHeader(
		controller.UserInfo,
		input,
	)

	inspectionLotSpecDetailsRes := controller.createInspectionLotRequestSpecDetails(
		controller.UserInfo,
		input,
	)

	inspectionLotComponentCompositionsRes := controller.createInspectionLotRequestComponentCompositions(
		controller.UserInfo,
		input,
	)

	inspectionLotInspectionsRes := controller.createInspectionLotRequestInspections(
		controller.UserInfo,
		input,
	)

	inspectionLotMillBoxInterfaceRes := *controller.createInspectionLotRequestMillBoxInterface(
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

	controller.fin(
		&inspectionLotHeaderRes,
		inspectionLotSpecDetailsRes,
		inspectionLotComponentCompositionsRes,
		inspectionLotInspectionsRes,
		&inspectionLotMillBoxInterfaceRes,
		&businessPartnerRes,
		plantRes,
	)
}

func (
	controller *InspectionLotSingleUnitMillBoxInterfaceController,
) fin(
	inspectionLotHeaderRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	inspectionLotSpecDetailsRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	inspectionLotComponentCompositionsRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	inspectionLotInspectionsRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	inspectionLotMillBoxInterfaceRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	//plantMapper := services.PlantMapper(
	//	plantRes.Message.General,
	//)

	data := apiOutputFormatter.InspectionLot{}

	for _, v := range *inspectionLotHeaderRes.Message.Header {
		data.InspectionLotHeader = append(data.InspectionLotHeader,
			apiOutputFormatter.InspectionLotHeader{
				InspectionLot:                      v.InspectionLot,
				InspectionPlantBusinessPartnerName: businessPartnerMapper[v.InspectionPlantBusinessPartner].BusinessPartnerName,
				InspectionLotDate:                  v.InspectionLotDate,
				InspectionSpecification:            v.InspectionSpecification,
				Product:                            v.Product,
				ProductionOrder:                    v.ProductionOrder,
				ProductionOrderItem:                v.ProductionOrderItem,
			},
		)
	}

	for _, v := range *inspectionLotSpecDetailsRes.Message.SpecDetail {
		data.InspectionLotSpecDetail = append(data.InspectionLotSpecDetail,
			apiOutputFormatter.InspectionLotSpecDetail{
				InspectionLot:   v.InspectionLot,
				SpecType:        v.SpecType,
				UpperLimitValue: v.UpperLimitValue,
				LowerLimitValue: v.LowerLimitValue,
				StandardValue:   v.StandardValue,
				SpecTypeUnit:    v.SpecTypeUnit,
			},
		)
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

	for _, v := range *inspectionLotInspectionsRes.Message.Inspection {
		data.InspectionLotInspection = append(data.InspectionLotInspection,
			apiOutputFormatter.InspectionLotInspection{
				InspectionLot:                            v.InspectionLot,
				Inspection:                               v.Inspection,
				InspectionType:                           v.InspectionType,
				InspectionTypeCertificateValueInText:     v.InspectionTypeCertificateValueInText,
				InspectionTypeCertificateValueInQuantity: v.InspectionTypeCertificateValueInQuantity,
				InspectionTypeValueUnit:                  v.InspectionTypeValueUnit,
			},
		)
	}

	// accepter 対応
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
