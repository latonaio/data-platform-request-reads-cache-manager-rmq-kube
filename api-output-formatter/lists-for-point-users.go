package apiOutputFormatter

type ListsForPointUsers struct {
	ContentList    	[]ContentList	`json:"ContentList"`
}

type ContentList struct {
	ListObject			int		`json:"ListObject"`
	ListObjectType		string	`json:"ListObjectType"`
	Introduction		*string	`json:"Introduction"`
	LocalSubRegion		*string `json:"LocalRegion"`
	LocalSubRegionName	*string `json:"LocalSubRegionName"`
	LocalRegion			*string `json:"LocalRegion"`
	LocalRegionName		*string `json:"LocalRegionName"`
	Tag1				*string	`json:"Tag1"`
	Tag2				*string	`json:"Tag2"`
	Tag3				*string	`json:"Tag3"`
	Tag4				*string	`json:"Tag4"`
}
