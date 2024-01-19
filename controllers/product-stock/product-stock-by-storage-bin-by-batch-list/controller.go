package controllersProductStockByStorageBinByBatchList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBatchMasterRecord "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/batch-master-record"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesRequestsProductStock "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-stock/product-stock"
	apiModuleRuntimesResponsesBatchMasterRecord "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/batch-master-record"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProductStock "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-stock"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ProductStockByStorageBinByBatchListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const (
	buyer  = "buyer"
	seller = "seller"
)

func (controller *ProductStockByStorageBinByBatchListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	_, _ = controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "product-stock"
	redisKeyCategory2 := "by-storage-bin-by-batch-list"
	userType := controller.GetString(":userType") // buyer or seller
	product := controller.GetString("product")
	plant := controller.GetString("plant")

	productStockByStorageBinByBatchHeader := apiInputReader.ProductStock{}

	if userType == buyer {
		productStockByStorageBinByBatchHeader = apiInputReader.ProductStock{
			ProductStockHeader: &apiInputReader.ProductStockHeader{
				Product:         product,
				BusinessPartner: *controller.UserInfo.BusinessPartner,
				Plant:           plant,
			},
			ProductStockByStorageBinByBatchHeader: &apiInputReader.ProductStockByStorageBinByBatchHeader{
				Product:         product,
				BusinessPartner: *controller.UserInfo.BusinessPartner,
				Plant:           plant,
				Buyer:           controller.UserInfo.BusinessPartner,
			},
		}
	}

	if userType == seller {
		productStockByStorageBinByBatchHeader = apiInputReader.ProductStock{
			ProductStockHeader: &apiInputReader.ProductStockHeader{
				Product:         product,
				BusinessPartner: *controller.UserInfo.BusinessPartner,
				Plant:           plant,
			},
			ProductStockByStorageBinByBatchHeader: &apiInputReader.ProductStockByStorageBinByBatchHeader{
				Product:         product,
				BusinessPartner: *controller.UserInfo.BusinessPartner,
				Plant:           plant,
				Seller:          controller.UserInfo.BusinessPartner,
			},
		}
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
		var responseData apiOutputFormatter.ProductStock

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
			controller.request(productStockByStorageBinByBatchHeader)
		}()
	} else {
		controller.request(productStockByStorageBinByBatchHeader)
	}
}

func (
	controller *ProductStockByStorageBinByBatchListController,
) createProductStockRequestProductStockByStorageBinByBatch(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductStock,
) *apiModuleRuntimesResponsesProductStock.ProductStockRes {
	responseJsonData := apiModuleRuntimesResponsesProductStock.ProductStockRes{}
	responseBody := apiModuleRuntimesRequestsProductStock.ProductStockReads(
		requestPram,
		input,
		&controller.Controller,
		"ProductStocksByStorageBinByBatch",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("ProductStockReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductStockByStorageBinByBatchListController,
) createProductStockRequestProductStock(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductStock,
) *apiModuleRuntimesResponsesProductStock.ProductStockRes {
	responseJsonData := apiModuleRuntimesResponsesProductStock.ProductStockRes{}
	responseBody := apiModuleRuntimesRequestsProductStock.ProductStockReads(
		requestPram,
		input,
		&controller.Controller,
		"ProductStock",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductStockRequestProductStock Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductStockByStorageBinByBatchListController,
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
	controller *ProductStockByStorageBinByBatchListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	productStockRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*productStockRes.Message.ProductStock))
	for i, v := range *productStockRes.Message.ProductStock {
		input[i].Plant = v.Plant
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
	controller *ProductStockByStorageBinByBatchListController,
) createBatchMasterRecordRequestHeader(
	requestPram *apiInputReader.Request,
	productStockByStorageBinByBatchRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
) *apiModuleRuntimesResponsesBatchMasterRecord.BatchMasterRecordRes {
	batches := make([]apiModuleRuntimesRequestsBatchMasterRecord.Batch,
		len(*productStockByStorageBinByBatchRes.Message.ProductStockByStorageBinByBatch),
	)
	responseJsonData := apiModuleRuntimesResponsesBatchMasterRecord.BatchMasterRecordRes{}

	for _, v := range *productStockByStorageBinByBatchRes.Message.ProductStockByStorageBinByBatch {
		batches = append(batches, apiModuleRuntimesRequestsBatchMasterRecord.Batch{
			Product:         v.Product,
			BusinessPartner: *requestPram.BusinessPartner,
			Plant:           v.Plant,
			Batch:           v.Batch,
		})
	}

	responseBody := apiModuleRuntimesRequestsBatchMasterRecord.BatchMasterRecordReads(
		requestPram,
		batches,
		&controller.Controller,
		"Batches",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createBatchMasterRecordRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductStockByStorageBinByBatchListController,
) request(
	input apiInputReader.ProductStock,
) {
	defer services.Recover(controller.CustomLogger)

	productStockRes := controller.createProductStockRequestProductStock(
		controller.UserInfo,
		input,
	)

	productStockByStorageBinByBatchRes := controller.createProductStockRequestProductStockByStorageBinByBatch(
		controller.UserInfo,
		input,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		productStockRes,
	)

	batchMasterRecordRes := controller.createBatchMasterRecordRequestHeader(
		controller.UserInfo,
		productStockByStorageBinByBatchRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	controller.fin(
		productStockRes,
		productStockByStorageBinByBatchRes,
		batchMasterRecordRes,
		plantRes,
		productDocRes,
	)
}

func (
	controller *ProductStockByStorageBinByBatchListController,
) fin(
	productStockRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
	productStockByStorageBinByBatchRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
	batchMasterRecordRes *apiModuleRuntimesResponsesBatchMasterRecord.BatchMasterRecordRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {
	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	batchMapper := services.BatchMapper(
		batchMasterRecordRes.Message.Batch,
	)

	data := apiOutputFormatter.ProductStock{}

	for _, v := range *productStockRes.Message.ProductStock {
		img := services.ReadProductImage(
			productDocRes,
			v.BusinessPartner,
			v.Product,
		)

		plantName := fmt.Sprintf("%s", plantMapper[v.Plant].PlantName)

		data.ProductStockHeader = append(data.ProductStockHeader,
			apiOutputFormatter.ProductStockHeader{
				ProductStock: v.ProductStock,
				PlantName:    plantName,
				Images: apiOutputFormatter.Images{
					Product: img,
				},
			},
		)
	}

	for _, v := range *productStockByStorageBinByBatchRes.Message.ProductStockByStorageBinByBatch {

		validityStartDate := fmt.Sprintf("%s", batchMapper[v.Batch].ValidityStartDate)
		validityStartTime := fmt.Sprintf("%s", batchMapper[v.Batch].ValidityStartTime)
		validityEndDate := fmt.Sprintf("%s", batchMapper[v.Batch].ValidityEndDate)
		validityEndTime := fmt.Sprintf("%s", batchMapper[v.Batch].ValidityEndTime)

		data.ProductStockByStorageBinByBatchHeader = append(data.ProductStockByStorageBinByBatchHeader,
			apiOutputFormatter.ProductStockByStorageBinByBatchHeader{
				Plant:             v.Plant,
				StorageLocation:   v.StorageLocation,
				StorageBin:        v.StorageBin,
				Batch:             v.Batch,
				ProductStock:      v.ProductStock,
				ValidityStartDate: validityStartDate,
				ValidityStartTime: validityStartTime,
				ValidityEndDate:   validityEndDate,
				ValidityEndTime:   validityEndTime,
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
