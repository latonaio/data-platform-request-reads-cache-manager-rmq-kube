package controllersPriceMasterDetailList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequests "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsPriceMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/price-master"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesPriceMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/price-master"
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

type PriceMasterDetailListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *PriceMasterDetailListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	priceMaster, _ := controller.GetInt("priceMaster")
	redisKeyCategory1 := "price-master"
	redisKeyCategory2 := "detail-list"
	redisKeyCategory3 := priceMaster
	userType := controller.GetString("userType")

	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	priceMasterHeaderDetails := apiInputReader.PriceMaster{
		PriceMasterHeader: &apiInputReader.PriceMasterHeader{
			PriceMaster: priceMaster,
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
		var responseData apiOutputFormatter.PriceMaster

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
			controller.request(priceMasterHeaderDetails)
		}()
	} else {
		controller.request(priceMasterHeaderDetails)
	}
}

func (
	controller *PriceMasterDetailListController,
) createPriceMasterRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.PriceMaster,
) *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes {
	responseJsonData := apiModuleRuntimesResponsesPriceMaster.PriceMasterRes{}
	responseBody := apiModuleRuntimesRequestsPriceMaster.PriceMasterReads(
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
		controller.CustomLogger.Error("createPriceMasterRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PriceMasterDetailListController,
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
	controller *PriceMasterDetailListController,
) request(
	input apiInputReader.PriceMaster,
) {
	defer services.Recover(controller.CustomLogger)

	bHeaderRes := controller.createPriceMasterRequestHeader(
		controller.UserInfo,
		input,
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
		pRes,
		pdRes,
	)
}

func (
	controller *PriceMasterDetailListController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	bRes *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, 0)
	isMarkedForDeletion := false

	for _, v := range *bRes.Message.Header {
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
	controller *PriceMasterDetailListController,
) fin(
	bHeaderRes *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	pRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	pdRes *apiModuleRuntimesResponsesProductMasterDoc.ProductMasterDocRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	descriptionMapper := services.ProductDescByBPMapper(
		pRes.Message.ProductDescByBP,
	)

	data := apiOutputFormatter.PriceMaster{}

	for _, v := range *bHeaderRes.Message.Header {
		img := services.CreateProductImage(
			pdRes,
			*requestPram.BusinessPartnerID,
			v.Product,
		)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)

		data.PriceMasterHeader = append(data.PriceMasterHeader,
			apiOutputFormatter.PriceMasterHeader{
				SupplyChainRelationshipID:	v.Product,
				Buyer:						v.Buyer,
				BuyerName:					businessPartnerMapper[v.Buyer].BusinessPartnerName,
				Seller:						v.Seller,
				SellerName:					businessPartnerMapper[v.Seller].BusinessPartnerName,
			},
		)
		
		data.PriceMasterDetailHeader = append(data.PriceMasterDetailHeader,
			apiOutputFormatter.PriceMasterDetailHeader{
				Product:                  v.Product,
				ProductDescription:       &productDescription,
				ConditionType:            v.ConditionType,
				ConditionRateValue:       v.ConditionRateValue,
				ConditionRateValueUnit:   v.ConditionRateValueUnit,
				ConditionScaleQuantity:   v.ConditionScaleQuantity,
                ConditionCurrency         v.ConditionCurrency,
                ConditionRecord           v.ConditionRecord,
                ConditionSequentialNumber v.ConditionSequentialNumber,
                IsMarkedForDeletion       v.IsMarkedForDeletion,
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
