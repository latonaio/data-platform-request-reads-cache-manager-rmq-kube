package dpfm_api_output_formatter

type EquipmentList struct {
	Equipments []Equipment `json:"Equipments"`
}

type Equipment struct {
	Equipment           int     `json:"Equipment"`
	EquipmentName       string  `json:"EquipmentName"`
	EquipmentTypeName   *string `json:"EquipmentTypeName"`
	PlantName           *string `json:"PlantName"`
	ValidityStartDate   *string `json:"ValidityStartDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	Images              Images  `json:"Images"`
}
