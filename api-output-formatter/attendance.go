package apiOutputFormatter

type Attendance struct {
	AttendanceHeader          []AttendanceHeader          `json:"AttendanceHeader"`
	AttendanceHeaderWithEvent []AttendanceHeaderWithEvent `json:"AttendanceHeaderWithEvent"`
	MountPath                 *string                     `json:"mount_path"`
	Accepter                  []string                    `json:"Accepter"`
}

type AttendanceHeader struct {
	Attendance           int    `json:"Attendance"`
	AttendanceDate       string `json:"AttendanceDate"`
	AttendanceTime       string `json:"AttendanceTime"`
	Attender             int    `json:"Attender"`
	AttendanceObjectType string `json:"AttendanceObjectType"`
	AttendanceObject     int    `json:"AttendanceObject"`
	Participation        *int   `json:"Participation"`
	Invitation           *int   `json:"Invitation"`
	CreationDate         string `json:"CreationDate"`
	CreationTime         string `json:"CreationTime"`
	IsCancelled          *bool  `json:"IsCancelled"`
	Images               Images `json:"Images"`
}

type AttendanceHeaderWithEvent struct {
	Attendance              int    `json:"Attendance"`
	AttendanceDate          string `json:"AttendanceDate"`
	AttendanceTime          string `json:"AttendanceTime"`
	Attender                int    `json:"Attender"`
	AttendanceObjectType    string `json:"AttendanceObjectType"`
	AttendanceObject        int    `json:"AttendanceObject"`
	Participation           *int   `json:"Participation"`
	Invitation              *int   `json:"Invitation"`
	CreationDate            string `json:"CreationDate"`
	CreationTime            string `json:"CreationTime"`
	IsCancelled             *bool  `json:"IsCancelled"`
	Event                   int    `json:"Event"`
	EventOperationStartDate string `json:"EventOperationStartDate"`
	EventOperationStartTime string `json:"EventOperationStartTime"`
	EventOperationEndDate   string `json:"EventOperationEndDate"`
	EventOperationEndTime   string `json:"EventOperationEndTime"`
	EventDescription        string `json:"EventDescription"`
	EventSite               int    `json:"EventSite"`
	Site					int	   `json:"Site"`
	SiteDescription         string `json:"SiteDescription"`
	Images                  Images `json:"Images"`
}
