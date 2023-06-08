package models

type StorageBinReq struct {
	ConnectionKey     string     `json:"connection_key"`
	Result            bool       `json:"result"`
	RedisKey          string     `json:"redis_key"`
	Filepath          string     `json:"filepath"`
	APIStatusCode     int        `json:"api_status_code"`
	RuntimeSessionID  string     `json:"runtime_session_id"`
	BusinessPartnerID *int       `json:"business_partner"`
	ServiceLabel      string     `json:"service_label"`
	APIType           string     `json:"APIType"`
	General           StorageBin `json:"StorageBin"`
	APISchema         string     `json:"api_schema"`
	Accepter          []string   `json:"accepter"`
	Deleted           bool       `json:"deleted"`
}

type StorageBin struct {
	BusinessPartner    int      `json:"BusinessPartner"`
	Plant              string   `json:"Plant"`
	StorageLocation    string   `json:"StorageLocation"`
	StorageBin         string   `json:"StorageBin"`
	StorageType        *string  `json:"StorageType"`
	XCoordinates       *float32 `json:"XCoordinates"`
	YCoordinates       *float32 `json:"YCoordinates"`
	ZCoordinates       *float32 `json:"ZCoordinates"`
	Capacity           *float32 `json:"Capacity"`
	CapacityUnit       *string  `json:"CapacityUnit"`
	CapacityWeight     *float32 `json:"CapacityWeight"`
	CapacityWeightUnit *string  `json:"CapacityWeightUnit"`
}
