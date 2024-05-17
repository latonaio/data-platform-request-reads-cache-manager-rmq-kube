package apiModuleRuntimesResponsesMessageType

type MessageTypeRes struct {
	Message MessageTypeGlobal `json:"message,omitempty"`
}

type MessageTypeGlobal struct {
	MessageType    *[]MessageType     `json:"MessageType,omitempty"`
	Text           *[]Text            `json:"Text,omitempty"`
}

type MessageType struct {
	MessageType             string	`json:"MessageType"`
	CreationDate		    string	`json:"CreationDate"`
	LastChangeDate		    string	`json:"LastChangeDate"`
	IsMarkedForDeletion	    *bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	MessageType             string  `json:"MessageType"`
	Language                string  `json:"Language"`
	MessageTypeName         string  `json:"MessageTypeName"`
	CreationDate            string  `json:"CreationDate"`
	LastChangeDate          string  `json:"LastChangeDate"`
	IsMarkedForDeletion     *bool   `json:"IsMarkedForDeletion"`
}
