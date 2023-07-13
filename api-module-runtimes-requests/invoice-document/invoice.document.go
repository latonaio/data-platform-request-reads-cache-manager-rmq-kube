package apiModuleRuntimesRequestsInvoiceDocument

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type InvoiceDocumentReq struct {
	Header   Header   `json:"InvoiceDocument"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	InvoiceDocument                   int      `json:"InvoiceDocument"`
	SupplyChainRelationshipID         *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID  *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID  *int     `json:"SupplyChainRelationshipPaymentID"`
	BillToParty                       *int     `json:"BillToParty"`
	BillFromParty                     *int     `json:"BillFromParty"`
	BillToCountry                     *string  `json:"BillToCountry"`
	BillFromCountry                   *string  `json:"BillFromCountry"`
	Payer                             *int     `json:"Payer"`
	Payee                             *int     `json:"Payee"`
	InvoiceDocumentDate               *string  `json:"InvoiceDocumentDate"`
	InvoiceDocumentTime               *string  `json:"InvoiceDocumentTime"`
	InvoicePeriodStartDate            *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate              *string  `json:"InvoicePeriodEndDate"`
	AccountingPostingDate             *string  `json:"AccountingPostingDate"`
	IsExportImport                    *bool    `json:"IsExportImport"`
	HeaderBillingIsConfirmed          *bool    `json:"HeaderBillingIsConfirmed"`
	HeaderBillingConfStatus           *string  `json:"HeaderBillingConfStatus"`
	TotalNetAmount                    *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                    *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                  *float32 `json:"TotalGrossAmount"`
	TransactionCurrency               *string  `json:"TransactionCurrency"`
	Incoterms                         *string  `json:"Incoterms"`
	PaymentTerms                      *string  `json:"PaymentTerms"`
	DueCalculationBaseDate            *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                    *string  `json:"PaymentDueDate"`
	NetPaymentDays                    *int     `json:"NetPaymentDays"`
	PaymentMethod                     *string  `json:"PaymentMethod"`
	ExternalReferenceDocument         *string  `json:"ExternalReferenceDocument"`
	DocumentHeaderText                *string  `json:"DocumentHeaderText"`
	HeaderIsCleared                   *bool    `json:"HeaderIsCleared"`
	HeaderPaymentBlockStatus          *bool    `json:"HeaderPaymentBlockStatus"`
	HeaderPaymentRequisitionIsCreated *bool    `json:"HeaderPaymentRequisitionIsCreated"`
	CreationDate                      *string  `json:"CreationDate"`
	CreationTime                      *string  `json:"CreationTime"`
	LastChangeDate                    *string  `json:"LastChangeDate"`
	LastChangeTime                    *string  `json:"LastChangeTime"`
	IsCancelled                       *bool    `json:"IsCancelled"`
	Item                              []Item      `json:"Item"`
}

type Item struct {
	InvoiceDocument                         int      `json:"InvoiceDocument"`
	InvoiceDocumentItem                     int      `json:"InvoiceDocumentItem"`
	InvoiceDocumentItemCategory             *string  `json:"InvoiceDocumentItemCategory"`
	SupplyChainRelationshipID               *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID       *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID  *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	InvoiceDocumentItemText                 *string  `json:"InvoiceDocumentItemText"`
	InvoiceDocumentItemTextByBuyer          *string  `json:"InvoiceDocumentItemTextByBuyer"`
	InvoiceDocumentItemTextBySeller         *string  `json:"InvoiceDocumentItemTextBySeller"`
	Product                                 *string  `json:"Product"`
	ProductGroup                            *string  `json:"ProductGroup"`
	ProductStandardID                       *string  `json:"ProductStandardID"`
	ItemBillingIsConfirmed                  *bool    `json:"ItemBillingIsConfirmed"`
	Buyer                                   *int     `json:"Buyer"`
	Seller                                  *int     `json:"Seller"`
	DeliverToParty                          *int     `json:"DeliverToParty"`
	DeliverFromParty                        *int     `json:"DeliverFromParty"`
	DeliverToPlant                          *string  `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation           *string  `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlant                        *string  `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation         *string  `json:"DeliverFromPlantStorageLocation"`
	ProductionPlantBusinessPartner          *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                         *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation          *string  `json:"ProductionPlantStorageLocation"`
	ServicesRenderedDate                    *string  `json:"ServicesRenderedDate"`
	InvoiceQuantity                         *float32 `json:"InvoiceQuantity"`
	InvoiceQuantityUnit                     *string  `json:"InvoiceQuantityUnit"`
	InvoiceQuantityInBaseUnit               *float32 `json:"InvoiceQuantityInBaseUnit"`
	BaseUnit                                *string  `json:"BaseUnit"`
	DeliveryUnit                            *string  `json:"DeliveryUnit"`
	ActualGoodsIssueDate                    *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                    *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate                  *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                  *string  `json:"ActualGoodsReceiptTime"`
	ItemGrossWeight                         *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                           *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                          *string  `json:"ItemWeightUnit"`
	NetAmount                               *float32 `json:"NetAmount"`
	TaxAmount                               *float32 `json:"TaxAmount"`
	GrossAmount                             *float32 `json:"GrossAmount"`
	GoodsIssueOrReceiptSlipNumber           *string  `json:"GoodsIssueOrReceiptSlipNumber"`
	TransactionCurrency                     *string  `json:"TransactionCurrency"`
	PricingDate                             *string  `json:"PricingDate"`
	TransactionTaxClassification            *string  `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   *string  `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry *string  `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                *string  `json:"DefinedTaxClassification"`
	Project                                 *int     `json:"Project"`
	WBSElement                              *int     `json:"WBSElement"`
	OrderID                                 *int     `json:"OrderID"`
	OrderItem                               *int     `json:"OrderItem"`
	OrderType                               *string  `json:"OrderType"`
	ContractType                            *string  `json:"ContractType"`
	OrderVaridityStartDate                  *string  `json:"OrderVaridityStartDate"`
	OrderVaridityEndDate                    *string  `json:"OrderVaridityEndDate"`
	InvoicePeriodStartDate                  *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                    *string  `json:"InvoicePeriodEndDate"`
	DeliveryDocument                        *int     `json:"DeliveryDocument"`
	DeliveryDocumentItem                    *int     `json:"DeliveryDocumentItem"`
	OriginDocument                          *int     `json:"OriginDocument"`
	OriginDocumentItem                      *int     `json:"OriginDocumentItem"`
	ReferenceDocument                       *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                   *int     `json:"ReferenceDocumentItem"`
	ExternalReferenceDocument               *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem           *string  `json:"ExternalReferenceDocumentItem"`
	TaxCode                                 *string  `json:"TaxCode"`
	TaxRate                                 *float32 `json:"TaxRate"`
	CountryOfOrigin                         *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                 *string  `json:"CountryOfOriginLanguage"`
	Equipment                        		*int     `json:"Equipment"`
	ItemPaymentRequisitionIsCreated         *bool    `json:"ItemPaymentRequisitionIsCreated"`
	ItemIsCleared                           *bool    `json:"ItemIsCleared"`
	ItemPaymentBlockStatus                  *bool    `json:"ItemPaymentBlockStatus"`
	CreationDate                            *string  `json:"CreationDate"`
	CreationTime                            *string  `json:"CreationTime"`
	LastChangeDate                          *string  `json:"LastChangeDate"`
	LastChangeTime                          *string  `json:"LastChangeTime"`
	IsCancelled                             *bool    `json:"IsCancelled"`
}

func CreateInvoiceDocumentRequestHeaderByBillToParty(
	requestPram *apiInputReader.Request,
	invoiceDocumentHeader *apiInputReader.InvoiceDocumentHeader,
) InvoiceDocumentReq {
	req := InvoiceDocumentReq{
		Header: Header{
			BillToParty:    	           		invoiceDocumentHeader.BillToParty,
			IsCancelled:						invoiceDocumentHeader.IsCancelled,
			IsMarkedForDeletion: 				invoiceDocumentHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersByBillToParty",
		},
	}
	return req
}

func CreateInvoiceDocumentRequestHeaderByBillFromParty(
	requestPram *apiInputReader.Request,
	invoiceDocumentHeader *apiInputReader.InvoiceDocumentHeader,
) InvoiceDocumentReq {
	req := InvoiceDocumentReq{
		Header: Header{
			BillFromParty:      				invoiceDocumentHeader.BillFromParty,
			IsCancelled:						invoiceDocumentHeader.IsCancelled,
			IsMarkedForDeletion: 				invoiceDocumentHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersByBillFromParty",
		},
	}
	return req
}

func CreateInvoiceDocumentRequestItems(
	requestPram *apiInputReader.Request,
	invoiceDocumentItems *apiInputReader.InvoiceDocumentItems,
) InvoiceDocumentReq {
	req := InvoiceDocumentReq{
		Header: Header{
			InvoiceDocument: invoiceDocumentItems.OrderID,
			Item: []Item{
				{
					IsCancelled: 		 invoiceDocumentItems.IsCancelled,
					IsMarkedForDeletion: invoiceDocumentItems.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Items",
		},
	}
	return req
}

func InvoiceDocumentReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.InvoiceDocument,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_INVOICE_DOCUMENT_SRV"
	aPIType := "reads"

	var request InvoiceDocumentReq

	if accepter == "HeadersByBillToParty" {
		request = CreateInvoiceDocumentRequestHeaderByBillToParty(
			requestPram,
			&apiInputReader.InvoiceDocumentHeader{
				BillToParty:		 		 		input.InvoiceDocumentHeader.BillToParty,
				IsCancelled: 		 				input.InvoiceDocumentHeader.IsCancelled,
				IsMarkedForDeletion: 				input.InvoiceDocumentHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "HeadersByBillFromParty" {
		request = CreateInvoiceDocumentRequestHeaderByBillFromParty(
			requestPram,
			&apiInputReader.InvoiceDocumentHeader{
				BillFromParty:						input.InvoiceDocumentHeader.BillFromParty,
				IsCancelled:						input.InvoiceDocumentHeader.IsCancelled,
				IsMarkedForDeletion:				input.InvoiceDocumentHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Items" {
		request = CreateInvoiceDocumentRequestItems(
			requestPram,
			&apiInputReader.InvoiceDocumentItems{
				InvoiceDocument:      input.InvoiceDocumentItems.InvoiceDocument,
				IsCancelled:		  input.InvoiceDocumentItems.IsCancelled,
				IsMarkedForDeletion:  input.InvoiceDocumentItems.IsMarkedForDeletion,
			},
		)
	}

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}
