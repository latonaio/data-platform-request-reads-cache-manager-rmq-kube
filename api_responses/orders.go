package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type OrdersRes struct {
	ConnectionKey       string         `json:"connection_key,omitempty"`
	Result              bool           `json:"result,omitempty"`
	RedisKey            string         `json:"redis_key,omitempty"`
	Filepath            string         `json:"filepath,omitempty"`
	APIStatusCode       int            `json:"api_status_code,omitempty"`
	RuntimeSessionID    string         `json:"runtime_session_id,omitempty"`
	BusinessPartnerID   *int           `json:"business_partner,omitempty"`
	ServiceLabel        string         `json:"service_label,omitempty"`
	APIType             string         `json:"api_type,omitempty"`
	Message             *OrdersMessage `json:"message,omitempty"`
	APISchema           string         `json:"api_schema,omitempty"`
	Accepter            []string       `json:"accepter,omitempty"`
	Deleted             bool           `json:"deleted,omitempty"`
	SQLUpdateResult     *bool          `json:"sql_update_result,omitempty"`
	SQLUpdateError      string         `json:"sql_update_error,omitempty"`
	SubfuncResult       *bool          `json:"subfunc_result,omitempty"`
	SubfuncError        string         `json:"subfunc_error,omitempty"`
	ExconfResult        *bool          `json:"exconf_result,omitempty"`
	ExconfError         string         `json:"exconf_error,omitempty"`
	APIProcessingResult *bool          `json:"api_processing_result,omitempty"`
	APIProcessingError  string         `json:"api_processing_error,omitempty"`
}

type OrdersMessage struct {
	Header             *[]OrdersHeader             `json:"Header,omitempty"`
	Item               *[]OrdersItem               `json:"Item,omitempty"`
	ItemPricingElement *[]OrdersItemPricingElement `json:"ItemPricingElement,omitempty"`
	ItemScheduleLine   *[]ItemScheduleLine         `json:"ItemScheduleLine,omitempty"`
	Address            *[]OrdersAddress            `json:"Address,omitempty"`
	Partner            *[]OrdersPartner            `json:"Partner,omitempty"`
	HeaderDoc          *[]OrdersHeaderDoc          `json:"HeaderDoc,omitempty"`
	HeadersBySeller    *[]HeadersBySeller          `json:"HeadersBySeller,omitempty"`
	HeadersByBuyer     *[]HeadersByBuyer           `json:"HeadersByBuyer,omitempty"`
}
type OrdersHeader struct {
	OrderID                          int      `json:"OrderID,omitempty"`
	OrderDate                        *string  `json:"OrderDate,omitempty"`
	OrderType                        *string  `json:"OrderType,omitempty"`
	SupplyChainRelationshipID        *int     `json:"SupplyChainRelationshipID,omitempty"`
	SupplyChainRelationshipBillingID *int     `json:"SupplyChainRelationshipBillingID,omitempty"`
	SupplyChainRelationshipPaymentID *int     `json:"SupplyChainRelationshipPaymentID,omitempty"`
	Buyer                            *int     `json:"Buyer,omitempty"`
	Seller                           *int     `json:"Seller,omitempty"`
	BillToParty                      *int     `json:"BillToParty,omitempty"`
	BillFromParty                    *int     `json:"BillFromParty,omitempty"`
	BillToCountry                    *string  `json:"BillToCountry,omitempty"`
	BillFromCountry                  *string  `json:"BillFromCountry,omitempty"`
	Payer                            *int     `json:"Payer,omitempty"`
	Payee                            *int     `json:"Payee,omitempty"`
	CreationDate                     *string  `json:"CreationDate,omitempty"`
	LastChangeDate                   *string  `json:"LastChangeDate,omitempty"`
	ContractType                     *string  `json:"ContractType,omitempty"`
	OrderValidityStartDate           *string  `json:"OrderValidityStartDate,omitempty"`
	OrderValidityEndDate             *string  `json:"OrderValidityEndDate,omitempty"`
	InvoicePeriodStartDate           *string  `json:"InvoicePeriodStartDate,omitempty"`
	InvoicePeriodEndDate             *string  `json:"InvoicePeriodEndDate,omitempty"`
	TotalNetAmount                   *float32 `json:"TotalNetAmount,omitempty"`
	TotalTaxAmount                   *float32 `json:"TotalTaxAmount,omitempty"`
	TotalGrossAmount                 *float32 `json:"TotalGrossAmount,omitempty"`
	HeaderDeliveryStatus             *string  `json:"HeaderDeliveryStatus,omitempty"`
	HeaderBillingStatus              *string  `json:"HeaderBillingStatus,omitempty"`
	HeaderDocReferenceStatus         *string  `json:"HeaderDocReferenceStatus,omitempty"`
	TransactionCurrency              *string  `json:"TransactionCurrency,omitempty"`
	PricingDate                      *string  `json:"PricingDate,omitempty"`
	PriceDetnExchangeRate            *float32 `json:"PriceDetnExchangeRate,omitempty"`
	RequestedDeliveryDate            *string  `json:"RequestedDeliveryDate,omitempty"`
	HeaderCompleteDeliveryIsDefined  *bool    `json:"HeaderCompleteDeliveryIsDefined,omitempty"`
	Incoterms                        *string  `json:"Incoterms,omitempty"`
	PaymentTerms                     *string  `json:"PaymentTerms"`
	PaymentTermsName                 *string  `json:"PaymentTermsName"`
	PaymentMethod                    *string  `json:"PaymentMethod"`
	PaymentMethodName                *string  `json:"PaymentMethodName"`
	ReferenceDocument                *int     `json:"ReferenceDocument,omitempty"`
	ReferenceDocumentItem            *int     `json:"ReferenceDocumentItem,omitempty"`
	AccountAssignmentGroup           *string  `json:"AccountAssignmentGroup,omitempty"`
	AccountingExchangeRate           *float32 `json:"AccountingExchangeRate,omitempty"`
	InvoiceDocumentDate              *string  `json:"InvoiceDocumentDate,omitempty"`
	IsExportImport                   *bool    `json:"IsExportImport,omitempty"`
	HeaderText                       *string  `json:"HeaderText,omitempty"`
	HeaderBlockStatus                *bool    `json:"HeaderBlockStatus,omitempty"`
	HeaderDeliveryBlockStatus        *bool    `json:"HeaderDeliveryBlockStatus,omitempty"`
	HeaderBillingBlockStatus         *bool    `json:"HeaderBillingBlockStatus,omitempty"`
	IsCancelled                      bool     `json:"IsCancelled,omitempty"`
	IsMarkedForDeletion              bool     `json:"IsMarkedForDeletion,omitempty"`
}

type OrdersPartner struct {
	OrderID                 int     `json:"OrderID,omitempty"`
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

type OrdersAddress struct {
	OrderID     int     `json:"OrderID,omitempty"`
	AddressID   int     `json:"AddressID,omitempty"`
	PostalCode  *string `json:"PostalCode,omitempty"`
	LocalRegion *string `json:"LocalRegion,omitempty"`
	Country     *string `json:"Country,omitempty"`
	District    *string `json:"District,omitempty"`
	StreetName  *string `json:"StreetName,omitempty"`
	CityName    *string `json:"CityName,omitempty"`
	Building    *string `json:"Building,omitempty"`
	Floor       *int    `json:"Floor,omitempty"`
	Room        *int    `json:"Room,omitempty"`
}

type HeadersBySeller struct {
	OrderID                          int     `json:"OrderID,omitempty"`
	HeaderDeliveryStatus             *string `json:"HeaderDeliveryStatus,omitempty"`
	DeliverToBusinessPartnerFullName *string `json:"DeliverToBusinessPartnerFullName,omitempty"`
	SellerBusinessPartnerFullName    *string `json:"SellerBusinessPartnerFullName,omitempty"`
}

type HeadersByBuyer struct {
	OrderID                          int     `json:"OrderID,omitempty"`
	HeaderDeliveryStatus             *string `json:"HeaderDeliveryStatus,omitempty"`
	DeliverToBusinessPartnerFullName *string `json:"DeliverToBusinessPartnerFullName,omitempty"`
	BuyerBusinessPartnerFullName     *string `json:"BuyerBusinessPartnerFullName,omitempty"`
}

type OrdersHeaderDoc struct {
	OrderID                  int     `json:"OrderID,omitempty"`
	DocType                  string  `json:"DocType,omitempty"`
	DocVersionID             int     `json:"DocVersionID,omitempty"`
	DocID                    string  `json:"DocID,omitempty"`
	FileExtension            *string `json:"FileExtension,omitempty"`
	FileName                 *string `json:"FileName,omitempty"`
	FilePath                 *string `json:"FilePath,omitempty"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner,omitempty"`
}

type OrdersItem struct {
	OrderID                                       int      `json:"OrderID,omitempty"`
	OrderItem                                     int      `json:"OrderItem,omitempty"`
	OrderItemCategory                             *string  `json:"OrderItemCategory,omitempty"`
	SupplyChainRelationshipID                     *int     `json:"SupplyChainRelationshipID,omitempty"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID,omitempty"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID,omitempty"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID,omitempty"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID,omitempty"`
	OrderItemText                                 *string  `json:"OrderItemText,omitempty"`
	OrderItemTextByBuyer                          *string  `json:"OrderItemTextByBuyer,omitempty"`
	OrderItemTextBySeller                         *string  `json:"OrderItemTextBySeller,omitempty"`
	Product                                       *string  `json:"Product,omitempty"`
	ProductStandardID                             *string  `json:"ProductStandardID,omitempty"`
	ProductGroup                                  *string  `json:"ProductGroup,omitempty"`
	BaseUnit                                      *string  `json:"BaseUnit,omitempty"`
	PricingDate                                   *string  `json:"PricingDate,omitempty"`
	PriceDetnExchangeRate                         *float32 `json:"PriceDetnExchangeRate,omitempty"`
	RequestedDeliveryDate                         *string  `json:"RequestedDeliveryDate,omitempty"`
	RequestedDeliveryTime                         *string  `json:"RequestedDeliveryTime,omitempty"`
	DeliverToParty                                *int     `json:"DeliverToParty,omitempty"`
	DeliverFromParty                              *int     `json:"DeliverFromParty,omitempty"`
	CreationDate                                  *string  `json:"CreationDate,omitempty"`
	CreationTime                                  *string  `json:"CreationTime,omitempty"`
	LastChangeDate                                *string  `json:"LastChangeDate,omitempty"`
	LastChangeTime                                *string  `json:"LastChangeTime,omitempty"`
	DeliverToPlant                                *string  `json:"DeliverToPlant,omitempty"`
	DeliverToPlantTimeZone                        *string  `json:"DeliverToPlantTimeZone,omitempty"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation,omitempty"`
	ProductIsBatchManagedInDeliverToPlant         *bool    `json:"ProductIsBatchManagedInDeliverToPlant,omitempty"`
	BatchMgmtPolicyInDeliverToPlant               *string  `json:"BatchMgmtPolicyInDeliverToPlant,omitempty"`
	DeliverToPlantBatch                           *string  `json:"DeliverToPlantBatch,omitempty"`
	DeliverToPlantBatchValidityStartDate          *string  `json:"DeliverToPlantBatchValidityStartDate,omitempty"`
	DeliverToPlantBatchValidityStartTime          *string  `json:"DeliverToPlantBatchValidityStartTime,omitempty"`
	DeliverToPlantBatchValidityEndDate            *string  `json:"DeliverToPlantBatchValidityEndDate,omitempty"`
	DeliverToPlantBatchValidityEndTime            *string  `json:"DeliverToPlantBatchValidityEndTime,omitempty"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant,omitempty"`
	DeliverFromPlantTimeZone                      *string  `json:"DeliverFromPlantTimeZone,omitempty"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation,omitempty"`
	ProductIsBatchManagedInDeliverFromPlant       *bool    `json:"ProductIsBatchManagedInDeliverFromPlant,omitempty"`
	BatchMgmtPolicyInDeliverFromPlant             *string  `json:"BatchMgmtPolicyInDeliverFromPlant,omitempty"`
	DeliverFromPlantBatch                         *string  `json:"DeliverFromPlantBatch,omitempty"`
	DeliverFromPlantBatchValidityStartDate        *string  `json:"DeliverFromPlantBatchValidityStartDate,omitempty"`
	DeliverFromPlantBatchValidityStartTime        *string  `json:"DeliverFromPlantBatchValidityStartTime,omitempty"`
	DeliverFromPlantBatchValidityEndDate          *string  `json:"DeliverFromPlantBatchValidityEndDate,omitempty"`
	DeliverFromPlantBatchValidityEndTime          *string  `json:"DeliverFromPlantBatchValidityEndTime,omitempty"`
	DeliveryUnit                                  *string  `json:"DeliveryUnit,omitempty"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner,omitempty"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant,omitempty"`
	StockConfirmationPlantTimeZone                *string  `json:"StockConfirmationPlantTimeZone,omitempty"`
	ProductIsBatchManagedInStockConfirmationPlant *bool    `json:"ProductIsBatchManagedInStockConfirmationPlant,omitempty"`
	BatchMgmtPolicyInStockConfirmationPlant       *string  `json:"BatchMgmtPolicyInStockConfirmationPlant,omitempty"`
	StockConfirmationPlantBatch                   *string  `json:"StockConfirmationPlantBatch,omitempty"`
	StockConfirmationPlantBatchValidityStartDate  *string  `json:"StockConfirmationPlantBatchValidityStartDate,omitempty"`
	StockConfirmationPlantBatchValidityStartTime  *string  `json:"StockConfirmationPlantBatchValidityStartTime,omitempty"`
	StockConfirmationPlantBatchValidityEndDate    *string  `json:"StockConfirmationPlantBatchValidityEndDate,omitempty"`
	StockConfirmationPlantBatchValidityEndTime    *string  `json:"StockConfirmationPlantBatchValidityEndTime,omitempty"`
	ServicesRenderingDate                         *string  `json:"ServicesRenderingDate,omitempty"`
	OrderQuantityInBaseUnit                       *float32 `json:"OrderQuantityInBaseUnit,omitempty"`
	OrderQuantityInDeliveryUnit                   *float32 `json:"OrderQuantityInDeliveryUnit,omitempty"`
	StockConfirmationPolicy                       *string  `json:"StockConfirmationPolicy,omitempty"`
	StockConfirmationStatus                       *string  `json:"StockConfirmationStatus,omitempty"`
	ConfirmedOrderQuantityInBaseUnit              *float32 `json:"ConfirmedOrderQuantityInBaseUnit,omitempty"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit,omitempty"`
	ProductGrossWeight                            *float32 `json:"ProductGrossWeight,omitempty"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight,omitempty"`
	ProductNetWeight                              *float32 `json:"ProductNetWeight,omitempty"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight,omitempty"`
	InternalCapacityQuantity                      *float32 `json:"InternalCapacityQuantity,omitempty"`
	InternalCapacityQuantityUnit                  *string  `json:"InternalCapacityQuantityUnit,omitempty"`
	NetAmount                                     *float32 `json:"NetAmount,omitempty"`
	TaxAmount                                     *float32 `json:"TaxAmount,omitempty"`
	GrossAmount                                   *float32 `json:"GrossAmount,omitempty"`
	InvoiceDocumentDate                           *string  `json:"InvoiceDocumentDate,omitempty"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner,omitempty"`
	ProductionPlant                               *string  `json:"ProductionPlant,omitempty"`
	ProductionPlantTimeZone                       *string  `json:"ProductionPlantTimeZone,omitempty"`
	ProductionPlantStorageLocation                *string  `json:"ProductionPlantStorageLocation,omitempty"`
	ProductIsBatchManagedInProductionPlant        *bool    `json:"ProductIsBatchManagedInProductionPlant,omitempty"`
	BatchMgmtPolicyInProductionPlant              *string  `json:"BatchMgmtPolicyInProductionPlant,omitempty"`
	ProductionPlantBatch                          *string  `json:"ProductionPlantBatch,omitempty"`
	ProductionPlantBatchValidityStartDate         *string  `json:"ProductionPlantBatchValidityStartDate,omitempty"`
	ProductionPlantBatchValidityStartTime         *string  `json:"ProductionPlantBatchValidityStartTime,omitempty"`
	ProductionPlantBatchValidityEndDate           *string  `json:"ProductionPlantBatchValidityEndDate,omitempty"`
	ProductionPlantBatchValidityEndTime           *string  `json:"ProductionPlantBatchValidityEndTime,omitempty"`
	InspectionPlan                                *int     `json:"InspectionPlan,omitempty"`
	InspectionPlant                               *string  `json:"InspectionPlant,omitempty"`
	InspectionOrder                               *int     `json:"InspectionOrder,omitempty"`
	Incoterms                                     *string  `json:"Incoterms,omitempty"`
	TransactionTaxClassification                  *string  `json:"TransactionTaxClassification,omitempty"`
	ProductTaxClassificationBillToCountry         *string  `json:"ProductTaxClassificationBillToCountry,omitempty"`
	ProductTaxClassificationBillFromCountry       *string  `json:"ProductTaxClassificationBillFromCountry,omitempty"`
	DefinedTaxClassification                      *string  `json:"DefinedTaxClassification,omitempty"`
	AccountAssignmentGroup                        *string  `json:"AccountAssignmentGroup,omitempty"`
	ProductAccountAssignmentGroup                 *string  `json:"ProductAccountAssignmentGroup,omitempty"`
	PaymentTerms                                  *string  `json:"PaymentTerms,omitempty"`
	DueCalculationBaseDate                        *string  `json:"DueCalculationBaseDate,omitempty"`
	PaymentDueDate                                *string  `json:"PaymentDueDate,omitempty"`
	NetPaymentDays                                *int     `json:"NetPaymentDays,omitempty"`
	PaymentMethod                                 *string  `json:"PaymentMethod,omitempty"`
	Project                                       *string  `json:"Project,omitempty"`
	AccountingExchangeRate                        *float32 `json:"AccountingExchangeRate,omitempty"`
	ReferenceDocument                             *int     `json:"ReferenceDocument,omitempty"`
	ReferenceDocumentItem                         *int     `json:"ReferenceDocumentItem,omitempty"`
	ItemCompleteDeliveryIsDefined                 *bool    `json:"ItemCompleteDeliveryIsDefined,omitempty"`
	ItemDeliveryStatus                            *string  `json:"ItemDeliveryStatus,omitempty"`
	IssuingStatus                                 *string  `json:"IssuingStatus,omitempty"`
	ReceivingStatus                               *string  `json:"ReceivingStatus,omitempty"`
	ItemBillingStatus                             *string  `json:"ItemBillingStatus,omitempty"`
	TaxCode                                       *string  `json:"TaxCode,omitempty"`
	TaxRate                                       *float32 `json:"TaxRate,omitempty"`
	CountryOfOrigin                               *string  `json:"CountryOfOrigin,omitempty"`
	CountryOfOriginLanguage                       *string  `json:"CountryOfOriginLanguage,omitempty"`
	ItemBlockStatus                               *bool    `json:"ItemBlockStatus,omitempty"`
	ItemDeliveryBlockStatus                       *bool    `json:"ItemDeliveryBlockStatus,omitempty"`
	ItemBillingBlockStatus                        *bool    `json:"ItemBillingBlockStatus,omitempty"`
	IsCancelled                                   *bool    `json:"IsCancelled,omitempty"`
	IsMarkedForDeletion                           *bool    `json:"IsMarkedForDeletion,omitempty"`
}

type OrdersItemPricingElement struct {
	OrderID                    int      `json:"OrderID,omitempty"`
	OrderItem                  int      `json:"OrderItem,omitempty"`
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID,omitempty"`
	Buyer                      int      `json:"Buyer,omitempty"`
	Seller                     int      `json:"Seller,omitempty"`
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

type ItemScheduleLine struct {
	OrderID                                         int      `json:"OrderID,omitempty"`
	OrderItem                                       int      `json:"OrderItem,omitempty"`
	ScheduleLine                                    int      `json:"ScheduleLine,omitempty"`
	SupplyChainRelationshipID                       *int     `json:"SupplyChainRelationshipID,omitempty"`
	SupplyChainRelationshipStockConfPlantID         *int     `json:"SupplyChainRelationshipStockConfPlantID,omitempty"`
	Product                                         *string  `json:"Product,omitempty"`
	StockConfirmationBussinessPartner               *int     `json:"StockConfirmationBussinessPartner,omitempty"`
	StockConfirmationPlant                          *string  `json:"StockConfirmationPlant,omitempty"`
	StockConfirmationPlantTimeZone                  *string  `json:"StockConfirmationPlantTimeZone,omitempty"`
	StockConfirmationPlantBatch                     *string  `json:"StockConfirmationPlantBatch,omitempty"`
	StockConfirmationPlantBatchValidityStartDate    *string  `json:"StockConfirmationPlantBatchValidityStartDate,omitempty"`
	StockConfirmationPlantBatchValidityEndDate      *string  `json:"StockConfirmationPlantBatchValidityEndDate,omitempty"`
	RequestedDeliveryDate                           *string  `json:"RequestedDeliveryDate,omitempty"`
	RequestedDeliveryTime                           *string  `json:"RequestedDeliveryTime,omitempty"`
	ConfirmedDeliveryDate                           *string  `json:"ConfirmedDeliveryDate,omitempty"`
	ScheduleLineOrderQuantity                       *float32 `json:"ScheduleLineOrderQuantity,omitempty"`
	OriginalOrderQuantityInBaseUnit                 *float32 `json:"OriginalOrderQuantityInBaseUnit,omitempty"`
	ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit *float32 `json:"ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit,omitempty"`
	ConfirmedOrderQuantityByPDTAvailCheck           *float32 `json:"ConfirmedOrderQuantityByPDTAvailCheck,omitempty"`
	DeliveredQuantityInBaseUnit                     *float32 `json:"DeliveredQuantityInBaseUnit,omitempty"`
	UndeliveredQuantityInBaseUnit                   *float32 `json:"UndeliveredQuantityInBaseUnit,omitempty"`
	OpenConfirmedQuantityInBaseUnit                 *float32 `json:"OpenConfirmedQuantityInBaseUnit,omitempty"`
	StockIsFullyConfirmed                           *bool    `json:"StockIsFullyConfirmed,omitempty"`
	PlusMinusFlag                                   *string  `json:"PlusMinusFlag,omitempty"`
	ItemScheduleLineDeliveryBlockStatus             *bool    `json:"ItemScheduleLineDeliveryBlockStatus,omitempty"`
	IsCancelled                                     *bool    `json:"IsCancelled,omitempty"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion,omitempty"`
}

type OrdersItems struct {
	OrderID                     int      `json:"OrderID,omitempty"`
	OrderItem                   int      `json:"OrderItem,omitempty"`
	OrderItemText               *string  `json:"OrderItemText,omitempty"`
	OrderQuantityInDeliveryUnit *float32 `json:"OrderQuantityInDeliveryUnit,omitempty"`
	NetAmount                   *float32 `json:"NetAmount,omitempty"`
	Product                     *string  `json:"Product,omitempty"`
	ProductDescription          *string  `json:"ProductDescription,omitempty"`
	ConditionRateValue          *float32 `json:"ConditionRateValue,omitempty"`
	ConfirmedDeliveryDate       *string  `json:"ConfirmedDeliveryDate,omitempty"`
}

func CreateOrdersRes(msg rabbitmq.RabbitmqMessage) (*OrdersRes, error) {
	res := OrdersRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
