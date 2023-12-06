package delivery_document

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type DeliveryDocumentReq struct {
	Header   Header   `json:"DeliveryDocument"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	DeliveryDocument                       int      `json:"DeliveryDocument"`
	SupplyChainRelationshipID              *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipBillingID       *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID       *int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                                  *int     `json:"Buyer"`
	Seller                                 *int     `json:"Seller"`
	DeliverToParty                         *int     `json:"DeliverToParty"`
	DeliverFromParty                       *int     `json:"DeliverFromParty"`
	DeliverToPlant                         *string  `json:"DeliverToPlant"`
	DeliverFromPlant                       *string  `json:"DeliverFromPlant"`
	BillToParty                            *int     `json:"BillToParty"`
	BillFromParty                          *int     `json:"BillFromParty"`
	BillToCountry                          *string  `json:"BillToCountry"`
	BillFromCountry                        *string  `json:"BillFromCountry"`
	Payer                                  *int     `json:"Payer"`
	Payee                                  *int     `json:"Payee"`
	IsExportImport                         *bool    `json:"IsExportImport"`
	DeliverToPlantTimeZone                 *string  `json:"DeliverToPlantTimeZone"`
	DeliverFromPlantTimeZone               *string  `json:"DeliverFromPlantTimeZone"`
	ReferenceDocument                      *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                  *int     `json:"ReferenceDocumentItem"`
	OrderID                                *int     `json:"OrderID"`
	OrderItem                              *int     `json:"OrderItem"`
	ProductionOrder                        *int     `json:"ProductionOrder"`
	ProductionOrderItem                    *int     `json:"ProductionOrderItem"`
	Operations                             *int     `json:"Operations"`
	OperationsItem                         *int     `json:"OperationsItem"`
	BillOfMaterial                         *int     `json:"BillOfMaterial"`
	BillOfMaterialItem                     *int     `json:"BillOfMaterialItem"`
	ContractType                           *string  `json:"ContractType"`
	OrderValidityStartDate                 *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate                   *string  `json:"OrderValidityEndDate"`
	DeliveryDocumentDate                   *string  `json:"DeliveryDocumentDate"`
	PlannedGoodsIssueDate                  *string  `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                  *string  `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                *string  `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                *string  `json:"PlannedGoodsReceiptTime"`
	InvoiceDocumentDate                    *string  `json:"InvoiceDocumentDate"`
	HeaderCompleteDeliveryIsDefined        *bool    `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus                   *string  `json:"HeaderDeliveryStatus"`
	GoodsIssueOrReceiptSlipNumber          *string  `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderBillingStatus                    *string  `json:"HeaderBillingStatus"`
	HeaderBillingConfStatus                *string  `json:"HeaderBillingConfStatus"`
	HeaderBillingBlockStatus               *bool    `json:"HeaderBillingBlockStatus"`
	HeaderGrossWeight                      *float32 `json:"HeaderGrossWeight"`
	HeaderNetWeight                        *float32 `json:"HeaderNetWeight"`
	HeaderWeightUnit                       *string  `json:"HeaderWeightUnit"`
	Incoterms                              *string  `json:"Incoterms"`
	TransactionCurrency                    *string  `json:"TransactionCurrency"`
	HeaderDeliveryBlockStatus              *bool    `json:"HeaderDeliveryBlockStatus"`
	HeaderIssuingBlockStatus               *bool    `json:"HeaderIssuingBlockStatus"`
	HeaderReceivingBlockStatus             *bool    `json:"HeaderReceivingBlockStatus"`
	ExternalReferenceDocument              *string  `json:"ExternalReferenceDocument"`
	CreationDate                           *string  `json:"CreationDate"`
	CreationTime                           *string  `json:"CreationTime"`
	LastChangeDate                         *string  `json:"LastChangeDate"`
	LastChangeTime                         *string  `json:"LastChangeTime"`
	IsCancelled                            *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                    *bool    `json:"IsMarkedForDeletion"`
	Item                                   []Item   `json:"Item"`
}

type Item struct {
	DeliveryDocument                              int      `json:"DeliveryDocument"`
	DeliveryDocumentItem                          int      `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemCategory                  *string  `json:"DeliveryDocumentItemCategory"`
	SupplyChainRelationshipID                     *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID"`
	SupplyChainRelationshipBillingID              *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID              *int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                                         *int     `json:"Buyer"`
	Seller                                        *int     `json:"Seller"`
	DeliverToParty                                *int     `json:"DeliverToParty"`
	DeliverFromParty                              *int     `json:"DeliverFromParty"`
	DeliverToPlant                                *string  `json:"DeliverToPlant"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant"`
	BillToParty                                   *int     `json:"BillToParty"`
	BillFromParty                                 *int     `json:"BillFromParty"`
	BillToCountry                                 *string  `json:"BillToCountry"`
	BillFromCountry                               *string  `json:"BillFromCountry"`
	Payer                                         *int     `json:"Payer"`
	Payee                                         *int     `json:"Payee"`
	Product                                       *string  `json:"Product"`
	ProductStandardID                             *string  `json:"ProductStandardID"`
	ProductGroup                                  *string  `json:"ProductGroup"`
	BaseUnit                                      *string  `json:"BaseUnit"`
	DeliveryUnit                                  *string  `json:"DeliveryUnit"`
	OriginalQuantityInBaseUnit                    *float32 `json:"OriginalQuantityInBaseUnit"`
	OriginalQuantityInDeliveryUnit                *float32 `json:"OriginalQuantityInDeliveryUnit"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation"`
	ProductIsBatchManagedInDeliverToPlant         *bool    `json:"ProductIsBatchManagedInDeliverToPlant"`
	BatchMgmtPolicyInDeliverToPlant               *string  `json:"BatchMgmtPolicyInDeliverToPlant"`
	DeliverToPlantBatch                           *string  `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate          *string  `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityStartTime          *string  `json:"DeliverToPlantBatchValidityStartTime"`
	DeliverToPlantBatchValidityEndDate            *string  `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverToPlantBatchValidityEndTime            *string  `json:"DeliverToPlantBatchValidityEndTime"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation"`
	ProductIsBatchManagedInDeliverFromPlant       *bool    `json:"ProductIsBatchManagedInDeliverFromPlant"`
	BatchMgmtPolicyInDeliverFromPlant             *string  `json:"BatchMgmtPolicyInDeliverFromPlant"`
	DeliverFromPlantBatch                         *string  `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartDate        *string  `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityStartTime        *string  `json:"DeliverFromPlantBatchValidityStartTime"`
	DeliverFromPlantBatchValidityEndDate          *string  `json:"DeliverFromPlantBatchValidityEndDate"`
	DeliverFromPlantBatchValidityEndTime          *string  `json:"DeliverFromPlantBatchValidityEndTime"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant"`
	ProductIsBatchManagedInStockConfirmationPlant *bool    `json:"ProductIsBatchManagedInStockConfirmationPlant"`
	BatchMgmtPolicyInStockConfirmationPlant       *string  `json:"BatchMgmtPolicyInStockConfirmationPlant"`
	StockConfirmationPlantBatch                   *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate  *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityStartTime  *string  `json:"StockConfirmationPlantBatchValidityStartTime"`
	StockConfirmationPlantBatchValidityEndDate    *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	StockConfirmationPlantBatchValidityEndTime    *string  `json:"StockConfirmationPlantBatchValidityEndTime"`
	StockConfirmationPolicy                       *string  `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                       *string  `json:"StockConfirmationStatus"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation                *string  `json:"ProductionPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant        *bool    `json:"ProductIsBatchManagedInProductionPlant"`
	BatchMgmtPolicyInProductionPlant              *string  `json:"BatchMgmtPolicyInProductionPlant"`
	ProductionPlantBatch                          *string  `json:"ProductionPlantBatch"`
	ProductionPlantBatchValidityStartDate         *string  `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityStartTime         *string  `json:"ProductionPlantBatchValidityStartTime"`
	ProductionPlantBatchValidityEndDate           *string  `json:"ProductionPlantBatchValidityEndDate"`
	ProductionPlantBatchValidityEndTime           *string  `json:"ProductionPlantBatchValidityEndTime"`
	InspectionPlan                                *int     `json:"InspectionPlan"`
	InspectionPlant                               *string  `json:"InspectionPlant"`
	InspectionOrder                               *int     `json:"InspectionOrder"`
	DeliveryDocumentItemText                      *string  `json:"DeliveryDocumentItemText"`
	DeliveryDocumentItemTextByBuyer               *string  `json:"DeliveryDocumentItemTextByBuyer"`
	DeliveryDocumentItemTextBySeller              *string  `json:"DeliveryDocumentItemTextBySeller"`
	PlannedGoodsIssueDate                         *string  `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                         *string  `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                       *string  `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                       *string  `json:"PlannedGoodsReceiptTime"`
	PlannedGoodsIssueQuantity                     *float32 `json:"PlannedGoodsIssueQuantity"`
	PlannedGoodsIssueQtyInBaseUnit                *float32 `json:"PlannedGoodsIssueQtyInBaseUnit"`
	PlannedGoodsReceiptQuantity                   *float32 `json:"PlannedGoodsReceiptQuantity"`
	PlannedGoodsReceiptQtyInBaseUnit              *float32 `json:"PlannedGoodsReceiptQtyInBaseUnit"`
	ActualGoodsIssueDate                          *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                          *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate                        *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                        *string  `json:"ActualGoodsReceiptTime"`
	ActualGoodsIssueQuantity                      *float32 `json:"ActualGoodsIssueQuantity"`
	ActualGoodsIssueQtyInBaseUnit                 *float32 `json:"ActualGoodsIssueQtyInBaseUnit"`
	ActualGoodsReceiptQuantity                    *float32 `json:"ActualGoodsReceiptQuantity"`
	ActualGoodsReceiptQtyInBaseUnit               *float32 `json:"ActualGoodsReceiptQtyInBaseUnit"`
	QuantityPerPackage                            *float32 `json:"QuantityPerPackage"`
	ItemBillingStatus                             *string  `json:"ItemBillingStatus"`
	ItemCompleteDeliveryIsDefined                 *bool    `json:"ItemCompleteDeliveryIsDefined"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight"`
	InternalCapacityQuantity                      *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit                  *string  `json:"InternalCapacityQuantityUnit"`
	ItemIsBillingRelevant                         *bool    `json:"ItemIsBillingRelevant"`
	NetAmount                                     *float32 `json:"NetAmount"`
	TaxAmount                                     *float32 `json:"TaxAmount"`
	GrossAmount                                   *float32 `json:"GrossAmount"`
	OrderID                                       *int     `json:"OrderID"`
	OrderItem                                     *int     `json:"OrderItem"`
	ProductionOrder                               *int     `json:"ProductionOrder"`
	ProductionOrderItem                           *int     `json:"ProductionOrderItem"`
	BillOfMaterial                                *int     `json:"BillOfMaterial"`
	BillOfMaterialItem                            *int     `json:"BillOfMaterialItem"`
	Operations                                    *int     `json:"Operations"`
	OperationsItem                                *int     `json:"OperationsItem"`
	OrderType                                     *string  `json:"OrderType"`
	ContractType                                  *string  `json:"ContractType"`
	OrderValidityStartDate                        *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate                          *string  `json:"OrderValidityEndDate"`
	PaymentTerms                                  *string  `json:"PaymentTerms"`
	DueCalculationBaseDate                        *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                                *string  `json:"PaymentDueDate"`
	NetPaymentDays                                *int     `json:"NetPaymentDays"`
	PaymentMethod                                 *string  `json:"PaymentMethod"`
	InvoicePeriodStartDate                        *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                          *string  `json:"InvoicePeriodEndDate"`
	ConfirmedDeliveryDate                         *string  `json:"ConfirmedDeliveryDate"`
	Project                                       *int     `json:"Project"`
	WBSElement                                    *int     `json:"WBSElement"`
	ReferenceDocument                             *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                         *int     `json:"ReferenceDocumentItem"`
	TransactionTaxClassification                  *string  `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry         *string  `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry       *string  `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                      *string  `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                        *string  `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup                 *string  `json:"ProductAccountAssignmentGroup"`
	TaxCode                                       *string  `json:"TaxCode"`
	TaxRate                                       *float32 `json:"TaxRate"`
	CountryOfOrigin                               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                       *string  `json:"CountryOfOriginLanguage"`
	Equipment                                     *int     `json:"Equipment"`
	ItemDeliveryBlockStatus                       *bool    `json:"ItemDeliveryBlockStatus"`
	ItemIssuingBlockStatus                        *bool    `json:"ItemIssuingBlockStatus"`
	ItemReceivingBlockStatus                      *bool    `json:"ItemReceivingBlockStatus"`
	ItemBillingBlockStatus                        *bool    `json:"ItemBillingBlockStatus"`
	ExternalReferenceDocument                     *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem                 *string  `json:"ExternalReferenceDocumentItem"`
	CreationDate                                  *string  `json:"CreationDate"`
	CreationTime                                  *string  `json:"CreationTime"`
	LastChangeDate                                *string  `json:"LastChangeDate"`
	LastChangeTime                                *string  `json:"LastChangeTime"`
	IsCancelled                                   *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                           *bool    `json:"IsMarkedForDeletion"`
}

func CreateDeliveryDocumentRequestHeader(
	requestPram *apiInputReader.Request,
	deliveryDocumentHeader *apiInputReader.DeliveryDocumentHeader,
) DeliveryDocumentReq {
	req := DeliveryDocumentReq{
		Header: Header{
			DeliverToParty:                  deliveryDocumentHeader.DeliverToParty,
			HeaderCompleteDeliveryIsDefined: deliveryDocumentHeader.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryBlockStatus:       deliveryDocumentHeader.HeaderDeliveryBlockStatus,
			HeaderDeliveryStatus:            deliveryDocumentHeader.HeaderDeliveryStatus,
			IsCancelled:                     deliveryDocumentHeader.IsCancelled,
			IsMarkedForDeletion:             deliveryDocumentHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersByDeliverToParty",
		},
	}
	return req
}

func CreateDeliveryDocumentRequestHeaderByDeliverToParty(
	requestPram *apiInputReader.Request,
	deliveryDocumentHeader *apiInputReader.DeliveryDocumentHeader,
) DeliveryDocumentReq {
	req := DeliveryDocumentReq{
		Header: Header{
			DeliverToParty:                  deliveryDocumentHeader.DeliverToParty,
			HeaderCompleteDeliveryIsDefined: deliveryDocumentHeader.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryBlockStatus:       deliveryDocumentHeader.HeaderDeliveryBlockStatus,
			HeaderDeliveryStatus:            deliveryDocumentHeader.HeaderDeliveryStatus,
			IsCancelled:                     deliveryDocumentHeader.IsCancelled,
			IsMarkedForDeletion:             deliveryDocumentHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersByDeliverToParty",
		},
	}
	return req
}

func CreateDeliveryDocumentRequestHeaderByDeliverFromParty(
	requestPram *apiInputReader.Request,
	deliveryDocumentHeader *apiInputReader.DeliveryDocumentHeader,
) DeliveryDocumentReq {
	req := DeliveryDocumentReq{
		Header: Header{
			DeliverFromParty:                deliveryDocumentHeader.DeliverFromParty,
			HeaderCompleteDeliveryIsDefined: deliveryDocumentHeader.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryBlockStatus:       deliveryDocumentHeader.HeaderDeliveryBlockStatus,
			HeaderDeliveryStatus:            deliveryDocumentHeader.HeaderDeliveryStatus,
			IsCancelled:                     deliveryDocumentHeader.IsCancelled,
			IsMarkedForDeletion:             deliveryDocumentHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersByDeliverFromParty",
		},
	}
	return req
}

func CreateDeliveryDocumentRequestItem(
	requestPram *apiInputReader.Request,
	deliveryDocumentItem *apiInputReader.DeliveryDocumentItem,
) DeliveryDocumentReq {
	req := DeliveryDocumentReq{
		Header: Header{
			DeliveryDocument: deliveryDocumentItem.DeliveryDocument,
			Item: []Item{
				{
					DeliveryDocumentItem:          deliveryDocumentItem.DeliveryDocumentItem,
					ItemCompleteDeliveryIsDefined: deliveryDocumentItem.ItemCompleteDeliveryIsDefined,
					ItemDeliveryBlockStatus:       deliveryDocumentItem.ItemDeliveryBlockStatus,
					IsCancelled:                   deliveryDocumentItem.IsCancelled,
					IsMarkedForDeletion:           deliveryDocumentItem.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Item",
		},
	}
	return req
}

func CreateDeliveryDocumentRequestItems(
	requestPram *apiInputReader.Request,
	deliveryDocumentItems *apiInputReader.DeliveryDocumentItems,
) DeliveryDocumentReq {
	req := DeliveryDocumentReq{
		Header: Header{
			DeliveryDocument: deliveryDocumentItems.DeliveryDocument,
			Item: []Item{
				{
					DeliveryDocumentItem:          *deliveryDocumentItems.DeliveryDocumentItem,
					ItemCompleteDeliveryIsDefined: deliveryDocumentItems.ItemCompleteDeliveryIsDefined,
					ItemDeliveryBlockStatus:       deliveryDocumentItems.ItemDeliveryBlockStatus,
					IsCancelled:                   deliveryDocumentItems.IsCancelled,
					IsMarkedForDeletion:           deliveryDocumentItems.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Items",
		},
	}
	return req
}

func DeliveryDocumentReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.DeliveryDocument,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_DELIVERY_DOCUMENT_SRV"
	aPIType := "reads"

	var request DeliveryDocumentReq

	if accepter == "Header" {
		request = CreateDeliveryDocumentRequestHeader(
			requestPram,
			&apiInputReader.DeliveryDocumentHeader{
				DeliveryDocument:                input.DeliveryDocumentHeader.DeliveryDocument,
				HeaderCompleteDeliveryIsDefined: input.DeliveryDocumentHeader.HeaderCompleteDeliveryIsDefined,
				HeaderDeliveryBlockStatus:       input.DeliveryDocumentHeader.HeaderDeliveryBlockStatus,
				HeaderDeliveryStatus:            input.DeliveryDocumentHeader.HeaderDeliveryStatus,
				IsCancelled:                     input.DeliveryDocumentHeader.IsCancelled,
				IsMarkedForDeletion:             input.DeliveryDocumentHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "HeadersByDeliverToParty" {
		request = CreateDeliveryDocumentRequestHeaderByDeliverToParty(
			requestPram,
			&apiInputReader.DeliveryDocumentHeader{
				DeliverToParty:                  input.DeliveryDocumentHeader.DeliverToParty,
				HeaderCompleteDeliveryIsDefined: input.DeliveryDocumentHeader.HeaderCompleteDeliveryIsDefined,
				HeaderDeliveryBlockStatus:       input.DeliveryDocumentHeader.HeaderDeliveryBlockStatus,
				HeaderDeliveryStatus:            input.DeliveryDocumentHeader.HeaderDeliveryStatus,
				IsCancelled:                     input.DeliveryDocumentHeader.IsCancelled,
				IsMarkedForDeletion:             input.DeliveryDocumentHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "HeadersByDeliverFromParty" {
		request = CreateDeliveryDocumentRequestHeaderByDeliverFromParty(
			requestPram,
			&apiInputReader.DeliveryDocumentHeader{
				DeliverFromParty:                input.DeliveryDocumentHeader.DeliverFromParty,
				HeaderCompleteDeliveryIsDefined: input.DeliveryDocumentHeader.HeaderCompleteDeliveryIsDefined,
				HeaderDeliveryBlockStatus:       input.DeliveryDocumentHeader.HeaderDeliveryBlockStatus,
				HeaderDeliveryStatus:            input.DeliveryDocumentHeader.HeaderDeliveryStatus,
				IsCancelled:                     input.DeliveryDocumentHeader.IsCancelled,
				IsMarkedForDeletion:             input.DeliveryDocumentHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Item" {
		request = CreateDeliveryDocumentRequestItem(
			requestPram,
			&apiInputReader.DeliveryDocumentItem{
				DeliveryDocument:              input.DeliveryDocumentItems.DeliveryDocument,
				DeliveryDocumentItem:          *input.DeliveryDocumentItems.DeliveryDocumentItem,
				ItemCompleteDeliveryIsDefined: input.DeliveryDocumentItems.ItemCompleteDeliveryIsDefined,
				ItemDeliveryBlockStatus:       input.DeliveryDocumentItems.ItemDeliveryBlockStatus,
				IsCancelled:                   input.DeliveryDocumentItems.IsCancelled,
				IsMarkedForDeletion:           input.DeliveryDocumentItems.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Items" {
		request = CreateDeliveryDocumentRequestItems(
			requestPram,
			&apiInputReader.DeliveryDocumentItems{
				DeliveryDocument:              input.DeliveryDocumentItems.DeliveryDocument,
				DeliveryDocumentItem:          input.DeliveryDocumentItems.DeliveryDocumentItem,
				ItemCompleteDeliveryIsDefined: input.DeliveryDocumentItems.ItemCompleteDeliveryIsDefined,
				ItemDeliveryBlockStatus:       input.DeliveryDocumentItems.ItemDeliveryBlockStatus,
				IsCancelled:                   input.DeliveryDocumentItems.IsCancelled,
				IsMarkedForDeletion:           input.DeliveryDocumentItems.IsMarkedForDeletion,
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
