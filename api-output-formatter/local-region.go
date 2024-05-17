package apiOutputFormatter

type LocalRegion struct {
	LocalRegionLocalRegion    []LocalRegionLocalRegion    `json:"LocalRegionLocalRegion"`
	LocalRegionText           []LocalRegionText           `json:"LocalRegionText"`
	Accepter                  []string                    `json:"Accepter"`
}

type LocalRegionLocalRegion struct {
	LocalRegion    string    `json:"LocalRegion"`
	Country        string    `json:"Country"`
}

type LocalRegionText struct {
	LocalRegion        string `json:"LocalRegion"`
	Country            string `json:"Country"`
	Language           string `json:"Language"`
	LocalRegionName    string `json:"LocalRegionName"`
}
