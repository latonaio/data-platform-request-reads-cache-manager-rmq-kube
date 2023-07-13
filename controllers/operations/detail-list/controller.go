package controllersOperationsDetailList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequests "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests"
	apiModuleRuntimesRequestsOperations "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/operations"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master"
	apiModuleRuntimesResponsesOperations "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/operations"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master-doc"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"io/ioutil"
	"strconv"
	"strings"
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
	userType := controller.GetString("userType")

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
) *apiModuleRuntimesResponsesProductMasterDoc.ProductMasterDocRes {
	responseJsonData := apiModuleRuntimesResponsesProductMasterDoc.ProductMasterDocRes{}
	responseBody := apiModuleRuntimesRequests.ProductMasterDocReads(
		requestPram,
		&controller.Controller,
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
) request(
	input apiInputReader.Operations,
) {
	defer services.Recover(controller.CustomLogger)

	bHeaderRes := controller.createOperationsRequestHeader(
		controller.UserInfo,
		input,
	)

	bRes := controller.createOperationsRequestItems(
		controller.UserInfo,
		input,
	)

	plRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		bRes,
	)

	pRes := controller.createProductMasterRequestProductDescByBP(
		controller.UserInfo,
		bHeaderRes,
	)

	pdRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	controller.fin(
		bHeaderRes,
		bRes,
		plRes,
		pRes,
		pdRes,
	)
}

func (
	controller *OperationsDetailListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	bRes *apiModuleRuntimesResponsesOperations.OperationsRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	generals := make(apiModuleRuntimesRequests.PlantGenerals, len(*bRes.Message.Item))
	for i, v := range *bRes.Message.Item {
		generals[i].Plant = &v.ProductionPlant
		generals[i].Language = requestPram.Language
	}

	aPIServiceName := "DPFM_API_PLANT_SRV"
	aPIType := "reads"
	responseJsonData := apiModuleRuntimesResponsesPlant.PlantRes{}

	request := apiModuleRuntimesRequests.
		CreatePlantRequestGenerals(
			requestPram,
			generals,
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
	controller *OperationsDetailListController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	bRes *apiModuleRuntimesResponsesOperations.OperationsRes,
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
		controller.CustomLogger.Error("ProductMasterReadsProductDescsByBP Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OperationsDetailListController,
) fin(
	bHeaderRes *apiModuleRuntimesResponsesOperations.OperationsRes,
	bRes *apiModuleRuntimesResponsesOperations.OperationsRes,
	plRes *apiModuleRuntimesResponsesPlant.PlantRes,
	pRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	pdRes *apiModuleRuntimesResponsesProductMasterDoc.ProductMasterDocRes,
) {
	descriptionMapper := services.ProductDescByBPMapper(
		pRes.Message.ProductDescByBP,
	)

	plantMapper := services.PlantMapper(
		plRes.Message.Generals,
	)

	data := apiOutputFormatter.Operations{}

	for _, v := range *bHeaderRes.Message.Header {
		img := services.CreateProductImage(
			pdRes,
			v.OwnerProductionPlantBusinessPartner,
			v.Product,
		)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)

		data.OperationsHeader = append(data.OperationsHeader,
			apiOutputFormatter.OperationsHeader{
				Product:                  v.Product,
				Operations:               v.Operations,
				ProductDescription:       &productDescription,
				OwnerProductionPlant:     v.OwnerProductionPlant,
				OwnerProductionPlantName: plantMapper[v.OwnerProductionPlant].PlantName,
				ValidityStartDate:        v.ValidityStartDate,
				IsMarkedForDeletion:      v.IsMarkedForDeletion,
				Images: apiOutputFormatter.Images{
					Product: img,
				},
			},
		)
	}

	for _, v := range *bRes.Message.Item {
		data.OperationsItem = append(data.OperationsItem,
			apiOutputFormatter.OperationsItem{
				OperationsItem:          v.OperationsItem,
				OperationsText:          *v.OperationsText,
				ProductionPlant:         &v.ProductionPlant,
				ProductionPlantName:     plantMapper[v.ProductionPlantName].PlantName,
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
