package apiInputReader

type IncotermsGlobal struct {
	Incoterms     *Incoterms
	IncotermsText *IncotermsText
}

type Incoterms struct {
	Incoterms           string `json:"Incoterms"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type IncotermsText struct {
	Incoterms           string `json:"Incoterms"`
	Language            string `json:"Language"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
