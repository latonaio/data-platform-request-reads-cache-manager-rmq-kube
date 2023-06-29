package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type DeliveryDocumentRes struct {
	ConnectionKey       string                  `json:"connection_key,omitempty"`
	Result              bool                    `json:"result,omitempty"`
	RedisKey            string                  `json:"redis_key,omitempty"`
	Filepath            string                  `json:"filepath,omitempty"`
	APIStatusCode       int                     `json:"api_status_code,omitempty"`
	RuntimeSessionID    string                  `json:"runtime_session_id,omitempty"`
	BusinessPartnerID   *int                    `json:"business_partner,omitempty"`
	ServiceLabel        string                  `json:"service_label,omitempty"`
	APIType             string                  `json:"api_type,omitempty"`
	Message             DeliveryDocumentMessage `json:"message,omitempty"`
	APISchema           string                  `json:"api_schema,omitempty"`
	Accepter            []string                `json:"accepter,omitempty"`
	Deleted             bool                    `json:"deleted,omitempty"`
	SQLUpdateResult     *bool                   `json:"sql_update_result,omitempty"`
	SQLUpdateError      string                  `json:"sql_update_error,omitempty"`
	SubfuncResult       *bool                   `json:"subfunc_result,omitempty"`
	SubfuncError        string                  `json:"subfunc_error,omitempty"`
	ExconfResult        *bool                   `json:"exconf_result,omitempty"`
	ExconfError         string                  `json:"exconf_error,omitempty"`
	APIProcessingResult *bool                   `json:"api_processing_result,omitempty"`
	APIProcessingError  string                  `json:"api_processing_error,omitempty"`
}

type DeliveryDocumentMessage struct {
	Header           *[]DeliveryDocumentHeader  `json:"Header,omitempty"`
	Item             *[]DeliveryDocumentItem    `json:"Item,omitempty"`
	Partner          *[]DeliveryDocumentPartner `json:"Partner,omitempty"`
	Address          *[]DeliveryDocumentAddress `json:"Address,omitempty"`
	DeliverFromItems *[]DeliverFromItems        `json:"DeliverFromItems,omitempty"`
	DeliverToItems   *[]DeliverToItems          `json:"DeliverToItems,omitempty"`
}

type DeliveryDocumentHeader struct {
	DeliveryDocument                       int      `json:"DeliveryDocument,omitempty"`
	SupplyChainRelationshipID              *int     `json:"SupplyChainRelationshipID,omitempty"`
	SupplyChainRelationshipDeliveryID      *int     `json:"SupplyChainRelationshipDeliveryID,omitempty"`
	SupplyChainRelationshipDeliveryPlantID *int     `json:"SupplyChainRelationshipDeliveryPlantID,omitempty"`
	SupplyChainRelationshipBillingID       *int     `json:"SupplyChainRelationshipBillingID,omitempty"`
	SupplyChainRelationshipPaymentID       *int     `json:"SupplyChainRelationshipPaymentID,omitempty"`
	Buyer                                  *int     `json:"Buyer,omitempty"`
	Seller                                 *int     `json:"Seller,omitempty"`
	DeliverToParty                         *int     `json:"DeliverToParty,omitempty"`
	DeliverFromParty                       *int     `json:"DeliverFromParty,omitempty"`
	DeliverToPlant                         *string  `json:"DeliverToPlant,omitempty"`
	DeliverFromPlant                       *string  `json:"DeliverFromPlant,omitempty"`
	BillToParty                            *int     `json:"BillToParty,omitempty"`
	BillFromParty                          *int     `json:"BillFromParty,omitempty"`
	BillToCountry                          *string  `json:"BillToCountry,omitempty"`
	BillFromCountry                        *string  `json:"BillFromCountry,omitempty"`
	Payer                                  *int     `json:"Payer,omitempty"`
	Payee                                  *int     `json:"Payee,omitempty"`
	IsExportImport                         *bool    `json:"IsExportImport,omitempty"`
	DeliverToPlantTimeZone                 *string  `json:"DeliverToPlantTimeZone,omitempty"`
	DeliverFromPlantTimeZone               *string  `json:"DeliverFromPlantTimeZone,omitempty"`
	ReferenceDocument                      *int     `json:"ReferenceDocument,omitempty"`
	ReferenceDocumentItem                  *int     `json:"ReferenceDocumentItem,omitempty"`
	OrderID                                *int     `json:"OrderID,omitempty"`
	OrderItem                              *int     `json:"OrderItem,omitempty"`
	ContractType                           *string  `json:"ContractType,omitempty"`
	OrderValidityStartDate                 *string  `json:"OrderValidityStartDate,omitempty"`
	OrderValidityEndDate                   *string  `json:"OrderValidityEndDate,omitempty"`
	DocumentDate                           *string  `json:"DocumentDate,omitempty"`
	PlannedGoodsIssueDate                  *string  `json:"PlannedGoodsIssueDate,omitempty"`
	PlannedGoodsIssueTime                  *string  `json:"PlannedGoodsIssueTime,omitempty"`
	PlannedGoodsReceiptDate                *string  `json:"PlannedGoodsReceiptDate,omitempty"`
	PlannedGoodsReceiptTime                *string  `json:"PlannedGoodsReceiptTime,omitempty"`
	InvoiceDocumentDate                    *string  `json:"InvoiceDocumentDate,omitempty"`
	HeaderCompleteDeliveryIsDefined        *bool    `json:"HeaderCompleteDeliveryIsDefined,omitempty"`
	HeaderDeliveryStatus                   *string  `json:"HeaderDeliveryStatus,omitempty"`
	CreationDate                           *string  `json:"CreationDate,omitempty"`
	CreationTime                           *string  `json:"CreationTime,omitempty"`
	LastChangeDate                         *string  `json:"LastChangeDate,omitempty"`
	LastChangeTime                         *string  `json:"LastChangeTime,omitempty"`
	GoodsIssueOrReceiptSlipNumber          *string  `json:"GoodsIssueOrReceiptSlipNumber,omitempty"`
	HeaderBillingStatus                    *string  `json:"HeaderBillingStatus,omitempty"`
	HeaderBillingConfStatus                *string  `json:"HeaderBillingConfStatus,omitempty"`
	HeaderBillingBlockStatus               *bool    `json:"HeaderBillingBlockStatus,omitempty"`
	HeaderGrossWeight                      *float32 `json:"HeaderGrossWeight,omitempty"`
	HeaderNetWeight                        *float32 `json:"HeaderNetWeight,omitempty"`
	HeaderWeightUnit                       *string  `json:"HeaderWeightUnit,omitempty"`
	Incoterms                              *string  `json:"Incoterms,omitempty"`
	TransactionCurrency                    *string  `json:"TransactionCurrency,omitempty"`
	HeaderDeliveryBlockStatus              *bool    `json:"HeaderDeliveryBlockStatus,omitempty"`
	HeaderIssuingBlockStatus               *bool    `json:"HeaderIssuingBlockStatus,omitempty"`
	HeaderReceivingBlockStatus             *bool    `json:"HeaderReceivingBlockStatus,omitempty"`
	IsCancelled                            *bool    `json:"IsCancelled,omitempty"`
	IsMarkedForDeletion                    *bool    `json:"IsMarkedForDeletion,omitempty"`
}

type DeliveryDocumentHeaderDoc struct {
	DeliveryDocument         int     `json:"DeliveryDocument,omitempty"`
	DocType                  string  `json:"DocType,omitempty"`
	DocVersionID             int     `json:"DocVersionID,omitempty"`
	DocID                    string  `json:"DocID,omitempty"`
	FileExtension            *string `json:"FileExtension,omitempty"`
	FileName                 *string `json:"FileName,omitempty"`
	FilePath                 *string `json:"FilePath,omitempty"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner,omitempty"`
}

type DeliveryDocumentItem struct {
	DeliveryDocument                              int      `json:"DeliveryDocument,omitempty"`
	DeliveryDocumentItem                          int      `json:"DeliveryDocumentItem,omitempty"`
	DeliveryDocumentItemCategory                  *string  `json:"DeliveryDocumentItemCategory,omitempty"`
	SupplyChainRelationshipID                     *int     `json:"SupplyChainRelationshipID,omitempty"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID,omitempty"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID,omitempty"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID,omitempty"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID,omitempty"`
	SupplyChainRelationshipBillingID              *int     `json:"SupplyChainRelationshipBillingID,omitempty"`
	SupplyChainRelationshipPaymentID              *int     `json:"SupplyChainRelationshipPaymentID,omitempty"`
	Buyer                                         *int     `json:"Buyer,omitempty"`
	Seller                                        *int     `json:"Seller,omitempty"`
	DeliverToParty                                *int     `json:"DeliverToParty,omitempty"`
	DeliverFromParty                              *int     `json:"DeliverFromParty,omitempty"`
	DeliverToPlant                                *string  `json:"DeliverToPlant,omitempty"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant,omitempty"`
	BillToParty                                   *int     `json:"BillToParty,omitempty"`
	BillFromParty                                 *int     `json:"BillFromParty,omitempty"`
	BillToCountry                                 *string  `json:"BillToCountry,omitempty"`
	BillFromCountry                               *string  `json:"BillFromCountry,omitempty"`
	Payer                                         *int     `json:"Payer,omitempty"`
	Payee                                         *int     `json:"Payee,omitempty"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation,omitempty"`
	ProductIsBatchManagedInDeliverToPlant         *bool    `json:"ProductIsBatchManagedInDeliverToPlant,omitempty"`
	BatchMgmtPolicyInDeliverToPlant               *string  `json:"BatchMgmtPolicyInDeliverToPlant,omitempty"`
	DeliverToPlantBatch                           *string  `json:"DeliverToPlantBatch,omitempty"`
	DeliverToPlantBatchValidityStartDate          *string  `json:"DeliverToPlantBatchValidityStartDate,omitempty"`
	DeliverToPlantBatchValidityStartTime          *string  `json:"DeliverToPlantBatchValidityStartTime,omitempty"`
	DeliverToPlantBatchValidityEndDate            *string  `json:"DeliverToPlantBatchValidityEndDate,omitempty"`
	DeliverToPlantBatchValidityEndTime            *string  `json:"DeliverToPlantBatchValidityEndTime,omitempty"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation,omitempty"`
	ProductIsBatchManagedInDeliverFromPlant       *bool    `json:"ProductIsBatchManagedInDeliverFromPlant,omitempty"`
	BatchMgmtPolicyInDeliverFromPlant             *string  `json:"BatchMgmtPolicyInDeliverFromPlant,omitempty"`
	DeliverFromPlantBatch                         *string  `json:"DeliverFromPlantBatch,omitempty"`
	DeliverFromPlantBatchValidityStartDate        *string  `json:"DeliverFromPlantBatchValidityStartDate,omitempty"`
	DeliverFromPlantBatchValidityStartTime        *string  `json:"DeliverFromPlantBatchValidityStartTime,omitempty"`
	DeliverFromPlantBatchValidityEndDate          *string  `json:"DeliverFromPlantBatchValidityEndDate,omitempty"`
	DeliverFromPlantBatchValidityEndTime          *string  `json:"DeliverFromPlantBatchValidityEndTime,omitempty"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner,omitempty"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant,omitempty"`
	ProductIsBatchManagedInStockConfirmationPlant *bool    `json:"ProductIsBatchManagedInStockConfirmationPlant,omitempty"`
	BatchMgmtPolicyInStockConfirmationPlant       *string  `json:"BatchMgmtPolicyInStockConfirmationPlant,omitempty"`
	StockConfirmationPlantBatch                   *string  `json:"StockConfirmationPlantBatch,omitempty"`
	StockConfirmationPlantBatchValidityStartDate  *string  `json:"StockConfirmationPlantBatchValidityStartDate,omitempty"`
	StockConfirmationPlantBatchValidityStartTime  *string  `json:"StockConfirmationPlantBatchValidityStartTime,omitempty"`
	StockConfirmationPlantBatchValidityEndDate    *string  `json:"StockConfirmationPlantBatchValidityEndDate,omitempty"`
	StockConfirmationPlantBatchValidityEndTime    *string  `json:"StockConfirmationPlantBatchValidityEndTime,omitempty"`
	StockConfirmationPolicy                       *string  `json:"StockConfirmationPolicy,omitempty"`
	StockConfirmationStatus                       *string  `json:"StockConfirmationStatus,omitempty"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner,omitempty"`
	ProductionPlant                               *string  `json:"ProductionPlant,omitempty"`
	ProductionPlantStorageLocation                *string  `json:"ProductionPlantStorageLocation,omitempty"`
	ProductIsBatchManagedInProductionPlant        *bool    `json:"ProductIsBatchManagedInProductionPlant,omitempty"`
	BatchMgmtPolicyInProductionPlant              *string  `json:"BatchMgmtPolicyInProductionPlant,omitempty"`
	ProductionPlantBatch                          *string  `json:"ProductionPlantBatch,omitempty"`
	ProductionPlantBatchValidityStartDate         *string  `json:"ProductionPlantBatchValidityStartDate,omitempty"`
	ProductionPlantBatchValidityStartTime         *string  `json:"ProductionPlantBatchValidityStartTime,omitempty"`
	ProductionPlantBatchValidityEndDate           *string  `json:"ProductionPlantBatchValidityEndDate,omitempty"`
	ProductionPlantBatchValidityEndTime           *string  `json:"ProductionPlantBatchValidityEndTime,omitempty"`
	DeliveryDocumentItemText                      *string  `json:"DeliveryDocumentItemText,omitempty"`
	DeliveryDocumentItemTextByBuyer               *string  `json:"DeliveryDocumentItemTextByBuyer,omitempty"`
	DeliveryDocumentItemTextBySeller              *string  `json:"DeliveryDocumentItemTextBySeller,omitempty"`
	Product                                       *string  `json:"Product,omitempty"`
	ProductStandardID                             *string  `json:"ProductStandardID,omitempty"`
	ProductGroup                                  *string  `json:"ProductGroup,omitempty"`
	BaseUnit                                      *string  `json:"BaseUnit,omitempty"`
	OriginalQuantityInBaseUnit                    *float32 `json:"OriginalQuantityInBaseUnit,omitempty"`
	DeliveryUnit                                  *string  `json:"DeliveryUnit,omitempty"`
	PlannedGoodsIssueDate                         *string  `json:"PlannedGoodsIssueDate,omitempty"`
	PlannedGoodsIssueTime                         *string  `json:"PlannedGoodsIssueTime,omitempty"`
	PlannedGoodsReceiptDate                       *string  `json:"PlannedGoodsReceiptDate,omitempty"`
	PlannedGoodsReceiptTime                       *string  `json:"PlannedGoodsReceiptTime,omitempty"`
	PlannedGoodsIssueQuantity                     *float32 `json:"PlannedGoodsIssueQuantity,omitempty"`
	PlannedGoodsIssueQtyInBaseUnit                *float32 `json:"PlannedGoodsIssueQtyInBaseUnit,omitempty"`
	PlannedGoodsReceiptQuantity                   *float32 `json:"PlannedGoodsReceiptQuantity,omitempty"`
	PlannedGoodsReceiptQtyInBaseUnit              *float32 `json:"PlannedGoodsReceiptQtyInBaseUnit,omitempty"`
	ActualGoodsIssueDate                          *string  `json:"ActualGoodsIssueDate,omitempty"`
	ActualGoodsIssueTime                          *string  `json:"ActualGoodsIssueTime,omitempty"`
	ActualGoodsReceiptDate                        *string  `json:"ActualGoodsReceiptDate,omitempty"`
	ActualGoodsReceiptTime                        *string  `json:"ActualGoodsReceiptTime,omitempty"`
	ActualGoodsIssueQuantity                      *float32 `json:"ActualGoodsIssueQuantity,omitempty"`
	ActualGoodsIssueQtyInBaseUnit                 *float32 `json:"ActualGoodsIssueQtyInBaseUnit,omitempty"`
	ActualGoodsReceiptQuantity                    *float32 `json:"ActualGoodsReceiptQuantity,omitempty"`
	ActualGoodsReceiptQtyInBaseUnit               *float32 `json:"ActualGoodsReceiptQtyInBaseUnit,omitempty"`
	CreationDate                                  *string  `json:"CreationDate,omitempty"`
	CreationTime                                  *string  `json:"CreationTime,omitempty"`
	LastChangeDate                                *string  `json:"LastChangeDate,omitempty"`
	LastChangeTime                                *string  `json:"LastChangeTime,omitempty"`
	ItemBillingStatus                             *string  `json:"ItemBillingStatus,omitempty"`
	ItemCompleteDeliveryIsDefined                 *bool    `json:"ItemCompleteDeliveryIsDefined,omitempty"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight,omitempty"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight,omitempty"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit,omitempty"`
	InternalCapacityQuantity                      *float32 `json:"InternalCapacityQuantity,omitempty"`
	InternalCapacityQuantityUnit                  *string  `json:"InternalCapacityQuantityUnit,omitempty"`
	ItemIsBillingRelevant                         *bool    `json:"ItemIsBillingRelevant,omitempty"`
	NetAmount                                     *float32 `json:"NetAmount,omitempty"`
	TaxAmount                                     *float32 `json:"TaxAmount,omitempty"`
	GrossAmount                                   *float32 `json:"GrossAmount,omitempty"`
	OrderID                                       *int     `json:"OrderID,omitempty"`
	OrderItem                                     *int     `json:"OrderItem,omitempty"`
	OrderType                                     *string  `json:"OrderType,omitempty"`
	ContractType                                  *string  `json:"ContractType,omitempty"`
	OrderValidityStartDate                        *string  `json:"OrderValidityStartDate,omitempty"`
	OrderValidityEndDate                          *string  `json:"OrderValidityEndDate,omitempty"`
	PaymentTerms                                  *string  `json:"PaymentTerms,omitempty"`
	DueCalculationBaseDate                        *string  `json:"DueCalculationBaseDate,omitempty"`
	PaymentDueDate                                *string  `json:"PaymentDueDate,omitempty"`
	NetPaymentDays                                *int     `json:"NetPaymentDays,omitempty"`
	PaymentMethod                                 *string  `json:"PaymentMethod,omitempty"`
	InvoicePeriodStartDate                        *string  `json:"InvoicePeriodStartDate,omitempty"`
	InvoicePeriodEndDate                          *string  `json:"InvoicePeriodEndDate,omitempty"`
	ConfirmedDeliveryDate                         *string  `json:"ConfirmedDeliveryDate,omitempty"`
	Project                                       *string  `json:"Project,omitempty"`
	ReferenceDocument                             *int     `json:"ReferenceDocument,omitempty"`
	ReferenceDocumentItem                         *int     `json:"ReferenceDocumentItem,omitempty"`
	TransactionTaxClassification                  *string  `json:"TransactionTaxClassification,omitempty"`
	ProductTaxClassificationBillToCountry         *string  `json:"ProductTaxClassificationBillToCountry,omitempty"`
	ProductTaxClassificationBillFromCountry       *string  `json:"ProductTaxClassificationBillFromCountry,omitempty"`
	DefinedTaxClassification                      *string  `json:"DefinedTaxClassification,omitempty"`
	AccountAssignmentGroup                        *string  `json:"AccountAssignmentGroup,omitempty"`
	ProductAccountAssignmentGroup                 *string  `json:"ProductAccountAssignmentGroup,omitempty"`
	TaxCode                                       *string  `json:"TaxCode,omitempty"`
	TaxRate                                       *float32 `json:"TaxRate,omitempty"`
	CountryOfOrigin                               *string  `json:"CountryOfOrigin,omitempty"`
	CountryOfOriginLanguage                       *string  `json:"CountryOfOriginLanguage,omitempty"`
	ItemDeliveryBlockStatus                       *bool    `json:"ItemDeliveryBlockStatus,omitempty"`
	ItemIssuingBlockStatus                        *bool    `json:"ItemIssuingBlockStatus,omitempty"`
	ItemReceivingBlockStatus                      *bool    `json:"ItemReceivingBlockStatus,omitempty"`
	ItemBillingBlockStatus                        *bool    `json:"ItemBillingBlockStatus,omitempty"`
	IsCancelled                                   *bool    `json:"IsCancelled,omitempty"`
	IsMarkedForDeletion                           *bool    `json:"IsMarkedForDeletion,omitempty"`
	ProductionOrder                               *int     `json:"ProductionOrder,omitempty"`
	ProductionOrderItem                           *int     `json:"ProductionOrderItem,omitempty"`
	Operations                                    *int     `json:"Operations,omitempty"`
	OperationsItem                                *int     `json:"OperationsItem,omitempty"`
	BillOfMaterial                                *int     `json:"BillOfMaterial,omitempty"`
	BillOfMaterialItem                            *int     `json:"BillOfMaterialItem,omitempty"`
}

type DeliveryDocumentAddress struct {
	DeliveryDocument int     `json:"DeliveryDocument,omitempty"`
	AddressID        int     `json:"AddressID,omitempty"`
	PostalCode       *string `json:"PostalCode,omitempty"`
	LocalRegion      *string `json:"LocalRegion,omitempty"`
	Country          *string `json:"Country,omitempty"`
	District         *string `json:"District,omitempty"`
	StreetName       *string `json:"StreetName,omitempty"`
	CityName         *string `json:"CityName,omitempty"`
	Building         *string `json:"Building,omitempty"`
	Floor            *int    `json:"Floor,omitempty"`
	Room             *int    `json:"Room,omitempty"`
}

type DeliveryDocumentPartner struct {
	DeliveryDocument        int     `json:"DeliveryDocument,omitempty"`
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

type DeliverFromItems struct {
	DeliveryDocument                   int     `json:"DeliveryDocument,omitempty"`
	DeliverFromBusinessPartnerFullName *string `json:"DeliverFromBusinessPartnerFullName,omitempty"`
	DeliverFromBusinessPartnerName     *string `json:"DeliverFromBusinessPartnerName,omitempty"`
	DeliverToBusinessPartnerName       *string `json:"DeliverToBusinessPartnerName,omitempty"`
	DeliverToBusinessPartnerFullName   *string `json:"DeliverToBusinessPartnerFullName,omitempty"`
	HeaderDeliveryStatus               *string `json:"HeaderDeliveryStatus,omitempty"`
	ItemBillingStatus                  *string `json:"ItemBillingStatus,omitempty"`
	ConfirmedDeliveryDate              *string `json:"ConfirmedDeliveryDate,omitempty"`
}

type DeliverToItems struct {
	DeliveryDocument                   int     `json:"DeliveryDocument,omitempty"`
	HeaderDeliveryStatus               *string `json:"HeaderDeliveryStatus,omitempty"`
	DeliverToBusinessPartnerFullName   *string `json:"DeliverToBusinessPartnerFullName,omitempty"`
	DeliverToBusinessPartnerName       *string `json:"DeliverToBusinessPartnerName,omitempty"`
	DeliverFromBusinessPartnerFullName *string `json:"DeliverFromBusinessPartnerFullName,omitempty"`
	DeliverFromBusinessPartnerName     *string `json:"DeliverFromBusinessPartnerName,omitempty"`
	ItemBillingStatus                  *string `json:"ItemBillingStatus,omitempty"`
	ConfirmedDeliveryDate              *string `json:"ConfirmedDeliveryDate,omitempty"`
}

type HeadersByDeliverFromParty struct {
	DeliveryDocument                        int     `json:"DeliveryDocument,omitempty"`
	HeaderDeliveryStatus                    *string `json:"HeaderDeliveryStatus,omitempty"`
	DeliverFromPartyBusinessPartnerFullName *string `json:"DeliverFromPartyBusinessPartnerFullName,omitempty"`
}

type HeadersByDeliverToParty struct {
	DeliveryDocument                      int     `json:"DeliveryDocument,omitempty"`
	HeaderDeliveryStatus                  *string `json:"HeaderDeliveryStatus,omitempty"`
	DeliverToPartyBusinessPartnerFullName *string `json:"DeliverToPartyBusinessPartnerFullName,omitempty"`
}

type DeliveryDocumentItems struct {
	DeliveryDocument         int      `json:"DeliveryDocument,omitempty"`
	DeliveryDocumentItem     int      `json:"DeliveryDocumentItem,omitempty"`
	DeliveryDocumentItemText *string  `json:"DeliveryDocumentItemText,omitempty"`
	NetAmount                *float32 `json:"NetAmount,omitempty"`
	Product                  *string  `json:"Product,omitempty"`
	ProductDescription       *string  `json:"ProductDescription,omitempty"`
	ConditionRateValue       *float32 `json:"ConditionRateValue,omitempty"`
	ConfirmedDeliveryDate    *string  `json:"ConfirmedDeliveryDate,omitempty"`
}

func CreateDeliveryDocumentRes(msg rabbitmq.RabbitmqMessage) (*DeliveryDocumentRes, error) {
	res := DeliveryDocumentRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
