package apiInputReader

type CountryGlobal struct {
	Country     *Country
	CountryText *CountryText
}

type Country struct {
	Country             string `json:"Country"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type CountryText struct {
	Country             string `json:"Country"`
	Language            string `json:"Language"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
