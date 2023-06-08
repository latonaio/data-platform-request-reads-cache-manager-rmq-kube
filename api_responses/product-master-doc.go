package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type ProductMasterDocRes struct {
	ConnectionKey     string             `json:"connection_key,omitempty"`
	Result            bool               `json:"result,omitempty"`
	RedisKey          string             `json:"redis_key,omitempty"`
	Filepath          string             `json:"filepath,omitempty"`
	APIStatusCode     int                `json:"api_status_code,omitempty"`
	RuntimeSessionID  string             `json:"runtime_session_id,omitempty"`
	BusinessPartnerID *int               `json:"business_partner,omitempty"`
	ServiceLabel      string             `json:"service_label,omitempty"`
	APIType           string             `json:"api_type,omitempty"`
	Message           *ResponseHeaderDoc `json:"message,omitempty"`
	APISchema         string             `json:"api_schema,omitempty"`
	Accepter          []string           `json:"accepter,omitempty"`
	Deleted           bool               `json:"deleted,omitempty"`
}

type ResponseHeaderDoc struct {
	HeaderDoc *[]PMDHeaderDoc `json:"HeaderDoc"`
}

type PMDHeaderDoc struct {
	Product                  string `json:"Product"`
	DocType                  string `json:"DocType"`
	FileExtension            string `json:"FileExtension"`
	DocVersionID             int    `json:"DocVersionID"`
	DocID                    string `json:"DocID"`
	DocIssuerBusinessPartner int    `json:"DocIssuerBusinessPartner"`
	FilePath                 string `json:"FilePath"`
	FileName                 string `json:"FileName"`
}

func CreateProductMasterDocRes(msg rabbitmq.RabbitmqMessage) (*ProductMasterDocRes, error) {
	res := ProductMasterDocRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
