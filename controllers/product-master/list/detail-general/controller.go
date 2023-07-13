package controllersProductMasterDetailGeneral

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequests "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master-doc"
    apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
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

type ProductMasterDetailGeneralController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ProductMasterDetailGeneralController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	productMaster, _ := controller.GetInt("productMaster")
	redisKeyCategory1 := "product-master"
	redisKeyCategory2 := "detail-general"
	redisKeyCategory3 := productMaster
	userType := controller.GetString("userType")

	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	productMasterGeneralDetails := apiInputReader.ProductMaster{
		ProductMasterGeneral: &apiInputReader.ProductMasterGeneral{
			ProductMaster: productMaster,
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
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReads(
		requestPram,
		input,
		&controller.Controller,
		"General",
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
	controller *ProductMasterDetailGeneralController,
) request(
	input apiInputReader.ProductMaster,
) {
	defer services.Recover(controller.CustomLogger)

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
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, 0)
	isMarkedForDeletion := false

	for _, v := range *bRes.Message.General {
		productDescsByBP = append(productDescsByBP, apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BusinessPartner: *requestPram.BusinessPartnerID,
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
	pdRes *apiModuleRuntimesResponsesProductMasterDoc.ProductMasterDocRes,
) {
	descriptionMapper := services.ProductDescByBPMapper(
		pRes.Message.ProductDescByBP,
	)

	data := apiOutputFormatter.ProductMaster{}

	for _, v := range *bGeneralRes.Message.General {
		img := services.CreateProductImage(
			pdRes,
			*requestPram.BusinessPartnerID,
			v.Product,
		)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)

		data.ProductMasterGeneral = append(data.ProductMasterGeneral,
			apiOutputFormatter.ProductMasterGeneral{
				Product:					v.Product,
				ProductName:				&productDescription,
				ProductGroup:				v.ProductGroup,
				BaseUnit:					v.BaseUnit,
				ValidityStartDate:			v.ValidityStartDate,
			},
		)

		data.ProductMasterDetailGeneral = append(data.ProductMasterDetailGeneral,
			apiOutputFormatter.ProductMasterDetailGeneral{
				ProductType:					v.ProductType,
				GrossWeight:					v.GrossWeight,
				NetWeight:						v.NetWeight,
				WeightUnit:						v.WeightUnit,
				InternalCapacityQuantity:		v.InternalCapacityQuantity,
				InternalCapacityQuantityUnit:	v.InternalCapacityQuantityUnit,
				SizeOrDimensionText:			v.SizeOrDimensionText,
				ProductStandardID:				v.ProductStandardID,
				IndustryStandardName:			v.IndustryStandardName,
				CountryOfOrigin:				v.CountryOfOrigin,
				CountryOfOriginLanguage:		v.CountryOfOriginLanguage,
				BarcodeType:					v.BarcodeType,
				ProductAccountAssignmentGroup:	v.ProductAccountAssignmentGroup,
				CreationDate:					v.CreationDate,
				LastChangeDate:					v.LastChangeDate,
				IsMarkedForDeletion:			v.IsMarkedForDeletion,
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
