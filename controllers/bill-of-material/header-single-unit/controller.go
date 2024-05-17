package controllersBillOfMaterialHeaderSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBillOfMaterial "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/bill-of-material/bill-of-material"
	apiModuleRuntimesRequestsBillOfMaterialDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/bill-of-material/bill-of-material-doc"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesBillOfMaterial "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/bill-of-material"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
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

type BillOfMaterialHeaderSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *BillOfMaterialHeaderSingleUnitController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)

	redisKeyCategory1 := "bill-of-material"
	redisKeyCategory2 := "header-single-unit"
	billOfMaterial, _ := controller.GetInt("billOfMaterial")

	billOfMaterialHeader := apiInputReader.BillOfMaterial{
		BillOfMaterialHeader: &apiInputReader.BillOfMaterialHeader{
			BillOfMaterial:      billOfMaterial,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		BillOfMaterialHeaderDoc: &apiInputReader.BillOfMaterialHeaderDoc{
			BillOfMaterial: billOfMaterial,
			//DocType:					&docType,
			DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(billOfMaterial),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.BillOfMaterial

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
			controller.request(billOfMaterialHeader)
		}()
	} else {
		controller.request(billOfMaterialHeader)
	}
}

func (
	controller *BillOfMaterialHeaderSingleUnitController,
) createBillOfMaterialRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.BillOfMaterial,
) *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes {
	responseJsonData := apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes{}
	responseBody := apiModuleRuntimesRequestsBillOfMaterial.BillOfMaterialReads(
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
		controller.CustomLogger.Error("createBillOfMaterialRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BillOfMaterialHeaderSingleUnitController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	bRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, 0)
	isMarkedForDeletion := false

	for _, v := range *bRes.Message.Header {
		productDescsByBP = append(productDescsByBP, apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BusinessPartner: v.OwnerProductionPlantBusinessPartner,
					ProductDescByBP: []apiModuleRuntimesRequestsProductMaster.ProductDescByBP{
						{
							Language:            *requestPram.Language,
							IsMarkedForDeletion: &isMarkedForDeletion,
						},
					},
				},
			},
		})
	}

	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsProductDescsByBP(
		requestPram,
		productDescsByBP,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductMasterRequestProductDescByBP Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BillOfMaterialHeaderSingleUnitController,
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
	controller *BillOfMaterialHeaderSingleUnitController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	billOfMaterialRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*billOfMaterialRes.Message.Header))
	for i, v := range *billOfMaterialRes.Message.Header {
		input[i].Plant = v.OwnerProductionPlant
		input[i].BusinessPartner = v.OwnerProductionPlantBusinessPartner
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
	controller *BillOfMaterialHeaderSingleUnitController,
) createBillOfMaterialDocRequestHeaderDoc(
	requestPram *apiInputReader.Request,
	input apiInputReader.BillOfMaterial,
) *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialDocRes {
	responseJsonData := apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialDocRes{}
	responseBody := apiModuleRuntimesRequestsBillOfMaterialDoc.BillOfMaterialDocReads(
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
		controller.CustomLogger.Error("createBillOfMaterialDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BillOfMaterialHeaderSingleUnitController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	billOfMaterialRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*billOfMaterialRes.Message.Header))

	for _, v := range *billOfMaterialRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.OwnerProductionPlantBusinessPartner,
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
	controller *BillOfMaterialHeaderSingleUnitController,
) request(
	input apiInputReader.BillOfMaterial,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := controller.createBillOfMaterialRequestHeader(
		controller.UserInfo,
		input,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		headerRes,
	)

	productDescByBPRes := controller.createProductMasterRequestProductDescByBP(
		controller.UserInfo,
		headerRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		headerRes,
	)

	headerDocRes := controller.createBillOfMaterialDocRequestHeaderDoc(
		controller.UserInfo,
		input,
	)

	controller.fin(
		headerRes,
		plantRes,
		productDescByBPRes,
		productDocRes,
		&businessPartnerRes,
		headerDocRes,
	)
}

func (
	controller *BillOfMaterialHeaderSingleUnitController,
) fin(
	headerRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	productDescByBPRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	headerDocRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialDocRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	descriptionMapper := services.ProductDescByBPMapper(
		productDescByBPRes.Message.ProductDescByBP,
	)

	data := apiOutputFormatter.BillOfMaterial{}

	for _, v := range *headerRes.Message.Header {
		img := services.ReadProductImage(
			productDocRes,
			v.OwnerProductionPlantBusinessPartner,
			v.Product,
		)

		qrcode := services.CreateQRCodeBillOfMaterialHeaderDocImage(
			headerDocRes,
			v.BillOfMaterial,
		)

		documentImage := services.ReadDocumentImageBillOfMaterial(
			headerDocRes,
			v.BillOfMaterial,
		)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)
		plantName := fmt.Sprintf("%s", plantMapper[strconv.Itoa(v.OwnerProductionPlantBusinessPartner)].PlantName)

		data.BillOfMaterialHeader = append(data.BillOfMaterialHeader,
			apiOutputFormatter.BillOfMaterialHeader{
				BillOfMaterial:                          v.BillOfMaterial,
				Product:                                 v.Product,
				ProductDescription:                      &productDescription,
				OwnerProductionPlant:                    v.OwnerProductionPlant,
				OwnerProductionPlantName:                plantName,
				OwnerBusinessPartner:                    businessPartnerMapper[v.OwnerProductionPlantBusinessPartner].BusinessPartner,
				OwnerBusinessPartnerName:                businessPartnerMapper[v.OwnerProductionPlantBusinessPartner].BusinessPartnerName,
				ProductBaseUnit:                         v.ProductBaseUnit,
				ProductProductionUnit:                   v.ProductProductionUnit,
				ProductStandardQuantityInBaseUnit:       v.ProductStandardQuantityInBaseUnit,
				ProductStandardQuantityInProductionUnit: v.ProductStandardQuantityInProductionUnit,
				ValidityStartDate:                       v.ValidityStartDate,
				IsMarkedForDeletion:                     v.IsMarkedForDeletion,
				Images: apiOutputFormatter.Images{
					Product:                     img,
					QRCode:                      qrcode,
					DocumentImageBillOfMaterial: documentImage,
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
