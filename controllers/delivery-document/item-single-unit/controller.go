package controllersDeliveryDocumentSingleUnit

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsDeliveryDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/delivery-document/delivery-document"
	apiModuleRuntimesRequestsDeilveryDocumentDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/delivery-document/delivery-document-doc"
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

type DeliveryDocumentSingleUnitController struct {
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

func (controller *DeliveryDocumentSingleUnitController) Get() {
	//isReleased, _ := controller.GetBool("isReleased")
	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "delivery-document"
	redisKeyCategory2 := "delivery-document-item-single-unit"
	deliveryDocument, _ := controller.GetInt("deliveryDocument")
	deliveryDocumentItem, _ := controller.GetInt("deliveryDocumentItem")
	userType := controller.GetString(":userType")
	pDeliverToParty, _ := controller.GetInt("deliverToParty")
	pDeliverFromParty, _ := controller.GetInt("deliverFromParty")

	DeliveryDocumentSingleUnit := apiInputReader.DeliveryDocument{}

	headerCompleteDeliveryIsDefined := false
	headerDeliveryBlockStatus := false
	headerDeliveryStatus := "CL"
	isCancelled := false
	isMarkedForDeletion := false

	itemCompleteDeliveryIsDefined := false
	itemDeliveryBlockStatus := false

	if userType == deliverToParty {
		DeliveryDocumentSingleUnit = apiInputReader.DeliveryDocument{
			DeliveryDocumentHeader: &apiInputReader.DeliveryDocumentHeader{
				DeliveryDocument:                deliveryDocument,
				DeliverToParty:                  &pDeliverToParty,
				HeaderCompleteDeliveryIsDefined: &headerCompleteDeliveryIsDefined,
				HeaderDeliveryBlockStatus:       &headerDeliveryBlockStatus,
				HeaderDeliveryStatus:            &headerDeliveryStatus,
				IsCancelled:                     &isCancelled,
				IsMarkedForDeletion:             &isMarkedForDeletion,
			},
			DeliveryDocumentItems: &apiInputReader.DeliveryDocumentItems{
				DeliveryDocument:              deliveryDocument,
				DeliveryDocumentItem:          &deliveryDocumentItem,
				ItemCompleteDeliveryIsDefined: &itemCompleteDeliveryIsDefined,
				ItemDeliveryBlockStatus:       &itemDeliveryBlockStatus,
				IsCancelled:                   &isCancelled,
				IsMarkedForDeletion:           &isMarkedForDeletion,
			},
			//DeliveryDocumentDocItemDoc: &apiInputReader.DeliveryDocumentDocItemDoc{
			//	OrderID:                  orderId,
			//	OrderItem:                orderItem,
			//	DocType:                  "QRCODE",
			//	DocIssuerBusinessPartner: *controller.UserInfo.BusinessPartner,
			//},
		}
	} else {
		DeliveryDocumentSingleUnit = apiInputReader.DeliveryDocument{
			DeliveryDocumentHeader: &apiInputReader.DeliveryDocumentHeader{
				DeliveryDocument:                deliveryDocument,
				DeliverFromParty:                &pDeliverFromParty,
				HeaderCompleteDeliveryIsDefined: &headerCompleteDeliveryIsDefined,
				HeaderDeliveryBlockStatus:       &headerDeliveryBlockStatus,
				HeaderDeliveryStatus:            &headerDeliveryStatus,
				IsCancelled:                     &isCancelled,
				IsMarkedForDeletion:             &isMarkedForDeletion,
			},
			DeliveryDocumentItems: &apiInputReader.DeliveryDocumentItems{
				DeliveryDocument:              deliveryDocument,
				DeliveryDocumentItem:          &deliveryDocumentItem,
				ItemCompleteDeliveryIsDefined: &itemCompleteDeliveryIsDefined,
				ItemDeliveryBlockStatus:       &itemDeliveryBlockStatus,
				IsCancelled:                   &isCancelled,
				IsMarkedForDeletion:           &isMarkedForDeletion,
			},
			//DeliveryDocumentDocItemDoc: &apiInputReader.DeliveryDocumentDocItemDoc{
			//	OrderID:                  orderId,
			//	OrderItem:                orderItem,
			//	DocType:                  "QRCODE",
			//	DocIssuerBusinessPartner: *controller.UserInfo.BusinessPartner,
			//},
		}
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			redisKeyCategory2,
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
			controller.request(DeliveryDocumentSingleUnit)
		}()
	} else {
		controller.request(DeliveryDocumentSingleUnit)
	}
}

func (
	controller *DeliveryDocumentSingleUnitController,
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
		controller.CustomLogger.Error("createDeliveryDocumentRequestHeaderByDeliverToParty Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentSingleUnitController,
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
		controller.CustomLogger.Error("createDeliveryDocumentRequestHeaderByDeliverFromParty Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentSingleUnitController,
) createDeliveryDocumentRequestItem(
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
		controller.CustomLogger.Error("createDeliveryDocumentRequestItem Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentSingleUnitController,
) createDeliveryDocumentRequestItemScheduleLines(
	requestPram *apiInputReader.Request,
	input apiInputReader.DeliveryDocument,
) *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes {
	responseJsonData := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes{}
	responseBody := apiModuleRuntimesRequestsDeliveryDocument.DeliveryDocumentReads(
		requestPram,
		input,
		&controller.Controller,
		"ItemScheduleLines",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createDeliveryDocumentRequestItemScheduleLines Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentSingleUnitController,
) createDeliveryDocumentRequestItemPricingElements(
	requestPram *apiInputReader.Request,
	input apiInputReader.DeliveryDocument,
) *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes {
	responseJsonData := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes{}
	responseBody := apiModuleRuntimesRequestsDeliveryDocument.DeliveryDocumentReads(
		requestPram,
		input,
		&controller.Controller,
		"ItemPricingElements",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createDeliveryDocumentRequestItemPricingElements Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentSingleUnitController,
) createDeliveryDocumentDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.DeliveryDocument,
) *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentDocRes {
	responseJsonData := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentDocRes{}
	responseBody := apiModuleRuntimesRequestsDeilveryDocumentDoc.DeliveryDocumentDocReads(
		requestPram,
		input,
		&controller.Controller,
		"DeliveryDocumentDoc",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createDeliveryDocumentDocRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentSingleUnitController,
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
	controller *DeliveryDocumentSingleUnitController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	deliveryDocumentItemRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*deliveryDocumentItemRes.Message.Item))

	for _, v := range *deliveryDocumentItemRes.Message.Item {
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
	controller *DeliveryDocumentSingleUnitController,
) createPlantRequest(
	requestPram *apiInputReader.Request,
	deliveryDocumentItemRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*deliveryDocumentItemRes.Message.Item))

	for _, v := range *deliveryDocumentItemRes.Message.Item {
		input = append(input, apiModuleRuntimesRequestsPlant.General{
			Plant: v.DeliverToPlant,
		})
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
		controller.CustomLogger.Error("createPlantRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentSingleUnitController,
) request(
	input apiInputReader.DeliveryDocument,
) {
	defer services.Recover(controller.CustomLogger)

	deliveryDocumentHeaderRes := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes{}

	if input.DeliveryDocumentHeader.DeliverToParty != nil {
		deliveryDocumentHeaderRes = *controller.createDeliveryDocumentRequestHeaderByDeliverToParty(
			controller.UserInfo,
			input,
		)
	}

	if input.DeliveryDocumentHeader.DeliverFromParty != nil {
		deliveryDocumentHeaderRes = *controller.createDeliveryDocumentRequestHeaderByDeliverFromParty(
			controller.UserInfo,
			input,
		)
	}

	deliveryDocumentItemRes := controller.createDeliveryDocumentRequestItem(
		controller.UserInfo,
		input,
	)

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		deliveryDocumentItemRes,
	)

	productDocRes := controller.createProductMasterDocRequest(
		controller.UserInfo,
	)

	plantRes := *controller.createPlantRequest(
		controller.UserInfo,
		deliveryDocumentItemRes,
	)

	deliveryDocumentItemDocRes := controller.createDeliveryDocumentDocRequest(
		controller.UserInfo,
		input,
	)

	controller.fin(
		&deliveryDocumentHeaderRes,
		deliveryDocumentItemRes,
		&businessPartnerRes,
		productDocRes,
		&plantRes,
		deliveryDocumentItemDocRes,
	)
}

func (
	controller *DeliveryDocumentSingleUnitController,
) fin(
	deliveryDocumentHeaderRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
	deliveryDocumentItemRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
	deliveryDocumentItemDocRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentDocRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)
	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	data := apiOutputFormatter.DeliveryDocument{}

	for _, v := range *deliveryDocumentItemRes.Message.Item {
		img := services.CreateProductImage(
			productDocRes,
			*controller.UserInfo.BusinessPartner,
			v.Product,
		)

		qrcode := services.CreateQRCodeDeliveryDocumentItemDocImage(
			deliveryDocumentItemDocRes,
		)

		data.DeliveryDocumentSingleUnit = append(data.DeliveryDocumentSingleUnit,
			apiOutputFormatter.DeliveryDocumentSingleUnit{
				DeliveryDocument:        v.DeliveryDocument,
				DeliveryDocumentItem:    v.DeliveryDocumentItem,
				PlannedGoodsIssueDate:   v.PlannedGoodsIssueDate,
				PlannedGoodsIssueTime:   v.PlannedGoodsIssueTime,
				PlannedGoodsReceiptDate: v.PlannedGoodsReceiptDate,
				PlannedGoodsReceiptTime: v.PlannedGoodsReceiptTime,
				DeliverToParty:          v.DeliverToParty,
				DeliverToPartyName:      businessPartnerMapper[v.DeliverToParty].BusinessPartnerName,
				DeliverToPlant:          v.DeliverToPlant,
				DeliverToPlantName:      plantMapper[v.DeliverToPlant].PlantName,
				DeliverFromParty:        v.DeliverFromParty,
				DeliverFromPartyName:    businessPartnerMapper[v.DeliverFromParty].BusinessPartnerName,
				DeliverFromPlant:        v.DeliverFromPlant,
				DeliverFromPlantName:    plantMapper[v.DeliverFromPlant].PlantName,

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
