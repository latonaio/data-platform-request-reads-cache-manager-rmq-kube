package apiModuleRuntimesResponsesQuotations

type QuotationsRes struct {
	Message Quotations `json:"message,omitempty"`
}

type Quotations struct {
	Header *[]Header `json:"Header,omitempty"`
	//	Item   *[]Item   `json:"Item,omitempty"`
}

type Header struct {
	Quotation                        int      `json:"Quotation"`
	QuotationDate                    string   `json:"QuotationDate"`
	QuotationType                    string   `json:"QuotationType"`
	QuotationStatus                  string   `json:"QuotationStatus"`
	SupplyChainRelationshipID        int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID *int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int      `json:"Buyer"`
	Seller                           int      `json:"Seller"`
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
	TotalNetAmount                   float32  `json:"TotalNetAmount"`
	TotalTaxAmount                   float32  `json:"TotalTaxAmount"`
	TotalGrossAmount                 float32  `json:"TotalGrossAmount"`
	HeaderOrderIsDefined             *bool    `json:"HeaderOrderIsDefined"`
	TransactionCurrency              string   `json:"TransactionCurrency"`
	PricingDate                      string   `json:"PricingDate"`
	PriceDetnExchangeRate            *string  `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate            string   `json:"RequestedDeliveryDate"`
	OrderProbabilityInPercent        *float32 `json:"OrderProbabilityInPercent"`
	ExpectedOrderNetAmount           *float32 `json:"ExpectedOrderNetAmount"`
	Incoterms                        *string  `json:"Incoterms"`
	PaymentTerms                     string   `json:"PaymentTerms"`
	PaymentMethod                    string   `json:"PaymentMethod"`
	ReferenceDocument                *int     `json:"ReferenceDocument"`
	AccountAssignmentGroup           string   `json:"AccountAssignmentGroup"`
	AccountingExchangeRate           *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate              string   `json:"InvoiceDocumentDate"`
	IsExportImport                   *bool    `json:"IsExportImport"`
	HeaderText                       *bool    `json:"HeaderText"`
	HeaderIsClosed                   *bool    `json:"HeaderIsClosed"`
	HeaderBlockStatus                *bool    `json:"HeaderBlockStatus"`
	ExternalReferenceDocument        *string  `json:"ExternalReferenceDocument"`
	CreationDate                     string   `json:"CreationDate"`
	LastChangeDate                   string   `json:"LastChangeDate"`
	IsCancelled                      *bool    `json:"IsCancelled"`
	IsMarkedForDeletion              *bool    `json:"IsMarkedForDeletion"`
}
