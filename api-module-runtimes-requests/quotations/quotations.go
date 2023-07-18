package apiModuleRuntimesRequestsQuotations

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type QuotationsReq struct {
	Header   Header   `json:"Quotations"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Quotation                        int      `json:"Quotation"`
	QuotationDate                    *string  `json:"QuotationDate"`
	QuoationType                     *string  `json:"QuoationType"`
	QuoationStatus                   *string  `json:"QuoationStatus"`
	SupplyChainRelationshipID        *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID *int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            *int     `json:"Buyer"`
	Seller                           *int     `json:"Seller"`
	BillToParty                      *int     `json:"BillToParty"`
	BillFromParty                    *int     `json:"BillFromParty"`
	BillToCountry                    *int     `json:"BillToCountry"`
	BillFromCountry                  *int     `json:"BillFromCountry"`
	Payer                            *int     `json:"Payer"`
	Payee                            *int     `json:"Payee"`
	ContractType                     *string  `json:"ContractType"`
	BindingPeriodValidityStartDate   *string  `json:"BindingPeriodValidityStartDate"`
	BindingPeriodValidityEndDate     *string  `json:"BindingPeriodValidityEndDate"`
	OrderValidityStartDate           *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate             *string  `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate           *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate             *string  `json:"InvoicePeriodEndDate"`
	TotalNetAmount                   *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                   *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                 *float32 `json:"TotalGrossAmount"`
	HeaderOrderIsDefined             *bool    `json:"HeaderOrderIsDefined"`
	TransactionCurrency              *string  `json:"TransactionCurrency"`
	PricingDate                      *string  `json:"PricingDate"`
	PriceDetnExchangeRate            *string  `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate            *string  `json:"RequestedDeliveryDate"`
	OrderProbabilityInPercent        *float32 `json:"OrderProbabilityInPercent"`
	ExpectedOrderNetAmount           *float32 `json:"ExpectedOrderNetAmount"`
	Incoterms                        *string  `json:"Incoterms"`
	PaymentTerms                     *string  `json:"PaymentTerms"`
	PaymentMethod                    *string  `json:"PaymentMethod"`
	ReferenceDocument                *int     `json:"ReferenceDocument"`
	AccountAssignmentGroup           *string  `json:"AccountAssignmentGroup"`
	AccountingExchangeRate           *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate              *string  `json:"InvoiceDocumentDate"`
	IsExportImport                   *bool    `json:"IsExportImport"`
	HeaderText                       *bool    `json:"HeaderText"`
	HeaderIsClosed                   *bool    `json:"HeaderIsClosed"`
	HeaderBlockStatus                *bool    `json:"HeaderBlockStatus"`
	CreationDate                     *string  `json:"CreationDate"`
	LastChangeDate                   *string  `json:"LastChangeDate"`
	IsCancelled                      *bool    `json:"IsCancelled"`
	IsMarkedForDeletion              *bool    `json:"IsMarkedForDeletion"`
	//	Item                        	 []Item      `json:"Item"`
}

func CreateQuotationsRequestHeaderByBuyer(
	requestPram *apiInputReader.Request,
	quotationsHeader *apiInputReader.QuotationsHeader,
) QuotationsReq {
	req := QuotationsReq{
		Header: Header{
			Buyer:                quotationsHeader.Buyer,
			HeaderOrderIsDefined: quotationsHeader.HeaderOrderIsDefined,
			HeaderIsClosed:       quotationsHeader.HeaderIsClosed,
			HeaderBlockStatus:    quotationsHeader.HeaderBlockStatus,
			IsCancelled:          quotationsHeader.IsCancelled,
			IsMarkedForDeletion:  quotationsHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersByBuyer",
		},
	}
	return req
}

func CreateQuotationsRequestHeaderBySeller(
	requestPram *apiInputReader.Request,
	quotationsHeader *apiInputReader.QuotationsHeader,
) QuotationsReq {
	req := QuotationsReq{
		Header: Header{
			Seller:               quotationsHeader.Seller,
			HeaderOrderIsDefined: quotationsHeader.HeaderOrderIsDefined,
			HeaderIsClosed:       quotationsHeader.HeaderIsClosed,
			HeaderBlockStatus:    quotationsHeader.HeaderBlockStatus,
			IsCancelled:          quotationsHeader.IsCancelled,
			IsMarkedForDeletion:  quotationsHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersBySeller",
		},
	}
	return req
}

func CreateQuotationsRequestItems(
	requestPram *apiInputReader.Request,
	quotationsItems *apiInputReader.QuotationsItems,
) QuotationsReq {
	req := QuotationsReq{
		Header: Header{
			Quotation: quotationsItems.Quotation,
			//Item: []Item{
			//	{
			//		IsCancelled:         quotationsItems.IsCancelled,
			//		IsMarkedForDeletion: quotationsItems.IsMarkedForDeletion,
			//	},
			//},
		},
		Accepter: []string{
			"Items",
		},
	}
	return req
}

func QuotationsReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Quotations,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_ORDERS_SRV"
	aPIType := "reads"

	var request QuotationsReq

	if accepter == "HeadersByBuyer" {
		request = CreateQuotationsRequestHeaderByBuyer(
			requestPram,
			&apiInputReader.QuotationsHeader{
				Buyer:                input.QuotationsHeader.Buyer,
				HeaderOrderIsDefined: input.QuotationsHeader.HeaderOrderIsDefined,
				HeaderIsClosed:       input.QuotationsHeader.HeaderIsClosed,
				HeaderBlockStatus:    input.QuotationsHeader.HeaderBlockStatus,
				IsCancelled:          input.QuotationsHeader.IsCancelled,
				IsMarkedForDeletion:  input.QuotationsHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "HeadersBySeller" {
		request = CreateQuotationsRequestHeaderBySeller(
			requestPram,
			&apiInputReader.QuotationsHeader{
				Seller:               input.QuotationsHeader.Seller,
				HeaderOrderIsDefined: input.QuotationsHeader.HeaderOrderIsDefined,
				HeaderIsClosed:       input.QuotationsHeader.HeaderIsClosed,
				HeaderBlockStatus:    input.QuotationsHeader.HeaderBlockStatus,
				IsCancelled:          input.QuotationsHeader.IsCancelled,
				IsMarkedForDeletion:  input.QuotationsHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Items" {
		request = CreateQuotationsRequestItems(
			requestPram,
			&apiInputReader.QuotationsItems{
				Quotation:           input.QuotationsItems.Quotation,
				IsCancelled:         input.QuotationsItems.IsCancelled,
				IsMarkedForDeletion: input.QuotationsItems.IsMarkedForDeletion,
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
