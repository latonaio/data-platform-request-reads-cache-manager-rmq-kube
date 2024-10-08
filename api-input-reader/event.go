package apiInputReader

type Event struct {
	EventHeader              	*EventHeader
	EventPartner             	*EventPartner
	EventAddress             	*EventAddress
    EventCampaign            	*EventCampaign
    EventGame                	*EventGame
	EventParticipation			*EventParticipation
	EventAttendance				*EventAttendance
    EventPointTransaction    	*EventPointTransaction
    EventPointConditionElement	*EventPointConditionElement
	EventCounter              	*EventCounter
	EventDocHeaderDoc        	*EventDocHeaderDoc
}

type EventHeader struct {
	Event                    int     `json:"Event"`
	EventOwner               *int    `json:"EventOwner"`
	IsReleased               *bool   `json:"IsReleased"`
	IsCancelled              *bool   `json:"IsCancelled"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

type EventPartner struct {
	Event				int		`json:"Event"`
	PartnerFunction		string	`json:"PartnerFunction"`
	BusinessPartner		int		`json:"BusinessPartner"`
}

type EventAddress struct {
	Event			int		`json:"Event"`
	AddressID		int		`json:"AddressID"`
	LocalSubRegion 	*string `json:"LocalSubRegion"`
	LocalRegion 	*string `json:"LocalRegion"`
}

type EventCampaign struct {
	Event		             int     `json:"Event"`
	Campaign				 int	 `json:"Campaign"`
	IsReleased               *bool   `json:"IsReleased"`
	IsCancelled              *bool   `json:"IsCancelled"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

type EventGame struct {
	Event		             int     `json:"Event"`
	Game				 	 int	 `json:"Game"`
	IsReleased               *bool   `json:"IsReleased"`
	IsCancelled              *bool   `json:"IsCancelled"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

type EventParticipation struct {
	Event							int		`json:"Event"`
	Participator					int		`json:"Participator"`
	Participation					*int	`json:"Participation"`
	IsCancelled						*bool	`json:"IsCancelled"`
}

type EventAttendance struct {
	Event							int		`json:"Event"`
	Attender						int		`json:"Attender"`
	Attendance						*int	`json:"Attendance"`
	IsCancelled						*bool	`json:"IsCancelled"`
}

type EventPointTransaction struct {
	Event		             		int     `json:"Event"`
	Sender				 	 		int	 	`json:"Sender"`
	Receiver				 		int	 	`json:"Receiver"`
	PointConditionRecord	 		int	 	`json:"PointConditionRecord"`
	PointConditionSequentialNumber	int	 	`json:"PointConditionSequentialNumber"`
	IsCancelled              		*bool   `json:"IsCancelled"`
}

type EventPointConditionElement struct {
	Event		             		int     `json:"Event"`
	PointConditionRecord	 		int	 	`json:"PointConditionRecord"`
	PointConditionSequentialNumber	int	 	`json:"PointConditionSequentialNumber"`
	IsReleased               		*bool   `json:"IsReleased"`
	IsCancelled              		*bool   `json:"IsCancelled"`
	IsMarkedForDeletion      		*bool   `json:"IsMarkedForDeletion"`
}

type EventCounter struct {
	Event                    int     `json:"Event"`
}

type EventDocHeaderDoc struct {
	Event            		 int     `json:"Event"`
	DocType                  *string `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}
