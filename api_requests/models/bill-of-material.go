package models

type BillOfMaterialReq struct {
	ConnectionKey     string               `json:"connection_key"`
	Result            bool                 `json:"result"`
	RedisKey          string               `json:"redis_key"`
	Filepath          string               `json:"filepath"`
	APIStatusCode     int                  `json:"api_status_code"`
	RuntimeSessionID  string               `json:"runtime_session_id"`
	BusinessPartnerID *int                 `json:"business_partner"`
	ServiceLabel      string               `json:"service_label"`
	APIType           string               `json:"api_type"`
	Header            BillOfMaterialHeader `json:"BillOfMaterial"`
	APISchema         string               `json:"api_schema"`
	Accepter          []string             `json:"accepter"`
	Deleted           bool                 `json:"deleted"`
}

type BillOfMaterialHeader struct {
	BillOfMaterial              int                        `json:"BillOfMaterial"`
	BillOfMaterialType          *string                    `json:"BillOfMaterialType"`
	Product                     *string                    `json:"Product"`
	OwnerBusinessPartner        *int                       `json:"OwnerBusinessPartner"`
	OwnerPlant                  *string                    `json:"OwnerPlant"`
	BOMAlternativeText          *string                    `json:"BOMAlternativeText"`
	BOMHeaderBaseUnit           *string                    `json:"BOMHeaderBaseUnit"`
	BOMHeaderQuantityInBaseUnit *float32                   `json:"BOMHeaderQuantityInBaseUnit"`
	ValidityStartDate           *string                    `json:"ValidityStartDate"`
	ValidityEndDate             *string                    `json:"ValidityEndDate"`
	CreationDate                *string                    `json:"CreationDate"`
	LastChangeDate              *string                    `json:"LastChangeDate"`
	BillOfMaterialHeaderText    *string                    `json:"BillOfMaterialHeaderText"`
	IsMarkedForDeletion         *bool                      `json:"IsMarkedForDeletion"`
	Item                        []BillOfMaterialHeaderItem `json:"Item"`
}

type BillOfMaterialHeaderItem struct {
	BillOfMaterial                  int      `json:"BillOfMaterial"`
	BillOfMaterialItem              int      `json:"BillOfMaterialItem"`
	ComponentProduct                *string  `json:"ComponentProduct"`
	ComponentProductBusinessPartner *int     `json:"ComponentProductBusinessPartner"`
	StockConfirmationPlant          *string  `json:"StockConfirmationPlant"`
	BOMAlternativeText              *string  `json:"BOMAlternativeText"`
	BOMItemBaseUnit                 *string  `json:"BOMItemBaseUnit"`
	BOMItemQuantityInBaseUnit       *float32 `json:"BOMItemQuantityInBaseUnit"`
	ComponentScrapInPercent         *int     `json:"ComponentScrapInPercent"`
	ValidityStartDate               *string  `json:"ValidityStartDate"`
	ValidityEndDate                 *string  `json:"ValidityEndDate"`
	BillOfMaterialItemText          *string  `json:"BillOfMaterialItemText"`
	IsMarkedForDeletion             *bool    `json:"IsMarkedForDeletion"`
}
