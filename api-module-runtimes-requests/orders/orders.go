package apiModuleRuntimesRequestsOrders

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type OrdersReq struct {
	Header   Header   `json:"Orders"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	OrderID                          int      `json:"OrderID"`
	OrderDate                        *string  `json:"OrderDate"`
	OrderType                        *string  `json:"OrderType"`
	SupplyChainRelationshipID        *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID *int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            *int     `json:"Buyer"`
	Seller                           *int     `json:"Seller"`
	BillToParty                      *int     `json:"BillToParty"`
	BillFromParty                    *int     `json:"BillFromParty"`
	BillToCountry                    *string  `json:"BillToCountry"`
	BillFromCountry                  *string  `json:"BillFromCountry"`
	Payer                            *int     `json:"Payer"`
	Payee                            *int     `json:"Payee"`
	ContractType                     *string  `json:"ContractType"`
	OrderValidityStartDate           *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate             *string  `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate           *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate             *string  `json:"InvoicePeriodEndDate"`
	TotalNetAmount                   *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                   *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                 *float32 `json:"TotalGrossAmount"`
	HeaderDeliveryStatus             *string  `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus              *string  `json:"HeaderBillingStatus"`
	HeaderDocReferenceStatus         *string  `json:"HeaderDocReferenceStatus"`
	TransactionCurrency              *string  `json:"TransactionCurrency"`
	PricingDate                      *string  `json:"PricingDate"`
	PriceDetnExchangeRate            *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate            *string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime            *string  `json:"RequestedDeliveryTime"`
	HeaderCompleteDeliveryIsDefined  *bool    `json:"HeaderCompleteDeliveryIsDefined"`
	Incoterms                        *string  `json:"Incoterms"`
	PaymentTerms                     *string  `json:"PaymentTerms"`
	PaymentMethod                    *string  `json:"PaymentMethod"`
	ReferenceDocument                *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem            *int     `json:"ReferenceDocumentItem"`
	AccountAssignmentGroup           *string  `json:"AccountAssignmentGroup"`
	AccountingExchangeRate           *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate              *string  `json:"InvoiceDocumentDate"`
	IsExportImport                   *bool    `json:"IsExportImport"`
	HeaderText                       *string  `json:"HeaderText"`
	HeaderBlockStatus                *bool    `json:"HeaderBlockStatus"`
	HeaderDeliveryBlockStatus        *bool    `json:"HeaderDeliveryBlockStatus"`
	HeaderBillingBlockStatus         *bool    `json:"HeaderBillingBlockStatus"`
	CreationDate                     *string  `json:"CreationDate"`
	CreationTime                     *string  `json:"CreationTime"`
	LastChangeDate                   *string  `json:"LastChangeDate"`
	LastChangeTime                   *string  `json:"LastChangeTime"`
	IsCancelled                      *bool    `json:"IsCancelled"`
	IsMarkedForDeletion              *bool    `json:"IsMarkedForDeletion"`
	Item                             []Item   `json:"Item"`
}

type Item struct {
	OrderID                                       int      `json:"OrderID"`
	OrderItem                                     int      `json:"OrderItem"`
	OrderItemCategory                             *string  `json:"OrderItemCategory"`
	SupplyChainRelationshipID                     *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID"`
	OrderItemText                                 *string  `json:"OrderItemText"`
	OrderItemTextByBuyer                          *string  `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller                         *string  `json:"OrderItemTextBySeller"`
	Product                                       *string  `json:"Product"`
	ProductStandardID                             *string  `json:"ProductStandardID"`
	ProductGroup                                  *string  `json:"ProductGroup"`
	BaseUnit                                      *string  `json:"BaseUnit"`
	BillOfMaterial                                *int     `json:"BillOfMaterial"`
	BillOfMaterialItem                            *int     `json:"BillOfMaterialItem"`
	PricingDate                                   *string  `json:"PricingDate"`
	PriceDetnExchangeRate                         *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate                         *string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                         *string  `json:"RequestedDeliveryTime"`
	DeliverToParty                                *int     `json:"DeliverToParty"`
	DeliverFromParty                              *int     `json:"DeliverFromParty"`
	DeliverToPlant                                *string  `json:"DeliverToPlant"`
	DeliverToPlantTimeZone                        *string  `json:"DeliverToPlantTimeZone"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation"`
	ProductIsBatchManagedInDeliverToPlant         *bool    `json:"ProductIsBatchManagedInDeliverToPlant"`
	BatchMgmtPolicyInDeliverToPlant               *string  `json:"BatchMgmtPolicyInDeliverToPlant"`
	DeliverToPlantBatch                           *string  `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate          *string  `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityStartTime          *string  `json:"DeliverToPlantBatchValidityStartTime"`
	DeliverToPlantBatchValidityEndDate            *string  `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverToPlantBatchValidityEndTime            *string  `json:"DeliverToPlantBatchValidityEndTime"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant"`
	DeliverFromPlantTimeZone                      *string  `json:"DeliverFromPlantTimeZone"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation"`
	ProductIsBatchManagedInDeliverFromPlant       *bool    `json:"ProductIsBatchManagedInDeliverFromPlant"`
	BatchMgmtPolicyInDeliverFromPlant             *string  `json:"BatchMgmtPolicyInDeliverFromPlant"`
	DeliverFromPlantBatch                         *string  `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartDate        *string  `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityStartTime        *string  `json:"DeliverFromPlantBatchValidityStartTime"`
	DeliverFromPlantBatchValidityEndDate          *string  `json:"DeliverFromPlantBatchValidityEndDate"`
	DeliverFromPlantBatchValidityEndTime          *string  `json:"DeliverFromPlantBatchValidityEndTime"`
	DeliveryUnit                                  *string  `json:"DeliveryUnit"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone                *string  `json:"StockConfirmationPlantTimeZone"`
	ProductIsBatchManagedInStockConfirmationPlant *bool    `json:"ProductIsBatchManagedInStockConfirmationPlant"`
	BatchMgmtPolicyInStockConfirmationPlant       *string  `json:"BatchMgmtPolicyInStockConfirmationPlant"`
	StockConfirmationPlantBatch                   *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate  *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityStartTime  *string  `json:"StockConfirmationPlantBatchValidityStartTime"`
	StockConfirmationPlantBatchValidityEndDate    *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	StockConfirmationPlantBatchValidityEndTime    *string  `json:"StockConfirmationPlantBatchValidityEndTime"`
	ServicesRenderingDate                         *string  `json:"ServicesRenderingDate"`
	OrderQuantityInBaseUnit                       *float32 `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit                   *float32 `json:"OrderQuantityInDeliveryUnit"`
	QuantityPerPackage                            *float32 `json:"QuantityPerPackage"`
	StockConfirmationPolicy                       *string  `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                       *string  `json:"StockConfirmationStatus"`
	ConfirmedOrderQuantityInBaseUnit              *float32 `json:"ConfirmedOrderQuantityInBaseUnit"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit"`
	ProductGrossWeight                            *float32 `json:"ProductGrossWeight"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight"`
	ProductNetWeight                              *float32 `json:"ProductNetWeight"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight"`
	InternalCapacityQuantity                      *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit                  *string  `json:"InternalCapacityQuantityUnit"`
	NetAmount                                     *float32 `json:"NetAmount"`
	TaxAmount                                     *float32 `json:"TaxAmount"`
	GrossAmount                                   *float32 `json:"GrossAmount"`
	InvoiceDocumentDate                           *string  `json:"InvoiceDocumentDate"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string  `json:"ProductionPlant"`
	ProductionPlantTimeZone                       *string  `json:"ProductionPlantTimeZone"`
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
	Incoterms                                     *string  `json:"Incoterms"`
	TransactionTaxClassification                  *string  `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry         *string  `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry       *string  `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                      *string  `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                        *string  `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup                 *string  `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                                  *string  `json:"PaymentTerms"`
	DueCalculationBaseDate                        *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                                *string  `json:"PaymentDueDate"`
	NetPaymentDays                                *int     `json:"NetPaymentDays"`
	PaymentMethod                                 *string  `json:"PaymentMethod"`
	Project                                       *int     `json:"Project"`
	WBSElement                                    *int     `json:"WBSElement"`
	AccountingExchangeRate                        *float32 `json:"AccountingExchangeRate"`
	ReferenceDocument                             *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                         *int     `json:"ReferenceDocumentItem"`
	ItemCompleteDeliveryIsDefined                 *bool    `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus                            *string  `json:"ItemDeliveryStatus"`
	IssuingStatus                                 *string  `json:"IssuingStatus"`
	ReceivingStatus                               *string  `json:"ReceivingStatus"`
	ItemBillingStatus                             *string  `json:"ItemBillingStatus"`
	TaxCode                                       *string  `json:"TaxCode"`
	TaxRate                                       *float32 `json:"TaxRate"`
	CountryOfOrigin                               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                       *string  `json:"CountryOfOriginLanguage"`
	Equipment                                     *int     `json:"Equipment"`
	PlannedFreight                                *int     `json:"PlannedFreight"`
	FreightOrder                                  *int     `json:"FreightOrder"`
	ItemBlockStatus                               *bool    `json:"ItemBlockStatus"`
	ItemDeliveryBlockStatus                       *bool    `json:"ItemDeliveryBlockStatus"`
	ItemBillingBlockStatus                        *bool    `json:"ItemBillingBlockStatus"`
	CreationDate                                  *string  `json:"CreationDate"`
	CreationTime                                  *string  `json:"CreationTime"`
	LastChangeDate                                *string  `json:"LastChangeDate"`
	LastChangeTime                                *string  `json:"LastChangeTime"`
	IsCancelled                                   *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                           *bool    `json:"IsMarkedForDeletion"`
}

func CreateOrdersRequestHeaderByBuyer(
	requestPram *apiInputReader.Request,
	ordersHeader *apiInputReader.OrdersHeader,
) OrdersReq {
	req := OrdersReq{
		Header: Header{
			Buyer:                           ordersHeader.Buyer,
			HeaderCompleteDeliveryIsDefined: ordersHeader.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryBlockStatus:       ordersHeader.HeaderDeliveryBlockStatus,
			HeaderDeliveryStatus:            ordersHeader.HeaderDeliveryStatus,
			IsCancelled:                     ordersHeader.IsCancelled,
			IsMarkedForDeletion:             ordersHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersByBuyer",
		},
	}
	return req
}

func CreateOrdersRequestHeaderBySeller(
	requestPram *apiInputReader.Request,
	ordersHeader *apiInputReader.OrdersHeader,
) OrdersReq {
	req := OrdersReq{
		Header: Header{
			Seller:                          ordersHeader.Seller,
			HeaderCompleteDeliveryIsDefined: ordersHeader.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryBlockStatus:       ordersHeader.HeaderDeliveryBlockStatus,
			HeaderDeliveryStatus:            ordersHeader.HeaderDeliveryStatus,
			IsCancelled:                     ordersHeader.IsCancelled,
			IsMarkedForDeletion:             ordersHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersBySeller",
		},
	}
	return req
}

func CreateOrdersRequestItems(
	requestPram *apiInputReader.Request,
	ordersItems *apiInputReader.OrdersItems,
) OrdersReq {
	req := OrdersReq{
		Header: Header{
			OrderID: ordersItems.OrderID,
			Item: []Item{
				{
					ItemCompleteDeliveryIsDefined: ordersItems.ItemCompleteDeliveryIsDefined,
					ItemDeliveryBlockStatus:       ordersItems.ItemDeliveryBlockStatus,
					ItemDeliveryStatus:            ordersItems.ItemDeliveryStatus,
					IsCancelled:                   ordersItems.IsCancelled,
					IsMarkedForDeletion:           ordersItems.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Items",
		},
	}
	return req
}

func OrdersReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Orders,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_ORDERS_SRV"
	aPIType := "reads"

	var request OrdersReq

	if accepter == "HeadersByBuyer" {
		request = CreateOrdersRequestHeaderByBuyer(
			requestPram,
			&apiInputReader.OrdersHeader{
				Buyer:                           input.OrdersHeader.Buyer,
				HeaderCompleteDeliveryIsDefined: input.OrdersHeader.HeaderCompleteDeliveryIsDefined,
				HeaderDeliveryBlockStatus:       input.OrdersHeader.HeaderDeliveryBlockStatus,
				HeaderDeliveryStatus:            input.OrdersHeader.HeaderDeliveryStatus,
				IsCancelled:                     input.OrdersHeader.IsCancelled,
				IsMarkedForDeletion:             input.OrdersHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "HeadersBySeller" {
		request = CreateOrdersRequestHeaderBySeller(
			requestPram,
			&apiInputReader.OrdersHeader{
				Seller:                          input.OrdersHeader.Seller,
				HeaderCompleteDeliveryIsDefined: input.OrdersHeader.HeaderCompleteDeliveryIsDefined,
				HeaderDeliveryBlockStatus:       input.OrdersHeader.HeaderDeliveryBlockStatus,
				HeaderDeliveryStatus:            input.OrdersHeader.HeaderDeliveryStatus,
				IsCancelled:                     input.OrdersHeader.IsCancelled,
				IsMarkedForDeletion:             input.OrdersHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Items" {
		request = CreateOrdersRequestItems(
			requestPram,
			&apiInputReader.OrdersItems{
				OrderID:                       input.OrdersItems.OrderID,
				ItemCompleteDeliveryIsDefined: input.OrdersItems.ItemCompleteDeliveryIsDefined,
				ItemDeliveryBlockStatus:       input.OrdersItems.ItemDeliveryBlockStatus,
				ItemDeliveryStatus:            input.OrdersItems.ItemDeliveryStatus,
				IsCancelled:                   input.OrdersItems.IsCancelled,
				IsMarkedForDeletion:           input.OrdersItems.IsMarkedForDeletion,
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
