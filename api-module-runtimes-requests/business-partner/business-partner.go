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
	BusinessPartner               int          `json:"BusinessPartner"`
	BusinessPartnerFullName       *string      `json:"BusinessPartnerFullName"`
	BusinessPartnerName           *string      `json:"BusinessPartnerName"`
	Industry                      *string      `json:"Industry"`
	LegalEntityRegistration       *string      `json:"LegalEntityRegistration"`
	Country                       *string      `json:"Country"`
	Language                      *string      `json:"Language"`
	Currency                      *string      `json:"Currency"`
	OrganizationBPName1           *string      `json:"OrganizationBPName1"`
	OrganizationBPName2           *string      `json:"OrganizationBPName2"`
	OrganizationBPName3           *string      `json:"OrganizationBPName3"`
	OrganizationBPName4           *string      `json:"OrganizationBPName4"`
	BPTag1                        *string      `json:"BPTag1"`
	BPTag2                        *string      `json:"BPTag2"`
	BPTag3                        *string      `json:"BPTag3"`
	BPTag4                        *string      `json:"BPTag4"`
	OrganizationFoundationDate    *string      `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   *string      `json:"OrganizationLiquidationDate"`
	BusinessPartnerBirthplaceName *string      `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      *string      `json:"BusinessPartnerDeathDate"`
	GroupBusinessPartnerName1     *string      `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     *string      `json:"GroupBusinessPartnerName2"`
	AddressID                     *int         `json:"AddressID"`
	BusinessPartnerIDByExtSystem  *string      `json:"BusinessPartnerIDByExtSystem"`
	BusinessPartnerIsBlocked      *bool        `json:"BusinessPartnerIsBlocked"`
	CreationDate                  *string      `json:"CreationDate"`
	LastChangeDate                *string      `json:"LastChangeDate"`
	IsMarkedForDeletion           *bool        `json:"IsMarkedForDeletion"`
	Accounting                    []Accounting `json:"Accounting"`
	FinInst                       []FinInst    `json:"FinInst"`
	Role                          []Role       `json:"Role"`
}

type Accounting struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	ChartOfAccounts     *string `json:"ChartOfAccounts"`
	FiscalYearVariant   *string `json:"FiscalYearVariant"`
	CreationDate        *string `json:"CreationDate"`
	LastChangeDate      *string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
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

type Role struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	ValidityStartDate   string `json:"ValidityStartDate"`
}

func CreateBusinessPartnerRequestGeneralsByBusinessPartners(
	requestPram *apiInputReader.Request,
	generals []General,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		Generals: generals,
		Accepter: []string{
			"GeneralsByBusinessPartners",
		},
	}
	return req
}

func CreateBusinessPartnerRequestGenerals(
	requestPram *apiInputReader.Request,
	input apiInputReader.BusinessPartner,
) BusinessPartnerReq {
	isMarkedForDeletion := false

	req := BusinessPartnerReq{
		General: General{
			IsMarkedForDeletion: &isMarkedForDeletion,
			//IsMarkedForDeletion: requestPram.IsMarkedForDeletion,
		},
		Accepter: []string{
			"Generals",
		},
	}
	return req
}

func BusinessPartnerReads(
	requestPram *apiInputReader.Request,
	generals []General,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "reads"

	var request BusinessPartnerReq

	if accepter == "GeneralsByBusinessPartners" {
		request = CreateBusinessPartnerRequestGeneralsByBusinessPartners(
			requestPram,
			generals,
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
	)

	return responseBody
}
