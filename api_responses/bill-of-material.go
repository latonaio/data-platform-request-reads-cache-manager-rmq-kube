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
	BillOfMaterial              int      `json:"BillOfMaterial,omitempty"`
	BillOfMaterialType          *string  `json:"BillOfMaterialType,omitempty"`
	Product                     *string  `json:"Product,omitempty"`
	OwnerBusinessPartner        *int     `json:"OwnerBusinessPartner,omitempty"`
	OwnerPlant                  *string  `json:"OwnerPlant,omitempty"`
	BOMAlternativeText          *string  `json:"BOMAlternativeText,omitempty"`
	BOMHeaderBaseUnit           *string  `json:"BOMHeaderBaseUnit,omitempty"`
	BOMHeaderQuantityInBaseUnit *float32 `json:"BOMHeaderQuantityInBaseUnit,omitempty"`
	ValidityStartDate           *string  `json:"ValidityStartDate,omitempty"`
	ValidityEndDate             *string  `json:"ValidityEndDate,omitempty"`
	CreationDate                *string  `json:"CreationDate,omitempty"`
	LastChangeDate              *string  `json:"LastChangeDate,omitempty"`
	BillOfMaterialHeaderText    *string  `json:"BillOfMaterialHeaderText,omitempty"`
	IsMarkedForDeletion         *bool    `json:"IsMarkedForDeletion,omitempty"`
}

type BillOfMaterialItem struct {
	BillOfMaterial                  int      `json:"BillOfMaterial,omitempty"`
	BillOfMaterialItem              int      `json:"BillOfMaterialItem,omitempty"`
	ComponentProduct                *string  `json:"ComponentProduct,omitempty"`
	ComponentProductBusinessPartner *int     `json:"ComponentProductBusinessPartner,omitempty"`
	StockConfirmationPlant          *string  `json:"StockConfirmationPlant,omitempty"`
	BOMAlternativeText              *string  `json:"BOMAlternativeText,omitempty"`
	BOMItemBaseUnit                 *string  `json:"BOMItemBaseUnit,omitempty"`
	BOMItemQuantityInBaseUnit       *float32 `json:"BOMItemQuantityInBaseUnit,omitempty"`
	ComponentScrapInPercent         *int     `json:"ComponentScrapInPercent,omitempty"`
	ValidityStartDate               *string  `json:"ValidityStartDate,omitempty"`
	ValidityEndDate                 *string  `json:"ValidityEndDate,omitempty"`
	BillOfMaterialItemText          *string  `json:"BillOfMaterialItemText,omitempty"`
	IsMarkedForDeletion             *bool    `json:"IsMarkedForDeletion,omitempty"`
}

func CreateBillOfMaterialRes(msg rabbitmq.RabbitmqMessage) (*BillOfMaterialRes, error) {
	res := BillOfMaterialRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
