package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type InvoiceRes struct {
	ConnectionKey       string   `json:"connection_key,omitempty"`
	Result              bool     `json:"result,omitempty"`
	RedisKey            string   `json:"redis_key,omitempty"`
	Filepath            string   `json:"filepath,omitempty"`
	APIStatusCode       int      `json:"api_status_code,omitempty"`
	RuntimeSessionID    string   `json:"runtime_session_id,omitempty"`
	BusinessPartnerID   *int     `json:"business_partner,omitempty"`
	ServiceLabel        string   `json:"service_label,omitempty"`
	APIType             string   `json:"api_type,omitempty"`
	Message             Invoice  `json:"message,omitempty"`
	APISchema           string   `json:"api_schema,omitempty"`
	Accepter            []string `json:"accepter,omitempty"`
	Deleted             bool     `json:"deleted,omitempty"`
	SQLUpdateResult     *bool    `json:"sql_update_result,omitempty"`
	SQLUpdateError      string   `json:"sql_update_error,omitempty"`
	SubfuncResult       *bool    `json:"subfunc_result,omitempty"`
	SubfuncError        string   `json:"subfunc_error,omitempty"`
	ExconfResult        *bool    `json:"exconf_result,omitempty"`
	ExconfError         string   `json:"exconf_error,omitempty"`
	APIProcessingResult *bool    `json:"api_processing_result,omitempty"`
	APIProcessingError  string   `json:"api_processing_error,omitempty"`
}

type Invoice struct {
	Header             *[]InvoiceHeader             `json:"Header,omitempty"`
	Item               *[]InvoiceItem               `json:"Item,omitempty"`
	ItemPricingElement *[]InvoiceItemPricingElement `json:"ItemPricingElement,omitempty"`
	Address            *[]InvoiceAddress            `json:"Address,omitempty"`
	Partner            *[]InvoicePartner            `json:"Partner,omitempty"`
	HeaderDoc          *[]InvoiceHeaderDoc          `json:"HeaderDoc,omitempty"`
}

type InvoiceHeader struct {
	InvoiceDocument                   int      `json:"InvoiceDocument"`
	CreationDate                      *string  `json:"CreationDate"`
	CreationTime                      *string  `json:"CreationTime"`
	LastChangeDate                    *string  `json:"LastChangeDate"`
	LastChangeTime                    *string  `json:"LastChangeTime"`
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
	IsCancelled                       *bool    `json:"IsCancelled"`
}

type InvoiceHeaderDoc struct {
	InvoiceDocument          int     `json:"InvoiceDocument,omitempty"`
	DocType                  string  `json:"DocType,omitempty"`
	DocVersionID             int     `json:"DocVersionID,omitempty"`
	DocID                    string  `json:"DocID,omitempty"`
	FileExtension            *string `json:"FileExtension,omitempty"`
	FileName                 *string `json:"FileName,omitempty"`
	FilePath                 *string `json:"FilePath,omitempty"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner,omitempty"`
}

type InvoicePartner struct {
	InvoiceDocument         int     `json:"InvoiceDocument,omitempty"`
	PartnerFunction         string  `json:"PartnerFunction,omitempty"`
	BusinessPartner         int     `json:"BusinessPartner,omitempty"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName,omitempty"`
	BusinessPartnerName     *string `json:"BusinessPartnerName,omitempty"`
	Organization            *string `json:"Organization,omitempty"`
	Country                 *string `json:"Country,omitempty"`
	Language                *string `json:"Language,omitempty"`
	Currency                *string `json:"Currency,omitempty"`
	ExternalDocumentID      *string `json:"ExternalDocumentID,omitempty"`
	AddressID               *int    `json:"AddressID,omitempty"`
}

type InvoiceItem struct {
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
	CreationDate                            *string  `json:"CreationDate"`
	CreationTime                            *string  `json:"CreationTime"`
	LastChangeDate                          *string  `json:"LastChangeDate"`
	LastChangeTime                          *string  `json:"LastChangeTime"`
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
	Project                                 *string  `json:"Project"`
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
	ItemPaymentRequisitionIsCreated         *bool    `json:"ItemPaymentRequisitionIsCreated"`
	ItemIsCleared                           *bool    `json:"ItemIsCleared"`
	ItemPaymentBlockStatus                  *bool    `json:"ItemPaymentBlockStatus"`
	IsCancelled                             *bool    `json:"IsCancelled"`
}

type InvoiceAddress struct {
	InvoiceDocument int     `json:"InvoiceDocument,omitempty"`
	AddressID       int     `json:"AddressID,omitempty"`
	PostalCode      *string `json:"PostalCode,omitempty"`
	LocalRegion     *string `json:"LocalRegion,omitempty"`
	Country         *string `json:"Country,omitempty"`
	District        *string `json:"District,omitempty"`
	StreetName      *string `json:"StreetName,omitempty"`
	CityName        *string `json:"CityName,omitempty"`
	Building        *string `json:"Building,omitempty"`
	Floor           *int    `json:"Floor,omitempty"`
	Room            *int    `json:"Room,omitempty"`
}

type InvoiceItemPricingElement struct {
	InvoiceDocument            int      `json:"InvoiceDocument,omitempty"`
	InvoiceDocumentItem        int      `json:"InvoiceDocumentItem,omitempty"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter,omitempty"`
	ConditionRecord            *int     `json:"ConditionRecord,omitempty"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber,omitempty"`
	ConditionType              *string  `json:"ConditionType,omitempty"`
	PricingDate                *string  `json:"PricingDate,omitempty"`
	ConditionRateValue         *float32 `json:"ConditionRateValue,omitempty"`
	ConditionCurrency          *string  `json:"ConditionCurrency,omitempty"`
	ConditionQuantity          *float32 `json:"ConditionQuantity,omitempty"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit,omitempty"`
	TaxCode                    *string  `json:"TaxCode,omitempty"`
	ConditionAmount            *float32 `json:"ConditionAmount,omitempty"`
	TransactionCurrency        *string  `json:"TransactionCurrency,omitempty"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged,omitempty"`
}

func CreateInvoiceRes(msg rabbitmq.RabbitmqMessage) (*InvoiceRes, error) {
	res := InvoiceRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
