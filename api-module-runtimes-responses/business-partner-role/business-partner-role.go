package apiModuleRuntimesResponsesBusinessPartnerRole

type BusinessPartnerRoleRes struct {
	Message BusinessPartnerRoleGlobal `json:"message,omitempty"`
}

type BusinessPartnerRoleGlobal struct {
	BusinessPartnerRole    *[]BusinessPartnerRole    `json:"BusinessPartnerRole,omitempty"`
	Text                   *[]Text                   `json:"Text,omitempty"`
}

type BusinessPartnerRole struct {
	BusinessPartnerRole	string	`json:"BusinessPartnerRole"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	BusinessPartnerRole     string  `json:"BusinessPartnerRole"`
	Language          	    string  `json:"Language"`
	BusinessPartnerRoleName	string 	`json:"BusinessPartnerRoleName"`
	CreationDate		    string	`json:"CreationDate"`
	LastChangeDate		    string	`json:"LastChangeDate"`
	IsMarkedForDeletion	    *bool	`json:"IsMarkedForDeletion"`
}
