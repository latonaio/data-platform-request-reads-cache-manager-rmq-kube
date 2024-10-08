package apiOutputFormatter

type BusinessPartner struct {
	BusinessPartnerGeneral       			[]BusinessPartnerGeneral			`json:"Generals"`
	BusinessPartnerGeneralDoc    			[]BusinessPartnerGeneralDoc			`json:"GeneralDoc"`
	BusinessPartnerBPRole        			[]BusinessPartnerBPRole				`json:"Role"`
	BusinessPartnerPerson        			[]BusinessPartnerPerson				`json:"Person"`
	BusinessPartnerAddress       			[]BusinessPartnerAddress			`json:"Address"`
	BusinessPartnerSNS           			[]BusinessPartnerSNS				`json:"SNS"`
	BusinessPartnerGPS           			[]BusinessPartnerGPS				`json:"GPS"`
	BusinessPartnerRank          			[]BusinessPartnerRank				`json:"Rank"`
	BusinessPartnerPersonOrganization		[]BusinessPartnerPersonOrganization	`json:"PersonOrganization"`
	BusinessPartnerDetailGeneral			[]BusinessPartnerDetailGeneral		`json:"DetailGeneral"`
}

type BusinessPartnerGeneral struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerName string `json:"BusinessPartnerName"`
	Withdrawal          *bool  `json:"Withdrawal"`
	IsReleased          *bool  `json:"IsReleased"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
	Images              Images `json:"Images"`
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

type BusinessPartnerBPRole struct {
	BusinessPartner			int    `json:"BusinessPartner"`
	BusinessPartnerRole		string `json:"BusinessPartnerRole"`
	BusinessPartnerRoleName string `json:"BusinessPartnerRoleName"`
}

type BusinessPartnerPerson struct {
	BusinessPartner					int		`json:"BusinessPartner"`
	BusinessPartnerType				string	`json:"BusinessPartnerType"`
	FirstName						string	`json:"FirstName"`
	LastName						string	`json:"LastName"`
	FullName						string	`json:"FullName"`
	MiddleName						*string	`json:"MiddleName"`
	NickName						string	`json:"NickName"`
	Gender							string	`json:"Gender"`
	Language						string	`json:"Language"`
	LanguageName					string	`json:"LanguageName"`
	CorrespondenceLanguage			*string	`json:"CorrespondenceLanguage"`
	BirthDate						string	`json:"BirthDate"`
	Nationality						string	`json:"Nationality"`
	NationalityName					string	`json:"NationalityName"`
	EmailAddress					*string	`json:"EmailAddress"`
	MobilePhoneNumber				*string	`json:"MobilePhoneNumber"`
	ProfileComment					*string	`json:"ProfileComment"`
	PreferableLocalSubRegion		string  `json:"PreferableLocalSubRegion"`
	PreferableLocalSubRegionName	string  `json:"PreferableLocalSubRegionName"`
	PreferableLocalRegion			string  `json:"PreferableLocalRegion"`
	PreferableLocalRegionName		string  `json:"PreferableLocalRegionName"`
	PreferableCountry				string  `json:"PreferableCountry"`
	ActPurpose						string  `json:"ActPurpose"`
	ActPurposeName					string  `json:"ActPurposeName"`
	TermsOfUseIsConfirmed			*bool   `json:"TermsOfUseIsConfirmed"`
	Images                  		Images  `json:"Images"`
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

type BusinessPartnerSNS struct {
	BusinessPartner     int		`json:"BusinessPartner"`
	XURL  				*string	`json:"XURL"`
	InstagramURL     	*string	`json:"InstagramURL"`
	TikTokURL         	*string	`json:"TikTokURL"`
	PointAppsURL        string	`json:"PointAppsURL"`
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
	RankName			string `json:"RankName"`
	ValidityStartDate   string `json:"ValidityStartDate"`
	ValidityEndDate     string `json:"ValidityEndDate"`
}

type BusinessPartnerPersonOrganization struct {
	BusinessPartner        			int     `json:"BusinessPartner"`
	BusinessPartnerType    			string  `json:"BusinessPartnerType"`
	OrganizationBusinessPartner		int		`json:"OrganizationBusinessPartner"`
	OrganizationBusinessPartnerName	string  `json:"OrganizationBusinessPartnerName"`
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
