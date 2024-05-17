package apiInputReader

type MessageTypeGlobal struct {
	MessageType     *MessageType
	MessageTypeText *MessageTypeText
}

type MessageType struct {
	MessageType            string `json:"MessageType"`
	IsMarkedForDeletion    *bool  `json:"IsMarkedForDeletion"`
}

type MessageTypeText struct {
	MessageType            string `json:"MessageType"`
	Language               string `json:"Language"`
	IsMarkedForDeletion    *bool  `json:"IsMarkedForDeletion"`
}
