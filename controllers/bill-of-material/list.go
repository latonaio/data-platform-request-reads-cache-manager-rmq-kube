package controllers

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequests "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests"
	apiModuleRuntimesResponses "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"io/ioutil"
	"strings"
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
	userType := controller.GetString("userType")

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
		var responseData apiOutputFormatter.BillOfMaterialList

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
			controller.request(isMarkedForDeletion)
		}()
	} else {
		controller.request(isMarkedForDeletion)
	}
}

func (
	controller *BillOfMaterialListController,
) createBillOfMaterialRequestHeaderByOwnerProductionPlantBP(
	requestPram *apiInputReader.Request,
	isMarkedForDeletion bool,
) *apiModuleRuntimesResponses.BillOfMaterialRes {
	responseJsonData := apiModuleRuntimesResponses.BillOfMaterialRes{}
	responseBody := apiModuleRuntimesRequests.BillOfMaterialReads(
		requestPram,
		isMarkedForDeletion,
		&controller.Controller,
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
	controller *BillOfMaterialListController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	bRes *apiModuleRuntimesResponses.BillOfMaterialRes,
) *apiModuleRuntimesResponses.ProductMasterRes {
	descByBP := make([]apiModuleRuntimesRequests.ProductDescByBP, 0)
	for _, v := range *bRes.Message.Header {
		descByBP = append(descByBP, apiModuleRuntimesRequests.ProductDescByBP{
			Product:         v.Product,
			BusinessPartner: v.OwnerProductionPlantBusinessPartner,
			Language:        *requestPram.Language,
		})
	}

	responseJsonData := apiModuleRuntimesResponses.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequests.ProductMasterReads(
		requestPram,
		descByBP,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("ProductMasterReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BillOfMaterialListController,
) createProductMasterDocRequest(
	requestPram *apiInputReader.Request,
) *apiModuleRuntimesResponses.ProductMasterDocRes {
	responseJsonData := apiModuleRuntimesResponses.ProductMasterDocRes{}
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
	controller *BillOfMaterialListController,
) request(
	isMarkedForDeletion bool,
) {
	defer services.Recover(controller.CustomLogger)

	bRes := controller.createBillOfMaterialRequestHeaderByOwnerProductionPlantBP(
		controller.UserInfo,
		isMarkedForDeletion,
	)

	pRes := controller.createProductMasterRequestProductDescByBP(
		controller.UserInfo,
		bRes,
	)

	pdRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	plRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		bRes,
	)

	controller.fin(
		bRes,
		pRes,
		pdRes,
		plRes,
	)
}

func (
	controller *BillOfMaterialListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	bRes *apiModuleRuntimesResponses.BillOfMaterialRes,
) *apiModuleRuntimesResponses.PlantRes {
	generals := make(apiModuleRuntimesRequests.PlantGenerals, len(*bRes.Message.Header))
	for i, v := range *bRes.Message.Header {
		generals[i].Plant = &v.OwnerProductionPlant
		generals[i].Language = requestPram.Language
	}

	aPIServiceName := "DPFM_API_PLANT_SRV"
	aPIType := "reads"
	responseJsonData := apiModuleRuntimesResponses.PlantRes{}

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
	}

	return &responseJsonData
}

func (
	controller *BillOfMaterialListController,
) fin(
	bRes *apiModuleRuntimesResponses.BillOfMaterialRes,
	pRes *apiModuleRuntimesResponses.ProductMasterRes,
	pdRes *apiModuleRuntimesResponses.ProductMasterDocRes,
	plRes *apiModuleRuntimesResponses.PlantRes,
) {
	descriptionMapper := services.DescriptionMapper(
		pRes.Message.ProductDescByBP,
	)

	plantMapper := services.PlantMapper(
		plRes.Message.Generals,
	)

	data := apiOutputFormatter.BillOfMaterialList{}

	for _, v := range *bRes.Message.Header {
		img := services.CreateProductImage(
			pdRes,
			v.OwnerProductionPlantBusinessPartner,
			v.Product,
		)

		data.BillOfMaterials = append(data.BillOfMaterials,
			apiOutputFormatter.BillOfMaterial{
				Product:                  v.Product,
				BillOfMaterial:           v.BillOfMaterial,
				ProductDescription:       descriptionMapper[v.Product].ProductDescription,
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
