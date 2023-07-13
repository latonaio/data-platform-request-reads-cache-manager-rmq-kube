package controllersInvoiceDocumentList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequests "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsInvoiceDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/invoice-document"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesInvoiceDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/invoice-document"
	apiModuleRuntimesResponses "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master-doc"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type InvoiceDocumentListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const (
	billToParty  = "billToParty"
	billFromParty = "billFromParty"
)

func (controller *InvoiceDocumentListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	isMarkedForDeletion, _ := controller.GetBool("isMarkedForDeletion")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "invoiceDocument"
	redisKeyCategory2 := "list"
	userType := controller.GetString("userType") // billToParty or billFromParty
	billToPartyValue, _ := controller.GetInt("billToParty")
	billFromPartyValue, _ := controller.GetInt("billFromParty")

	invoiceDocumentHeader := apiInputReader.InvoiceDocument{}

	if userType == billToParty {
		invoiceDocumentHeader = apiInputReader.InvoiceDocument{
			InvoiceDocumentHeader: &apiInputReader.InvoiceDocumentHeader{
				BillToParty:      &billToPartyValue,
				IsMarkedForDeletion: &isMarkedForDeletion,
			},
		}
	}

	if userType == billFromParty {
		invoiceDocumentHeader = apiInputReader.InvoiceDocument{
			InvoiceDocumentHeader:  &apiInputReader.InvoiceDocumentHeader{
				BillFromParty:    &billFromPartyValue,
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
		var responseData apiOutputFormatter.InvoiceDocument

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
			controller.request(invoiceDocumentHeader)
		}()
	} else {
		controller.request(invoiceDocumentHeader)
	}
}

func (
	controller *InvoiceDocumentListController,
) createInvoiceDocumentRequestHeaderByBillToParty(
	requestPram *apiInputReader.Request,
	input apiInputReader.InvoiceDocument,
) *apiModuleRuntimesResponsesInvoiceDocument.InvoiceDocumentRes {
	responseJsonData := apiModuleRuntimesResponsesInvoiceDocument.InvoiceDocumentRes{}
	responseBody := apiModuleRuntimesRequestsInvoiceDocument.InvoiceDocumentReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersByBillToParty",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("InvoiceDocumentReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *InvoiceDocumentListController,
) createInvoiceDocumentRequestHeaderByBillFromParty(
	requestPram *apiInputReader.Request,
	input apiInputReader.InvoiceDocument,
) *apiModuleRuntimesResponsesInvoiceDocument.InvoiceDocumentRes {
	responseJsonData := apiModuleRuntimesResponsesInvoiceDocument.InvoiceDocumentRes{}
	responseBody := apiModuleRuntimesRequestsInvoiceDocument.InvoiceDocumentReads(
		requestPram,
		input,
		&controller.Controller,
		"HeadersByBillFromParty",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("InvoiceDocumentReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *InvoiceDocumentListController,
) createBusinessPartnerRequestByBillToParty(
	requestPram *apiInputReader.Request,
	invoiceDocumentRes *apiModuleRuntimesResponsesInvoiceDocument.InvoiceDocumentRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, 0)

	for _, v := range *invoiceDocumentRes.Message.Header {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.BillToParty,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReads(
		requestPram,
		generals,
		&controller.Controller,
		"GeneralsByBusinessPartners",
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
	controller *InvoiceDocumentListController,
) createBusinessPartnerRequestByBillFromParty(
	requestPram *apiInputReader.Request,
	invoiceDocumentRes *apiModuleRuntimesResponsesInvoiceDocument.InvoiceDocumentRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	generals := make([]apiModuleRuntimesRequestsBusinessPartner.General, 0)

	for _, v := range *invoiceDocumentRes.Message.Header {
		generals = append(generals, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.BillFromParty,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReads(
		requestPram,
		generals,
		&controller.Controller,
		"GeneralsByBusinessPartners",
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
	controller *InvoiceDocumentListController,
) createProductMasterDocRequest(
	requestPram *apiInputReader.Request,
) *apiModuleRuntimesResponses.ProductMasterDocRes {
	responseJsonData := apiModuleRuntimesResponses.ProductMasterDocRes{}
	responseBody := apiModuleRuntimesRequests.ProductMasterDocReads(
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
	controller *InvoiceDocumentListController,
) request(
	input apiInputReader.InvoiceDocument,
) {
	defer services.Recover(controller.CustomLogger)

	invoiceDocumentRes := apiModuleRuntimesResponsesInvoiceDocument.InvoiceDocumentRes{}
	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	if input.InvoiceDocumentHeader.BillToParty != nil {
		invoiceDocumentRes = *controller.createInvoiceDocumentRequestHeaderByBillToParty(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequestByBillToParty(
			controller.UserInfo,
			&invoiceDocumentRes,
		)
	}

	if input.InvoiceDocumentHeader.BillFromParty != nil {
		invoiceDocumentRes = *controller.createInvoiceDocumentRequestHeaderByBillFromParty(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequestByBillFromParty(
			controller.UserInfo,
			&invoiceDocumentRes,
		)
	}

	controller.fin(
		&invoiceDocumentRes,
		&businessPartnerRes,
	)
}

func (
	controller *InvoiceDocumentListController,
) fin(
	invoiceDocumentRes *apiModuleRuntimesResponsesInvoiceDocument.InvoiceDocumentRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.InvoiceDocument{}

	for _, v := range *invoiceDocumentRes.Message.Header {

		data.InvoiceDocumentHeader = append(data.InvoiceDocumentHeader,
			apiOutputFormatter.InvoiceDocumentHeader{
				InvoiceDocument:	    	v.InvoiceDocument,
				BillToParty:            	v.BillToParty,
				BillToPartyName:        	businessPartnerMapper[v.BillToParty].BusinessPartnerName,
				BillFromParty:          	v.BillFromParty,
				BillFromPartyName:      	businessPartnerMapper[v.BillFromParty].BusinessPartnerName,
				InvoiceDocumentDate			v.InvoiceDocumentDate,
				PaymentDueDate				v.PaymentDueDate,
				HeaderBillingIsConfirmed	v.HeaderBillingIsConfirmed,
				IsCancelled			    	v.IsCancelled,
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
