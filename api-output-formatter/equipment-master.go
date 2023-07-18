package apiOutputFormatter

type EquipmentMaster struct {
	EquipmentMasterGeneral       []EquipmentMasterGeneral       `json:"Generals"`
	EquipmentMasterDetailGeneral []EquipmentMasterDetailGeneral `json:"DetailGeneral"`
}

type EquipmentMasterGeneral struct {
	Equipment            int    `json:"Equipment"`
	EquipmentName        string `json:"EquipmentName"`
	EquipmentType        string `json:"EquipmentType"`
	EquipmentTypeName    string `json:"EquipmentTypeName"`
	MaintenancePlant     string `json:"MaintenancePlant"`
	MaintenancePlantName string `json:"MaintenancePlantName"`
	BaseUnit             string `json:"BaseUnit"`
	ValidityStartDate    string `json:"ValidityStartDate"`
	IsMarkedForDeletion  *bool  `json:"IsMarkedForDeletion"`
	Images               Images `json:"Images"`
}

type EquipmentMasterDetailGeneral struct {
	EquipmentCategory                   *string  `json:"EquipmentCategory"`
	TechnicalObjectType                 *string  `json:"TechnicalObjectType"`
	GrossWeight                         *float32 `json:"GrossWeight"`
	NetWeight                           *float32 `json:"NetWeight"`
	WeightUnit                          *string  `json:"WeightUnit"`
	SizeOrDimensionText                 *string  `json:"SizeOrDimensionText"`
	InventoryNumber                     *string  `json:"InventoryNumber"`
	OperationStartDate                  *string  `json:"OperationStartDate"`
	OperationStartTime                  *string  `json:"OperationStartTime"`
	OperationEndDate                    *string  `json:"OperationEndDate"`
	OperationEndTime                    *string  `json:"OperationEndTime"`
	EquipmentStandardID                 *string  `json:"EquipmentStandardID"`
	EquipmentIndustryStandardName       *string  `json:"EquipmentIndustryStandardName"`
	AcquisitionDate                     *string  `json:"AcquisitionDate"`
	Manufacturer                        *int     `json:"Manufacturer"`
	ManufacturerCountry                 *string  `json:"ManufacturerCountry"`
	ManufacturerPartNmbr                *string  `json:"ManufacturerPartNmbr"`
	ManufacturerSerialNumber            *string  `json:"ManufacturerSerialNumber"`
	MaintenancePlantBusinessPartner     int      `json:"MaintenancePlantBusinessPartner"`
	MaintenancePlantBusinessPartnerName string   `json:"MaintenancePlantBusinessPartnerName"`
	MaintenancePlant                    string   `json:"MaintenancePlant"`
	MaintenancePlantName                string   `json:"MaintenancePlantName"`
	WorkCenter                          *int     `json:"WorkCenter"`
	Project                             *int     `json:"Project"`
	WBSElement                          *int     `json:"WBSElement"`
	FunctionalLocation                  *string  `json:"FunctionalLocation"`
	SuperordinateEquipment              *int     `json:"SuperordinateEquipment"`
	EquipmentIsAvailable                *bool    `json:"EquipmentIsAvailable"`
	EquipmentIsInstalled                *bool    `json:"EquipmentIsInstalled"`
	EquipHasSubOrdinateEquipment        *bool    `json:"EquipHasSubOrdinateEquipment"`
	MasterFixedAsset                    *string  `json:"MasterFixedAsset"`
	FixedAsset                          *string  `json:"FixedAsset"`
	CreationDate                        string   `json:"CreationDate"`
	LastChangeDate                      string   `json:"LastChangeDate"`
	IsMarkedForDeletion                 *bool    `json:"IsMarkedForDeletion"`
}
