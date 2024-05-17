package apiOutputFormatter

type Country struct {
	CountryCountry    []CountryCountry    `json:"CountryCountry"`
	CountryText       []CountryText       `json:"CountryText"`
	Accepter          []string            `json:"Accepter"`
}

type CountryCountry struct {
	Country        string `json:"Country"`
}

type CountryText struct {
	Country        string `json:"Country"`
	Language       string `json:"Language"`
	CountryName    string `json:"CountryName"`
}
