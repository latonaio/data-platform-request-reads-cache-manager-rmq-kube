package apiInputReader

type Friend struct {
	FriendGeneral    *FriendGeneral
	FriendGeneralDoc *FriendGeneralDoc
}

type FriendGeneral struct {
	BusinessPartner     int   `json:"BusinessPartner"`
	Friend              int   `json:"Friend"`
	FriendIsBlocked     bool  `json:"FriendIsBlocked"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type FriendGeneralDoc struct {
	BusinessPartner          int     `json:"BusinessPartner"`
	DocType                  *string `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}
