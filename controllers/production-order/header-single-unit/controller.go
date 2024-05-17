package controllersProductionOrderHeaderSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesRequestsProductionOrder "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/production-order/production-order"
	apiModuleRuntimesRequestsProductionOrderDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/production-order/production-order-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiModuleRuntimesResponsesProductionOrder "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/production-order"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type ProductionOrderHeaderSingleUnitController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const ()

func (controller *ProductionOrderHeaderSingleUnitController) Get() {
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "productionOrder"
	redisKeyCategory2 := "header-single-unit"
	productionOrder, _ := controller.GetInt("productionOrder")

	isReleased := false
	isCancelled := false
	isMarkedForDeletion := false

	docType := "QRCODE"

	productionOrderHeaderSingleUnit := apiInputReader.ProductionOrder{
		ProductionOrderHeader: &apiInputReader.ProductionOrderHeader{
			ProductionOrder:     productionOrder,
			IsReleased:          &isReleased,
			IsCancelled:         &isCancelled,
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		ProductionOrderDocHeaderDoc: &apiInputReader.ProductionOrderDocHeaderDoc{
			ProductionOrder:          &productionOrder,
			DocType:                  &docType,
			DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
			strconv.Itoa(productionOrder),
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.ProductionOrder

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
			controller.request(productionOrderHeaderSingleUnit)
		}()
	} else {
		controller.request(productionOrderHeaderSingleUnit)
	}
}

func (
	controller *ProductionOrderHeaderSingleUnitController,
) createProductionOrderRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductionOrder,
) *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes {
	responseJsonData := apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes{}
	responseBody := apiModuleRuntimesRequestsProductionOrder.ProductionOrderReads(
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
		controller.CustomLogger.Error("createProductionOrderRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderHeaderSingleUnitController,
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
		controller.CustomLogger.Error("createProductMasterDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderHeaderSingleUnitController,
) createProductMasterRequestGeneral(
	requestPram *apiInputReader.Request,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}
	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsGeneral(
		requestPram,
		apiModuleRuntimesRequestsProductMaster.General{},
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
	controller *ProductionOrderHeaderSingleUnitController,
) createProductionOrderDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.ProductionOrder,
) *apiModuleRuntimesResponsesProductionOrder.ProductionOrderDocRes {
	responseJsonData := apiModuleRuntimesResponsesProductionOrder.ProductionOrderDocRes{}
	responseBody := apiModuleRuntimesRequestsProductionOrderDoc.ProductionOrderDocReads(
		requestPram,
		input,
		&controller.Controller,
		"HeaderDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductionOrderDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderHeaderSingleUnitController,
) createProductMasterRequestBPPlant(
	requestPram *apiInputReader.Request,
	productionOrderRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}

	bpPlant := apiModuleRuntimesRequestsProductMaster.General{}

	for _, v := range *productionOrderRes.Message.Header {
		bpPlant = apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BPPlant: []apiModuleRuntimesRequestsProductMaster.BPPlant{
						{
							BusinessPartner: v.OwnerProductionPlantBusinessPartner,
							Plant:           v.OwnerProductionPlant,
						},
					},
				},
			},
		}
	}

	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsByBPPlant(
		requestPram,
		bpPlant,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductMasterRequestBPPlant Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderHeaderSingleUnitController,
) createProductMasterRequestProduction(
	requestPram *apiInputReader.Request,
	productionOrderRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterRes{}

	production := apiModuleRuntimesRequestsProductMaster.General{}

	for _, v := range *productionOrderRes.Message.Header {
		production = apiModuleRuntimesRequestsProductMaster.General{
			Product: v.Product,
			BusinessPartner: []apiModuleRuntimesRequestsProductMaster.BusinessPartner{
				{
					BPPlant: []apiModuleRuntimesRequestsProductMaster.BPPlant{
						{
							BusinessPartner: v.OwnerProductionPlantBusinessPartner,
							Plant:           v.OwnerProductionPlant,
						},
					},
				},
			},
		}
	}

	responseBody := apiModuleRuntimesRequestsProductMaster.ProductMasterReadsProduction(
		requestPram,
		production,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createProductMasterRequestProduction Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderHeaderSingleUnitController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	productionOrderItemRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*productionOrderItemRes.Message.Header))

	for _, v := range *productionOrderItemRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.OwnerProductionPlantBusinessPartner,
		})
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Buyer,
		})
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Seller,
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
	controller *ProductionOrderHeaderSingleUnitController,
) createPlantRequestGenerals(
	requestPram *apiInputReader.Request,
	productionOrderHeaderRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*productionOrderHeaderRes.Message.Header))
	for i, v := range *productionOrderHeaderRes.Message.Header {
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
	controller *ProductionOrderHeaderSingleUnitController,
) request(
	input apiInputReader.ProductionOrder,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := controller.createProductionOrderRequestHeader(
		controller.UserInfo,
		input,
	)

	productMasterGeneralRes := controller.createProductMasterRequestGeneral(
		controller.UserInfo,
	)

	productMasterBPPlantRes := controller.createProductMasterRequestBPPlant(
		controller.UserInfo,
		headerRes,
	)

	productMasterProductionRes := controller.createProductMasterRequestProduction(
		controller.UserInfo,
		headerRes,
	)

	plantRes := controller.createPlantRequestGenerals(
		controller.UserInfo,
		headerRes,
	)

	productMasterDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		headerRes,
	)

	headerDocRes := controller.createProductionOrderDocRequest(
		controller.UserInfo,
		input,
	)

	controller.fin(
		headerRes,
		productMasterGeneralRes,
		productMasterBPPlantRes,
		productMasterProductionRes,
		plantRes,
		productMasterDocRes,
		&businessPartnerRes,
		headerDocRes,
	)
}

func (
	controller *ProductionOrderHeaderSingleUnitController,
) fin(
	headerRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderRes,
	productMasterGeneralRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productMasterBPPlantRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	productMasterProductionRes *apiModuleRuntimesResponsesProductMaster.ProductMasterRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	productMasterDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	headerDocRes *apiModuleRuntimesResponsesProductionOrder.ProductionOrderDocRes,
) {
	//generalMapper := services.GeneralsMapper(
	//	productMasterGeneralRes.Message.General,
	//)

	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.ProductionOrder{}

	for _, v := range *headerRes.Message.Header {
		img := services.ReadProductImage(
			productMasterDocRes,
			v.OwnerProductionPlantBusinessPartner,
			v.Product,
		)

		qrcode := services.CreateQRCodeProductionOrderHeaderDocImage(
			headerDocRes,
			v.ProductionOrder,
		)

		var buyerName string

		buyerNameMapperName := businessPartnerMapper[v.Buyer].BusinessPartnerName
		if &buyerNameMapperName != nil {
			buyerName = buyerNameMapperName
		}

		var sellerName string

		sellerNameMapperName := businessPartnerMapper[v.Seller].BusinessPartnerName
		if &sellerNameMapperName != nil {
			sellerName = sellerNameMapperName
		}

		var ownerProductionPlantBusinessPartnerName *string

		ownerProductionPlantBusinessPartnerNameMapperName := businessPartnerMapper[v.OwnerProductionPlantBusinessPartner].BusinessPartnerName
		if &ownerProductionPlantBusinessPartnerName != nil {
			ownerProductionPlantBusinessPartnerName = &ownerProductionPlantBusinessPartnerNameMapperName
		}

		productionPlantName := fmt.Sprintf("%s", plantMapper[strconv.Itoa(v.OwnerProductionPlantBusinessPartner)].PlantName)

		data.ProductionOrderHeaderSingleUnit = append(data.ProductionOrderHeaderSingleUnit,
			apiOutputFormatter.ProductionOrderHeaderSingleUnit{
				ProductionOrder:                         v.ProductionOrder,
				ProductionOrderDate:                     v.ProductionOrderDate,
				Product:                                 v.Product,
				Buyer:                                   v.Buyer,
				BuyerName:                               buyerName,
				Seller:                                  v.Seller,
				SellerName:                              sellerName,
				InspectionLot:                           v.InspectionLot,
				OwnerProductionPlantBusinessPartner:     v.OwnerProductionPlantBusinessPartner,
				OwnerProductionPlantBusinessPartnerName: ownerProductionPlantBusinessPartnerNameMapperName,
				OwnerProductionPlant:                    v.OwnerProductionPlant,
				OwnerProductionPlantName:                productionPlantName,
				ProductionOrderQuantityInBaseUnit:       v.ProductionOrderQuantityInBaseUnit,
				ProductionOrderQuantityInDestinationProductionUnit: v.ProductionOrderQuantityInDestinationProductionUnit,
				Images: apiOutputFormatter.Images{
					Product: img,
					QRCode:  qrcode,
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
