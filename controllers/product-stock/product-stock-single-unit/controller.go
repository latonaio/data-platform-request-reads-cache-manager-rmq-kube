package controllersProductStockSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	//apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesRequestsProductStock "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-stock/product-stock"
	apiModuleRuntimesRequestsProductStockDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-stock/product-stock-doc"
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

type ProductStockSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const ()

func (controller *ProductStockSingleUnitController) Get() {
	//isReleased, _ := controller.GetBool("isReleased")
	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "product-stock"
	redisKeyCategory2 := "product-stock-single-unit"
	product := controller.GetString("product")
	plant := controller.GetString("plant")

	productStockSingleUnit := apiInputReader.ProductStock{
		ProductStockHeader: &apiInputReader.ProductStockHeader{
			Product:         product,
			BusinessPartner: *controller.UserInfo.BusinessPartner,
			Plant:           plant,
		},
		ProductStockDocProductStockDoc: &apiInputReader.ProductStockDocProductStockDoc{
			Product:                  &product,
			BusinessPartner:          controller.UserInfo.BusinessPartner,
			DocType:                  "QRCODE",
			DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
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
			controller.request(productStockSingleUnit)
		}()
	} else {
		controller.request(productStockSingleUnit)
	}
}

func (
	controller *ProductStockSingleUnitController,
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
	controller *ProductStockSingleUnitController,
) createProductStockDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductStock,
) *apiModuleRuntimesResponsesProductStock.ProductStockDocRes {
	responseJsonData := apiModuleRuntimesResponsesProductStock.ProductStockDocRes{}
	responseBody := apiModuleRuntimesRequestsProductStockDoc.ProductStockDocReads(
		requestPram,
		input,
		&controller.Controller,
		"ProductStockDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductStockDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductStockSingleUnitController,
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
	controller *ProductStockSingleUnitController,
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
	controller *ProductStockSingleUnitController,
) request(
	input apiInputReader.ProductStock,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	productStockRes := controller.createProductStockRequestProductStock(
		controller.UserInfo,
		input,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		productStockRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	productStockDocRes := controller.createProductStockDocRequest(
		controller.UserInfo,
		input,
	)

	controller.fin(
		productStockRes,
		plantRes,
		productDocRes,
		productStockDocRes,
	)
}

func (
	controller *ProductStockSingleUnitController,
) fin(
	productStockRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	productStockDocRes *apiModuleRuntimesResponsesProductStock.ProductStockDocRes,
) {

	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	data := apiOutputFormatter.ProductStock{}

	for _, v := range *productStockRes.Message.ProductStock {
		img := services.ReadProductImage(
			productDocRes,
			v.BusinessPartner,
			v.Product,
		)

		qrcode := services.CreateQRCodeProductStockDocImage(
			productStockDocRes,
		)

		plantName := fmt.Sprintf("%s", plantMapper[v.Plant].PlantName)

		data.ProductStockSingleUnit = append(data.ProductStockSingleUnit,
			apiOutputFormatter.ProductStockSingleUnit{
				ProductStock: v.ProductStock,
				PlantName:    plantName,
				Images: apiOutputFormatter.Images{
					Product: img,
					QRCode:  qrcode,
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
