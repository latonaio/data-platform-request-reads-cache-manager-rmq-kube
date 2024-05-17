package apiInputReader

type ActPurposeGlobal struct {
	ActPurpose     *ActPurpose
	ActPurposeText *ActPurposeText
}

type ActPurpose struct {
	ActPurpose          string `json:"ActPurpose"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type ActPurposeText struct {
	ActPurpose          string `json:"ActPurpose"`
	Language            string `json:"Language"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
