package controllersBillOfMaterialDetailList

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
	"strconv"
)

type BillOfMaterialDetailListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *BillOfMaterialDetailListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	billOfMaterial, _ := controller.GetInt("billOfMaterial")
	redisKeyCategory1 := "bill-of-material"
	redisKeyCategory2 := "detail-list"
	redisKeyCategory3 := billOfMaterial
	userType := controller.GetString(":userType")

	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	billOfMaterialItems := apiInputReader.BillOfMaterial{
		BillOfMaterialHeader: &apiInputReader.BillOfMaterialHeader{
			BillOfMaterial: billOfMaterial,
		},
		BillOfMaterialItems: &apiInputReader.BillOfMaterialItems{
			BillOfMaterial:      billOfMaterial,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
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
			controller.request(billOfMaterialItems)
		}()
	} else {
		controller.request(billOfMaterialItems)
	}
}

func (
	controller *BillOfMaterialDetailListController,
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
	controller *BillOfMaterialDetailListController,
) createBillOfMaterialRequestItems(
	requestPram *apiInputReader.Request,
	input apiInputReader.BillOfMaterial,
) *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes {
	responseJsonData := apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes{}
	responseBody := apiModuleRuntimesRequestsBillOfMaterial.BillOfMaterialReads(
		requestPram,
		input,
		&controller.Controller,
		"Items",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("BillOfMaterialReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BillOfMaterialDetailListController,
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
	controller *BillOfMaterialDetailListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	plantRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*plantRes.Message.Item))
	for i, v := range *plantRes.Message.Item {
		input[i].Plant = v.ProductionPlant
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
	controller *BillOfMaterialDetailListController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	bRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, len(*bRes.Message.Header))
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
		controller.CustomLogger.Error("ProductMasterReadsProductDescsByBP Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BillOfMaterialDetailListController,
) request(
	input apiInputReader.BillOfMaterial,
) {
	defer services.Recover(controller.CustomLogger)

	headerRes := controller.createBillOfMaterialRequestHeader(
		controller.UserInfo,
		input,
	)

	itemRes := controller.createBillOfMaterialRequestItems(
		controller.UserInfo,
		input,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		itemRes,
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
		itemRes,
		plantRes,
		productDescByBPRes,
		productDocRes,
	)
}

func (
	controller *BillOfMaterialDetailListController,
) fin(
	headerRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
	itemRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	productDescByBPRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {

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

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)
		plantName := fmt.Sprintf("%s", plantMapper[v.OwnerProductionPlant].PlantName)

		data.BillOfMaterialHeaderWithItem = append(data.BillOfMaterialHeaderWithItem,
			apiOutputFormatter.BillOfMaterialHeaderWithItem{
				Product:                  v.Product,
				BillOfMaterial:           v.BillOfMaterial,
				ProductDescription:       productDescription,
				OwnerProductionPlant:     v.OwnerProductionPlant,
				OwnerProductionPlantName: plantName,
				ValidityStartDate:        v.ValidityStartDate,
				Images: apiOutputFormatter.Images{
					Product: img,
				},
			},
		)
	}

	for _, v := range *itemRes.Message.Item {
		plantName := fmt.Sprintf("%s", plantMapper[v.StockConfirmationPlant].PlantName)

		data.BillOfMaterialItem = append(data.BillOfMaterialItem,
			apiOutputFormatter.BillOfMaterialItem{
				ComponentProduct:                           v.ComponentProduct,
				BillOfMaterialItem:                         v.BillOfMaterialItem,
				BillOfMaterialItemText:                     *v.BillOfMaterialItemText,
				StockConfirmationPlant:                     &v.StockConfirmationPlant,
				StockConfirmationPlantName:                 &plantName,
				ComponentProductStandardQuantityInBaseUnit: &v.ComponentProductStandardQuantityInBaseUnit,
				ComponentProductBaseUnit:                   &v.ComponentProductBaseUnit,
				ValidityStartDate:                          v.ValidityStartDate,
				IsMarkedForDeletion:                        v.IsMarkedForDeletion,
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
