package apiOutputFormatter

type BusinessPartner struct {
	BusinessPartnerGeneral       []BusinessPartnerGeneral       `json:"Generals"`
	BusinessPartnerGeneralDoc    []BusinessPartnerGeneralDoc    `json:"GeneralDoc"`
	BusinessPartnerPerson        []BusinessPartnerPerson        `json:"Person"`
	BusinessPartnerAddress       []BusinessPartnerAddress       `json:"Address"`
	BusinessPartnerGPS           []BusinessPartnerGPS           `json:"GPS"`
	BusinessPartnerRank          []BusinessPartnerRank          `json:"Rank"`
	BusinessPartnerDetailGeneral []BusinessPartnerDetailGeneral `json:"DetailGeneral"`
}

type BusinessPartnerGeneral struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerName string `json:"BusinessPartnerName"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type BusinessPartnerGeneralDoc struct {
	BusinessPartner          int	`json:"BusinessPartner"`
	BusinessPartnerName		 string `json:"BusinessPartnerName"`
	DocType                  string `json:"DocType"`
	FileExtension            string `json:"FileExtension"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
	FilePath                 string `json:"FilePath"`
	FileName                 string `json:"FileName"`
	Images                   Images `json:"Images"`
}

type BusinessPartnerPerson struct {
	BusinessPartner				int		`json:"BusinessPartner"`
	BusinessPartnerType			string	`json:"BusinessPartnerType"`
	FirstName					*string	`json:"FirstName"`
	LastName					*string	`json:"LastName"`
	FullName					*string	`json:"FullName"`
	MiddleName					*string	`json:"MiddleName"`
	NickName					string	`json:"NickName"`
	Gender						string	`json:"Gender"`
	Language					string	`json:"Language"`
	CorrespondenceLanguage		*string	`json:"CorrespondenceLanguage"`
	BirthDate					*string	`json:"BirthDate"`
	Nationality					string	`json:"Nationality"`
	EmailAddress				*string	`json:"EmailAddress"`
	MobilePhoneNumber			*string	`json:"MobilePhoneNumber"`
	ProfileComment				*string	`json:"ProfileComment"`
	PreferableLocalSubRegion	string  `json:"PreferableLocalSubRegion"`
	PreferableLocalRegion		string  `json:"PreferableLocalRegion"`
	PreferableCountry			string  `json:"PreferableCountry"`
	ActPurpose					string  `json:"ActPurpose"`
}

type BusinessPartnerAddress struct {
	BusinessPartner    int     `json:"BusinessPartner"`
	AddressID          int     `json:"AddressID"`
	PostalCode         *string `json:"PostalCode"`
	LocalSubRegion     string  `json:"LocalSubRegion"`
	LocalSubRegionName string  `json:"LocalSubRegionName"`
	LocalRegion        string  `json:"LocalRegion"`
	LocalRegionName    string  `json:"LocalRegionName"`
	Country            *string `json:"Country"`
	StreetName         *string `json:"StreetName"`
	CityName           *string `json:"CityName"`
	Building           *string `json:"Building"`
	Floor              *int    `json:"Floor"`
	Room               *int    `json:"Room"`
	Site               *int    `json:"Site"`
}

type BusinessPartnerGPS struct {
	BusinessPartner     int		`json:"BusinessPartner"`
	XCoordinate     	float32	`json:"XCoordinate"`
	YCoordinate     	float32	`json:"YCoordinate"`
	ZCoordinate     	float32	`json:"ZCoordinate"`
	LocalSubRegion  	string	`json:"LocalSubRegion"`
	LocalRegion     	string	`json:"LocalRegion"`
	Country         	string	`json:"Country"`
	LastChangeDate      string	`json:"LastChangeDate"`
	LastChangeTime      string	`json:"LastChangeTime"`
}

type BusinessPartnerRank struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	RankType            string `json:"RankType"`
	Rank                int    `json:"Rank"`
	ValidityStartDate   string `json:"ValidityStartDate"`
	ValidityEndDate     string `json:"ValidityEndDate"`
}

type BusinessPartnerDetailGeneral struct {
	BusinessPartnerFullName  *string `json:"BusinessPartnerFullName"`
	Industry                 *string `json:"Industry"`
	LegalEntityRegistration  *string `json:"LegalEntityRegistration"`
	Country                  string  `json:"Country"`
	Language                 *string `json:"Language"`
	Currency                 *string `json:"Currency"`
	Representative           *string `json:"Representative"`
	PhoneNumber           	 *string `json:"PhoneNumber"`
	AddressID                *int    `json:"AddressID"`
	BusinessPartnerIsBlocked *bool   `json:"BusinessPartnerIsBlocked"`
	CreationDate             string  `json:"CreationDate"`
	LastChangeDate           string  `json:"LastChangeDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}
