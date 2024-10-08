package apiOutputFormatter

type Article struct {
	ArticleHeader             	[]ArticleHeader             	`json:"ArticleHeader"`
	ArticlePartner            	[]ArticlePartner            	`json:"ArticlePartner"`
	ArticleAddress            	[]ArticleAddress            	`json:"ArticleAddress"`
	ArticleAddressWithHeader  	[]ArticleAddressWithHeader  	`json:"ArticleAddressWithHeader"`
	ArticlePartnerWithAddress 	[]ArticlePartnerWithAddress 	`json:"ArticlePartnerWithAddress"`
	ArticleCounter				[]ArticleCounter				`json:"ArticleCounter"`
	MountPath               	*string                   		`json:"mount_path"`
	Accepter                	[]string                  		`json:"Accepter"`
}

type ArticleHeader struct {
	Article                   			int     `json:"Article"`
	ArticleType               			string  `json:"ArticleType"`
	ArticleTypeName           			string  `json:"ArticleTypeName"`
	ArticleOwner              			int     `json:"ArticleOwner"`
	ArticleOwnerName          			string  `json:"ArticleOwnerName"`
	ArticleOwnerBusinessPartnerRole		string	`json:"ArticleOwnerBusinessPartnerRole"`
	ArticleOwnerBusinessPartnerRoleName	string	`json:"ArticleOwnerBusinessPartnerRoleName"`
	PersonResponsible       			string  `json:"PersonResponsible"`
	ValidityStartDate       			string  `json:"ValidityStartDate"`
	ValidityStartTime       			string  `json:"ValidityStartTime"`
	ValidityEndDate         			string  `json:"ValidityEndDate"`
	ValidityEndTime         			string  `json:"ValidityEndTime"`
	Description             			string  `json:"Description"`
	LongText                			string  `json:"LongText"`
	Introduction            			*string `json:"Introduction"`
	Site                    			int     `json:"Site"`
	SiteDescription            			string  `json:"SiteDescription"`
	Shop                    			*int    `json:"Shop"`
	ShopDescription            			*string `json:"ShopDescription"`
	Tag1                    			*string `json:"Tag1"`
	Tag2                    			*string `json:"Tag2"`
	Tag3                    			*string `json:"Tag3"`
	Tag4                    			*string `json:"Tag4"`
	DistributionProfile     			string  `json:"DistributionProfile"`
	DistributionProfileName 			string  `json:"DistributionProfileName"`
	QuestionnaireType					*string `json:"QuestionnaireType"`
	QuestionnaireTypeName				*string `json:"QuestionnaireTypeName"`
	QuestionnaireTemplate				*string `json:"QuestionnaireTemplate"`
	QuestionnaireTemplateName			*string `json:"QuestionnaireTemplateName"`
	LastChangeDate                  	string  `json:"LastChangeDate"`
	LastChangeTime                  	string  `json:"LastChangeTime"`
	CreateUser						 	int	 	`json:"CreateUser"`
	CreateUserFullName				 	*string `json:"CreateUserFullName"`
	CreateUserNickName				 	*string `json:"CreateUserNickName"`
	LastChangeUser					 	int 	`json:"LastChangeUser"`
	LastChangeUserFullName			 	*string `json:"LastChangeUserFullName"`
	LastChangeUserNickName			 	*string `json:"LastChangeUserNickName"`
	NumberOfLikes						*int	`json:"NumberOfLikes"`
	Images                  			Images  `json:"Images"`
}

type ArticlePartner struct {
	Article                   int     `json:"Article"`
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

type ArticleAddress struct {
	Article              int     `json:"Article"`
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
}

type ArticleCounter struct {
	Article					int		`json:"Article"`
	NumberOfLikes			int		`json:"NumberOfLikes"`
}

type ArticleAddressWithHeader struct {
	Article           int     `json:"Article"`
	AddressID         int     `json:"AddressID"`
	LocalSubRegion    string  `json:"LocalSubRegion"`
	LocalRegion       string  `json:"LocalRegion"`
	ArticleType       string  `json:"ArticleType"`
	ArticleOwner      int     `json:"ArticleOwner"`
	ArticleOwnerName  string  `json:"ArticleOwnerName"`
	ValidityStartDate string  `json:"ValidityStartDate"`
	ValidityStartTime string  `json:"ValidityStartTime"`
	ValidityEndDate   string  `json:"ValidityEndDate"`
	ValidityEndTime   string  `json:"ValidityEndTime"`
	Description       string  `json:"Description"`
	LongText          string  `json:"LongText"`
	Introduction      *string `json:"Introduction"`
	Site              int     `json:"Site"`
	Tag1              *string `json:"Tag1"`
	Tag2              *string `json:"Tag2"`
	Tag3              *string `json:"Tag3"`
	Tag4              *string `json:"Tag4"`
	LastChangeDate    string  `json:"LastChangeDate"`
	LastChangeTime    string  `json:"LastChangeTime"`
	NumberOfLikes	  *int	  `json:"NumberOfLikes"`
	Images            Images  `json:"Images"`
}

type ArticlePartnerWithAddress struct {
	Article                 int     `json:"Article"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	AddressID               *int    `json:"AddressID"`
	PostalCode              *string `json:"PostalCode"`
	LocalRegion             string  `json:"LocalRegion"`
	LocalRegionName         string  `json:"LocalRegionName"`
	Country                 string  `json:"Country"`
	StreetName              *string `json:"StreetName"`
	CityName                *string `json:"CityName"`
	Building                *string `json:"Building"`
	Floor                   *int    `json:"Floor"`
	Room                    *int    `json:"Room"`
}
