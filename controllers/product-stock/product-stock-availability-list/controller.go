package controllersProductStockAvailabilityList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesRequestsProductStock "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-stock/product-stock"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProductStock "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-stock"

	//apiModuleRuntimesResponsesProductStock "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-stock"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ProductStockAvailabilityListController struct {
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

func (controller *ProductStockAvailabilityListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	_, _ = controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "product-stock-availability"
	redisKeyCategory2 := "list"
	userType := controller.GetString(":userType") // buyer or seller

	productStockAvailabilityHeader := apiInputReader.ProductStock{}

	if userType == buyer {
		productStockAvailabilityHeader = apiInputReader.ProductStock{
			ProductStockAvailabilityHeader: &apiInputReader.ProductStockAvailabilityHeader{
				Buyer: controller.UserInfo.BusinessPartner,
			},
		}
	}

	if userType == seller {
		productStockAvailabilityHeader = apiInputReader.ProductStock{
			ProductStockAvailabilityHeader: &apiInputReader.ProductStockAvailabilityHeader{
				Seller: controller.UserInfo.BusinessPartner,
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
			controller.request(productStockAvailabilityHeader)
		}()
	} else {
		controller.request(productStockAvailabilityHeader)
	}
}

func (
	controller *ProductStockAvailabilityListController,
) createProductStockRequestProductStockAvailabilityByBuyer(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductStock,
) *apiModuleRuntimesResponsesProductStock.ProductStockRes {
	responseJsonData := apiModuleRuntimesResponsesProductStock.ProductStockRes{}
	responseBody := apiModuleRuntimesRequestsProductStock.ProductStockAvailabilityReads(
		requestPram,
		input,
		&controller.Controller,
		"ProductStockAvailabilitiesByBuyer",
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
	controller *ProductStockAvailabilityListController,
) createProductStockRequestProductStockAvailabilityBySeller(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductStock,
) *apiModuleRuntimesResponsesProductStock.ProductStockRes {
	responseJsonData := apiModuleRuntimesResponsesProductStock.ProductStockRes{}
	responseBody := apiModuleRuntimesRequestsProductStock.ProductStockAvailabilityReads(
		requestPram,
		input,
		&controller.Controller,
		"ProductStockAvailabilitiesBySeller",
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
	controller *ProductStockAvailabilityListController,
) createProductMasterRequestProductDescByBPByBuyer(
	requestPram *apiInputReader.Request,
	pdByBuyerRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, len(*pdByBuyerRes.Message.ProductStockAvailability))
	isMarkedForDeletion := false

	for _, v := range *pdByBuyerRes.Message.ProductStockAvailability {
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
	controller *ProductStockAvailabilityListController,
) createProductMasterRequestProductDescByBPBySeller(
	requestPram *apiInputReader.Request,
	pdByBuyerRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, len(*pdByBuyerRes.Message.ProductStockAvailability))
	isMarkedForDeletion := false

	for _, v := range *pdByBuyerRes.Message.ProductStockAvailability {
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
	controller *ProductStockAvailabilityListController,
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
	controller *ProductStockAvailabilityListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	productStockAvailabilityRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*productStockAvailabilityRes.Message.ProductStockAvailability))

	for _, v := range *productStockAvailabilityRes.Message.ProductStockAvailability {
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
		controller.CustomLogger.Error("BusinessPartnerReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductStockAvailabilityListController,
) request(
	input apiInputReader.ProductStock,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := apiModuleRuntimesResponsesProductStock.ProductStockRes{}
	//productMasterRes := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	if input.ProductStockAvailabilityHeader.Buyer != nil {
		headerRes = *controller.createProductStockRequestProductStockAvailabilityByBuyer(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&headerRes,
		)
	}

	if input.ProductStockAvailabilityHeader.Seller != nil {
		headerRes = *controller.createProductStockRequestProductStockAvailabilityBySeller(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&headerRes,
		)
	}

	controller.fin(
		&headerRes,
		&businessPartnerRes,
	)
}

func (
	controller *ProductStockAvailabilityListController,
) fin(
	headerRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.ProductStock{}

	for _, v := range *headerRes.Message.ProductStockAvailability {

		data.ProductStockAvailabilityHeader = append(data.ProductStockAvailabilityHeader,
			apiOutputFormatter.ProductStockAvailabilityHeader{
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
