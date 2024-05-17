package apiModuleRuntimesResponsesDistributionProfile

type DistributionProfileRes struct {
	Message DistributionProfileGlobal `json:"message,omitempty"`
}

type DistributionProfileGlobal struct {
	DistributionProfile    *[]DistributionProfile    `json:"DistributionProfile,omitempty"`
	Text                   *[]Text                   `json:"Text,omitempty"`
}

type DistributionProfile struct {
	DistributionProfile string	`json:"DistributionProfile"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type Text struct {
	DistributionProfile        string  `json:"DistributionProfile"`
	Language          	       string  `json:"Language"`
	DistributionProfileName	   string  `json:"DistributionProfileName"`
	CreationDate		       string  `json:"CreationDate"`
	LastChangeDate		       string  `json:"LastChangeDate"`
	IsMarkedForDeletion	       *bool   `json:"IsMarkedForDeletion"`
}
