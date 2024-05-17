package apiOutputFormatter

type EventType struct {
	EventTypeEventType    []EventTypeEventType    `json:"EventTypeEventType"`
	EventTypeText         []EventTypeText         `json:"EventTypeText"`
	Accepter              []string                `json:"Accepter"`
}

type EventTypeEventType struct {
	EventType            string	`json:"EventType"`
}

type EventTypeText struct {
	EventType            string `json:"EventType"`
	Language             string `json:"Language"`
	EventTypeName        string `json:"EventTypeName"`
}
