package apiModuleRuntimesResponsesAuthenticator

type AuthenticatorRes struct {
	Message Authenticator `json:"message,omitempty"`
}

type Authenticator struct {
	User						*[]User						`json:"User,omitempty"`
	InitialEmailAuth			*[]InitialEmailAuth			`json:"InitialEmailAuth,omitempty"`
	InitialSMSAuth				*[]InitialSMSAuth			`json:"InitialSMSAuth,omitempty"`
	SMSAuth						*[]SMSAuth					`json:"SMSAuth,omitempty"`
	InitialGoogleAccountAuth	*[]InitialGoogleAccountAuth	`json:"InitialGoogleAccountAuth,omitempty"`
	GoogleAccountAuth			*[]GoogleAccountAuth		`json:"GoogleAccountAuth,omitempty"`
	InstagramAuth				*[]InstagramAuth			`json:"InstagramAuth,omitempty"`
}

type User struct {
	UserID				string	`json:"UserID"`
	BusinessPartner		int		`json:"BusinessPartner"`
	Password			string	`json:"Password"`
	Qos					string	`json:"Qos"`
	IsEncrypt			bool	`json:"IsEncrypt"`
	Language			string	`json:"Language"`
	LastLoginDate		*string	`json:"LastLoginDate"`
	LastLoginTime		*string	`json:"LastLoginTime"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type InitialEmailAuth struct {
	EmailAddress		string	`json:"EmailAddress"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type InitialSMSAuth struct {
	MobilePhoneNumber	string	`json:"MobilePhoneNumber"`
	AuthenticationCode	int		`json:"AuthenticationCode"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type SMSAuth struct {
	UserID				string	`json:"UserID"`
	MobilePhoneNumber	string	`json:"MobilePhoneNumber"`
	AuthenticationCode	int		`json:"AuthenticationCode"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type InitialGoogleAccountAuth struct {
	EmailAddress		string	`json:"EmailAddress"`
	GoogleID			string	`json:"GoogleID"`
	AccessToken			string	`json:"AccessToken"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type GoogleAccountAuth struct {
	UserID				string	`json:"UserID"`
	EmailAddress		string	`json:"EmailAddress"`
	GoogleID			string	`json:"GoogleID"`
	AccessToken			string	`json:"AccessToken"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type InstagramAuth struct {
	UserID				string	`json:"UserID"`
	InstagramID			string	`json:"InstagramID"`
	AccessToken			string	`json:"AccessToken"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
