package apiModuleRuntimesResponsesCountry

type CountryRes struct {
	Message CountryGlobal `json:"message,omitempty"`
}

type CountryGlobal struct {
	Country    *[]Country    `json:"Country,omitempty"`
	Text       *[]Text       `json:"Text,omitempty"`
}

type Country struct {
	Country			    string	`json:"Country"`
	GlobalRegion		string	`json:"GlobalRegion"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	Country             string  `json:"Country"`
	Language          	string  `json:"Language"`
	CountryName		    string 	`json:"CountryName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
