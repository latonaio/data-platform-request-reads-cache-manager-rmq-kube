package apiInputReader

type Message struct {
	MessageHeader    *MessageHeader
}

type MessageHeader struct {
	Message    	        int		`json:"Message"`
	MessageType	        *string	`json:"MessageType"`
	Sender			    *int    `json:"Sender"`
	Receiver		    *int    `json:"Receiver"`
	MessageIsSent		*bool	`json:"MessageIsSent"`
	IsMarkedForDeletion *bool	`json:"IsMarkedForDeletion"`
}
