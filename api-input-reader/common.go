package apiInputReader

type Request struct {
	BusinessPartner     *int    `json:"BusinessPartner"`
	UserID              *string `json:"UserId"`
	User                *string `json:"User"`
	Language            *string `json:"Language"`
	UserType            *string `json:"UserType"`
	IsCancelled         *bool   `json:"IsCancelled"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	RuntimeSessionID    *string `json:"RuntimeSessionId"`
}

type RedisCacheApiName map[string]map[string]interface{}
