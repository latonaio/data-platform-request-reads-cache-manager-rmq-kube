package apiInputReader

type Message struct {
	MessageHeader			*MessageHeader
	MessageDocHeaderDoc		*MessageDocHeaderDoc
}

type MessageHeader struct {
	Message    	        int		`json:"Message"`
	MessageType	        *string	`json:"MessageType"`
	Sender			    *int    `json:"Sender"`
	Receiver		    *int    `json:"Receiver"`
	MessageIsSent		*bool	`json:"MessageIsSent"`
	IsMarkedForDeletion *bool	`json:"IsMarkedForDeletion"`
}

type MessageDocHeaderDoc struct {
	Message            		 int     `json:"Message"`
	DocType                  *string `json:"DocType"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}
