package controllersInspectionLotSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsInspectionLot "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/inspection-lot/inspection-lot"
	apiModuleRuntimesRequestsInspectionLotDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/inspection-lot/inspection-lot-doc"
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
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type InspectionLotSingleUnitController struct {
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

func (controller *InspectionLotSingleUnitController) Get() {
	//isReleased, _ := controller.GetBool("isReleased")
	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "inspection-lot"
	redisKeyCategory2 := "inspection-lot-header-single-unit"
	inspectionLot, _ := controller.GetInt("inspectionLot")
	//	userType := controller.GetString(":userType")
	//	pBuyer, _ := controller.GetInt("buyer")
	//	pSeller, _ := controller.GetInt("seller")

	InspectionLotSingleUnit := apiInputReader.InspectionLot{}

	isReleased := true
	isMarkedForDeletion := false

	docType := "QRCODE"

	InspectionLotSingleUnit = apiInputReader.InspectionLot{
		InspectionLotHeader: &apiInputReader.InspectionLotHeader{
			InspectionLot:       inspectionLot,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		InspectionLotSpecDetails: &apiInputReader.InspectionLotSpecDetails{
			InspectionLot:       inspectionLot,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		InspectionLotComponentCompositions: &apiInputReader.InspectionLotComponentCompositions{
			InspectionLot:       inspectionLot,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		InspectionLotInspections: &apiInputReader.InspectionLotInspections{
			InspectionLot:       inspectionLot,
			IsReleased:          &isReleased,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		InspectionLotDocHeaderDoc: &apiInputReader.InspectionLotDocHeaderDoc{
			InspectionLot:            inspectionLot,
			DocType:                  &docType,
			DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
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
	controller *InspectionLotSingleUnitController,
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
	controller *InspectionLotSingleUnitController,
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
	controller *InspectionLotSingleUnitController,
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
	controller *InspectionLotSingleUnitController,
) createPlantRequest(
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
	controller *InspectionLotSingleUnitController,
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
	controller *InspectionLotSingleUnitController,
) createInspectionLotDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.InspectionLot,
) *apiModuleRuntimesResponsesInspectionLot.InspectionLotDocRes {
	responseJsonData := apiModuleRuntimesResponsesInspectionLot.InspectionLotDocRes{}
	responseBody := apiModuleRuntimesRequestsInspectionLotDoc.InspectionLotDocReads(
		requestPram,
		input,
		&controller.Controller,
		"HeaderDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createInspectionLotDocRequestHeaderDoc Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *InspectionLotSingleUnitController,
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
	controller *InspectionLotSingleUnitController,
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
	controller *InspectionLotSingleUnitController,
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

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		&inspectionLotHeaderRes,
	)

	plantRes := controller.createPlantRequest(
		controller.UserInfo,
		&inspectionLotHeaderRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	inspectionLotHeaderDocRes := controller.createInspectionLotDocRequest(
		controller.UserInfo,
		input,
	)

	controller.fin(
		&inspectionLotHeaderRes,
		inspectionLotSpecDetailsRes,
		inspectionLotComponentCompositionsRes,
		inspectionLotInspectionsRes,
		&businessPartnerRes,
		plantRes,
		productDocRes,
		inspectionLotHeaderDocRes,
	)
}

func (
	controller *InspectionLotSingleUnitController,
) fin(
	inspectionLotHeaderRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	inspectionLotSpecDetailsRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	inspectionLotComponentCompositionsRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	inspectionLotInspectionsRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	inspectionLotHeaderDocRes *apiModuleRuntimesResponsesInspectionLot.InspectionLotDocRes,
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

		qrcode := services.CreateQRCodeInspectionLotHeaderDocImage(
			inspectionLotHeaderDocRes,
			v.InspectionLot,
		)

		//documentImage := services.ReadDocumentImageInspectionLot(
		//	inspectionLotHeaderDocRes,
		//	v.InspectionLot,
		//)

		data.InspectionLotSingleUnit = append(data.InspectionLotSingleUnit,
			apiOutputFormatter.InspectionLotSingleUnit{
				InspectionLot:                      v.InspectionLot,
				InspectionLotDate:                  v.InspectionLotDate,
				InspectionPlantBusinessPartner:     v.InspectionPlantBusinessPartner,
				InspectionPlantBusinessPartnerName: businessPartnerMapper[v.InspectionPlantBusinessPartner].BusinessPartnerName,
				InspectionPlant:                    v.InspectionPlant,
				InspectionPlantName:                plantMapper[strconv.Itoa(v.InspectionPlantBusinessPartner)].PlantName,
				Product:                            v.Product,
				ProductSpecification:               v.ProductSpecification,
				ProductionOrder:                    v.ProductionOrder,
				ProductionOrderItem:                v.ProductionOrderItem,
				UsageControlChain:                  v.UsageControlChain,
				CertificateAuthorityChain:          v.CertificateAuthorityChain,

				Images: apiOutputFormatter.Images{
					Product: img,
					QRCode:  qrcode,
					//DocumentImageInspectionLot: documentImage,
				},
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
