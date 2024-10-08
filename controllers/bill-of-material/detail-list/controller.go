package controllersBillOfMaterialDetailList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBillOfMaterial "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/bill-of-material/bill-of-material"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesBillOfMaterial "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/bill-of-material"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type BillOfMaterialDetailListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *BillOfMaterialDetailListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	billOfMaterial, _ := controller.GetInt("billOfMaterial")
	redisKeyCategory1 := "bill-of-material"
	redisKeyCategory2 := "detail-list"
	redisKeyCategory3 := billOfMaterial

	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	isMarkedForDeletion := false

	billOfMaterialItems := apiInputReader.BillOfMaterial{
		BillOfMaterialHeader: &apiInputReader.BillOfMaterialHeader{
			BillOfMaterial: billOfMaterial,
		},
		BillOfMaterialItems: &apiInputReader.BillOfMaterialItems{
			BillOfMaterial:      billOfMaterial,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(redisKeyCategory3),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.BillOfMaterial

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
			controller.request(billOfMaterialItems)
		}()
	} else {
		controller.request(billOfMaterialItems)
	}
}

func (
	controller *BillOfMaterialDetailListController,
) createBillOfMaterialRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.BillOfMaterial,
) *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes {
	responseJsonData := apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes{}
	responseBody := apiModuleRuntimesRequestsBillOfMaterial.BillOfMaterialReads(
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
		controller.CustomLogger.Error("createBillOfMaterialRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BillOfMaterialDetailListController,
) createBillOfMaterialRequestItems(
	requestPram *apiInputReader.Request,
	input apiInputReader.BillOfMaterial,
) *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes {
	responseJsonData := apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes{}
	responseBody := apiModuleRuntimesRequestsBillOfMaterial.BillOfMaterialReads(
		requestPram,
		input,
		&controller.Controller,
		"Items",
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
	controller *BillOfMaterialDetailListController,
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
	controller *BillOfMaterialDetailListController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	billOfMaterialRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*billOfMaterialRes.Message.Header))
	for i, v := range *billOfMaterialRes.Message.Header {
		input[i].Plant = v.OwnerProductionPlant
		input[i].BusinessPartner = v.OwnerProductionPlantBusinessPartner
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
	controller *BillOfMaterialDetailListController,
) createProductMasterRequestProductDescByBP(
	requestPram *apiInputReader.Request,
	bRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	productDescsByBP := make([]apiModuleRuntimesRequestsProductMaster.General, len(*bRes.Message.Header))
	isMarkedForDeletion := false

	for _, v := range *bRes.Message.Header {
		productDescsByBP = append(productDescsByBP, apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BusinessPartner: v.OwnerProductionPlantBusinessPartner,
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
	controller *BillOfMaterialDetailListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	billOfMaterialHeaderRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
	billOfMaterialItemRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input []apiModuleRuntimesRequestsBusinessPartner.General

	for _, v := range *billOfMaterialHeaderRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Buyer,
		})
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Seller,
		})
	}

	for _, v := range *billOfMaterialItemRes.Message.Item {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.ComponentProductBuyer,
		})
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.ComponentProductSeller,
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
	controller *BillOfMaterialDetailListController,
) request(
	input apiInputReader.BillOfMaterial,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := controller.createBillOfMaterialRequestHeader(
		controller.UserInfo,
		input,
	)

	itemRes := controller.createBillOfMaterialRequestItems(
		controller.UserInfo,
		input,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		headerRes,
	)

	productDescByBPRes := controller.createProductMasterRequestProductDescByBP(
		controller.UserInfo,
		headerRes,
	)

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		headerRes,
		itemRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	controller.fin(
		headerRes,
		itemRes,
		plantRes,
		&businessPartnerRes,
		productDescByBPRes,
		productDocRes,
	)
}

func (
	controller *BillOfMaterialDetailListController,
) fin(
	headerRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
	itemRes *apiModuleRuntimesResponsesBillOfMaterial.BillOfMaterialRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	productDescByBPRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	descriptionMapper := services.ProductDescByBPMapper(
		productDescByBPRes.Message.ProductDescByBP,
	)

	data := apiOutputFormatter.BillOfMaterial{}

	for _, v := range *headerRes.Message.Header {
		img := services.ReadProductImage(
			productDocRes,
			v.OwnerProductionPlantBusinessPartner,
			v.Product,
		)

		productDescription := fmt.Sprintf("%s", descriptionMapper[v.Product].ProductDescription)
		plantName := fmt.Sprintf("%s", plantMapper[strconv.Itoa(v.OwnerProductionPlantBusinessPartner)].PlantName)

		data.BillOfMaterialHeader = append(data.BillOfMaterialHeader,
			apiOutputFormatter.BillOfMaterialHeader{
				BillOfMaterial:                          v.BillOfMaterial,
				Product:                                 v.Product,
				ProductDescription:                      &productDescription,
				Seller:                                  v.Seller,
				SellerName:                              businessPartnerMapper[v.Seller].BusinessPartnerName,
				Buyer:                                   v.Buyer,
				BuyerName:                               businessPartnerMapper[v.Buyer].BusinessPartnerName,
				OwnerProductionPlant:                    v.OwnerProductionPlant,
				OwnerProductionPlantName:                plantName,
				ProductBaseUnit:                         v.ProductBaseUnit,
				ProductProductionUnit:                   v.ProductProductionUnit,
				ProductStandardQuantityInBaseUnit:       v.ProductStandardQuantityInBaseUnit,
				ProductStandardQuantityInProductionUnit: v.ProductStandardQuantityInProductionUnit,
				ValidityStartDate:                       v.ValidityStartDate,
				Images: apiOutputFormatter.Images{
					Product: img,
				},
			},
		)
	}

	for _, v := range *itemRes.Message.Item {
		plantName := fmt.Sprintf("%s", plantMapper[v.StockConfirmationPlant].PlantName)
		img := services.ReadProductImage(
			productDocRes,
			v.ProductionPlantBusinessPartner,
			v.Product,
		)

		data.BillOfMaterialItem = append(data.BillOfMaterialItem,
			apiOutputFormatter.BillOfMaterialItem{
				BillOfMaterial:                                 v.BillOfMaterial,
				BillOfMaterialItem:                             v.BillOfMaterialItem,
				ComponentProduct:                               v.ComponentProduct,
				ComponentProductBuyer:                          v.ComponentProductBuyer,
				ComponentProductBuyerName:                      businessPartnerMapper[v.ComponentProductBuyer].BusinessPartnerName,
				ComponentProductSeller:                         v.ComponentProductSeller,
				ComponentProductSellerName:                     businessPartnerMapper[v.ComponentProductSeller].BusinessPartnerName,
				StockConfirmationPlant:                         v.StockConfirmationPlant,
				StockConfirmationPlantName:                     plantName,
				ComponentProductStandardQuantityInBaseUnit:     v.ComponentProductStandardQuantityInBaseUnit,
				ComponentProductStandardQuantityInDeliveryUnit: v.ComponentProductStandardQuantityInDeliveryUnit,
				ComponentProductBaseUnit:                       v.ComponentProductBaseUnit,
				ComponentProductDeliveryUnit:                   v.ComponentProductDeliveryUnit,
				BillOfMaterialItemText:                         v.BillOfMaterialItemText,
				ValidityStartDate:                              v.ValidityStartDate,
				IsMarkedForDeletion:                            v.IsMarkedForDeletion,
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
