package models

type ProductStockReq struct {
	ConnectionKey     string       `json:"connection_key"`
	Result            bool         `json:"result"`
	RedisKey          string       `json:"redis_key"`
	Filepath          string       `json:"filepath"`
	APIStatusCode     int          `json:"api_status_code"`
	RuntimeSessionID  string       `json:"runtime_session_id"`
	BusinessPartnerID *int         `json:"business_partner"`
	ServiceLabel      string       `json:"service_label"`
	APIType           string       `json:"api_type"`
	ProductStock      ProductStock `json:"ProductStock"`
	APISchema         string       `json:"api_schema"`
	Accepter          []string     `json:"accepter"`
	Deleted           bool         `json:"deleted"`
}

type ProductStock struct {
	Product                   string                     `json:"Product"`
	BusinessPartner           int                        `json:"BusinessPartner"`
	Plant                     string                     `json:"Plant"`
	InventoryStockType        *string                    `json:"InventoryStockType"`
	InventorySpecialStockType *string                    `json:"InventorySpecialStockType"`
	ProductStock              *float32                   `json:"ProductStock"`
	ProductStockAvailability  []ProductStockAvailability `json:"ProductStockAvailability"`
}

type ProductStockAvailability struct {
	Product                      string   `json:"Product"`
	BusinessPartner              int      `json:"BusinessPartner"`
	Plant                        string   `json:"Plant"`
	ProductStockAvailabilityDate string   `json:"ProductStockAvailabilityDate"`
	InventoryStockType           *string  `json:"InventoryStockType"`
	InventorySpecialStockType    *string  `json:"InventorySpecialStockType"`
	AvailableProductStock        *float32 `json:"AvailableProductStock"`
}
