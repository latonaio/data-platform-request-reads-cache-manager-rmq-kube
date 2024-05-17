package site

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type SiteReq struct {
	Header   Header   `json:"Site"`
	Headers  []Header `json:"Sites"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Site                         int       `json:"Site"`
	SiteType                     *string   `json:"SiteType"`
	SiteOwner                    *int      `json:"SiteOwner"`
	SiteOwnerBusinessPartnerRole *string   `json:"SiteOwnerBusinessPartnerRole"`
	Brand						 *int	   `json:"Brand"`
	PersonResponsible            *string   `json:"PersonResponsible"`
	URL				             *string   `json:"URL"`
	ValidityStartDate            *string   `json:"ValidityStartDate"`
	ValidityStartTime            *string   `json:"ValidityStartTime"`
	ValidityEndDate              *string   `json:"ValidityEndDate"`
	ValidityEndTime              *string   `json:"ValidityEndTime"`
	DailyOperationStartTime		 *string   `json:"DailyOperationStartTime"`
	DailyOperationEndTime		 *string   `json:"DailyOperationEndTime"`
	Description                  *string   `json:"Description"`
	LongText                     *string   `json:"LongText"`
	Introduction                 *string   `json:"Introduction"`
	OperationRemarks			 *string   `json:"OperationRemarks"`
	PhoneNumber					 *string   `json:"PhoneNumber"`
	SuperiorSite				 *int	   `json:"SuperiorSite"`
	Project                      *int      `json:"Project"`
	WBSElement                   *int      `json:"WBSElement"`
	Tag1                         *string   `json:"Tag1"`
	Tag2                         *string   `json:"Tag2"`
	Tag3                         *string   `json:"Tag3"`
	Tag4                         *string   `json:"Tag4"`
	CreationDate                 *string   `json:"CreationDate"`
	CreationTime                 *string   `json:"CreationTime"`
	LastChangeDate               *string   `json:"LastChangeDate"`
	LastChangeTime               *string   `json:"LastChangeTime"`
	CreateUser					 *int	   `json:"CreateUser"`
	LastChangeUser				 *int	   `json:"LastChangeUser"`
	IsReleased					 *bool	   `json:"IsReleased"`
	IsMarkedForDeletion          *bool     `json:"IsMarkedForDeletion"`
	Partner                      []Partner `json:"Partner"`
	Address                      []Address `json:"Address"`
}

type Partner struct {
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
}

type Address struct {
	Site           int      `json:"Site"`
	AddressID      int      `json:"AddressID"`
	PostalCode     *string  `json:"PostalCode"`
	LocalSubRegion *string  `json:"LocalSubRegion"`
	LocalRegion    *string  `json:"LocalRegion"`
	Country        *string  `json:"Country"`
	GlobalRegion   *string  `json:"GlobalRegion"`
	TimeZone       *string  `json:"TimeZone"`
	District       *string  `json:"District"`
	StreetName     *string  `json:"StreetName"`
	CityName       *string  `json:"CityName"`
	Building       *string  `json:"Building"`
	Floor          *int     `json:"Floor"`
	Room           *int     `json:"Room"`
	XCoordinate    *float32 `json:"XCoordinate"`
	YCoordinate    *float32 `json:"YCoordinate"`
	ZCoordinate    *float32 `json:"ZCoordinate"`
}

func CreateSiteRequestHeader(
	requestPram *apiInputReader.Request,
	siteHeader *apiInputReader.SiteHeader,
) SiteReq {
	req := SiteReq{
		Header: Header{
			Site:                siteHeader.Site,
			IsReleased:			 siteHeader.IsReleased,
			IsMarkedForDeletion: siteHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreateSiteRequestHeaders(
	requestPram *apiInputReader.Request,
	siteHeaders *apiInputReader.SiteHeader,
) SiteReq {
	req := SiteReq{
		Header: Header{
			IsReleased:				siteHeaders.IsReleased,
			IsMarkedForDeletion:	siteHeaders.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Headers",
		},
	}
	return req
}

func CreateSiteRequestHeadersBySites(
	requestPram *apiInputReader.Request,
	siteHeaders []Header,
) SiteReq {
	req := SiteReq{
		Headers: siteHeaders,
		Accepter: []string{
			"HeadersBySites",
		},
	}
	return req
}

func CreateSiteRequestPartner(
	requestPram *apiInputReader.Request,
	sitePartner *apiInputReader.SitePartner,
) SiteReq {
	req := SiteReq{
		Header: Header{
			Site: sitePartner.Site,
			Partner: []Partner{
				{
					PartnerFunction: sitePartner.PartnerFunction,
					BusinessPartner: sitePartner.BusinessPartner,
				},
			},
		},
		Accepter: []string{
			"Partner",
		},
	}
	return req
}

func CreateSiteRequestPartners(
	requestPram *apiInputReader.Request,
	sitePartners *apiInputReader.SitePartner,
) SiteReq {
	req := SiteReq{
		Header: Header{
			Site: sitePartners.Site,
			Partner: []Partner{
				{
					//					IsMarkedForDeletion:           sitePartners.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Partners",
		},
	}
	return req
}

func CreateSiteRequestAddress(
	requestPram *apiInputReader.Request,
	siteAddress *apiInputReader.SiteAddress,
) SiteReq {
	req := SiteReq{
		Header: Header{
			Site: siteAddress.Site,
			Address: []Address{
				{
					AddressID: siteAddress.AddressID,
				},
			},
		},
		Accepter: []string{
			"Address",
		},
	}
	return req
}

func CreateSiteRequestAddresses(
	requestPram *apiInputReader.Request,
	siteAddresses *apiInputReader.SiteAddress,
) SiteReq {
	req := SiteReq{
		Header: Header{
			Site: siteAddresses.Site,
			Address: []Address{
				{
					//					IsMarkedForDeletion:           siteAddresses.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Addresses",
		},
	}
	return req
}

func CreateSiteRequestAddressesByLocalSubRegion(
	requestPram *apiInputReader.Request,
	siteAddresses *apiInputReader.SiteAddress,
) SiteReq {
	req := SiteReq{
		Header: Header{
			Address: []Address{
				{
					LocalSubRegion: siteAddresses.LocalSubRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalSubRegion",
		},
	}
	return req
}

func CreateSiteRequestAddressesByLocalSubRegions(
	requestPram *apiInputReader.Request,
	siteAddresses *apiInputReader.SiteAddress,
) SiteReq {
	req := SiteReq{
		Header: Header{
			Address: []Address{
				{
					LocalSubRegion: siteAddresses.LocalSubRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalSubRegions",
		},
	}
	return req
}

func CreateSiteRequestAddressesByLocalRegion(
	requestPram *apiInputReader.Request,
	siteAddresses *apiInputReader.SiteAddress,
) SiteReq {
	req := SiteReq{
		Header: Header{
			Address: []Address{
				{
					LocalRegion: siteAddresses.LocalRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalRegion",
		},
	}
	return req
}

func CreateSiteRequestAddressesByLocalRegions(
	requestPram *apiInputReader.Request,
	siteAddresses *apiInputReader.SiteAddress,
) SiteReq {
	req := SiteReq{
		Header: Header{
			Site: siteAddresses.Site,
			Address: []Address{
				{
					LocalRegion: siteAddresses.LocalRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalRegions",
		},
	}
	return req
}

func SiteReadsHeadersBySites(
	requestPram *apiInputReader.Request,
	input []Header,
	controller *beego.Controller,
) []byte {

	aPIServiceName := "DPFM_API_SITE_SRV"
	aPIType := "reads"

	request := CreateSiteRequestHeadersBySites(
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
	)

	return responseBody
}

func SiteReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Site,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_SITE_SRV"
	aPIType := "reads"

	var request SiteReq

	if accepter == "Header" {
		request = CreateSiteRequestHeader(
			requestPram,
			&apiInputReader.SiteHeader{
				Site:                input.SiteHeader.Site,
				IsReleased:          input.SiteHeader.IsReleased,
				IsMarkedForDeletion: input.SiteHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Headers" {
		request = CreateSiteRequestHeaders(
			requestPram,
			&apiInputReader.SiteHeader{
				IsReleased:          input.SiteHeader.IsReleased,
				IsMarkedForDeletion: input.SiteHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Partner" {
		request = CreateSiteRequestPartner(
			requestPram,
			&apiInputReader.SitePartner{
				Site:            input.SitePartner.Site,
				PartnerFunction: input.SitePartner.PartnerFunction,
				BusinessPartner: input.SitePartner.BusinessPartner,
			},
		)
	}

	if accepter == "Partners" {
		request = CreateSiteRequestPartners(
			requestPram,
			&apiInputReader.SitePartner{
				Site: input.SitePartner.Site,
			},
		)
	}

	if accepter == "Address" {
		request = CreateSiteRequestAddress(
			requestPram,
			&apiInputReader.SiteAddress{
				Site:      input.SiteAddress.Site,
				AddressID: input.SiteAddress.AddressID,
			},
		)
	}

	if accepter == "Addresses" {
		request = CreateSiteRequestAddresses(
			requestPram,
			&apiInputReader.SiteAddress{
				Site: input.SiteAddress.Site,
			},
		)
	}

	if accepter == "AddressesByLocalSubRegion" {
		request = CreateSiteRequestAddressesByLocalSubRegion(
			requestPram,
			&apiInputReader.SiteAddress{
				LocalSubRegion: input.SiteAddress.LocalSubRegion,
			},
		)
	}

	if accepter == "AddressesByLocalSubRegions" {
		request = CreateSiteRequestAddressesByLocalSubRegions(
			requestPram,
			&apiInputReader.SiteAddress{
				LocalSubRegion: input.SiteAddress.LocalSubRegion,
			},
		)
	}

	if accepter == "AddressesByLocalRegion" {
		request = CreateSiteRequestAddressesByLocalRegion(
			requestPram,
			&apiInputReader.SiteAddress{
				LocalRegion: input.SiteAddress.LocalRegion,
			},
		)
	}

	if accepter == "AddressesByLocalRegions" {
		request = CreateSiteRequestAddressesByLocalRegions(
			requestPram,
			&apiInputReader.SiteAddress{
				LocalRegion: input.SiteAddress.LocalRegion,
			},
		)
	}

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
	)

	return responseBody
}
