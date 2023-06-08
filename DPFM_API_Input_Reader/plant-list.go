package dpfm_api_input_reader

type PlantList struct {
	UIKeyGeneralUserID          string          `json:"ui_key_general_user_id"`
	UIKeyGeneralUserLanguage    string          `json:"ui_key_general_user_language"`
	UIKeyGeneralBusinessPartner string          `json:"ui_key_general_business_partner"`
	UIFunction                  string          `json:"ui_function"`
	UIKeyFunctionURL            string          `json:"ui_key_function_url"`
	RuntimeSessionID            string          `json:"runtime_session_id"`
	Params                      PlantListParams `json:"Params"`
	ReqReceiveQueue             *string         `json:"responseReceiveQueue"`
}

type PlantListParams struct {
	BusinessPartner int `json:"BusinessPartner"`
}
