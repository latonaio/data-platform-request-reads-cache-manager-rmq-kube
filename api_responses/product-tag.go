package apiresponses

import (
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type ProductTagRes struct {
	ConnectionKey       string            `json:"connection_key,omitempty"`
	Result              bool              `json:"result,omitempty"`
	RedisKey            string            `json:"redis_key,omitempty"`
	Filepath            string            `json:"filepath,omitempty"`
	APIStatusCode       int               `json:"api_status_code,omitempty"`
	RuntimeSessionID    string            `json:"runtime_session_id,omitempty"`
	BusinessPartnerID   *int              `json:"business_partner,omitempty"`
	ServiceLabel        string            `json:"service_label,omitempty"`
	APIType             string            `json:"api_type,omitempty"`
	Message             ProductTagMessage `json:"message,omitempty"`
	APISchema           string            `json:"api_schema,omitempty"`
	Accepter            []string          `json:"accepter,omitempty"`
	Deleted             bool              `json:"deleted,omitempty"`
	SQLUpdateResult     *bool             `json:"sql_update_result,omitempty"`
	SQLUpdateError      string            `json:"sql_update_error,omitempty"`
	SubfuncResult       *bool             `json:"subfunc_result,omitempty"`
	SubfuncError        string            `json:"subfunc_error,omitempty"`
	ExconfResult        *bool             `json:"exconf_result,omitempty"`
	ExconfError         string            `json:"exconf_error,omitempty"`
	APIProcessingResult *bool             `json:"api_processing_result,omitempty"`
	APIProcessingError  string            `json:"api_processing_error,omitempty"`
}

type ProductTagMessage struct {
	ProductTag *[]ProductTag `json:"ProductTag,omitempty"`
}

type ProductTagReads struct {
	ConnectionKey string `json:"connection_key,omitempty"`
	Result        bool   `json:"result,omitempty"`
	RedisKey      string `json:"redis_key,omitempty"`
	Filepath      string `json:"filepath,omitempty"`
	Product       string `json:"Product,omitempty"`
	APISchema     string `json:"api_schema,omitempty"`
	MaterialCode  string `json:"material_code,omitempty"`
	Deleted       string `json:"deleted,omitempty"`
}

type ProductTag struct {
	Key      string `json:"key,omitempty"`
	DocCount int    `json:"doc_count,omitempty"`
}

func CreateProductTagRes(msg rabbitmq.RabbitmqMessage) (*ProductTagRes, error) {
	res := ProductTagRes{}
	err := json.Unmarshal(msg.Raw(), &res)
	if err != nil {
		return nil, xerrors.Errorf("unmarshal error: %w", err)
	}
	return &res, nil
}
