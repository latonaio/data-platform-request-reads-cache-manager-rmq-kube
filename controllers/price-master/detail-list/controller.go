package controllersPriceMasterDetailList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsPriceMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/price-master"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesPriceMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/price-master"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
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
	redisKeyCategory1 := "price-master"
	redisKeyCategory2 := "detail-list"
	redisKeyCategory3 := "priceMaster"
	userType := controller.GetString(":userType")
	buyer, _ := controller.GetInt("buyer")
	seller, _ := controller.GetInt("seller")
	supplyChainRelationshipID, _ := controller.GetInt("supplyChainRelationshipID")

	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	priceMasterHeaderDetails := apiInputReader.PriceMaster{
		PriceMasterHeader: &apiInputReader.PriceMasterHeader{
			SupplyChainRelationshipID: supplyChainRelationshipID,
			Buyer:                     &buyer,
			Seller:                    &seller,
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
) createPriceMasterRequestHeaderByBuyer(
	requestPram *apiInputReader.Request,
	input apiInputReader.PriceMaster,
) *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes {
	responseJsonData := apiModuleRuntimesResponsesPriceMaster.PriceMasterRes{}
	responseBody := apiModuleRuntimesRequestsPriceMaster.PriceMasterReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersByBuyer",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("PriceMasterReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PriceMasterDetailListController,
) createPriceMasterRequestHeaderBySeller(
	requestPram *apiInputReader.Request,
	input apiInputReader.PriceMaster,
) *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes {
	responseJsonData := apiModuleRuntimesResponsesPriceMaster.PriceMasterRes{}
	responseBody := apiModuleRuntimesRequestsPriceMaster.PriceMasterReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersBySeller",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("PriceMasterReads Unmarshal error")
	}

	return &responseJsonData
}

//func (
//	controller *PriceMasterDetailListController,
//) createPriceMasterRequestHeader(
//	requestPram *apiInputReader.Request,
//	input apiInputReader.PriceMaster,
//) *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes {
//	responseJsonData := apiModuleRuntimesResponsesPriceMaster.PriceMasterRes{}
//	responseBody := apiModuleRuntimesRequestsPriceMaster.PriceMasterReads(
//		requestPram,
//		input,
//		&controller.Controller,
//		"Header",
//	)
//
//	err := json.Unmarshal(responseBody, &responseJsonData)
//	if err != nil {
//		services.HandleError(
//			&controller.Controller,
//			err,
//			nil,
//		)
//		controller.CustomLogger.Error("createPriceMasterRequestHeader Unmarshal error")
//	}
//
//	return &responseJsonData
//}

func (
	controller *PriceMasterDetailListController,
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
	controller *PriceMasterDetailListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	priceMasterRes *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*priceMasterRes.Message.Header))

	for _, v := range *priceMasterRes.Message.Header {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Buyer,
		})
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Seller,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsGeneralsByBusinessPartners(
		requestPram,
		generals,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("BusinessPartnerGeneralReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PriceMasterDetailListController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	bRes *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, len(*bRes.Message.Header))
	isMarkedForDeletion := false

	for _, v := range *bRes.Message.Header {
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
	controller *PriceMasterDetailListController,
) request(
	input apiInputReader.PriceMaster,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	//	priceMasterHeaderRes := controller.createPriceMasterRequestHeader(
	//		controller.UserInfo,
	//		input,
	//	)

	headerRes := apiModuleRuntimesResponsesPriceMaster.PriceMasterRes{}

	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	if input.PriceMasterHeader.Buyer != nil {
		headerRes = *controller.createPriceMasterRequestHeaderByBuyer(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&headerRes,
		)
	}

	if input.PriceMasterHeader.Seller != nil {
		headerRes = *controller.createPriceMasterRequestHeaderBySeller(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&headerRes,
		)
	}

	productDescByBPRes := controller.createProductMasterRequestProductDescByBP(
		controller.UserInfo,
		&headerRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	controller.fin(
		&headerRes,
		&businessPartnerRes,
		productDescByBPRes,
		productDocRes,
	)
}

func (
	controller *PriceMasterDetailListController,
) fin(
	headerRes *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	productDescByBPRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	descriptionMapper := services.ProductDescByBPMapper(
		productDescByBPRes.Message.ProductDescByBP,
	)

	data := apiOutputFormatter.PriceMaster{}

	for _, v := range *headerRes.Message.Header {
		img := services.ReadProductImage(
			productDocRes,
			*controller.UserInfo.BusinessPartner,
			v.Product,
		)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)

		data.PriceMasterHeader = append(data.PriceMasterHeader,
			apiOutputFormatter.PriceMasterHeader{
				SupplyChainRelationshipID: v.SupplyChainRelationshipID,
				Buyer:                     v.Buyer,
				BuyerName:                 businessPartnerMapper[v.Buyer].BusinessPartnerName,
				Seller:                    v.Seller,
				SellerName:                businessPartnerMapper[v.Seller].BusinessPartnerName,
			},
		)

		data.PriceMasterDetailHeader = append(data.PriceMasterDetailHeader,
			apiOutputFormatter.PriceMasterDetailHeader{
				Product:                   v.Product,
				ProductDescription:        productDescription,
				ConditionType:             v.ConditionType,
				ConditionRateValue:        v.ConditionRateValue,
				ConditionRateValueUnit:    v.ConditionRateValueUnit,
				ConditionScaleQuantity:    v.ConditionScaleQuantity,
				ConditionCurrency:         v.ConditionCurrency,
				ConditionRecord:           v.ConditionRecord,
				ConditionSequentialNumber: v.ConditionSequentialNumber,
				IsMarkedForDeletion:       v.IsMarkedForDeletion,
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
