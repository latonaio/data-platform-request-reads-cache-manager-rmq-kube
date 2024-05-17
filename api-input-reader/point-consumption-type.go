package apiInputReader

type PointConsumptionTypeGlobal struct {
	PointConsumptionType     *PointConsumptionType
	PointConsumptionTypeText *PointConsumptionTypeText
}

type PointConsumptionType struct {
	PointConsumptionType   string `json:"PointConsumptionType"`
	IsMarkedForDeletion    *bool  `json:"IsMarkedForDeletion"`
}

type PointConsumptionTypeText struct {
	PointConsumptionType    string `json:"PointConsumptionType"`
	Language                string `json:"Language"`
	IsMarkedForDeletion     *bool  `json:"IsMarkedForDeletion"`
}
