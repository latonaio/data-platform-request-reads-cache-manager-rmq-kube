package apiInputReader

type ObjectTypeGlobal struct {
	ObjectType     *ObjectType
	ObjectTypeText *ObjectTypeText
}

type ObjectType struct {
	ObjectType          string `json:"ObjectType"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type ObjectTypeText struct {
	ObjectType          string `json:"ObjectType"`
	Language            string `json:"Language"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
