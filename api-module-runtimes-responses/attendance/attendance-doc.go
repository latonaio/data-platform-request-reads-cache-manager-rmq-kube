package apiModuleRuntimesResponsesAttendance

type AttendanceDocRes struct {
	Message  *ResponseGeneralDoc `json:"message,omitempty"`
	Accepter []string            `json:"accepter,omitempty"`
}

type ResponseGeneralDoc struct {
	HeaderDoc *[]HeaderDoc `json:"HeaderDoc"`
}

type HeaderDoc struct {
	Attendance               int    `json:"Attendance"`
	DocType                  string `json:"DocType"`
	FileExtension            string `json:"FileExtension"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
	FilePath                 string `json:"FilePath"`
	FileName                 string `json:"FileName"`
}
