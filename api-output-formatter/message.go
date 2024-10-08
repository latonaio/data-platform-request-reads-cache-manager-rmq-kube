package apiOutputFormatter

type Message struct {
	MessageHeader    []MessageHeader    `json:"MessageHeader"`
	MountPath        *string            `json:"mount_path"`
	Accepter         []string           `json:"Accepter"`
}

type MessageHeader struct {
	Message				int		`json:"Message"`
	MessageType			string	`json:"MessageType"`
	MessageTypeName		string	`json:"MessageTypeName"`
	Sender				int		`json:"Sender"`
	SenderName			string	`json:"SenderName"`
	Receiver			int		`json:"Receiver"`
	ReceiverName		string 	`json:"ReceiverName"`
	Language			string	`json:"Language"`
	Title				string  `json:"Title"`
	LongText			string	`json:"LongText"`
	MessageIsRead		bool	`json:"MessageIsRead"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsCancelled			*bool	`json:"IsCancelled"`
	Images              Images  `json:"Images"`
}
