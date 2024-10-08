package apiModuleRuntimesResponsesArticle

type ArticleRes struct {
	Message Article `json:"message,omitempty"`
}

type Article struct {
	Header                *[]Header                `json:"Header"`
	Partner               *[]Partner               `json:"Partner"`
	Address               *[]Address               `json:"Address"`
	Counter               *[]Counter               `json:"Counter"`
}

type Header struct {
	Article                         int     `json:"Article"`
	ArticleType                     string  `json:"ArticleType"`
	ArticleOwner                    int     `json:"ArticleOwner"`
	ArticleOwnerBusinessPartnerRole string  `json:"ArticleOwnerBusinessPartnerRole"`
	PersonResponsible               string  `json:"PersonResponsible"`
	ValidityStartDate               string  `json:"ValidityStartDate"`
	ValidityStartTime               string  `json:"ValidityStartTime"`
	ValidityEndDate                 string  `json:"ValidityEndDate"`
	ValidityEndTime                 string  `json:"ValidityEndTime"`
	Description                     string  `json:"Description"`
	LongText                        string  `json:"LongText"`
	Introduction                    *string `json:"Introduction"`
	Site                            int     `json:"Site"`
	Shop						    *int	  `json:"Shop"`
	Project                         *int    `json:"Project"`
	WBSElement                      *int    `json:"WBSElement"`
	Tag1                            *string `json:"Tag1"`
	Tag2                            *string `json:"Tag2"`
	Tag3                            *string `json:"Tag3"`
	Tag4                            *string `json:"Tag4"`
	DistributionProfile             string  `json:"DistributionProfile"`
	QuestionnaireType			    *string `json:"QuestionnaireType"`
	QuestionnaireTemplate		    *string `json:"QuestionnaireTemplate"`
	CreationDate                    string  `json:"CreationDate"`
	CreationTime                    string  `json:"CreationTime"`
	LastChangeDate                  string  `json:"LastChangeDate"`
	LastChangeTime                  string  `json:"LastChangeTime"`
	CreateUser					    int	  `json:"CreateUser"`
	LastChangeUser				    int	  `json:"LastChangeUser"`
	IsReleased                      *bool   `json:"IsReleased"`
	IsMarkedForDeletion             *bool   `json:"IsMarkedForDeletion"`
}

type Partner struct {
	Article                 int     `json:"Article"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
}

type Address struct {
	Article        int      `json:"Article"`
	AddressID      int      `json:"AddressID"`
	PostalCode     string   `json:"PostalCode"`
	LocalSubRegion string   `json:"LocalSubRegion"`
	LocalRegion    string   `json:"LocalRegion"`
	Country        string   `json:"Country"`
	GlobalRegion   string   `json:"GlobalRegion"`
	TimeZone       string   `json:"TimeZone"`
	District       *string  `json:"District"`
	StreetName     *string  `json:"StreetName"`
	CityName       *string  `json:"CityName"`
	Building       *string  `json:"Building"`
	Floor          *int     `json:"Floor"`
	Room           *int     `json:"Room"`
	XCoordinate    *float32 `json:"XCoordinate"`
	YCoordinate    *float32 `json:"YCoordinate"`
	ZCoordinate    *float32 `json:"ZCoordinate"`
	Site           *int     `json:"Site"`
}

type Counter struct {
	Article					int		`json:"Article"`
	NumberOfLikes			int		`json:"NumberOfLikes"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	LastChangeDate			string	`json:"LastChangeDate"`
	LastChangeTime			string	`json:"LastChangeTime"`
}
