package controllersInvoiceDocumentDetailList

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiModuleRuntimesRequestsDeliveryDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/delivery-document"
	apiModuleRuntimesRequestsInvoiceDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/invoice-document"
	apiModuleRuntimesRequestsProductMasterDoc "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/product-master/product-master-doc"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesDeliveryDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/delivery-document"
	apiModuleRuntimesResponsesInvoiceDocument "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/invoice-document"
	apiModuleRuntimesResponsesProductMaster "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/product-master"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type InvoiceDocumentDetailListController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const (
	billToParty   = "billToParty"
	billFromParty = "billFromParty"
)

func (controller *InvoiceDocumentDetailListController) Get() {
	//aPIType := controller.Ctx.Input.Param(":aPIType")
	invoiceDocument, _ := controller.GetInt("invoiceDocument")
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "invoice-document"
	redisKeyCategory2 := "detail-list"
	userType := controller.GetString(":userType") // billToParty or billFromParty
	billToPartyValue, _ := controller.GetInt("billToParty")
	billFromPartyValue, _ := controller.GetInt("billFromParty")

	invoiceDocumentHeader := apiInputReader.InvoiceDocument{}

	if userType == billToParty {
		invoiceDocumentHeader = apiInputReader.InvoiceDocument{
			InvoiceDocumentHeader: &apiInputReader.InvoiceDocumentHeader{
				InvoiceDocument: invoiceDocument,
				BillToParty:     &billToPartyValue,
				// todo 確認
				//IsCancelled:			&isCancelled,
			},
			InvoiceDocumentItems: &apiInputReader.InvoiceDocumentItems{
				InvoiceDocument: invoiceDocument,
				//IsCancelled:			&isCancelled,
			},
		}
	}

	if userType == billFromParty {
		invoiceDocumentHeader = apiInputReader.InvoiceDocument{
			InvoiceDocumentHeader: &apiInputReader.InvoiceDocumentHeader{
				InvoiceDocument: invoiceDocument,
				BillFromParty:   &billFromPartyValue,
				//IsCancelled:     &isCancelled,
			},
			InvoiceDocumentItems: &apiInputReader.InvoiceDocumentItems{
				InvoiceDocument: invoiceDocument,
				//IsCancelled:     &isCancelled,
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
	controller *InvoiceDocumentDetailListController,
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
	controller *InvoiceDocumentDetailListController,
) createInvoiceDocumentRequestItems(
	requestPram *apiInputReader.Request,
	input apiInputReader.InvoiceDocument,
) *apiModuleRuntimesResponsesInvoiceDocument.InvoiceDocumentRes {
	responseJsonData := apiModuleRuntimesResponsesInvoiceDocument.InvoiceDocumentRes{}
	responseBody := apiModuleRuntimesRequestsInvoiceDocument.InvoiceDocumentReads(
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
		controller.CustomLogger.Error("InvoiceDocumentReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *InvoiceDocumentDetailListController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	invoiceDocumentRes *apiModuleRuntimesResponsesInvoiceDocument.InvoiceDocumentRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*invoiceDocumentRes.Message.Header))

	for _, v := range *invoiceDocumentRes.Message.Header {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.BillToParty,
		})
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.BillFromParty,
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
	controller *InvoiceDocumentDetailListController,
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
	controller *InvoiceDocumentDetailListController,
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
	controller *InvoiceDocumentDetailListController,
) request(
	input apiInputReader.InvoiceDocument,
) {
	defer services.Recover(controller.CustomLogger)

	headerRes := apiModuleRuntimesResponsesInvoiceDocument.InvoiceDocumentRes{}
	businessPartnerRes := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}

	if input.InvoiceDocumentHeader.BillToParty != nil {
		headerRes = *controller.createInvoiceDocumentRequestHeaderByBillToParty(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&headerRes,
		)
	}

	if input.InvoiceDocumentHeader.BillFromParty != nil {
		headerRes = *controller.createInvoiceDocumentRequestHeaderByBillToParty(
			controller.UserInfo,
			input,
		)
		businessPartnerRes = *controller.createBusinessPartnerRequest(
			controller.UserInfo,
			&headerRes,
		)
	}

	itemRes := controller.createInvoiceDocumentRequestItems(
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
		productDocRes,
	)
}

func (
	controller *InvoiceDocumentDetailListController,
) fin(
	headerRes *apiModuleRuntimesResponsesInvoiceDocument.InvoiceDocumentRes,
	itemRes *apiModuleRuntimesResponsesInvoiceDocument.InvoiceDocumentRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
	productDocRes *apiModuleRuntimesResponsesProductMaster.ProductMasterDocRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.InvoiceDocument{}

	for _, v := range *headerRes.Message.Header {
		//img := services.CreateProductImage(
		//	productDocRes,
		//	v.BillToParty,	//BillFromPartyの対応が必要
		//	v.Product,
		//)

		data.InvoiceDocumentHeaderWithItem = append(data.InvoiceDocumentHeaderWithItem,
			apiOutputFormatter.InvoiceDocumentHeaderWithItem{
				InvoiceDocument:     v.InvoiceDocument,
				InvoiceDocumentDate: v.InvoiceDocumentDate,
				BillToParty:         v.BillToParty,
				BillToPartyName:     businessPartnerMapper[v.BillToParty].BusinessPartnerName,
				BillFromParty:       v.BillFromParty,
				BillFromPartyName:   businessPartnerMapper[v.BillFromParty].BusinessPartnerName,
			},
		)
	}

	for _, v := range *itemRes.Message.Item {
		data.InvoiceDocumentItem = append(data.InvoiceDocumentItem,
			apiOutputFormatter.InvoiceDocumentItem{
				InvoiceDocumentItem:             v.InvoiceDocumentItem,
				Product:                         v.Product,
				InvoiceDocumentItemTextByBuyer:  v.InvoiceDocumentItemTextByBuyer,
				InvoiceDocumentItemTextBySeller: v.InvoiceDocumentItemTextBySeller,
				InvoiceQuantity:                 v.InvoiceQuantity,
				InvoiceQuantityUnit:             v.InvoiceQuantityUnit,
				ActualGoodsIssueDate:            v.ActualGoodsIssueDate,
				ActualGoodsIssueTime:            v.ActualGoodsIssueTime,
				ActualGoodsReceiptDate:          v.ActualGoodsReceiptDate,
				ItemBillingIsConfirmed:          v.ItemBillingIsConfirmed,
				IsCancelled:                     v.IsCancelled,
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
