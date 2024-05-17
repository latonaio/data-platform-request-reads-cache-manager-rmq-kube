package apiOutputFormatter

type SiteType struct {
	SiteTypeSiteType    []SiteTypeSiteType    `json:"SiteTypeSiteType"`
	SiteTypeText        []SiteTypeText        `json:"SiteTypeText"`
	Accepter            []string              `json:"Accepter"`
}

type SiteTypeSiteType struct {
	SiteType            string	`json:"SiteType"`
}

type SiteTypeText struct {
	SiteType            string `json:"SiteType"`
	Language            string `json:"Language"`
	SiteTypeName        string `json:"SiteTypeName"`
}
