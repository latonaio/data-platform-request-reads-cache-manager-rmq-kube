package apiModuleRuntimesRequestsCertificateAuthorityChainWithUsageControlChain

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type CertificateAuthorityChainWithUsageControlChainReq struct {
	CertificateAuthorityChain CertificateAuthorityChain `json:"CertificateAuthorityChain"`
	UsageControlChain         UsageControlChain         `json:"UsageControl"`
	Accepter                  []string                  `json:"accepter"`
}

type CertificateAuthorityChain struct {
	CertificateAuthorityChain string `json:"CertificateAuthorityChain"`
	DataIssuer                int    `json:"DataIssuer"`
	DataAuthorizer            int    `json:"DataAuthorizer"`
	DataDistributor           int    `json:"DataDistributor"`
	CreationDate              string `json:"CreationDate"`
	CreationTime              string `json:"CreationTime"`
	LastChangeDate            string `json:"LastChangeDate"`
	LastChangeTime            string `json:"LastChangeTime"`
	IsMarkedForDeletion       *bool  `json:"IsMarkedForDeletion"`
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

func CreateCertificateAuthorityChainWithUsageControlChainRequestCertificateAuthorityChain(
	requestPram *apiInputReader.Request,
	certificateAuthorityChainWithUsageControlChain *apiInputReader.CertificateAuthorityChainWithUsageControlChain,
) CertificateAuthorityChainWithUsageControlChainReq {
	req := CertificateAuthorityChainWithUsageControlChainReq{
		CertificateAuthorityChain: CertificateAuthorityChain{
			CertificateAuthorityChain: certificateAuthorityChainWithUsageControlChain.CertificateAuthorityChain,
			IsMarkedForDeletion:       certificateAuthorityChainWithUsageControlChain.IsMarkedForDeletion,
		},
		Accepter: []string{
			"CertificateAuthorityChain",
		},
	}
	return req
}

func CreateCertificateAuthorityChainWithUsageControlChainRequestUsageControlChain(
	requestPram *apiInputReader.Request,
	certificateAuthorityChainWithUsageControlChain *apiInputReader.CertificateAuthorityChainWithUsageControlChain,
) CertificateAuthorityChainWithUsageControlChainReq {
	req := CertificateAuthorityChainWithUsageControlChainReq{
		UsageControlChain: UsageControlChain{
			UsageControlChain:   certificateAuthorityChainWithUsageControlChain.UsageControlChain,
			IsMarkedForDeletion: certificateAuthorityChainWithUsageControlChain.IsMarkedForDeletion,
		},
		Accepter: []string{
			"UsageControlChain",
		},
	}
	return req
}

func CertificateAuthorityChainReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.CertificateAuthorityChainWithUsageControlChainGlobal,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_CERTIFICATE_AUTHORITY_CHAIN_SRV"
	aPIType := "reads"

	var request CertificateAuthorityChainWithUsageControlChainReq

	if accepter == "CertificateAuthorityChain" {
		request = CreateCertificateAuthorityChainWithUsageControlChainRequestCertificateAuthorityChain(
			requestPram,
			&apiInputReader.CertificateAuthorityChainWithUsageControlChain{
				CertificateAuthorityChain: input.CertificateAuthorityChainWithUsageControlChain.CertificateAuthorityChain,
				IsMarkedForDeletion:       input.CertificateAuthorityChainWithUsageControlChain.IsMarkedForDeletion,
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

func UsageControlChainReads(
	requestPram *apiInputReader.Request,
	input apiInputReader.CertificateAuthorityChainWithUsageControlChainGlobal,
	controller *beego.Controller,
	accepter string,
) []byte {
	aPIServiceName := "DPFM_API_USAGE_CONTROL_CHAIN_SRV"
	aPIType := "reads"

	var request CertificateAuthorityChainWithUsageControlChainReq

	if accepter == "UsageControlChain" {
		request = CreateCertificateAuthorityChainWithUsageControlChainRequestUsageControlChain(
			requestPram,
			&apiInputReader.CertificateAuthorityChainWithUsageControlChain{
				UsageControlChain:   input.CertificateAuthorityChainWithUsageControlChain.UsageControlChain,
				IsMarkedForDeletion: input.CertificateAuthorityChainWithUsageControlChain.IsMarkedForDeletion,
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
