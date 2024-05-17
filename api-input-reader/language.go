package apiInputReader

type LanguageGlobal struct {
	Language     *Language
	LanguageText *LanguageText
}

type Language struct {
	Language            string `json:"Language"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type LanguageText struct {
	Language                string `json:"Language"`
	CorrespondenceLanguage  string `json:"CorrespondenceLanguage"`
	IsMarkedForDeletion     *bool  `json:"IsMarkedForDeletion"`
}
