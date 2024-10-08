package apiOutputFormatter

type Authenticator struct {
	AuthenticatorUser						[]AuthenticatorUser						`json:"User"`
	AuthenticatorInitialEmailAuth			[]AuthenticatorInitialEmailAuth			`json:"InitialEmailAuth"`
	AuthenticatorInitialSMSAuth				[]AuthenticatorInitialSMSAuth			`json:"InitialSMSAuth"`
	AuthenticatorSMSAuth					[]AuthenticatorSMSAuth					`json:"SMSAuth"`
	AuthenticatorInitialGoogleAccountAuth	[]AuthenticatorInitialGoogleAccountAuth	`json:"InitialGoogleAccountAuth"`
	AuthenticatorGoogleAccountAuth			[]AuthenticatorGoogleAccountAuth		`json:"GoogleAccountAuth"`
	AuthenticatorInstagramAuth				[]AuthenticatorInstagramAuth			`json:"InstagramAuth"`
}

type AuthenticatorUser struct {
	UserID				string	`json:"UserID"`
	BusinessPartner		int		`json:"BusinessPartner"`
	Password			string	`json:"Password"`
	Language			string	`json:"Language"`
	LastLoginDate		*string	`json:"LastLoginDate"`
	LastLoginTime		*string	`json:"LastLoginTime"`
}

type AuthenticatorInitialEmailAuth struct {
	EmailAddress		string	`json:"EmailAddress"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
}

type AuthenticatorInitialSMSAuth struct {
	MobilePhoneNumber	string	`json:"MobilePhoneNumber"`
	AuthenticationCode	int		`json:"AuthenticationCode"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
}

type AuthenticatorSMSAuth struct {
	UserID				string	`json:"UserID"`
	MobilePhoneNumber	string	`json:"MobilePhoneNumber"`
	AuthenticationCode	int		`json:"AuthenticationCode"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
}

type AuthenticatorInitialGoogleAccountAuth struct {
	EmailAddress		string	`json:"EmailAddress"`
	GoogleID			string	`json:"GoogleID"`
	AccessToken			string	`json:"AccessToken"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type AuthenticatorGoogleAccountAuth struct {
	UserID				string	`json:"UserID"`
	EmailAddress		string	`json:"EmailAddress"`
	GoogleID			string	`json:"GoogleID"`
	AccessToken			string	`json:"AccessToken"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
}

type AuthenticatorInstagramAuth struct {
	UserID				string	`json:"UserID"`
	InstagramID			string	`json:"InstagramID"`
	AccessToken			string	`json:"AccessToken"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
