package apiModuleRuntimesResponsesProject

type ProjectRes struct {
	Message ProjectGlobal `json:"message,omitempty"`
}

type ProjectGlobal struct {
	Project    *[]Project    `json:"Project,omitempty"`
	WBSElement *[]WBSElement `json:"WBSElement,omitempty"`
}

type Project struct {
	Project               int     `json:"Project"`
	ProjectCode           string  `json:"ProjectCode"`
	ProjectDescription    string  `json:"ProjectDescription"`
	OwnerBusinessPartner  int     `json:"OwnerBusinessPartner"`
	OwnerPlant            string  `json:"OwnerPlant"`
	ProjectProfile        *string `json:"ProjectProfile"`
	ResponsiblePerson     *int    `json:"ResponsiblePerson"`
	ResponsiblePersonName *string `json:"ResponsiblePersonName"`
	PlannedStartDate      *string `json:"PlannedStartDate"`
	PlannedEndDate        *string `json:"PlannedEndDate"`
	ActualStartDate       *string `json:"ActualStartDate"`
	ActualEndDate         *string `json:"ActualEndDate"`
	CreationDate          string  `json:"CreationDate"`
	LastChangeDate        string  `json:"LastChangeDate"`
	IsMarkedForDeletion   *bool   `json:"IsMarkedForDeletion"`
}

type WBSElement struct {
	Project               int     `json:"Project"`
	WBSElement            int     `json:"WBSElement"`
	WBSElementCode        string  `json:"WBSElementCode"`
	WBSElementDescription string  `json:"WBSElementDescription"`
	BusinessPartner       int     `json:"BusinessPartner"`
	Plant                 string  `json:"Plant"`
	ResponsiblePerson     *int    `json:"ResponsiblePerson"`
	ResponsiblePersonName *string `json:"ResponsiblePersonName"`
	PlannedStartDate      *string `json:"PlannedStartDate"`
	PlannedEndDate        *string `json:"PlannedEndDate"`
	ActualStartDate       *string `json:"ActualStartDate"`
	ActualEndDate         *string `json:"ActualEndDate"`
	CreationDate          string  `json:"CreationDate"`
	LastChangeDate        string  `json:"LastChangeDate"`
	IsMarkedForDeletion   *bool   `json:"IsMarkedForDeletion"`
}
