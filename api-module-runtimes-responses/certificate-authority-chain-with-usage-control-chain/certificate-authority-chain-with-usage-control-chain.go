package apiModuleRuntimesResponsesCertificateAuthorityChainWithUsageControlChain

type CertificateAuthorityChainWithUsageControlChainRes struct {
	Message CertificateAuthorityChainWithUsageControlChainGlobal `json:"message,omitempty"`
}

type CertificateAuthorityChainWithUsageControlChainGlobal struct {
	CertificateAuthorityChain	*[]CertificateAuthorityChain	`json:"CertificateAuthority,omitempty"`
	UsageControlChain			*[]UsageControlChain			`json:"UsageControl,omitempty"`
}

type CertificateAuthorityChain struct {
	CertificateAuthorityChain      string   `json:"CertificateAuthorityChain"`
	DataIssuer               	   int		`json:"DataIssuer"`
	DataAuthorizer                 int		`json:"DataAuthorizer"`
	DataDistributor                int		`json:"DataDistributor"`
	CreationDate                   string   `json:"CreationDate"`
	CreationTime                   string   `json:"CreationTime"`
	LastChangeDate                 string   `json:"LastChangeDate"`
	LastChangeTime                 string   `json:"LastChangeTime"`
	IsMarkedForDeletion            *bool    `json:"IsMarkedForDeletion"`
}

type UsageControlChain struct {
	UsageControlChain              string   `json:"UsageControlChain"`
	UsageControlLess               *bool    `json:"UsageControlLess"`
	Perpetual                      *bool    `json:"Perpetual"`
	Rental                         *bool    `json:"Rental"`
	Duration                       *float64 `json:"Duration"`
	DurationUnit                   *string  `json:"DurationUnit"`
	ValidityStartDate              *string  `json:"ValidityStartDate"`
	ValidityStartTime              *string  `json:"ValidityStartTime"`
	ValidityEndDate                *string  `json:"ValidityEndDate"`
	ValidityEndTime                *string  `json:"ValidityEndTime"`
	DeleteAfterValidityEnd         *bool    `json:"DeleteAfterValidityEnd"`
	ServiceLabelRestriction        *string  `json:"ServiceLabelRestriction"`
	ApplicationRestriction         *string  `json:"ApplicationRestriction"`
	PurposeRestriction             *string  `json:"PurposeRestriction"`
	BusinessPartnerRoleRestriction *string  `json:"BusinessPartnerRoleRestriction"`
	DataStateRestriction           *string  `json:"DataStateRestriction"`
	NumberOfUsageRestriction       *int     `json:"NumberOfUsageRestriction"`
	NumberOfActualUsage            *int     `json:"NumberOfActualUsage"`
	IPAddressRestriction           *string  `json:"IPAddressRestriction"`
	MACAddressRestriction          *string  `json:"MACAddressRestriction"`
	ModifyIsAllowed                *bool    `json:"ModifyIsAllowed"`
	LocalLoggingIsAllowed          *bool    `json:"LocalLoggingIsAllowed"`
	RemoteNotificationIsAllowed    *string  `json:"RemoteNotificationIsAllowed"`
	DistributeOnlyIfEncrypted      *bool    `json:"DistributeOnlyIfEncrypted"`
	AttachPolicyWhenDistribute     *bool    `json:"AttachPolicyWhenDistribute"`
	PostalCode                     *string  `json:"PostalCode"`
	LocalSubRegion                 *string  `json:"LocalSubRegion"`
	LocalRegion                    *string  `json:"LocalRegion"`
	Country                        *string  `json:"Country"`
	GlobalRegion                   *string  `json:"GlobalRegion"`
	TimeZone                       *string  `json:"TimeZone"`
	CreationDate                   string   `json:"CreationDate"`
	CreationTime                   string   `json:"CreationTime"`
	LastChangeDate                 string   `json:"LastChangeDate"`
	LastChangeTime                 string   `json:"LastChangeTime"`
	IsMarkedForDeletion            *bool    `json:"IsMarkedForDeletion"`
}
