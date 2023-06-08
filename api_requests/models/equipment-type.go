package models

type EquipmentTypeReq struct {
	ConnectionKey     string          `json:"connection_key"`
	Result            bool            `json:"result"`
	RedisKey          string          `json:"redis_key"`
	Filepath          string          `json:"filepath"`
	APIStatusCode     int             `json:"api_status_code"`
	RuntimeSessionID  string          `json:"runtime_session_id"`
	BusinessPartnerID *int            `json:"business_partner"`
	ServiceLabel      string          `json:"service_label"`
	APIType           string          `json:"api_type"`
	EquipmentType     []EquipmentType `json:"EquipmentType"`
	APISchema         string          `json:"api_schema"`
	Accepter          []string        `json:"accepter"`
	Deleted           bool            `json:"deleted"`
}

type EquipmentType struct {
	EquipmentType     string              `json:"EquipmentType"`
	EquipmentTypeText []EquipmentTypeText `json:"EquipmentTypeText"`
}

type EquipmentTypeText struct {
	Language          string `json:"Language"`
	EquipmentTypeName string `json:"EquipmentTypeName"`
}
