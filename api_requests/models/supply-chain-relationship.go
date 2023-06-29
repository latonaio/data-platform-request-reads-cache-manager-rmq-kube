package models

type SupplyChainRelationshipGeneralReq struct {
	ConnectionKey                  string                         `json:"connection_key"`
	Result                         bool                           `json:"result"`
	RedisKey                       string                         `json:"redis_key"`
	Filepath                       string                         `json:"filepath"`
	APIStatusCode                  int                            `json:"api_status_code"`
	RuntimeSessionID               string                         `json:"runtime_session_id"`
	BusinessPartnerID              *int                           `json:"business_partner"`
	ServiceLabel                   string                         `json:"service_label"`
	APIType                        string                         `json:"api_type"`
	SupplyChainRelationshipGeneral SupplyChainRelationshipGeneral `json:"SupplyChainRelationshipGeneral"`
	APISchema                      string                         `json:"api_schema"`
	Accepter                       []string                       `json:"accepter"`
	Deleted                        bool                           `json:"deleted"`
}

type SupplyChainRelationshipGeneral struct {
	SupplyChainRelationshipID *int `json:"SupplyChainRelationshipID"`
	Buyer                     *int `json:"Buyer"`
	Seller                    *int `json:"Seller"`
}

type SupplyChainRelationshipReq struct {
	ConnectionKey     string     `json:"connection_key"`
	Result            bool       `json:"result"`
	RedisKey          string     `json:"redis_key"`
	Filepath          string     `json:"filepath"`
	APIStatusCode     int        `json:"api_status_code"`
	RuntimeSessionID  string     `json:"runtime_session_id"`
	BusinessPartnerID *int       `json:"business_partner"`
	ServiceLabel      string     `json:"service_label"`
	APIType           string     `json:"api_type"`
	General           SCRGeneral `json:"SupplyChainRelationship"`
	APISchema         string     `json:"api_schema"`
	Accepter          []string   `json:"accepter"`
	Deleted           bool       `json:"deleted"`
}

type SCRGeneral struct {
	SupplyChainRelationshipID int                       `json:"SupplyChainRelationshipID"`
	Buyer                     *int                      `json:"Buyer"`
	Seller                    *int                      `json:"Seller"`
	CreationDate              *string                   `json:"CreationDate"`
	LastChangeDate            *string                   `json:"LastChangeDate"`
	IsMarkedForDeletion       *bool                     `json:"IsMarkedForDeletion"`
	Transaction               []Transaction             `json:"Transaction"`
	DeliveryRelation          []DeliveryRelation        `json:"DeliveryRelation"`
	BillingRelation           []BillingRelation         `json:"BillingRelation"`
	ProductionPlantRelation   []ProductionPlantRelation `json:"ProductionPlantRelation"`
	StockConfPlantRelation    []StockConfPlantRelation  `json:"StockConfPlantRelation"`
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
	SupplyChainRelationshipID         int                     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID int                     `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             int                     `json:"Buyer"`
	Seller                            int                     `json:"Seller"`
	DeliverToParty                    int                     `json:"DeliverToParty"`
	DeliverFromParty                  int                     `json:"DeliverFromParty"`
	DefaultRelation                   *bool                   `json:"DefaultRelation"`
	CreationDate                      *string                 `json:"CreationDate"`
	LastChangeDate                    *string                 `json:"LastChangeDate"`
	IsMarkedForDeletion               *bool                   `json:"IsMarkedForDeletion"`
	DeliveryPlantRelation             []DeliveryPlantRelation `json:"DeliveryPlantRelation"`
}

type BillingRelation struct {
	SupplyChainRelationshipID        int               `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int               `json:"SupplyChainRelationshipBillingID"`
	Buyer                            int               `json:"Buyer"`
	Seller                           int               `json:"Seller"`
	BillToParty                      int               `json:"BillToParty"`
	BillFromParty                    int               `json:"BillFromParty"`
	DefaultRelation                  *bool             `json:"DefaultRelation"`
	BillToCountry                    *string           `json:"BillToCountry"`
	BillFromCountry                  *string           `json:"BillFromCountry"`
	IsExportImport                   *bool             `json:"IsExportImport"`
	TransactionTaxCategory           *string           `json:"TransactionTaxCategory"`
	TransactionTaxClassification     *string           `json:"TransactionTaxClassification"`
	CreationDate                     *string           `json:"CreationDate"`
	LastChangeDate                   *string           `json:"LastChangeDate"`
	IsMarkedForDeletion              *bool             `json:"IsMarkedForDeletion"`
	PaymentRelation                  []PaymentRelation `json:"PaymentRelation"`
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
	SupplyChainRelationshipID              int                            `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int                            `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int                            `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int                            `json:"Buyer"`
	Seller                                 int                            `json:"Seller"`
	DeliverToParty                         int                            `json:"DeliverToParty"`
	DeliverFromParty                       int                            `json:"DeliverFromParty"`
	DeliverToPlant                         string                         `json:"DeliverToPlant"`
	DeliverFromPlant                       string                         `json:"DeliverFromPlant"`
	DefaultRelation                        *bool                          `json:"DefaultRelation"`
	MRPType                                *string                        `json:"MRPType"`
	MRPController                          *string                        `json:"MRPController"`
	CreationDate                           *string                        `json:"CreationDate"`
	LastChangeDate                         *string                        `json:"LastChangeDate"`
	IsMarkedForDeletion                    *bool                          `json:"IsMarkedForDeletion"`
	DeliveryPlantRelationProduct           []DeliveryPlantRelationProduct `json:"DeliveryPlantRelationProduct"`
}

type DeliveryPlantRelationProduct struct {
	SupplyChainRelationshipID                 int                                   `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID         int                                   `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID    int                                   `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                     int                                   `json:"Buyer"`
	Seller                                    int                                   `json:"Seller"`
	DeliverToParty                            int                                   `json:"DeliverToParty"`
	DeliverFromToParty                        int                                   `json:"DeliverFromToParty"`
	DeliverToPlant                            string                                `json:"DeliverToPlant"`
	DeliverFromPlant                          string                                `json:"DeliverFromPlant"`
	Product                                   string                                `json:"Product"`
	DeliverToPlantStorageLocation             *string                               `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlantStorageLocation           *string                               `json:"DeliverFromPlantStorageLocation"`
	DeliveryUnit                              *string                               `json:"DeliveryUnit"`
	QuantityPerPackage                        *float32                              `json:"QuantityPerPackage"`
	MRPType                                   *string                               `json:"MRPType"`
	MRPController                             *string                               `json:"MRPController"`
	ReorderThresholdQuantity                  *float32                              `json:"ReorderThresholdQuantity"`
	PlanningTimeFence                         *int                                  `json:"PlanningTimeFence"`
	MRPPlanningCalendar                       *string                               `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32                              `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *int                                  `json:"SafetyDuration"`
	MaximumStockQuantityInBaseUnit            *float32                              `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32                              `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32                              `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32                              `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32                              `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32                              `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32                              `json:"MaximumDeliveryQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool                                 `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDurationInDays            *int                                  `json:"StandardDeliveryDurationInDays"`
	IsAutoOrderCreationAllowed                *bool                                 `json:"IsAutoOrderCreationAllowed"`
	IsOrderAcknowledgementRequired            *bool                                 `json:"IsOrderAcknowledgementRequired"`
	CreationDate                              *string                               `json:"CreationDate"`
	LastChangeDate                            *string                               `json:"LastChangeDate"`
	IsMarkedForDeletion                       *bool                                 `json:"IsMarkedForDeletion"`
	DeliveryPlantRelationProductMRPArea       []DeliveryPlantRelationProductMRPArea `json:"DeliveryPlantRelationProductMRPArea"`
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
	DeliverToPlantStorageLocation             *string  `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlantStorageLocation           *string  `json:"DeliverFromPlantStorageLocation"`
	DeliveryUnit                              *string  `json:"DeliveryUnit"`
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
	SupplyChainRelationshipID               int                             `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID int                             `json:"SupplyChainRelationshipStockConfPlantID"`
	Buyer                                   int                             `json:"Buyer"`
	Seller                                  int                             `json:"Seller"`
	StockConfirmationBusinessPartner        int                             `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                  string                          `json:"StockConfirmationPlant"`
	DefaultRelation                         *bool                           `json:"DefaultRelation"`
	MRPType                                 *string                         `json:"MRPType"`
	MRPController                           *string                         `json:"MRPController"`
	CreationDate                            *string                         `json:"CreationDate"`
	LastChangeDate                          *string                         `json:"LastChangeDate"`
	IsMarkedForDeletion                     *bool                           `json:"IsMarkedForDeletion"`
	StockConfPlantRelationProduct           []StockConfPlantRelationProduct `json:"StockConfPlantRelationProduct"`
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
	SupplyChainRelationshipID                int                                 `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID int                                 `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                    int                                 `json:"Buyer"`
	Seller                                   int                                 `json:"Seller"`
	ProductionPlantBusinessPartner           int                                 `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string                              `json:"ProductionPlant"`
	DefaultRelation                          *bool                               `json:"DefaultRelation"`
	MRPType                                  *string                             `json:"MRPType"`
	MRPController                            *string                             `json:"MRPController"`
	CreationDate                             *string                             `json:"CreationDate"`
	LastChangeDate                           *string                             `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool                               `json:"IsMarkedForDeletion"`
	ProductionPlantRelationProductMRP        []ProductionPlantRelationProductMRP `json:"ProductionPlantRelationProductMRP"`
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
