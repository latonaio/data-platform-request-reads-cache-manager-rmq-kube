package models

type WorkCenterReq struct {
	ConnectionKey     string            `json:"connection_key"`
	Result            bool              `json:"result"`
	RedisKey          string            `json:"redis_key"`
	Filepath          string            `json:"filepath"`
	APIStatusCode     int               `json:"api_status_code"`
	RuntimeSessionID  string            `json:"runtime_session_id"`
	BusinessPartnerID *int              `json:"business_partner"`
	ServiceLabel      string            `json:"service_label"`
	APIType           string            `json:"api_type"`
	General           WorkCenterGeneral `json:"WorkCenter"`
	APISchema         string            `json:"api_schema"`
	Accepter          []string          `json:"accepter"`
	Deleted           bool              `json:"deleted"`
}

type WorkCenterGeneral struct {
	WorkCenter                   int     `json:"WorkCenter"`
	WorkCenterType               *string `json:"WorkCenterType"`
	WorkCenterName               *string `json:"WorkCenterName"`
	BusinessPartner              *int    `json:"BusinessPartner"`
	Plant                        *string `json:"Plant"`
	WorkCenterCategory           *string `json:"WorkCenterCategory"`
	WorkCenterResponsible        *string `json:"WorkCenterResponsible"`
	SupplyArea                   *string `json:"SupplyArea"`
	WorkCenterUsage              *string `json:"WorkCenterUsage"`
	MatlCompIsMarkedForBackflush *bool   `json:"MatlCompIsMarkedForBackflush"`
	WorkCenterLocation           *string `json:"WorkCenterLocation"`
	CapacityInternalID           *string `json:"CapacityInternalID"`
	CapacityCategoryCode         *string `json:"CapacityCategoryCode"`
	ValidityStartDate            *string `json:"ValidityStartDate"`
	ValidityEndDate              *string `json:"ValidityEndDate"`
	IsMarkedForDeletion          *bool   `json:"IsMarkedForDeletion"`
}
