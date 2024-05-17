package apiInputReader

type LocalRegionGlobal struct {
	LocalRegion     *LocalRegion
	LocalRegionText *LocalRegionText
}

type LocalRegion struct {
	LocalRegion         string `json:"LocalRegion"`
    Country             string `json:"Country"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type LocalRegionText struct {
	LocalRegion         string `json:"LocalRegion"`
    Country             string `json:"Country"`
	Language            string `json:"Language"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
