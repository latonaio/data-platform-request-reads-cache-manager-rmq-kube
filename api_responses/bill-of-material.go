package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type BillOfMaterialRes struct {
	ConnectionKey       string         `json:"connection_key,omitempty"`
	Result              bool           `json:"result,omitempty"`
	RedisKey            string         `json:"redis_key,omitempty"`
	Filepath            string         `json:"filepath,omitempty"`
	APIStatusCode       int            `json:"api_status_code,omitempty"`
	RuntimeSessionID    string         `json:"runtime_session_id,omitempty"`
	BusinessPartnerID   *int           `json:"business_partner,omitempty"`
	ServiceLabel        string         `json:"service_label,omitempty"`
	APIType             string         `json:"api_type,omitempty"`
	Message             BillOfMaterial `json:"message,omitempty"`
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

type BillOfMaterial struct {
	Header *[]BillOfMaterialHeader `json:"Header,omitempty"`
	Item   *[]BillOfMaterialItem   `json:"Item,omitempty"`
}

type BillOfMaterialHeader struct {
	BillOfMaterial                           int     `json:"BillOfMaterial"`
	BillOfMaterialType                       string  `json:"BillOfMaterialType"`
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID int     `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  string  `json:"Product"`
	Buyer                                    int     `json:"Buyer"`
	Seller                                   int     `json:"Seller"`
	DepartureDeliverFromParty                int     `json:"DepartureDeliverFromParty"`
	DepartureDeliverFromPlant                string  `json:"DepartureDeliverFromPlant"`
	DestinationDeliverToParty                int     `json:"DestinationDeliverToParty"`
	DestinationDeliverToPlant                string  `json:"DestinationDeliverToPlant"`
	OwnerProductionPlantBusinessPartner      int     `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlant                     string  `json:"OwnerProductionPlant"`
	ProductBaseUnit                          string  `json:"ProductBaseUnit"`
	ProductDeliveryUnit                      string  `json:"ProductDeliveryUnit"`
	ProductProductionUnit                    string  `json:"ProductProductionUnit"`
	ProductStandardQuantityInBaseUnit        float32 `json:"ProductStandardQuantityInBaseUnit"`
	ProductStandardQuantityInDeliveryUnit    float32 `json:"ProductStandardQuantityInDeliveryUnit"`
	ProductStandardQuantityInProductionUnit  float32 `json:"ProductStandardQuantityInProductionUnit"`
	BillOfMaterialHeaderText                 *string `json:"BillOfMaterialHeaderText"`
	ValidityStartDate                        *string `json:"ValidityStartDate"`
	ValidityEndDate                          *string `json:"ValidityEndDate"`
	CreationDate                             string  `json:"CreationDate"`
	LastChangeDate                           string  `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool   `json:"IsMarkedForDeletion"`
}

type BillOfMaterialItem struct {
	BillOfMaterial                                  int      `json:"BillOfMaterial"`
	BillOfMaterialItem                              int      `json:"BillOfMaterialItem"`
	SupplyChainRelationshipID                       int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID               int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID          int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID         int      `json:"SupplyChainRelationshipStockConfPlantID"`
	Product                                         *string  `json:"Product"`
	ProductionPlantBusinessPartner                  *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                                 *string  `json:"ProductionPlant"`
	ComponentProduct                                *string  `json:"ComponentProduct"`
	ComponentProductBuyer                           *int     `json:"ComponentProductBuyer"`
	ComponentProductSeller                          *int     `json:"ComponentProductSeller"`
	ComponentProductDeliverFromParty                *int     `json:"ComponentProductDeliverFromParty"`
	ComponentProductDeliverFromPlant                *string  `json:"ComponentProductDeliverFromPlant"`
	ComponentProductDeliverToParty                  *int     `json:"ComponentProductDeliverToParty"`
	ComponentProductDeliverToPlant                  *string  `json:"ComponentProductDeliverToPlant"`
	StockConfirmationBusinessPartner                *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                          *string  `json:"StockConfirmationPlant"`
	ComponentProductStandardQuantityInBaseUnuit     *float32 `json:"ComponentProductStandardQuantityInBaseUnuit"`
	ComponentProductStandardQuantityInDeliveryUnuit *float32 `json:"ComponentProductStandardQuantityInDeliveryUnuit"`
	ComponentProductBaseUnit                        *string  `json:"ComponentProductBaseUnit"`
	ComponentProductDeliveryUnit                    *string  `json:"ComponentProductDeliveryUnit"`
	ComponentProductStandardScrapInPercent          *float32 `json:"ComponentProductStandardScrapInPercent"`
	IsMarkedForBackflush                            *bool    `json:"IsMarkedForBackflush"`
	BillOfMaterialItemText                          *string  `json:"BillOfMaterialItemText"`
	ValidityStartDate                               *string  `json:"ValidityStartDate"`
	ValidityEndDate                                 *string  `json:"ValidityEndDate"`
	CreationDate                                    *string  `json:"CreationDate"`
	LastChangeDate                                  *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion"`
}

func CreateBillOfMaterialRes(msg rabbitmq.RabbitmqMessage) (*BillOfMaterialRes, error) {
	res := BillOfMaterialRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
