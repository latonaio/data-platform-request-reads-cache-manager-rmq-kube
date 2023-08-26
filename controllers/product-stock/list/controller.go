package controllersProductStockList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesRequestsProductStock "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-stock"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProductStock "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-stock"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ProductStockListController struct {
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

func (controller *ProductStockListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	_, _ = controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "product-stock"
	redisKeyCategory2 := "list"
	userType := controller.GetString(":userType") // buyer or seller

	productStockHeader := apiInputReader.ProductStock{}

	if userType == buyer {
		productStockHeader = apiInputReader.ProductStock{
			ProductStockHeader: &apiInputReader.ProductStockHeader{
				Buyer: controller.UserInfo.BusinessPartner,
			},
		}
	}

	if userType == seller {
		productStockHeader = apiInputReader.ProductStock{
			ProductStockHeader: &apiInputReader.ProductStockHeader{
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
			controller.request(productStockHeader)
		}()
	} else {
		controller.request(productStockHeader)
	}
}

func (
	controller *ProductStockListController,
) createProductStockRequestHeaderByBuyer(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductStock,
) *apiModuleRuntimesResponsesProductStock.ProductStockRes {
	responseJsonData := apiModuleRuntimesResponsesProductStock.ProductStockRes{}
	responseBody := apiModuleRuntimesRequestsProductStock.ProductStockReads(
		requestPram,
		input,
		&controller.Controller,
		"ProductStocksByBuyer",
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
	controller *ProductStockListController,
) createProductStockRequestHeaderBySeller(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductStock,
) *apiModuleRuntimesResponsesProductStock.ProductStockRes {
	responseJsonData := apiModuleRuntimesResponsesProductStock.ProductStockRes{}
	responseBody := apiModuleRuntimesRequestsProductStock.ProductStockReads(
		requestPram,
		input,
		&controller.Controller,
		"ProductStocksBySeller",
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
	controller *ProductStockListController,
) createProductMasterRequestProductDescByBPByBuyer(
	requestPram *apiInputReader.Request,
	pdByBuyerRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, len(*pdByBuyerRes.Message.Header))
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
	controller *ProductStockListController,
) createProductMasterRequestProductDescByBPBySeller(
	requestPram *apiInputReader.Request,
	pdByBuyerRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, len(*pdByBuyerRes.Message.Header))
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
	controller *ProductStockListController,
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
	controller *ProductStockListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	productStockRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*productStockRes.Message.Header))

	for _, v := range *productStockRes.Message.Header {
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
	controller *ProductStockListController,
) request(
	input apiInputReader.ProductStock,
) {
	defer services.Recover(controller.CustomLogger)

	headerRes := apiModuleRuntimesResponsesProductStock.ProductStockRes{}
	//productMasterRes := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	if input.ProductStockHeader.Buyer != nil {
		headerRes = *controller.createProductStockRequestHeaderByBuyer(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&headerRes,
		)
	}

	if input.ProductStockHeader.Seller != nil {
		headerRes = *controller.createProductStockRequestHeaderBySeller(
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
	controller *ProductStockListController,
) fin(
	headerRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.ProductStock{}

	for _, v := range *headerRes.Message.Header {

		data.ProductStockHeader = append(data.ProductStockHeader,
			apiOutputFormatter.ProductStockHeader{
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
