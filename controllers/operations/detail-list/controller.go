package controllersOperationsDetailList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsOperations "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/operations"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesOperations "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/operations"
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

type OperationsDetailListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *OperationsDetailListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	operations, _ := controller.GetInt("operations")
	redisKeyCategory1 := "operations"
	redisKeyCategory2 := "detail-list"
	redisKeyCategory3 := operations
	userType := controller.GetString(":userType")

	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	operationsItems := apiInputReader.Operations{
		OperationsHeader: &apiInputReader.OperationsHeader{
			Operations: operations,
		},
		OperationsItems: &apiInputReader.OperationsItems{
			Operations:          operations,
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
		var responseData apiOutputFormatter.Operations

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
			controller.request(operationsItems)
		}()
	} else {
		controller.request(operationsItems)
	}
}

func (
	controller *OperationsDetailListController,
) createOperationsRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.Operations,
) *apiModuleRuntimesResponsesOperations.OperationsRes {
	responseJsonData := apiModuleRuntimesResponsesOperations.OperationsRes{}
	responseBody := apiModuleRuntimesRequestsOperations.OperationsReads(
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
		controller.CustomLogger.Error("createOperationsRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OperationsDetailListController,
) createOperationsRequestItems(
	requestPram *apiInputReader.Request,
	input apiInputReader.Operations,
) *apiModuleRuntimesResponsesOperations.OperationsRes {
	responseJsonData := apiModuleRuntimesResponsesOperations.OperationsRes{}
	responseBody := apiModuleRuntimesRequestsOperations.OperationsReads(
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
		controller.CustomLogger.Error("OperationsReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OperationsDetailListController,
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
	controller *OperationsDetailListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	plantRes *apiModuleRuntimesResponsesOperations.OperationsRes,
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
	controller *OperationsDetailListController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	pdBPRes *apiModuleRuntimesResponsesOperations.OperationsRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, len(*pdBPRes.Message.Header))
	isMarkedForDeletion := false

	for _, v := range *pdBPRes.Message.Header {
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
	controller *OperationsDetailListController,
) request(
	input apiInputReader.Operations,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := controller.createOperationsRequestHeader(
		controller.UserInfo,
		input,
	)

	itemRes := controller.createOperationsRequestItems(
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

	controller.fin(
		headerRes,
		itemRes,
		plantRes,
		productDescByBPRes,
		productDocRes,
	)
}

func (
	controller *OperationsDetailListController,
) fin(
	headerRes *apiModuleRuntimesResponsesOperations.OperationsRes,
	itemRes *apiModuleRuntimesResponsesOperations.OperationsRes,
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

	data := apiOutputFormatter.Operations{}

	for _, v := range *headerRes.Message.Header {
		img := services.ReadProductImage(
			productDocRes,
			v.OwnerProductionPlantBusinessPartner,
			v.Product,
		)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)
		plantName := fmt.Sprintf("%s", plantMapper[v.OwnerProductionPlant].PlantName)

		data.OperationsHeader = append(data.OperationsHeader,
			apiOutputFormatter.OperationsHeader{
				Product:                  &v.Product,
				Operations:               v.Operations,
				ProductDescription:       &productDescription,
				OwnerProductionPlant:     &v.OwnerProductionPlant,
				OwnerProductionPlantName: &plantName,
				ValidityStartDate:        v.ValidityStartDate,
				IsMarkedForDeletion:      v.IsMarkedForDeletion,
				Images: apiOutputFormatter.Images{
					Product: img,
				},
			},
		)
	}

	for _, v := range *headerRes.Message.Item {
		data.OperationsItem = append(data.OperationsItem,
			apiOutputFormatter.OperationsItem{
				OperationsItem:          v.OperationsItem,
				OperationsText:          v.OperationsText,
				ProductionPlant:         v.ProductionPlant,
				ProductionPlantName:     plantMapper[v.ProductionPlant].PlantName,
				StandardLotSizeQuantity: v.StandardLotSizeQuantity,
				OperationsUnit:          v.OperationsUnit,
				ValidityStartDate:       v.ValidityStartDate,
				IsMarkedForDeletion:     v.IsMarkedForDeletion,
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
