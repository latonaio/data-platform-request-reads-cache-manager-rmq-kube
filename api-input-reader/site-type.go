package apiInputReader

type SiteTypeGlobal struct {
	SiteType     *SiteType
	SiteTypeText *SiteTypeText
}

type SiteType struct {
	SiteType            string `json:"SiteType"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type SiteTypeText struct {
	SiteType            string `json:"SiteType"`
	Language            string `json:"Language"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
