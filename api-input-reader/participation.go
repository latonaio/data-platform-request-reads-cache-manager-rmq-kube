package apiInputReader

type Participation struct {
	ParticipationHeader			*ParticipationHeader
	ParticipationDocHeaderDoc	*ParticipationDocHeaderDoc
}

type ParticipationHeader struct {
	Participation				int		`json:"Participation"`
	Participator				*int	`json:"Participator"`
	ParticipationObjectType		*string	`json:"ParticipationObjectType"`
	ParticipationObject			*int	`json:"ParticipationObject"`
	Attendance					*int	`json:"Attendance"`
	Invitation					*int	`json:"Invitation"`
	IsCancelled					*bool	`json:"IsCancelled"`
}

type ParticipationDocHeaderDoc struct {
	Participation            int     `json:"Participation"`
	DocType                  *string `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}
