package apiModuleRuntimesResponsesLanguage

type LanguageRes struct {
	Message LanguageGlobal `json:"message,omitempty"`
}

type LanguageGlobal struct {
	Language    *[]Language    `json:"Language,omitempty"`
	Text        *[]Text        `json:"Text,omitempty"`
}

type Language struct {
	Language			string	`json:"Language"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	Language                  string   `json:"Language"`
	CorrespondenceLanguage    string   `json:"CorrespondenceLanguage"`
	LanguageName              string   `json:"LanguageName"`
	CreationDate              string   `json:"CreationDate"`
	LastChangeDate		      string   `json:"LastChangeDate"`
	IsMarkedForDeletion	      *bool    `json:"IsMarkedForDeletion"`
}
