package apiOutputFormatter

type InspectionLot struct {
	InspectionLotHeader               []InspectionLotHeader               `json:"Header"`
	InspectionLotSingleUnit           []InspectionLotSingleUnit           `json:"InspectionLotSingleUnit"`
	InspectionLotPartner              []InspectionLotPartner              `json:"Partner"`
	InspectionLotSpecDetail           []InspectionLotSpecDetail           `json:"SpecDetail"`
	InspectionLotComponentComposition []InspectionLotComponentComposition `json:"ComponentComposition"`
	InspectionLotInspection           []InspectionLotInspection           `json:"Inspection"`
	InspectionLotPdfMountPath         *string                             `json:"InspectionLotPdfMountPath"`
}

type InspectionLotHeader struct {
	InspectionLot                      int     `json:"InspectionLot"`
	InspectionLotDate                  string  `json:"InspectionLotDate"`
	InspectionPlantBusinessPartner     int     `json:"InspectionPlantBusinessPartner"`
	InspectionPlantBusinessPartnerName string  `json:"InspectionPlantBusinessPartnerName"`
	InspectionPlant                    string  `json:"InspectionPlant"`
	InspectionPlantName                string  `json:"InspectionPlantNamejk"`
	Product                            string  `json:"Product"`
	ProductSpecification               *string `json:"ProductSpecification"`
	InspectionSpecification            *string `json:"InspectionSpecification"`
	ProductionOrder                    *int    `json:"ProductionOrder"`
	ProductionOrderItem                *int    `json:"ProductionOrderItem"`
	CertificateAuthorityChain          *string `json:"CertificateAuthorityChain"`
	UsageControlChain                  *string `json:"UsageControlChain"`
}

type InspectionLotPartner struct {
	InspectionLot		int `json:"InspectionLot"`
	PartnerFunction     string `json:"PartnerFunction"`
	BusinessPartner     string `json:"BusinessPartner"`
	BusinessPartnerName string `json:"BusinessPartnerName"`
	Product             string `json:"Product"`
	InspectionLotDate   string `json:"InspectionLotDate"`
}

type InspectionLotSingleUnit struct {
	InspectionLot                      int     `json:"InspectionLot"`
	InspectionLotDate                  string  `json:"InspectionLotDate"`
	InspectionPlantBusinessPartner     int     `json:"InspectionPlantBusinessPartner"`
	InspectionPlantBusinessPartnerName string  `json:"InspectionPlantBusinessPartnerName"`
	InspectionPlant                    string  `json:"InspectionPlant"`
	InspectionPlantName                string  `json:"InspectionPlantName"`
	Product                            string  `json:"Product"`
	ProductSpecification               *string `json:"ProductSpecification"`
	ProductionOrder                    *int    `json:"ProductionOrder"`
	ProductionOrderItem                *int    `json:"ProductionOrderItem"`
	CertificateAuthorityChain          *string `json:"CertificateAuthorityChain"`
	UsageControlChain                  *string `json:"UsageControlChain"`
	Images                             Images  `json:"Images"`
}

type InspectionLotSpecDetail struct {
	InspectionLot   int     `json:"InspectionLot"`
	SpecType        string  `json:"SpecType"`
	SpecTypeText    string  `json:"SpecTypeText"`
	UpperLimitValue float32 `json:"UpperLimitValue"`
	LowerLimitValue float32 `json:"LowerLimitValue"`
	StandardValue   float32 `json:"StandardValue"`
	SpecTypeUnit    string  `json:"SpecTypeUnit"`
	Formula         *string `json:"Formula"`
}

type InspectionLotComponentComposition struct {
	InspectionLot                              int     `json:"InspectionLot"`
	ComponentCompositionType                   string  `json:"ComponentCompositionType"`
	ComponentCompositionTypeText               string  `json:"ComponentCompositionTypeText"`
	ComponentCompositionUpperLimitInPercent    float32 `json:"ComponentCompositionUpperLimitInPercent"`
	ComponentCompositionLowerLimitInPercent    float32 `json:"ComponentCompositionLowerLimitInPercent"`
	ComponentCompositionStandardValueInPercent float32 `json:"ComponentCompositionStandardValueInPercent"`
}

type InspectionLotInspection struct {
	InspectionLot                            int      `json:"InspectionLot"`
	Inspection                               int      `json:"Inspection"`
	InspectionType                           string   `json:"InspectionType"`
	InspectionTypeValueUnit                  *string  `json:"InspectionTypeValueUnit"`
	InspectionTypePlannedValue               *float32 `json:"InspectionTypePlannedValue"`
	InspectionTypeCertificateType            *string  `json:"InspectionTypeCertificateType"`
	InspectionTypeCertificateValueInText     *string  `json:"InspectionTypeCertificateValueInText"`
	InspectionTypeCertificateValueInQuantity *float32 `json:"InspectionTypeCertificateValueInQuantity"`
	InspectionLotInspectionText              *string  `json:"InspectionLotInspectionText"`
}
