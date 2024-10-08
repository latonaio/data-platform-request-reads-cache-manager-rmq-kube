package apiModuleRuntimesResponsesParticipation

type ParticipationRes struct {
	Message Participation `json:"message,omitempty"`
}

type Participation struct {
	Header    *[]Header    `json:"Header"`
}

type Header struct {
	Participation				int		`json:"Participation"`
	ParticipationDate			string	`json:"ParticipationDate"`
	ParticipationTime			string	`json:"ParticipationTime"`
	Participator				int		`json:"Participator"`
	ParticipationObjectType		string	`json:"ParticipationObjectType"`
	ParticipationObject			int		`json:"ParticipationObject"`
	Attendance					*int	`json:"Attendance"`
	Invitation					*int	`json:"Invitation"`
	CreationDate				string	`json:"CreationDate"`
	CreationTime				string	`json:"CreationTime"`
	IsCancelled					*bool	`json:"IsCancelled"`
}
