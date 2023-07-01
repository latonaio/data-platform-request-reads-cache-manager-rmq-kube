package apiresponses

import (
	"encoding/json"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type SupplyChainRelationshipExconfRes struct {
	ConnectionKey                           string                                  `json:"connection_key"`
	Result                                  bool                                    `json:"result"`
	RedisKey                                string                                  `json:"redis_key"`
	Filepath                                string                                  `json:"filepath"`
	APIStatusCode                           int                                     `json:"api_status_code"`
	RuntimeSessionID                        string                                  `json:"runtime_session_id"`
	BusinessPartnerID                       *int                                    `json:"business_partner"`
	ServiceLabel                            string                                  `json:"service_label"`
	SupplyChainRelationshipGeneral          SupplyChainRelationshipGeneral          `json:"SupplyChainRelationshipGeneral"`
	SupplyChainRelationshipDeliveryRelation SupplyChainRelationshipDeliveryRelation `json:"SupplyChainRelationshipDeliveryRelation"`
	SupplyChainRelationshipDeliveryPlant    SupplyChainRelationshipDeliveryPlant    `json:"SupplyChainRelationshipDeliveryPlant"`
	SupplyChainRelationshipBilling          SupplyChainRelationshipBilling          `json:"SupplyChainRelationshipBilling"`
	SupplyChainRelationshipPayment          SupplyChainRelationshipPayment          `json:"SupplyChainRelationshipPayment"`
	SupplyChainRelationshipTransaction      SupplyChainRelationshipTransaction      `json:"SupplyChainRelationshipTransaction"`
	APISchema                               string                                  `json:"api_schema"`
	Accepter                                []string                                `json:"accepter"`
	Deleted                                 bool                                    `json:"deleted"`
}

type SupplyChainRelationshipGeneral struct {
	SupplyChainRelationshipID int  `json:"SupplyChainRelationshipID"`
	Buyer                     int  `json:"Buyer"`
	Seller                    int  `json:"Seller"`
	ExistenceConf             bool `json:"ExistenceConf"`
}

type SupplyChainRelationshipDeliveryRelation struct {
	SupplyChainRelationshipID         int  `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID int  `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             int  `json:"Buyer"`
	Seller                            int  `json:"Seller"`
	DeliverToParty                    int  `json:"DeliverToParty"`
	DeliverFromParty                  int  `json:"DeliverFromParty"`
	ExistenceConf                     bool `json:"ExistenceConf"`
}

type SupplyChainRelationshipDeliveryPlant struct {
	SupplyChainRelationshipDeliveryID      int    `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int    `json:"SupplyChainRelationshipDeliveryPlantID"`
	DeliverToParty                         string `json:"DeliverToParty"`
	DeliverFromParty                       string `json:"DeliverFromParty"`
	DeliverToPlant                         string `json:"DeliverToPlant"`
	DeliverFromPlant                       string `json:"DeliverFromPlant"`
	IsMarkedForDeletion                    bool   `json:"IsMarkedForDeletion"`
	ExistenceConf                          bool   `json:"ExistenceConf"`
}

type SupplyChainRelationshipBilling struct {
	SupplyChainRelationshipBillingID int    `json:"SupplyChainRelationshipBillingID"`
	BillToParty                      string `json:"BillToParty"`
	BillFromParty                    string `json:"BillFromParty"`
	IsMarkedForDeletion              bool   `json:"IsMarkedForDeletion"`
	ExistenceConf                    bool   `json:"ExistenceConf"`
}

type SupplyChainRelationshipPayment struct {
	SupplyChainRelationshipPaymentID int    `json:"SupplyChainRelationshipPaymentID"`
	Payer                            string `json:"Payer"`
	Payee                            string `json:"Payee"`
	IsMarkedForDeletion              bool   `json:"IsMarkedForDeletion"`
	ExistenceConf                    bool   `json:"ExistenceConf"`
}

type SupplyChainRelationshipTransaction struct {
	TransactionCurrency    string `json:"TransactionCurrency"`
	PaymentTerms           string `json:"PaymentTerms"`
	PaymentMethod          string `json:"PaymentMethod"`
	Incoterms              string `json:"Incoterms"`
	AccountAssignmentGroup int    `json:"AccountAssignmentGroup"`
	QuotationIsBlocked     bool   `json:"QuotationIsBlocked"`
	OrderIsBlocked         bool   `json:"OrderIsBlocked"`
	DeliveryIsBlocked      bool   `json:"DeliveryIsBlocked"`
	BillingIsBlocked       bool   `json:"BillingIsBlocked"`
	IsMarkedForDeletion    bool   `json:"IsMarkedForDeletion"`
	ExistenceConf          bool   `json:"ExistenceConf"`
}

type SupplyChainRelationshipGeneralExconfRes struct {
	ConnectionKey           string                        `json:"connection_key"`
	Result                  bool                          `json:"result"`
	RedisKey                string                        `json:"redis_key"`
	Filepath                string                        `json:"filepath"`
	APIStatusCode           int                           `json:"api_status_code"`
	RuntimeSessionID        string                        `json:"runtime_session_id"`
	BusinessPartnerID       *int                          `json:"business_partner"`
	ServiceLabel            string                        `json:"service_label"`
	SupplyChainRelationship SupplyChainRelationshipExconf `json:"SupplyChainRelationship"`
	APISchema               string                        `json:"api_schema"`
	Accepter                []string                      `json:"accepter"`
	Deleted                 bool                          `json:"deleted"`
}

type SupplyChainRelationshipExconf struct {
	General           *SupplyChainRelationshipExconfGeneral `json:"General"`
	Delivery          *DeliveryExconf                       `json:"Delivery"`
	DeliveryPlant     *DeliveryPlantExconf                  `json:"DeliveryPlant"`
	Billing           *BillingExconf                        `json:"Billing"`
	Payment           *PaymentExconf                        `json:"Payment"`
	TransactionExconf *TransactionExconf                    `json:"TransactionExconf"`
}

type SupplyChainRelationshipExconfGeneral struct {
	Product       string `json:"Product"`
	ExistenceConf bool   `json:"ExistenceConf"`
}

type DeliveryExconf struct {
	SupplyChainRelationshipDeliveryID int    `json:"SupplyChainRelationshipDeliveryID"`
	DeliverToParty                    string `json:"DeliverToParty"`
	DeliverFromParty                  string `json:"DeliverFromParty"`
	IsMarkedForDeletion               bool   `json:"IsMarkedForDeletion"`
	ExistenceConf                     bool   `json:"ExistenceConf"`
}

type DeliveryPlantExconf struct {
	SupplyChainRelationshipDeliveryID      int    `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int    `json:"SupplyChainRelationshipDeliveryPlantID"`
	DeliverToParty                         string `json:"DeliverToParty"`
	DeliverFromParty                       string `json:"DeliverFromParty"`
	DeliverToPlant                         string `json:"DeliverToPlant"`
	DeliverFromPlant                       string `json:"DeliverFromPlant"`
	IsMarkedForDeletion                    bool   `json:"IsMarkedForDeletion"`
	ExistenceConf                          bool   `json:"ExistenceConf"`
}

type BillingExconf struct {
	SupplyChainRelationshipBillingID int    `json:"SupplyChainRelationshipBillingID"`
	BillToParty                      string `json:"BillToParty"`
	BillFromParty                    string `json:"BillFromParty"`
	IsMarkedForDeletion              bool   `json:"IsMarkedForDeletion"`
	ExistenceConf                    bool   `json:"ExistenceConf"`
}

type PaymentExconf struct {
	SupplyChainRelationshipPaymentID int    `json:"SupplyChainRelationshipPaymentID"`
	Payer                            string `json:"Payer"`
	Payee                            string `json:"Payee"`
	IsMarkedForDeletion              bool   `json:"IsMarkedForDeletion"`
	ExistenceConf                    bool   `json:"ExistenceConf"`
}

type TransactionExconf struct {
	TransactionCurrency    string `json:"TransactionCurrency"`
	PaymentTerms           string `json:"PaymentTerms"`
	PaymentMethod          string `json:"PaymentMethod"`
	Incoterms              string `json:"Incoterms"`
	AccountAssignmentGroup int    `json:"AccountAssignmentGroup"`
	QuotationIsBlocked     bool   `json:"QuotationIsBlocked"`
	OrderIsBlocked         bool   `json:"OrderIsBlocked"`
	DeliveryIsBlocked      bool   `json:"DeliveryIsBlocked"`
	BillingIsBlocked       bool   `json:"BillingIsBlocked"`
	IsMarkedForDeletion    bool   `json:"IsMarkedForDeletion"`
	ExistenceConf          bool   `json:"ExistenceConf"`
}

func CreateSupplyChainRelationshipExconfRes(msg rabbitmq.RabbitmqMessage) (*SupplyChainRelationshipExconfRes, error) {
	res := SupplyChainRelationshipExconfRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
