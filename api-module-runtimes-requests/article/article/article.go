package article

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type ArticleReq struct {
	Header   Header   `json:"Article"`
	Headers  []Header `json:"Articles"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Article                         int                     `json:"Article"`
	ArticleType                     *string                 `json:"ArticleType"`
	ArticleOwner                    *int                    `json:"ArticleOwner"`
	ArticleOwnerBusinessPartnerRole *string                 `json:"ArticleOwnerBusinessPartnerRole"`
	PersonResponsible               *string                 `json:"PersonResponsible"`
	ValidityStartDate               *string                 `json:"ValidityStartDate"`
	ValidityStartTime               *string                 `json:"ValidityStartTime"`
	ValidityEndDate                 *string                 `json:"ValidityEndDate"`
	ValidityEndTime                 *string                 `json:"ValidityEndTime"`
	Description                     *string                 `json:"Description"`
	LongText                        *string                 `json:"LongText"`
	Introduction                    *string                 `json:"Introduction"`
	Site                            *int                    `json:"Site"`
	Shop                            *int                    `json:"Shop"`
	Project                         *int                    `json:"Project"`
	WBSElement                      *int                    `json:"WBSElement"`
	Tag1                            *string                 `json:"Tag1"`
	Tag2                            *string                 `json:"Tag2"`
	Tag3                            *string                 `json:"Tag3"`
	Tag4                            *string                 `json:"Tag4"`
	DistributionProfile             *string                 `json:"DistributionProfile"`
	QuestionnaireType               *string                 `json:"QuestionnaireType"`
	QuestionnaireTemplate           *string                 `json:"QuestionnaireTemplate"`
	CreationDate                    *string                 `json:"CreationDate"`
	CreationTime                    *string                 `json:"CreationTime"`
	LastChangeDate                  *string                 `json:"LastChangeDate"`
	LastChangeTime                  *string                 `json:"LastChangeTime"`
	CreateUser                      *int                    `json:"CreateUser"`
	LastChangeUser                  *int                    `json:"LastChangeUser"`
	IsReleased                      *bool                   `json:"IsReleased"`
	IsMarkedForDeletion             *bool                   `json:"IsMarkedForDeletion"`
	Partner                         []Partner               `json:"Partner"`
	Address                         []Address               `json:"Address"`
	Counter							[]Counter               `json:"Counter"`
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
	Site           *int     `json:"Site"`
}

type Counter struct {
	Article                int     `json:"Article"`
	NumberOfLikes          *int    `json:"NumberOfLikes"`
	CreationDate           *string `json:"CreationDate"`
	CreationTime           *string `json:"CreationTime"`
	LastChangeDate         *string `json:"LastChangeDate"`
	LastChangeTime         *string `json:"LastChangeTime"`
}

func CreateArticleRequestHeader(
	requestPram *apiInputReader.Request,
	articleHeader *apiInputReader.ArticleHeader,
) ArticleReq {
	req := ArticleReq{
		Header: Header{
			Article:             articleHeader.Article,
			IsReleased:          articleHeader.IsReleased,
			IsMarkedForDeletion: articleHeader.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreateArticleRequestHeaders(
	requestPram *apiInputReader.Request,
	articleHeaders *apiInputReader.ArticleHeader,
) ArticleReq {
	req := ArticleReq{
		Header: Header{
			IsReleased:          articleHeaders.IsReleased,
			IsMarkedForDeletion: articleHeaders.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Headers",
		},
	}
	return req
}

func CreateArticleRequestHeadersByArticles(
	requestPram *apiInputReader.Request,
	articleHeaders []Header,
) ArticleReq {
	req := ArticleReq{
		Headers: articleHeaders,
		Accepter: []string{
			"HeadersByArticles",
		},
	}
	return req
}

func CreateArticleRequestHeadersByArticleOwner(
	requestPram *apiInputReader.Request,
	articleHeaders *apiInputReader.ArticleHeader,
) ArticleReq {
	req := ArticleReq{
		Header: Header{
			ArticleOwner:	     articleHeaders.ArticleOwner,
			IsReleased:          articleHeaders.IsReleased,
			IsMarkedForDeletion: articleHeaders.IsMarkedForDeletion,
		},
		Accepter: []string{
			"HeadersByArticleOwner",
		},
	}
	return req
}

func CreateArticleRequestPartner(
	requestPram *apiInputReader.Request,
	articlePartner *apiInputReader.ArticlePartner,
) ArticleReq {
	req := ArticleReq{
		Header: Header{
			Article: articlePartner.Article,
			Partner: []Partner{
				{
					PartnerFunction: articlePartner.PartnerFunction,
					BusinessPartner: articlePartner.BusinessPartner,
				},
			},
		},
		Accepter: []string{
			"Partner",
		},
	}
	return req
}

func CreateArticleRequestPartners(
	requestPram *apiInputReader.Request,
	articlePartners *apiInputReader.ArticlePartner,
) ArticleReq {
	req := ArticleReq{
		Header: Header{
			Article: articlePartners.Article,
			Partner: []Partner{
				{
					//					IsMarkedForDeletion:           articlePartners.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Partners",
		},
	}
	return req
}

func CreateArticleRequestAddress(
	requestPram *apiInputReader.Request,
	articleAddress *apiInputReader.ArticleAddress,
) ArticleReq {
	req := ArticleReq{
		Header: Header{
			Article: articleAddress.Article,
			Address: []Address{
				{
					AddressID: articleAddress.AddressID,
				},
			},
		},
		Accepter: []string{
			"Address",
		},
	}
	return req
}

func CreateArticleRequestAddresses(
	requestPram *apiInputReader.Request,
	articleAddresses *apiInputReader.ArticleAddress,
) ArticleReq {
	req := ArticleReq{
		Header: Header{
			Article: articleAddresses.Article,
			Address: []Address{
				{
					//					IsMarkedForDeletion:           articleAddresses.IsMarkedForDeletion,
				},
			},
		},
		Accepter: []string{
			"Addresses",
		},
	}
	return req
}

func CreateArticleRequestAddressesByLocalSubRegion(
	requestPram *apiInputReader.Request,
	articleAddresses *apiInputReader.ArticleAddress,
) ArticleReq {
	req := ArticleReq{
		Header: Header{
			Address: []Address{
				{
					LocalSubRegion: articleAddresses.LocalSubRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalSubRegion",
		},
	}
	return req
}

func CreateArticleRequestAddressesByLocalSubRegions(
	requestPram *apiInputReader.Request,
	articleAddresses *apiInputReader.ArticleAddress,
) ArticleReq {
	req := ArticleReq{
		Header: Header{
			Address: []Address{
				{
					LocalSubRegion: articleAddresses.LocalSubRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalSubRegions",
		},
	}
	return req
}

func CreateArticleRequestAddressesByLocalRegion(
	requestPram *apiInputReader.Request,
	articleAddresses *apiInputReader.ArticleAddress,
) ArticleReq {
	req := ArticleReq{
		Header: Header{
			Address: []Address{
				{
					LocalRegion: articleAddresses.LocalRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalRegion",
		},
	}
	return req
}

func CreateArticleRequestAddressesByLocalRegions(
	requestPram *apiInputReader.Request,
	articleAddresses *apiInputReader.ArticleAddress,
) ArticleReq {
	req := ArticleReq{
		Header: Header{
			Article: articleAddresses.Article,
			Address: []Address{
				{
					LocalRegion: articleAddresses.LocalRegion,
				},
			},
		},
		Accepter: []string{
			"AddressesByLocalRegions",
		},
	}
	return req
}

func CreateArticleRequestCounter(
	requestPram *apiInputReader.Request,
	articleCounter *apiInputReader.ArticleCounter,
) ArticleReq {
	req := ArticleReq{
		Header: Header{
			Article: articleCounter.Article,
			Counter: []Counter{
				{
				},
			},
		},
		Accepter: []string{
			"Counter",
		},
	}
	return req
}

func CreateArticleRequestCountersByArticles(
	requestPram *apiInputReader.Request,
	articleHeaders []Header,
) ArticleReq {
	req := ArticleReq{
		Headers: articleHeaders,
		Accepter: []string{
			"CountersByArticles",
		},
	}
	return req
}

func ArticleReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.Article,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_ARTICLE_SRV"
	aPIType := "reads"

	var request ArticleReq

	if accepter == "Header" {
		request = CreateArticleRequestHeader(
			requestPram,
			&apiInputReader.ArticleHeader{
				Article:               input.ArticleHeader.Article,
				IsReleased:          input.ArticleHeader.IsReleased,
				IsMarkedForDeletion: input.ArticleHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Headers" {
		request = CreateArticleRequestHeaders(
			requestPram,
			&apiInputReader.ArticleHeader{
				IsReleased:          input.ArticleHeader.IsReleased,
				IsMarkedForDeletion: input.ArticleHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "HeadersByArticleOwner" {
		request = CreateArticleRequestHeadersByArticleOwner(
			requestPram,
			&apiInputReader.ArticleHeader{
				ArticleOwner:	     input.ArticleHeader.ArticleOwner,
				IsReleased:          input.ArticleHeader.IsReleased,
				IsMarkedForDeletion: input.ArticleHeader.IsMarkedForDeletion,
			},
		)
	}

	if accepter == "Partner" {
		request = CreateArticleRequestPartner(
			requestPram,
			&apiInputReader.ArticlePartner{
				Article:           input.ArticlePartner.Article,
				PartnerFunction: input.ArticlePartner.PartnerFunction,
				BusinessPartner: input.ArticlePartner.BusinessPartner,
			},
		)
	}

	if accepter == "Partners" {
		request = CreateArticleRequestPartners(
			requestPram,
			&apiInputReader.ArticlePartner{
				Article: input.ArticlePartner.Article,
			},
		)
	}

	if accepter == "Address" {
		request = CreateArticleRequestAddress(
			requestPram,
			&apiInputReader.ArticleAddress{
				Article:     input.ArticleAddress.Article,
				AddressID: input.ArticleAddress.AddressID,
			},
		)
	}

	if accepter == "Addresses" {
		request = CreateArticleRequestAddresses(
			requestPram,
			&apiInputReader.ArticleAddress{
				Article: input.ArticleAddress.Article,
			},
		)
	}

	if accepter == "AddressesByLocalSubRegion" {
		request = CreateArticleRequestAddressesByLocalSubRegion(
			requestPram,
			&apiInputReader.ArticleAddress{
				LocalSubRegion: input.ArticleAddress.LocalSubRegion,
			},
		)
	}

	if accepter == "AddressesByLocalSubRegions" {
		request = CreateArticleRequestAddressesByLocalSubRegions(
			requestPram,
			&apiInputReader.ArticleAddress{
				LocalSubRegion: input.ArticleAddress.LocalSubRegion,
			},
		)
	}

	if accepter == "AddressesByLocalRegion" {
		request = CreateArticleRequestAddressesByLocalRegion(
			requestPram,
			&apiInputReader.ArticleAddress{
				LocalRegion: input.ArticleAddress.LocalRegion,
			},
		)
	}

	if accepter == "AddressesByLocalRegions" {
		request = CreateArticleRequestAddressesByLocalRegions(
			requestPram,
			&apiInputReader.ArticleAddress{
				LocalRegion: input.ArticleAddress.LocalRegion,
			},
		)
	}

	if accepter == "Counter" {
		request = CreateArticleRequestCounter(
			requestPram,
			&apiInputReader.ArticleCounter{
				Article:	input.ArticleCounter.Article,
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
		requestPram,
	)

	return responseBody
}

func ArticleReadsHeadersByArticles(
	requestPram *apiInputReader.Request,
	input []Header,
	controller *beego.Controller,
) []byte {

	aPIServiceName := "DPFM_API_ARTICLE_SRV"
	aPIType := "reads"

	request := CreateArticleRequestHeadersByArticles(
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

func ArticleReadsCountersByArticles(
	requestPram *apiInputReader.Request,
	input []Header,
	controller *beego.Controller,
) []byte {

	//if accepter == "CountersByArticles" {
	//	request = CreateArticleRequestCountersByArticles(
	//		requestPram,
	//		&apiInputReader.ArticleCounter{
	//			Article:    input.ArticleCounter.Article,
	//		},
	//	)
	//}

	aPIServiceName := "DPFM_API_ARTICLE_SRV"
	aPIType := "reads"

	request := CreateArticleRequestCountersByArticles(
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
