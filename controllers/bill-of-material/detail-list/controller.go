package controllersBillOfMaterialDetailList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequests "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests"
	apiModuleRuntimesRequestsBillOfMaterial "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/bill-of-material"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master"
	"data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/bill-of-material"
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
	userType := controller.GetString("userType")

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
	controller *BillOfMaterialDetailListController,
) request(
	input apiInputReader.BillOfMaterial,
) {
	defer services.Recover(controller.CustomLogger)

	bHeaderRes := controller.createBillOfMaterialRequestHeader(
		controller.UserInfo,
		input,
	)

	bRes := controller.createBillOfMaterialRequestItems(
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
	controller *BillOfMaterialDetailListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	bRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
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
	controller *BillOfMaterialDetailListController,
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
		controller.CustomLogger.Error("ProductMasterReadsProductDescsByBP Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BillOfMaterialDetailListController,
) fin(
	bHeaderRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
	bRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
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

	data := apiOutputFormatter.BillOfMaterial{}

	for _, v := range *bHeaderRes.Message.Header {
		img := services.CreateProductImage(
			pdRes,
			v.OwnerProductionPlantBusinessPartner,
			v.Product,
		)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)

		data.BillOfMaterialHeader = append(data.BillOfMaterialHeader,
			apiOutputFormatter.BillOfMaterialHeader{
				Product:                  v.Product,
				BillOfMaterial:           v.BillOfMaterial,
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
		data.BillOfMaterialItem = append(data.BillOfMaterialItem,
			apiOutputFormatter.BillOfMaterialItem{
				ComponentProduct:                           v.ComponentProduct,
				BillOfMaterialItem:                         v.BillOfMaterialItem,
				BillOfMaterialItemText:                     *v.BillOfMaterialItemText,
				StockConfirmationPlant:                     &v.StockConfirmationPlant,
				StockConfirmationPlantName:                 plantMapper[v.StockConfirmationPlant].PlantName,
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
