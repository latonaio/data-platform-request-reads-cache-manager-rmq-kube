package apiInputReader

type ProjectGlobal struct {
	Project           *Project
	ProjectWBSElement *ProjectWBSElement
}

type Project struct {
	Project             int   `json:"Project"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type ProjectWBSElement struct {
	Project             int   `json:"Project"`
	WBSElement          int   `json:"WBSElement"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}
