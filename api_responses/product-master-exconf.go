package apiresponses

type ProductMasterExconfRes struct {
	ConnectionKey     string              `json:"connection_key"`
	Result            bool                `json:"result"`
	RedisKey          string              `json:"redis_key"`
	Filepath          string              `json:"filepath"`
	APIStatusCode     int                 `json:"api_status_code"`
	RuntimeSessionID  string              `json:"runtime_session_id"`
	BusinessPartnerID *int                `json:"business_partner"`
	ServiceLabel      string              `json:"service_label"`
	ProductMaster     ProductMasterExconf `json:"ProductMaster"`
	APISchema         string              `json:"api_schema"`
	Accepter          []string            `json:"accepter"`
	Deleted           bool                `json:"deleted"`
}

type ProductMasterExconf struct {
	General         *ProductMasterExconfGeneral `json:"General"`
	BusinessPartner *BusinessPartnerExconf      `json:"BusinessPartner"`
	BPPlant         *BPPlantExconf              `json:"BPPlant"`
	StorageLocation *StorageLocationExconf      `json:"StorageLocation"`
	MRPArea         *MRPAreaExconf              `json:"MRPArea"`
	WorkScheduling  *WorkSchedulingExconf       `json:"WorkScheduling"`
	Accounting      *AccountingExconf           `json:"Accounting"`
}

type ProductMasterExconfGeneral struct {
	Product       string `json:"Product"`
	ExistenceConf bool   `json:"ExistenceConf"`
}

type BusinessPartnerExconf struct {
	Product           string `json:"Product"`
	BusinessPartner   int    `json:"BusinessPartner"`
	ValidityEndDate   string `json:"ValidityEndDate"`
	ValidityStartDate string `json:"ValidityStartDate"`
	ExistenceConf     bool   `json:"ExistenceConf"`
}

type BPPlantExconf struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}

type StorageLocationExconf struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	StorageLocation string `json:"StorageLocation"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}

type MRPAreaExconf struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	MRPArea         string `json:"MRPArea"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}

type WorkSchedulingExconf struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}

type AccountingExconf struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	ExistenceConf   bool   `json:"ExistenceConf"`
}
