package controllersDeliveryDocumentItem

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

type DeliveryDocumentItemController struct {
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

func (controller *DeliveryDocumentItemController) Get() {
	//isReleased, _ := controller.GetBool("isReleased")
	//isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "delivery-document"
	redisKeyCategory2 := "delivery-document-item"
	deliveryDocument, _ := controller.GetInt("deliveryDocument")
	userType := controller.GetString(":userType")
	pDeliverToParty, _ := controller.GetInt("deliverToParty")
	pDeliverFromParty, _ := controller.GetInt("deliverFromParty")

	DeliveryDocumentItem := apiInputReader.DeliveryDocument{}

	headerCompleteDeliveryIsDefined := false
	headerDeliveryBlockStatus := false
	headerDeliveryStatus := "CL"
	isCancelled := false
	isMarkedForDeletion := false

	itemCompleteDeliveryIsDefined := false
	itemDeliveryBlockStatus := false

	docType := "IMAGE"

	if userType == deliverToParty {
		DeliveryDocumentItem = apiInputReader.DeliveryDocument{
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
				ItemCompleteDeliveryIsDefined: &itemCompleteDeliveryIsDefined,
				ItemDeliveryBlockStatus:       &itemDeliveryBlockStatus,
				IsCancelled:                   &isCancelled,
				IsMarkedForDeletion:           &isMarkedForDeletion,
			},
			DeliveryDocumentDocItemDoc: &apiInputReader.DeliveryDocumentDocItemDoc{
				DeliveryDocument:         deliveryDocument,
				DocType:                  &docType,
				DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
			},
		}
	} else {
		DeliveryDocumentItem = apiInputReader.DeliveryDocument{
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
				ItemCompleteDeliveryIsDefined: &itemCompleteDeliveryIsDefined,
				ItemDeliveryBlockStatus:       &itemDeliveryBlockStatus,
				IsCancelled:                   &isCancelled,
				IsMarkedForDeletion:           &isMarkedForDeletion,
			},
			DeliveryDocumentDocItemDoc: &apiInputReader.DeliveryDocumentDocItemDoc{
				DeliveryDocument:         deliveryDocument,
				DocType:                  &docType,
				DocIssuerBusinessPartner: controller.UserInfo.BusinessPartner,
			},
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
			controller.request(DeliveryDocumentItem)
		}()
	} else {
		controller.request(DeliveryDocumentItem)
	}
}

func (
	controller *DeliveryDocumentItemController,
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
	controller *DeliveryDocumentItemController,
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
	controller *DeliveryDocumentItemController,
) createDeliveryDocumentRequestHeader(
	requestPram *apiInputReader.Request,
	input apiInputReader.DeliveryDocument,
) *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes {
	responseJsonData := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes{}
	responseBody := apiModuleRuntimesRequestsDeliveryDocument.DeliveryDocumentReads(
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
		controller.CustomLogger.Error("DeliveryDocumentReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentItemController,
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
	controller *DeliveryDocumentItemController,
) createDeliveryDocumentDocRequest(
	requestPram *apiInputReader.Request,
	input apiInputReader.DeliveryDocument,
) *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentDocRes {
	responseJsonData := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentDocRes{}
	responseBody := apiModuleRuntimesRequestsDeilveryDocumentDoc.DeliveryDocumentDocReads(
		requestPram,
		input,
		&controller.Controller,
		"ItemDoc",
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
	controller *DeliveryDocumentItemController,
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
	controller *DeliveryDocumentItemController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	deliveryDocumentItemRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	var input []apiModuleRuntimesRequestsBusinessPartner.General

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
	controller *DeliveryDocumentItemController,
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
	controller *DeliveryDocumentItemController,
) request(
	input apiInputReader.DeliveryDocument,
) {
	defer services.Recover(controller.CustomLogger)

	deliveryDocumentHeaderRes := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes{}

	deliveryDocumentHeaderRes = *controller.createDeliveryDocumentRequestHeader(
		controller.UserInfo,
		input,
	)

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
	controller *DeliveryDocumentItemController,
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

	for _, v := range *deliveryDocumentHeaderRes.Message.Header {
		//img := services.ReadProductImage(
		//	productDocRes,
		//	v.DeliverToParty,	//DeliverFromPartyの対応が必要
		//	v.Product,
		//)

		data.DeliveryDocumentHeaderWithItem = append(data.DeliveryDocumentHeaderWithItem,
			apiOutputFormatter.DeliveryDocumentHeaderWithItem{
				DeliveryDocument:        v.DeliveryDocument,
				DeliveryDocumentDate:    v.DeliveryDocumentDate,
				DeliverToParty:          v.DeliverToParty,
				DeliverToPartyName:      businessPartnerMapper[v.DeliverToParty].BusinessPartnerName,
				DeliverFromParty:        v.DeliverFromParty,
				DeliverFromPartyName:    businessPartnerMapper[v.DeliverFromParty].BusinessPartnerName,
				DeliverToPlant:          v.DeliverToPlant,
				DeliverToPlantName:      plantMapper[v.DeliverToPlant].PlantName,
				DeliverFromPlant:        v.DeliverFromPlant,
				DeliverFromPlantName:    plantMapper[v.DeliverFromPlant].PlantName,
				PlannedGoodsIssueDate:   v.PlannedGoodsIssueDate,
				PlannedGoodsIssueTime:   v.PlannedGoodsIssueTime,
				PlannedGoodsReceiptDate: v.PlannedGoodsReceiptDate,
				PlannedGoodsReceiptTime: v.PlannedGoodsReceiptTime,
				HeaderGrossWeight:       *v.HeaderGrossWeight,
				HeaderNetWeight:         *v.HeaderNetWeight,
				HeaderWeightUnit:        *v.HeaderWeightUnit,
			},
		)
	}

	for _, v := range *deliveryDocumentItemRes.Message.Item {
		img := services.ReadProductImage(
			productDocRes,
			*controller.UserInfo.BusinessPartner,
			v.Product,
		)

		documentImage := services.ReadDocumentImageDeliveryDocument(
			deliveryDocumentItemDocRes,
			v.DeliveryDocument,
			v.DeliveryDocumentItem,
		)

		data.DeliveryDocumentItem = append(data.DeliveryDocumentItem,
			apiOutputFormatter.DeliveryDocumentItem{
				DeliveryDocumentItem:           v.DeliveryDocumentItem,
				Product:                        v.Product,
				DeliveryDocumentItemText:       *v.DeliveryDocumentItemText,
				PlannedGoodsIssueDate:          v.PlannedGoodsIssueDate,
				PlannedGoodsIssueTime:          v.PlannedGoodsIssueTime,
				PlannedGoodsReceiptDate:        v.PlannedGoodsReceiptDate,
				PlannedGoodsReceiptTime:        v.PlannedGoodsReceiptTime,
				PlannedGoodsIssueQuantity:      v.PlannedGoodsIssueQuantity,
				PlannedGoodsIssueQtyInBaseUnit: v.PlannedGoodsIssueQtyInBaseUnit,
				DeliveryUnit:                   v.DeliveryUnit,
				BaseUnit:                       v.BaseUnit,

				Images: apiOutputFormatter.Images{
					Product:                       img,
					DocumentImageDeliveryDocument: documentImage,
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
