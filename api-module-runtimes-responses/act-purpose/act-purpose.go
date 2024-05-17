package apiModuleRuntimesResponsesActPurpose

type ActPurposeRes struct {
	Message ActPurposeGlobal `json:"message,omitempty"`
}

type ActPurposeGlobal struct {
	ActPurpose   *[]ActPurpose   `json:"ActPurpose,omitempty"`
	Text         *[]Text         `json:"Text,omitempty"`
}

type ActPurpose struct {
	ActPurpose			string	`json:"ActPurpose"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	ActPurpose     		string  `json:"ActPurpose"`
	Language          	string  `json:"Language"`
	ActPurposeName		string 	`json:"ActPurposeName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
