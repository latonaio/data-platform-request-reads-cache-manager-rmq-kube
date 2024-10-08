package apiOutputFormatter

type Site struct {
	SiteHeader             []SiteHeader             `json:"SiteHeader"`
	SitePartner            []SitePartner            `json:"SitePartner"`
	SiteAddress            []SiteAddress            `json:"SiteAddress"`
	SiteAddressWithHeader  []SiteAddressWithHeader  `json:"SiteAddressWithHeader"`
	SitePartnerWithAddress []SitePartnerWithAddress `json:"SitePartnerWithAddress"`
	SiteCounter		   	   []SiteCounter			`json:"SiteCounter"`
	MountPath              *string                  `json:"mount_path"`
	Accepter               []string                 `json:"Accepter"`
}

type SiteHeader struct {
	Site                             int     `json:"Site"`
	SiteType                         string  `json:"SiteType"`
	SiteTypeName                     string  `json:"SiteTypeName"`
	SiteOwner                        *int    `json:"SiteOwner"`
	SiteOwnerName                    *string `json:"SiteOwnerName"`
	SiteOwnerBusinessPartnerRole     *string `json:"SiteOwnerBusinessPartnerRole"`
	SiteOwnerBusinessPartnerRoleName *string `json:"SiteOwnerBusinessPartnerRoleName"`
	Brand							 *int	 `json:"Brand"`
	BrandDescription				 *string `json:"BrandDescription"`
	PersonResponsible                string  `json:"PersonResponsible"`
	URL								 *string `json:"URL"`
	DailyOperationStartTime          string  `json:"DailyOperationStartTime"`
	DailyOperationEndTime            string  `json:"DailyOperationEndTime"`
	Description                      string  `json:"Description"`
	LongText                         string  `json:"LongText"`
	Introduction                     *string `json:"Introduction"`
	OperationRemarks                 *string `json:"OperationRemarks"`
	PhoneNumber						 *string `json:"PhoneNumber"`
	AvailabilityOfParking			 *bool	 `json:"AvailabilityOfParking"`
	NumberOfParkingSpaces			 *int	 `json:"NumberOfParkingSpaces"`
	SuperiorSite					 *int 	 `json:"SuperiorSite"`
	Tag1                             *string `json:"Tag1"`
	Tag2                             *string `json:"Tag2"`
	Tag3                             *string `json:"Tag3"`
	Tag4                             *string `json:"Tag4"`
	LastChangeDate					 string  `json:"LastChangeDate"`
	LastChangeTime					 string  `json:"LastChangeTime"`
	CreateUser						 int	 `json:"CreateUser"`
	CreateUserFullName				 *string `json:"CreateUserFullName"`
	CreateUserNickName				 *string `json:"CreateUserNickName"`
	LastChangeUser					 int 	 `json:"LastChangeUser"`
	LastChangeUserFullName			 *string `json:"LastChangeUserFullName"`
	LastChangeUserNickName			 *string `json:"LastChangeUserNickName"`
	NumberOfLikes					 *int	 `json:"NumberOfLikes"`
	Images                           Images  `json:"Images"`
}

type SitePartner struct {
	Site                    int     `json:"Site"`
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
	CityName                *string `json:"CityName"`
	StreetName              *string `json:"StreetName"`
	PostalCode              *string `json:"PostalCode"`
	AddressIdentifier       *string `json:"AddressIdentifier"`
	LocalRegionName         *string `json:"LocalRegionName"`
}

type SiteAddress struct {
	Site				int		`json:"Site"`
	AddressID			int		`json:"AddressID"`
	PostalCode  		string 	`json:"PostalCode"`
	LocalSubRegion		string  `json:"LocalSubRegion"`
	LocalSubRegionName	string  `json:"LocalSubRegionName"`
	LocalRegion			string  `json:"LocalRegion"`
	LocalRegionName		string  `json:"LocalRegionName"`
	Country				*string `json:"Country"`
	GlobalRegion		string 	`json:"GlobalRegion"`
	TimeZone			string 	`json:"TimeZone"`
	District			*string `json:"District"`
	StreetName			*string `json:"StreetName"`
	CityName			*string `json:"CityName"`
	Building			*string `json:"Building"`
	Floor				*int	`json:"Floor"`
	Room				*int	`json:"Room"`
	XCoordinate			*float32 `json:"XCoordinate"`
	YCoordinate			*float32 `json:"YCoordinate"`
	ZCoordinate			*float32 `json:"ZCoordinate"`
}

type SiteCounter struct {
	Site					int		`json:"Site"`
	NumberOfLikes			int		`json:"NumberOfLikes"`
}

type SiteAddressWithHeader struct {
	Site              int     `json:"Site"`
	AddressID         int     `json:"AddressID"`
	LocalSubRegion    string  `json:"LocalSubRegion"`
	LocalRegion       string  `json:"LocalRegion"`
	SiteType          string  `json:"SiteType"`
	SiteOwner         int     `json:"SiteOwner"`
	SiteOwnerName     string  `json:"SiteOwnerName"`
	ValidityStartDate string  `json:"ValidityStartDate"`
	ValidityStartTime string  `json:"ValidityStartTime"`
	ValidityEndDate   string  `json:"ValidityEndDate"`
	ValidityEndTime   string  `json:"ValidityEndTime"`
	Description       string  `json:"Description"`
	LongText          string  `json:"LongText"`
	Introduction      *string `json:"Introduction"`
	Tag1              *string `json:"Tag1"`
	Tag2              *string `json:"Tag2"`
	Tag3              *string `json:"Tag3"`
	Tag4              *string `json:"Tag4"`
	LastChangeDate	  string  `json:"LastChangeDate"`
	LastChangeTime	  string  `json:"LastChangeTime"`
	NumberOfLikes	  *int	  `json:"NumberOfLikes"`
	PostalCode  	  string  `json:"PostalCode"`
	Country     	  string  `json:"Country"`
	GlobalRegion   	  string  `json:"GlobalRegion"`
	TimeZone   		  string  `json:"TimeZone"`
	District    	  *string `json:"District"`
	StreetName  	  *string `json:"StreetName"`
	CityName    	  *string `json:"CityName"`
	Building    	  *string `json:"Building"`
	Floor       	  *int	  `json:"Floor"`
	Room        	  *int	  `json:"Room"`
	XCoordinate 	  *float32 `json:"XCoordinate"`
	YCoordinate 	  *float32 `json:"YCoordinate"`
	ZCoordinate 	  *float32 `json:"ZCoordinate"`
	Images            Images  `json:"Images"`
}

type SitePartnerWithAddress struct {
	Site                    int     `json:"Site"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	AddressID               *int    `json:"AddressID"`
	PostalCode  	  		string  `json:"PostalCode"`
	Country     	  		string  `json:"Country"`
	GlobalRegion   	  		string  `json:"GlobalRegion"`
	TimeZone   		  		string  `json:"TimeZone"`
	District    	  		*string `json:"District"`
	StreetName  	  		*string `json:"StreetName"`
	CityName    	  		*string `json:"CityName"`
	Building    	  		*string `json:"Building"`
	Floor       	  		*int    `json:"Floor"`
	Room        	  		*int    `json:"Room"`
	XCoordinate 	  		*float32 `json:"XCoordinate"`
	YCoordinate 	  		*float32 `json:"YCoordinate"`
	ZCoordinate 	  		*float32 `json:"ZCoordinate"`
}
