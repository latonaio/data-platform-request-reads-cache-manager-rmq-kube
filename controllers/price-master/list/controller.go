package controllersPriceMasterList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequests "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsPriceMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/price-master"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesPriceMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/price-master"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponses "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master-doc"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type PriceMasterListController struct {
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

func (controller *PriceMasterListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "priceMaster"
	redisKeyCategory2 := "list"
	userType := controller.GetString(":userType") // buyer or seller

	priceMasterHeader := apiInputReader.PriceMaster{}

	if userType == buyer {
		priceMasterHeader = apiInputReader.PriceMaster{
			PriceMasterHeader: &apiInputReader.PriceMasterHeader{
				Buyer:               controller.UserInfo.BusinessPartner,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		}
	}

	if userType == seller {
		priceMasterHeader = apiInputReader.PriceMaster{
			PriceMasterHeader: &apiInputReader.PriceMasterHeader{
				Seller:              controller.UserInfo.BusinessPartner,
				IsMarkedForDeletion: &isMarkedForDeletion,
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
			controller.request(priceMasterHeader)
		}()
	} else {
		controller.request(priceMasterHeader)
	}
}

func (
	controller *PriceMasterListController,
) createPriceMasterRequestHeaderByBuyer(
	requestPram *apiInputReader.Request,
	input apiInputReader.PriceMaster,
) *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes {
	responseJsonData := apiModuleRuntimesResponsesPriceMaster.PriceMasterRes{}
	responseBody := apiModuleRuntimesRequestsPriceMaster.PriceMasterReads(
		requestPram,
		input,
		&controller.Controller,
		"PriceMastersByBuyer",
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
	controller *PriceMasterListController,
) createPriceMasterRequestHeaderBySeller(
	requestPram *apiInputReader.Request,
	input apiInputReader.PriceMaster,
) *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes {
	responseJsonData := apiModuleRuntimesResponsesPriceMaster.PriceMasterRes{}
	responseBody := apiModuleRuntimesRequestsPriceMaster.PriceMasterReads(
		requestPram,
		input,
		&controller.Controller,
		"PriceMastersBySeller",
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
	controller *PriceMasterListController,
) createProductMasterRequestProductDescByBPByBuyer(
	requestPram *apiInputReader.Request,
	pdByBuyerRes *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, 0)
	isMarkedForDeletion := false

	for _, v := range *pdByBuyerRes.Message.Header {
		productDescsByBP = append(productDescsByBP, apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BusinessPartner: v.Buyer,
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
	controller *PriceMasterListController,
) createProductMasterRequestProductDescByBPBySeller(
	requestPram *apiInputReader.Request,
	pdByBuyerRes *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, 0)
	isMarkedForDeletion := false

	for _, v := range *pdByBuyerRes.Message.Header {
		productDescsByBP = append(productDescsByBP, apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BusinessPartner: v.Seller,
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
	controller *PriceMasterListController,
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
	controller *PriceMasterListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	priceMasterRes *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, 0)

	for _, v := range *priceMasterRes.Message.Header {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Buyer,
		})
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Seller,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReads(
		requestPram,
		generals,
		&controller.Controller,
		"GeneralsByBusinessPartners",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("BusinessPartnerReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PriceMasterListController,
) request(
	input apiInputReader.PriceMaster,
) {
	defer services.Recover(controller.CustomLogger)

	priceMastersRes := apiModuleRuntimesResponsesPriceMaster.PriceMasterRes{}
	//productMasterRes := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	if input.PriceMasterHeader.Buyer != nil {
		priceMastersRes = *controller.createPriceMasterRequestHeaderByBuyer(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&priceMastersRes,
		)
	}

	if input.PriceMasterHeader.Seller != nil {
		priceMastersRes = *controller.createPriceMasterRequestHeaderBySeller(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&priceMastersRes,
		)
	}

	controller.fin(
		&priceMastersRes,
		&businessPartnerRes,
	)
}

func (
	controller *PriceMasterListController,
) fin(
	priceMasterRes *apiModuleRuntimesResponsesPriceMaster.PriceMasterRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.PriceMaster{}

	for _, v := range *priceMasterRes.Message.Header {

		data.PriceMasterHeader = append(data.PriceMasterHeader,
			apiOutputFormatter.PriceMasterHeader{
				SupplyChainRelationshipID: v.SupplyChainRelationshipID,
				Buyer:                     v.Buyer,
				BuyerName:                 businessPartnerMapper[v.Buyer].BusinessPartnerName,
				Seller:                    v.Seller,
				SellerName:                businessPartnerMapper[v.Seller].BusinessPartnerName,
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
