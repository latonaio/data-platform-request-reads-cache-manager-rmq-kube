package controllersBillOfMaterialList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBillOfMaterial "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/bill-of-material"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesBillOfMaterial "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/bill-of-material"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type BillOfMaterialListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const (
	OwnerProductionPlantBusinessPartner = "ownerProductionPlantBusinessPartner"
)

func (controller *BillOfMaterialListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)

	redisKeyCategory1 := "bill-of-material"
	redisKeyCategory2 := "list"
	//userType := OwnerProductionPlantBusinessPartner
	billOfMaterial, _ := controller.GetInt("billOfMaterial")
	userType := controller.GetString(":userType")

	billOfMaterialHeader := apiInputReader.BillOfMaterial{
		BillOfMaterialHeader: &apiInputReader.BillOfMaterialHeader{
			BillOfMaterial:      billOfMaterial,
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
	controller *BillOfMaterialListController,
) createBillOfMaterialRequestHeaderByOwnerProductionPlantBP(
	requestPram *apiInputReader.Request,
	input apiInputReader.BillOfMaterial,
) *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes {
	responseJsonData := apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes{}
	responseBody := apiModuleRuntimesRequestsBillOfMaterial.BillOfMaterialReads(
		requestPram,
		input,
		&controller.Controller,
		"HeaderByOwnerProductionPlantBP",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createBillOfMaterialRequestHeaderByOwnerProductionPlantBP Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BillOfMaterialListController,
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
	controller *BillOfMaterialListController,
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
	controller *BillOfMaterialListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	billOfMaterialRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*billOfMaterialRes.Message.Header))
	for i, v := range *billOfMaterialRes.Message.Header {
		input[i].Plant = v.OwnerProductionPlant
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
	controller *BillOfMaterialListController,
) request(
	input apiInputReader.BillOfMaterial,
) {
	defer services.Recover(controller.CustomLogger)

	headerRes := controller.createBillOfMaterialRequestHeaderByOwnerProductionPlantBP(
		controller.UserInfo,
		input,
	)

	plantlRes := controller.createPlantRequestGenerals(
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

	controller.fin(
		headerRes,
		plantlRes,
		productDescByBPRes,
		productDocRes,
	)
}

func (
	controller *BillOfMaterialListController,
) fin(
	headerRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
	plantlRes *apiModuleRuntimesResponsesPlant.PlantRes,
	productDescByBPRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {

	plantMapper := services.PlantMapper(
		plantlRes.Message.General,
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

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)
		plantName := fmt.Sprintf("%s", plantMapper[v.OwnerProductionPlant].PlantName)

		data.BillOfMaterialHeader = append(data.BillOfMaterialHeader,
			apiOutputFormatter.BillOfMaterialHeader{
				Product:                  v.Product,
				BillOfMaterial:           v.BillOfMaterial,
				ProductDescription:       &productDescription,
				OwnerProductionPlant:     v.OwnerProductionPlant,
				OwnerProductionPlantName: &plantName,
				ValidityStartDate:        v.ValidityStartDate,
				IsMarkedForDeletion:      v.IsMarkedForDeletion,
				Images: apiOutputFormatter.Images{
					Product: img,
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
