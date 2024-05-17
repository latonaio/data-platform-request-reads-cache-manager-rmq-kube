package apiModuleRuntimesResponsesFriend

type FriendRes struct {
	Message Friend `json:"message,omitempty"`
}

type Friend struct {
	General    *[]General    `json:"General"`
}

type General struct {
	BusinessPartner				int		`json:"BusinessPartner"`
	Friend						int		`json:"Friend"`
	BPBusinessPartnerType		string	`json:"BPBusinessPartnerType"`
	FriendBusinessPartnerType	string	`json:"FriendBusinessPartnerType"`
	RankType					string	`json:"RankType"`
	Rank						int		`json:"Rank"`
	FriendIsBlocked				bool	`json:"FriendIsBlocked"`
	CreationDate				string	`json:"CreationDate"`
	CreationTime				string	`json:"CreationTime"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}
