package apiModuleRuntimesResponsesMillSheetPdf

type MillSheetPdfRes struct {
	Message   MillSheetPdf `json:"message,omitempty"`
	MountPath *string      `json:"mount_path"`
}

type MillSheetPdf struct {
	Header     []Header   `json:"Header"`
	GeneralDoc GeneralDoc `json:"GeneralDoc"`
}

type Header struct {
	OrderID                            string                  `json:"OrderID"`
	OrderItem                          int                     `json:"OrderItem"`
	BuyerName                          string                  `json:"BuyerName"`
	SellerName                         string                  `json:"SellerName"`
	Product                            string                  `json:"Product"`
	SizeOrDimensionText                string                  `json:"SizeOrDimensionText"`
	OrderQuantityInBaseUnit            float32                 `json:"OrderQuantityInBaseUnit"`
	ProductSpecification               string                  `json:"ProductSpecification"`
	MarkingOfMaterial                  string                  `json:"MarkingOfMaterial"`
	ProductionVersion                  *int                    `json:"ProductionVersion"`
	ProductionVersionItem              *int                    `json:"ProductionVersionItem"`
	ProductionOrder                    *int                    `json:"ProductionOrder"`
	ProductionOrderItem                *int                    `json:"ProductionOrderItem"`
	Contract                           *int                    `json:"Contract"`
	ContractItem                       *int                    `json:"ContractItem"`
	Project                            int                     `json:"Project"`
	WBSElement                         int                     `json:"WBSElement"`
	InspectionLot                      string                  `json:"InspectionLot"`
	InspectionPlantBusinessPartnerName string                  `json:"InspectionPlantBusinessPartnerName"`
	InspectionLotDate                  string                  `json:"InspectionLotDate"`
	InspectionSpecification            string                  `json:"InspectionSpecification"`
	HeatNumber                         string                  `json:"HeatNumber"`
	DrawingNo                          string                  `json:"DrawingNo"`
	PurchaseOrderNo                    string                  `json:"PurchaseOrderNo"`
	SpecDetails                        []SpecDetails           `json:"SpecDetails"`
	ComponentCompositions              []ComponentCompositions `json:"ComponentCompositions"`
	Inspections                        []Inspections           `json:"Inspections"`
	Remarks                            string                  `json:"Remarks"`
	ChiefOfInspectionSection           string                  `json:"ChiefOfInspectionSection"`
}

type SpecDetails struct {
	OrderID          int      `json:"OrderID"`
	OrderItem        int      `json:"OrderItem"`
	InspectionLot    int      `json:"InspectionLot"`
	SpecType         string   `json:"SpecType"`
	UpperLimitValue  *float32 `json:"UpperLimitValue"`
	LowerLimitValue  *float32 `json:"LowerLimitValue"`
	StandardValue    *float32 `json:"StandardValue"`
	SpecTypeUnit     *string  `json:"SpecTypeUnit"`
	HeatNumber       *string  `json:"HeatNumber"`
	SpecTypeTextInJA string   `json:"SpecTypeTextInJA"`
}

type ComponentCompositions struct {
	OrderID                                    int      `json:"OrderID"`
	OrderItem                                  int      `json:"OrderItem"`
	InspectionLot                              int      `json:"InspectionLot"`
	ComponentCompositionType                   string   `json:"ComponentCompositionType"`
	ComponentCompositionUpperLimitInPercent    *float32 `json:"ComponentCompositionUpperLimitInPercent"`
	ComponentCompositionLowerLimitInPercent    *float32 `json:"ComponentCompositionLowerLimitInPercent"`
	ComponentCompositionStandardValueInPercent *float32 `json:"ComponentCompositionStandardValueInPercent"`
	HeatNumber                                 *string  `json:"HeatNumber"`
}

type Inspections struct {
	OrderID                                  int      `json:"OrderID"`
	OrderItem                                int      `json:"OrderItem"`
	InspectionLot                            int      `json:"InspectionLot"`
	Inspection                               int      `json:"Inspection"`
	InspectionType                           string   `json:"InspectionType"`
	InspectionTypeCertificateValueInText     *string  `json:"InspectionTypeCertificateValueInText"`
	InspectionTypeCertificateValueInQuantity *float32 `json:"InspectionTypeCertificateValueInQuantity"`
	InspectionTypeValueUnit                  *string  `json:"InspectionTypeValueUnit"`
	InspectionTypeTextInJA                   string   `json:"InspectionTypeTextInJA"`
}

type GeneralDoc struct {
	Product                  string  `json:"Product"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}
