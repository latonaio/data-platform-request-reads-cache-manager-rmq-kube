package controllersDeliveryDocumentDetailList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsDeliveryDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/delivery-document"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesDeliveryDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/delivery-document"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type DeliveryDocumentDetailListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const (
	deliverToParty   = "deliverToParty"
	deliverFromParty = "deliverFromParty"
)

func (controller *DeliveryDocumentDetailListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "deliveryDocument"
	redisKeyCategory2 := "list"
	userType := controller.GetString(":userType") // deliverToParty or deliverFromParty
	deliverToPartyValue, _ := controller.GetInt("deliverToParty")
	deliverFromPartyValue, _ := controller.GetInt("deliverFromParty")

	deliveryDocumentHeader := apiInputReader.DeliveryDocument{}

	if userType == deliverToParty {
		deliveryDocumentHeader = apiInputReader.DeliveryDocument{
			DeliveryDocumentHeader: &apiInputReader.DeliveryDocumentHeader{
				DeliverToParty:      &deliverToPartyValue,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		}
	}

	if userType == deliverFromParty {
		deliveryDocumentHeader = apiInputReader.DeliveryDocument{
			DeliveryDocumentHeader: &apiInputReader.DeliveryDocumentHeader{
				DeliverFromParty:    &deliverFromPartyValue,
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
		var responseData apiOutputFormatter.DeliveryDocument

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
			controller.request(deliveryDocumentHeader)
		}()
	} else {
		controller.request(deliveryDocumentHeader)
	}
}

func (
	controller *DeliveryDocumentDetailListController,
) createDeliveryDocumentRequestHeaderByDeliverToParty(
	requestPram *apiInputReader.Request,
	input apiInputReader.DeliveryDocument,
) *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes {
	responseJsonData := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes{}
	responseBody := apiModuleRuntimesRequestsDeliveryDocument.DeliveryDocumentReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersByDeliverToParty",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("DeliveryDocumentReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentDetailListController,
) createDeliveryDocumentRequestHeaderByDeliverFromParty(
	requestPram *apiInputReader.Request,
	input apiInputReader.DeliveryDocument,
) *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes {
	responseJsonData := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes{}
	responseBody := apiModuleRuntimesRequestsDeliveryDocument.DeliveryDocumentReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersByDeliverFromParty",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("DeliveryDocumentReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentDetailListController,
) createDeliveryDocumentRequestItems(
	requestPram *apiInputReader.Request,
	input apiInputReader.DeliveryDocument,
) *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes {
	responseJsonData := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes{}
	responseBody := apiModuleRuntimesRequestsDeliveryDocument.DeliveryDocumentReads(
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
		controller.CustomLogger.Error("DeliveryDocumentReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentDetailListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	businessPartnerRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, 0)

	for _, v := range *businessPartnerRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.DeliverToParty,
		})
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.DeliverFromParty,
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
	controller *DeliveryDocumentDetailListController,
) createPlantRequestGeneralsByDeliverToParty(
	requestPram *apiInputReader.Request,
	plantRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, 0)

	for _, v := range *plantRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsPlant.General{
			BusinessPartner: v.DeliverToParty,
			Plant:           v.DeliverToPlant,
		})
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
		controller.CustomLogger.Error("PlantReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentDetailListController,
) createPlantRequestGeneralsByDeliverFromParty(
	requestPram *apiInputReader.Request,
	plantRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, 0)

	for _, v := range *plantRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsPlant.General{
			Plant: v.DeliverFromPlant,
		})
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
		controller.CustomLogger.Error("PlantReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentDetailListController,
) createProductMasterDocRequest(
	requestPram *apiInputReader.Request,
) *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes {
	responseJsonData := apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes{}
	responseBody := apiModuleRuntimesRequestsProductMasterDoc.ProductMasterDocReads(
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
	controller *DeliveryDocumentDetailListController,
) createBusinessPartnerRequestByDeliverToParty(
	requestPram *apiInputReader.Request,
	deliveryDocumentRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, 0)

	for _, v := range *deliveryDocumentRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.DeliverToParty,
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
		controller.CustomLogger.Error("BusinessPartnerReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentDetailListController,
) createBusinessPartnerRequestByDeliverFromParty(
	requestPram *apiInputReader.Request,
	deliveryDocumentRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, 0)

	for _, v := range *deliveryDocumentRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.DeliverFromParty,
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
		controller.CustomLogger.Error("BusinessPartnerReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentDetailListController,
) request(
	input apiInputReader.DeliveryDocument,
) {
	defer services.Recover(controller.CustomLogger)

	headerRes := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes{}
	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	plantRes := apiModuleRuntimesResponsesPlant.PlantRes{}

	if input.DeliveryDocumentHeader.DeliverToParty != nil {
		headerRes = *controller.createDeliveryDocumentRequestHeaderByDeliverToParty(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequestByDeliverToParty(
			controller.UserInfo,
			&headerRes,
		)
		plantRes = *controller.createPlantRequestGeneralsByDeliverToParty(
			controller.UserInfo,
			&headerRes,
		)
	}

	if input.DeliveryDocumentHeader.DeliverFromParty != nil {
		headerRes = *controller.createDeliveryDocumentRequestHeaderByDeliverFromParty(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequestByDeliverFromParty(
			controller.UserInfo,
			&headerRes,
		)
		plantRes = *controller.createPlantRequestGeneralsByDeliverFromParty(
			controller.UserInfo,
			&headerRes,
		)
	}

	itemRes := controller.createDeliveryDocumentRequestItems(
		controller.UserInfo,
		input,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	controller.fin(
		&headerRes,
		itemRes,
		&businessPartnerRes,
		&plantRes,
		productDocRes,
	)
}

func (
	controller *DeliveryDocumentDetailListController,
) fin(
	headerRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
	itemRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {
	//businessPartnerMapper := services.BusinessPartnerNameMapper(
	//	businessPartnerRes,
	//)
	plantMapper := services.PlantMapper(
		plantRes.Message.Generals,
	)

	data := apiOutputFormatter.DeliveryDocument{}

	for _, v := range *headerRes.Message.Header {
		//img := services.CreateProductImage(
		//	productDocRes,
		//	v.DeliverToParty,	//DeliverFromPartyの対応が必要
		//	v.Product,
		//)

		data.DeliveryDocumentHeaderWithItem = append(data.DeliveryDocumentHeaderWithItem,
			apiOutputFormatter.DeliveryDocumentHeaderWithItem{
				DeliveryDocument:        v.DeliveryDocument,
				DeliverToPlant:          v.DeliverToPlant,
				DeliverToPlantName:      plantMapper[v.DeliverToPlant].PlantName,
				DeliverFromPlant:        v.DeliverFromPlant,
				DeliverFromPlantName:    plantMapper[v.DeliverFromPlant].PlantName,
				PlannedGoodsReceiptDate: v.PlannedGoodsReceiptDate,
				PlannedGoodsReceiptTime: v.PlannedGoodsReceiptTime,
			},
		)
	}

	for _, v := range *itemRes.Message.Item {
		data.DeliveryDocumentItem = append(data.DeliveryDocumentItem,
			apiOutputFormatter.DeliveryDocumentItem{
				DeliveryDocumentItem:                 v.DeliveryDocumentItem,
				Product:                              v.Product,
				DeliveryDocumentItemItemTextByBuyer:  *v.DeliveryDocumentItemTextByBuyer,
				DeliveryDocumentItemItemTextBySeller: *v.DeliveryDocumentItemTextBySeller,
				PlannedGoodsIssueQuantity:            v.PlannedGoodsIssueQuantity,
				DeliveryUnit:                         v.DeliveryUnit,
				PlannedGoodsIssueDate:                v.PlannedGoodsIssueDate,
				PlannedGoodsIssueTime:                v.PlannedGoodsIssueTime,
				PlannedGoodsReceiptDate:              v.PlannedGoodsReceiptDate,
				PlannedGoodsReceiptTime:              v.PlannedGoodsReceiptTime,
				IsCancelled:                          v.IsCancelled,
				IsMarkedForDeletion:                  v.IsMarkedForDeletion,
				//Images: apiOutputFormatter.Images{
				//	Product: img,
				//},
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
