package apiInputReader

type WorkCenter struct {
	WorkCenterGeneral *WorkCenterGeneral
}

type WorkCenterGeneral struct {
	WorkCenter          int   `json:"WorkCenter"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
