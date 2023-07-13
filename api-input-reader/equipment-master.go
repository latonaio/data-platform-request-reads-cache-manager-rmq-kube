package apiInputReader

type EquipmentMaster struct {
	EquipmentMasterGeneral     *EquipmentMasterGeneral
}

type EquipmentMasterGeneral struct {
	Equipment                  int       `json:"Equipment"`
	IsMarkedForDeletion        *bool     `json:"IsMarkedForDeletion"`
}
