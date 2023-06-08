package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type BusinessPartnerRes struct {
	ConnectionKey       string     `json:"connection_key,omitempty"`
	Result              bool       `json:"result,omitempty"`
	RedisKey            string     `json:"redis_key,omitempty"`
	Filepath            string     `json:"filepath,omitempty"`
	APIStatusCode       int        `json:"api_status_code,omitempty"`
	RuntimeSessionID    string     `json:"runtime_session_id,omitempty"`
	BusinessPartnerID   *int       `json:"business_partner,omitempty"`
	ServiceLabel        string     `json:"service_label,omitempty"`
	APIType             string     `json:"api_type,omitempty"`
	Message             *BPMessage `json:"message,omitempty"`
	APISchema           string     `json:"api_schema,omitempty"`
	Accepter            []string   `json:"accepter,omitempty"`
	Deleted             bool       `json:"deleted,omitempty"`
	SQLUpdateResult     *bool      `json:"sql_update_result,omitempty"`
	SQLUpdateError      string     `json:"sql_update_error,omitempty"`
	SubfuncResult       *bool      `json:"subfunc_result,omitempty"`
	SubfuncError        string     `json:"subfunc_error,omitempty"`
	ExconfResult        *bool      `json:"exconf_result,omitempty"`
	ExconfError         string     `json:"exconf_error,omitempty"`
	APIProcessingResult *bool      `json:"api_processing_result,omitempty"`
	APIProcessingError  string     `json:"api_processing_error,omitempty"`
}

type BPMessage struct {
	General    *BPGeneral      `json:"General"`
	Generals   *[]BPGeneral    `json:"Generals"`
	Role       *[]Role         `json:"Role"`
	FinInst    *[]FinInst      `json:"FinInst"`
	Accounting *[]BPAccounting `json:"Accounting"`
}

type BusinessPartnerGeneral struct {
	ConnectionKey string `json:"connection_key,omitempty"`
	Result        bool   `json:"result,omitempty"`
	RedisKey      string `json:"redis_key,omitempty"`
	Filepath      string `json:"filepath,omitempty"`
	Product       string `json:"Product,omitempty"`
	APISchema     string `json:"api_schema,omitempty"`
	MaterialCode  string `json:"material_code,omitempty"`
	Deleted       string `json:"deleted,omitempty"`
}

type BPGeneral struct {
	BusinessPartner               int     `json:"BusinessPartner"`
	BusinessPartnerFullName       *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName           *string `json:"BusinessPartnerName"`
	CreationDate                  *string `json:"CreationDate"`
	CreationTime                  *string `json:"CreationTime"`
	Industry                      *string `json:"Industry"`
	LegalEntityRegistration       *string `json:"LegalEntityRegistration"`
	Country                       *string `json:"Country"`
	Language                      *string `json:"Language"`
	Currency                      *string `json:"Currency"`
	LastChangeDate                *string `json:"LastChangeDate"`
	LastChangeTime                *string `json:"LastChangeTime"`
	OrganizationBPName1           *string `json:"OrganizationBPName1"`
	OrganizationBPName2           *string `json:"OrganizationBPName2"`
	OrganizationBPName3           *string `json:"OrganizationBPName3"`
	OrganizationBPName4           *string `json:"OrganizationBPName4"`
	BPTag1                        *string `json:"BPTag1"`
	BPTag2                        *string `json:"BPTag2"`
	BPTag3                        *string `json:"BPTag3"`
	BPTag4                        *string `json:"BPTag4"`
	OrganizationFoundationDate    *string `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   *string `json:"OrganizationLiquidationDate"`
	BusinessPartnerBirthplaceName *string `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      *string `json:"BusinessPartnerDeathDate"`
	BusinessPartnerIsBlocked      *bool   `json:"BusinessPartnerIsBlocked"`
	GroupBusinessPartnerName1     *string `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     *string `json:"GroupBusinessPartnerName2"`
	AddressID                     *int    `json:"AddressID"`
	BusinessPartnerIDByExtSystem  *string `json:"BusinessPartnerIDByExtSystem"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}

type GeneralPDF struct {
	BusinessPartner          int     `json:"BusinessPartner,omitempty"`
	DocType                  string  `json:"DocType,omitempty"`
	DocVersionID             int     `json:"DocVersionID,omitempty"`
	DocID                    string  `json:"DocID,omitempty"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner,omitempty"`
	FileName                 *string `json:"FileName,omitempty"`
}

type Role struct {
	BusinessPartner     int    `json:"BusinessPartner,omitempty"`
	BusinessPartnerRole string `json:"BusinessPartnerRole,omitempty"`
	ValidityEndDate     string `json:"ValidityEndDate,omitempty"`
	ValidityStartDate   string `json:"ValidityStartDate,omitempty"`
}

type FinInst struct {
	BusinessPartner           int     `json:"BusinessPartner,omitempty"`
	FinInstIdentification     int     `json:"FinInstIdentification,omitempty"`
	ValidityEndDate           string  `json:"ValidityEndDate,omitempty"`
	ValidityStartDate         string  `json:"ValidityStartDate,omitempty"`
	FinInstCountry            *string `json:"FinInstCountry,omitempty"`
	FinInstCode               *string `json:"FinInstCode,omitempty"`
	FinInstBranchCode         *string `json:"FinInstBranchCode,omitempty"`
	FinInstFullCode           *string `json:"FinInstFullCode,omitempty"`
	FinInstName               *string `json:"FinInstName,omitempty"`
	FinInstBranchName         *string `json:"FinInstBranchName,omitempty"`
	SWIFTCode                 *string `json:"SWIFTCode,omitempty"`
	InternalFinInstCustomerID *int    `json:"InternalFinInstCustomerID,omitempty"`
	InternalFinInstAccountID  *int    `json:"InternalFinInstAccountID,omitempty"`
	FinInstControlKey         *string `json:"FinInstControlKey,omitempty"`
	FinInstAccountName        *string `json:"FinInstAccountName,omitempty"`
	FinInstAccount            *string `json:"FinInstAccount,omitempty"`
	HouseBank                 *string `json:"HouseBank,omitempty"`
	HouseBankAccount          *string `json:"HouseBankAccount,omitempty"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion,omitempty"`
}

type Relationship struct {
	BusinessPartner             int     `json:"BusinessPartner,omitempty"`
	RelationshipNumber          int     `json:"RelationshipNumber,omitempty"`
	ValidityEndDate             string  `json:"ValidityEndDate,omitempty"`
	ValidityStartDate           string  `json:"ValidityStartDate,omitempty"`
	RelationshipCategory        *string `json:"RelationshipCategory,omitempty"`
	RelationshipBusinessPartner *int    `json:"RelationshipBusinessPartner,omitempty"`
	BusinessPartnerPerson       *string `json:"BusinessPartnerPerson,omitempty"`
	IsStandardRelationship      *bool   `json:"IsStandardRelationship,omitempty"`
	IsMarkedForDeletion         *bool   `json:"IsMarkedForDeletion,omitempty"`
}

type BPAccounting struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	ChartOfAccounts     *string `json:"ChartOfAccounts"`
	FiscalYearVariant   *string `json:"FiscalYearVariant"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func CreateBusinessPartnerRes(msg rabbitmq.RabbitmqMessage) (*BusinessPartnerRes, error) {
	res := BusinessPartnerRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
