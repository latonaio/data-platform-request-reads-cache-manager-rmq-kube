package apiInputReader

type PointConditionTypeGlobal struct {
	PointConditionType     *PointConditionType
	PointConditionTypeText *PointConditionTypeText
}

type PointConditionType struct {
	PointConditionType    string `json:"PointConditionType"`
	IsMarkedForDeletion   *bool  `json:"IsMarkedForDeletion"`
}

type PointConditionTypeText struct {
	PointConditionType     string `json:"PointConditionType"`
	Language               string `json:"Language"`
	IsMarkedForDeletion    *bool  `json:"IsMarkedForDeletion"`
}
