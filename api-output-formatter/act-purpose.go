package apiOutputFormatter

type ActPurpose struct {
	ActPurposeActPurpose  []ActPurposeActPurpose  `json:"ActPurposeActPurpose"`
	ActPurposeText        []ActPurposeText        `json:"ActPurposeText"`
	Accepter              []string                `json:"Accepter"`
}

type ActPurposeActPurpose struct {
	ActPurpose            string `json:"ActPurpose"`
}

type ActPurposeText struct {
	ActPurpose            string `json:"ActPurpose"`
	Language              string `json:"Language"`
	ActPurposeName        string `json:"ActPurposeName"`
}
