package apiOutputFormatter

type PointConditionType struct {
	PointConditionTypePointConditionType    []PointConditionTypePointConditionType    `json:"PointConditionTypePointConditionType"`
	PointConditionTypeText                  []PointConditionTypeText                  `json:"PointConditionTypeText"`
	Accepter                                []string                                  `json:"Accepter"`
}

type PointConditionTypePointConditionType struct {
	PointConditionType        string	`json:"PointConditionType"`
}

type PointConditionTypeText struct {
	PointConditionType        string `json:"PointConditionType"`
	Language                  string `json:"Language"`
	PointConditionTypeName    string `json:"PointConditionTypeName"`
}
