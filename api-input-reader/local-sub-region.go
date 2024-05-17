package apiInputReader

type LocalSubRegionGlobal struct {
	LocalSubRegion     *LocalSubRegion
	LocalSubRegionText *LocalSubRegionText
}

type LocalSubRegion struct {
	LocalSubRegion      string `json:"LocalSubRegion"`
	LocalRegion         string `json:"LocalRegion"`
    Country             string `json:"Country"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type LocalSubRegionText struct {
	LocalSubRegion      string `json:"LocalSubRegion"`
	LocalRegion         string `json:"LocalRegion"`
    Country             string `json:"Country"`
	Language            string `json:"Language"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
