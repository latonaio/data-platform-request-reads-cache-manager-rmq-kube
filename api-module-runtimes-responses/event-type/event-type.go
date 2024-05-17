package apiModuleRuntimesResponsesEventType

type EventTypeRes struct {
	Message EventTypeGlobal `json:"message,omitempty"`
}

type EventTypeGlobal struct {
	EventType    *[]EventType    `json:"EventType,omitempty"`
	Text         *[]Text         `json:"Text,omitempty"`
}

type EventType struct {
	EventType			string	`json:"EventType"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	EventType     		string  `json:"EventType"`
	Language          	string  `json:"Language"`
	EventTypeName		string 	`json:"EventTypeName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
