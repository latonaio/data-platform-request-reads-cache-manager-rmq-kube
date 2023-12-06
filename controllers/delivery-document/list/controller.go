package controllersDeliveryDocumentList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsDeliveryDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/delivery-document/delivery-document"
	apiModuleRuntimesRequestsPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/plant"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesDeliveryDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/delivery-document"
	apiModuleRuntimesResponsesPlant "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/plant"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type DeliveryDocumentListController struct {
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

func (controller *DeliveryDocumentListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "deliveryDocument"
	redisKeyCategory2 := "list"
	userType := controller.GetString(":userType") // deliverToParty or deliverFromParty

	deliveryDocumentHeader := apiInputReader.DeliveryDocument{}

	headerCompleteDeliveryIsDefined := false
	headerDeliveryBlockStatus := false
	headerDeliveryStatus := "CL"

	if userType == deliverToParty {
		deliveryDocumentHeader = apiInputReader.DeliveryDocument{
			DeliveryDocumentHeader: &apiInputReader.DeliveryDocumentHeader{
				DeliverToParty:                  controller.UserInfo.BusinessPartner,
				HeaderCompleteDeliveryIsDefined: &headerCompleteDeliveryIsDefined,
				HeaderDeliveryBlockStatus:       &headerDeliveryBlockStatus,
				HeaderDeliveryStatus:            &headerDeliveryStatus,
				IsMarkedForDeletion:             &isMarkedForDeletion,
			},
		}
	}

	if userType == deliverFromParty {
		deliveryDocumentHeader = apiInputReader.DeliveryDocument{
			DeliveryDocumentHeader: &apiInputReader.DeliveryDocumentHeader{
				DeliverFromParty:                controller.UserInfo.BusinessPartner,
				HeaderCompleteDeliveryIsDefined: &headerCompleteDeliveryIsDefined,
				HeaderDeliveryBlockStatus:       &headerDeliveryBlockStatus,
				HeaderDeliveryStatus:            &headerDeliveryStatus,
				IsMarkedForDeletion:             &isMarkedForDeletion,
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
	controller *DeliveryDocumentListController,
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
	controller *DeliveryDocumentListController,
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
	controller *DeliveryDocumentListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	deliveryDocumentRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*deliveryDocumentRes.Message.Header))

	for _, v := range *deliveryDocumentRes.Message.Header {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.DeliverToParty,
		})
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.DeliverFromParty,
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
	controller *DeliveryDocumentListController,
) createPlantRequest(
	requestPram *apiInputReader.Request,
	plantRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
) *apiModuleRuntimesResponsesPlant.PlantRes {
	input := make([]apiModuleRuntimesRequestsPlant.General, len(*plantRes.Message.Header))

	for _, v := range *plantRes.Message.Header {
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
		controller.CustomLogger.Error("PlantReadsGeneralsByPlants Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentListController,
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
		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&headerRes,
		)
		plantRes = *controller.createPlantRequest(
			controller.UserInfo,
			&headerRes,
		)
	}

	if input.DeliveryDocumentHeader.DeliverFromParty != nil {
		headerRes = *controller.createDeliveryDocumentRequestHeaderByDeliverFromParty(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&headerRes,
		)
		plantRes = *controller.createPlantRequest(
			controller.UserInfo,
			&headerRes,
		)
	}

	controller.fin(
		&headerRes,
		&businessPartnerRes,
		&plantRes,
	)
}

func (
	controller *DeliveryDocumentListController,
) fin(
	headerRes *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	plantRes *apiModuleRuntimesResponsesPlant.PlantRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)
	plantMapper := services.PlantMapper(
		plantRes.Message.General,
	)

	data := apiOutputFormatter.DeliveryDocument{}

	for _, v := range *headerRes.Message.Header {

		data.DeliveryDocumentHeader = append(data.DeliveryDocumentHeader,
			apiOutputFormatter.DeliveryDocumentHeader{
				DeliveryDocument:        v.DeliveryDocument,
				DeliverToParty:          v.DeliverToParty,
				DeliverToPartyName:      businessPartnerMapper[v.DeliverToParty].BusinessPartnerName,
				DeliverToPlant:          v.DeliverToPlant,
				DeliverToPlantName:      plantMapper[v.DeliverToPlant].PlantName,
				DeliverFromParty:        v.DeliverFromParty,
				DeliverFromPartyName:    businessPartnerMapper[v.DeliverFromParty].BusinessPartnerName,
				DeliverFromPlant:        v.DeliverFromPlant,
				DeliverFromPlantName:    plantMapper[v.DeliverFromPlant].PlantName,
				HeaderDeliveryStatus:    v.HeaderDeliveryStatus,
				HeaderBillingStatus:     v.HeaderBillingStatus,
				PlannedGoodsReceiptDate: v.PlannedGoodsReceiptDate,
				PlannedGoodsReceiptTime: v.PlannedGoodsReceiptTime,
				IsCancelled:             v.IsCancelled,
				IsMarkedForDeletion:     v.IsMarkedForDeletion,
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
