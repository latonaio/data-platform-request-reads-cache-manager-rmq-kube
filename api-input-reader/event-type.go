package apiInputReader

type EventTypeGlobal struct {
	EventType     *EventType
	EventTypeText *EventTypeText
}

type EventType struct {
	EventType           string `json:"EventType"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type EventTypeText struct {
	EventType           string `json:"EventType"`
	Language            string `json:"Language"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
