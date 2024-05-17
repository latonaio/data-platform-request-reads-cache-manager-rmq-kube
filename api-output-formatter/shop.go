package apiOutputFormatter

type Shop struct {
	ShopHeader             []ShopHeader             `json:"ShopHeader"`
	ShopPartner            []ShopPartner            `json:"ShopPartner"`
	ShopAddress            []ShopAddress            `json:"ShopAddress"`
	ShopAddressWithHeader  []ShopAddressWithHeader  `json:"ShopAddressWithHeader"`
	ShopPartnerWithAddress []ShopPartnerWithAddress `json:"ShopPartnerWithAddress"`
	MountPath              *string                  `json:"mount_path"`
	Accepter               []string                 `json:"Accepter"`
}

type ShopHeader struct {
	Shop							int		`json:"Shop"`
	ShopType						string	`json:"ShopType"`
	ShopOwner						int		`json:"ShopOwner"`
	ShopOwnerBusinessPartnerRole	string	`json:"ShopOwnerBusinessPartnerRole"`
	Brand							*int	`json:"Brand"`
	PersonResponsible				string	`json:"PersonResponsible"`
	URL								*string	`json:"URL"`
	ValidityStartDate				string	`json:"ValidityStartDate"`
	ValidityStartTime				string	`json:"ValidityStartTime"`
	ValidityEndDate					string	`json:"ValidityEndDate"`
	ValidityEndTime					string	`json:"ValidityEndTime"`
	DailyOperationStartTime			string	`json:"DailyOperationStartTime"`
	DailyOperationEndTime			string	`json:"DailyOperationEndTime"`
	Description						string	`json:"Description"`
	LongText						string	`json:"LongText"`
	Introduction					*string	`json:"Introduction"`
	OperationRemarks				*string	`json:"OperationRemarks"`
	PhoneNumber						*string	`json:"PhoneNumber"`
	AvailabilityOfParking			*bool	`json:"AvailabilityOfParking"`
	NumberOfParkingSpaces			*int	`json:"NumberOfParkingSpaces"`
	Site							int		`json:"Site"`
	Project							*int	`json:"Project"`
	WBSElement						*int	`json:"WBSElement"`
	Tag1							*string	`json:"Tag1"`
	Tag2							*string	`json:"Tag2"`
	Tag3							*string	`json:"Tag3"`
	Tag4							*string	`json:"Tag4"`
	PointConsumptionType      		string  `json:"PointConsumptionType"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	CreateUser						int		`json:"CreateUser"`
	LastChangeUser					int		`json:"LastChangeUser"`
	IsReleased						*bool	`json:"IsReleased"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
	Images                           Images  `json:"Images"`
}

type ShopPartner struct {
	Shop                    int     `json:"Shop"`
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

type ShopAddress struct {
	Shop				int		`json:"Shop"`
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
	Site			    *int	 `json:"Site"`
}

type ShopAddressWithHeader struct {
	Shop              int     `json:"Shop"`
	AddressID         int     `json:"AddressID"`
	LocalSubRegion    string  `json:"LocalSubRegion"`
	LocalRegion       string  `json:"LocalRegion"`
	ShopType          string  `json:"ShopType"`
	ShopOwner         int     `json:"ShopOwner"`
	ShopOwnerName     string  `json:"ShopOwnerName"`
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

type ShopPartnerWithAddress struct {
	Shop                    int     `json:"Shop"`
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
	Site			        *int	 `json:"Site"`
}
