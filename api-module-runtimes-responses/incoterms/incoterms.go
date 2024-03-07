package apiModuleRuntimesResponsesIncoterms

type IncotermsRes struct {
	Message IncotermsGlobal `json:"message,omitempty"`
}

type IncotermsGlobal struct {
	Incoterms    *[]Incoterms    `json:"Incoterms,omitempty"`
	Text         *[]Text         `json:"Text,omitempty"`
}

type Incoterms struct {
	Incoterms			string	`json:"Incoterms"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	Incoterms     		string  `json:"Incoterms"`
	Language          	string  `json:"Language"`
	IncotermsName		string 	`json:"IncotermsName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
