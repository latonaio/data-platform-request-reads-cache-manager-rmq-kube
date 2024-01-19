package apiModuleRuntimesResponsesInspectionLot

type InspectionLotDocRes struct {
	Message  *ResponseGeneralDoc `json:"message,omitempty"`
	Accepter []string            `json:"accepter,omitempty"`
}

type ResponseGeneralDoc struct {
	HeaderDoc      *[]HeaderDoc    `json:"HeaderDoc"`
	OperationDoc   *[]OperationDoc `json:"OperationDoc"`
}

type HeaderDoc struct {
	InspectionLot            int    `json:"InspectionLot"`
	DocType                  string `json:"DocType"`
	FileExtension            string `json:"FileExtension"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
	FilePath                 string `json:"FilePath"`
	FileName                 string `json:"FileName"`
}

type OperationDoc struct {
	InspectionLot            int    `json:"InspectionLot"`
	InspectionLotOperation   int    `json:"InspectionLotOperation"`
	DocType                  string `json:"DocType"`
	FileExtension            string `json:"FileExtension"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
	FilePath                 string `json:"FilePath"`
	FileName                 string `json:"FileName"`
}
