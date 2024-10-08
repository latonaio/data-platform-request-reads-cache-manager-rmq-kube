package apiOutputFormatter

type Event struct {
	EventHeader             	[]EventHeader             		`json:"EventHeader"`
	EventPartner            	[]EventPartner            		`json:"EventPartner"`
	EventAddress            	[]EventAddress            		`json:"EventAddress"`
	EventParticipation			[]EventParticipation			`json:"EventParticipation"`
	EventAttendance				[]EventAttendance				`json:"EventAttendance"`
	EventPointConditionElement	[]EventPointConditionElement	`json:"EventPointConditionElement"`
	EventAddressWithHeader  	[]EventAddressWithHeader  		`json:"EventAddressWithHeader"`
	EventPartnerWithAddress 	[]EventPartnerWithAddress 		`json:"EventPartnerWithAddress"`
	EventCounter				[]EventCounter					`json:"EventCounter"`
	MountPath               	*string                   		`json:"mount_path"`
	Accepter                	[]string                  		`json:"Accepter"`
}

type EventHeader struct {
	Event                   			int     `json:"Event"`
	EventType               			string  `json:"EventType"`
	EventTypeName           			string  `json:"EventTypeName"`
	EventOwner              			int     `json:"EventOwner"`
	EventOwnerName          			string  `json:"EventOwnerName"`
	EventOwnerBusinessPartnerRole		string	`json:"EventOwnerBusinessPartnerRole"`
	EventOwnerBusinessPartnerRoleName	string	`json:"EventOwnerBusinessPartnerRoleName"`
	PersonResponsible       			string  `json:"PersonResponsible"`
	URL									*string	`json:"URL"`
	ValidityStartDate       			string  `json:"ValidityStartDate"`
	ValidityStartTime       			string  `json:"ValidityStartTime"`
	ValidityEndDate         			string  `json:"ValidityEndDate"`
	ValidityEndTime         			string  `json:"ValidityEndTime"`
	OperationStartDate					string	`json:"OperationStartDate"`
	OperationStartTime					string	`json:"OperationStartTime"`
	OperationEndDate					string	`json:"OperationEndDate"`
	OperationEndTime					string	`json:"OperationEndTime"`
	Description             			string  `json:"Description"`
	LongText                			string  `json:"LongText"`
	Introduction            			*string `json:"Introduction"`
	Site                    			int     `json:"Site"`
	SiteDescription            			string  `json:"SiteDescription"`
	Capacity							int		`json:"Capacity"`
	Shop                    			*int    `json:"Shop"`
	ShopDescription            			*string `json:"ShopDescription"`
	Tag1                    			*string `json:"Tag1"`
	Tag2                    			*string `json:"Tag2"`
	Tag3                    			*string `json:"Tag3"`
	Tag4                    			*string `json:"Tag4"`
	DistributionProfile     			string  `json:"DistributionProfile"`
	DistributionProfileName 			string  `json:"DistributionProfileName"`
	PointConditionType      			string  `json:"PointConditionType"`
	PointConditionTypeName  			string  `json:"PointConditionTypeName"`
	QuestionnaireType					*string `json:"QuestionnaireType"`
	QuestionnaireTypeName				*string `json:"QuestionnaireTypeName"`
	QuestionnaireTemplate				*string `json:"QuestionnaireTemplate"`
	QuestionnaireTemplateName			*string `json:"QuestionnaireTemplateName"`
	LastChangeDate                		string  `json:"LastChangeDate"`
	LastChangeTime                		string  `json:"LastChangeTime"`
	CreateUser						 	int	 	`json:"CreateUser"`
	CreateUserFullName				 	*string `json:"CreateUserFullName"`
	CreateUserNickName				 	*string `json:"CreateUserNickName"`
	LastChangeUser					 	int 	`json:"LastChangeUser"`
	LastChangeUserFullName			 	*string `json:"LastChangeUserFullName"`
	LastChangeUserNickName			 	*string `json:"LastChangeUserNickName"`
	NumberOfLikes						*int	`json:"NumberOfLikes"`
	NumberOfParticipations				*int	`json:"NumberOfParticipations"`
	NumberOfAttendances					*int	`json:"NumberOfAttendances"`
	Images                  			Images  `json:"Images"`
}

type EventPartner struct {
	Event                   int     `json:"Event"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
	CityName                *string `json:"CityName"`
	StreetName              *string `json:"StreetName"`
	PostalCode              *string `json:"PostalCode"`
	AddressIdentifier       *string `json:"AddressIdentifier"`
	LocalRegionName         *string `json:"LocalRegionName"`
}

type EventAddress struct {
	Event              int     `json:"Event"`
	AddressID          int     `json:"AddressID"`
	PostalCode         *string `json:"PostalCode"`
	LocalSubRegion     string  `json:"LocalSubRegion"`
	LocalSubRegionName string  `json:"LocalSubRegionName"`
	LocalRegion        string  `json:"LocalRegion"`
	LocalRegionName    string  `json:"LocalRegionName"`
	Country            *string `json:"Country"`
	StreetName         *string `json:"StreetName"`
	CityName           *string `json:"CityName"`
	Building           *string `json:"Building"`
	Floor              *int    `json:"Floor"`
	Room               *int    `json:"Room"`
}

type EventParticipation struct {
	Event							int		`json:"Event"`
	Participator					int		`json:"Participator"`
	Participation					int		`json:"Participation"`
	IsCancelled						*bool	`json:"IsCancelled"`
}

type EventAttendance struct {
	Event							int		`json:"Event"`
	Attender						int		`json:"Attender"`
	Attendance						int		`json:"Attendance"`
	Participation					int		`json:"Participation"`
	IsCancelled						*bool	`json:"IsCancelled"`
}

type EventPointConditionElement struct {
	Event							int		`json:"Event"`
	PointConditionRecord			int		`json:"PointConditionRecord"`
	PointConditionSequentialNumber	int		`json:"PointConditionSequentialNumber"`
	PointSymbol						string	`json:"PointSymbol"`
	Sender							int		`json:"Sender"`
	PointTransactionType			string	`json:"PointTransactionType"`
	PointConditionType				string	`json:"PointConditionType"`
	PointConditionRateValue			float32	`json:"PointConditionRateValue"`
	PointConditionRatio				float32	`json:"PointConditionRatio"`
	PlusMinus						string	`json:"PlusMinus"`
}

type EventCounter struct {
	Event					int		`json:"Event"`
	NumberOfLikes			int		`json:"NumberOfLikes"`
	NumberOfParticipations	int		`json:"NumberOfParticipations"`
	NumberOfAttendances		int		`json:"NumberOfAttendances"`
}

type EventAddressWithHeader struct {
	Event             		int     `json:"Event"`
	AddressID         		int     `json:"AddressID"`
	LocalSubRegion    		string  `json:"LocalSubRegion"`
	LocalRegion       		string  `json:"LocalRegion"`
	EventType         		string  `json:"EventType"`
	EventOwner        		int     `json:"EventOwner"`
	EventOwnerName    		string  `json:"EventOwnerName"`
	ValidityStartDate 		string  `json:"ValidityStartDate"`
	ValidityStartTime 		string  `json:"ValidityStartTime"`
	ValidityEndDate   		string  `json:"ValidityEndDate"`
	ValidityEndTime   		string  `json:"ValidityEndTime"`
	Description       		string  `json:"Description"`
	LongText          		string  `json:"LongText"`
	Introduction      		*string `json:"Introduction"`
	Site              		int     `json:"Site"`
	Capacity		  		int		`json:"Capacity"`
	Tag1              		*string `json:"Tag1"`
	Tag2              		*string `json:"Tag2"`
	Tag3              		*string `json:"Tag3"`
	Tag4              		*string `json:"Tag4"`
	LastChangeDate          string  `json:"LastChangeDate"`
	LastChangeTime          string  `json:"LastChangeTime"`
	NumberOfLikes			*int	`json:"NumberOfLikes"`
	NumberOfParticipations	*int	`json:"NumberOfParticipations"`
	NumberOfAttendances		*int	`json:"NumberOfAttendances"`
	Images            		Images  `json:"Images"`
}

type EventPartnerWithAddress struct {
	Event                   int     `json:"Event"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	AddressID               *int    `json:"AddressID"`
	PostalCode              *string `json:"PostalCode"`
	LocalRegion             string  `json:"LocalRegion"`
	LocalRegionName         string  `json:"LocalRegionName"`
	Country                 string  `json:"Country"`
	StreetName              *string `json:"StreetName"`
	CityName                *string `json:"CityName"`
	Building                *string `json:"Building"`
	Floor                   *int    `json:"Floor"`
	Room                    *int    `json:"Room"`
}
