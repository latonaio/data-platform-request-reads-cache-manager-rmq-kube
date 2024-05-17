package apiOutputFormatter

type Language struct {
	LanguageLanguage    []LanguageLanguage    `json:"LanguageLanguage"`
	LanguageText        []LanguageText        `json:"LanguageText"`
	Accepter            []string              `json:"Accepter"`
}

type LanguageLanguage struct {
	Language            string	`json:"Language"`
}

type LanguageText struct {
	Language                  string `json:"Language"`
	CorrespondenceLanguage    string `json:"CorrespondenceLanguage"`
	LanguageName              string `json:"LanguageName"`
}
