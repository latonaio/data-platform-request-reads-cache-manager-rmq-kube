package apiModuleRuntimesResponsesMessage

type MessageRes struct {
	Message Message `json:"message,omitempty"`
}

type Message struct {
	Header *[]Header `json:"Header"`
}

type Header struct {
	Message				int		`json:"Message"`
	MessageType			string	`json:"MessageType"`
	Sender				int		`json:"Sender"`
	Receiver			int		`json:"Receiver"`
	Language			string	`json:"Language"`
	Title				string	`json:"Title"`
	LongText			string	`json:"LongText"`
	MessageIsSent		bool	`json:"MessageIsSent"`
	MessageIsRead		bool	`json:"MessageIsRead"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsCancelled			*bool	`json:"IsCancelled"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
