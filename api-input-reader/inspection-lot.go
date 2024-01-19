package apiInputReader

type InspectionLot struct {
	InspectionLotHeader               *InspectionLotHeader
	InspectionLotSpecDetail           *InspectionLotSpecDetail
	InspectionLotComponentComposition *InspectionLotComponentComposition
	InspectionLotInspection           *InspectionLotInspection
	InspectionLotDocOperationDoc      *InspectionLotDocOperationDoc
}

type InspectionLotHeader struct {
	InspectionLot       int   `json:"InspectionLot"`
	IsReleased          *bool `json:"IsReleased"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type InspectionLotSpecDetail struct {
	InspectionLot       int    `json:"InspectionLot"`
	SpecType            string `json:"SpecType"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type InspectionLotComponentComposition struct {
	InspectionLot            int    `json:"InspectionLot"`
	ComponentCompositionType string `json:"ComponentCompositionType"`
	IsMarkedForDeletion      *bool  `json:"IsMarkedForDeletion"`
}

type InspectionLotInspection struct {
	InspectionLot       int   `json:"InspectionLot"`
	Inspection          int   `json:"Inspection"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type InspectionLotOperation struct {
	InspectionLot       int   `json:"InspectionLot"`
	Operations          int   `json:"Operations"`
	OperationsItem      int   `json:"OperationsItem"`
	OperationID         int   `json:"OperationID"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type InspectionLotDocOperationDoc struct {
	InspectionLot            int    `json:"InspectionLot"`
	Operations               int    `json:"Operations"`
	OperationsItem           int    `json:"OperationsItem"`
	OperationID              int    `json:"OperationID"`
	DocType                  string `json:"DocType"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
}
