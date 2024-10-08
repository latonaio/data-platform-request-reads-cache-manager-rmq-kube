package controllersProductMasterDetailGeneral

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

type ProductMasterDetailGeneralController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ProductMasterDetailGeneralController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	product := controller.GetString("product")
	redisKeyCategory1 := "product-master"
	redisKeyCategory2 := "detail-general"
	redisKeyCategory3 := product
	userType := controller.GetString(":userType")

	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	productMasterGeneralDetails := apiInputReader.ProductMaster{
		ProductMasterGeneral: &apiInputReader.ProductMasterGeneral{
			Product: product,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			redisKeyCategory3,
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
			controller.request(productMasterGeneralDetails)
		}()
	} else {
		controller.request(productMasterGeneralDetails)
	}
}

func (
	controller *ProductMasterDetailGeneralController,
) createProductMasterRequestGeneral(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductMaster,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsGenerals(
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
		controller.CustomLogger.Error("createProductMasterRequestGeneral Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductMasterDetailGeneralController,
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
	controller *ProductMasterDetailGeneralController,
) request(
	input apiInputReader.ProductMaster,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	bGeneralRes := controller.createProductMasterRequestGeneral(
		controller.UserInfo,
		input,
	)

	pRes := controller.createProductMasterRequestProductDescByBP(
		controller.UserInfo,
		bGeneralRes,
	)

	pdRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	controller.fin(
		bGeneralRes,
		pRes,
		pdRes,
	)
}

func (
	controller *ProductMasterDetailGeneralController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	bRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, len(*bRes.Message.General))
	isMarkedForDeletion := false

	for _, v := range *bRes.Message.General {
		productDescsByBP = append(productDescsByBP, apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BusinessPartner: *requestPram.BusinessPartner,
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
	controller *ProductMasterDetailGeneralController,
) fin(
	bGeneralRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	pRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	pdRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {
	descriptionMapper := services.ProductDescByBPMapper(
		pRes.Message.ProductDescByBP,
	)

	data := apiOutputFormatter.ProductMaster{}

	for _, v := range *bGeneralRes.Message.General {
		img := services.ReadProductImage(
			pdRes,
			*controller.UserInfo.BusinessPartner,
			v.Product,
		)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)

		data.ProductMasterGeneralWithOthers = append(data.ProductMasterGeneralWithOthers,
			apiOutputFormatter.ProductMasterGeneralWithOthers{
				Product:           v.Product,
				ProductName:       &productDescription,
				ProductGroup:      v.ProductGroup,
				BaseUnit:          v.BaseUnit,
				ValidityStartDate: v.ValidityStartDate,
			},
		)

		data.ProductMasterDetailGeneral = append(data.ProductMasterDetailGeneral,
			apiOutputFormatter.ProductMasterDetailGeneral{
				ProductType:                   v.ProductType,
				GrossWeight:                   v.GrossWeight,
				NetWeight:                     v.NetWeight,
				WeightUnit:                    v.WeightUnit,
				InternalCapacityQuantity:      v.InternalCapacityQuantity,
				InternalCapacityQuantityUnit:  v.InternalCapacityQuantityUnit,
				SizeOrDimensionText:           v.SizeOrDimensionText,
				ProductStandardID:             v.ProductStandardID,
				IndustryStandardName:          v.IndustryStandardName,
				CountryOfOrigin:               v.CountryOfOrigin,
				CountryOfOriginLanguage:       v.CountryOfOriginLanguage,
				BarcodeType:                   v.BarcodeType,
				ProductAccountAssignmentGroup: v.ProductAccountAssignmentGroup,
				CreationDate:                  v.CreationDate,
				LastChangeDate:                v.LastChangeDate,
				IsMarkedForDeletion:           v.IsMarkedForDeletion,
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
