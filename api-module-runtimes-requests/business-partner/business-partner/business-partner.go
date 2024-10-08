package apiModuleRuntimesRequestsBusinessPartner

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type BusinessPartnerReq struct {
	BusinessPartnerID *int      `json:"business_partner"`
	General           General   `json:"BusinessPartner"`
	Generals          []General `json:"BusinessPartners"`
	Accepter          []string  `json:"accepter"`
}

type General struct {
	BusinessPartner               int                  `json:"BusinessPartner"`
	BusinessPartnerType           *string              `json:"BusinessPartnerType"`
	BusinessPartnerFullName       *string              `json:"BusinessPartnerFullName"`
	BusinessPartnerName           *string              `json:"BusinessPartnerName"`
	Industry                      *string              `json:"Industry"`
	LegalEntityRegistration       *string              `json:"LegalEntityRegistration"`
	Country                       *string              `json:"Country"`
	Language                      *string              `json:"Language"`
	Currency                      *string              `json:"Currency"`
	Representative                *string              `json:"Representative"`
	PhoneNumber                   *string              `json:"PhoneNumber"`
	OrganizationBPName1           *string              `json:"OrganizationBPName1"`
	OrganizationBPName2           *string              `json:"OrganizationBPName2"`
	OrganizationBPName3           *string              `json:"OrganizationBPName3"`
	OrganizationBPName4           *string              `json:"OrganizationBPName4"`
	BPTag1                        *string              `json:"BPTag1"`
	BPTag2                        *string              `json:"BPTag2"`
	BPTag3                        *string              `json:"BPTag3"`
	BPTag4                        *string              `json:"BPTag4"`
	OrganizationFoundationDate    *string              `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   *string              `json:"OrganizationLiquidationDate"`
	BusinessPartnerBirthplaceName *string              `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      *string              `json:"BusinessPartnerDeathDate"`
	GroupBusinessPartnerName1     *string              `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     *string              `json:"GroupBusinessPartnerName2"`
	AddressID                     *int                 `json:"AddressID"`
	BusinessPartnerIDByExtSystem  *string              `json:"BusinessPartnerIDByExtSystem"`
	BusinessPartnerIsBlocked      *bool                `json:"BusinessPartnerIsBlocked"`
	CertificateAuthorityChain     *string              `json:"CertificateAuthorityChain"`
	UsageControlChain             *string              `json:"UsageControlChain"`
	Withdrawal           		  *bool				   `json:"Withdrawal"`
	CreationDate                  *string              `json:"CreationDate"`
	LastChangeDate                *string              `json:"LastChangeDate"`
	IsReleased           		  *bool   			   `json:"IsReleased"`
	IsMarkedForDeletion           *bool                `json:"IsMarkedForDeletion"`
	Role                          []Role               `json:"Role"`
	Person                        []Person             `json:"Person"`
	Address                       []Address            `json:"Address"`
	SNS                           []SNS                `json:"SNS"`
	GPS                           []GPS                `json:"GPS"`
	Rank                          []Rank               `json:"Rank"`
	PersonOrganization            []PersonOrganization `json:"PersonOrganization"`
	FinInst                       []FinInst            `json:"FinInst"`
	Accounting                    []Accounting         `json:"Accounting"`
}

type Role struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	BusinessPartnerRole string  `json:"BusinessPartnerRole"`
	ValidityStartDate   *string `json:"ValidityStartDate"`
	ValidityEndDate     *string `json:"ValidityEndDate"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type Person struct {
	BusinessPartner          int     `json:"BusinessPartner"`
	BusinessPartnerType      *string `json:"BusinessPartnerType"`
	FirstName                *string `json:"FirstName"`
	LastName                 *string `json:"LastName"`
	FullName                 *string `json:"FullName"`
	MiddleName               *string `json:"MiddleName"`
	NickName                 *string `json:"NickName"`
	Gender                   *string `json:"Gender"`
	Language                 *string `json:"Language"`
	CorrespondenceLanguage   *string `json:"CorrespondenceLanguage"`
	BirthDate                *string `json:"BirthDate"`
	Nationality              *string `json:"Nationality"`
	EmailAddress             *string `json:"EmailAddress"`
	MobilePhoneNumber        *string `json:"MobilePhoneNumber"`
	ProfileComment           *string `json:"ProfileComment"`
	PreferableLocalSubRegion *string `json:"PreferableLocalSubRegion"`
	PreferableLocalRegion    *string `json:"PreferableLocalRegion"`
	PreferableCountry        *string `json:"PreferableCountry"`
	ActPurpose               *string `json:"ActPurpose"`
	TermsOfUseIsConfirmed	 *bool   `json:"TermsOfUseIsConfirmed"`
	CreationDate             *string `json:"CreationDate"`
	LastChangeDate           *string `json:"LastChangeDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

type Address struct {
	BusinessPartner int      `json:"BusinessPartner"`
	AddressID       int      `json:"AddressID"`
	PostalCode      *string  `json:"PostalCode"`
	LocalSubRegion  *string  `json:"LocalSubRegion"`
	LocalRegion     *string  `json:"LocalRegion"`
	Country         *string  `json:"Country"`
	GlobalRegion    *string  `json:"GlobalRegion"`
	TimeZone        *string  `json:"TimeZone"`
	District        *string  `json:"District"`
	StreetName      *string  `json:"StreetName"`
	CityName        *string  `json:"CityName"`
	Building        *string  `json:"Building"`
	Floor           *int     `json:"Floor"`
	Room            *int     `json:"Room"`
	XCoordinate     *float32 `json:"XCoordinate"`
	YCoordinate     *float32 `json:"YCoordinate"`
	ZCoordinate     *float32 `json:"ZCoordinate"`
	Site            *int     `json:"Site"`
}

type SNS struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	BusinessPartnerType *string `json:"BusinessPartnerType"`
	XURL                *string `json:"XURL"`
	InstagramURL        *string `json:"InstagramURL"`
	TikTokURL           *string `json:"TikTokURL"`
	PointAppsURL        *string	`json:"PointAppsURL"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type GPS struct {
	BusinessPartner     int      `json:"BusinessPartner"`
	BusinessPartnerType *string  `json:"BusinessPartnerType"`
	XCoordinate         *float32 `json:"XCoordinate"`
	YCoordinate         *float32 `json:"YCoordinate"`
	ZCoordinate         *float32 `json:"ZCoordinate"`
	LocalSubRegion      *string  `json:"LocalSubRegion"`
	LocalRegion         *string  `json:"LocalRegion"`
	Country             *string  `json:"Country"`
	CreationDate        *string  `json:"CreationDate"`
	CreationTime        *string  `json:"CreationTime"`
	LastChangeDate      *string  `json:"LastChangeDate"`
	LastChangeTime      *string  `json:"LastChangeTime"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}

type Rank struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	RankType            string  `json:"RankType"`
	Rank                *int    `json:"Rank"`
	ValidityStartDate   *string `json:"ValidityStartDate"`
	ValidityEndDate     *string `json:"ValidityEndDate"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type PersonOrganization struct {
	BusinessPartner             int     `json:"BusinessPartner"`
	BusinessPartnerType         *string `json:"BusinessPartnerType"`
	OrganizationBusinessPartner *int    `json:"OrganizationBusinessPartner"`
	CreationDate                *string `json:"CreationDate"`
	LastChangeDate              *string `json:"LastChangeDate"`
	IsMarkedForDeletion         *bool   `json:"IsMarkedForDeletion"`
}

type FinInst struct {
	BusinessPartner           int     `json:"BusinessPartner"`
	FinInstIdentification     int     `json:"FinInstIdentification"`
	FinInstCountry            *string `json:"FinInstCountry"`
	FinInstCode               *string `json:"FinInstCode"`
	FinInstBranchCode         *string `json:"FinInstBranchCode"`
	FinInstFullCode           *string `json:"FinInstFullCode"`
	FinInstName               *string `json:"FinInstName"`
	FinInstBranchName         *string `json:"FinInstBranchName"`
	SWIFTCode                 *string `json:"SWIFTCode"`
	InternalFinInstCustomerID *int    `json:"InternalFinInstCustomerID"`
	InternalFinInstAccountID  *int    `json:"InternalFinInstAccountID"`
	FinInstControlKey         *string `json:"FinInstControlKey"`
	FinInstAccountName        *string `json:"FinInstAccountName"`
	FinInstAccount            *string `json:"FinInstAccount"`
	HouseBank                 *string `json:"HouseBank"`
	HouseBankAccount          *string `json:"HouseBankAccount"`
	CreationDate              *string `json:"CreationDate"`
	LastChangeDate            *string `json:"LastChangeDate"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

type Accounting struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	ChartOfAccounts     *string `json:"ChartOfAccounts"`
	FiscalYearVariant   *string `json:"FiscalYearVariant"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func CreateBusinessPartnerRequestGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.BusinessPartner,
) BusinessPartnerReq {
	isMarkedForDeletion := false

	req := BusinessPartnerReq{
		General: General{
			IsMarkedForDeletion: &isMarkedForDeletion,
		},
		Accepter: []string{
			"Generals",
		},
	}
	return req
}

func CreateBusinessPartnerRequestGeneralsByBusinessPartners(
	requestPram *apiInputReader.Request,
	input []General,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		Generals: input,
		Accepter: []string{
			"GeneralsByBusinessPartners",
		},
	}
	return req
}

func CreateBusinessPartnerRequestRole(
	requestPram *apiInputReader.Request,
	input Role,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		General: General{
			BusinessPartner: input.BusinessPartner,
			Role: []Role{
				{},
			},
		},
		Accepter: []string{
			"Role",
		},
	}
	return req
}

func CreateBusinessPartnerRequestPerson(
	requestPram *apiInputReader.Request,
	input Person,
) BusinessPartnerReq {
	// TODO いったん固定値
	isMarkedForDeletion := false

	req := BusinessPartnerReq{
		General: General{
			BusinessPartner: input.BusinessPartner,
			Person: []Person{
				{
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Person",
		},
	}
	return req
}

func CreateBusinessPartnerRequestAddresses(
	requestPram *apiInputReader.Request,
	input Address,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		General: General{
			BusinessPartner: input.BusinessPartner,
		},
		Accepter: []string{
			"Addresses",
		},
	}
	return req
}

func CreateBusinessPartnerRequestSNS(
	requestPram *apiInputReader.Request,
	input SNS,
) BusinessPartnerReq {
	// TODO いったん固定値
	isMarkedForDeletion := false

	req := BusinessPartnerReq{
		General: General{
			BusinessPartner: input.BusinessPartner,
			SNS: []SNS{
				{
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"SNS",
		},
	}
	return req
}

func CreateBusinessPartnerRequestGPS(
	requestPram *apiInputReader.Request,
	input GPS,
) BusinessPartnerReq {
	// TODO いったん固定値
	isMarkedForDeletion := false

	req := BusinessPartnerReq{
		General: General{
			BusinessPartner: input.BusinessPartner,
			GPS: []GPS{
				{
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"GPS",
		},
	}
	return req
}

func CreateBusinessPartnerRequestRank(
	requestPram *apiInputReader.Request,
	input Rank,
) BusinessPartnerReq {
	// TODO いったん固定値
	isMarkedForDeletion := false

	req := BusinessPartnerReq{
		General: General{
			BusinessPartner: input.BusinessPartner,
			Rank: []Rank{
				{
					RankType:            input.RankType,
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Rank",
		},
	}
	return req
}

func CreateBusinessPartnerRequestPersonsByBusinessPartners(
	requestPram *apiInputReader.Request,
	input []General,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		Generals: input,
		Accepter: []string{
			"PersonsByBusinessPartners",
		},
	}
	return req
}

func CreateBusinessPartnerRequestPersonOrganization(
	requestPram *apiInputReader.Request,
	input PersonOrganization,
) BusinessPartnerReq {
	// TODO いったん固定値
	isMarkedForDeletion := false

	req := BusinessPartnerReq{
		General: General{
			BusinessPartner: input.BusinessPartner,
			PersonOrganization: []PersonOrganization{
				{
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"PersonOrganization",
		},
	}
	return req
}

func BusinessPartnerReadsGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.BusinessPartner,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "reads"

	var request BusinessPartnerReq

	request = CreateBusinessPartnerRequestGenerals(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}

func BusinessPartnerReadsGeneralsByBusinessPartners(
	requestPram *apiInputReader.Request,
	input []General,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "reads"

	var request BusinessPartnerReq

	request = CreateBusinessPartnerRequestGeneralsByBusinessPartners(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}

func BusinessPartnerReadsRole(
	requestPram *apiInputReader.Request,
	input Role,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "reads"

	var request BusinessPartnerReq

	request = CreateBusinessPartnerRequestRole(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}

func BusinessPartnerReadsPerson(
	requestPram *apiInputReader.Request,
	input Person,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "reads"

	var request BusinessPartnerReq

	request = CreateBusinessPartnerRequestPerson(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}

func BusinessPartnerReadsAddresses(
	requestPram *apiInputReader.Request,
	input Address,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "reads"

	var request BusinessPartnerReq

	request = CreateBusinessPartnerRequestAddresses(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}

func BusinessPartnerReadsSNS(
	requestPram *apiInputReader.Request,
	input SNS,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "reads"

	var request BusinessPartnerReq

	request = CreateBusinessPartnerRequestSNS(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}

func BusinessPartnerReadsGPS(
	requestPram *apiInputReader.Request,
	input GPS,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "reads"

	var request BusinessPartnerReq

	request = CreateBusinessPartnerRequestGPS(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}

func BusinessPartnerReadsRank(
	requestPram *apiInputReader.Request,
	input Rank,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "reads"

	var request BusinessPartnerReq

	request = CreateBusinessPartnerRequestRank(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}

func BusinessPartnerReadsPersonsByBusinessPartners(
	requestPram *apiInputReader.Request,
	input []General,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "reads"

	var request BusinessPartnerReq

	request = CreateBusinessPartnerRequestPersonsByBusinessPartners(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}

func BusinessPartnerReadsPersonOrganization(
	requestPram *apiInputReader.Request,
	input PersonOrganization,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "reads"

	var request BusinessPartnerReq

	request = CreateBusinessPartnerRequestPersonOrganization(
		requestPram,
		input,
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
		requestPram,
	)

	return responseBody
}
