package controllersPurchaseRequisitionList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsPurchaseRequisition "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/purchase-requisition"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesPurchaseRequisition "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/purchase-requisition"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type PurchaseRequisitionListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const (
	buyer = "buyer"
	//	seller = "seller"
)

func (controller *PurchaseRequisitionListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	isCancelled, _ := controller.GetBool("isCancelled")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "purchase-requisition"
	redisKeyCategory2 := "list"
	userType := controller.GetString(":userType") // buyer

	purchaseRequisitionHeader := apiInputReader.PurchaseRequisition{}

	headerCompleteOrderIsDefined := false
	headerOrderStatus := "CL"
	isCancelled = false
	isMarkedForDeletion = false

	if userType == buyer {
		purchaseRequisitionHeader = apiInputReader.PurchaseRequisition{
			PurchaseRequisitionHeader: &apiInputReader.PurchaseRequisitionHeader{
				Buyer:                        controller.UserInfo.BusinessPartner,
				HeaderCompleteOrderIsDefined: &headerCompleteOrderIsDefined,
				HeaderOrderStatus:            &headerOrderStatus,
				IsCancelled:                  &isCancelled,
				IsMarkedForDeletion:          &isMarkedForDeletion,
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
		var responseData apiOutputFormatter.PurchaseRequisition

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
			controller.request(purchaseRequisitionHeader)
		}()
	} else {
		controller.request(purchaseRequisitionHeader)
	}
}

func (
	controller *PurchaseRequisitionListController,
) createPurchaseRequisitionRequestHeaderByBuyer(
	requestPram *apiInputReader.Request,
	input apiInputReader.PurchaseRequisition,
) *apiModuleRuntimesResponsesPurchaseRequisition.PurchaseRequisitionRes {
	responseJsonData := apiModuleRuntimesResponsesPurchaseRequisition.PurchaseRequisitionRes{}
	responseBody := apiModuleRuntimesRequestsPurchaseRequisition.PurchaseRequisitionReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersByBuyer",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("PurchaseRequisitionReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PurchaseRequisitionListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	businessPartnerRes *apiModuleRuntimesResponsesPurchaseRequisition.PurchaseRequisitionRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*businessPartnerRes.Message.Header))

	for _, v := range *businessPartnerRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.Buyer,
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
	controller *PurchaseRequisitionListController,
) request(
	input apiInputReader.PurchaseRequisition,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := apiModuleRuntimesResponsesPurchaseRequisition.PurchaseRequisitionRes{}

	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	headerRes = *controller.createPurchaseRequisitionRequestHeaderByBuyer(
		controller.UserInfo,
		input,
	)

	businessPartnerRes = *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		&headerRes,
	)

	controller.fin(
		&headerRes,
		&businessPartnerRes,
	)
}

func (
	controller *PurchaseRequisitionListController,
) fin(
	headerRes *apiModuleRuntimesResponsesPurchaseRequisition.PurchaseRequisitionRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.PurchaseRequisition{}

	for _, v := range *headerRes.Message.Header {

		data.PurchaseRequisitionHeader = append(data.PurchaseRequisitionHeader,
			apiOutputFormatter.PurchaseRequisitionHeader{
				PurchaseRequisition:     v.PurchaseRequisition,
				Buyer:                   v.Buyer,
				BuyerName:               businessPartnerMapper[v.Buyer].BusinessPartnerName,
				HeaderOrdertatus:        &v.HeaderOrderStatus,
				PurchaseRequisitionType: v.PurchaseRequisitionType,
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
