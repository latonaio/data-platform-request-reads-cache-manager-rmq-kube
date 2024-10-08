package controllersInspectionLotSpecDetail

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
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

type InspectionLotSpecDetailController struct {
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

func (controller *InspectionLotSpecDetailController) Get() {
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "inspection-lot"
	redisKeyCategory2 := "inspection-lot-spec-detail"
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
		InspectionLotSpecDetails: &apiInputReader.InspectionLotSpecDetails{
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
	controller *InspectionLotSpecDetailController,
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
	controller *InspectionLotSpecDetailController,
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
	controller *InspectionLotSpecDetailController,
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
	controller *InspectionLotSpecDetailController,
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
	controller *InspectionLotSpecDetailController,
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
	controller *InspectionLotSpecDetailController,
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
		inspectionLotSpecDetailsRes,
		&businessPartnerRes,
		plantRes,
		productDocRes,
	)
}

func (
	controller *InspectionLotSpecDetailController,
) fin(
	inspectionLotHeaderRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	inspectionLotSpecDetailsRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
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

	for _, v := range *inspectionLotSpecDetailsRes.Message.SpecDetail {
		data.InspectionLotSpecDetail = append(data.InspectionLotSpecDetail,
			apiOutputFormatter.InspectionLotSpecDetail{
				InspectionLot:   v.InspectionLot,
				SpecType:        v.SpecType,
				UpperLimitValue: v.UpperLimitValue,
				LowerLimitValue: v.LowerLimitValue,
				StandardValue:   v.StandardValue,
				SpecTypeUnit:    v.SpecTypeUnit,
				Formula:         v.Formula,
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
