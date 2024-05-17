package controllersProductStockAvailabilityDetailList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesRequestsProductStock "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-stock/product-stock"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProductStock "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-stock"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ProductStockAvailabilityDetailListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *ProductStockAvailabilityDetailListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "product-stock-availability"
	redisKeyCategory2 := "detail-list"
	redisKeyCategory3 := controller.GetString("availableProductStock")
	supplyChainRelationshipID, _ := controller.GetInt("supplyChainRelationshipID")
	userType := controller.GetString(":userType")
	buyer, _ := controller.GetInt("buyer")
	seller, _ := controller.GetInt("seller")

	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")

	productStockAvailabilityHeaderDetails := apiInputReader.ProductStock{
		ProductStockAvailabilityHeader: &apiInputReader.ProductStockAvailabilityHeader{
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
			controller.request(productStockAvailabilityHeaderDetails)
		}()
	} else {
		controller.request(productStockAvailabilityHeaderDetails)
	}
}

func (
	controller *ProductStockAvailabilityDetailListController,
) createProductStockRequestProductStockAvailabilityByBuyer(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductStock,
) *apiModuleRuntimesResponsesProductStock.ProductStockRes {
	responseJsonData := apiModuleRuntimesResponsesProductStock.ProductStockRes{}
	responseBody := apiModuleRuntimesRequestsProductStock.ProductStockAvailabilityReads(
		requestPram,
		input,
		&controller.Controller,
		"ProductStcokAvailabilitiesByBuyer",
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
	controller *ProductStockAvailabilityDetailListController,
) createProductStockRequestProductStockAvailabilityBySeller(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductStock,
) *apiModuleRuntimesResponsesProductStock.ProductStockRes {
	responseJsonData := apiModuleRuntimesResponsesProductStock.ProductStockRes{}
	responseBody := apiModuleRuntimesRequestsProductStock.ProductStockAvailabilityReads(
		requestPram,
		input,
		&controller.Controller,
		"ProductStcokAvailabilitiesBySeller",
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

//func (
//	controller *ProductStockAvailabilityDetailListController,
//) createProductStockAvailabilityRequestHeader(
//	requestPram *apiInputReader.Request,
//	input apiInputReader.ProductStockAvailability,
//) *apiModuleRuntimesResponsesProductStockAvailability.ProductStockAvailabilityRes {
//	responseJsonData := apiModuleRuntimesResponsesProductStockAvailability.ProductStockAvailabilityRes{}
//	responseBody := apiModuleRuntimesRequestsProductStockAvailability.ProductStockAvailabilityReads(
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
//		controller.CustomLogger.Error("createProductStockAvailabilityRequestHeader Unmarshal error")
//	}
//
//	return &responseJsonData
//}

func (
	controller *ProductStockAvailabilityDetailListController,
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
	controller *ProductStockAvailabilityDetailListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	headerRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*headerRes.Message.ProductStockAvailability))

	for _, v := range *headerRes.Message.ProductStockAvailability {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Buyer,
		})
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Seller,
		})
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.BusinessPartner,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsGeneralsByBusinessPartners(
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
		controller.CustomLogger.Error("BusinessPartnerGeneralReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductStockAvailabilityDetailListController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	headerRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, len(*headerRes.Message.ProductStockAvailability))
	isMarkedForDeletion := false

	for _, v := range *headerRes.Message.ProductStockAvailability {
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
	controller *ProductStockAvailabilityDetailListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	plantRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*plantRes.Message.ProductStockAvailability))
	for i, v := range *plantRes.Message.ProductStockAvailability {
		input[i].Plant = v.Plant
	}

	responseJsonData := apiModuleRuntimesResponsesPlant.PlantRes{}
	responseBody := apiModuleRuntimesRequestsPlant.PlantReadsGeneralsByPlants(
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
		controller.CustomLogger.Error("createPlantRequestGenerals Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductStockAvailabilityDetailListController,
) request(
	input apiInputReader.ProductStock,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	//	productStockAvailabilityHeaderRes := controller.createProductStockAvailabilityRequestHeader(
	//		controller.UserInfo,
	//		input,
	//	)

	headerRes := apiModuleRuntimesResponsesProductStock.ProductStockRes{}

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

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		&headerRes,
	)

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
		plantRes,
		productDescByBPRes,
		productDocRes,
	)
}

func (
	controller *ProductStockAvailabilityDetailListController,
) fin(
	headerRes *apiModuleRuntimesResponsesProductStock.ProductStockRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	plantlRes *apiModuleRuntimesResponsesPlant.PlantRes,
	productDescByBPRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	plantMapper := services.PlantMapper(
		plantlRes.Message.General,
	)

	descriptionMapper := services.ProductDescByBPMapper(
		productDescByBPRes.Message.ProductDescByBP,
	)

	data := apiOutputFormatter.ProductStock{}

	for _, v := range *headerRes.Message.ProductStockAvailability {
		img := services.ReadProductImage(
			productDocRes,
			*controller.UserInfo.BusinessPartner,
			v.Product,
		)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)

		data.ProductStockAvailabilityHeader = append(data.ProductStockAvailabilityHeader,
			apiOutputFormatter.ProductStockAvailabilityHeader{
				SupplyChainRelationshipID: v.SupplyChainRelationshipID,
				Buyer:                     v.Buyer,
				BuyerName:                 businessPartnerMapper[v.Buyer].BusinessPartnerName,
				Seller:                    v.Seller,
				SellerName:                businessPartnerMapper[v.Seller].BusinessPartnerName,
			},
		)

		data.ProductStockAvailabilityDetailHeader = append(data.ProductStockAvailabilityDetailHeader,
			apiOutputFormatter.ProductStockAvailabilityDetailHeader{
				BusinessPartner:              v.BusinessPartner,
				BusinessPartnerName:          businessPartnerMapper[v.BusinessPartner].BusinessPartnerName,
				Plant:                        v.Plant,
				PlantName:                    plantMapper[v.Plant].PlantName,
				Product:                      v.Product,
				ProductDescription:           productDescription,
				AvailableProductStock:        v.AvailableProductStock,
				DeliverToPlant:               v.DeliverToPlant,
				DeliverToPlantName:           plantMapper[v.DeliverToPlant].PlantName,
				DeliverFromPlant:             v.DeliverFromPlant,
				DeliverFromPlantName:         plantMapper[v.DeliverFromPlant].PlantName,
				ProductStockAvailabilityDate: v.ProductStockAvailabilityDate,
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
