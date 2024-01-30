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
	OrderID   int `json:"OrderID"`
	OrderItem int `json:"OrderItem"`
	//	↑OrderNo                   			string	`json:"orderNo"`
	BuyerName string `json:"BuyerName"`
	//	↑Customer                  			string	`json:"customer"`
	SellerName string `json:"SellerName"`
	Product    string `json:"Product"`
	//	↑Product                   			string	`json:"product"`
	SizeOrDimensionText     string  `json:"SizeOrDimensionText"`
	OrderQuantityInBaseUnit float32 `json:"OrderQuantityInBaseUnit"`
	//	Quantity                   			string	`json:"quantity"`
	ProductSpecification string `json:"ProductSpecification"`
	//	↑StandardOfRawMaterial     			string	`json:"standardOfRawMaterial"`
	MarkingOfMaterial string `json:"MarkingOfMaterial"`
	//	↑MarkingOfMaterial         			string	`json:"markingOfMaterial"`
	ProductionVersion     *int `json:"ProductionVersion"`
	ProductionVersionItem *int `json:"ProductionVersionItem"`
	ProductionOrder       *int `json:"ProductionOrder"`
	ProductionOrderItem   *int `json:"ProductionOrderItem"`
	//	↑MfgNo                     			string	`json:"mfgNo"`
	Contract      *int `json:"Contract"`
	ContractItem  *int `json:"ContractItem"`
	Project       int  `json:"Project"`
	WBSElement    int  `json:"WBSElement"`
	InspectionLot int  `json:"InspectionLot"`
	//	↑CertificateNo             			string	`json:"certificateNo"`
	InspectionPlantBusinessPartnerName string `json:"InspectionPlantBusinessPartnerName"`
	//	↑RawMaterialMaker					string	`json:"rawMaterialMaker"`
	InspectionLotDate string `json:"InspectionLotDate"`
	//	↑Date                      	string	`json:"date"`
	InspectionSpecification string `json:"InspectionSpecification"`
	//	↑SpecificationForInspection	string	`json:"specificationForInspection"`
	HeatNumber      string        `json:"HeatNumber"`
	DrawingNo       string        `json:"DrawingNo"`
	PurchaseOrderNo string        `json:"PurchaseOrderNo"`
	SpecDetails     []SpecDetails `json:"SpecDetails"`
	//	MechanicalProperties       []MechanicalProperty  	`json:"mechanicalProperties"`
	ComponentCompositions []ComponentCompositions `json:"ComponentCompositions"`
	//	ChemicalCompositions       []ChemicalComposition 	`json:"chemicalCompositions"`
	Inspections []Inspections `json:"Inspections"`
	//	Specifications             []Specification       	`json:"specifications"`
	Remarks                  string `json:"Remarks"`
	ChiefOfInspectionSection string `json:"ChiefOfInspectionSection"`
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

//type MechanicalProperty struct {
//	LabelEn string `json:"labelEn"`
//	Label   string `json:"label"`
//	Unit    string `json:"unit,omitempty"`
//	Min     string `json:"min,omitempty"`
//	Max     string `json:"max,omitempty"`
//	HeatNo  string `json:"heatNo,omitempty"`
//}

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

//type ChemicalComposition struct {
//	Label  string `json:"label"`
//	Min    string `json:"min,omitempty"`
//	Max    string `json:"max,omitempty"`
//	HeatNo string `json:"heatNo,omitempty"`
//}

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

//type Specification struct {
//	LabelEn string `json:"labelEn"`
//	Label   string `json:"label"`
//	Value1  string `json:"value1,omitempty"`
//	Value2  string `json:"value2,omitempty"`
//}

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
