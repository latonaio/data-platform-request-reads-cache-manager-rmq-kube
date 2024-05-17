package apiOutputFormatter

type LocalSubRegion struct {
	LocalSubRegionLocalSubRegion    []LocalSubRegionLocalSubRegion    `json:"LocalSubRegionLocalSubRegion"`
	LocalSubRegionText              []LocalSubRegionText              `json:"LocalSubRegionText"`
	Accepter                        []string                          `json:"Accepter"`
}

type LocalSubRegionLocalSubRegion struct {
	LocalSubRegion    string    `json:"LocalSubRegion"`
	LocalRegion       string    `json:"LocalRegion"`
	Country           string    `json:"Country"`
}

type LocalSubRegionText struct {
	LocalSubRegion        string `json:"LocalSubRegion"`
	LocalRegion           string `json:"LocalRegion"`
	Country               string `json:"Country"`
	Language              string `json:"Language"`
	LocalSubRegionName    string `json:"LocalSubRegionName"`
}
