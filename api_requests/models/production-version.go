package models

type ProductionVersionReq struct {
	ConnectionKey     string            `json:"connection_key"`
	Result            bool              `json:"result"`
	RedisKey          string            `json:"redis_key"`
	Filepath          string            `json:"filepath"`
	APIStatusCode     int               `json:"api_status_code"`
	RuntimeSessionID  string            `json:"runtime_session_id"`
	BusinessPartnerID *int              `json:"business_partner"`
	ServiceLabel      string            `json:"service_label"`
	APIType           string            `json:"api_type"`
	ProductionVersion ProductionVersion `json:"ProductionVersion"`
	APISchema         string            `json:"api_schema"`
	Accepter          []string          `json:"accepter"`
	Deleted           bool              `json:"deleted"`
}

type ProductionVersion struct {
	ProductionVersion       int                     `json:"ProductionVersion"`
	Product                 string                  `json:"Product"`
	OwnerBusinessPartner    int                     `json:"OwnerBusinessPartner"`
	OwnerPlant              string                  `json:"OwnerPlant"`
	BillOfMaterial          int                     `json:"BillOfMaterial"`
	Operations              int                     `json:"Operations"`
	ProductionVersionText   *string                 `json:"ProductionVersionText"`
	ProductionVersionStatus *string                 `json:"ProductionVersionStatus"`
	ValidityStartDate       *string                 `json:"ValidityStartDate"`
	ValidityEndDate         *string                 `json:"ValidityEndDate"`
	CreationDate            *string                 `json:"CreationDate"`
	LastChangeDate          *string                 `json:"LastChangeDate"`
	IsLocked                *bool                   `json:"IsLocked"`
	IsMarkedForDeletion     *bool                   `json:"IsMarkedForDeletion"`
	Item                    []ProductionVersionItem `json:"Item"`
}

type ProductionVersionItem struct {
	ProductionVersion       int     `json:"ProductionVersion"`
	ProductionVersionItem   int     `json:"ProductionVersionItem"`
	Product                 string  `json:"Product"`
	BusinessPartner         int     `json:"BusinessPartner"`
	Plant                   string  `json:"Plant"`
	BillOfMaterial          int     `json:"BillOfMaterial"`
	Operations              int     `json:"Operations"`
	ProductionVersionText   *string `json:"ProductionVersionText"`
	ProductionVersionStatus *string `json:"ProductionVersionStatus"`
	ValidityStartDate       *string `json:"ValidityStartDate"`
	ValidityEndDate         *string `json:"ValidityEndDate"`
	CreationDate            *string `json:"CreationDate"`
	LastChangeDate          *string `json:"LastChangeDate"`
	IsLocked                *bool   `json:"IsLocked"`
	IsMarkedForDeletion     *bool   `json:"IsMarkedForDeletion"`
}
