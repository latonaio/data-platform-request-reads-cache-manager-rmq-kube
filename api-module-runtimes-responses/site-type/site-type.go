package apiModuleRuntimesResponsesSiteType

type SiteTypeRes struct {
	Message SiteTypeGlobal `json:"message,omitempty"`
}

type SiteTypeGlobal struct {
	SiteType     *[]SiteType    `json:"SiteType,omitempty"`
	Text         *[]Text        `json:"Text,omitempty"`
}

type SiteType struct {
	SiteType			string	`json:"SiteType"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	SiteType     		string  `json:"SiteType"`
	Language          	string  `json:"Language"`
	SiteTypeName		string 	`json:"SiteTypeName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
