package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type SupplyChainRelationshipRes struct {
	ConnectionKey       string                         `json:"connection_key,omitempty"`
	Result              bool                           `json:"result,omitempty"`
	RedisKey            string                         `json:"redis_key,omitempty"`
	Filepath            string                         `json:"filepath,omitempty"`
	APIStatusCode       int                            `json:"api_status_code,omitempty"`
	RuntimeSessionID    string                         `json:"runtime_session_id,omitempty"`
	BusinessPartnerID   *int                           `json:"business_partner,omitempty"`
	ServiceLabel        string                         `json:"service_label,omitempty"`
	APIType             string                         `json:"api_type,omitempty"`
	Message             SupplyChainRelationshipMessage `json:"message,omitempty"`
	APISchema           string                         `json:"api_schema,omitempty"`
	Accepter            []string                       `json:"accepter,omitempty"`
	Deleted             bool                           `json:"deleted,omitempty"`
	SQLUpdateResult     *bool                          `json:"sql_update_result,omitempty"`
	SQLUpdateError      string                         `json:"sql_update_error,omitempty"`
	SubfuncResult       *bool                          `json:"subfunc_result,omitempty"`
	SubfuncError        string                         `json:"subfunc_error,omitempty"`
	ExconfResult        *bool                          `json:"exconf_result,omitempty"`
	ExconfError         string                         `json:"exconf_error,omitempty"`
	APIProcessingResult *bool                          `json:"api_processing_result,omitempty"`
	APIProcessingError  string                         `json:"api_processing_error,omitempty"`
}

type SupplyChainRelationshipMessage struct {
	General                             *[]SCRGeneral                          `json:"General"`
	Transaction                         *[]Transaction                         `json:"Transaction"`
	DeliveryRelation                    *[]DeliveryRelation                    `json:"DeliveryRelation"`
	BillingRelation                     *[]BillingRelation                     `json:"BillingRelation"`
	PaymentRelation                     *[]PaymentRelation                     `json:"PaymentRelation"`
	DeliveryPlantRelation               *[]DeliveryPlantRelation               `json:"DeliveryPlantRelation"`
	DeliveryPlantRelationProduct        *[]DeliveryPlantRelationProduct        `json:"DeliveryPlantRelationProduct"`
	DeliveryPlantRelationProductMRPArea *[]DeliveryPlantRelationProductMRPArea `json:"DeliveryPlantRelationProductMRPArea"`
	StockConfPlantRelation              *[]StockConfPlantRelation              `json:"StockConfPlantRelation"`
	StockConfPlantRelationProduct       *[]StockConfPlantRelationProduct       `json:"StockConfPlantRelationProduct"`
	ProductionPlantRelation             *[]ProductionPlantRelation             `json:"ProductionPlantRelation"`
	ProductionPlantRelationProductMRP   *[]ProductionPlantRelationProductMRP   `json:"ProductionPlantRelationProductMRP"`
}

type SCRGeneral struct {
	SupplyChainRelationshipID int     `json:"SupplyChainRelationshipID"`
	Buyer                     *int    `json:"Buyer"`
	Seller                    *int    `json:"Seller"`
	CreationDate              *string `json:"CreationDate"`
	LastChangeDate            *string `json:"LastChangeDate"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

type Transaction struct {
	SupplyChainRelationshipID int     `json:"SupplyChainRelationshipID"`
	Buyer                     int     `json:"Buyer"`
	Seller                    int     `json:"Seller"`
	TransactionCurrency       *string `json:"TransactionCurrency"`
	PaymentTerms              *string `json:"PaymentTerms"`
	PaymentMethod             *string `json:"PaymentMethod"`
	Incoterms                 *string `json:"Incoterms"`
	AccountAssignmentGroup    *string `json:"AccountAssignmentGroup"`
	CreationDate              *string `json:"CreationDate"`
	LastChangeDate            *string `json:"LastChangeDate"`
	QuotationIsBlocked        *bool   `json:"QuotationIsBlocked"`
	OrderIsBlocked            *bool   `json:"OrderIsBlocked"`
	DeliveryIsBlocked         *bool   `json:"DeliveryIsBlocked"`
	BillingIsBlocked          *bool   `json:"BillingIsBlocked"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

type DeliveryRelation struct {
	SupplyChainRelationshipID         int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID int     `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             int     `json:"Buyer"`
	BuyerName                         *string `json:"BuyerName"`
	Seller                            int     `json:"Seller"`
	SellerName                        *string `json:"SellerName"`
	DeliverToParty                    int     `json:"DeliverToParty"`
	DeliverToPartyName                *string `json:"DeliverToPartyName"`
	DeliverFromParty                  int     `json:"DeliverFromParty"`
	DeliverFromPartyName              *string `json:"DeliverFromPartyName"`
	DefaultRelation                   *bool   `json:"DefaultRelation"`
	CreationDate                      *string `json:"CreationDate"`
	LastChangeDate                    *string `json:"LastChangeDate"`
	IsMarkedForDeletion               *bool   `json:"IsMarkedForDeletion"`
}

type BillingRelation struct {
	SupplyChainRelationshipID        int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int     `json:"SupplyChainRelationshipBillingID"`
	Buyer                            int     `json:"Buyer"`
	Seller                           int     `json:"Seller"`
	BillToParty                      int     `json:"BillToParty"`
	BillFromParty                    int     `json:"BillFromParty"`
	DefaultRelation                  *bool   `json:"DefaultRelation"`
	BillToCountry                    string  `json:"BillToCountry"`
	BillFromCountry                  string  `json:"BillFromCountry"`
	IsExportImport                   *bool   `json:"IsExportImport"`
	TransactionTaxCategory           *string `json:"TransactionTaxCategory"`
	TransactionTaxClassification     *string `json:"TransactionTaxClassification"`
	CreationDate                     *string `json:"CreationDate"`
	LastChangeDate                   *string `json:"LastChangeDate"`
	IsMarkedForDeletion              *bool   `json:"IsMarkedForDeletion"`
}

type PaymentRelation struct {
	SupplyChainRelationshipID        int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int     `json:"Buyer"`
	Seller                           int     `json:"Seller"`
	BillToParty                      int     `json:"BillToParty"`
	BillFromParty                    int     `json:"BillFromParty"`
	Payer                            int     `json:"Payer"`
	Payee                            int     `json:"Payee"`
	DefaultRelation                  *bool   `json:"DefaultRelation"`
	PayerHouseBank                   *string `json:"PayerHouseBank"`
	PayerHouseBankAccount            *string `json:"PayerHouseBankAccount"`
	PayeeHouseBank                   *string `json:"PayeeHouseBank"`
	PayeeHouseBankAccount            *string `json:"PayeeHouseBankAccount"`
	CreationDate                     *string `json:"CreationDate"`
	LastChangeDate                   *string `json:"LastChangeDate"`
	IsMarkedForDeletion              *bool   `json:"IsMarkedForDeletion"`
}

type DeliveryPlantRelation struct {
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int     `json:"Buyer"`
	Seller                                 int     `json:"Seller"`
	DeliverToParty                         int     `json:"DeliverToParty"`
	DeliverToPartyName                     *string `json:"DeliverToPartyName"`
	DeliverFromParty                       int     `json:"DeliverFromParty"`
	DeliverFromPartyName                   *string `json:"DeliverFromPartyName"`
	DeliverToPlant                         string  `json:"DeliverToPlant"`
	DeliverToPlantName                     *string `json:"DeliverToPlantName"`
	DeliverFromPlant                       string  `json:"DeliverFromPlant"`
	DeliverFromPlantName                   *string `json:"DeliverFromPlantName"`
	DefaultRelation                        *bool   `json:"DefaultRelation"`
	MRPType                                *string `json:"MRPType"`
	MRPController                          *string `json:"MRPController"`
	CreationDate                           *string `json:"CreationDate"`
	LastChangeDate                         *string `json:"LastChangeDate"`
	IsMarkedForDeletion                    *bool   `json:"IsMarkedForDeletion"`
}

type DeliveryPlantRelationProduct struct {
	SupplyChainRelationshipID                 int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID         int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID    int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                     int      `json:"Buyer"`
	Seller                                    int      `json:"Seller"`
	DeliverToParty                            int      `json:"DeliverToParty"`
	DeliverFromToParty                        int      `json:"DeliverFromToParty"`
	DeliverToPlant                            string   `json:"DeliverToPlant"`
	DeliverFromPlant                          string   `json:"DeliverFromPlant"`
	Product                                   string   `json:"Product"`
	DeliverToPlantStorageLocation             string   `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlantStorageLocation           string   `json:"DeliverFromPlantStorageLocation"`
	DeliveryUnit                              string   `json:"DeliveryUnit"`
	QuantityPerPackage                        *float32 `json:"QuantityPerPackage"`
	MRPType                                   *string  `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	ReorderThresholdQuantity                  *float32 `json:"ReorderThresholdQuantity"`
	PlanningTimeFence                         *int     `json:"PlanningTimeFence"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *int     `json:"SafetyDuration"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32 `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDurationInDays            *int     `json:"StandardDeliveryDurationInDays"`
	IsAutoOrderCreationAllowed                *bool    `json:"IsAutoOrderCreationAllowed"`
	IsOrderAcknowledgementRequired            *bool    `json:"IsOrderAcknowledgementRequired"`
	CreationDate                              *string  `json:"CreationDate"`
	LastChangeDate                            *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}

type DeliveryPlantRelationProductMRPArea struct {
	SupplyChainRelationshipID                 int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID         int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID    int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                     int      `json:"Buyer"`
	Seller                                    int      `json:"Seller"`
	DeliverToParty                            int      `json:"DeliverToParty"`
	DeliverFromParty                          int      `json:"DeliverFromParty"`
	DeliverToPlant                            string   `json:"DeliverToPlant"`
	DeliverFromPlant                          string   `json:"DeliverFromPlant"`
	Product                                   string   `json:"Product"`
	DeliverToPlantStorageLocation             string   `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlantStorageLocation           string   `json:"DeliverFromPlantStorageLocation"`
	DeliveryUnit                              string   `json:"DeliveryUnit"`
	QuantityPerPackage                        *float32 `json:"QuantityPerPackage"`
	MRPType                                   *string  `json:"MRPType"`
	MRPArea                                   string   `json:"MRPArea"`
	MRPController                             *string  `json:"MRPController"`
	ReorderThresholdQuantity                  *float32 `json:"ReorderThresholdQuantity"`
	PlanningTimeFence                         *int     `json:"PlanningTimeFence"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *int     `json:"SafetyDuration"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32 `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDurationInDays            *int     `json:"StandardDeliveryDurationInDays"`
	IsAutoOrderCreationAllowed                *bool    `json:"IsAutoOrderCreationAllowed"`
	IsOrderAcknowledgementRequired            *bool    `json:"IsOrderAcknowledgementRequired"`
	CreationDate                              *string  `json:"CreationDate"`
	LastChangeDate                            *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}

type StockConfPlantRelation struct {
	SupplyChainRelationshipID               int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID int     `json:"SupplyChainRelationshipStockConfPlantID"`
	Buyer                                   int     `json:"Buyer"`
	Seller                                  int     `json:"Seller"`
	StockConfirmationBusinessPartner        int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                  string  `json:"StockConfirmationPlant"`
	DefaultRelation                         *bool   `json:"DefaultRelation"`
	MRPType                                 *string `json:"MRPType"`
	MRPController                           *string `json:"MRPController"`
	CreationDate                            *string `json:"CreationDate"`
	LastChangeDate                          *string `json:"LastChangeDate"`
	IsMarkedForDeletion                     *bool   `json:"IsMarkedForDeletion"`
}

type StockConfPlantRelationProduct struct {
	SupplyChainRelationshipID               int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID int     `json:"SupplyChainRelationshipStockConfPlantID"`
	Buyer                                   int     `json:"Buyer"`
	Seller                                  int     `json:"Seller"`
	StockConfirmationBusinessPartner        int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                  string  `json:"StockConfirmationPlant"`
	Product                                 string  `json:"Product"`
	MRPType                                 *string `json:"MRPType"`
	MRPController                           *string `json:"MRPController"`
	CreationDate                            *string `json:"CreationDate"`
	LastChangeDate                          *string `json:"LastChangeDate"`
	IsMarkedForDeletion                     *bool   `json:"IsMarkedForDeletion"`
}

type ProductionPlantRelation struct {
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID int     `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                    int     `json:"Buyer"`
	Seller                                   int     `json:"Seller"`
	ProductionPlantBusinessPartner           int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string  `json:"ProductionPlant"`
	DefaultRelation                          *bool   `json:"DefaultRelation"`
	MRPType                                  *string `json:"MRPType"`
	MRPController                            *string `json:"MRPController"`
	CreationDate                             *string `json:"CreationDate"`
	LastChangeDate                           *string `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool   `json:"IsMarkedForDeletion"`
}

type ProductionPlantRelationProductMRP struct {
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID int     `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                    int     `json:"Buyer"`
	Seller                                   int     `json:"Seller"`
	ProductionPlantBusinessPartner           int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string  `json:"ProductionPlant"`
	Product                                  string  `json:"Product"`
	ProductionPlantStorageLocation           *string `json:"ProductionPlantStorageLocation"`
	MRPType                                  *string `json:"MRPType"`
	MRPController                            *string `json:"MRPController"`
	CreationDate                             *string `json:"CreationDate"`
	LastChangeDate                           *string `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool   `json:"IsMarkedForDeletion"`
}

func CreateSupplyChainRelationshipRes(msg rabbitmq.RabbitmqMessage) (*SupplyChainRelationshipRes, error) {
	res := SupplyChainRelationshipRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
