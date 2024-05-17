package apiOutputFormatter

type Authenticator struct {
	AuthenticatorUser    []AuthenticatorUser    `json:"User"`
}

type AuthenticatorUser struct {
	UserID				string	`json:"UserID"`
	BusinessPartner		int		`json:"BusinessPartner"`
	Password			string	`json:"Password"`
	Language			string	`json:"Language"`
	LastLoginDate		*string	`json:"LastLoginDate"`
	LastLoginTime		*string	`json:"LastLoginTime"`
}
