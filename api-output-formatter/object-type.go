package apiOutputFormatter

type ObjectType struct {
	ObjectTypeObjectType    []ObjectTypeObjectType    `json:"ObjectTypeObjectType"`
	ObjectTypeText          []ObjectTypeText          `json:"ObjectTypeText"`
	Accepter                []string                  `json:"Accepter"`
}

type ObjectTypeObjectType struct {
	ObjectType          string `json:"ObjectType"`
}

type ObjectTypeText struct {
	ObjectType          string `json:"ObjectType"`
	Language            string `json:"Language"`
	ObjectTypeName      string `json:"ObjectTypeName"`
}
