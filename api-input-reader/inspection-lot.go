package apiInputReader

type InspectionLot struct {
	InspectionLotHeader                *InspectionLotHeader
	InspectionLotPartner               *InspectionLotPartner
	InspectionLotPartners              *InspectionLotPartners
	InspectionLotSpecDetail            *InspectionLotSpecDetail
	InspectionLotSpecDetails           *InspectionLotSpecDetails
	InspectionLotComponentComposition  *InspectionLotComponentComposition
	InspectionLotComponentCompositions *InspectionLotComponentCompositions
	InspectionLotInspection            *InspectionLotInspection
	InspectionLotInspections           *InspectionLotInspections
	InspectionLotDocOperationDoc       *InspectionLotDocOperationDoc
	InspectionLotDocHeaderDoc          *InspectionLotDocHeaderDoc
}

type InspectionLotHeader struct {
	InspectionLot       int   `json:"InspectionLot"`
	IsReleased          *bool `json:"IsReleased"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type InspectionLotPartner struct {
	InspectionLot   int    `json:"InspectionLot"`
	PartnerFunction string `json:"PartnerFunction"`
	BusinessPartner int    `json:"BusinessPartner"`
}

type InspectionLotPartners struct {
	InspectionLot int `json:"InspectionLot"`
}

type InspectionLotSpecDetail struct {
	InspectionLot       int    `json:"InspectionLot"`
	SpecType            string `json:"SpecType"`
	IsReleased          *bool  `json:"IsReleased"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type InspectionLotSpecDetails struct {
	InspectionLot       int    `json:"InspectionLot"`
	SpecType            string `json:"SpecType"`
	IsReleased          *bool  `json:"IsReleased"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type InspectionLotComponentComposition struct {
	InspectionLot            int    `json:"InspectionLot"`
	ComponentCompositionType string `json:"ComponentCompositionType"`
	IsReleased               *bool  `json:"IsReleased"`
	IsMarkedForDeletion      *bool  `json:"IsMarkedForDeletion"`
}

type InspectionLotComponentCompositions struct {
	InspectionLot            int    `json:"InspectionLot"`
	ComponentCompositionType string `json:"ComponentCompositionType"`
	IsReleased               *bool  `json:"IsReleased"`
	IsMarkedForDeletion      *bool  `json:"IsMarkedForDeletion"`
}

type InspectionLotInspection struct {
	InspectionLot       int   `json:"InspectionLot"`
	Inspection          int   `json:"Inspection"`
	IsReleased          *bool `json:"IsReleased"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type InspectionLotInspections struct {
	InspectionLot       int   `json:"InspectionLot"`
	Inspection          int   `json:"Inspection"`
	IsReleased          *bool `json:"IsReleased"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type InspectionLotOperation struct {
	InspectionLot       int   `json:"InspectionLot"`
	Operations          int   `json:"Operations"`
	OperationsItem      int   `json:"OperationsItem"`
	OperationID         int   `json:"OperationID"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type InspectionLotDocHeaderDoc struct {
	InspectionLot            int     `json:"InspectionLot"`
	DocType                  *string `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type InspectionLotDocOperationDoc struct {
	InspectionLot            int    `json:"InspectionLot"`
	Operations               int    `json:"Operations"`
	OperationsItem           int    `json:"OperationsItem"`
	OperationID              int    `json:"OperationID"`
	DocType                  string `json:"DocType"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
}
