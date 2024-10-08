package apiModuleRuntimesResponsesObjectType

type ObjectTypeRes struct {
	Message ObjectTypeGlobal `json:"message,omitempty"`
}

type ObjectTypeGlobal struct {
	ObjectType   *[]ObjectType  `json:"ObjectType,omitempty"`
	Text         *[]Text        `json:"Text,omitempty"`
}

type ObjectType struct {
	ObjectType			string	`json:"ObjectType"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	ObjectType     		string  `json:"ObjectType"`
	Language          	string  `json:"Language"`
	ObjectTypeName		string 	`json:"ObjectTypeName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
