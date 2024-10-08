package apiInputReader

type Attendance struct {
	AttendanceHeader       *AttendanceHeader
	AttendanceDocHeaderDoc *AttendanceDocHeaderDoc
}

type AttendanceHeader struct {
	Attendance           int     `json:"Attendance"`
	Attender             *int    `json:"Attender"`
	AttendanceObjectType *string `json:"AttendanceObjectType"`
	AttendanceObject     *int    `json:"AttendanceObject"`
	Participation        *int    `json:"Participation"`
	Invitation           *int    `json:"Invitation"`
	IsCancelled          *bool   `json:"IsCancelled"`
}

type AttendanceDocHeaderDoc struct {
	Attendance               int     `json:"Attendance"`
	DocType                  *string `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}
