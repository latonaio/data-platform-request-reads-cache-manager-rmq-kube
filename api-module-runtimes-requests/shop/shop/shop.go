package shop

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type ShopReq struct {
	Header   Header   `json:"Shop"`
	Headers  []Header `json:"Shops"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Shop                         int       `json:"Shop"`
	ShopType                     *string   `json:"ShopType"`
	ShopOwner                    *int      `json:"ShopOwner"`
	ShopOwnerBusinessPartnerRole *string   `json:"ShopOwnerBusinessPartnerRole"`
	Brand						 *int	   `json:"Brand"`
	PersonResponsible            *string   `json:"PersonResponsible"`
	URL							 *string   `json:"URL"`
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
	Site						 *int	   `json:"Site"`
	Project                      *int      `json:"Project"`
	WBSElement                   *int      `json:"WBSElement"`
	Tag1                         *string   `json:"Tag1"`
	Tag2                         *string   `json:"Tag2"`
	Tag3                         *string   `json:"Tag3"`
	Tag4                         *string   `json:"Tag4"`
	PointConsumptionType         *string   `json:"PointConsumptionType"`
	CreationDate                 *string   `json:"CreationDate"`
	CreationTime                 *string   `json:"CreationTime"`
	LastChangeDate               *string   `json:"LastChangeDate"`
	LastChangeTime               *string   `json:"LastChangeTime"`
	CreateUser					 *int      `json:"CreateUser"`
	LastChangeUser				 *int      `json:"LastChangeUser"`
	IsReleased					 *bool	   `json:"IsReleased"`
	IsMarkedForDeletion          *bool     `json:"IsMarkedForDeletion"`
	Partner                      []Partner `json:"Partner"`
	Address                      []Address `json:"Address"`
}

type Partner struct {
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
}

type Address struct {
	Shop           int      `json:"Shop"`
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
	Site		   *int		`json:"Site"`
}

func CreateShopRequestHeader(
	requestPram *apiInputReader.Request,
	shopHeader *apiInputReader.ShopHeader,
) ShopReq {
	req := ShopReq{
		Header: Header{
			Shop:                shopHeader.Shop,
			IsReleased:			 shopHeader.IsReleased,
			IsMarkedForDeletion: shopHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreateShopRequestHeaders(
	requestPram *apiInputReader.Request,
	shopHeaders *apiInputReader.ShopHeader,
) ShopReq {
	req := ShopReq{
		Header: Header{
			IsReleased:				shopHeaders.IsReleased,
			IsMarkedForDeletion:	shopHeaders.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Headers",
		},
	}
	return req
}

func CreateShopRequestHeadersByShops(
	requestPram *apiInputReader.Request,
	shopHeaders []Header,
) ShopReq {
	req := ShopReq{
		Headers: shopHeaders,
		Accepter: []string{
			"HeadersByShops",
		},
	}
	return req
}

func CreateShopRequestPartner(
	requestPram *apiInputReader.Request,
	shopPartner *apiInputReader.ShopPartner,
) ShopReq {
	req := ShopReq{
		Header: Header{
			Shop: shopPartner.Shop,
			Partner: []Partner{
				{
					PartnerFunction: shopPartner.PartnerFunction,
					BusinessPartner: shopPartner.BusinessPartner,
				},
			},
		},
		Accepter: []string{
			"Partner",
		},
	}
	return req
}

func CreateShopRequestPartners(
	requestPram *apiInputReader.Request,
	shopPartners *apiInputReader.ShopPartner,
) ShopReq {
	req := ShopReq{
		Header: Header{
			Shop: shopPartners.Shop,
			Partner: []Partner{
				{
					//					IsMarkedForDeletion:           shopPartners.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Partners",
		},
	}
	return req
}

func CreateShopRequestAddress(
	requestPram *apiInputReader.Request,
	shopAddress *apiInputReader.ShopAddress,
) ShopReq {
	req := ShopReq{
		Header: Header{
			Shop: shopAddress.Shop,
			Address: []Address{
				{
					AddressID: shopAddress.AddressID,
				},
			},
		},
		Accepter: []string{
			"Address",
		},
	}
	return req
}

func CreateShopRequestAddresses(
	requestPram *apiInputReader.Request,
	shopAddresses *apiInputReader.ShopAddress,
) ShopReq {
	req := ShopReq{
		Header: Header{
			Shop: shopAddresses.Shop,
			Address: []Address{
				{
					//					IsMarkedForDeletion:           shopAddresses.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Addresses",
		},
	}
	return req
}

func CreateShopRequestAddressesByLocalSubRegion(
	requestPram *apiInputReader.Request,
	shopAddresses *apiInputReader.ShopAddress,
) ShopReq {
	req := ShopReq{
		Header: Header{
			Address: []Address{
				{
					LocalSubRegion: shopAddresses.LocalSubRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalSubRegion",
		},
	}
	return req
}

func CreateShopRequestAddressesByLocalSubRegions(
	requestPram *apiInputReader.Request,
	shopAddresses *apiInputReader.ShopAddress,
) ShopReq {
	req := ShopReq{
		Header: Header{
			Address: []Address{
				{
					LocalSubRegion: shopAddresses.LocalSubRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalSubRegions",
		},
	}
	return req
}

func CreateShopRequestAddressesByLocalRegion(
	requestPram *apiInputReader.Request,
	shopAddresses *apiInputReader.ShopAddress,
) ShopReq {
	req := ShopReq{
		Header: Header{
			Address: []Address{
				{
					LocalRegion: shopAddresses.LocalRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalRegion",
		},
	}
	return req
}

func CreateShopRequestAddressesByLocalRegions(
	requestPram *apiInputReader.Request,
	shopAddresses *apiInputReader.ShopAddress,
) ShopReq {
	req := ShopReq{
		Header: Header{
			Shop: shopAddresses.Shop,
			Address: []Address{
				{
					LocalRegion: shopAddresses.LocalRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalRegions",
		},
	}
	return req
}

func ShopReadsHeader(
	requestPram *apiInputReader.Request,
	input []Header,
	controller *beego.Controller,
) []byte {

	aPIServiceName := "DPFM_API_SHOP_SRV"
	aPIType := "reads"

	request := CreateShopRequestHeader(
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

func ShopReadsHeadersByShops(
	requestPram *apiInputReader.Request,
	input []Header,
	controller *beego.Controller,
) []byte {

	aPIServiceName := "DPFM_API_SHOP_SRV"
	aPIType := "reads"

	request := CreateShopRequestHeadersByShops(
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

func ShopReadsAddresses(
	requestPram *apiInputReader.Request,
	input []Header,
	controller *beego.Controller,
) []byte {

	aPIServiceName := "DPFM_API_SHOP_SRV"
	aPIType := "reads"

	request := CreateShopRequestAddresses(
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

func ShopReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Shop,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_SHOP_SRV"
	aPIType := "reads"

	var request ShopReq

	if accepter == "Header" {
		request = CreateShopRequestHeader(
			requestPram,
			&apiInputReader.ShopHeader{
				Shop:                input.ShopHeader.Shop,
				IsReleased:          input.ShopHeader.IsReleased,
				IsMarkedForDeletion: input.ShopHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Headers" {
		request = CreateShopRequestHeaders(
			requestPram,
			&apiInputReader.ShopHeader{
				IsReleased:          input.ShopHeader.IsReleased,
				IsMarkedForDeletion: input.ShopHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Partner" {
		request = CreateShopRequestPartner(
			requestPram,
			&apiInputReader.ShopPartner{
				Shop:            input.ShopPartner.Shop,
				PartnerFunction: input.ShopPartner.PartnerFunction,
				BusinessPartner: input.ShopPartner.BusinessPartner,
			},
		)
	}

	if accepter == "Partners" {
		request = CreateShopRequestPartners(
			requestPram,
			&apiInputReader.ShopPartner{
				Shop: input.ShopPartner.Shop,
			},
		)
	}

	if accepter == "Address" {
		request = CreateShopRequestAddress(
			requestPram,
			&apiInputReader.ShopAddress{
				Shop:      input.ShopAddress.Shop,
				AddressID: input.ShopAddress.AddressID,
			},
		)
	}

	if accepter == "Addresses" {
		request = CreateShopRequestAddresses(
			requestPram,
			&apiInputReader.ShopAddress{
				Shop: input.ShopAddress.Shop,
			},
		)
	}

	if accepter == "AddressesByLocalSubRegion" {
		request = CreateShopRequestAddressesByLocalSubRegion(
			requestPram,
			&apiInputReader.ShopAddress{
				LocalSubRegion: input.ShopAddress.LocalSubRegion,
			},
		)
	}

	if accepter == "AddressesByLocalSubRegions" {
		request = CreateShopRequestAddressesByLocalSubRegions(
			requestPram,
			&apiInputReader.ShopAddress{
				LocalSubRegion: input.ShopAddress.LocalSubRegion,
			},
		)
	}

	if accepter == "AddressesByLocalRegion" {
		request = CreateShopRequestAddressesByLocalRegion(
			requestPram,
			&apiInputReader.ShopAddress{
				LocalRegion: input.ShopAddress.LocalRegion,
			},
		)
	}

	if accepter == "AddressesByLocalRegions" {
		request = CreateShopRequestAddressesByLocalRegions(
			requestPram,
			&apiInputReader.ShopAddress{
				LocalRegion: input.ShopAddress.LocalRegion,
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
