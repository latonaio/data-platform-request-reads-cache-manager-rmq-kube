package apiOutputFormatter

type Friend struct {
	FriendGeneral   []FriendGeneral    `json:"FriendGeneral"`
	MountPath       *string            `json:"mount_path"`
	Accepter        []string           `json:"Accepter"`
}

type FriendGeneral struct {
	BusinessPartner				int		`json:"BusinessPartner"`
	Friend						int		`json:"Friend"`
    FriendNickName              string  `json:"FriendNickName"`
	BPBusinessPartnerType		string	`json:"BPBusinessPartnerType"`
	FriendBusinessPartnerType	string	`json:"FriendBusinessPartnerType"`
	RankType					string	`json:"RankType"`
	RankTypeName				string	`json:"RankTypeName"`
	Rank						int		`json:"Rank"`
	FriendIsBlocked				bool	`json:"FriendIsBlocked"`
	Images                  	Images  `json:"Images"`
}
