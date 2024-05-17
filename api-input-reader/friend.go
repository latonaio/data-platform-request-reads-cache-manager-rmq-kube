package apiInputReader

type Friend struct {
	FriendGeneral *FriendGeneral
}

type FriendGeneral struct {
	BusinessPartner     int   `json:"BusinessPartner"`
	Friend              int   `json:"Friend"`
	FriendIsBlocked     bool  `json:"FriendIsBlocked"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
