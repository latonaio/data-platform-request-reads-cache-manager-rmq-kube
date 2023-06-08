package dpfm_api_output_formatter

type EquipmentDetail struct {
	Equipment         int                      `json:"Equipment"`
	EquipmentName     string                   `json:"EquipmentName"`
	EquipmentTypeName string                   `json:"EquipmentTypeName"`
	PlantName         string                   `json:"PlantName"`
	ValidityStartDate string                   `Json:"ValidityStartDate"`
	EquipmentDetail   []EquipmentDetailGeneral `json:"DeliveryDocumentDetail"`
	EquipmentBP       []EquipmentBP            `json:"EquipmentBP"`
	EquipmentOwnerBP  []EquipmentOwnerBP       `json:"EquipmentOwnerBP"`
}

type EquipmentDetailGeneral struct {
	EquipmentCategory        string  `json:"EquipmentCategory"`
	TechnicalObjectType      string  `json:"TechnicalObjectType"`
	GrossWeight              float32 `json:"GrossWeight"`
	NetWeight                float32 `json:"NetWeight"`
	WeightUnit               string  `json:"WeightUnit"`
	SizeOrDimensionText      string  `json:"SizeOrDimensionText"`
	OperationStartDate       string  `json:"OperationStartDate"`
	OperationStartTime       string  `json:"OperationStartTime"`
	OperationEndDate         string  `json:"OperationEndDate"`
	OperationEndTime         string  `json:"OperationEndTime"`
	AcquisitionDate          string  `json:"AcquisitionDate"`
	BusinessPartnerName      string  `json:"BusinessPartnerName"`
	ManifacturerSerialNumber string  `json:"ManifacturerSerialNumber"`
	MasterFixedAsset         string  `json:"MasterFixedAsset"`
	FixedAsset               string  `json:"FixedAsset"`
	ValidityEndDate          string  `json:"ValidityEndDate"`
	IsMarkedForDeletion      bool    `json:"IsMarkedForDeletion"`
}

type EquipmentBP struct {
	EquipmentPartnerObjectNmbr int    `json:"EquipmentPartnerObjectNmbr"`
	BusinessPartner            int    `json:"BusinessPartner"`
	PartnerFunction            string `json:"PartnerFunction"`
	//BP名称が入ります（名称、リクエスト先不明）

	ValidityStartDate string `json:"ValidityStartDate"`
	ValidityEndDate   string `json:"ValidityEndDate"`
}

type EquipmentOwnerBP struct {
}
