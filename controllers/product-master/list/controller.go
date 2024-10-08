package controllersProductMasterList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ProductMasterListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

//const (
//	buyer	= "buyer"
//	seller	= "seller"
//)

func (controller *ProductMasterListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "product-master"
	redisKeyCategory2 := "list"
	userType := controller.GetString(":userType")

	productMasterGeneral := apiInputReader.ProductMaster{
		ProductMasterGeneral: &apiInputReader.ProductMasterGeneral{},
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
		var responseData apiOutputFormatter.ProductMaster

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
			controller.request(productMasterGeneral)
		}()
	} else {
		controller.request(productMasterGeneral)
	}
}

func (
	controller *ProductMasterListController,
) createProductMasterRequestGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductMaster,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsGenerals(
		requestPram,
		apiInputReader.ProductMaster{},
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductMasterRequestGenerals Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductMasterListController,
) createProductMasterRequestProductDescription(
	requestPram *apiInputReader.Request,
	pRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescription := make([]apiModuleRuntimesRequestsProductMaster.General, len(*pRes.Message.General))
	isMarkedForDeletion := false

	for _, v := range *pRes.Message.General {
		productDescription = append(productDescription, apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			ProductDescription: []apiModuleRuntimesRequestsProductMaster.ProductDescription{
				{
					Language:            *requestPram.Language,
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		})
	}

	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsProductDescriptions(
		requestPram,
		productDescription,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductMasterRequestProductDescription Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductMasterListController,
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
	controller *ProductMasterListController,
) request(
	input apiInputReader.ProductMaster,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	gsRes := controller.createProductMasterRequestGenerals(
		controller.UserInfo,
		input,
	)

	pdRes := controller.createProductMasterRequestProductDescription(
		controller.UserInfo,
		gsRes,
	)

	//pMDocRes := controller.createProductMasterDocRequest(
	//	controller.UserInfo,
	//)

	controller.fin(
		gsRes,
		pdRes,
		//pMDocRes,
	)
}

func (
	controller *ProductMasterListController,
) fin(
	gsRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	pdRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	// pMDocRes *apiModuleRuntimesResponsesProductMasterDoc.ProductMasterDocRes,
) {
	productDescriptionMapper := services.ProductDescriptionMapper(
		pdRes.Message.ProductDescription,
	)

	data := apiOutputFormatter.ProductMaster{}

	for _, v := range *gsRes.Message.General {
		productDescription := fmt.Sprintf("%s", productDescriptionMapper[v.Product].ProductDescription)

		data.ProductMasterGeneral = append(data.ProductMasterGeneral,
			apiOutputFormatter.ProductMasterGeneral{
				Product:             v.Product,
				ProductName:         &productDescription,
				ProductGroup:        v.ProductGroup,
				BaseUnit:            v.BaseUnit,
				ValidityStartDate:   v.ValidityStartDate,
				IsMarkedForDeletion: v.IsMarkedForDeletion,
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
