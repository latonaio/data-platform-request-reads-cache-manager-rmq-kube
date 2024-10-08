package controllersQuotationsList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsQuotations "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/quotations"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesQuotations "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/quotations"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type QuotationsListController struct {
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

func (controller *QuotationsListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	_, _ = controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	redisKeyCategory1 := "quotations"
	redisKeyCategory2 := "list"
	userType := controller.GetString(":userType") // buyer or seller

	quotationsHeader := apiInputReader.Quotations{}

	if userType == buyer {
		quotationsHeader = apiInputReader.Quotations{
			QuotationsHeader: &apiInputReader.QuotationsHeader{
				Buyer: controller.UserInfo.BusinessPartner,
				//IsMarkedForDeletion: &isMarkedForDeletion,
			},
		}
	}

	if userType == seller {
		quotationsHeader = apiInputReader.Quotations{
			QuotationsHeader: &apiInputReader.QuotationsHeader{
				Seller: controller.UserInfo.BusinessPartner,
				//IsMarkedForDeletion: &isMarkedForDeletion,
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
		var responseData apiOutputFormatter.Quotations

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
			controller.request(quotationsHeader)
		}()
	} else {
		controller.request(quotationsHeader)
	}
}

func (
	controller *QuotationsListController,
) createQuotationsRequestHeaderByBuyer(
	requestPram *apiInputReader.Request,
	input apiInputReader.Quotations,
) *apiModuleRuntimesResponsesQuotations.QuotationsRes {
	responseJsonData := apiModuleRuntimesResponsesQuotations.QuotationsRes{}
	responseBody := apiModuleRuntimesRequestsQuotations.QuotationsReads(
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
		controller.CustomLogger.Error("QuotationsReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *QuotationsListController,
) createQuotationsRequestHeaderBySeller(
	requestPram *apiInputReader.Request,
	input apiInputReader.Quotations,
) *apiModuleRuntimesResponsesQuotations.QuotationsRes {
	responseJsonData := apiModuleRuntimesResponsesQuotations.QuotationsRes{}
	responseBody := apiModuleRuntimesRequestsQuotations.QuotationsReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersBySeller",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("QuotationsReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *QuotationsListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	businessPartnerRes *apiModuleRuntimesResponsesQuotations.QuotationsRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*businessPartnerRes.Message.Header))

	for _, v := range *businessPartnerRes.Message.Header {
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
	controller *QuotationsListController,
) request(
	input apiInputReader.Quotations,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	headerRes := apiModuleRuntimesResponsesQuotations.QuotationsRes{}
	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	if input.QuotationsHeader.Buyer != nil {
		headerRes = *controller.createQuotationsRequestHeaderByBuyer(
			controller.UserInfo,
			input,
		)

		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&headerRes,
		)
	}

	if input.QuotationsHeader.Seller != nil {
		headerRes = *controller.createQuotationsRequestHeaderBySeller(
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
	controller *QuotationsListController,
) fin(
	headerRes *apiModuleRuntimesResponsesQuotations.QuotationsRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.Quotations{}

	for _, v := range *headerRes.Message.Header {

		data.QuotationsHeader = append(data.QuotationsHeader,
			apiOutputFormatter.QuotationsHeader{
				Quotation:            v.Quotation,
				Buyer:                v.Buyer,
				BuyerName:            businessPartnerMapper[v.Buyer].BusinessPartnerName,
				Seller:               v.Seller,
				SellerName:           businessPartnerMapper[v.Seller].BusinessPartnerName,
				HeaderOrderIsDefined: v.HeaderOrderIsDefined,
				QuotationType:        v.QuotationType,
				IsCancelled:          v.IsCancelled,
				IsMarkedForDeletion:  v.IsMarkedForDeletion,
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
