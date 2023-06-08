package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type ProductStockRes struct {
	ConnectionKey       string       `json:"connection_key,omitempty"`
	Result              bool         `json:"result,omitempty"`
	RedisKey            string       `json:"redis_key,omitempty"`
	Filepath            string       `json:"filepath,omitempty"`
	APIStatusCode       int          `json:"api_status_code,omitempty"`
	RuntimeSessionID    string       `json:"runtime_session_id,omitempty"`
	BusinessPartnerID   *int         `json:"business_partner,omitempty"`
	ServiceLabel        string       `json:"service_label,omitempty"`
	APIType             string       `json:"api_type,omitempty"`
	Message             StockMessage `json:"message,omitempty"`
	APISchema           string       `json:"api_schema,omitempty"`
	Accepter            []string     `json:"accepter,omitempty"`
	Deleted             bool         `json:"deleted,omitempty"`
	SQLUpdateResult     *bool        `json:"sql_update_result,omitempty"`
	SQLUpdateError      string       `json:"sql_update_error,omitempty"`
	SubfuncResult       *bool        `json:"subfunc_result,omitempty"`
	SubfuncError        string       `json:"subfunc_error,omitempty"`
	ExconfResult        *bool        `json:"exconf_result,omitempty"`
	ExconfError         string       `json:"exconf_error,omitempty"`
	APIProcessingResult *bool        `json:"api_processing_result,omitempty"`
	APIProcessingError  string       `json:"api_processing_error,omitempty"`
}

type StockMessage struct {
	ProductStock             *ProductStock               `json:"ProductStock,omitempty"`
	ProductStockAvailability *[]ProductStockAvailability `json:"ProductStockAvailability,omitempty"`
}

type ProductStockReads struct {
	ConnectionKey string `json:"connection_key,omitempty"`
	Result        bool   `json:"result,omitempty"`
	RedisKey      string `json:"redis_key,omitempty"`
	Filepath      string `json:"filepath,omitempty"`
	Product       string `json:"Product,omitempty"`
	APISchema     string `json:"api_schema,omitempty"`
	MaterialCode  string `json:"material_code,omitempty"`
	Deleted       string `json:"deleted,omitempty"`
}

type ProductStock struct {
	Product                   string   `json:"Product"`
	BusinessPartner           int      `json:"BusinessPartner"`
	Plant                     string   `json:"Plant"`
	InventoryStockType        *string  `json:"InventoryStockType"`
	InventorySpecialStockType *string  `json:"InventorySpecialStockType"`
	ProductStock              *float32 `json:"ProductStock"`
}

type ProductStockAvailability struct {
	Product                      string   `json:"Product,omitempty"`
	BusinessPartner              int      `json:"BusinessPartner,omitempty"`
	Plant                        string   `json:"Plant,omitempty"`
	ProductStockAvailabilityDate string   `json:"ProductStockAvailabilityDate,omitempty"`
	InventoryStockType           *string  `json:"InventoryStockType,omitempty"`
	InventorySpecialStockType    *string  `json:"InventorySpecialStockType,omitempty"`
	AvailableProductStock        *float32 `json:"AvailableProductStock,omitempty"`
}

func CreateProductStockRes(msg rabbitmq.RabbitmqMessage) (*ProductStockRes, error) {
	res := ProductStockRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
