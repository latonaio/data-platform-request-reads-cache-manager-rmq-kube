package apiModuleRuntimesResponsesLocalRegion

type LocalRegionRes struct {
	Message LocalRegionGlobal `json:"message,omitempty"`
}

type LocalRegionGlobal struct {
	LocalRegion  *[]LocalRegion  `json:"LocalRegion,omitempty"`
	Text         *[]Text         `json:"Text,omitempty"`
}

type LocalRegion struct {
	LocalRegion			string	`json:"LocalRegion"`
    Country             string  `json:"Country"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	LocalRegion         string  `json:"LocalRegion"`
    Country             string  `json:"Country"`
	Language          	string  `json:"Language"`
	LocalRegionName		string 	`json:"LocalRegionName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
