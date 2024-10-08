package apiOutputFormatter

type Participation struct {
	ParticipationHeader          []ParticipationHeader          `json:"ParticipationHeader"`
	ParticipationHeaderWithEvent []ParticipationHeaderWithEvent `json:"ParticipationHeaderWithEvent"`
	MountPath                    *string                        `json:"mount_path"`
	Accepter                     []string                       `json:"Accepter"`
}

type ParticipationHeader struct {
	Participation           int    `json:"Participation"`
	ParticipationDate       string `json:"ParticipationDate"`
	ParticipationTime       string `json:"ParticipationTime"`
	Participator            int    `json:"Participator"`
	ParticipationObjectType string `json:"ParticipationObjectType"`
	ParticipationObject     int    `json:"ParticipationObject"`
	Attendance              *int   `json:"Attendance"`
	Invitation              *int   `json:"Invitation"`
	CreationDate            string `json:"CreationDate"`
	CreationTime            string `json:"CreationTime"`
	IsCancelled             *bool  `json:"IsCancelled"`
	Images                  Images `json:"Images"`
}

type ParticipationHeaderWithEvent struct {
	Participation           int    `json:"Participation"`
	ParticipationDate       string `json:"ParticipationDate"`
	ParticipationTime       string `json:"ParticipationTime"`
	Participator            int    `json:"Participator"`
	ParticipationObjectType string `json:"ParticipationObjectType"`
	ParticipationObject     int    `json:"ParticipationObject"`
	Attendance              *int   `json:"Attendance"`
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
