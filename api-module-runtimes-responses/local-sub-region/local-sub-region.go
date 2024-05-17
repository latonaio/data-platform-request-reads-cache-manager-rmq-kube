package apiModuleRuntimesResponsesLocalSubRegion

type LocalSubRegionRes struct {
	Message LocalSubRegionGlobal `json:"message,omitempty"`
}

type LocalSubRegionGlobal struct {
	LocalSubRegion    *[]LocalSubRegion  `json:"LocalSubRegion,omitempty"`
	Text              *[]Text            `json:"Text,omitempty"`
}

type LocalSubRegion struct {
	LocalSubRegion		string	`json:"LocalSubRegion"`
	LocalRegion		    string	`json:"LocalRegion"`
    Country             string  `json:"Country"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	LocalSubRegion      string  `json:"LocalSubRegion"`
	LocalRegion		    string	`json:"LocalRegion"`
    Country             string  `json:"Country"`
	Language          	string  `json:"Language"`
	LocalSubRegionName	string 	`json:"LocalSubRegionName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
